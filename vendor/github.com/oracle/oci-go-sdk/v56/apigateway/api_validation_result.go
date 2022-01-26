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

// ApiValidationResultResultEnum Enum with underlying type: string
type ApiValidationResultResultEnum string

// Set of constants representing the allowable values for ApiValidationResultResultEnum
const (
	ApiValidationResultResultError   ApiValidationResultResultEnum = "ERROR"
	ApiValidationResultResultWarning ApiValidationResultResultEnum = "WARNING"
	ApiValidationResultResultOk      ApiValidationResultResultEnum = "OK"
	ApiValidationResultResultFailed  ApiValidationResultResultEnum = "FAILED"
)

var mappingApiValidationResultResult = map[string]ApiValidationResultResultEnum{
	"ERROR":   ApiValidationResultResultError,
	"WARNING": ApiValidationResultResultWarning,
	"OK":      ApiValidationResultResultOk,
	"FAILED":  ApiValidationResultResultFailed,
}

// GetApiValidationResultResultEnumValues Enumerates the set of values for ApiValidationResultResultEnum
func GetApiValidationResultResultEnumValues() []ApiValidationResultResultEnum {
	values := make([]ApiValidationResultResultEnum, 0)
	for _, v := range mappingApiValidationResultResult {
		values = append(values, v)
	}
	return values
}
