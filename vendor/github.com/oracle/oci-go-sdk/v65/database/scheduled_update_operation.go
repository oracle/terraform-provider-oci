// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ScheduledUpdateOperation A scheduled update operation for an Autonomous AI Database. Exactly one of `isScheduleUpdateToEarliest`, `timeScheduledUpdate`, or `isDisableUpdateSchedule` must be specified for each operation.
type ScheduledUpdateOperation interface {

	// Schedule to earliest available time slot.
	GetIsScheduleUpdateToEarliest() *bool

	// Scheduled time (RFC3339).
	GetTimeScheduledUpdate() *common.SDKTime

	// Cancel existing scheduled update.
	GetIsDisableUpdateSchedule() *bool
}

type scheduledupdateoperation struct {
	JsonData                   []byte
	IsScheduleUpdateToEarliest *bool           `mandatory:"false" json:"isScheduleUpdateToEarliest"`
	TimeScheduledUpdate        *common.SDKTime `mandatory:"false" json:"timeScheduledUpdate"`
	IsDisableUpdateSchedule    *bool           `mandatory:"false" json:"isDisableUpdateSchedule"`
	OperationType              string          `json:"operationType"`
}

// UnmarshalJSON unmarshals json
func (m *scheduledupdateoperation) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerscheduledupdateoperation scheduledupdateoperation
	s := struct {
		Model Unmarshalerscheduledupdateoperation
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.IsScheduleUpdateToEarliest = s.Model.IsScheduleUpdateToEarliest
	m.TimeScheduledUpdate = s.Model.TimeScheduledUpdate
	m.IsDisableUpdateSchedule = s.Model.IsDisableUpdateSchedule
	m.OperationType = s.Model.OperationType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *scheduledupdateoperation) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.OperationType {
	case "EXTERNAL_LOCATION_UPDATE":
		mm := ExternalLocationScheduleOperation{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for ScheduledUpdateOperation: %s.", m.OperationType)
		return *m, nil
	}
}

// GetIsScheduleUpdateToEarliest returns IsScheduleUpdateToEarliest
func (m scheduledupdateoperation) GetIsScheduleUpdateToEarliest() *bool {
	return m.IsScheduleUpdateToEarliest
}

// GetTimeScheduledUpdate returns TimeScheduledUpdate
func (m scheduledupdateoperation) GetTimeScheduledUpdate() *common.SDKTime {
	return m.TimeScheduledUpdate
}

// GetIsDisableUpdateSchedule returns IsDisableUpdateSchedule
func (m scheduledupdateoperation) GetIsDisableUpdateSchedule() *bool {
	return m.IsDisableUpdateSchedule
}

func (m scheduledupdateoperation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m scheduledupdateoperation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ScheduledUpdateOperationOperationTypeEnum Enum with underlying type: string
type ScheduledUpdateOperationOperationTypeEnum string

// Set of constants representing the allowable values for ScheduledUpdateOperationOperationTypeEnum
const (
	ScheduledUpdateOperationOperationTypeExternalLocationUpdate ScheduledUpdateOperationOperationTypeEnum = "EXTERNAL_LOCATION_UPDATE"
)

var mappingScheduledUpdateOperationOperationTypeEnum = map[string]ScheduledUpdateOperationOperationTypeEnum{
	"EXTERNAL_LOCATION_UPDATE": ScheduledUpdateOperationOperationTypeExternalLocationUpdate,
}

var mappingScheduledUpdateOperationOperationTypeEnumLowerCase = map[string]ScheduledUpdateOperationOperationTypeEnum{
	"external_location_update": ScheduledUpdateOperationOperationTypeExternalLocationUpdate,
}

// GetScheduledUpdateOperationOperationTypeEnumValues Enumerates the set of values for ScheduledUpdateOperationOperationTypeEnum
func GetScheduledUpdateOperationOperationTypeEnumValues() []ScheduledUpdateOperationOperationTypeEnum {
	values := make([]ScheduledUpdateOperationOperationTypeEnum, 0)
	for _, v := range mappingScheduledUpdateOperationOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetScheduledUpdateOperationOperationTypeEnumStringValues Enumerates the set of values in String for ScheduledUpdateOperationOperationTypeEnum
func GetScheduledUpdateOperationOperationTypeEnumStringValues() []string {
	return []string{
		"EXTERNAL_LOCATION_UPDATE",
	}
}

// GetMappingScheduledUpdateOperationOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingScheduledUpdateOperationOperationTypeEnum(val string) (ScheduledUpdateOperationOperationTypeEnum, bool) {
	enum, ok := mappingScheduledUpdateOperationOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
