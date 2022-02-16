// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package tfresource

import (
	"fmt"
	"net"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/globalvar"

	"github.com/stretchr/testify/assert"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"

	"github.com/oracle/oci-go-sdk/v58/common"
)

type TestOCIResponse struct {
	statusCode int
	header     map[string][]string
}

type retryTestInput struct {
	serviceName              string
	disableNotFoundRetries   bool
	httpResponseStatusCode   int
	header                   map[string][]string
	expectedRetryTimeSeconds int
	responseError            error
	jitterMode               bool
	optionals                []interface{}
}

func (response TestOCIResponse) HTTPResponse() *http.Response {
	result := http.Response{}
	result.Request = &http.Request{}
	result.StatusCode = response.statusCode
	result.Header = http.Header(response.header)
	return &result
}

func retryLoop(t *testing.T, r *retryTestInput) {
	retryPolicy := GetRetryPolicy(r.disableNotFoundRetries, r.serviceName, r.optionals...)
	startTime := time.Now()

	for i := uint(1); true; i++ {
		operationResponse := common.NewOCIOperationResponse(TestOCIResponse{statusCode: r.httpResponseStatusCode, header: r.header}, r.responseError, i)
		expectedShouldRetry := GetElapsedRetryDuration(startTime) < (time.Duration(r.expectedRetryTimeSeconds) * time.Second)
		actualShouldRetry := retryPolicy.ShouldRetryOperation(operationResponse)
		if actualShouldRetry != expectedShouldRetry {
			t.Errorf("Expected shouldRetry to return %v for attempt %v", expectedShouldRetry, i)
			return
		}

		if !actualShouldRetry {
			fmt.Println("Timeout exceeded; no retry.")
			return
		}

		waitTime := retryPolicy.NextDuration(operationResponse)
		fmt.Printf("Attempt #%v: Will wait for %v\n", i, waitTime.Round(time.Second))

		if r.jitterMode {
			expectedWaitTimeMax := time.Duration(2*i*i) * time.Second
			if i > quadraticBackoffCap {
				expectedWaitTimeMax = time.Duration(2*quadraticBackoffCap*quadraticBackoffCap) * time.Second
			}
			if waitTime >= expectedWaitTimeMax || waitTime < minRetryBackoff {
				t.Errorf("Expected wait time to be between %v and %v for attempt %v, but got %v", minRetryBackoff, expectedWaitTimeMax, i, waitTime)
				return
			}
		} else {

		}

		time.Sleep(waitTime)
	}
}

func TestNetTimeoutError(t *testing.T) {
	errNet := net.DNSError{
		Err:       "Timeout",
		IsTimeout: true,
	}
	assert.Equal(t, common.IsNetworkError(&errNet), true)

	if httpreplay.ModeRecordReplay() {
		t.Skip("Skip Retry Tests in HttpReplay mode.")
	}
	ShortRetryTime = 15 * time.Second
	LongRetryTime = 30 * time.Second
	ConfiguredRetryDuration = nil
	r := retryTestInput{
		serviceName:              "core",
		httpResponseStatusCode:   404,
		header:                   map[string][]string{},
		responseError:            &errNet,
		expectedRetryTimeSeconds: 15,
		jitterMode:               true,
	}
	retryLoop(t, &r)
}

func TestNetTemporaryError(t *testing.T) {
	errNet := net.DNSError{
		Err:         "Temporary",
		IsTemporary: true,
	}
	assert.Equal(t, common.IsNetworkError(&errNet), true)

	if httpreplay.ModeRecordReplay() {
		t.Skip("Skip Retry Tests in HttpReplay mode.")
	}
	ShortRetryTime = 15 * time.Second
	LongRetryTime = 30 * time.Second
	ConfiguredRetryDuration = nil
	r := retryTestInput{
		serviceName:              "core",
		httpResponseStatusCode:   404,
		header:                   map[string][]string{},
		responseError:            &errNet,
		expectedRetryTimeSeconds: 15,
		jitterMode:               true,
	}
	retryLoop(t, &r)
}

// Test a simple retry loop, simulating a 429 rate error
// issue-routing-tag: terraform/default
func TestUnitRetryLoop_basic(t *testing.T) {
	if httpreplay.ModeRecordReplay() {
		t.Skip("Skip Retry Tests in HttpReplay mode.")
	}
	ShortRetryTime = 15 * time.Second
	LongRetryTime = 30 * time.Second
	ConfiguredRetryDuration = nil
	r := retryTestInput{
		serviceName:              "core",
		httpResponseStatusCode:   404,
		header:                   map[string][]string{},
		responseError:            fmt.Errorf("Retriable error"),
		expectedRetryTimeSeconds: 0,
		jitterMode:               true,
	}
	retryLoop(t, &r)
}

// Configured retry timeout should be used for 429/500 errors
// issue-routing-tag: terraform/default
func TestUnitRetryLoop_configuredRetry(t *testing.T) {
	if httpreplay.ModeRecordReplay() {
		t.Skip("Skip Retry Tests in HttpReplay mode.")
	}
	ShortRetryTime = 15 * time.Second
	LongRetryTime = 15 * time.Second
	tmp := time.Duration(30 * time.Second)
	ConfiguredRetryDuration = &tmp
	r := retryTestInput{
		serviceName:              "core",
		httpResponseStatusCode:   429,
		header:                   map[string][]string{},
		responseError:            fmt.Errorf("Too many requests. "),
		expectedRetryTimeSeconds: 30,
		jitterMode:               true,
	}
	retryLoop(t, &r)
}

// issue-routing-tag: terraform/default
func TestUnitRetryLoop_outOfCapacity(t *testing.T) {
	if httpreplay.ModeRecordReplay() {
		t.Skip("Skip Retry Tests in HttpReplay mode.")
	}
	ShortRetryTime = 15 * time.Second
	LongRetryTime = 15 * time.Second
	tmp := time.Duration(30 * time.Second)
	ConfiguredRetryDuration = &tmp
	r := retryTestInput{
		serviceName:              "core",
		httpResponseStatusCode:   500,
		header:                   map[string][]string{},
		responseError:            fmt.Errorf("Out of host capacity. "),
		expectedRetryTimeSeconds: 0,
		jitterMode:               true,
	}
	retryLoop(t, &r)
}

// Even if a retry timeout is configured, it should be ignored for errors that are not 429/500
// issue-routing-tag: terraform/default
func TestUnitRetryLoop_configuredRetryWith404(t *testing.T) {
	if httpreplay.ModeRecordReplay() {
		t.Skip("Skip Retry Tests in HttpReplay mode.")
	}
	ShortRetryTime = 15 * time.Second
	LongRetryTime = 15 * time.Second
	tmp := time.Duration(60 * time.Second)
	ConfiguredRetryDuration = &tmp
	r := retryTestInput{
		serviceName:              "core",
		httpResponseStatusCode:   404,
		header:                   map[string][]string{},
		responseError:            fmt.Errorf("Retriable error"),
		expectedRetryTimeSeconds: 0,
		jitterMode:               true,
	}
	retryLoop(t, &r)
}

// Test concurrent retry loops
// issue-routing-tag: terraform/default
func TestUnitRetryLoop_concurrent(t *testing.T) {
	if httpreplay.ModeRecordReplay() {
		t.Skip("Skip Retry Tests in HttpReplay mode.")
	}
	ShortRetryTime = 15 * time.Second
	LongRetryTime = 15 * time.Second
	tmp := time.Duration(30 * time.Second)
	ConfiguredRetryDuration = &tmp
	r := retryTestInput{
		serviceName:              "core",
		httpResponseStatusCode:   500,
		header:                   map[string][]string{},
		responseError:            fmt.Errorf("Retriable error"),
		expectedRetryTimeSeconds: 30,
		jitterMode:               true,
	}

	workerFunc := func(t *testing.T, wg *sync.WaitGroup) {
		retryLoop(t, &r)
		wg.Done()
	}

	waitGroup := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		go workerFunc(t, &waitGroup)
	}

	waitGroup.Wait()
}

// issue-routing-tag: terraform/default
func TestUnitRetryKMSThrottling(t *testing.T) {
	if httpreplay.ModeRecordReplay() {
		t.Skip("Skip Retry Tests in HttpReplay mode.")
	}
	ShortRetryTime = 15 * time.Second
	LongRetryTime = 15 * time.Second
	ConfiguredRetryDuration = nil

	r := retryTestInput{
		serviceName:            "kms",
		httpResponseStatusCode: 429,
		header: map[string][]string{
			"retry-after": []string{"2"},
		},
		responseError:            fmt.Errorf("Retriable error"),
		expectedRetryTimeSeconds: 15,
		jitterMode:               false,
	}
	retryLoop(t, &r)
}

// issue-routing-tag: terraform/default
func TestUnitRetrySubnet409Conflict(t *testing.T) {
	if httpreplay.ModeRecordReplay() {
		t.Skip("Skip Retry Tests in HttpReplay mode.")
	}
	ShortRetryTime = 15 * time.Second
	LongRetryTime = 30 * time.Second
	ConfiguredRetryDuration = nil

	var subnetOptionals []interface{} = make([]interface{}, 2)
	subnetOptionals[0] = globalvar.SubnetService
	subnetOptionals[1] = globalvar.DeleteResource

	r := retryTestInput{
		serviceName:              coreService,
		httpResponseStatusCode:   409,
		header:                   map[string][]string{},
		responseError:            fmt.Errorf("Conflict"),
		optionals:                subnetOptionals,
		expectedRetryTimeSeconds: 30,
		jitterMode:               true,
	}
	retryLoop(t, &r)
}

// issue-routing-tag: terraform/default
func TestUnitRetrySubnet409OtherErrorMessage(t *testing.T) {
	if httpreplay.ModeRecordReplay() {
		t.Skip("Skip Retry Tests in HttpReplay mode.")
	}
	ShortRetryTime = 15 * time.Second
	LongRetryTime = 30 * time.Second
	ConfiguredRetryDuration = nil

	var subnetOptionals []interface{} = make([]interface{}, 2)
	subnetOptionals[0] = globalvar.SubnetService
	subnetOptionals[1] = globalvar.DeleteResource

	r := retryTestInput{
		serviceName:              coreService,
		httpResponseStatusCode:   409,
		header:                   map[string][]string{},
		responseError:            fmt.Errorf("other error message"),
		optionals:                subnetOptionals,
		expectedRetryTimeSeconds: 15,
		jitterMode:               true,
	}
	retryLoop(t, &r)
}

// issue-routing-tag: terraform/default
func TestUnitRetryDatabase(t *testing.T) {
	if httpreplay.ModeRecordReplay() {
		t.Skip("Skip Retry Tests in HttpReplay mode.")
	}
	ShortRetryTime = 15 * time.Second
	LongRetryTime = 30 * time.Second
	ConfiguredRetryDuration = nil

	r := retryTestInput{
		serviceName:              databaseService,
		httpResponseStatusCode:   409,
		header:                   map[string][]string{},
		responseError:            fmt.Errorf("InvalidatedRetryToken"),
		expectedRetryTimeSeconds: 0,
		jitterMode:               true,
	}
	retryLoop(t, &r)
}

// issue-routing-tag: terraform/default
func TestUnitRetryIdentity409ErrorInvalidatedRetryToken(t *testing.T) {
	if httpreplay.ModeRecordReplay() {
		t.Skip("Skip Retry Tests in HttpReplay mode.")
	}
	ShortRetryTime = 15 * time.Second
	LongRetryTime = 30 * time.Second
	ConfiguredRetryDuration = nil

	r := retryTestInput{
		serviceName:              identityService,
		httpResponseStatusCode:   409,
		header:                   map[string][]string{},
		responseError:            fmt.Errorf("InvalidatedRetryToken"),
		expectedRetryTimeSeconds: 0,
		jitterMode:               true,
	}
	retryLoop(t, &r)
}

// issue-routing-tag: terraform/default
func TestUnitRetryIdentity409ErrorNotAuthorizedOrResourceAlreadyExists(t *testing.T) {
	if httpreplay.ModeRecordReplay() {
		t.Skip("Skip Retry Tests in HttpReplay mode.")
	}
	ShortRetryTime = 15 * time.Second
	LongRetryTime = 30 * time.Second
	ConfiguredRetryDuration = nil

	r := retryTestInput{
		serviceName:              identityService,
		httpResponseStatusCode:   409,
		header:                   map[string][]string{},
		responseError:            fmt.Errorf("NotAuthorizedOrResourceAlreadyExists"),
		expectedRetryTimeSeconds: 30,
		jitterMode:               true,
	}
	retryLoop(t, &r)
}

// issue-routing-tag: terraform/default
func TestUnitRetryObjectStorage(t *testing.T) {
	if httpreplay.ModeRecordReplay() {
		t.Skip("Skip Retry Tests in HttpReplay mode.")
	}
	ShortRetryTime = 15 * time.Second
	LongRetryTime = 30 * time.Second
	ConfiguredRetryDuration = nil

	r := retryTestInput{
		serviceName:              objectstorageService,
		httpResponseStatusCode:   409,
		header:                   map[string][]string{},
		responseError:            fmt.Errorf("NotAuthorizedOrResourceAlreadyExists"),
		expectedRetryTimeSeconds: 30,
		jitterMode:               true,
	}
	retryLoop(t, &r)
}

// issue-routing-tag: terraform/default
func TestUnitRetryDbHomeWith404Error(t *testing.T) {
	if httpreplay.ModeRecordReplay() {
		t.Skip("Skip Retry Tests in HttpReplay mode.")
	}
	ShortRetryTime = 15 * time.Second
	LongRetryTime = 30 * time.Second
	ConfiguredRetryDuration = nil

	r := retryTestInput{
		serviceName:              databaseService,
		httpResponseStatusCode:   404,
		header:                   map[string][]string{},
		responseError:            fmt.Errorf("NotAuthorizedOrNotFound"),
		expectedRetryTimeSeconds: 0,
		jitterMode:               true,
		optionals:                []interface{}{GetDbHomeRetryDurationFunction(20 * time.Second)},
	}
	retryLoop(t, &r)
}

// issue-routing-tag: terraform/default
func TestUnitRetryDbHomeWithConflictingStateError(t *testing.T) {
	if httpreplay.ModeRecordReplay() {
		t.Skip("Skip Retry Tests in HttpReplay mode.")
	}
	ShortRetryTime = 15 * time.Second
	LongRetryTime = 30 * time.Second
	ConfiguredRetryDuration = nil

	r := retryTestInput{
		serviceName:              databaseService,
		httpResponseStatusCode:   409,
		header:                   map[string][]string{},
		responseError:            fmt.Errorf("The existing Db System with ID blahblahblah has a conflicting state of UPDATING."),
		expectedRetryTimeSeconds: 30,
		jitterMode:               true,
		optionals:                []interface{}{GetDbHomeRetryDurationFunction(20 * time.Second)},
	}
	retryLoop(t, &r)
}

// issue-routing-tag: terraform/default
func TestUnitRetryDbHomeWithInvalidatedRetryTokenError(t *testing.T) {
	if httpreplay.ModeRecordReplay() {
		t.Skip("Skip Retry Tests in HttpReplay mode.")
	}
	ShortRetryTime = 15 * time.Second
	LongRetryTime = 30 * time.Second
	ConfiguredRetryDuration = nil

	r := retryTestInput{
		serviceName:              databaseService,
		httpResponseStatusCode:   409,
		header:                   map[string][]string{},
		responseError:            fmt.Errorf("InvalidatedRetryToken"),
		expectedRetryTimeSeconds: 0,
		jitterMode:               true,
		optionals:                []interface{}{GetDbHomeRetryDurationFunction(20 * time.Second)},
	}
	retryLoop(t, &r)
}
