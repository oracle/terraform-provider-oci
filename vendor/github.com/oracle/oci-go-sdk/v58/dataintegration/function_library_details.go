// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// FunctionLibraryDetails The details including name, description for the function library, which is a container for user defined functions.
type FunctionLibraryDetails struct {

	// Generated key that can be used in API calls to identify FunctionLibrary.
	Key *string `mandatory:"true" json:"key"`

	// The type of the object.
	ModelType FunctionLibraryDetailsModelTypeEnum `mandatory:"true" json:"modelType"`

	// The version of the object that is used to track changes in the object instance.
	ObjectVersion *int `mandatory:"true" json:"objectVersion"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"false" json:"name"`

	// A user defined description for the FunctionLibrary.
	Description *string `mandatory:"false" json:"description"`

	// The category name.
	CategoryName *string `mandatory:"false" json:"categoryName"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
	Identifier *string `mandatory:"false" json:"identifier"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	RegistryMetadata *RegistryMetadata `mandatory:"false" json:"registryMetadata"`
}

func (m FunctionLibraryDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FunctionLibraryDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingFunctionLibraryDetailsModelTypeEnum(string(m.ModelType)); !ok && m.ModelType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ModelType: %s. Supported values are: %s.", m.ModelType, strings.Join(GetFunctionLibraryDetailsModelTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// FunctionLibraryDetailsModelTypeEnum Enum with underlying type: string
type FunctionLibraryDetailsModelTypeEnum string

// Set of constants representing the allowable values for FunctionLibraryDetailsModelTypeEnum
const (
	FunctionLibraryDetailsModelTypeFunctionLibrary FunctionLibraryDetailsModelTypeEnum = "FUNCTION_LIBRARY"
)

var mappingFunctionLibraryDetailsModelTypeEnum = map[string]FunctionLibraryDetailsModelTypeEnum{
	"FUNCTION_LIBRARY": FunctionLibraryDetailsModelTypeFunctionLibrary,
}

// GetFunctionLibraryDetailsModelTypeEnumValues Enumerates the set of values for FunctionLibraryDetailsModelTypeEnum
func GetFunctionLibraryDetailsModelTypeEnumValues() []FunctionLibraryDetailsModelTypeEnum {
	values := make([]FunctionLibraryDetailsModelTypeEnum, 0)
	for _, v := range mappingFunctionLibraryDetailsModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetFunctionLibraryDetailsModelTypeEnumStringValues Enumerates the set of values in String for FunctionLibraryDetailsModelTypeEnum
func GetFunctionLibraryDetailsModelTypeEnumStringValues() []string {
	return []string{
		"FUNCTION_LIBRARY",
	}
}

// GetMappingFunctionLibraryDetailsModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFunctionLibraryDetailsModelTypeEnum(val string) (FunctionLibraryDetailsModelTypeEnum, bool) {
	mappingFunctionLibraryDetailsModelTypeEnumIgnoreCase := make(map[string]FunctionLibraryDetailsModelTypeEnum)
	for k, v := range mappingFunctionLibraryDetailsModelTypeEnum {
		mappingFunctionLibraryDetailsModelTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingFunctionLibraryDetailsModelTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
