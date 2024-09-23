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

// SuppressionCondition Precondition for an alarm suppression within the suppression date and time range (`timeSuppressFrom` to `timeSuppressUntil`).
type SuppressionCondition interface {
}

type suppressioncondition struct {
	JsonData      []byte
	ConditionType string `json:"conditionType"`
}

// UnmarshalJSON unmarshals json
func (m *suppressioncondition) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalersuppressioncondition suppressioncondition
	s := struct {
		Model Unmarshalersuppressioncondition
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ConditionType = s.Model.ConditionType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *suppressioncondition) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ConditionType {
	case "RECURRENCE":
		mm := Recurrence{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for SuppressionCondition: %s.", m.ConditionType)
		return *m, nil
	}
}

func (m suppressioncondition) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m suppressioncondition) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SuppressionConditionConditionTypeEnum Enum with underlying type: string
type SuppressionConditionConditionTypeEnum string

// Set of constants representing the allowable values for SuppressionConditionConditionTypeEnum
const (
	SuppressionConditionConditionTypeRecurrence SuppressionConditionConditionTypeEnum = "RECURRENCE"
)

var mappingSuppressionConditionConditionTypeEnum = map[string]SuppressionConditionConditionTypeEnum{
	"RECURRENCE": SuppressionConditionConditionTypeRecurrence,
}

var mappingSuppressionConditionConditionTypeEnumLowerCase = map[string]SuppressionConditionConditionTypeEnum{
	"recurrence": SuppressionConditionConditionTypeRecurrence,
}

// GetSuppressionConditionConditionTypeEnumValues Enumerates the set of values for SuppressionConditionConditionTypeEnum
func GetSuppressionConditionConditionTypeEnumValues() []SuppressionConditionConditionTypeEnum {
	values := make([]SuppressionConditionConditionTypeEnum, 0)
	for _, v := range mappingSuppressionConditionConditionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSuppressionConditionConditionTypeEnumStringValues Enumerates the set of values in String for SuppressionConditionConditionTypeEnum
func GetSuppressionConditionConditionTypeEnumStringValues() []string {
	return []string{
		"RECURRENCE",
	}
}

// GetMappingSuppressionConditionConditionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSuppressionConditionConditionTypeEnum(val string) (SuppressionConditionConditionTypeEnum, bool) {
	enum, ok := mappingSuppressionConditionConditionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
