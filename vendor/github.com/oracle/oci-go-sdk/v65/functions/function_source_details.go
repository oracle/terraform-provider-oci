// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Functions Service API
//
// API for the Functions service.
//

package functions

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// FunctionSourceDetails The source details for the Function. The function can be created from various sources.
type FunctionSourceDetails interface {
}

type functionsourcedetails struct {
	JsonData   []byte
	SourceType string `json:"sourceType"`
}

// UnmarshalJSON unmarshals json
func (m *functionsourcedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerfunctionsourcedetails functionsourcedetails
	s := struct {
		Model Unmarshalerfunctionsourcedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.SourceType = s.Model.SourceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *functionsourcedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.SourceType {
	case "PRE_BUILT_FUNCTIONS":
		mm := PreBuiltFunctionSourceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for FunctionSourceDetails: %s.", m.SourceType)
		return *m, nil
	}
}

func (m functionsourcedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m functionsourcedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// FunctionSourceDetailsSourceTypeEnum Enum with underlying type: string
type FunctionSourceDetailsSourceTypeEnum string

// Set of constants representing the allowable values for FunctionSourceDetailsSourceTypeEnum
const (
	FunctionSourceDetailsSourceTypePreBuiltFunctions FunctionSourceDetailsSourceTypeEnum = "PRE_BUILT_FUNCTIONS"
)

var mappingFunctionSourceDetailsSourceTypeEnum = map[string]FunctionSourceDetailsSourceTypeEnum{
	"PRE_BUILT_FUNCTIONS": FunctionSourceDetailsSourceTypePreBuiltFunctions,
}

var mappingFunctionSourceDetailsSourceTypeEnumLowerCase = map[string]FunctionSourceDetailsSourceTypeEnum{
	"pre_built_functions": FunctionSourceDetailsSourceTypePreBuiltFunctions,
}

// GetFunctionSourceDetailsSourceTypeEnumValues Enumerates the set of values for FunctionSourceDetailsSourceTypeEnum
func GetFunctionSourceDetailsSourceTypeEnumValues() []FunctionSourceDetailsSourceTypeEnum {
	values := make([]FunctionSourceDetailsSourceTypeEnum, 0)
	for _, v := range mappingFunctionSourceDetailsSourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetFunctionSourceDetailsSourceTypeEnumStringValues Enumerates the set of values in String for FunctionSourceDetailsSourceTypeEnum
func GetFunctionSourceDetailsSourceTypeEnumStringValues() []string {
	return []string{
		"PRE_BUILT_FUNCTIONS",
	}
}

// GetMappingFunctionSourceDetailsSourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFunctionSourceDetailsSourceTypeEnum(val string) (FunctionSourceDetailsSourceTypeEnum, bool) {
	enum, ok := mappingFunctionSourceDetailsSourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
