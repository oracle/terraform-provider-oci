// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strings"
	"time"

	oci_common "github.com/oracle/oci-go-sdk/v45/common"
)

type errorTypeEnum string

const (
	ServiceError         errorTypeEnum = "ServiceError"
	TimeoutError         errorTypeEnum = "TimeoutError"
	UnexpectedStateError errorTypeEnum = "UnexpectedStateError"
	WorkRequestError     errorTypeEnum = "WorkRequestError"
)

type customError struct {
	TypeOfError   errorTypeEnum
	ErrorCode     int
	ErrorCodeName string
	Service       string
	Message       string
	OpcRequestID  string
	ResourceOCID  string
	Suggestion    string
	VersionError  string
}

// Create new error format for Terraform output
func newCustomError(sync interface{}, err error) error {
	var tfError customError

	// Service error
	if failure, isServiceError := oci_common.IsServiceError(err); isServiceError {
		tfError = customError{
			TypeOfError:   ServiceError,
			ErrorCode:     failure.GetHTTPStatusCode(),
			ErrorCodeName: failure.GetCode(),
			Message:       failure.GetMessage(),
			OpcRequestID:  failure.GetOpcRequestID(),
			Service:       getServiceName(sync),
		}
	} else if strings.Contains(err.Error(), "timeout while waiting for state") {
		// Timeout error
		tfError = customError{
			TypeOfError:   TimeoutError,
			ErrorCodeName: "Operation Timeout",
			Message:       err.Error(),
			Service:       getServiceName(sync),
		}
		// Unexpected state error
	} else if strings.Contains(err.Error(), "unexpected state") {
		tfError = customError{
			TypeOfError:   UnexpectedStateError,
			ErrorCodeName: "Unexpected LifeCycle state",
			Message:       err.Error(),
			Service:       getServiceName(sync),
			ResourceOCID:  getResourceOCID(sync),
		}
	} else {
		// Terraform error return as is
		return err
	}

	tfError.VersionError = getVersionAndDateError()
	tfError.Suggestion = getSuggestionFromError(tfError)
	return tfError.Error()
}

func (tfE customError) Error() error {
	switch tfE.TypeOfError {
	case ServiceError:
		return fmt.Errorf("%d-%s \n"+
			"%s \n"+
			"Service: %s \n"+
			"Error Message: %s \n"+
			"OPC request ID: %s \n"+
			"Suggestion: %s\n",
			tfE.ErrorCode, tfE.ErrorCodeName, tfE.VersionError, tfE.Service, tfE.Message, tfE.OpcRequestID, tfE.Suggestion)
	case TimeoutError:
		return fmt.Errorf("%s \n"+
			"%s \n"+
			"Service: %s \n"+
			"Error Message: %s \n"+
			"Suggestion: %s\n",
			tfE.ErrorCodeName, tfE.VersionError, tfE.Service, tfE.Message, tfE.Suggestion)
	case UnexpectedStateError:
		return fmt.Errorf("%s \n"+
			"%s \n"+
			"Service: %s \n"+
			"Error Message: %s \n"+
			"Resource OCID: %s \n"+
			"Suggestion: %s\n",
			tfE.ErrorCodeName, tfE.VersionError, tfE.Service, tfE.Message, tfE.ResourceOCID, tfE.Suggestion)
	default:
		return fmt.Errorf(tfE.Message)
	}
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

func handleError(sync interface{}, err error) error {
	if err != nil {
		tfError := newCustomError(sync, err)
		return tfError
	}
	return err
}

func getServiceName(sync interface{}) string {
	syncTypeName := reflect.TypeOf(sync).String()
	if strings.Contains(syncTypeName, "ResourceCrud") {
		serviceName := syncTypeName[strings.Index(syncTypeName, ".")+1 : strings.Index(syncTypeName, "ResourceCrud")]
		return removeDuplicate(serviceName)
	}
	if strings.Contains(syncTypeName, "DataSourcesCrud") {
		serviceName := syncTypeName[strings.Index(syncTypeName, ".")+1 : strings.Index(syncTypeName, "DataSourcesCrud")]
		return removeDuplicate(serviceName)
	}
	if strings.Contains(syncTypeName, "DataSourceCrud") {
		serviceName := syncTypeName[strings.Index(syncTypeName, ".")+1 : strings.Index(syncTypeName, "DataSourceCrud")]
		return removeDuplicate(serviceName)
	}
	log.Printf("[DEBUG] Can't get the service name for: %v", syncTypeName)
	return ""
}

func removeDuplicate(name string) string {
	re := regexp.MustCompile(`[A-Z][^A-Z]*`)
	subMatchAll := re.FindAllString(name, -1)
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range subMatchAll {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return strings.Join(list, " ")
}

// Use to get OCID from refresh state only
func getResourceOCID(sync interface{}) string {
	if syn, ok := sync.(StatefulResource); ok {
		return syn.ID()
	}
	return ""
}

func getVersionAndDateError() string {
	result := fmt.Sprintf("Provider version: %s, released on %s. ", Version, ReleaseDate)
	today := time.Now()
	releaseDate, _ := time.Parse("2006-01-02", ReleaseDate)
	days := today.Sub(releaseDate).Hours() / 24

	if days > 8 {
		versionOld := int(days / 7)
		result += fmt.Sprintf("This provider is %v update(s) behind to current.", versionOld)
	}
	return result
}
