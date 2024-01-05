// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ValidationRequestPolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingValidationRequestPolicyValidationModeEnum(string(m.ValidationMode)); !ok && m.ValidationMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ValidationMode: %s. Supported values are: %s.", m.ValidationMode, strings.Join(GetValidationRequestPolicyValidationModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ValidationRequestPolicyValidationModeEnum Enum with underlying type: string
type ValidationRequestPolicyValidationModeEnum string

// Set of constants representing the allowable values for ValidationRequestPolicyValidationModeEnum
const (
	ValidationRequestPolicyValidationModeEnforcing  ValidationRequestPolicyValidationModeEnum = "ENFORCING"
	ValidationRequestPolicyValidationModePermissive ValidationRequestPolicyValidationModeEnum = "PERMISSIVE"
	ValidationRequestPolicyValidationModeDisabled   ValidationRequestPolicyValidationModeEnum = "DISABLED"
)

var mappingValidationRequestPolicyValidationModeEnum = map[string]ValidationRequestPolicyValidationModeEnum{
	"ENFORCING":  ValidationRequestPolicyValidationModeEnforcing,
	"PERMISSIVE": ValidationRequestPolicyValidationModePermissive,
	"DISABLED":   ValidationRequestPolicyValidationModeDisabled,
}

var mappingValidationRequestPolicyValidationModeEnumLowerCase = map[string]ValidationRequestPolicyValidationModeEnum{
	"enforcing":  ValidationRequestPolicyValidationModeEnforcing,
	"permissive": ValidationRequestPolicyValidationModePermissive,
	"disabled":   ValidationRequestPolicyValidationModeDisabled,
}

// GetValidationRequestPolicyValidationModeEnumValues Enumerates the set of values for ValidationRequestPolicyValidationModeEnum
func GetValidationRequestPolicyValidationModeEnumValues() []ValidationRequestPolicyValidationModeEnum {
	values := make([]ValidationRequestPolicyValidationModeEnum, 0)
	for _, v := range mappingValidationRequestPolicyValidationModeEnum {
		values = append(values, v)
	}
	return values
}

// GetValidationRequestPolicyValidationModeEnumStringValues Enumerates the set of values in String for ValidationRequestPolicyValidationModeEnum
func GetValidationRequestPolicyValidationModeEnumStringValues() []string {
	return []string{
		"ENFORCING",
		"PERMISSIVE",
		"DISABLED",
	}
}

// GetMappingValidationRequestPolicyValidationModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingValidationRequestPolicyValidationModeEnum(val string) (ValidationRequestPolicyValidationModeEnum, bool) {
	enum, ok := mappingValidationRequestPolicyValidationModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
