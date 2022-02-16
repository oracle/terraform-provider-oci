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

// ApiValidationResult The result of single validation.
type ApiValidationResult struct {

	// Name of the validation.
	Name *string `mandatory:"true" json:"name"`

	// Result of the validation.
	Result ApiValidationResultResultEnum `mandatory:"true" json:"result"`
}

func (m ApiValidationResult) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ApiValidationResult) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingApiValidationResultResultEnum(string(m.Result)); !ok && m.Result != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Result: %s. Supported values are: %s.", m.Result, strings.Join(GetApiValidationResultResultEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ApiValidationResultResultEnum Enum with underlying type: string
type ApiValidationResultResultEnum string

// Set of constants representing the allowable values for ApiValidationResultResultEnum
const (
	ApiValidationResultResultError   ApiValidationResultResultEnum = "ERROR"
	ApiValidationResultResultWarning ApiValidationResultResultEnum = "WARNING"
	ApiValidationResultResultOk      ApiValidationResultResultEnum = "OK"
	ApiValidationResultResultFailed  ApiValidationResultResultEnum = "FAILED"
)

var mappingApiValidationResultResultEnum = map[string]ApiValidationResultResultEnum{
	"ERROR":   ApiValidationResultResultError,
	"WARNING": ApiValidationResultResultWarning,
	"OK":      ApiValidationResultResultOk,
	"FAILED":  ApiValidationResultResultFailed,
}

// GetApiValidationResultResultEnumValues Enumerates the set of values for ApiValidationResultResultEnum
func GetApiValidationResultResultEnumValues() []ApiValidationResultResultEnum {
	values := make([]ApiValidationResultResultEnum, 0)
	for _, v := range mappingApiValidationResultResultEnum {
		values = append(values, v)
	}
	return values
}

// GetApiValidationResultResultEnumStringValues Enumerates the set of values in String for ApiValidationResultResultEnum
func GetApiValidationResultResultEnumStringValues() []string {
	return []string{
		"ERROR",
		"WARNING",
		"OK",
		"FAILED",
	}
}

// GetMappingApiValidationResultResultEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingApiValidationResultResultEnum(val string) (ApiValidationResultResultEnum, bool) {
	mappingApiValidationResultResultEnumIgnoreCase := make(map[string]ApiValidationResultResultEnum)
	for k, v := range mappingApiValidationResultResultEnum {
		mappingApiValidationResultResultEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingApiValidationResultResultEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
