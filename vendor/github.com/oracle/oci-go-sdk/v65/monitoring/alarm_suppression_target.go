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

// AlarmSuppressionTarget The target of the alarm suppression.
type AlarmSuppressionTarget interface {
}

type alarmsuppressiontarget struct {
	JsonData   []byte
	TargetType string `json:"targetType"`
}

// UnmarshalJSON unmarshals json
func (m *alarmsuppressiontarget) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshaleralarmsuppressiontarget alarmsuppressiontarget
	s := struct {
		Model Unmarshaleralarmsuppressiontarget
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.TargetType = s.Model.TargetType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *alarmsuppressiontarget) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.TargetType {
	case "ALARM":
		mm := AlarmSuppressionAlarmTarget{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "COMPARTMENT":
		mm := AlarmSuppressionCompartmentTarget{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for AlarmSuppressionTarget: %s.", m.TargetType)
		return *m, nil
	}
}

func (m alarmsuppressiontarget) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m alarmsuppressiontarget) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AlarmSuppressionTargetTargetTypeEnum Enum with underlying type: string
type AlarmSuppressionTargetTargetTypeEnum string

// Set of constants representing the allowable values for AlarmSuppressionTargetTargetTypeEnum
const (
	AlarmSuppressionTargetTargetTypeAlarm       AlarmSuppressionTargetTargetTypeEnum = "ALARM"
	AlarmSuppressionTargetTargetTypeCompartment AlarmSuppressionTargetTargetTypeEnum = "COMPARTMENT"
)

var mappingAlarmSuppressionTargetTargetTypeEnum = map[string]AlarmSuppressionTargetTargetTypeEnum{
	"ALARM":       AlarmSuppressionTargetTargetTypeAlarm,
	"COMPARTMENT": AlarmSuppressionTargetTargetTypeCompartment,
}

var mappingAlarmSuppressionTargetTargetTypeEnumLowerCase = map[string]AlarmSuppressionTargetTargetTypeEnum{
	"alarm":       AlarmSuppressionTargetTargetTypeAlarm,
	"compartment": AlarmSuppressionTargetTargetTypeCompartment,
}

// GetAlarmSuppressionTargetTargetTypeEnumValues Enumerates the set of values for AlarmSuppressionTargetTargetTypeEnum
func GetAlarmSuppressionTargetTargetTypeEnumValues() []AlarmSuppressionTargetTargetTypeEnum {
	values := make([]AlarmSuppressionTargetTargetTypeEnum, 0)
	for _, v := range mappingAlarmSuppressionTargetTargetTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAlarmSuppressionTargetTargetTypeEnumStringValues Enumerates the set of values in String for AlarmSuppressionTargetTargetTypeEnum
func GetAlarmSuppressionTargetTargetTypeEnumStringValues() []string {
	return []string{
		"ALARM",
		"COMPARTMENT",
	}
}

// GetMappingAlarmSuppressionTargetTargetTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAlarmSuppressionTargetTargetTypeEnum(val string) (AlarmSuppressionTargetTargetTypeEnum, bool) {
	enum, ok := mappingAlarmSuppressionTargetTargetTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
