package provider

import (
	"math/rand"
	"strings"
	"time"

	oci_common "github.com/oracle/oci-go-sdk/common"
)

const (
	quadraticBackoffCap  = 12              // This corresponds to a 2*12*12=288 second cap on retry wait times (~5 minutes)
	minRetryBackoff      = 1 * time.Second // Must wait for at least 1 second before retrying
	identityService      = "identity"
	objectstorageService = "object_storage"
)

var shortRetryTime = 2 * time.Minute
var longRetryTime = 10 * time.Minute
var configuredRetryDuration *time.Duration

func init() {
	rand.Seed(time.Now().UnixNano())
}

func getRetryBackoffDuration(response oci_common.OCIOperationResponse, disableNotFoundRetries bool, service string, startTime time.Time) time.Duration {
	// Avoid having a very large retry backoff
	attempt := response.AttemptNumber
	if attempt > quadraticBackoffCap {
		attempt = quadraticBackoffCap
	}
	retryBackoffRange := time.Duration(2*attempt*attempt)*time.Second - minRetryBackoff

	// Jitter the backoff time. The actual backoff time might be anywhere within the minimum and quadratic backoff time to avoid clustering.
	backoffDuration := time.Duration(rand.Int63n(int64(retryBackoffRange+1))) + minRetryBackoff

	// If we are about to exceed the retry duration; then reduce the backoff so that next attempt happens roughly when
	// the entire retry duration is supposed to expire. Jitter is necessary again to avoid clustering.
	expectedRetryDuration := getExpectedRetryDuration(response, disableNotFoundRetries, service)
	timeWaited := getElapsedRetryDuration(startTime)
	if timeWaited < expectedRetryDuration && timeWaited+backoffDuration > expectedRetryDuration {
		extraJitterRange := int64(float64(expectedRetryDuration) * 0.05)
		finalBackoffDuration := expectedRetryDuration - timeWaited + time.Duration(rand.Int63n(extraJitterRange+1)) + minRetryBackoff
		if finalBackoffDuration < backoffDuration {
			backoffDuration = finalBackoffDuration
		}
	}
	return backoffDuration
}

func getElapsedRetryDuration(firstAttemptTime time.Time) time.Duration {
	return time.Now().Sub(firstAttemptTime)
}

func getExpectedRetryDuration(response oci_common.OCIOperationResponse, disableNotFoundRetries bool, service string) time.Duration {
	if response.Response == nil || response.Response.HTTPResponse() == nil {
		return 0
	}

	statusCode := response.Response.HTTPResponse().StatusCode
	e := response.Error

	if statusCode >= 200 && statusCode < 300 {
		return 0
	}

	switch statusCode {
	case 400, 401, 403:
		return 0
	case 404:
		if disableNotFoundRetries {
			return 0
		}
		if service == identityService || service == objectstorageService {
			return longRetryTime
		}
	case 409:
		if e != nil && strings.Contains(e.Error(), "InvalidatedRetryToken") {
			return 0
		}
		if service == identityService && e != nil &&
			(strings.Contains(e.Error(), "CompartmentAlreadyExists") ||
				strings.Contains(e.Error(), "TagDefinitionAlreadyExists") ||
				strings.Contains(e.Error(), "TenantCapacityExceeded") ||
				strings.Contains(e.Error(),
					"TagNamespaceAlreadyExists")) {
			return 0
		}
		if e != nil && strings.Contains(e.Error(), "NotAuthorizedOrResourceAlreadyExists") && (service == identityService || service == objectstorageService) {
			return longRetryTime
		}
	case 412:
		return 0
	case 429:
		if configuredRetryDuration != nil {
			return *configuredRetryDuration
		}
		return longRetryTime
	case 500:
		if configuredRetryDuration != nil {
			return *configuredRetryDuration
		}
		if service == objectstorageService {
			return longRetryTime
		}
	}
	return shortRetryTime
}

func shouldRetry(response oci_common.OCIOperationResponse, disableNotFoundRetries bool, service string, startTime time.Time) bool {
	return getElapsedRetryDuration(startTime) < getExpectedRetryDuration(response, disableNotFoundRetries, service)
}

// Because this function notes the start time for making should retry decisions, it's advised
// for this function call to be made immediately before the client API call.
func getRetryPolicy(disableNotFoundRetries bool, service string) *oci_common.RetryPolicy {
	startTime := time.Now()
	retryPolicy := &oci_common.RetryPolicy{
		MaximumNumberAttempts: 0,
		ShouldRetryOperation: func(response oci_common.OCIOperationResponse) bool {
			return shouldRetry(response, disableNotFoundRetries, service, startTime)
		},
		NextDuration: func(response oci_common.OCIOperationResponse) time.Duration {
			return getRetryBackoffDuration(response, disableNotFoundRetries, service, startTime)
		},
	}

	return retryPolicy
}
