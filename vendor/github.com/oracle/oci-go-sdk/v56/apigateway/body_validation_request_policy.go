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
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// BodyValidationRequestPolicy Validate the payload body of the incoming API requests on a specific route.
type BodyValidationRequestPolicy struct {

	// The content of the request body. The key is a media type range (https://tools.ietf.org/html/rfc7231#appendix-D)
	// subset restricted to the following schema
	//     key ::= (
	//           / (  "*" "/" "*" )
	//           / ( type "/" "*" )
	//           / ( type "/" subtype )
	//           )
	// For requests that match multiple keys, only the most specific key is applicable.
	// e.g. `text/plain` overrides `text/*`
	Content map[string]ContentValidation `mandatory:"true" json:"content"`

	// Validation behavior mode.
	// In `ENFORCING` mode, upon a validation failure, the request will be rejected with a 4xx response
	// and not sent to the backend.
	// In `PERMISSIVE` mode, the result of the validation will be exposed as metrics while the request
	// will follow the normal path.
	// `DISABLED` type turns the validation off.
	ValidationMode BodyValidationRequestPolicyValidationModeEnum `mandatory:"false" json:"validationMode,omitempty"`

	// Determines if the request body is required in the request.
	Required *bool `mandatory:"false" json:"required"`
}

func (m BodyValidationRequestPolicy) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *BodyValidationRequestPolicy) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ValidationMode BodyValidationRequestPolicyValidationModeEnum `json:"validationMode"`
		Required       *bool                                         `json:"required"`
		Content        map[string]contentvalidation                  `json:"content"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.ValidationMode = model.ValidationMode

	m.Required = model.Required

	m.Content = make(map[string]ContentValidation)
	for k, v := range model.Content {
		nn, e = v.UnmarshalPolymorphicJSON(v.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Content[k] = nn.(ContentValidation)
		} else {
			m.Content[k] = nil
		}
	}

	return
}

// BodyValidationRequestPolicyValidationModeEnum Enum with underlying type: string
type BodyValidationRequestPolicyValidationModeEnum string

// Set of constants representing the allowable values for BodyValidationRequestPolicyValidationModeEnum
const (
	BodyValidationRequestPolicyValidationModeEnforcing  BodyValidationRequestPolicyValidationModeEnum = "ENFORCING"
	BodyValidationRequestPolicyValidationModePermissive BodyValidationRequestPolicyValidationModeEnum = "PERMISSIVE"
	BodyValidationRequestPolicyValidationModeDisabled   BodyValidationRequestPolicyValidationModeEnum = "DISABLED"
)

var mappingBodyValidationRequestPolicyValidationMode = map[string]BodyValidationRequestPolicyValidationModeEnum{
	"ENFORCING":  BodyValidationRequestPolicyValidationModeEnforcing,
	"PERMISSIVE": BodyValidationRequestPolicyValidationModePermissive,
	"DISABLED":   BodyValidationRequestPolicyValidationModeDisabled,
}

// GetBodyValidationRequestPolicyValidationModeEnumValues Enumerates the set of values for BodyValidationRequestPolicyValidationModeEnum
func GetBodyValidationRequestPolicyValidationModeEnumValues() []BodyValidationRequestPolicyValidationModeEnum {
	values := make([]BodyValidationRequestPolicyValidationModeEnum, 0)
	for _, v := range mappingBodyValidationRequestPolicyValidationMode {
		values = append(values, v)
	}
	return values
}
