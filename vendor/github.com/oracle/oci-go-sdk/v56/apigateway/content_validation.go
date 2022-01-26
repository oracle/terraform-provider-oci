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
		return *m, nil
	}
}

func (m contentvalidation) String() string {
	return common.PointerString(m)
}

// ContentValidationValidationTypeEnum Enum with underlying type: string
type ContentValidationValidationTypeEnum string

// Set of constants representing the allowable values for ContentValidationValidationTypeEnum
const (
	ContentValidationValidationTypeNone ContentValidationValidationTypeEnum = "NONE"
)

var mappingContentValidationValidationType = map[string]ContentValidationValidationTypeEnum{
	"NONE": ContentValidationValidationTypeNone,
}

// GetContentValidationValidationTypeEnumValues Enumerates the set of values for ContentValidationValidationTypeEnum
func GetContentValidationValidationTypeEnumValues() []ContentValidationValidationTypeEnum {
	values := make([]ContentValidationValidationTypeEnum, 0)
	for _, v := range mappingContentValidationValidationType {
		values = append(values, v)
	}
	return values
}
