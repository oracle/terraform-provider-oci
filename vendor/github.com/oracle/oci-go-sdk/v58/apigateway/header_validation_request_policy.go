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

// HeaderValidationRequestPolicy Validate the HTTP headers on the incoming API requests on a specific route.
type HeaderValidationRequestPolicy struct {

	// Validation behavior mode.
	// In `ENFORCING` mode, upon a validation failure, the request will be rejected with a 4xx response
	// and not sent to the backend.
	// In `PERMISSIVE` mode, the result of the validation will be exposed as metrics while the request
	// will follow the normal path.
	// `DISABLED` type turns the validation off.
	ValidationMode HeaderValidationRequestPolicyValidationModeEnum `mandatory:"false" json:"validationMode,omitempty"`

	Headers []HeaderValidationItem `mandatory:"false" json:"headers"`
}

func (m HeaderValidationRequestPolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HeaderValidationRequestPolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingHeaderValidationRequestPolicyValidationModeEnum(string(m.ValidationMode)); !ok && m.ValidationMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ValidationMode: %s. Supported values are: %s.", m.ValidationMode, strings.Join(GetHeaderValidationRequestPolicyValidationModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// HeaderValidationRequestPolicyValidationModeEnum Enum with underlying type: string
type HeaderValidationRequestPolicyValidationModeEnum string

// Set of constants representing the allowable values for HeaderValidationRequestPolicyValidationModeEnum
const (
	HeaderValidationRequestPolicyValidationModeEnforcing  HeaderValidationRequestPolicyValidationModeEnum = "ENFORCING"
	HeaderValidationRequestPolicyValidationModePermissive HeaderValidationRequestPolicyValidationModeEnum = "PERMISSIVE"
	HeaderValidationRequestPolicyValidationModeDisabled   HeaderValidationRequestPolicyValidationModeEnum = "DISABLED"
)

var mappingHeaderValidationRequestPolicyValidationModeEnum = map[string]HeaderValidationRequestPolicyValidationModeEnum{
	"ENFORCING":  HeaderValidationRequestPolicyValidationModeEnforcing,
	"PERMISSIVE": HeaderValidationRequestPolicyValidationModePermissive,
	"DISABLED":   HeaderValidationRequestPolicyValidationModeDisabled,
}

// GetHeaderValidationRequestPolicyValidationModeEnumValues Enumerates the set of values for HeaderValidationRequestPolicyValidationModeEnum
func GetHeaderValidationRequestPolicyValidationModeEnumValues() []HeaderValidationRequestPolicyValidationModeEnum {
	values := make([]HeaderValidationRequestPolicyValidationModeEnum, 0)
	for _, v := range mappingHeaderValidationRequestPolicyValidationModeEnum {
		values = append(values, v)
	}
	return values
}

// GetHeaderValidationRequestPolicyValidationModeEnumStringValues Enumerates the set of values in String for HeaderValidationRequestPolicyValidationModeEnum
func GetHeaderValidationRequestPolicyValidationModeEnumStringValues() []string {
	return []string{
		"ENFORCING",
		"PERMISSIVE",
		"DISABLED",
	}
}

// GetMappingHeaderValidationRequestPolicyValidationModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHeaderValidationRequestPolicyValidationModeEnum(val string) (HeaderValidationRequestPolicyValidationModeEnum, bool) {
	mappingHeaderValidationRequestPolicyValidationModeEnumIgnoreCase := make(map[string]HeaderValidationRequestPolicyValidationModeEnum)
	for k, v := range mappingHeaderValidationRequestPolicyValidationModeEnum {
		mappingHeaderValidationRequestPolicyValidationModeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingHeaderValidationRequestPolicyValidationModeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
