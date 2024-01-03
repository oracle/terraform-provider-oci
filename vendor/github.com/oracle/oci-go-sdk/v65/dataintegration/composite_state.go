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

// CompositeState The composite state object provides information on the state of a task or schedule.
type CompositeState struct {

	// The type of the Composite State Aggregator.
	CompositeStateAggregator CompositeStateCompositeStateAggregatorEnum `mandatory:"false" json:"compositeStateAggregator,omitempty"`

	// Generated key that can be used in API calls to identify Composite State.
	Key *string `mandatory:"false" json:"key"`

	// The type of the object.
	ModelType *string `mandatory:"false" json:"modelType"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"false" json:"name"`

	// Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
	Identifier *string `mandatory:"false" json:"identifier"`

	// Detailed description for the object.
	Description *string `mandatory:"false" json:"description"`

	// The version of the object that is used to track changes in the object instance.
	ObjectVersion *int `mandatory:"false" json:"objectVersion"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Map that stores all the States for a given Task or Schedule
	AllStatesMap map[string]State `mandatory:"false" json:"allStatesMap"`

	RegistryMetadata *RegistryMetadata `mandatory:"false" json:"registryMetadata"`

	Metadata *ObjectMetadata `mandatory:"false" json:"metadata"`
}

func (m CompositeState) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CompositeState) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCompositeStateCompositeStateAggregatorEnum(string(m.CompositeStateAggregator)); !ok && m.CompositeStateAggregator != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CompositeStateAggregator: %s. Supported values are: %s.", m.CompositeStateAggregator, strings.Join(GetCompositeStateCompositeStateAggregatorEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CompositeStateCompositeStateAggregatorEnum Enum with underlying type: string
type CompositeStateCompositeStateAggregatorEnum string

// Set of constants representing the allowable values for CompositeStateCompositeStateAggregatorEnum
const (
	CompositeStateCompositeStateAggregatorTaskSchedule CompositeStateCompositeStateAggregatorEnum = "TASK_SCHEDULE"
	CompositeStateCompositeStateAggregatorTask         CompositeStateCompositeStateAggregatorEnum = "TASK"
	CompositeStateCompositeStateAggregatorTaskOperator CompositeStateCompositeStateAggregatorEnum = "TASK_OPERATOR"
)

var mappingCompositeStateCompositeStateAggregatorEnum = map[string]CompositeStateCompositeStateAggregatorEnum{
	"TASK_SCHEDULE": CompositeStateCompositeStateAggregatorTaskSchedule,
	"TASK":          CompositeStateCompositeStateAggregatorTask,
	"TASK_OPERATOR": CompositeStateCompositeStateAggregatorTaskOperator,
}

var mappingCompositeStateCompositeStateAggregatorEnumLowerCase = map[string]CompositeStateCompositeStateAggregatorEnum{
	"task_schedule": CompositeStateCompositeStateAggregatorTaskSchedule,
	"task":          CompositeStateCompositeStateAggregatorTask,
	"task_operator": CompositeStateCompositeStateAggregatorTaskOperator,
}

// GetCompositeStateCompositeStateAggregatorEnumValues Enumerates the set of values for CompositeStateCompositeStateAggregatorEnum
func GetCompositeStateCompositeStateAggregatorEnumValues() []CompositeStateCompositeStateAggregatorEnum {
	values := make([]CompositeStateCompositeStateAggregatorEnum, 0)
	for _, v := range mappingCompositeStateCompositeStateAggregatorEnum {
		values = append(values, v)
	}
	return values
}

// GetCompositeStateCompositeStateAggregatorEnumStringValues Enumerates the set of values in String for CompositeStateCompositeStateAggregatorEnum
func GetCompositeStateCompositeStateAggregatorEnumStringValues() []string {
	return []string{
		"TASK_SCHEDULE",
		"TASK",
		"TASK_OPERATOR",
	}
}

// GetMappingCompositeStateCompositeStateAggregatorEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCompositeStateCompositeStateAggregatorEnum(val string) (CompositeStateCompositeStateAggregatorEnum, bool) {
	enum, ok := mappingCompositeStateCompositeStateAggregatorEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
