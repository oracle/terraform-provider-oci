// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"strings"
	"time"

	"context"

	"github.com/hashicorp/terraform/helper/resource"
	oci_common "github.com/oracle/oci-go-sdk/common"
	oci_waas "github.com/oracle/oci-go-sdk/waas"
)

const (
	certificateService              = "certificate"
	waasDeleteConflictRetryDuration = 60 * time.Minute
)

var waasServiceExpectedRetryDurationMap = map[string]serviceExpectedRetryDurationFunc{
	certificateService: getWaasCertificateExpectedRetryDuration,
}

func waasWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		//Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		//Make sure we stop on default rules
		if shouldRetry(response, false, "waas", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if okeRes, ok := response.Response.(oci_waas.GetWorkRequestResponse); ok {
			return okeRes.TimeFinished == nil
		}
		return false
	}
}

//waasWaitForWorkRequest custom logic to extract an identifier from a workRequest
func waasWaitForWorkRequest(wId *string, entityType string, action oci_waas.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_waas.WaasClient) (*string, error) {
	retryPolicy := getRetryPolicy(disableFoundRetries, "waas")
	retryPolicy.ShouldRetryOperation = waasWorkRequestShouldRetryFunc(timeout)

	response := oci_waas.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_waas.WorkRequestStatusInProgress),
			string(oci_waas.WorkRequestStatusAccepted),
			string(oci_waas.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_waas.WorkRequestStatusSucceeded),
			string(oci_waas.WorkRequestStatusFailed),
			string(oci_waas.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_waas.GetWorkRequestRequest{
					WorkRequestId: wId,
					RequestMetadata: oci_common.RequestMetadata{
						RetryPolicy: retryPolicy,
					},
				})
			wr := &response.WorkRequest
			return wr, string(wr.Status), err
		},
		Timeout: timeout,
	}
	if _, e := stateConf.WaitForState(); e != nil {
		return nil, e
	}

	var identifier *string
	//The work request response contains an array of objects that finished the operation
	for _, res := range response.Resources {
		if strings.Contains(strings.ToLower(*res.EntityType), entityType) {
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest didn't do all its intended tasks, if the errors is set; so we should check for it
	var workRequestErr error
	if len(response.Errors) > 0 {
		errorMessage := getErrorFromWaasWorkRequest(response)
		workRequestErr = fmt.Errorf("work request did not succeed, workId: %s, entity: %s, action: %s. Message: %s", *wId, entityType, action, errorMessage)
	}

	return identifier, workRequestErr
}

func getErrorFromWaasWorkRequest(response oci_waas.GetWorkRequestResponse) string {
	allErrs := make([]string, 0)
	for _, wrkErr := range response.Errors {
		allErrs = append(allErrs, *wrkErr.Message)
	}
	errorMessage := strings.Join(allErrs, "\n")
	return errorMessage
}

func getWaasExpectedRetryDuration(response oci_common.OCIOperationResponse, disableNotFoundRetries bool, optionals ...string) time.Duration {
	if len(optionals) > 0 {
		if expectedRetryDurationFunc, ok := waasServiceExpectedRetryDurationMap[optionals[0]]; ok {
			return expectedRetryDurationFunc(response, disableNotFoundRetries, optionals[1:]...)
		}
	}
	return getDefaultExpectedRetryDuration(response, disableNotFoundRetries)

}

func getWaasCertificateExpectedRetryDuration(response oci_common.OCIOperationResponse, disableNotFoundRetries bool, optionals ...string) time.Duration {
	defaultRetryTime := getDefaultExpectedRetryDuration(response, disableNotFoundRetries)
	if response.Response == nil || response.Response.HTTPResponse() == nil {
		return defaultRetryTime
	}
	if len(optionals) > 0 {
		switch optionals[0] {
		case deleteResource:
			switch statusCode := response.Response.HTTPResponse().StatusCode; statusCode {
			case 409:
				if e := response.Error; e != nil && strings.Contains(e.Error(), "IncorrectState") {
					defaultRetryTime = waasDeleteConflictRetryDuration
				}
			}
		}
	}
	return defaultRetryTime

}
