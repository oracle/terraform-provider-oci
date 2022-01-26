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

// QueryParameterValidationRequestPolicy Validate the URL query parameters on the incoming API requests on a specific route.
type QueryParameterValidationRequestPolicy struct {

	// Validation behavior mode.
	// In `ENFORCING` mode, upon a validation failure, the request will be rejected with a 4xx response
	// and not sent to the backend.
	// In `PERMISSIVE` mode, the result of the validation will be exposed as metrics while the request
	// will follow the normal path.
	// `DISABLED` type turns the validation off.
	ValidationMode QueryParameterValidationRequestPolicyValidationModeEnum `mandatory:"false" json:"validationMode,omitempty"`

	Parameters []QueryParameterValidationItem `mandatory:"false" json:"parameters"`
}

func (m QueryParameterValidationRequestPolicy) String() string {
	return common.PointerString(m)
}

// QueryParameterValidationRequestPolicyValidationModeEnum Enum with underlying type: string
type QueryParameterValidationRequestPolicyValidationModeEnum string

// Set of constants representing the allowable values for QueryParameterValidationRequestPolicyValidationModeEnum
const (
	QueryParameterValidationRequestPolicyValidationModeEnforcing  QueryParameterValidationRequestPolicyValidationModeEnum = "ENFORCING"
	QueryParameterValidationRequestPolicyValidationModePermissive QueryParameterValidationRequestPolicyValidationModeEnum = "PERMISSIVE"
	QueryParameterValidationRequestPolicyValidationModeDisabled   QueryParameterValidationRequestPolicyValidationModeEnum = "DISABLED"
)

var mappingQueryParameterValidationRequestPolicyValidationMode = map[string]QueryParameterValidationRequestPolicyValidationModeEnum{
	"ENFORCING":  QueryParameterValidationRequestPolicyValidationModeEnforcing,
	"PERMISSIVE": QueryParameterValidationRequestPolicyValidationModePermissive,
	"DISABLED":   QueryParameterValidationRequestPolicyValidationModeDisabled,
}

// GetQueryParameterValidationRequestPolicyValidationModeEnumValues Enumerates the set of values for QueryParameterValidationRequestPolicyValidationModeEnum
func GetQueryParameterValidationRequestPolicyValidationModeEnumValues() []QueryParameterValidationRequestPolicyValidationModeEnum {
	values := make([]QueryParameterValidationRequestPolicyValidationModeEnum, 0)
	for _, v := range mappingQueryParameterValidationRequestPolicyValidationMode {
		values = append(values, v)
	}
	return values
}
