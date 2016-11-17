package baremetal

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type requestor interface {
	request(method string, reqOpts requestOptions) (r *requestResponse, e error)
	getRequest(reqOpts requestOptions) (resp *requestResponse, e error)
	deleteRequest(reqOpts requestOptions) (e error)
}

type apiRequestor struct {
	httpClient *http.Client
	authInfo   *authenticationInfo
	urlBuilder urlBuilderFn
}

func newCoreAPIRequestor(authInfo *authenticationInfo, tr *http.Transport) (r *apiRequestor) {
	return &apiRequestor{
		httpClient: &http.Client{
			Transport: tr,
		},
		authInfo:   authInfo,
		urlBuilder: buildCoreURL,
	}
}

func newIdentityAPIRequestor(authInfo *authenticationInfo, tr *http.Transport) (r *apiRequestor) {
	return &apiRequestor{
		httpClient: &http.Client{
			Transport: tr,
		},
		authInfo:   authInfo,
		urlBuilder: buildIdentityURL,
	}
}

func (api *apiRequestor) deleteRequest(reqOpts requestOptions) (e error) {
	var url string
	if url, e = reqOpts.url(api.urlBuilder); e != nil {
		return
	}

	var req *http.Request
	if req, e = http.NewRequest(http.MethodDelete, url, nil); e != nil {
		return
	}

	req.Header = reqOpts.header()

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

func (api *apiRequestor) getRequest(reqOpts requestOptions) (getResp *requestResponse, e error) {
	var url string
	if url, e = reqOpts.url(api.urlBuilder); e != nil {
		return
	}

	var req *http.Request
	if req, e = http.NewRequest(http.MethodGet, url, nil); e != nil {
		return
	}

	req.Header = reqOpts.header()

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

func (api *apiRequestor) request(method string, reqOpts requestOptions) (r *requestResponse, e error) {
	var jsonBuffer []byte
	if jsonBuffer, e = reqOpts.getBody(); e != nil {
		return
	}

	buffer := bytes.NewBuffer(jsonBuffer)

	var url string
	if url, e = reqOpts.url(api.urlBuilder); e != nil {
		return
	}

	var req *http.Request
	if req, e = http.NewRequest(method, url, buffer); e != nil {
		return
	}
	req.Header = reqOpts.header()

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
