// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Load Balancing API
//
// API for the Load Balancing service. Use this API to manage load balancers, backend sets, and related items. For more
// information, see Overview of Load Balancing (https://docs.cloud.oracle.com/iaas/Content/Balance/Concepts/balanceoverview.htm).
//

package loadbalancer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// WorkRequestError An object returned in the event of a work request error.
type WorkRequestError struct {
	ErrorCode WorkRequestErrorErrorCodeEnum `mandatory:"true" json:"errorCode"`

	// A human-readable error string.
	Message *string `mandatory:"true" json:"message"`
}

func (m WorkRequestError) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m WorkRequestError) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingWorkRequestErrorErrorCodeEnum(string(m.ErrorCode)); !ok && m.ErrorCode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ErrorCode: %s. Supported values are: %s.", m.ErrorCode, strings.Join(GetWorkRequestErrorErrorCodeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// WorkRequestErrorErrorCodeEnum Enum with underlying type: string
type WorkRequestErrorErrorCodeEnum string

// Set of constants representing the allowable values for WorkRequestErrorErrorCodeEnum
const (
	WorkRequestErrorErrorCodeBadInput      WorkRequestErrorErrorCodeEnum = "BAD_INPUT"
	WorkRequestErrorErrorCodeInternalError WorkRequestErrorErrorCodeEnum = "INTERNAL_ERROR"
)

var mappingWorkRequestErrorErrorCodeEnum = map[string]WorkRequestErrorErrorCodeEnum{
	"BAD_INPUT":      WorkRequestErrorErrorCodeBadInput,
	"INTERNAL_ERROR": WorkRequestErrorErrorCodeInternalError,
}

var mappingWorkRequestErrorErrorCodeEnumLowerCase = map[string]WorkRequestErrorErrorCodeEnum{
	"bad_input":      WorkRequestErrorErrorCodeBadInput,
	"internal_error": WorkRequestErrorErrorCodeInternalError,
}

// GetWorkRequestErrorErrorCodeEnumValues Enumerates the set of values for WorkRequestErrorErrorCodeEnum
func GetWorkRequestErrorErrorCodeEnumValues() []WorkRequestErrorErrorCodeEnum {
	values := make([]WorkRequestErrorErrorCodeEnum, 0)
	for _, v := range mappingWorkRequestErrorErrorCodeEnum {
		values = append(values, v)
	}
	return values
}

// GetWorkRequestErrorErrorCodeEnumStringValues Enumerates the set of values in String for WorkRequestErrorErrorCodeEnum
func GetWorkRequestErrorErrorCodeEnumStringValues() []string {
	return []string{
		"BAD_INPUT",
		"INTERNAL_ERROR",
	}
}

// GetMappingWorkRequestErrorErrorCodeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkRequestErrorErrorCodeEnum(val string) (WorkRequestErrorErrorCodeEnum, bool) {
	enum, ok := mappingWorkRequestErrorErrorCodeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
