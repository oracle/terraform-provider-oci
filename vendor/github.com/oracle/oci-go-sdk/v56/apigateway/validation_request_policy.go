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

// ValidationRequestPolicy Top-level validation policy mixin (not directly used).
type ValidationRequestPolicy struct {

	// Validation behavior mode.
	// In `ENFORCING` mode, upon a validation failure, the request will be rejected with a 4xx response
	// and not sent to the backend.
	// In `PERMISSIVE` mode, the result of the validation will be exposed as metrics while the request
	// will follow the normal path.
	// `DISABLED` type turns the validation off.
	ValidationMode ValidationRequestPolicyValidationModeEnum `mandatory:"false" json:"validationMode,omitempty"`
}

func (m ValidationRequestPolicy) String() string {
	return common.PointerString(m)
}

// ValidationRequestPolicyValidationModeEnum Enum with underlying type: string
type ValidationRequestPolicyValidationModeEnum string

// Set of constants representing the allowable values for ValidationRequestPolicyValidationModeEnum
const (
	ValidationRequestPolicyValidationModeEnforcing  ValidationRequestPolicyValidationModeEnum = "ENFORCING"
	ValidationRequestPolicyValidationModePermissive ValidationRequestPolicyValidationModeEnum = "PERMISSIVE"
	ValidationRequestPolicyValidationModeDisabled   ValidationRequestPolicyValidationModeEnum = "DISABLED"
)

var mappingValidationRequestPolicyValidationMode = map[string]ValidationRequestPolicyValidationModeEnum{
	"ENFORCING":  ValidationRequestPolicyValidationModeEnforcing,
	"PERMISSIVE": ValidationRequestPolicyValidationModePermissive,
	"DISABLED":   ValidationRequestPolicyValidationModeDisabled,
}

// GetValidationRequestPolicyValidationModeEnumValues Enumerates the set of values for ValidationRequestPolicyValidationModeEnum
func GetValidationRequestPolicyValidationModeEnumValues() []ValidationRequestPolicyValidationModeEnum {
	values := make([]ValidationRequestPolicyValidationModeEnum, 0)
	for _, v := range mappingValidationRequestPolicyValidationMode {
		values = append(values, v)
	}
	return values
}
