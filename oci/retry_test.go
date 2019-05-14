// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/oracle/oci-go-sdk/common"
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
	retryPolicy := getRetryPolicy(r.disableNotFoundRetries, r.serviceName, r.optionals...)
	startTime := time.Now()

	for i := uint(1); true; i++ {
		operationResponse := common.NewOCIOperationResponse(TestOCIResponse{statusCode: r.httpResponseStatusCode, header: r.header}, r.responseError, i)

		expectedShouldRetry := getElapsedRetryDuration(startTime) < (time.Duration(r.expectedRetryTimeSeconds) * time.Second)
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
		fmt.Printf("Attempt #%v: Will wait for %v ms\n", i, waitTime.Nanoseconds()/1000000)

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

// Test a simple retry loop, simulating a 429 rate error
func TestRetryLoop_basic(t *testing.T) {
	shortRetryTime = 15 * time.Second
	longRetryTime = 30 * time.Second
	configuredRetryDuration = nil
	r := retryTestInput{
		serviceName:              "core",
		httpResponseStatusCode:   404,
		header:                   map[string][]string{},
		responseError:            fmt.Errorf("Retriable error"),
		expectedRetryTimeSeconds: 15,
		jitterMode:               true,
	}
	retryLoop(t, &r)
}

// Configured retry timeout should be used for 429/500 errors
func TestRetryLoop_configuredRetry(t *testing.T) {
	shortRetryTime = 15 * time.Second
	longRetryTime = 15 * time.Second
	tmp := time.Duration(30 * time.Second)
	configuredRetryDuration = &tmp
	r := retryTestInput{
		serviceName:              "core",
		httpResponseStatusCode:   429,
		header:                   map[string][]string{},
		responseError:            fmt.Errorf("Retriable error"),
		expectedRetryTimeSeconds: 30,
		jitterMode:               true,
	}
	retryLoop(t, &r)
}

// Even if a retry timeout is configured, it should be ignored for errors that are not 429/500
func TestRetryLoop_configuredRetryWith404(t *testing.T) {
	shortRetryTime = 15 * time.Second
	longRetryTime = 15 * time.Second
	tmp := time.Duration(60 * time.Second)
	configuredRetryDuration = &tmp
	r := retryTestInput{
		serviceName:              "core",
		httpResponseStatusCode:   404,
		header:                   map[string][]string{},
		responseError:            fmt.Errorf("Retriable error"),
		expectedRetryTimeSeconds: 15,
		jitterMode:               true,
	}
	retryLoop(t, &r)
}

// Test concurrent retry loops
func TestRetryLoop_concurrent(t *testing.T) {
	shortRetryTime = 15 * time.Second
	longRetryTime = 15 * time.Second
	tmp := time.Duration(30 * time.Second)
	configuredRetryDuration = &tmp
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

func TestRetryKMSThrottling(t *testing.T) {
	shortRetryTime = 15 * time.Second
	longRetryTime = 15 * time.Second
	configuredRetryDuration = nil

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

func TestRetrySubnet409Conflict(t *testing.T) {
	shortRetryTime = 15 * time.Second
	longRetryTime = 30 * time.Second
	configuredRetryDuration = nil

	var subnetOptionals []interface{} = make([]interface{}, 2)
	subnetOptionals[0] = subnetService
	subnetOptionals[1] = deleteResource

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

func TestRetrySubnet409OtherErrorMessage(t *testing.T) {
	shortRetryTime = 15 * time.Second
	longRetryTime = 30 * time.Second
	configuredRetryDuration = nil

	var subnetOptionals []interface{} = make([]interface{}, 2)
	subnetOptionals[0] = subnetService
	subnetOptionals[1] = deleteResource

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

func TestRetryDatabase(t *testing.T) {
	shortRetryTime = 15 * time.Second
	longRetryTime = 30 * time.Second
	configuredRetryDuration = nil

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

func TestRetryIdentity409ErrorInvalidatedRetryToken(t *testing.T) {
	shortRetryTime = 15 * time.Second
	longRetryTime = 30 * time.Second
	configuredRetryDuration = nil

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

func TestRetryIdentity409ErrorNotAuthorizedOrResourceAlreadyExists(t *testing.T) {
	shortRetryTime = 15 * time.Second
	longRetryTime = 30 * time.Second
	configuredRetryDuration = nil

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

func TestRetryObjectStorage(t *testing.T) {
	shortRetryTime = 15 * time.Second
	longRetryTime = 30 * time.Second
	configuredRetryDuration = nil

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
