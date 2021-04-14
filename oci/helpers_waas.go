// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"strings"
	"time"

	oci_common "github.com/oracle/oci-go-sdk/v39/common"
)

const (
	certificateService              = "certificate"
	waasDeleteConflictRetryDuration = 60 * time.Minute
)

var waasServiceExpectedRetryDurationMap = map[string]serviceExpectedRetryDurationFunc{
	certificateService: getWaasCertificateExpectedRetryDuration,
}

func getWaasExpectedRetryDuration(response oci_common.OCIOperationResponse, disableNotFoundRetries bool, optionals ...interface{}) time.Duration {
	if len(optionals) > 0 {
		if key, ok := optionals[0].(string); ok {
			if expectedRetryDurationFunc, ok := waasServiceExpectedRetryDurationMap[key]; ok {
				return expectedRetryDurationFunc(response, disableNotFoundRetries, optionals[1:]...)
			}
		}
	}
	return getDefaultExpectedRetryDuration(response, disableNotFoundRetries)

}

func getWaasCertificateExpectedRetryDuration(response oci_common.OCIOperationResponse, disableNotFoundRetries bool, optionals ...interface{}) time.Duration {
	defaultRetryTime := getDefaultExpectedRetryDuration(response, disableNotFoundRetries)
	if response.Response == nil || response.Response.HTTPResponse() == nil {
		return defaultRetryTime
	}
	if len(optionals) > 0 {
		if key, ok := optionals[0].(string); ok {
			switch key {
			case deleteResource:
				switch statusCode := response.Response.HTTPResponse().StatusCode; statusCode {
				case 409:
					if e := response.Error; e != nil && strings.Contains(e.Error(), "IncorrectState") {
						defaultRetryTime = waasDeleteConflictRetryDuration
					}
				}
			}
		}
	}
	return defaultRetryTime

}
