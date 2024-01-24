// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package tfresource

import (
	"fmt"
	"net"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/globalvar"

	"github.com/stretchr/testify/assert"

	"github.com/oracle/terraform-provider-oci/httpreplay"

	"github.com/oracle/oci-go-sdk/v65/common"
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
type retryByEcInputRetry struct {
	serviceName              string
	disableNotFoundRetries   bool
	httpResponseStatusCode   int
	header                   map[string][]string
	expectedRetryTimeSeconds int
	responseError            error
	jitterMode               bool
	optionals                []interface{}
	EndOfWindowTime          *time.Time
}
type TestOCIOperationResponse struct {
	// Response from OCI Operation
	Response TestOCIResponse

	// Error from OCI Operation
	Error error

	// Operation Attempt Number (one-based)
	AttemptNumber uint

	// End of eventually consistent effects, or nil if no such effects
	EndOfWindowTime *time.Time

	// Backoff scaling factor (only used for dealing with eventual consistency)
	BackoffScalingFactor float64

	// Time of the initial attempt
	InitialAttemptTime time.Time
}

// ServiceError models all potential errors generated the service call
type TestServiceError interface {
	// The http status code of the error
	GetHTTPStatusCode() int

	// The human-readable error string as sent by the service
	GetMessage() string

	// A short error code that defines the error, meant for programmatic parsing.
	// See https://docs.cloud.oracle.com/Content/API/References/apierrors.htm
	GetCode() string

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	GetOpcRequestID() string
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
	ShortRetryTime = 1 * time.Second
	LongRetryTime = 2 * time.Second
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
	ShortRetryTime = 1 * time.Second
	LongRetryTime = 2 * time.Second
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
	ShortRetryTime = 1 * time.Second
	LongRetryTime = 2 * time.Second
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
	ShortRetryTime = 1 * time.Second
	LongRetryTime = 2 * time.Second
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
	ShortRetryTime = 1 * time.Second
	LongRetryTime = 2 * time.Second
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
	ShortRetryTime = 1 * time.Second
	LongRetryTime = 2 * time.Second
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
	ShortRetryTime = 1 * time.Second
	LongRetryTime = 2 * time.Second
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
	ShortRetryTime = 1 * time.Second
	LongRetryTime = 2 * time.Second
	ConfiguredRetryDuration = nil

	r := retryTestInput{
		serviceName:            "kms",
		httpResponseStatusCode: 429,
		header: map[string][]string{
			"retry-after": []string{"2"},
		},
		responseError:            fmt.Errorf("Retriable error"),
		expectedRetryTimeSeconds: 1,
		jitterMode:               false,
	}
	retryLoop(t, &r)
}

// issue-routing-tag: terraform/default
func TestUnitRetrySubnet409Conflict(t *testing.T) {
	if httpreplay.ModeRecordReplay() {
		t.Skip("Skip Retry Tests in HttpReplay mode.")
	}
	ShortRetryTime = 1 * time.Second
	LongRetryTime = 2 * time.Second
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
		expectedRetryTimeSeconds: 2,
		jitterMode:               true,
	}
	retryLoop(t, &r)
}

// issue-routing-tag: terraform/default
func TestUnitRetrySubnet409OtherErrorMessage(t *testing.T) {
	if httpreplay.ModeRecordReplay() {
		t.Skip("Skip Retry Tests in HttpReplay mode.")
	}
	ShortRetryTime = 1 * time.Second
	LongRetryTime = 2 * time.Second
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
		expectedRetryTimeSeconds: 1,
		jitterMode:               true,
	}
	retryLoop(t, &r)
}

// issue-routing-tag: terraform/default
func TestUnitRetryDatabase(t *testing.T) {
	if httpreplay.ModeRecordReplay() {
		t.Skip("Skip Retry Tests in HttpReplay mode.")
	}
	ShortRetryTime = 1 * time.Second
	LongRetryTime = 2 * time.Second
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
	ShortRetryTime = 1 * time.Second
	LongRetryTime = 2 * time.Second
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
	ShortRetryTime = 1 * time.Second
	LongRetryTime = 2 * time.Second
	ConfiguredRetryDuration = nil

	r := retryTestInput{
		serviceName:              identityService,
		httpResponseStatusCode:   409,
		header:                   map[string][]string{},
		responseError:            fmt.Errorf("NotAuthorizedOrResourceAlreadyExists"),
		expectedRetryTimeSeconds: 2,
		jitterMode:               true,
	}
	retryLoop(t, &r)
}

// issue-routing-tag: terraform/default
func TestUnitRetryObjectStorage(t *testing.T) {
	if httpreplay.ModeRecordReplay() {
		t.Skip("Skip Retry Tests in HttpReplay mode.")
	}
	ShortRetryTime = 1 * time.Second
	LongRetryTime = 2 * time.Second
	ConfiguredRetryDuration = nil

	r := retryTestInput{
		serviceName:              objectstorageService,
		httpResponseStatusCode:   409,
		header:                   map[string][]string{},
		responseError:            fmt.Errorf("NotAuthorizedOrResourceAlreadyExists"),
		expectedRetryTimeSeconds: 2,
		jitterMode:               true,
	}
	retryLoop(t, &r)
}

// issue-routing-tag: terraform/default
func TestUnitRetryDbHomeWith404Error(t *testing.T) {
	if httpreplay.ModeRecordReplay() {
		t.Skip("Skip Retry Tests in HttpReplay mode.")
	}
	ShortRetryTime = 1 * time.Second
	LongRetryTime = 2 * time.Second
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
	ShortRetryTime = 1 * time.Second
	LongRetryTime = 2 * time.Second
	ConfiguredRetryDuration = nil

	r := retryTestInput{
		serviceName:              databaseService,
		httpResponseStatusCode:   409,
		header:                   map[string][]string{},
		responseError:            fmt.Errorf("The existing Db System with ID blahblahblah has a conflicting state of UPDATING."),
		expectedRetryTimeSeconds: 2,
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
	ShortRetryTime = 1 * time.Second
	LongRetryTime = 2 * time.Second
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

// issue-routing-tag: terraform/default
func TestUnitisRetriableByEcServiceError(t *testing.T) {

	type args struct {
		response common.OCIOperationResponse
	}
	type testFormat struct {
		name   string
		args   args
		output bool
	}

	tests := []testFormat{
		{
			name:   "Test response with nil value",
			args:   args{response: common.OCIOperationResponse{Error: nil}},
			output: false,
		},
		{
			name:   "Test response with status code 404",
			args:   args{response: common.OCIOperationResponse{Error: fmt.Errorf("Retriable error")}},
			output: false,
		},
	}
	for _, test := range tests {
		t.Logf("Running %s", test.name)
		if res, _ := isRetriableByEc(test.args.response); res != test.output {
			t.Errorf("Output %t not equal to expected %t", res, test.output)
		}

	}
}

// issue-routing-tag: terraform/default
func TestUnitisRemainingEventualConsistencyDuration(t *testing.T) {
	if httpreplay.ModeRecordReplay() {
		t.Skip("Skip Retry Tests in HttpReplay mode.")
	}
	ShortRetryTime = 1 * time.Second
	LongRetryTime = 2 * time.Second
	ConfiguredRetryDuration = nil
	isServiceErrorVar = func(err error) (failure common.ServiceError, ok bool) {
		return nil, true
	}
	isErrorAffectedByEventualConsistency = func(Error error) bool {
		return false
	}
	r := common.OCIOperationResponse{
		Response:             TestOCIResponse{},
		Error:                fmt.Errorf("InvalidatedRetryToken"),
		AttemptNumber:        0,
		EndOfWindowTime:      nil,
		BackoffScalingFactor: 0,
		InitialAttemptTime:   time.Now(),
	}
	RetriableByEcFlag, _ := isRetriableByEc(r)
	assert.Equal(t, RetriableByEcFlag, false)
}

// issue-routing-tag: terraform/default
func TestUnitRemainingEventualConsistencyDuration(t *testing.T) {
	if httpreplay.ModeRecordReplay() {
		t.Skip("Skip Retry Tests in HttpReplay mode.")
	}
	ShortRetryTime = 1 * time.Second
	LongRetryTime = 2 * time.Second
	ConfiguredRetryDuration = nil
	r := common.OCIOperationResponse{
		Response:             TestOCIResponse{},
		Error:                fmt.Errorf("InvalidatedRetryToken"),
		AttemptNumber:        0,
		EndOfWindowTime:      nil,
		BackoffScalingFactor: 0,
		InitialAttemptTime:   time.Now(),
	}
	assert.Nil(t, getRemainingEventualConsistencyDuration(r))
	if res := getRemainingEventualConsistencyDuration(r); res != nil {
		t.Errorf("Time duration : %s", res)
	}
}

// issue-routing-tag: terraform/default
func TestUnitIdentityExpectedRetryDuration(t *testing.T) {
	if httpreplay.ModeRecordReplay() {
		t.Skip("Skip Retry Tests in HttpReplay mode.")
	}
	ShortRetryTime = 1 * time.Second
	LongRetryTime = 2 * time.Second
	ConfiguredRetryDuration = nil

	type args struct {
		response               common.OCIOperationResponse
		disableNotFoundRetries bool
		optionals              interface{}
	}
	type testFormat struct {
		name   string
		args   args
		output time.Duration
	}

	tests := []testFormat{
		{
			name:   "Test response with nil value",
			args:   args{response: common.OCIOperationResponse{Error: nil, Response: TestOCIResponse{}}, disableNotFoundRetries: true},
			output: 1 * time.Second,
		},
		{
			name:   "Test response with status code 404",
			args:   args{response: common.OCIOperationResponse{Error: nil, Response: TestOCIResponse{statusCode: 404}}, disableNotFoundRetries: true},
			output: 0,
		},
		{
			name:   "Test response with status code 409",
			args:   args{response: common.OCIOperationResponse{Error: fmt.Errorf("NotAuthorizedOrResourceAlreadyExists"), Response: TestOCIResponse{statusCode: 409}}, disableNotFoundRetries: true},
			output: 2 * time.Second,
		},
	}
	for _, test := range tests {
		t.Logf("Running %s", test.name)
		if res := getIdentityExpectedRetryDuration(test.args.response, test.args.disableNotFoundRetries); res != test.output {
			t.Errorf("Output %s not equal to expected %s", res, test.output)
		}

	}
}

// issue-routing-tag: terraform/default
func TestUnitObjectstorageServiceExpectedRetryDuration(t *testing.T) {
	if httpreplay.ModeRecordReplay() {
		t.Skip("Skip Retry Tests in HttpReplay mode.")
	}
	ShortRetryTime = 1 * time.Second
	LongRetryTime = 2 * time.Second

	type args struct {
		response               common.OCIOperationResponse
		disableNotFoundRetries bool
		optionals              interface{}
	}
	type testFormat struct {
		name   string
		args   args
		output time.Duration
	}

	tests := []testFormat{
		{
			name:   "Test response with nil value",
			args:   args{response: common.OCIOperationResponse{Error: nil, Response: TestOCIResponse{}}, disableNotFoundRetries: true},
			output: 1 * time.Second,
		},
		{
			name:   "Test response with status code 400",
			args:   args{response: common.OCIOperationResponse{Error: nil, Response: TestOCIResponse{statusCode: 400}}, disableNotFoundRetries: true},
			output: 0,
		},
		{
			name:   "Test response with status code 404",
			args:   args{response: common.OCIOperationResponse{Error: nil, Response: TestOCIResponse{statusCode: 404}}, disableNotFoundRetries: true},
			output: 0,
		},
		{
			name:   "Test response with status code 409",
			args:   args{response: common.OCIOperationResponse{Error: fmt.Errorf("NotAuthorizedOrResourceAlreadyExists"), Response: TestOCIResponse{statusCode: 409}}, disableNotFoundRetries: true},
			output: 2 * time.Second,
		},
		{
			name:   "Test response with status code 500",
			args:   args{response: common.OCIOperationResponse{Error: nil, Response: TestOCIResponse{statusCode: 500}}, disableNotFoundRetries: true},
			output: 2 * time.Second,
		},
	}
	for _, test := range tests {
		t.Logf("Running %s", test.name)
		if res := getObjectstorageServiceExpectedRetryDuration(test.args.response, test.args.disableNotFoundRetries); res != test.output {
			t.Errorf("Output %s not equal to expected %s", res, test.output)
		}

	}
}

// issue-routing-tag: terraform/default
func TestUnitObjectstorageServiceConfiguredRetryDuration(t *testing.T) {
	if httpreplay.ModeRecordReplay() {
		t.Skip("Skip Retry Tests in HttpReplay mode.")
	}
	ShortRetryTime = 1 * time.Second
	LongRetryTime = 2 * time.Second

	r := retryTestInput{
		serviceName:              objectstorageService,
		httpResponseStatusCode:   429,
		header:                   map[string][]string{},
		responseError:            fmt.Errorf("Out of host capacity"),
		expectedRetryTimeSeconds: 2,
		jitterMode:               true,
	}
	retryLoop(t, &r)
}

// issue-routing-tag: terraform/default
func TestUnitLogAnalyticsExpectedRetryDuration(t *testing.T) {
	if httpreplay.ModeRecordReplay() {
		t.Skip("Skip Retry Tests in HttpReplay mode.")
	}
	ShortRetryTime = 1 * time.Second
	LongRetryTime = 2 * time.Second

	r := retryTestInput{
		serviceName:              logAnalyticsService,
		httpResponseStatusCode:   304,
		header:                   map[string][]string{},
		responseError:            fmt.Errorf("Not Modified"),
		expectedRetryTimeSeconds: 0,
		jitterMode:               true,
	}
	retryLoop(t, &r)
}

// issue-routing-tag: terraform/default
func TestUnitWaasExpectedRetryDuration(t *testing.T) {
	if httpreplay.ModeRecordReplay() {
		t.Skip("Skip Retry Tests in HttpReplay mode.")
	}

	type args struct {
		response               common.OCIOperationResponse
		disableNotFoundRetries bool
		optionals              interface{}
	}
	type testFormat struct {
		name   string
		args   args
		output time.Duration
	}
	var subnetOptionals = make([]interface{}, 1)
	subnetOptionals[0] = "delete"
	tests := []testFormat{
		{
			name:   "Test response with optionals parameter",
			args:   args{response: common.OCIOperationResponse{Error: nil, Response: TestOCIResponse{statusCode: 409}}, disableNotFoundRetries: true, optionals: subnetOptionals[0]},
			output: 1 * time.Second,
		},
	}
	for _, test := range tests {
		t.Logf("Running %s", test.name)
		if res := getWaasExpectedRetryDuration(test.args.response, test.args.disableNotFoundRetries, test.args.optionals); res != test.output {
			t.Errorf("Output %s not equal to expected %s", res, test.output)
		}

	}
}

// issue-routing-tag: terraform/default
func TestUnitWaasCertificateExpectedRetryDuration(t *testing.T) {
	if httpreplay.ModeRecordReplay() {
		t.Skip("Skip Retry Tests in HttpReplay mode.")
	}
	type args struct {
		response               common.OCIOperationResponse
		disableNotFoundRetries bool
		optionals              interface{}
	}
	type testFormat struct {
		name   string
		args   args
		output time.Duration
	}
	var subnetOptionals = make([]interface{}, 1)
	subnetOptionals[0] = "delete"
	tests := []testFormat{
		{
			name:   "Test response with nil value",
			args:   args{response: common.OCIOperationResponse{Error: nil, Response: TestOCIResponse{}}, disableNotFoundRetries: true},
			output: 1 * time.Second,
		},
		{
			name:   "Test response with optionals parameter",
			args:   args{response: common.OCIOperationResponse{Error: nil, Response: TestOCIResponse{statusCode: 409}}, disableNotFoundRetries: true, optionals: subnetOptionals[0]},
			output: 1 * time.Second,
		},
	}
	for _, test := range tests {
		t.Logf("Running %s", test.name)
		if res := getWaasCertificateExpectedRetryDuration(test.args.response, test.args.disableNotFoundRetries, test.args.optionals); res != test.output {
			t.Errorf("Output %s not equal to expected %s", res, test.output)
		}

	}
}

// issue-routing-tag: terraform/default
func TestUnitDefaultExpectedRetryDuration(t *testing.T) {
	type args struct {
		response               common.OCIOperationResponse
		disableNotFoundRetries bool
	}
	type testFormat struct {
		name   string
		args   args
		output time.Duration
	}

	tests := []testFormat{
		{
			name:   "Test response with nil value",
			args:   args{response: common.OCIOperationResponse{Error: nil, Response: nil}, disableNotFoundRetries: true},
			output: 0,
		},
		{
			name:   "Test response with status code 404",
			args:   args{response: common.OCIOperationResponse{Error: nil, Response: TestOCIResponse{statusCode: 404}}, disableNotFoundRetries: true},
			output: 0,
		},
		{
			name:   "Test response with status code 400",
			args:   args{response: common.OCIOperationResponse{Error: nil, Response: TestOCIResponse{statusCode: 400}}, disableNotFoundRetries: true},
			output: 0,
		},
		{
			name:   "Test response with status code 412",
			args:   args{response: common.OCIOperationResponse{Error: nil, Response: TestOCIResponse{statusCode: 412}}, disableNotFoundRetries: true},
			output: 0,
		},
	}
	for _, test := range tests {
		t.Logf("Running %s", test.name)
		if res := GetDefaultExpectedRetryDuration(test.args.response, test.args.disableNotFoundRetries); res != test.output {
			t.Errorf("Output %s not equal to expected %s", res, test.output)
		}

	}
}
