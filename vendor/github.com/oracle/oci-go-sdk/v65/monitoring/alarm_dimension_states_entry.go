// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Monitoring API
//
// Use the Monitoring API to manage metric queries and alarms for assessing the health, capacity, and performance of your cloud resources.
// Endpoints vary by operation. For PostMetric, use the `telemetry-ingestion` endpoints; for all other operations, use the `telemetry` endpoints.
// For more information, see
// the Monitoring documentation (https://docs.cloud.oracle.com/iaas/Content/Monitoring/home.htm).
//

package monitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AlarmDimensionStatesEntry A timestamped alarm state entry for a metric stream.
type AlarmDimensionStatesEntry struct {

	// Indicator of the metric stream associated with the alarm state entry. Includes one or more dimension key-value pairs.
	Dimensions map[string]string `mandatory:"true" json:"dimensions"`

	// Transition state (status value) associated with the alarm state entry.
	// Example: `FIRING`
	Status AlarmDimensionStatesEntryStatusEnum `mandatory:"true" json:"status"`

	// Transition time associated with the alarm state entry. Format defined by RFC3339.
	// Example: `2022-02-01T01:02:29.600Z`
	Timestamp *common.SDKTime `mandatory:"true" json:"timestamp"`
}

func (m AlarmDimensionStatesEntry) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AlarmDimensionStatesEntry) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAlarmDimensionStatesEntryStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetAlarmDimensionStatesEntryStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AlarmDimensionStatesEntryStatusEnum Enum with underlying type: string
type AlarmDimensionStatesEntryStatusEnum string

// Set of constants representing the allowable values for AlarmDimensionStatesEntryStatusEnum
const (
	AlarmDimensionStatesEntryStatusFiring AlarmDimensionStatesEntryStatusEnum = "FIRING"
	AlarmDimensionStatesEntryStatusOk     AlarmDimensionStatesEntryStatusEnum = "OK"
)

var mappingAlarmDimensionStatesEntryStatusEnum = map[string]AlarmDimensionStatesEntryStatusEnum{
	"FIRING": AlarmDimensionStatesEntryStatusFiring,
	"OK":     AlarmDimensionStatesEntryStatusOk,
}

var mappingAlarmDimensionStatesEntryStatusEnumLowerCase = map[string]AlarmDimensionStatesEntryStatusEnum{
	"firing": AlarmDimensionStatesEntryStatusFiring,
	"ok":     AlarmDimensionStatesEntryStatusOk,
}

// GetAlarmDimensionStatesEntryStatusEnumValues Enumerates the set of values for AlarmDimensionStatesEntryStatusEnum
func GetAlarmDimensionStatesEntryStatusEnumValues() []AlarmDimensionStatesEntryStatusEnum {
	values := make([]AlarmDimensionStatesEntryStatusEnum, 0)
	for _, v := range mappingAlarmDimensionStatesEntryStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetAlarmDimensionStatesEntryStatusEnumStringValues Enumerates the set of values in String for AlarmDimensionStatesEntryStatusEnum
func GetAlarmDimensionStatesEntryStatusEnumStringValues() []string {
	return []string{
		"FIRING",
		"OK",
	}
}

// GetMappingAlarmDimensionStatesEntryStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAlarmDimensionStatesEntryStatusEnum(val string) (AlarmDimensionStatesEntryStatusEnum, bool) {
	enum, ok := mappingAlarmDimensionStatesEntryStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
