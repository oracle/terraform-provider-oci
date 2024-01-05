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
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ValidationFailurePolicy Policy for defining behaviour on validation failure.
type ValidationFailurePolicy interface {
}

type validationfailurepolicy struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *validationfailurepolicy) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalervalidationfailurepolicy validationfailurepolicy
	s := struct {
		Model Unmarshalervalidationfailurepolicy
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *validationfailurepolicy) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "MODIFY_RESPONSE":
		mm := ModifyResponseValidationFailurePolicy{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OAUTH2":
		mm := OAuth2ResponseValidationFailurePolicy{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ValidationFailurePolicy: %s.", m.Type)
		return *m, nil
	}
}

func (m validationfailurepolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m validationfailurepolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ValidationFailurePolicyTypeEnum Enum with underlying type: string
type ValidationFailurePolicyTypeEnum string

// Set of constants representing the allowable values for ValidationFailurePolicyTypeEnum
const (
	ValidationFailurePolicyTypeModifyResponse ValidationFailurePolicyTypeEnum = "MODIFY_RESPONSE"
	ValidationFailurePolicyTypeOauth2         ValidationFailurePolicyTypeEnum = "OAUTH2"
)

var mappingValidationFailurePolicyTypeEnum = map[string]ValidationFailurePolicyTypeEnum{
	"MODIFY_RESPONSE": ValidationFailurePolicyTypeModifyResponse,
	"OAUTH2":          ValidationFailurePolicyTypeOauth2,
}

var mappingValidationFailurePolicyTypeEnumLowerCase = map[string]ValidationFailurePolicyTypeEnum{
	"modify_response": ValidationFailurePolicyTypeModifyResponse,
	"oauth2":          ValidationFailurePolicyTypeOauth2,
}

// GetValidationFailurePolicyTypeEnumValues Enumerates the set of values for ValidationFailurePolicyTypeEnum
func GetValidationFailurePolicyTypeEnumValues() []ValidationFailurePolicyTypeEnum {
	values := make([]ValidationFailurePolicyTypeEnum, 0)
	for _, v := range mappingValidationFailurePolicyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetValidationFailurePolicyTypeEnumStringValues Enumerates the set of values in String for ValidationFailurePolicyTypeEnum
func GetValidationFailurePolicyTypeEnumStringValues() []string {
	return []string{
		"MODIFY_RESPONSE",
		"OAUTH2",
	}
}

// GetMappingValidationFailurePolicyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingValidationFailurePolicyTypeEnum(val string) (ValidationFailurePolicyTypeEnum, bool) {
	enum, ok := mappingValidationFailurePolicyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
