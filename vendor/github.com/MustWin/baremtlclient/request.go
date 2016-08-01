package baremtlsdk

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type authenticationInfo struct {
	privateRSAKey  *rsa.PrivateKey
	tenancyOCID    string
	userOCID       string
	keyFingerPrint string
}

type authenticateRequest struct {
	*authenticationInfo
	*http.Request
}

type requestResponse struct {
	header http.Header
	body   []byte
}
type requestor interface {
	request(method, url string, body interface{}, headers http.Header) (r *requestResponse, e error)
	getRequest(urlStr string, headers http.Header) (resp *requestResponse, e error)
	deleteRequest(urlStr string, headers http.Header) (e error)
}

type apiRequestor struct {
	httpClient *http.Client
	authInfo   *authenticationInfo
}

func newAPIRequestor(authInfo *authenticationInfo, tr *http.Transport) (r *apiRequestor) {

	return &apiRequestor{
		httpClient: &http.Client{
			Transport: tr,
		},
		authInfo: authInfo,
	}
}

func (api *apiRequestor) deleteRequest(urlStr string, headers http.Header) (e error) {
	var req *http.Request
	if req, e = http.NewRequest(http.MethodDelete, urlStr, nil); e != nil {
		return
	}

	if headers != nil {
		req.Header = headers
	}

	if e = createAuthorizationHeader(req, api.authInfo, []byte{}); e != nil {
		return
	}

	var resp *http.Response
	if resp, e = api.httpClient.Do(req); e != nil {
		return
	}

	if resp.StatusCode != http.StatusOK {
		decoder := json.NewDecoder(resp.Body)
		var err Error

		if e = decoder.Decode(&err); e != nil {
			return e
		}

		err.OPCRequestID = resp.Header.Get(headerOPCRequestID)
		return &err
	}

	return
}

func (api *apiRequestor) getRequest(urlStr string, headers http.Header) (getResp *requestResponse, e error) {

	var req *http.Request
	if req, e = http.NewRequest(http.MethodGet, urlStr, nil); e != nil {
		return
	}

	if headers != nil {
		req.Header = headers
	}

	if e = createAuthorizationHeader(req, api.authInfo, []byte{}); e != nil {
		return
	}

	var resp *http.Response
	if resp, e = api.httpClient.Do(req); e != nil {
		return
	}

	var reader bytes.Buffer
	_, e = reader.ReadFrom(resp.Body)
	resp.Body.Close()
	if e != nil {
		return
	}

	if resp.StatusCode != http.StatusOK {
		e = getErrorFromResponse(&reader, resp)
		return
	}

	getResp = &requestResponse{
		header: resp.Header,
		body:   reader.Bytes(),
	}

	return
}

func (api *apiRequestor) request(method, urlStr string, body interface{}, headers http.Header) (r *requestResponse, e error) {
	var jsonBuffer []byte
	if jsonBuffer, e = json.Marshal(body); e != nil {
		return
	}

	buffer := bytes.NewBuffer(jsonBuffer)

	var req *http.Request
	if req, e = http.NewRequest(method, urlStr, buffer); e != nil {
		return
	}

	if headers != nil {
		req.Header = headers
	}

	if e = createAuthorizationHeader(req, api.authInfo, jsonBuffer); e != nil {
		return
	}

	var resp *http.Response
	if resp, e = api.httpClient.Do(req); e != nil {
		return
	}

	var reader bytes.Buffer
	_, e = reader.ReadFrom(resp.Body)
	resp.Body.Close()

	if e != nil {
		return
	}

	// we still have to check response code, if we get non 200 response
	// body will contain an error object which we'll Unmarshal and send
	// back as an Error
	if resp.StatusCode != http.StatusOK {
		e = getErrorFromResponse(&reader, resp)
		return
	}

	r = &requestResponse{
		header: resp.Header,
		body:   reader.Bytes(),
	}

	return
}

func getErrorFromResponse(body io.Reader, resp *http.Response) (e error) {
	var apiError Error
	decoder := json.NewDecoder(body)
	if e = decoder.Decode(&apiError); e != nil {
		return
	}

	if opcRequestID := resp.Header.Get(headerOPCRequestID); opcRequestID != "" {
		apiError.OPCRequestID = opcRequestID
	}

	return &apiError
}

func (a *authenticationInfo) getKeyID() string {
	return fmt.Sprintf("%s/%s/%s", a.tenancyOCID, a.userOCID, a.keyFingerPrint)
}

func createAuthorizationHeader(request *http.Request, auth *authenticationInfo, body []byte) (e error) {
	addRequiredRequestHeaders(request, body)
	var sig string

	if sig, e = computeSignature(request, auth.privateRSAKey); e != nil {
		return
	}

	signedHeaders := getSigningHeaders(request.Method)
	headers := concantenateHeaders(signedHeaders)

	authValue := fmt.Sprintf("Signature version=\"1\" signature=\"%s\",headers=\"%s\","+
		"algorithm=\"rsa-sha256\",keyId=\"%s\"", sig, headers, auth.getKeyID())

	request.Header.Add("authorization", authValue)

	return
}

func concantenateHeaders(headers []string) (concatenated string) {

	for _, header := range headers {
		if len(concatenated) > 0 {
			concatenated += " "
		}
		concatenated += header
	}

	return
}

func getSigningHeaders(method string) []string {
	result := []string{
		"date",
		"(request-target)",
	}

	if method == http.MethodPost || method == http.MethodPut {
		result = append(result, "content-length", "content-type", "x-content-sha256")
	}

	return result
}

func computeSignature(request *http.Request, privateKey *rsa.PrivateKey) (sig string, e error) {
	signingString := getSigningString(request)
	hasher := sha256.New()
	hasher.Write([]byte(signingString))
	hashed := hasher.Sum(nil)
	var unencodedSig []byte
	unencodedSig, e = rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
	if e != nil {
		return
	}

	sig = base64.StdEncoding.EncodeToString(unencodedSig)

	return

}

func getSigningString(request *http.Request) string {
	signingHeaders := getSigningHeaders(request.Method)
	signingString := ""
	for _, header := range signingHeaders {
		if signingString != "" {
			signingString += "\n"
		}

		if header == "(request-target)" {
			signingString += fmt.Sprintf("%s: %s", header, getRequestTarget(request))
		} else {
			signingString += fmt.Sprintf("%s: %s", header, request.Header.Get(header))
		}
	}

	return signingString

}

func getRequestTarget(request *http.Request) string {
	lowercaseMethod := strings.ToLower(request.Method)
	return fmt.Sprintf("%s %s", lowercaseMethod, request.URL.RequestURI())
}

func addIfNotPresent(dest *http.Header, key, value string) {
	if dest.Get(key) == "" {
		dest.Set(key, value)
	}
}

func getBodyHash(body []byte) string {
	hash := sha256.Sum256(body)
	return base64.StdEncoding.EncodeToString(hash[:])
}

func addRequiredRequestHeaders(request *http.Request, body []byte) {
	addIfNotPresent(&request.Header, "content-type", "application/json")
	addIfNotPresent(&request.Header, "date", time.Now().UTC().Format(http.TimeFormat))
	addIfNotPresent(&request.Header, "accept", "*/*")

	if request.Method == http.MethodPost || request.Method == http.MethodPut {
		addIfNotPresent(&request.Header, "content-length", strconv.FormatInt(request.ContentLength, 10))

		if request.ContentLength > 0 {
			addIfNotPresent(&request.Header, "x-content-sha256", getBodyHash(body))
		}

	}
}
