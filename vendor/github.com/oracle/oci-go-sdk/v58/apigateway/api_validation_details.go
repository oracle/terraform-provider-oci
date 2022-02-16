// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// API Gateway API
//
// API for the API Gateway service. Use this API to manage gateways, deployments, and related items.
// For more information, see
// Overview of API Gateway (https://docs.cloud.oracle.com/iaas/Content/APIGateway/Concepts/apigatewayoverview.htm).
//

package apigateway

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// ApiValidationDetails Detail of an error or warning.
type ApiValidationDetails struct {

	// Name of the validation.
	Name *string `mandatory:"true" json:"name"`

	// Result of the validation.
	Result ApiValidationDetailsResultEnum `mandatory:"true" json:"result"`

	// Details of validation.
	Details []ApiValidationDetail `mandatory:"false" json:"details"`
}

func (m ApiValidationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ApiValidationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingApiValidationDetailsResultEnum(string(m.Result)); !ok && m.Result != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Result: %s. Supported values are: %s.", m.Result, strings.Join(GetApiValidationDetailsResultEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ApiValidationDetailsResultEnum Enum with underlying type: string
type ApiValidationDetailsResultEnum string

// Set of constants representing the allowable values for ApiValidationDetailsResultEnum
const (
	ApiValidationDetailsResultError   ApiValidationDetailsResultEnum = "ERROR"
	ApiValidationDetailsResultWarning ApiValidationDetailsResultEnum = "WARNING"
	ApiValidationDetailsResultOk      ApiValidationDetailsResultEnum = "OK"
	ApiValidationDetailsResultFailed  ApiValidationDetailsResultEnum = "FAILED"
)

var mappingApiValidationDetailsResultEnum = map[string]ApiValidationDetailsResultEnum{
	"ERROR":   ApiValidationDetailsResultError,
	"WARNING": ApiValidationDetailsResultWarning,
	"OK":      ApiValidationDetailsResultOk,
	"FAILED":  ApiValidationDetailsResultFailed,
}

// GetApiValidationDetailsResultEnumValues Enumerates the set of values for ApiValidationDetailsResultEnum
func GetApiValidationDetailsResultEnumValues() []ApiValidationDetailsResultEnum {
	values := make([]ApiValidationDetailsResultEnum, 0)
	for _, v := range mappingApiValidationDetailsResultEnum {
		values = append(values, v)
	}
	return values
}

// GetApiValidationDetailsResultEnumStringValues Enumerates the set of values in String for ApiValidationDetailsResultEnum
func GetApiValidationDetailsResultEnumStringValues() []string {
	return []string{
		"ERROR",
		"WARNING",
		"OK",
		"FAILED",
	}
}

// GetMappingApiValidationDetailsResultEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingApiValidationDetailsResultEnum(val string) (ApiValidationDetailsResultEnum, bool) {
	mappingApiValidationDetailsResultEnumIgnoreCase := make(map[string]ApiValidationDetailsResultEnum)
	for k, v := range mappingApiValidationDetailsResultEnum {
		mappingApiValidationDetailsResultEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingApiValidationDetailsResultEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
