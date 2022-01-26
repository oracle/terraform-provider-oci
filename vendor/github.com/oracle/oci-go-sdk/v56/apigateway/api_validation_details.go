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
	"github.com/oracle/oci-go-sdk/v56/common"
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

// ApiValidationDetailsResultEnum Enum with underlying type: string
type ApiValidationDetailsResultEnum string

// Set of constants representing the allowable values for ApiValidationDetailsResultEnum
const (
	ApiValidationDetailsResultError   ApiValidationDetailsResultEnum = "ERROR"
	ApiValidationDetailsResultWarning ApiValidationDetailsResultEnum = "WARNING"
	ApiValidationDetailsResultOk      ApiValidationDetailsResultEnum = "OK"
	ApiValidationDetailsResultFailed  ApiValidationDetailsResultEnum = "FAILED"
)

var mappingApiValidationDetailsResult = map[string]ApiValidationDetailsResultEnum{
	"ERROR":   ApiValidationDetailsResultError,
	"WARNING": ApiValidationDetailsResultWarning,
	"OK":      ApiValidationDetailsResultOk,
	"FAILED":  ApiValidationDetailsResultFailed,
}

// GetApiValidationDetailsResultEnumValues Enumerates the set of values for ApiValidationDetailsResultEnum
func GetApiValidationDetailsResultEnumValues() []ApiValidationDetailsResultEnum {
	values := make([]ApiValidationDetailsResultEnum, 0)
	for _, v := range mappingApiValidationDetailsResult {
		values = append(values, v)
	}
	return values
}
