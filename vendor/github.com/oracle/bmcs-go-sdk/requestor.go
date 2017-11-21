// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
	"time"
)

type requestor interface {
	request(method string, reqOpts request) (r *response, e error)
	getRequest(reqOpts request) (resp *response, e error)
	postRequest(reqOpts request) (resp *response, e error)
	deleteRequest(reqOpts request) (e error)
}

type apiRequestor struct {
	httpClient             *http.Client
	authInfo               *authenticationInfo
	urlBuilder             urlBuilderFn
	urlTemplate            string
	userAgent              string
	region                 string
	shortRetryTime         time.Duration
	longRetryTime          time.Duration
	randGen                *rand.Rand
	disableAutoRetries     bool
	disableNotFoundRetries bool
}

func newCoreAPIRequestor(authInfo *authenticationInfo, nco *NewClientOptions) (r *apiRequestor) {
	return &apiRequestor{
		httpClient: &http.Client{
			Transport: nco.Transport,
		},
		authInfo:               authInfo,
		urlBuilder:             buildCoreURL,
		urlTemplate:            nco.UrlTemplate,
		userAgent:              nco.UserAgent,
		region:                 nco.Region,
		shortRetryTime:         nco.ShortRetryTime,
		longRetryTime:          nco.LongRetryTime,
		randGen:                nco.RandGen,
		disableAutoRetries:     nco.DisableAutoRetries,
		disableNotFoundRetries: nco.DisableNotFoundRetries,
	}
}

func newObjectStorageAPIRequestor(authInfo *authenticationInfo, nco *NewClientOptions) (r *apiRequestor) {
	return &apiRequestor{
		httpClient: &http.Client{
			Transport: nco.Transport,
		},
		authInfo:               authInfo,
		urlBuilder:             buildObjectStorageURL,
		urlTemplate:            nco.UrlTemplate,
		userAgent:              nco.UserAgent,
		region:                 nco.Region,
		shortRetryTime:         nco.ShortRetryTime,
		longRetryTime:          nco.LongRetryTime,
		randGen:                nco.RandGen,
		disableAutoRetries:     nco.DisableAutoRetries,
		disableNotFoundRetries: nco.DisableNotFoundRetries,
	}
}

func newDatabaseAPIRequestor(authInfo *authenticationInfo, nco *NewClientOptions) (r *apiRequestor) {
	return &apiRequestor{
		httpClient: &http.Client{
			Transport: nco.Transport,
		},
		authInfo:               authInfo,
		urlBuilder:             buildDatabaseURL,
		urlTemplate:            nco.UrlTemplate,
		userAgent:              nco.UserAgent,
		region:                 nco.Region,
		shortRetryTime:         nco.ShortRetryTime,
		longRetryTime:          nco.LongRetryTime,
		randGen:                nco.RandGen,
		disableAutoRetries:     nco.DisableAutoRetries,
		disableNotFoundRetries: nco.DisableNotFoundRetries,
	}
}

func newIdentityAPIRequestor(authInfo *authenticationInfo, nco *NewClientOptions) (r *apiRequestor) {
	return &apiRequestor{
		httpClient: &http.Client{
			Transport: nco.Transport,
		},
		authInfo:               authInfo,
		urlBuilder:             buildIdentityURL,
		urlTemplate:            nco.UrlTemplate,
		userAgent:              nco.UserAgent,
		region:                 nco.Region,
		shortRetryTime:         nco.ShortRetryTime,
		longRetryTime:          nco.LongRetryTime,
		randGen:                nco.RandGen,
		disableAutoRetries:     nco.DisableAutoRetries,
		disableNotFoundRetries: nco.DisableNotFoundRetries,
	}
}

func newLoadBalancerAPIRequestor(authInfo *authenticationInfo, nco *NewClientOptions) (r *apiRequestor) {
	return &apiRequestor{
		httpClient: &http.Client{
			Transport: nco.Transport,
		},
		authInfo:               authInfo,
		urlBuilder:             buildLoadBalancerURL,
		urlTemplate:            nco.UrlTemplate,
		userAgent:              nco.UserAgent,
		region:                 nco.Region,
		shortRetryTime:         nco.ShortRetryTime,
		longRetryTime:          nco.LongRetryTime,
		randGen:                nco.RandGen,
		disableAutoRetries:     nco.DisableAutoRetries,
		disableNotFoundRetries: nco.DisableNotFoundRetries,
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

func (api *apiRequestor) postRequest(reqOpts request) (postResp *response, e error) {
	if postResp, e = api.request(http.MethodPost, reqOpts); e != nil {
		return
	}
	return
}

func (api *apiRequestor) request(method string, reqOpts request) (r *response, e error) {
	return submitRequestWithRetries(api, method, reqOpts, generateRetryToken(api.randGen),
		"", -1, 0, 1)
}

func submitRequestWithRetries(api *apiRequestor, method string, reqOpts request, generatedRetryToken string,
	currentErrorCode string, retryTimeRemaining time.Duration, timeWaited time.Duration, retryNum uint) (r *response, e error) {
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
	if url, e = reqOpts.marshalURL(api.urlTemplate, api.region, api.urlBuilder); e != nil {
		return
	}

	var req *http.Request
	if req, e = http.NewRequest(method, url, buffer); e != nil {
		return
	}
	req.Header = reqOpts.marshalHeader()

	//add random retry token if user hasn't added one so that we can safely retry requests
	if _, present := req.Header[retryTokenKey]; !api.disableAutoRetries &&
		!present &&
		method != http.MethodDelete &&
		method != http.MethodGet {
		req.Header[retryTokenKey] = []string{generatedRetryToken}
	}

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

		apiError := getErrorFromResponse(&reader, resp)
		if api.disableAutoRetries {
			e = &apiError
			return
		}
		errorCodeStr := fmt.Sprintf("%s:%s", apiError.Status, apiError.Code)
		if retryNum == 1 {
			retryTimeRemaining = getMaxRetryTimeInSeconds(api, apiError, req.URL.String(), method, api.disableNotFoundRetries)
			currentErrorCode = errorCodeStr
		} else if currentErrorCode != errorCodeStr {
			retryTimeRemaining = getMaxRetryTimeInSeconds(api, apiError, req.URL.String(), method, api.disableNotFoundRetries) - timeWaited
			currentErrorCode = errorCodeStr
		}
		if retryTimeRemaining > 0 {
			timeSlept := polynomialBackoffSleep(retryNum, retryTimeRemaining)
			return submitRequestWithRetries(api, method, reqOpts, generatedRetryToken,
				currentErrorCode, retryTimeRemaining-timeSlept, timeWaited+timeSlept, retryNum+1)
		} else {
			e = &apiError
			return
		}
	}

	r = &response{
		header: resp.Header,
		body:   reader.Bytes(),
	}

	return
}

var sleep = time.Sleep

func polynomialBackoffSleep(retryNum uint, retryTimeRemaining time.Duration) time.Duration {
	secondsToSleep := time.Duration(retryNum*retryNum) * time.Second
	if retryTimeRemaining < secondsToSleep {
		secondsToSleep = retryTimeRemaining
	}
	if os.Getenv("DEBUG") != "" {
		log.Printf("[DEBUG] Got a retriable error. Waiting %d seconds and trying again...", int(secondsToSleep.Seconds()))
	}
	if os.Getenv("TEST") != "true" {
		sleep(secondsToSleep)
	}
	return secondsToSleep
}

func getMaxRetryTimeInSeconds(api *apiRequestor, e Error, requestURL string, method string, disableNotFoundRetries bool) time.Duration {
	switch e.Status {
	case "400":
		return 0
	case "401":
		return 0
	case "403":
		return 0
	case "404":
		if disableNotFoundRetries {
			return 0
		}
		if method == http.MethodDelete {
			return 0
		}
		if requestServiceCheck(requestURL, identityServiceAPI) ||
			requestServiceCheck(requestURL, objectStorageServiceAPI) {
			return api.longRetryTime
		}
	case "409":
		if e.Code == "InvalidatedRetryToken" || e.Code == "CompartmentAlreadyExists" {
			return 0
		} else if e.Code == "NotAuthorizedOrResourceAlreadyExists" {
			if requestServiceCheck(requestURL, identityServiceAPI) ||
				requestServiceCheck(requestURL, objectStorageServiceAPI) {
				return api.longRetryTime
			}
		}
	case "412":
		return 0
	case "429":
		return api.longRetryTime
	case "500":
		if requestServiceCheck(requestURL, objectStorageServiceAPI) {
			return api.longRetryTime
		}
	}
	return api.shortRetryTime
}

func requestServiceCheck(requestURL string, service string) bool {
	if service == "" {
		return false
	}
	return strings.HasPrefix(requestURL, urlPrefix+service)
}

/* Generates a random alphanumeric string.
 * Used for generating a retry token so that the SDK can safely retry operations.
 */
func generateRetryToken(randGen *rand.Rand) string {
	alphanumericChars := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	retryToken := make([]rune, generatedRetryTokenLength)
	for i := range retryToken {
		retryToken[i] = alphanumericChars[randGen.Intn(len(alphanumericChars))]
	}
	return string(retryToken)
}
