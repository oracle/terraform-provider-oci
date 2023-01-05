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

// Operation The operation object.
type Operation interface {
	GetOperationAttributes() AbstractOperationAttributes

	GetMetadata() *ObjectMetadata
}

type operation struct {
	JsonData            []byte
	OperationAttributes AbstractOperationAttributes `mandatory:"false" json:"operationAttributes"`
	Metadata            *ObjectMetadata             `mandatory:"false" json:"metadata"`
	ModelType           string                      `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *operation) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshaleroperation operation
	s := struct {
		Model Unmarshaleroperation
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.OperationAttributes = s.Model.OperationAttributes
	m.Metadata = s.Model.Metadata
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *operation) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "PROCEDURE":
		mm := OperationFromProcedure{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "API":
		mm := OperationFromApi{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for Operation: %s.", m.ModelType)
		return *m, nil
	}
}

//GetOperationAttributes returns OperationAttributes
func (m operation) GetOperationAttributes() AbstractOperationAttributes {
	return m.OperationAttributes
}

//GetMetadata returns Metadata
func (m operation) GetMetadata() *ObjectMetadata {
	return m.Metadata
}

func (m operation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m operation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OperationModelTypeEnum Enum with underlying type: string
type OperationModelTypeEnum string

// Set of constants representing the allowable values for OperationModelTypeEnum
const (
	OperationModelTypeProcedure OperationModelTypeEnum = "PROCEDURE"
	OperationModelTypeApi       OperationModelTypeEnum = "API"
)

var mappingOperationModelTypeEnum = map[string]OperationModelTypeEnum{
	"PROCEDURE": OperationModelTypeProcedure,
	"API":       OperationModelTypeApi,
}

var mappingOperationModelTypeEnumLowerCase = map[string]OperationModelTypeEnum{
	"procedure": OperationModelTypeProcedure,
	"api":       OperationModelTypeApi,
}

// GetOperationModelTypeEnumValues Enumerates the set of values for OperationModelTypeEnum
func GetOperationModelTypeEnumValues() []OperationModelTypeEnum {
	values := make([]OperationModelTypeEnum, 0)
	for _, v := range mappingOperationModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOperationModelTypeEnumStringValues Enumerates the set of values in String for OperationModelTypeEnum
func GetOperationModelTypeEnumStringValues() []string {
	return []string{
		"PROCEDURE",
		"API",
	}
}

// GetMappingOperationModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationModelTypeEnum(val string) (OperationModelTypeEnum, bool) {
	enum, ok := mappingOperationModelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
