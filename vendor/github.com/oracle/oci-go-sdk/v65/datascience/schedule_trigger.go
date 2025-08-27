// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ScheduleTrigger The trigger of the schedule can be UNIX cron or iCal expression or simple interval
type ScheduleTrigger interface {

	// The schedule starting date time, if null, System set the time when schedule is created.
	// Format is defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	GetTimeStart() *common.SDKTime

	// The schedule end date time, if null, the schedule will never expire.
	// Format is defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	GetTimeEnd() *common.SDKTime
}

type scheduletrigger struct {
	JsonData    []byte
	TimeStart   *common.SDKTime `mandatory:"false" json:"timeStart"`
	TimeEnd     *common.SDKTime `mandatory:"false" json:"timeEnd"`
	TriggerType string          `json:"triggerType"`
}

// UnmarshalJSON unmarshals json
func (m *scheduletrigger) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerscheduletrigger scheduletrigger
	s := struct {
		Model Unmarshalerscheduletrigger
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.TimeStart = s.Model.TimeStart
	m.TimeEnd = s.Model.TimeEnd
	m.TriggerType = s.Model.TriggerType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *scheduletrigger) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.TriggerType {
	case "ICAL":
		mm := ScheduleICalTrigger{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "INTERVAL":
		mm := ScheduleIntervalTrigger{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CRON":
		mm := ScheduleCronTrigger{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for ScheduleTrigger: %s.", m.TriggerType)
		return *m, nil
	}
}

// GetTimeStart returns TimeStart
func (m scheduletrigger) GetTimeStart() *common.SDKTime {
	return m.TimeStart
}

// GetTimeEnd returns TimeEnd
func (m scheduletrigger) GetTimeEnd() *common.SDKTime {
	return m.TimeEnd
}

func (m scheduletrigger) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m scheduletrigger) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ScheduleTriggerTriggerTypeEnum Enum with underlying type: string
type ScheduleTriggerTriggerTypeEnum string

// Set of constants representing the allowable values for ScheduleTriggerTriggerTypeEnum
const (
	ScheduleTriggerTriggerTypeCron     ScheduleTriggerTriggerTypeEnum = "CRON"
	ScheduleTriggerTriggerTypeInterval ScheduleTriggerTriggerTypeEnum = "INTERVAL"
	ScheduleTriggerTriggerTypeIcal     ScheduleTriggerTriggerTypeEnum = "ICAL"
)

var mappingScheduleTriggerTriggerTypeEnum = map[string]ScheduleTriggerTriggerTypeEnum{
	"CRON":     ScheduleTriggerTriggerTypeCron,
	"INTERVAL": ScheduleTriggerTriggerTypeInterval,
	"ICAL":     ScheduleTriggerTriggerTypeIcal,
}

var mappingScheduleTriggerTriggerTypeEnumLowerCase = map[string]ScheduleTriggerTriggerTypeEnum{
	"cron":     ScheduleTriggerTriggerTypeCron,
	"interval": ScheduleTriggerTriggerTypeInterval,
	"ical":     ScheduleTriggerTriggerTypeIcal,
}

// GetScheduleTriggerTriggerTypeEnumValues Enumerates the set of values for ScheduleTriggerTriggerTypeEnum
func GetScheduleTriggerTriggerTypeEnumValues() []ScheduleTriggerTriggerTypeEnum {
	values := make([]ScheduleTriggerTriggerTypeEnum, 0)
	for _, v := range mappingScheduleTriggerTriggerTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetScheduleTriggerTriggerTypeEnumStringValues Enumerates the set of values in String for ScheduleTriggerTriggerTypeEnum
func GetScheduleTriggerTriggerTypeEnumStringValues() []string {
	return []string{
		"CRON",
		"INTERVAL",
		"ICAL",
	}
}

// GetMappingScheduleTriggerTriggerTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingScheduleTriggerTriggerTypeEnum(val string) (ScheduleTriggerTriggerTypeEnum, bool) {
	enum, ok := mappingScheduleTriggerTriggerTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
