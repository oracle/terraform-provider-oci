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

// SdkLanguageOptionalParameters List of additional applicable parameters for any given target language.
type SdkLanguageOptionalParameters struct {

	// Name of the parameter.
	ParamName *string `mandatory:"true" json:"paramName"`

	// Display name of the parameter.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Description for the parameter.
	Description *string `mandatory:"false" json:"description"`

	// Information on whether the parameter is required or not.
	IsRequired *bool `mandatory:"false" json:"isRequired"`

	// Maximum size as input value for this parameter.
	MaxSize *float32 `mandatory:"false" json:"maxSize"`

	// The input type for this param.
	// - Input type is ENUM when only specific list of input strings are allowed.
	// - Input type is EMAIL when input type is an email ID.
	// - Input type is URI when input type is an URI.
	// - Input type is STRING in all other cases.
	InputType SdkLanguageOptionalParametersInputTypeEnum `mandatory:"false" json:"inputType,omitempty"`

	// List of allowed input values.
	// Example: `[{"name": "name1", "description": "description1"}, ...]`
	AllowedValues []SdkLanguageOptionalParametersAllowedValue `mandatory:"false" json:"allowedValues"`
}

func (m SdkLanguageOptionalParameters) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SdkLanguageOptionalParameters) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSdkLanguageOptionalParametersInputTypeEnum(string(m.InputType)); !ok && m.InputType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InputType: %s. Supported values are: %s.", m.InputType, strings.Join(GetSdkLanguageOptionalParametersInputTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SdkLanguageOptionalParametersInputTypeEnum Enum with underlying type: string
type SdkLanguageOptionalParametersInputTypeEnum string

// Set of constants representing the allowable values for SdkLanguageOptionalParametersInputTypeEnum
const (
	SdkLanguageOptionalParametersInputTypeEnumvalue SdkLanguageOptionalParametersInputTypeEnum = "ENUM"
	SdkLanguageOptionalParametersInputTypeEmail     SdkLanguageOptionalParametersInputTypeEnum = "EMAIL"
	SdkLanguageOptionalParametersInputTypeUri       SdkLanguageOptionalParametersInputTypeEnum = "URI"
	SdkLanguageOptionalParametersInputTypeString    SdkLanguageOptionalParametersInputTypeEnum = "STRING"
)

var mappingSdkLanguageOptionalParametersInputTypeEnum = map[string]SdkLanguageOptionalParametersInputTypeEnum{
	"ENUM":   SdkLanguageOptionalParametersInputTypeEnumvalue,
	"EMAIL":  SdkLanguageOptionalParametersInputTypeEmail,
	"URI":    SdkLanguageOptionalParametersInputTypeUri,
	"STRING": SdkLanguageOptionalParametersInputTypeString,
}

// GetSdkLanguageOptionalParametersInputTypeEnumValues Enumerates the set of values for SdkLanguageOptionalParametersInputTypeEnum
func GetSdkLanguageOptionalParametersInputTypeEnumValues() []SdkLanguageOptionalParametersInputTypeEnum {
	values := make([]SdkLanguageOptionalParametersInputTypeEnum, 0)
	for _, v := range mappingSdkLanguageOptionalParametersInputTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSdkLanguageOptionalParametersInputTypeEnumStringValues Enumerates the set of values in String for SdkLanguageOptionalParametersInputTypeEnum
func GetSdkLanguageOptionalParametersInputTypeEnumStringValues() []string {
	return []string{
		"ENUM",
		"EMAIL",
		"URI",
		"STRING",
	}
}

// GetMappingSdkLanguageOptionalParametersInputTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSdkLanguageOptionalParametersInputTypeEnum(val string) (SdkLanguageOptionalParametersInputTypeEnum, bool) {
	mappingSdkLanguageOptionalParametersInputTypeEnumIgnoreCase := make(map[string]SdkLanguageOptionalParametersInputTypeEnum)
	for k, v := range mappingSdkLanguageOptionalParametersInputTypeEnum {
		mappingSdkLanguageOptionalParametersInputTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingSdkLanguageOptionalParametersInputTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
