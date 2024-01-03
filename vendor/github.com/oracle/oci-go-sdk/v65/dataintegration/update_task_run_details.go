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

// UpdateTaskRunDetails Properties used in task run update operations.
type UpdateTaskRunDetails struct {

	// The key of the object.
	Key *string `mandatory:"false" json:"key"`

	// The status of the object.
	Status UpdateTaskRunDetailsStatusEnum `mandatory:"false" json:"status,omitempty"`

	// The type of the object.
	ModelType *string `mandatory:"false" json:"modelType"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"false" json:"name"`

	// Detailed description for the object.
	Description *string `mandatory:"false" json:"description"`

	// The version of the object that is used to track changes in the object instance.
	ObjectVersion *int `mandatory:"false" json:"objectVersion"`

	// Optional task schedule key reference.
	TaskScheduleKey *string `mandatory:"false" json:"taskScheduleKey"`

	RegistryMetadata *RegistryMetadata `mandatory:"false" json:"registryMetadata"`
}

func (m UpdateTaskRunDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateTaskRunDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateTaskRunDetailsStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetUpdateTaskRunDetailsStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateTaskRunDetailsStatusEnum Enum with underlying type: string
type UpdateTaskRunDetailsStatusEnum string

// Set of constants representing the allowable values for UpdateTaskRunDetailsStatusEnum
const (
	UpdateTaskRunDetailsStatusTerminating UpdateTaskRunDetailsStatusEnum = "TERMINATING"
)

var mappingUpdateTaskRunDetailsStatusEnum = map[string]UpdateTaskRunDetailsStatusEnum{
	"TERMINATING": UpdateTaskRunDetailsStatusTerminating,
}

var mappingUpdateTaskRunDetailsStatusEnumLowerCase = map[string]UpdateTaskRunDetailsStatusEnum{
	"terminating": UpdateTaskRunDetailsStatusTerminating,
}

// GetUpdateTaskRunDetailsStatusEnumValues Enumerates the set of values for UpdateTaskRunDetailsStatusEnum
func GetUpdateTaskRunDetailsStatusEnumValues() []UpdateTaskRunDetailsStatusEnum {
	values := make([]UpdateTaskRunDetailsStatusEnum, 0)
	for _, v := range mappingUpdateTaskRunDetailsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateTaskRunDetailsStatusEnumStringValues Enumerates the set of values in String for UpdateTaskRunDetailsStatusEnum
func GetUpdateTaskRunDetailsStatusEnumStringValues() []string {
	return []string{
		"TERMINATING",
	}
}

// GetMappingUpdateTaskRunDetailsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateTaskRunDetailsStatusEnum(val string) (UpdateTaskRunDetailsStatusEnum, bool) {
	enum, ok := mappingUpdateTaskRunDetailsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
