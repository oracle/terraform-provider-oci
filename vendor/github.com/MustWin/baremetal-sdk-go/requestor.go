// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
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
	region     string
}

func newCoreAPIRequestor(authInfo *authenticationInfo, nco *NewClientOptions) (r *apiRequestor) {
	return &apiRequestor{
		httpClient: &http.Client{
			Transport: nco.Transport,
		},
		authInfo:   authInfo,
		urlBuilder: buildCoreURL,
		userAgent:  nco.UserAgent,
		region:     nco.Region,
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
		region:     nco.Region,
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
		region:     nco.Region,
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
		region:     nco.Region,
	}
}

func newLoadBalancerAPIRequestor(authInfo *authenticationInfo, nco *NewClientOptions) (r *apiRequestor) {
	return &apiRequestor{
		httpClient: &http.Client{
			Transport: nco.Transport,
		},
		authInfo:   authInfo,
		urlBuilder: buildLoadBalancerURL,
		userAgent:  nco.UserAgent,
		region:     nco.Region,
	}
}

func (api *apiRequestor) deleteRequest(reqOpts request) (e error) {
	_, e = api.request(http.MethodDelete, reqOpts)
	return
}

func (api *apiRequestor) getRequest(reqOpts request) (getResp *response, e error) {
	if getResp, e = api.request(http.MethodGet, reqOpts); e != nil {
		return
	}
	return
}

func (api *apiRequestor) request(method string, reqOpts request) (r *response, e error) {
	var jsonBuffer []byte
	var buffer *bytes.Buffer
	if method == http.MethodDelete || method == http.MethodGet {
		buffer = bytes.NewBuffer([]byte{})
	} else {
		if jsonBuffer, e = reqOpts.marshalBody(); e != nil {
			return
		}
		buffer = bytes.NewBuffer(jsonBuffer)
	}

	var url string
	if url, e = reqOpts.marshalURL(api.region, api.urlBuilder); e != nil {
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
	if e != nil {
		log.Printf("[WARN] Could not get HTTP authorization header, error: %#v\n", e)
		return
	}

	if os.Getenv("DEBUG") != "" {
		reqdump, err := httputil.DumpRequestOut(req, true)
		if err == nil {
			log.Printf("[DEBUG] HTTP Request: %v\n", string(reqdump))
		} else {
			log.Printf("[WARN] Could not dump HTTP Request: %#v, error: %#v\n", req, err)
		}
	}

	var resp *http.Response
	resp, e = api.httpClient.Do(req)
	if e != nil {
		log.Printf("[WARN] Could not get HTTP Response, error: %#v\n", e)
		return
	}

	if os.Getenv("DEBUG") != "" {
		respdump, err := httputil.DumpResponse(resp, true)
		if err == nil {
			log.Printf("[DEBUG] HTTP Response: %v\n", string(respdump))
		} else {
			log.Printf("[WARN] Could not dump HTTP Response: %#v, error: %#v\n", resp, err)
		}
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
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		e = getErrorFromResponse(&reader, resp)
		return
	}

	r = &response{
		header: resp.Header,
		body:   reader.Bytes(),
	}

	return
}
