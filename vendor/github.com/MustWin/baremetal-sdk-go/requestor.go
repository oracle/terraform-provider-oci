// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

import (
	"bytes"
	"encoding/json"
	"net/http"
	//"fmt"
	//"net/http/httputil"
)

type requestor interface {
	request(method string, reqOpts request) (r *response, e error)
	getRequest(reqOpts request) (resp *response, e error)
	deleteRequest(reqOpts request) (e error)
}

type apiRequestor struct {
	httpClient *http.Client
	authInfo   *authenticationInfo
	urlBuilder urlBuilderFn
	userAgent  string
}

func newCoreAPIRequestor(authInfo *authenticationInfo, nco *NewClientOptions) (r *apiRequestor) {
	return &apiRequestor{
		httpClient: &http.Client{
			Transport: nco.Transport,
		},
		authInfo:   authInfo,
		urlBuilder: buildCoreURL,
		userAgent:  nco.UserAgent,
	}
}

func newObjectStorageAPIRequestor(authInfo *authenticationInfo, nco *NewClientOptions) (r *apiRequestor) {
	return &apiRequestor{
		httpClient: &http.Client{
			Transport: nco.Transport,
		},
		authInfo:   authInfo,
		urlBuilder: buildObjectStorageURL,
		userAgent:  nco.UserAgent,
	}
}

func newDatabaseAPIRequestor(authInfo *authenticationInfo, nco *NewClientOptions) (r *apiRequestor) {
	return &apiRequestor{
		httpClient: &http.Client{
			Transport: nco.Transport,
		},
		authInfo:   authInfo,
		urlBuilder: buildDatabaseURL,
		userAgent:  nco.UserAgent,
	}
}

func newIdentityAPIRequestor(authInfo *authenticationInfo, nco *NewClientOptions) (r *apiRequestor) {
	return &apiRequestor{
		httpClient: &http.Client{
			Transport: nco.Transport,
		},
		authInfo:   authInfo,
		urlBuilder: buildIdentityURL,
		userAgent:  nco.UserAgent,
	}
}

func (api *apiRequestor) deleteRequest(reqOpts request) (e error) {
	var url string
	if url, e = reqOpts.marshalURL(api.urlBuilder); e != nil {
		return
	}

	var req *http.Request
	if req, e = http.NewRequest(http.MethodDelete, url, nil); e != nil {
		return
	}

	req.Header = reqOpts.marshalHeader()

	if e = createAuthorizationHeader(req, api.authInfo, api.userAgent, []byte{}); e != nil {
		return
	}

	var resp *http.Response
	if resp, e = api.httpClient.Do(req); e != nil {
		return
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
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

func (api *apiRequestor) getRequest(reqOpts request) (getResp *response, e error) {
	var url string
	if url, e = reqOpts.marshalURL(api.urlBuilder); e != nil {
		return
	}

	var req *http.Request
	if req, e = http.NewRequest(http.MethodGet, url, nil); e != nil {
		return
	}

	req.Header = reqOpts.marshalHeader()

	if e = createAuthorizationHeader(req, api.authInfo, api.userAgent, []byte{}); e != nil {
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

	getResp = &response{
		header: resp.Header,
		body:   reader.Bytes(),
	}

	return
}

func (api *apiRequestor) request(method string, reqOpts request) (r *response, e error) {
	var jsonBuffer []byte
	if jsonBuffer, e = reqOpts.marshalBody(); e != nil {
		return
	}

	buffer := bytes.NewBuffer(jsonBuffer)

	var url string
	if url, e = reqOpts.marshalURL(api.urlBuilder); e != nil {
		return
	}

	var req *http.Request
	if req, e = http.NewRequest(method, url, buffer); e != nil {
		return
	}
	req.Header = reqOpts.marshalHeader()

	if e = createAuthorizationHeader(req, api.authInfo, api.userAgent, jsonBuffer); e != nil {
		return
	}

	// fmt.Println("url")
	//rbody, _ := httputil.DumpRequest(req, true)
	//fmt.Println(string(rbody))
	var resp *http.Response
	if resp, e = api.httpClient.Do(req); e != nil {
		return
	}
	//rbody, _ = httputil.DumpResponse(resp, true)
	//fmt.Println(string(rbody))

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

	r = &response{
		header: resp.Header,
		body:   reader.Bytes(),
	}

	return
}
