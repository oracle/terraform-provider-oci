package baremetal

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type urlParts []interface{}

type requestOptions interface {
	header() http.Header
	getBody() interface{}
	url(urlBuilderFn) string
}

type sdkRequestOptions struct {
	body       interface{}
	name       resourceName
	options    []Options
	ocid       string
	ids        urlParts
	query      url.Values
	httpHeader http.Header
}

func (r *sdkRequestOptions) url(b urlBuilderFn) string {
	r.parseOptions()
	return b(r.name, r.query, r.ids...)
}

func (r *sdkRequestOptions) getBody() interface{} {
	r.parseOptions()
	return r.body
}

func (r *sdkRequestOptions) header() http.Header {
	r.parseOptions()
	return r.httpHeader
}

func (r *sdkRequestOptions) parseOptions() {
	if r.query == nil {
		r.query = url.Values{}
	}

	r.httpHeader = http.Header{}

	// Parse query options.
	if r.ocid != "" {
		r.query.Set(queryCompartmentID, r.ocid)
	}

	if len(r.options) > 0 {
		option := r.options[0]

		// Parse header options.
		if option.OPCIdempotencyToken != "" {
			r.httpHeader.Set(headerOPCIdempotencyToken, option.OPCIdempotencyToken)
		}
		if option.IfMatch != "" {
			r.httpHeader.Set(headerIfMatch, option.IfMatch)
		}

		if option.AvailabilityDomain != "" {
			r.query.Set(queryAvailabilityDomain, option.AvailabilityDomain)
		}
		if option.ImageID != "" {
			r.query.Set(queryImageID, option.ImageID)
		}
		if option.InstanceID != "" {
			r.query.Set(queryInstanceID, option.InstanceID)
		}
		if option.Limit > 0 {
			r.query.Set(queryLimit, strconv.FormatUint(option.Limit, 10))
		}
		if option.Page != "" {
			r.query.Set(queryPage, option.Page)
		}
		if option.VnicID != "" {
			r.query.Set(queryVnicID, option.VnicID)
		}
	}
}

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

func (r *requestResponse) unmarshal(resource interface{}) (e error) {
	var val interface{}
	if c, ok := resource.(Container); ok {
		val = c.GetList()
		c.SetNextPage(r.header.Get(headerOPCNextPage))
	} else {
		val = resource
	}

	if e = json.Unmarshal(r.body, val); e != nil {
		return
	}

	if rr, ok := resource.(Requestable); ok {
		rr.SetRequestID(r.header.Get(headerOPCRequestID))
	}
	if et, ok := resource.(ETagged); ok {
		et.SetETag(r.header.Get(headerETag))
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

	authValue := fmt.Sprintf("Signature headers=\"%s\",keyId=\"%s\",algorithm=\"rsa-sha256\",signature=\"%s\"", headers, auth.getKeyID(), sig)

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
