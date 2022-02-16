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

// CreateTaskRunDetails The properties used in task run create operations.
type CreateTaskRunDetails struct {

	// The key of the object.
	Key *string `mandatory:"false" json:"key"`

	// The type of the object.
	ModelType *string `mandatory:"false" json:"modelType"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"false" json:"name"`

	// Detailed description for the object.
	Description *string `mandatory:"false" json:"description"`

	ConfigProvider *CreateConfigProvider `mandatory:"false" json:"configProvider"`

	// Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
	Identifier *string `mandatory:"false" json:"identifier"`

	// Optional task schedule key reference.
	TaskScheduleKey *string `mandatory:"false" json:"taskScheduleKey"`

	// Reference Task Run Id to be used for re-run
	RefTaskRunId *string `mandatory:"false" json:"refTaskRunId"`

	// Supported re-run types
	ReRunType CreateTaskRunDetailsReRunTypeEnum `mandatory:"false" json:"reRunType,omitempty"`

	// Step Id for running from a certain step.
	StepId *string `mandatory:"false" json:"stepId"`

	RegistryMetadata *RegistryMetadata `mandatory:"false" json:"registryMetadata"`
}

func (m CreateTaskRunDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateTaskRunDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreateTaskRunDetailsReRunTypeEnum(string(m.ReRunType)); !ok && m.ReRunType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ReRunType: %s. Supported values are: %s.", m.ReRunType, strings.Join(GetCreateTaskRunDetailsReRunTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateTaskRunDetailsReRunTypeEnum Enum with underlying type: string
type CreateTaskRunDetailsReRunTypeEnum string

// Set of constants representing the allowable values for CreateTaskRunDetailsReRunTypeEnum
const (
	CreateTaskRunDetailsReRunTypeBeginning CreateTaskRunDetailsReRunTypeEnum = "BEGINNING"
	CreateTaskRunDetailsReRunTypeFailed    CreateTaskRunDetailsReRunTypeEnum = "FAILED"
	CreateTaskRunDetailsReRunTypeStep      CreateTaskRunDetailsReRunTypeEnum = "STEP"
)

var mappingCreateTaskRunDetailsReRunTypeEnum = map[string]CreateTaskRunDetailsReRunTypeEnum{
	"BEGINNING": CreateTaskRunDetailsReRunTypeBeginning,
	"FAILED":    CreateTaskRunDetailsReRunTypeFailed,
	"STEP":      CreateTaskRunDetailsReRunTypeStep,
}

// GetCreateTaskRunDetailsReRunTypeEnumValues Enumerates the set of values for CreateTaskRunDetailsReRunTypeEnum
func GetCreateTaskRunDetailsReRunTypeEnumValues() []CreateTaskRunDetailsReRunTypeEnum {
	values := make([]CreateTaskRunDetailsReRunTypeEnum, 0)
	for _, v := range mappingCreateTaskRunDetailsReRunTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateTaskRunDetailsReRunTypeEnumStringValues Enumerates the set of values in String for CreateTaskRunDetailsReRunTypeEnum
func GetCreateTaskRunDetailsReRunTypeEnumStringValues() []string {
	return []string{
		"BEGINNING",
		"FAILED",
		"STEP",
	}
}

// GetMappingCreateTaskRunDetailsReRunTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateTaskRunDetailsReRunTypeEnum(val string) (CreateTaskRunDetailsReRunTypeEnum, bool) {
	mappingCreateTaskRunDetailsReRunTypeEnumIgnoreCase := make(map[string]CreateTaskRunDetailsReRunTypeEnum)
	for k, v := range mappingCreateTaskRunDetailsReRunTypeEnum {
		mappingCreateTaskRunDetailsReRunTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingCreateTaskRunDetailsReRunTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
