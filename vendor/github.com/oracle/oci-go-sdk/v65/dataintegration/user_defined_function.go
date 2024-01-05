// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UserDefinedFunction The user defined function type contains the audit summary information and the definition of the user defined function.
type UserDefinedFunction struct {

	// Generated key that can be used in API calls to identify user defined function. On scenarios where reference to the user defined function is needed, a value can be passed in create.
	Key *string `mandatory:"false" json:"key"`

	// The type of the object.
	ModelType UserDefinedFunctionModelTypeEnum `mandatory:"false" json:"modelType,omitempty"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"false" json:"name"`

	// Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
	Identifier *string `mandatory:"false" json:"identifier"`

	// The version of the object that is used to track changes in the object instance.
	ObjectVersion *int `mandatory:"false" json:"objectVersion"`

	// An array of function signature.
	Signatures []FunctionSignature `mandatory:"false" json:"signatures"`

	Expr *Expression `mandatory:"false" json:"expr"`

	// Detailed description for the object.
	Description *string `mandatory:"false" json:"description"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	Metadata *ObjectMetadata `mandatory:"false" json:"metadata"`

	// A key map. If provided, key is replaced with generated key. This structure provides mapping between user provided key and generated key.
	KeyMap map[string]string `mandatory:"false" json:"keyMap"`
}

func (m UserDefinedFunction) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UserDefinedFunction) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUserDefinedFunctionModelTypeEnum(string(m.ModelType)); !ok && m.ModelType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ModelType: %s. Supported values are: %s.", m.ModelType, strings.Join(GetUserDefinedFunctionModelTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UserDefinedFunctionModelTypeEnum Enum with underlying type: string
type UserDefinedFunctionModelTypeEnum string

// Set of constants representing the allowable values for UserDefinedFunctionModelTypeEnum
const (
	UserDefinedFunctionModelTypeDisUserDefinedFunction UserDefinedFunctionModelTypeEnum = "DIS_USER_DEFINED_FUNCTION"
)

var mappingUserDefinedFunctionModelTypeEnum = map[string]UserDefinedFunctionModelTypeEnum{
	"DIS_USER_DEFINED_FUNCTION": UserDefinedFunctionModelTypeDisUserDefinedFunction,
}

var mappingUserDefinedFunctionModelTypeEnumLowerCase = map[string]UserDefinedFunctionModelTypeEnum{
	"dis_user_defined_function": UserDefinedFunctionModelTypeDisUserDefinedFunction,
}

// GetUserDefinedFunctionModelTypeEnumValues Enumerates the set of values for UserDefinedFunctionModelTypeEnum
func GetUserDefinedFunctionModelTypeEnumValues() []UserDefinedFunctionModelTypeEnum {
	values := make([]UserDefinedFunctionModelTypeEnum, 0)
	for _, v := range mappingUserDefinedFunctionModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUserDefinedFunctionModelTypeEnumStringValues Enumerates the set of values in String for UserDefinedFunctionModelTypeEnum
func GetUserDefinedFunctionModelTypeEnumStringValues() []string {
	return []string{
		"DIS_USER_DEFINED_FUNCTION",
	}
}

// GetMappingUserDefinedFunctionModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUserDefinedFunctionModelTypeEnum(val string) (UserDefinedFunctionModelTypeEnum, bool) {
	enum, ok := mappingUserDefinedFunctionModelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
