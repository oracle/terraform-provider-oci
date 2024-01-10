// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// REST API for Oracle Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service clusters. Build on Hadoop, Spark and Data Science distributions, which can be fully integrated with existing enterprise data in Oracle Database and Oracle applications.
//

package bds

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AutoScalePolicyDetails Details of an autoscale policy.
// You can create following types of autoscaling policies:
// - **MetricBasedVerticalScalingPolicy:** Vertical autoscaling action is triggered when a performance metric exceeds a threshold
// - **MetricBasedHorizontalScalingPolicy:** Horizontal autoscaling action is triggered when a performance metric exceeds a threshold
// - **ScheduleBasedVerticalScalingPolicy:** Vertical autoscaling action is triggered at the specific times that you schedule.
// - **ScheduleBasedHorizontalScalingPolicy:** Horizontal autoscaling action is triggered at the specific times that you schedule.
type AutoScalePolicyDetails interface {

	// The type of autoscaling trigger.
	GetTriggerType() AutoScalePolicyDetailsTriggerTypeEnum

	// The type of autoscaling action to take.
	GetActionType() AutoScalePolicyDetailsActionTypeEnum
}

type autoscalepolicydetails struct {
	JsonData    []byte
	TriggerType AutoScalePolicyDetailsTriggerTypeEnum `mandatory:"true" json:"triggerType"`
	ActionType  AutoScalePolicyDetailsActionTypeEnum  `mandatory:"true" json:"actionType"`
	PolicyType  string                                `json:"policyType"`
}

// UnmarshalJSON unmarshals json
func (m *autoscalepolicydetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerautoscalepolicydetails autoscalepolicydetails
	s := struct {
		Model Unmarshalerautoscalepolicydetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.TriggerType = s.Model.TriggerType
	m.ActionType = s.Model.ActionType
	m.PolicyType = s.Model.PolicyType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *autoscalepolicydetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.PolicyType {
	case "METRIC_BASED_VERTICAL_SCALING_POLICY":
		mm := MetricBasedVerticalScalingPolicyDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SCHEDULE_BASED_VERTICAL_SCALING_POLICY":
		mm := ScheduleBasedVerticalScalingPolicyDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SCHEDULE_BASED_HORIZONTAL_SCALING_POLICY":
		mm := ScheduleBasedHorizontalScalingPolicyDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "METRIC_BASED_HORIZONTAL_SCALING_POLICY":
		mm := MetricBasedHorizontalScalingPolicyDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for AutoScalePolicyDetails: %s.", m.PolicyType)
		return *m, nil
	}
}

// GetTriggerType returns TriggerType
func (m autoscalepolicydetails) GetTriggerType() AutoScalePolicyDetailsTriggerTypeEnum {
	return m.TriggerType
}

// GetActionType returns ActionType
func (m autoscalepolicydetails) GetActionType() AutoScalePolicyDetailsActionTypeEnum {
	return m.ActionType
}

func (m autoscalepolicydetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m autoscalepolicydetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAutoScalePolicyDetailsTriggerTypeEnum(string(m.TriggerType)); !ok && m.TriggerType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TriggerType: %s. Supported values are: %s.", m.TriggerType, strings.Join(GetAutoScalePolicyDetailsTriggerTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutoScalePolicyDetailsActionTypeEnum(string(m.ActionType)); !ok && m.ActionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ActionType: %s. Supported values are: %s.", m.ActionType, strings.Join(GetAutoScalePolicyDetailsActionTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AutoScalePolicyDetailsTriggerTypeEnum Enum with underlying type: string
type AutoScalePolicyDetailsTriggerTypeEnum string

// Set of constants representing the allowable values for AutoScalePolicyDetailsTriggerTypeEnum
const (
	AutoScalePolicyDetailsTriggerTypeMetricBased   AutoScalePolicyDetailsTriggerTypeEnum = "METRIC_BASED"
	AutoScalePolicyDetailsTriggerTypeScheduleBased AutoScalePolicyDetailsTriggerTypeEnum = "SCHEDULE_BASED"
)

var mappingAutoScalePolicyDetailsTriggerTypeEnum = map[string]AutoScalePolicyDetailsTriggerTypeEnum{
	"METRIC_BASED":   AutoScalePolicyDetailsTriggerTypeMetricBased,
	"SCHEDULE_BASED": AutoScalePolicyDetailsTriggerTypeScheduleBased,
}

var mappingAutoScalePolicyDetailsTriggerTypeEnumLowerCase = map[string]AutoScalePolicyDetailsTriggerTypeEnum{
	"metric_based":   AutoScalePolicyDetailsTriggerTypeMetricBased,
	"schedule_based": AutoScalePolicyDetailsTriggerTypeScheduleBased,
}

// GetAutoScalePolicyDetailsTriggerTypeEnumValues Enumerates the set of values for AutoScalePolicyDetailsTriggerTypeEnum
func GetAutoScalePolicyDetailsTriggerTypeEnumValues() []AutoScalePolicyDetailsTriggerTypeEnum {
	values := make([]AutoScalePolicyDetailsTriggerTypeEnum, 0)
	for _, v := range mappingAutoScalePolicyDetailsTriggerTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAutoScalePolicyDetailsTriggerTypeEnumStringValues Enumerates the set of values in String for AutoScalePolicyDetailsTriggerTypeEnum
func GetAutoScalePolicyDetailsTriggerTypeEnumStringValues() []string {
	return []string{
		"METRIC_BASED",
		"SCHEDULE_BASED",
	}
}

// GetMappingAutoScalePolicyDetailsTriggerTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutoScalePolicyDetailsTriggerTypeEnum(val string) (AutoScalePolicyDetailsTriggerTypeEnum, bool) {
	enum, ok := mappingAutoScalePolicyDetailsTriggerTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutoScalePolicyDetailsActionTypeEnum Enum with underlying type: string
type AutoScalePolicyDetailsActionTypeEnum string

// Set of constants representing the allowable values for AutoScalePolicyDetailsActionTypeEnum
const (
	AutoScalePolicyDetailsActionTypeVerticalScaling   AutoScalePolicyDetailsActionTypeEnum = "VERTICAL_SCALING"
	AutoScalePolicyDetailsActionTypeHorizontalScaling AutoScalePolicyDetailsActionTypeEnum = "HORIZONTAL_SCALING"
)

var mappingAutoScalePolicyDetailsActionTypeEnum = map[string]AutoScalePolicyDetailsActionTypeEnum{
	"VERTICAL_SCALING":   AutoScalePolicyDetailsActionTypeVerticalScaling,
	"HORIZONTAL_SCALING": AutoScalePolicyDetailsActionTypeHorizontalScaling,
}

var mappingAutoScalePolicyDetailsActionTypeEnumLowerCase = map[string]AutoScalePolicyDetailsActionTypeEnum{
	"vertical_scaling":   AutoScalePolicyDetailsActionTypeVerticalScaling,
	"horizontal_scaling": AutoScalePolicyDetailsActionTypeHorizontalScaling,
}

// GetAutoScalePolicyDetailsActionTypeEnumValues Enumerates the set of values for AutoScalePolicyDetailsActionTypeEnum
func GetAutoScalePolicyDetailsActionTypeEnumValues() []AutoScalePolicyDetailsActionTypeEnum {
	values := make([]AutoScalePolicyDetailsActionTypeEnum, 0)
	for _, v := range mappingAutoScalePolicyDetailsActionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAutoScalePolicyDetailsActionTypeEnumStringValues Enumerates the set of values in String for AutoScalePolicyDetailsActionTypeEnum
func GetAutoScalePolicyDetailsActionTypeEnumStringValues() []string {
	return []string{
		"VERTICAL_SCALING",
		"HORIZONTAL_SCALING",
	}
}

// GetMappingAutoScalePolicyDetailsActionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutoScalePolicyDetailsActionTypeEnum(val string) (AutoScalePolicyDetailsActionTypeEnum, bool) {
	enum, ok := mappingAutoScalePolicyDetailsActionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutoScalePolicyDetailsPolicyTypeEnum Enum with underlying type: string
type AutoScalePolicyDetailsPolicyTypeEnum string

// Set of constants representing the allowable values for AutoScalePolicyDetailsPolicyTypeEnum
const (
	AutoScalePolicyDetailsPolicyTypeMetricBasedVerticalScalingPolicy     AutoScalePolicyDetailsPolicyTypeEnum = "METRIC_BASED_VERTICAL_SCALING_POLICY"
	AutoScalePolicyDetailsPolicyTypeMetricBasedHorizontalScalingPolicy   AutoScalePolicyDetailsPolicyTypeEnum = "METRIC_BASED_HORIZONTAL_SCALING_POLICY"
	AutoScalePolicyDetailsPolicyTypeScheduleBasedVerticalScalingPolicy   AutoScalePolicyDetailsPolicyTypeEnum = "SCHEDULE_BASED_VERTICAL_SCALING_POLICY"
	AutoScalePolicyDetailsPolicyTypeScheduleBasedHorizontalScalingPolicy AutoScalePolicyDetailsPolicyTypeEnum = "SCHEDULE_BASED_HORIZONTAL_SCALING_POLICY"
)

var mappingAutoScalePolicyDetailsPolicyTypeEnum = map[string]AutoScalePolicyDetailsPolicyTypeEnum{
	"METRIC_BASED_VERTICAL_SCALING_POLICY":     AutoScalePolicyDetailsPolicyTypeMetricBasedVerticalScalingPolicy,
	"METRIC_BASED_HORIZONTAL_SCALING_POLICY":   AutoScalePolicyDetailsPolicyTypeMetricBasedHorizontalScalingPolicy,
	"SCHEDULE_BASED_VERTICAL_SCALING_POLICY":   AutoScalePolicyDetailsPolicyTypeScheduleBasedVerticalScalingPolicy,
	"SCHEDULE_BASED_HORIZONTAL_SCALING_POLICY": AutoScalePolicyDetailsPolicyTypeScheduleBasedHorizontalScalingPolicy,
}

var mappingAutoScalePolicyDetailsPolicyTypeEnumLowerCase = map[string]AutoScalePolicyDetailsPolicyTypeEnum{
	"metric_based_vertical_scaling_policy":     AutoScalePolicyDetailsPolicyTypeMetricBasedVerticalScalingPolicy,
	"metric_based_horizontal_scaling_policy":   AutoScalePolicyDetailsPolicyTypeMetricBasedHorizontalScalingPolicy,
	"schedule_based_vertical_scaling_policy":   AutoScalePolicyDetailsPolicyTypeScheduleBasedVerticalScalingPolicy,
	"schedule_based_horizontal_scaling_policy": AutoScalePolicyDetailsPolicyTypeScheduleBasedHorizontalScalingPolicy,
}

// GetAutoScalePolicyDetailsPolicyTypeEnumValues Enumerates the set of values for AutoScalePolicyDetailsPolicyTypeEnum
func GetAutoScalePolicyDetailsPolicyTypeEnumValues() []AutoScalePolicyDetailsPolicyTypeEnum {
	values := make([]AutoScalePolicyDetailsPolicyTypeEnum, 0)
	for _, v := range mappingAutoScalePolicyDetailsPolicyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAutoScalePolicyDetailsPolicyTypeEnumStringValues Enumerates the set of values in String for AutoScalePolicyDetailsPolicyTypeEnum
func GetAutoScalePolicyDetailsPolicyTypeEnumStringValues() []string {
	return []string{
		"METRIC_BASED_VERTICAL_SCALING_POLICY",
		"METRIC_BASED_HORIZONTAL_SCALING_POLICY",
		"SCHEDULE_BASED_VERTICAL_SCALING_POLICY",
		"SCHEDULE_BASED_HORIZONTAL_SCALING_POLICY",
	}
}

// GetMappingAutoScalePolicyDetailsPolicyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutoScalePolicyDetailsPolicyTypeEnum(val string) (AutoScalePolicyDetailsPolicyTypeEnum, bool) {
	enum, ok := mappingAutoScalePolicyDetailsPolicyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
