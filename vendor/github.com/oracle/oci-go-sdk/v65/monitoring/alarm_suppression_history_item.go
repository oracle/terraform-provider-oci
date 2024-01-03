// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Monitoring API
//
// Use the Monitoring API to manage metric queries and alarms for assessing the health, capacity, and performance of your cloud resources.
// Endpoints vary by operation. For PostMetricData, use the `telemetry-ingestion` endpoints; for all other operations, use the `telemetry` endpoints.
// For more information, see
// the Monitoring documentation (https://docs.cloud.oracle.com/iaas/Content/Monitoring/home.htm).
//

package monitoring

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AlarmSuppressionHistoryItem A summary of properties for the specified alarm suppression history item.
type AlarmSuppressionHistoryItem struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the alarm suppression.
	SuppressionId *string `mandatory:"true" json:"suppressionId"`

	AlarmSuppressionTarget AlarmSuppressionTarget `mandatory:"true" json:"alarmSuppressionTarget"`

	// The level of this alarm suppression.
	// `ALARM` indicates a suppression of the entire alarm, regardless of dimension.
	// `DIMENSION` indicates a suppression configured for specified dimensions.
	Level AlarmSuppressionHistoryItemLevelEnum `mandatory:"true" json:"level"`

	// A user-friendly name for the alarm suppression. It does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The start date and time for the suppression actually starts, inclusive. Format defined by RFC3339.
	// Example: `2023-02-01T01:02:29.600Z`
	TimeEffectiveFrom *common.SDKTime `mandatory:"true" json:"timeEffectiveFrom"`

	// The end date and time for the suppression actually ends, inclusive. Format defined by RFC3339.
	// Example: `2023-02-01T02:02:29.600Z`
	TimeEffectiveUntil *common.SDKTime `mandatory:"true" json:"timeEffectiveUntil"`

	// Human-readable reason for this alarm suppression.
	// It does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	// Oracle recommends including tracking information for the event or associated work,
	// such as a ticket number.
	// Example: `Planned outage due to change IT-1234.`
	Description *string `mandatory:"false" json:"description"`

	// Configured dimension filter for suppressing alarm state entries that include the set of specified dimension key-value pairs.
	// Example: `{"resourceId": "ocid1.instance.region1.phx.exampleuniqueID"}`
	Dimensions map[string]string `mandatory:"false" json:"dimensions"`
}

func (m AlarmSuppressionHistoryItem) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AlarmSuppressionHistoryItem) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAlarmSuppressionHistoryItemLevelEnum(string(m.Level)); !ok && m.Level != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Level: %s. Supported values are: %s.", m.Level, strings.Join(GetAlarmSuppressionHistoryItemLevelEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *AlarmSuppressionHistoryItem) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description            *string                              `json:"description"`
		Dimensions             map[string]string                    `json:"dimensions"`
		SuppressionId          *string                              `json:"suppressionId"`
		AlarmSuppressionTarget alarmsuppressiontarget               `json:"alarmSuppressionTarget"`
		Level                  AlarmSuppressionHistoryItemLevelEnum `json:"level"`
		DisplayName            *string                              `json:"displayName"`
		TimeEffectiveFrom      *common.SDKTime                      `json:"timeEffectiveFrom"`
		TimeEffectiveUntil     *common.SDKTime                      `json:"timeEffectiveUntil"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.Dimensions = model.Dimensions

	m.SuppressionId = model.SuppressionId

	nn, e = model.AlarmSuppressionTarget.UnmarshalPolymorphicJSON(model.AlarmSuppressionTarget.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.AlarmSuppressionTarget = nn.(AlarmSuppressionTarget)
	} else {
		m.AlarmSuppressionTarget = nil
	}

	m.Level = model.Level

	m.DisplayName = model.DisplayName

	m.TimeEffectiveFrom = model.TimeEffectiveFrom

	m.TimeEffectiveUntil = model.TimeEffectiveUntil

	return
}

// AlarmSuppressionHistoryItemLevelEnum Enum with underlying type: string
type AlarmSuppressionHistoryItemLevelEnum string

// Set of constants representing the allowable values for AlarmSuppressionHistoryItemLevelEnum
const (
	AlarmSuppressionHistoryItemLevelAlarm     AlarmSuppressionHistoryItemLevelEnum = "ALARM"
	AlarmSuppressionHistoryItemLevelDimension AlarmSuppressionHistoryItemLevelEnum = "DIMENSION"
)

var mappingAlarmSuppressionHistoryItemLevelEnum = map[string]AlarmSuppressionHistoryItemLevelEnum{
	"ALARM":     AlarmSuppressionHistoryItemLevelAlarm,
	"DIMENSION": AlarmSuppressionHistoryItemLevelDimension,
}

var mappingAlarmSuppressionHistoryItemLevelEnumLowerCase = map[string]AlarmSuppressionHistoryItemLevelEnum{
	"alarm":     AlarmSuppressionHistoryItemLevelAlarm,
	"dimension": AlarmSuppressionHistoryItemLevelDimension,
}

// GetAlarmSuppressionHistoryItemLevelEnumValues Enumerates the set of values for AlarmSuppressionHistoryItemLevelEnum
func GetAlarmSuppressionHistoryItemLevelEnumValues() []AlarmSuppressionHistoryItemLevelEnum {
	values := make([]AlarmSuppressionHistoryItemLevelEnum, 0)
	for _, v := range mappingAlarmSuppressionHistoryItemLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetAlarmSuppressionHistoryItemLevelEnumStringValues Enumerates the set of values in String for AlarmSuppressionHistoryItemLevelEnum
func GetAlarmSuppressionHistoryItemLevelEnumStringValues() []string {
	return []string{
		"ALARM",
		"DIMENSION",
	}
}

// GetMappingAlarmSuppressionHistoryItemLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAlarmSuppressionHistoryItemLevelEnum(val string) (AlarmSuppressionHistoryItemLevelEnum, bool) {
	enum, ok := mappingAlarmSuppressionHistoryItemLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
