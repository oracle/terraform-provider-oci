// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// CreateScheduledTaskDetails Details for creating a scheduled task.
type CreateScheduledTaskDetails interface {

	// Compartment Identifier OCID  (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	GetCompartmentId() *string

	// A user-friendly name that is changeable and that does not have to be unique.
	// Format: a leading alphanumeric, followed by zero or more
	// alphanumerics, underscores, spaces, backslashes, or hyphens in any order).
	// No trailing spaces allowed.
	GetDisplayName() *string

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}
}

type createscheduledtaskdetails struct {
	JsonData      []byte
	CompartmentId *string                           `mandatory:"true" json:"compartmentId"`
	DisplayName   *string                           `mandatory:"false" json:"displayName"`
	FreeformTags  map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags   map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	Kind          string                            `json:"kind"`
}

// UnmarshalJSON unmarshals json
func (m *createscheduledtaskdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatescheduledtaskdetails createscheduledtaskdetails
	s := struct {
		Model Unmarshalercreatescheduledtaskdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.CompartmentId = s.Model.CompartmentId
	m.DisplayName = s.Model.DisplayName
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.Kind = s.Model.Kind

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createscheduledtaskdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Kind {
	case "STANDARD":
		mm := CreateStandardTaskDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ACCELERATION":
		mm := CreateAccelerationTaskDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetCompartmentId returns CompartmentId
func (m createscheduledtaskdetails) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetDisplayName returns DisplayName
func (m createscheduledtaskdetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetFreeformTags returns FreeformTags
func (m createscheduledtaskdetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m createscheduledtaskdetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m createscheduledtaskdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createscheduledtaskdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateScheduledTaskDetailsKindEnum Enum with underlying type: string
type CreateScheduledTaskDetailsKindEnum string

// Set of constants representing the allowable values for CreateScheduledTaskDetailsKindEnum
const (
	CreateScheduledTaskDetailsKindAcceleration CreateScheduledTaskDetailsKindEnum = "ACCELERATION"
	CreateScheduledTaskDetailsKindStandard     CreateScheduledTaskDetailsKindEnum = "STANDARD"
)

var mappingCreateScheduledTaskDetailsKindEnum = map[string]CreateScheduledTaskDetailsKindEnum{
	"ACCELERATION": CreateScheduledTaskDetailsKindAcceleration,
	"STANDARD":     CreateScheduledTaskDetailsKindStandard,
}

// GetCreateScheduledTaskDetailsKindEnumValues Enumerates the set of values for CreateScheduledTaskDetailsKindEnum
func GetCreateScheduledTaskDetailsKindEnumValues() []CreateScheduledTaskDetailsKindEnum {
	values := make([]CreateScheduledTaskDetailsKindEnum, 0)
	for _, v := range mappingCreateScheduledTaskDetailsKindEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateScheduledTaskDetailsKindEnumStringValues Enumerates the set of values in String for CreateScheduledTaskDetailsKindEnum
func GetCreateScheduledTaskDetailsKindEnumStringValues() []string {
	return []string{
		"ACCELERATION",
		"STANDARD",
	}
}

// GetMappingCreateScheduledTaskDetailsKindEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateScheduledTaskDetailsKindEnum(val string) (CreateScheduledTaskDetailsKindEnum, bool) {
	mappingCreateScheduledTaskDetailsKindEnumIgnoreCase := make(map[string]CreateScheduledTaskDetailsKindEnum)
	for k, v := range mappingCreateScheduledTaskDetailsKindEnum {
		mappingCreateScheduledTaskDetailsKindEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingCreateScheduledTaskDetailsKindEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
