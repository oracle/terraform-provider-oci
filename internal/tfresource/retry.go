// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package tfresource

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/globalvar"

	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	oci_common "github.com/oracle/oci-go-sdk/v56/common"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

const (
	quadraticBackoffCap             = 12              // This corresponds to a 2*12*12=288 second cap on retry wait times (~5 minutes)
	minRetryBackoff                 = 1 * time.Second // Must wait for at least 1 second before retrying
	databaseService                 = "database"
	identityService                 = "identity"
	coreService                     = "core"
	WaasService                     = "waas"
	kmsService                      = "kms"
	objectstorageService            = "object_storage"
	logAnalyticsService             = "log_analytics"
	updateResource                  = "update"
	createResource                  = "create"
	getResource                     = "get"
	waasDeleteConflictRetryDuration = 60 * time.Minute
	certificateService              = "certificate"
)

type ServiceExpectedRetryDurationFunc func(response oci_common.OCIOperationResponse, disableNotFoundRetries bool, optionals ...interface{}) time.Duration
type expectedRetryDurationFn func(response oci_common.OCIOperationResponse, disableNotFoundRetries bool, service string, optionals ...interface{}) time.Duration
type getRetryPolicyFunc func(disableNotFoundRetries bool, service string, optionals ...interface{}) *oci_common.RetryPolicy

var serviceExpectedRetryDurationMap = map[string]ServiceExpectedRetryDurationFunc{
	coreService:          getCoreExpectedRetryDuration,
	databaseService:      getDatabaseExpectedRetryDuration,
	identityService:      getIdentityExpectedRetryDuration,
	objectstorageService: getObjectstorageServiceExpectedRetryDuration,
	WaasService:          getWaasExpectedRetryDuration,
	logAnalyticsService:  getLogAnalyticsExpectedRetryDuration,
}
var serviceRetryPolicyFnMap = map[string]getRetryPolicyFunc{
	kmsService: kmsGetRetryPolicy,
}

var ShortRetryTime = 2 * time.Minute
var LongRetryTime = 10 * time.Minute
var ConfiguredRetryDuration *time.Duration

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GetRetryBackoffDuration(response oci_common.OCIOperationResponse, disableNotFoundRetries bool, service string, startTime time.Time, optionals ...interface{}) time.Duration {
	return getRetryBackoffDurationWithExpectedRetryDurationFn(response, disableNotFoundRetries, service, startTime, getExpectedRetryDuration, optionals...)
}

func getRetryBackoffDurationWithExpectedRetryDurationFn(response oci_common.OCIOperationResponse, disableNotFoundRetries bool, service string, startTime time.Time, expectedRetryDurationFn expectedRetryDurationFn, optionals ...interface{}) time.Duration {
	if httpreplay.ShouldRetryImmediately() {
		return 0
	}

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
	expectedRetryDuration := expectedRetryDurationFn(response, disableNotFoundRetries, service, optionals...)
	timeWaited := GetElapsedRetryDuration(startTime)
	if timeWaited < expectedRetryDuration && timeWaited+backoffDuration > expectedRetryDuration {
		extraJitterRange := int64(float64(expectedRetryDuration) * 0.05)
		finalBackoffDuration := expectedRetryDuration - timeWaited + time.Duration(rand.Int63n(extraJitterRange+1)) + minRetryBackoff
		if finalBackoffDuration < backoffDuration {
			backoffDuration = finalBackoffDuration
		}
	}
	utils.Logf("Time elapsed for retry: %v;  Expected retry duration: %v \n", timeWaited.Round(time.Second), expectedRetryDuration.Round(time.Second))
	return backoffDuration
}

func GetElapsedRetryDuration(firstAttemptTime time.Time) time.Duration {
	return time.Now().Sub(firstAttemptTime)
}

func getExpectedRetryDuration(response oci_common.OCIOperationResponse, disableNotFoundRetries bool, service string, optionals ...interface{}) time.Duration {
	// Get the override retry duration function if it exists. This gives the most granular control over what value to return, and is passed
	// into GetRetryPolicy function as an optional argument to override retry durations on a per API basis.
	if len(optionals) > 0 {
		if overrideRetryDurationFn, ok := optionals[0].(expectedRetryDurationFn); ok {
			return overrideRetryDurationFn(response, disableNotFoundRetries, service, optionals)
		}
	}

	// Use the service specific retry duration calculation if it exists
	if retryDurationFn, ok := serviceExpectedRetryDurationMap[service]; ok {
		return retryDurationFn(response, disableNotFoundRetries, optionals...)
	}

	// Use the default retry duration computation
	return GetDefaultExpectedRetryDuration(response, disableNotFoundRetries)
}

func GetDefaultExpectedRetryDuration(response oci_common.OCIOperationResponse, disableNotFoundRetries bool) time.Duration {
	defaultRetryTime := ShortRetryTime

	if oci_common.IsNetworkError(response.Error) {
		log.Printf("[DEBUG] Retrying for network error...")
		return defaultRetryTime
	}

	if ok, remainingEc := isRetriableByEc(response); ok && remainingEc != nil {
		log.Printf("[DEBUG] Retrying Eventual consistency...")
		return *remainingEc
	}

	if response.Response == nil || response.Response.HTTPResponse() == nil {
		return 0
	}

	statusCode := response.Response.HTTPResponse().StatusCode
	e := response.Error

	if statusCode >= 200 && statusCode < 300 {
		return 0
	}

	switch statusCode {
	case 400, 401, 403, 413:
		return 0
	case 404:
		return 0
	case 409:
		if isDisable409Retry, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("disable_409_retry", "false")); isDisable409Retry {
			log.Printf("[ERROR] Resource is in conflict state due to multiple update request: %v", e.Error())
			return 0
		}
		if e != nil && (strings.Contains(e.Error(), "InvalidatedRetryToken") ||
			strings.Contains(e.Error(), "BucketNotEmpty")) {
			return 0
		}
	case 412:
		return 0
	case 429:
		if ConfiguredRetryDuration != nil {
			return *ConfiguredRetryDuration
		}
		defaultRetryTime = LongRetryTime
	case 500:
		if e != nil && (strings.Contains(e.Error(), "Out of host capacity")) {
			return 0
		}
		if ConfiguredRetryDuration != nil {
			return *ConfiguredRetryDuration
		}
	}

	return defaultRetryTime
}

func isRetriableByEc(r oci_common.OCIOperationResponse) (bool, *time.Duration) {
	if _, ok := oci_common.IsServiceError(r.Error); ok {
		now := time.Now()
		if r.EndOfWindowTime == nil || r.EndOfWindowTime.Before(now) {
			// either no eventually consistent effects, or they have disappeared by now
			utils.Debugln(fmt.Sprintf("EC.ShouldRetryOperation, no EC or in the past, returning false: endOfWindowTime = %v, now = %v", r.EndOfWindowTime, now))
			return false, nil
		}
		// there were eventually consistent effects present at the time of the first request
		// and they could still affect the retries
		if oci_common.IsErrorAffectedByEventualConsistency(r.Error) {
			// and it's one of the three affected error codes
			utils.Debugln(fmt.Sprintf("EC.ShouldRetryOperation, affected by EC, EC is present: endOfWindowTime = %v, now = %v", r.EndOfWindowTime, now))
			return true, getRemainingEventualConsistencyDuration(r)
		}
		return false, nil
	}
	return false, nil
}

// returns the remaining time with eventual consistency, or nil if there is no eventual consistency
func getRemainingEventualConsistencyDuration(r oci_common.OCIOperationResponse) *time.Duration {
	endOfWindow := oci_common.EcContext.GetEndOfWindow()
	if endOfWindow != nil {
		// there was an eventually consistent request
		if endOfWindow.After(r.InitialAttemptTime) {
			// and the eventually consistent effects may still be present
			remaining := endOfWindow.Sub(time.Now())
			if remaining > 0 {
				return &remaining
			}
		}
	}

	return nil
}

func getIdentityExpectedRetryDuration(response oci_common.OCIOperationResponse, disableNotFoundRetries bool, optionals ...interface{}) time.Duration {
	defaultRetryTime := GetDefaultExpectedRetryDuration(response, disableNotFoundRetries)
	if oci_common.IsNetworkError(response.Error) {
		return defaultRetryTime
	}

	if ok, remainingEc := isRetriableByEc(response); ok && remainingEc != nil {
		log.Printf("[DEBUG] Retrying Eventual consistency...")
		return *remainingEc
	}

	if response.Response == nil || response.Response.HTTPResponse() == nil {
		return defaultRetryTime
	}
	e := response.Error
	switch statusCode := response.Response.HTTPResponse().StatusCode; statusCode {
	case 404:
		return 0
	case 409:
		if e := response.Error; e != nil {
			if strings.Contains(e.Error(), "CompartmentAlreadyExists") || strings.Contains(e.Error(), "TagDefinitionAlreadyExists") ||
				strings.Contains(e.Error(), "TenantCapacityExceeded") || strings.Contains(e.Error(), "TagNamespaceAlreadyExists") ||
				strings.Contains(e.Error(), "InvalidatedRetryToken") {
				defaultRetryTime = 0
			} else if strings.Contains(e.Error(), "NotAuthorizedOrResourceAlreadyExists") {
				defaultRetryTime = LongRetryTime
			}
		}
		if isDisable409Retry, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("disable_409_retry", "false")); isDisable409Retry {
			log.Printf("[ERROR] Resource is in conflict state due to multiple update request: %v", e.Error())
			return 0
		}
	}
	return defaultRetryTime
}

func getDatabaseExpectedRetryDuration(response oci_common.OCIOperationResponse, disableNotFoundRetries bool, optionals ...interface{}) time.Duration {
	defaultRetryTime := GetDefaultExpectedRetryDuration(response, disableNotFoundRetries)
	if oci_common.IsNetworkError(response.Error) {
		return defaultRetryTime
	}

	if ok, remainingEc := isRetriableByEc(response); ok && remainingEc != nil {
		log.Printf("[DEBUG] Retrying Eventual consistency...")
		return *remainingEc
	}

	if response.Response == nil || response.Response.HTTPResponse() == nil {
		return defaultRetryTime
	}
	e := response.Error
	switch statusCode := response.Response.HTTPResponse().StatusCode; statusCode {
	case 409:
		if isDisable409Retry, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("disable_409_retry", "false")); isDisable409Retry {
			log.Printf("[ERROR] Resource is in conflict state due to multiple update request: %v", e.Error())
			return 0
		}
		if e := response.Error; e != nil {
			if strings.Contains(e.Error(), "InvalidatedRetryToken") {
				defaultRetryTime = 0
			} else {
				defaultRetryTime = LongRetryTime
			}
		}
	}
	return defaultRetryTime
}

func getObjectstorageServiceExpectedRetryDuration(response oci_common.OCIOperationResponse, disableNotFoundRetries bool, optionals ...interface{}) time.Duration {
	defaultRetryTime := GetDefaultExpectedRetryDuration(response, disableNotFoundRetries)
	if oci_common.IsNetworkError(response.Error) {
		return defaultRetryTime
	}

	if ok, remainingEc := isRetriableByEc(response); ok && remainingEc != nil {
		log.Printf("[DEBUG] Retrying Eventual consistency...")
		return *remainingEc
	}

	if response.Response == nil || response.Response.HTTPResponse() == nil {
		return defaultRetryTime
	}
	e := response.Error
	switch statusCode := response.Response.HTTPResponse().StatusCode; statusCode {
	case 404:
		return 0
	case 409:
		if e := response.Error; e != nil {
			if strings.Contains(e.Error(), "NotAuthorizedOrResourceAlreadyExists") {
				defaultRetryTime = LongRetryTime
			}
		}
		if isDisable409Retry, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("disable_409_retry", "false")); isDisable409Retry {
			log.Printf("[ERROR] Resource is in conflict state due to multiple update request: %v", e.Error())
			return 0
		}
	case 500:
		if ConfiguredRetryDuration != nil {
			defaultRetryTime = *ConfiguredRetryDuration
		} else {
			defaultRetryTime = LongRetryTime
		}
	}

	return defaultRetryTime
}

func getLogAnalyticsExpectedRetryDuration(response oci_common.OCIOperationResponse, disableNotFoundRetries bool, optionals ...interface{}) time.Duration {
	defaultRetryTime := GetDefaultExpectedRetryDuration(response, disableNotFoundRetries)
	if response.Response == nil || response.Response.HTTPResponse() == nil {
		return defaultRetryTime
	}

	statusCode := response.Response.HTTPResponse().StatusCode
	// 304 (Not Modified) is a successful return code for Log Analytics. don't retry.
	if statusCode == 304 {
		return 0
	}

	return defaultRetryTime
}

func ShouldRetry(response oci_common.OCIOperationResponse, disableNotFoundRetries bool, service string, startTime time.Time, optionals ...interface {
}) bool {
	return GetElapsedRetryDuration(startTime) < getExpectedRetryDuration(response, disableNotFoundRetries, service, optionals...)
}

// Because this function notes the start time for making should retry decisions, it's advised
// for this function call to be made immediately before the client API call.
func GetRetryPolicy(disableNotFoundRetries bool, service string, optionals ...interface{}) *oci_common.RetryPolicy {
	if serviceRetryPolicyFn, ok := serviceRetryPolicyFnMap[service]; ok {
		return serviceRetryPolicyFn(disableNotFoundRetries, service, optionals...)
	}
	return getDefaultRetryPolicy(disableNotFoundRetries, service, optionals...)
}

func getDefaultRetryPolicy(disableNotFoundRetries bool, service string, optionals ...interface{}) *oci_common.RetryPolicy {
	startTime := time.Now()
	retryPolicy := &oci_common.RetryPolicy{
		MaximumNumberAttempts: 0,
		ShouldRetryOperation: func(response oci_common.OCIOperationResponse) bool {
			return ShouldRetry(response, disableNotFoundRetries, service, startTime, optionals...)
		},
		NextDuration: func(response oci_common.OCIOperationResponse) time.Duration {
			return GetRetryBackoffDuration(response, disableNotFoundRetries, service, startTime, optionals...)
		},
	}

	return retryPolicy
}

var CoreServiceExpectedRetryDurationMap = map[string]ServiceExpectedRetryDurationFunc{
	globalvar.SubnetService: getSubnetExpectedRetryDuration,
}

func getSubnetExpectedRetryDuration(response oci_common.OCIOperationResponse, disableNotFoundRetries bool, optionals ...interface{}) time.Duration {
	defaultRetryTime := GetDefaultExpectedRetryDuration(response, disableNotFoundRetries)
	if response.Response == nil || response.Response.HTTPResponse() == nil {
		return defaultRetryTime
	}
	e := response.Error
	if len(optionals) > 0 {
		if key, ok := optionals[0].(string); ok {
			switch key {
			case globalvar.DeleteResource:
				switch statusCode := response.Response.HTTPResponse().StatusCode; statusCode {
				case 409:
					if isDisable409Retry, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("disable_409_retry", "false")); isDisable409Retry {
						log.Printf("[ERROR] Resource is in conflict state due to multiple update request: %v", e.Error())
						return 0
					}
					if e := response.Error; e != nil {
						if strings.Contains(e.Error(), "Conflict") {
							defaultRetryTime = LongRetryTime
						}
					}
				}
			}
		}
	}
	return defaultRetryTime
}

func getCoreExpectedRetryDuration(response oci_common.OCIOperationResponse, disableNotFoundRetries bool, optionals ...interface{}) time.Duration {
	if len(optionals) > 0 {
		if key, ok := optionals[0].(string); ok {
			if expectedRetryDurationFunc, ok := CoreServiceExpectedRetryDurationMap[key]; ok {
				return expectedRetryDurationFunc(response, disableNotFoundRetries, optionals[1:]...)
			}
		}
	}
	return GetDefaultExpectedRetryDuration(response, disableNotFoundRetries)
}

func kmsGetRetryPolicy(disableNotFoundRetries bool, service string, optionals ...interface{}) *oci_common.RetryPolicy {
	startTime := time.Now()
	retryPolicy := &oci_common.RetryPolicy{
		MaximumNumberAttempts: 0,
		ShouldRetryOperation: func(response oci_common.OCIOperationResponse) bool {
			return ShouldRetry(response, disableNotFoundRetries, service, startTime, optionals...)
		},
		NextDuration: func(response oci_common.OCIOperationResponse) time.Duration {
			return getKmsNextRetryDuration(response, disableNotFoundRetries, startTime, optionals...)
		},
	}
	return retryPolicy
}

func getKmsNextRetryDuration(response oci_common.OCIOperationResponse, disableNotFoundRetries bool, startTime time.Time, optionals ...interface{}) time.Duration {
	if httpreplay.ShouldRetryImmediately() {
		return 0
	}
	defaultRetryTime := GetRetryBackoffDuration(response, disableNotFoundRetries, "kms", startTime, optionals...)
	if response.Response == nil || response.Response.HTTPResponse() == nil {
		return defaultRetryTime
	}
	switch statusCode := response.Response.HTTPResponse().StatusCode; statusCode {
	case 429:
		rawResponse := response.Response.HTTPResponse()
		if retryAfterVal := rawResponse.Header["retry-after"]; len(retryAfterVal) > 0 {
			if i, err := strconv.Atoi(retryAfterVal[0]); err == nil {
				return time.Duration(i) * time.Second
			}
		}
	}
	return defaultRetryTime
}

func getWaasExpectedRetryDuration(response oci_common.OCIOperationResponse, disableNotFoundRetries bool, optionals ...interface{}) time.Duration {
	if len(optionals) > 0 {
		if key, ok := optionals[0].(string); ok {
			if expectedRetryDurationFunc, ok := waasServiceExpectedRetryDurationMap[key]; ok {
				return expectedRetryDurationFunc(response, disableNotFoundRetries, optionals[1:]...)
			}
		}
	}
	return GetDefaultExpectedRetryDuration(response, disableNotFoundRetries)

}

var waasServiceExpectedRetryDurationMap = map[string]ServiceExpectedRetryDurationFunc{
	certificateService: getWaasCertificateExpectedRetryDuration,
}

func getWaasCertificateExpectedRetryDuration(response oci_common.OCIOperationResponse, disableNotFoundRetries bool, optionals ...interface{}) time.Duration {
	defaultRetryTime := GetDefaultExpectedRetryDuration(response, disableNotFoundRetries)
	if response.Response == nil || response.Response.HTTPResponse() == nil {
		return defaultRetryTime
	}
	e := response.Error
	if len(optionals) > 0 {
		if key, ok := optionals[0].(string); ok {
			switch key {
			case globalvar.DeleteResource:
				switch statusCode := response.Response.HTTPResponse().StatusCode; statusCode {
				case 409:
					if isDisable409Retry, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("disable_409_retry", "false")); isDisable409Retry {
						log.Printf("[ERROR] Resource is in conflict state due to multiple update request: %v", e.Error())
						return 0
					}
					if e := response.Error; e != nil && strings.Contains(e.Error(), "IncorrectState") {
						defaultRetryTime = waasDeleteConflictRetryDuration
					}
				}
			}
		}
	}
	return defaultRetryTime

}

func GetDbHomeRetryDurationFunction(retryTimeout time.Duration) expectedRetryDurationFn {
	return func(response oci_common.OCIOperationResponse, disableNotFoundRetries bool, service string, optionals ...interface{}) time.Duration {
		defaultRetryTime := GetDefaultExpectedRetryDuration(response, disableNotFoundRetries)
		if response.Response == nil || response.Response.HTTPResponse() == nil {
			return defaultRetryTime
		}
		e := response.Error
		switch statusCode := response.Response.HTTPResponse().StatusCode; statusCode {
		case 409:
			if isDisable409Retry, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("disable_409_retry", "false")); isDisable409Retry {
				log.Printf("[ERROR] Resource is in conflict state due to multiple update request: %v", e.Error())
				return 0
			}
			if e := response.Error; e != nil {
				if strings.Contains(e.Error(), "IncorrectState") {
					defaultRetryTime = retryTimeout
				} else if strings.Contains(e.Error(), "InvalidatedRetryToken") {
					defaultRetryTime = 0
				} else {
					defaultRetryTime = LongRetryTime
				}
			}
		}
		return defaultRetryTime
	}
}
