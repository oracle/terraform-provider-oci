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

// FunctionLibrary The FunctionLibrary type contains the audit summary information and the definition of the FunctionLibrary.
type FunctionLibrary struct {

	// Generated key that can be used in API calls to identify FunctionLibrary.
	Key *string `mandatory:"false" json:"key"`

	// The type of the object.
	ModelType FunctionLibraryModelTypeEnum `mandatory:"false" json:"modelType,omitempty"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"false" json:"name"`

	// A user defined description for the Function Library.
	Description *string `mandatory:"false" json:"description"`

	// The category name.
	CategoryName *string `mandatory:"false" json:"categoryName"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
	Identifier *string `mandatory:"false" json:"identifier"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// The version of the object that is used to track changes in the object instance.
	ObjectVersion *int `mandatory:"false" json:"objectVersion"`

	Metadata *ObjectMetadata `mandatory:"false" json:"metadata"`

	// A key map. If provided, the key is replaced with generated key. This structure provides mapping between user provided key and generated key.
	KeyMap map[string]string `mandatory:"false" json:"keyMap"`
}

func (m FunctionLibrary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FunctionLibrary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingFunctionLibraryModelTypeEnum(string(m.ModelType)); !ok && m.ModelType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ModelType: %s. Supported values are: %s.", m.ModelType, strings.Join(GetFunctionLibraryModelTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// FunctionLibraryModelTypeEnum Enum with underlying type: string
type FunctionLibraryModelTypeEnum string

// Set of constants representing the allowable values for FunctionLibraryModelTypeEnum
const (
	FunctionLibraryModelTypeFunctionLibrary FunctionLibraryModelTypeEnum = "FUNCTION_LIBRARY"
)

var mappingFunctionLibraryModelTypeEnum = map[string]FunctionLibraryModelTypeEnum{
	"FUNCTION_LIBRARY": FunctionLibraryModelTypeFunctionLibrary,
}

// GetFunctionLibraryModelTypeEnumValues Enumerates the set of values for FunctionLibraryModelTypeEnum
func GetFunctionLibraryModelTypeEnumValues() []FunctionLibraryModelTypeEnum {
	values := make([]FunctionLibraryModelTypeEnum, 0)
	for _, v := range mappingFunctionLibraryModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetFunctionLibraryModelTypeEnumStringValues Enumerates the set of values in String for FunctionLibraryModelTypeEnum
func GetFunctionLibraryModelTypeEnumStringValues() []string {
	return []string{
		"FUNCTION_LIBRARY",
	}
}

// GetMappingFunctionLibraryModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFunctionLibraryModelTypeEnum(val string) (FunctionLibraryModelTypeEnum, bool) {
	mappingFunctionLibraryModelTypeEnumIgnoreCase := make(map[string]FunctionLibraryModelTypeEnum)
	for k, v := range mappingFunctionLibraryModelTypeEnum {
		mappingFunctionLibraryModelTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingFunctionLibraryModelTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
