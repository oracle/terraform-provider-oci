package provider

import (
	"net/http"
	"strings"
	"time"

	oci_common "github.com/oracle/oci-go-sdk/common"
)

const (
	shortRetryTime       = 2 * time.Minute
	longRetryTime        = 10 * time.Minute
	identityService      = "identity"
	objectstorageService = "object_storage"
)

var timeWaitedCache map[uint]time.Duration

//attempt starts at 1
//quadratic backoff (attempt^2) with forced retries at shortRetryTime and longRetryTime
func getNextDuration(attempt uint) time.Duration {
	timeWaited := getTimeWaited(attempt)
	nextDuration := time.Duration(attempt*attempt) * time.Second
	if timeWaited < shortRetryTime && nextDuration+timeWaited > shortRetryTime {
		nextDuration = shortRetryTime - timeWaited
	}
	if timeWaited < longRetryTime && nextDuration+timeWaited > longRetryTime {
		nextDuration = longRetryTime - timeWaited
	}
	return nextDuration
}

func getTimeWaited(attempt uint) time.Duration {
	if timeWaitedCache == nil {
		timeWaitedCache = map[uint]time.Duration{}
	}
	if timeWaited, ok := timeWaitedCache[attempt]; ok {
		return timeWaited
	}
	if attempt <= 1 {
		return 0
	}
	timeWaited := getTimeWaited(attempt-1) + getNextDuration(attempt-1)
	timeWaitedCache[attempt] = timeWaited
	return timeWaited
}

func shouldRetry(response *http.Response, e error, attempt uint, disableNotFoundRetries bool, service string) bool {
	if disableAutoRetries {
		return false
	}
	if response == nil {
		return false
	}
	if response.StatusCode >= 200 && response.StatusCode < 300 {
		return false
	}

	timeWaited := getTimeWaited(attempt)
	shortTimeDecision := timeWaited < shortRetryTime
	longTimeDecision := timeWaited < longRetryTime
	switch response.StatusCode {
	case 400:
		return false
	case 401:
		return false
	case 403:
		return false
	case 404:
		if disableNotFoundRetries {
			return false
		}
		if service == identityService || service == objectstorageService {
			return longTimeDecision
		}
	case 409:
		if e != nil && strings.Contains(e.Error(), "InvalidatedRetryToken") {
			return false
		}
		if service == identityService && e != nil && strings.Contains(e.Error(), "CompartmentAlreadyExists") {
			return false
		}
		if e != nil && strings.Contains(e.Error(), "NotAuthorizedOrResourceAlreadyExists") && (service == identityService || service == objectstorageService) {
			return longTimeDecision
		}
	case 412:
		return false
	case 429:
		return longTimeDecision
	case 500:
		if service == objectstorageService {
			return longTimeDecision
		}
	}
	return shortTimeDecision
}

func getRetryOptions(disableNotFoundRetries bool, service string) []oci_common.RetryPolicyOption {
	retryOptions := []oci_common.RetryPolicyOption{
		oci_common.MaximumNumberAttempts(0),
		oci_common.GetNextDuration(getNextDuration),
		oci_common.ShouldRetryOperation(func(response *http.Response, e error, attempt uint) bool {
			return shouldRetry(response, e, attempt, disableNotFoundRetries, service)
		}),
	}

	return retryOptions
}
