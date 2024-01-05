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

// ContentValidation Content validation properties.
type ContentValidation interface {
}

type contentvalidation struct {
	JsonData       []byte
	ValidationType string `json:"validationType"`
}

// UnmarshalJSON unmarshals json
func (m *contentvalidation) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercontentvalidation contentvalidation
	s := struct {
		Model Unmarshalercontentvalidation
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ValidationType = s.Model.ValidationType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *contentvalidation) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ValidationType {
	case "NONE":
		mm := NoContentValidation{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ContentValidation: %s.", m.ValidationType)
		return *m, nil
	}
}

func (m contentvalidation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m contentvalidation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ContentValidationValidationTypeEnum Enum with underlying type: string
type ContentValidationValidationTypeEnum string

// Set of constants representing the allowable values for ContentValidationValidationTypeEnum
const (
	ContentValidationValidationTypeNone ContentValidationValidationTypeEnum = "NONE"
)

var mappingContentValidationValidationTypeEnum = map[string]ContentValidationValidationTypeEnum{
	"NONE": ContentValidationValidationTypeNone,
}

var mappingContentValidationValidationTypeEnumLowerCase = map[string]ContentValidationValidationTypeEnum{
	"none": ContentValidationValidationTypeNone,
}

// GetContentValidationValidationTypeEnumValues Enumerates the set of values for ContentValidationValidationTypeEnum
func GetContentValidationValidationTypeEnumValues() []ContentValidationValidationTypeEnum {
	values := make([]ContentValidationValidationTypeEnum, 0)
	for _, v := range mappingContentValidationValidationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetContentValidationValidationTypeEnumStringValues Enumerates the set of values in String for ContentValidationValidationTypeEnum
func GetContentValidationValidationTypeEnumStringValues() []string {
	return []string{
		"NONE",
	}
}

// GetMappingContentValidationValidationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingContentValidationValidationTypeEnum(val string) (ContentValidationValidationTypeEnum, bool) {
	enum, ok := mappingContentValidationValidationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
