// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// FunctionLibrarySummary The FunctionLibrary summary type contains the audit summary information and the definition of the Function Library.
type FunctionLibrarySummary struct {

	// Generated key that can be used in API calls to identify FunctionLibrary.
	Key *string `mandatory:"false" json:"key"`

	// The type of the object.
	ModelType FunctionLibrarySummaryModelTypeEnum `mandatory:"false" json:"modelType,omitempty"`

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

func (m FunctionLibrarySummary) String() string {
	return common.PointerString(m)
}

// FunctionLibrarySummaryModelTypeEnum Enum with underlying type: string
type FunctionLibrarySummaryModelTypeEnum string

// Set of constants representing the allowable values for FunctionLibrarySummaryModelTypeEnum
const (
	FunctionLibrarySummaryModelTypeFunctionLibrary FunctionLibrarySummaryModelTypeEnum = "FUNCTION_LIBRARY"
)

var mappingFunctionLibrarySummaryModelType = map[string]FunctionLibrarySummaryModelTypeEnum{
	"FUNCTION_LIBRARY": FunctionLibrarySummaryModelTypeFunctionLibrary,
}

// GetFunctionLibrarySummaryModelTypeEnumValues Enumerates the set of values for FunctionLibrarySummaryModelTypeEnum
func GetFunctionLibrarySummaryModelTypeEnumValues() []FunctionLibrarySummaryModelTypeEnum {
	values := make([]FunctionLibrarySummaryModelTypeEnum, 0)
	for _, v := range mappingFunctionLibrarySummaryModelType {
		values = append(values, v)
	}
	return values
}
