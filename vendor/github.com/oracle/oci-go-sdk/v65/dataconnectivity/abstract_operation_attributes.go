// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Connectivity Management API
//
// Use the Data Connectivity Management Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataconnectivity

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AbstractOperationAttributes The operation attributes
type AbstractOperationAttributes interface {
}

type abstractoperationattributes struct {
	JsonData  []byte
	ModelType string `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *abstractoperationattributes) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerabstractoperationattributes abstractoperationattributes
	s := struct {
		Model Unmarshalerabstractoperationattributes
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *abstractoperationattributes) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "GENERIC_REST_API_ATTRIBUTES":
		mm := GenericRestApiAttributes{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for AbstractOperationAttributes: %s.", m.ModelType)
		return *m, nil
	}
}

func (m abstractoperationattributes) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m abstractoperationattributes) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AbstractOperationAttributesModelTypeEnum Enum with underlying type: string
type AbstractOperationAttributesModelTypeEnum string

// Set of constants representing the allowable values for AbstractOperationAttributesModelTypeEnum
const (
	AbstractOperationAttributesModelTypeGenericRestApiAttributes AbstractOperationAttributesModelTypeEnum = "GENERIC_REST_API_ATTRIBUTES"
)

var mappingAbstractOperationAttributesModelTypeEnum = map[string]AbstractOperationAttributesModelTypeEnum{
	"GENERIC_REST_API_ATTRIBUTES": AbstractOperationAttributesModelTypeGenericRestApiAttributes,
}

var mappingAbstractOperationAttributesModelTypeEnumLowerCase = map[string]AbstractOperationAttributesModelTypeEnum{
	"generic_rest_api_attributes": AbstractOperationAttributesModelTypeGenericRestApiAttributes,
}

// GetAbstractOperationAttributesModelTypeEnumValues Enumerates the set of values for AbstractOperationAttributesModelTypeEnum
func GetAbstractOperationAttributesModelTypeEnumValues() []AbstractOperationAttributesModelTypeEnum {
	values := make([]AbstractOperationAttributesModelTypeEnum, 0)
	for _, v := range mappingAbstractOperationAttributesModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAbstractOperationAttributesModelTypeEnumStringValues Enumerates the set of values in String for AbstractOperationAttributesModelTypeEnum
func GetAbstractOperationAttributesModelTypeEnumStringValues() []string {
	return []string{
		"GENERIC_REST_API_ATTRIBUTES",
	}
}

// GetMappingAbstractOperationAttributesModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAbstractOperationAttributesModelTypeEnum(val string) (AbstractOperationAttributesModelTypeEnum, bool) {
	enum, ok := mappingAbstractOperationAttributesModelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
