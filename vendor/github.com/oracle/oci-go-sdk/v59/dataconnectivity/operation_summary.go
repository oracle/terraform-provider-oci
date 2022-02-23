// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Connectivity Management API
//
// Use the DCMS APIs to perform Metadata/Data operations.
//

package dataconnectivity

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v59/common"
	"strings"
)

// OperationSummary The operation summary object
type OperationSummary interface {
	GetMetadata() *ObjectMetadata
}

type operationsummary struct {
	JsonData  []byte
	Metadata  *ObjectMetadata `mandatory:"false" json:"metadata"`
	ModelType string          `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *operationsummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshaleroperationsummary operationsummary
	s := struct {
		Model Unmarshaleroperationsummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Metadata = s.Model.Metadata
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *operationsummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "PROCEDURE":
		mm := OperationSummaryFromProcedure{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetMetadata returns Metadata
func (m operationsummary) GetMetadata() *ObjectMetadata {
	return m.Metadata
}

func (m operationsummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m operationsummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OperationSummaryModelTypeEnum Enum with underlying type: string
type OperationSummaryModelTypeEnum string

// Set of constants representing the allowable values for OperationSummaryModelTypeEnum
const (
	OperationSummaryModelTypeProcedure OperationSummaryModelTypeEnum = "PROCEDURE"
)

var mappingOperationSummaryModelTypeEnum = map[string]OperationSummaryModelTypeEnum{
	"PROCEDURE": OperationSummaryModelTypeProcedure,
}

var mappingOperationSummaryModelTypeEnumLowerCase = map[string]OperationSummaryModelTypeEnum{
	"procedure": OperationSummaryModelTypeProcedure,
}

// GetOperationSummaryModelTypeEnumValues Enumerates the set of values for OperationSummaryModelTypeEnum
func GetOperationSummaryModelTypeEnumValues() []OperationSummaryModelTypeEnum {
	values := make([]OperationSummaryModelTypeEnum, 0)
	for _, v := range mappingOperationSummaryModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOperationSummaryModelTypeEnumStringValues Enumerates the set of values in String for OperationSummaryModelTypeEnum
func GetOperationSummaryModelTypeEnumStringValues() []string {
	return []string{
		"PROCEDURE",
	}
}

// GetMappingOperationSummaryModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationSummaryModelTypeEnum(val string) (OperationSummaryModelTypeEnum, bool) {
	enum, ok := mappingOperationSummaryModelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
