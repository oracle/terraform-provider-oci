// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"strconv"
	"time"

	oci_common "github.com/oracle/oci-go-sdk/v40/common"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

func kmsGetRetryPolicy(disableNotFoundRetries bool, service string, optionals ...interface{}) *oci_common.RetryPolicy {
	startTime := time.Now()
	retryPolicy := &oci_common.RetryPolicy{
		MaximumNumberAttempts: 0,
		ShouldRetryOperation: func(response oci_common.OCIOperationResponse) bool {
			return shouldRetry(response, disableNotFoundRetries, service, startTime, optionals...)
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
	defaultRetryTime := getRetryBackoffDuration(response, disableNotFoundRetries, "kms", startTime, optionals...)
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
