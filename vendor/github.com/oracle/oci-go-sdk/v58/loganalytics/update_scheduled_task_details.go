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

// UpdateScheduledTaskDetails The details for updating a schedule task.
type UpdateScheduledTaskDetails interface {

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

	// Schedules may be updated for task types SAVED_SEARCH and PURGE.
	// Note there may only be a single schedule for SAVED_SEARCH and PURGE scheduled tasks.
	GetSchedules() []Schedule
}

type updatescheduledtaskdetails struct {
	JsonData     []byte
	DisplayName  *string                           `mandatory:"false" json:"displayName"`
	FreeformTags map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags  map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	Schedules    json.RawMessage                   `mandatory:"false" json:"schedules"`
	Kind         string                            `json:"kind"`
}

// UnmarshalJSON unmarshals json
func (m *updatescheduledtaskdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdatescheduledtaskdetails updatescheduledtaskdetails
	s := struct {
		Model Unmarshalerupdatescheduledtaskdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DisplayName = s.Model.DisplayName
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.Schedules = s.Model.Schedules
	m.Kind = s.Model.Kind

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updatescheduledtaskdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Kind {
	case "STANDARD":
		mm := UpdateStandardTaskDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetDisplayName returns DisplayName
func (m updatescheduledtaskdetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetFreeformTags returns FreeformTags
func (m updatescheduledtaskdetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m updatescheduledtaskdetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetSchedules returns Schedules
func (m updatescheduledtaskdetails) GetSchedules() json.RawMessage {
	return m.Schedules
}

func (m updatescheduledtaskdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updatescheduledtaskdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateScheduledTaskDetailsKindEnum Enum with underlying type: string
type UpdateScheduledTaskDetailsKindEnum string

// Set of constants representing the allowable values for UpdateScheduledTaskDetailsKindEnum
const (
	UpdateScheduledTaskDetailsKindAcceleration UpdateScheduledTaskDetailsKindEnum = "ACCELERATION"
	UpdateScheduledTaskDetailsKindStandard     UpdateScheduledTaskDetailsKindEnum = "STANDARD"
)

var mappingUpdateScheduledTaskDetailsKindEnum = map[string]UpdateScheduledTaskDetailsKindEnum{
	"ACCELERATION": UpdateScheduledTaskDetailsKindAcceleration,
	"STANDARD":     UpdateScheduledTaskDetailsKindStandard,
}

// GetUpdateScheduledTaskDetailsKindEnumValues Enumerates the set of values for UpdateScheduledTaskDetailsKindEnum
func GetUpdateScheduledTaskDetailsKindEnumValues() []UpdateScheduledTaskDetailsKindEnum {
	values := make([]UpdateScheduledTaskDetailsKindEnum, 0)
	for _, v := range mappingUpdateScheduledTaskDetailsKindEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateScheduledTaskDetailsKindEnumStringValues Enumerates the set of values in String for UpdateScheduledTaskDetailsKindEnum
func GetUpdateScheduledTaskDetailsKindEnumStringValues() []string {
	return []string{
		"ACCELERATION",
		"STANDARD",
	}
}

// GetMappingUpdateScheduledTaskDetailsKindEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateScheduledTaskDetailsKindEnum(val string) (UpdateScheduledTaskDetailsKindEnum, bool) {
	mappingUpdateScheduledTaskDetailsKindEnumIgnoreCase := make(map[string]UpdateScheduledTaskDetailsKindEnum)
	for k, v := range mappingUpdateScheduledTaskDetailsKindEnum {
		mappingUpdateScheduledTaskDetailsKindEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingUpdateScheduledTaskDetailsKindEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
