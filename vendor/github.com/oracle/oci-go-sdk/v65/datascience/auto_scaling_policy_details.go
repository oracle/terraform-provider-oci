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

// AutoScalingPolicyDetails Details for an autoscaling policy to enable on the model deployment.
// Each autoscaling configuration can have one autoscaling policy.
// In a threshold-based autoscaling policy, an autoscaling action is triggered when a performance metric meets
// or exceeds a threshold.
type AutoScalingPolicyDetails interface {
}

type autoscalingpolicydetails struct {
	JsonData              []byte
	AutoScalingPolicyType string `json:"autoScalingPolicyType"`
}

// UnmarshalJSON unmarshals json
func (m *autoscalingpolicydetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerautoscalingpolicydetails autoscalingpolicydetails
	s := struct {
		Model Unmarshalerautoscalingpolicydetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.AutoScalingPolicyType = s.Model.AutoScalingPolicyType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *autoscalingpolicydetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.AutoScalingPolicyType {
	case "THRESHOLD":
		mm := ThresholdBasedAutoScalingPolicyDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for AutoScalingPolicyDetails: %s.", m.AutoScalingPolicyType)
		return *m, nil
	}
}

func (m autoscalingpolicydetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m autoscalingpolicydetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AutoScalingPolicyDetailsAutoScalingPolicyTypeEnum Enum with underlying type: string
type AutoScalingPolicyDetailsAutoScalingPolicyTypeEnum string

// Set of constants representing the allowable values for AutoScalingPolicyDetailsAutoScalingPolicyTypeEnum
const (
	AutoScalingPolicyDetailsAutoScalingPolicyTypeThreshold AutoScalingPolicyDetailsAutoScalingPolicyTypeEnum = "THRESHOLD"
)

var mappingAutoScalingPolicyDetailsAutoScalingPolicyTypeEnum = map[string]AutoScalingPolicyDetailsAutoScalingPolicyTypeEnum{
	"THRESHOLD": AutoScalingPolicyDetailsAutoScalingPolicyTypeThreshold,
}

var mappingAutoScalingPolicyDetailsAutoScalingPolicyTypeEnumLowerCase = map[string]AutoScalingPolicyDetailsAutoScalingPolicyTypeEnum{
	"threshold": AutoScalingPolicyDetailsAutoScalingPolicyTypeThreshold,
}

// GetAutoScalingPolicyDetailsAutoScalingPolicyTypeEnumValues Enumerates the set of values for AutoScalingPolicyDetailsAutoScalingPolicyTypeEnum
func GetAutoScalingPolicyDetailsAutoScalingPolicyTypeEnumValues() []AutoScalingPolicyDetailsAutoScalingPolicyTypeEnum {
	values := make([]AutoScalingPolicyDetailsAutoScalingPolicyTypeEnum, 0)
	for _, v := range mappingAutoScalingPolicyDetailsAutoScalingPolicyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAutoScalingPolicyDetailsAutoScalingPolicyTypeEnumStringValues Enumerates the set of values in String for AutoScalingPolicyDetailsAutoScalingPolicyTypeEnum
func GetAutoScalingPolicyDetailsAutoScalingPolicyTypeEnumStringValues() []string {
	return []string{
		"THRESHOLD",
	}
}

// GetMappingAutoScalingPolicyDetailsAutoScalingPolicyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutoScalingPolicyDetailsAutoScalingPolicyTypeEnum(val string) (AutoScalingPolicyDetailsAutoScalingPolicyTypeEnum, bool) {
	enum, ok := mappingAutoScalingPolicyDetailsAutoScalingPolicyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
