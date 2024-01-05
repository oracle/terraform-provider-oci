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

// DetailedDescription The detailed description of an object.
type DetailedDescription struct {

	// The type of the published object.
	ModelType DetailedDescriptionModelTypeEnum `mandatory:"false" json:"modelType,omitempty"`

	// Generated key that can be used in API calls to identify task. On scenarios where reference to the task is needed, a value can be passed in create.
	Key *string `mandatory:"false" json:"key"`

	// The object's model version.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"false" json:"name"`

	// The version of the object that is used to track changes in the object instance.
	ObjectVersion *int `mandatory:"false" json:"objectVersion"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
	Identifier *string `mandatory:"false" json:"identifier"`

	Metadata *ObjectMetadata `mandatory:"false" json:"metadata"`

	// Base64 encoded image to represent logo of the object.
	Logo *string `mandatory:"false" json:"logo"`

	// Base64 encoded rich text description of the object.
	DetailedDescription *string `mandatory:"false" json:"detailedDescription"`
}

func (m DetailedDescription) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DetailedDescription) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDetailedDescriptionModelTypeEnum(string(m.ModelType)); !ok && m.ModelType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ModelType: %s. Supported values are: %s.", m.ModelType, strings.Join(GetDetailedDescriptionModelTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DetailedDescriptionModelTypeEnum Enum with underlying type: string
type DetailedDescriptionModelTypeEnum string

// Set of constants representing the allowable values for DetailedDescriptionModelTypeEnum
const (
	DetailedDescriptionModelTypeDetailedDescription DetailedDescriptionModelTypeEnum = "DETAILED_DESCRIPTION"
)

var mappingDetailedDescriptionModelTypeEnum = map[string]DetailedDescriptionModelTypeEnum{
	"DETAILED_DESCRIPTION": DetailedDescriptionModelTypeDetailedDescription,
}

var mappingDetailedDescriptionModelTypeEnumLowerCase = map[string]DetailedDescriptionModelTypeEnum{
	"detailed_description": DetailedDescriptionModelTypeDetailedDescription,
}

// GetDetailedDescriptionModelTypeEnumValues Enumerates the set of values for DetailedDescriptionModelTypeEnum
func GetDetailedDescriptionModelTypeEnumValues() []DetailedDescriptionModelTypeEnum {
	values := make([]DetailedDescriptionModelTypeEnum, 0)
	for _, v := range mappingDetailedDescriptionModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDetailedDescriptionModelTypeEnumStringValues Enumerates the set of values in String for DetailedDescriptionModelTypeEnum
func GetDetailedDescriptionModelTypeEnumStringValues() []string {
	return []string{
		"DETAILED_DESCRIPTION",
	}
}

// GetMappingDetailedDescriptionModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDetailedDescriptionModelTypeEnum(val string) (DetailedDescriptionModelTypeEnum, bool) {
	enum, ok := mappingDetailedDescriptionModelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
