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
}

type retryTestInput struct {
	serviceName              string
	disableNotFoundRetries   bool
	httpResponseStatusCode   int
	expectedRetryTimeSeconds int
}

func (response TestOCIResponse) HTTPResponse() *http.Response {
	result := http.Response{}
	result.Request = &http.Request{}
	result.StatusCode = response.statusCode
	return &result
}

func retryLoop(t *testing.T, r *retryTestInput) {
	startTime := time.Now()
	for i := uint(1); true; i++ {
		operationResponse := common.NewOCIOperationResponse(TestOCIResponse{statusCode: r.httpResponseStatusCode}, fmt.Errorf("Retriable error"), i)

		expectedShouldRetry := getElapsedRetryDuration(startTime) < (time.Duration(r.expectedRetryTimeSeconds) * time.Second)
		actualShouldRetry := shouldRetry(operationResponse, r.disableNotFoundRetries, r.serviceName, startTime)
		if actualShouldRetry != expectedShouldRetry {
			t.Errorf("Expected shouldRetry to return %v for attempt %v", expectedShouldRetry, i)
			return
		}

		if !actualShouldRetry {
			fmt.Println("Timeout exceeded; no retry.")
			return
		}

		waitTime := getRetryBackoffDuration(operationResponse, r.disableNotFoundRetries, r.serviceName, startTime)
		fmt.Printf("Attempt #%v: Will wait for %v ms\n", i, waitTime.Nanoseconds()/1000000)
		expectedWaitTimeMax := time.Duration(2*i*i) * time.Second
		if i > quadraticBackoffCap {
			expectedWaitTimeMax = time.Duration(2*quadraticBackoffCap*quadraticBackoffCap) * time.Second
		}
		if waitTime >= expectedWaitTimeMax || waitTime < minRetryBackoff {
			t.Errorf("Expected wait time to be between %v and %v for attempt %v, but got %v", minRetryBackoff, expectedWaitTimeMax, i, waitTime)
			return
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
		expectedRetryTimeSeconds: 15,
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
		expectedRetryTimeSeconds: 30,
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
		expectedRetryTimeSeconds: 15,
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
		expectedRetryTimeSeconds: 30,
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
