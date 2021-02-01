// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	oci_common "github.com/oracle/oci-go-sdk/v33/common"
)

type customError struct {
	isServiceError bool
	ErrorCode      int
	ErrorCodeName  string
	Service        string
	Message        string
	OpcRequestID   string
	Suggestion     string
}

// Create new error format for Terraform output
func newCustomError(sync interface{}, err error) error {
	var tfError customError
	// Service error
	if failure, isServiceError := oci_common.IsServiceError(err); isServiceError {
		tfError = customError{
			isServiceError: true,
			ErrorCode:      failure.GetHTTPStatusCode(),
			ErrorCodeName:  failure.GetCode(),
			Message:        failure.GetMessage(),
			OpcRequestID:   failure.GetOpcRequestID(),
			Service:        getServiceName(sync),
		}
	} else {
		// Terraform error return as is
		tfError.isServiceError = false
		return err
	}
	tfError.Suggestion = getSuggestionFromError(tfError)
	return tfError.Error()
}

func (tfE customError) Error() error {
	return fmt.Errorf("%d-%s \n"+
		"Service: %s \n"+
		"Error Message: %s \n"+
		"OPC request ID: %s \n"+
		"Suggestion: %s\n",
		tfE.ErrorCode, tfE.ErrorCodeName, tfE.Service, tfE.Message, tfE.OpcRequestID, tfE.Suggestion)
}

func handleMissingResourceError(sync ResourceVoider, err *error) {

	if err != nil {
		// patch till OCE service returns correct error response code for invalid auth token
		if strings.Contains((*err).Error(), "IDCS token validation has failed") {
			return
		}

		if strings.Contains((*err).Error(), "does not exist") ||
			strings.Contains((*err).Error(), " not present in ") ||
			strings.Contains((*err).Error(), "not found") ||
			(strings.Contains((*err).Error(), "Load balancer") && strings.Contains((*err).Error(), " has no ")) ||
			strings.Contains(strings.ToLower((*err).Error()), "status code: 404") { // status code: 404 is not enough because the load balancer error responses don't include it for some reason
			log.Println("[DEBUG] Object does not exist, voiding resource and nullifying error")
			if sync != nil {
				sync.VoidState()
			}
			*err = nil
		}
	}
}

func handleServiceError(sync interface{}, err error) error {
	tfError := newCustomError(sync, err)
	return tfError
}

func getServiceName(sync interface{}) string {
	syncTypeName := reflect.TypeOf(sync).String()
	if strings.Contains(syncTypeName, "ResourceCrud") {
		return syncTypeName[strings.Index(syncTypeName, ".")+1 : strings.Index(syncTypeName, "ResourceCrud")]
	}
	return ""
}
