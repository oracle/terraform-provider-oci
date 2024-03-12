// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// ScalingConfiguration Scaling configuration for the metric expression rule.
type ScalingConfiguration interface {

	// The period of time that the condition defined in the alarm must persist before the alarm state
	// changes from "OK" to "FIRING" or vice versa. For example, a value of 5 minutes means that the
	// alarm must persist in breaching the condition for five minutes before the alarm updates its
	// state to "FIRING"; likewise, the alarm must persist in not breaching the condition for five
	// minutes before the alarm updates its state to "OK."
	// The duration is specified as a string in ISO 8601 format (`PT10M` for ten minutes or `PT1H`
	// for one hour). Minimum: PT3M. Maximum: PT1H. Default: PT3M.
	GetPendingDuration() *string

	// The value is used for adjusting the count of instances by.
	GetInstanceCountAdjustment() *int
}

type scalingconfiguration struct {
	JsonData                 []byte
	PendingDuration          *string `mandatory:"false" json:"pendingDuration"`
	InstanceCountAdjustment  *int    `mandatory:"false" json:"instanceCountAdjustment"`
	ScalingConfigurationType string  `json:"scalingConfigurationType"`
}

// UnmarshalJSON unmarshals json
func (m *scalingconfiguration) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerscalingconfiguration scalingconfiguration
	s := struct {
		Model Unmarshalerscalingconfiguration
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.PendingDuration = s.Model.PendingDuration
	m.InstanceCountAdjustment = s.Model.InstanceCountAdjustment
	m.ScalingConfigurationType = s.Model.ScalingConfigurationType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *scalingconfiguration) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ScalingConfigurationType {
	case "QUERY":
		mm := CustomExpressionQueryScalingConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "THRESHOLD":
		mm := PredefinedExpressionThresholdScalingConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ScalingConfiguration: %s.", m.ScalingConfigurationType)
		return *m, nil
	}
}

// GetPendingDuration returns PendingDuration
func (m scalingconfiguration) GetPendingDuration() *string {
	return m.PendingDuration
}

// GetInstanceCountAdjustment returns InstanceCountAdjustment
func (m scalingconfiguration) GetInstanceCountAdjustment() *int {
	return m.InstanceCountAdjustment
}

func (m scalingconfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m scalingconfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ScalingConfigurationScalingConfigurationTypeEnum Enum with underlying type: string
type ScalingConfigurationScalingConfigurationTypeEnum string

// Set of constants representing the allowable values for ScalingConfigurationScalingConfigurationTypeEnum
const (
	ScalingConfigurationScalingConfigurationTypeThreshold ScalingConfigurationScalingConfigurationTypeEnum = "THRESHOLD"
	ScalingConfigurationScalingConfigurationTypeQuery     ScalingConfigurationScalingConfigurationTypeEnum = "QUERY"
)

var mappingScalingConfigurationScalingConfigurationTypeEnum = map[string]ScalingConfigurationScalingConfigurationTypeEnum{
	"THRESHOLD": ScalingConfigurationScalingConfigurationTypeThreshold,
	"QUERY":     ScalingConfigurationScalingConfigurationTypeQuery,
}

var mappingScalingConfigurationScalingConfigurationTypeEnumLowerCase = map[string]ScalingConfigurationScalingConfigurationTypeEnum{
	"threshold": ScalingConfigurationScalingConfigurationTypeThreshold,
	"query":     ScalingConfigurationScalingConfigurationTypeQuery,
}

// GetScalingConfigurationScalingConfigurationTypeEnumValues Enumerates the set of values for ScalingConfigurationScalingConfigurationTypeEnum
func GetScalingConfigurationScalingConfigurationTypeEnumValues() []ScalingConfigurationScalingConfigurationTypeEnum {
	values := make([]ScalingConfigurationScalingConfigurationTypeEnum, 0)
	for _, v := range mappingScalingConfigurationScalingConfigurationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetScalingConfigurationScalingConfigurationTypeEnumStringValues Enumerates the set of values in String for ScalingConfigurationScalingConfigurationTypeEnum
func GetScalingConfigurationScalingConfigurationTypeEnumStringValues() []string {
	return []string{
		"THRESHOLD",
		"QUERY",
	}
}

// GetMappingScalingConfigurationScalingConfigurationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingScalingConfigurationScalingConfigurationTypeEnum(val string) (ScalingConfigurationScalingConfigurationTypeEnum, bool) {
	enum, ok := mappingScalingConfigurationScalingConfigurationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
