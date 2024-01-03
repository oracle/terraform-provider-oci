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

// MetricBasedHorizontalScalingPolicyDetails Details of a metric based horizontal autoscaling policy.
// In a metric-based autoscaling policy, an autoscaling action is triggered when a performance metric exceeds a threshold.
type MetricBasedHorizontalScalingPolicyDetails struct {
	ScaleOutConfig *MetricBasedHorizontalScaleOutConfig `mandatory:"false" json:"scaleOutConfig"`

	ScaleInConfig *MetricBasedHorizontalScaleInConfig `mandatory:"false" json:"scaleInConfig"`

	// The type of autoscaling trigger.
	TriggerType AutoScalePolicyDetailsTriggerTypeEnum `mandatory:"true" json:"triggerType"`

	// The type of autoscaling action to take.
	ActionType AutoScalePolicyDetailsActionTypeEnum `mandatory:"true" json:"actionType"`
}

// GetTriggerType returns TriggerType
func (m MetricBasedHorizontalScalingPolicyDetails) GetTriggerType() AutoScalePolicyDetailsTriggerTypeEnum {
	return m.TriggerType
}

// GetActionType returns ActionType
func (m MetricBasedHorizontalScalingPolicyDetails) GetActionType() AutoScalePolicyDetailsActionTypeEnum {
	return m.ActionType
}

func (m MetricBasedHorizontalScalingPolicyDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MetricBasedHorizontalScalingPolicyDetails) ValidateEnumValue() (bool, error) {
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

// MarshalJSON marshals to json representation
func (m MetricBasedHorizontalScalingPolicyDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeMetricBasedHorizontalScalingPolicyDetails MetricBasedHorizontalScalingPolicyDetails
	s := struct {
		DiscriminatorParam string `json:"policyType"`
		MarshalTypeMetricBasedHorizontalScalingPolicyDetails
	}{
		"METRIC_BASED_HORIZONTAL_SCALING_POLICY",
		(MarshalTypeMetricBasedHorizontalScalingPolicyDetails)(m),
	}

	return json.Marshal(&s)
}
