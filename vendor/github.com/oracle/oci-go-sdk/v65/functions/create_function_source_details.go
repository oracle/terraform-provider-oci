// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// CreateFunctionSourceDetails The source details for creating the Function. The function can be created from various sources.
type CreateFunctionSourceDetails interface {
}

type createfunctionsourcedetails struct {
	JsonData   []byte
	SourceType string `json:"sourceType"`
}

// UnmarshalJSON unmarshals json
func (m *createfunctionsourcedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatefunctionsourcedetails createfunctionsourcedetails
	s := struct {
		Model Unmarshalercreatefunctionsourcedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.SourceType = s.Model.SourceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createfunctionsourcedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.SourceType {
	case "PRE_BUILT_FUNCTIONS":
		mm := CreatePreBuiltFunctionSourceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ARCHIVE":
		mm := CreateArchiveFunctionSourceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CONTAINER_IMAGE":
		mm := CreateContainerImageFunctionSourceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CreateFunctionSourceDetails: %s.", m.SourceType)
		return *m, nil
	}
}

func (m createfunctionsourcedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createfunctionsourcedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateFunctionSourceDetailsSourceTypeEnum Enum with underlying type: string
type CreateFunctionSourceDetailsSourceTypeEnum string

// Set of constants representing the allowable values for CreateFunctionSourceDetailsSourceTypeEnum
const (
	CreateFunctionSourceDetailsSourceTypePreBuiltFunctions CreateFunctionSourceDetailsSourceTypeEnum = "PRE_BUILT_FUNCTIONS"
	CreateFunctionSourceDetailsSourceTypeArchive           CreateFunctionSourceDetailsSourceTypeEnum = "ARCHIVE"
	CreateFunctionSourceDetailsSourceTypeContainerImage    CreateFunctionSourceDetailsSourceTypeEnum = "CONTAINER_IMAGE"
)

var mappingCreateFunctionSourceDetailsSourceTypeEnum = map[string]CreateFunctionSourceDetailsSourceTypeEnum{
	"PRE_BUILT_FUNCTIONS": CreateFunctionSourceDetailsSourceTypePreBuiltFunctions,
	"ARCHIVE":             CreateFunctionSourceDetailsSourceTypeArchive,
	"CONTAINER_IMAGE":     CreateFunctionSourceDetailsSourceTypeContainerImage,
}

var mappingCreateFunctionSourceDetailsSourceTypeEnumLowerCase = map[string]CreateFunctionSourceDetailsSourceTypeEnum{
	"pre_built_functions": CreateFunctionSourceDetailsSourceTypePreBuiltFunctions,
	"archive":             CreateFunctionSourceDetailsSourceTypeArchive,
	"container_image":     CreateFunctionSourceDetailsSourceTypeContainerImage,
}

// GetCreateFunctionSourceDetailsSourceTypeEnumValues Enumerates the set of values for CreateFunctionSourceDetailsSourceTypeEnum
func GetCreateFunctionSourceDetailsSourceTypeEnumValues() []CreateFunctionSourceDetailsSourceTypeEnum {
	values := make([]CreateFunctionSourceDetailsSourceTypeEnum, 0)
	for _, v := range mappingCreateFunctionSourceDetailsSourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateFunctionSourceDetailsSourceTypeEnumStringValues Enumerates the set of values in String for CreateFunctionSourceDetailsSourceTypeEnum
func GetCreateFunctionSourceDetailsSourceTypeEnumStringValues() []string {
	return []string{
		"PRE_BUILT_FUNCTIONS",
		"ARCHIVE",
		"CONTAINER_IMAGE",
	}
}

// GetMappingCreateFunctionSourceDetailsSourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateFunctionSourceDetailsSourceTypeEnum(val string) (CreateFunctionSourceDetailsSourceTypeEnum, bool) {
	enum, ok := mappingCreateFunctionSourceDetailsSourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
