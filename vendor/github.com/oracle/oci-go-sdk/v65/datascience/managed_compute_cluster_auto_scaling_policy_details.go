// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// ManagedComputeClusterAutoScalingPolicyDetails Details for an autoscaling policy to enable on the managed compute cluster type compute target . Each autoscaling configuration can have one autoscaling policy.
type ManagedComputeClusterAutoScalingPolicyDetails interface {
}

type managedcomputeclusterautoscalingpolicydetails struct {
	JsonData              []byte
	AutoScalingPolicyType string `json:"autoScalingPolicyType"`
}

// UnmarshalJSON unmarshals json
func (m *managedcomputeclusterautoscalingpolicydetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalermanagedcomputeclusterautoscalingpolicydetails managedcomputeclusterautoscalingpolicydetails
	s := struct {
		Model Unmarshalermanagedcomputeclusterautoscalingpolicydetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.AutoScalingPolicyType = s.Model.AutoScalingPolicyType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *managedcomputeclusterautoscalingpolicydetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.AutoScalingPolicyType {
	case "THRESHOLD":
		mm := ManagedComputeClusterThresholdBasedAutoScalingPolicyDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for ManagedComputeClusterAutoScalingPolicyDetails: %s.", m.AutoScalingPolicyType)
		return *m, nil
	}
}

func (m managedcomputeclusterautoscalingpolicydetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m managedcomputeclusterautoscalingpolicydetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ManagedComputeClusterAutoScalingPolicyDetailsAutoScalingPolicyTypeEnum Enum with underlying type: string
type ManagedComputeClusterAutoScalingPolicyDetailsAutoScalingPolicyTypeEnum string

// Set of constants representing the allowable values for ManagedComputeClusterAutoScalingPolicyDetailsAutoScalingPolicyTypeEnum
const (
	ManagedComputeClusterAutoScalingPolicyDetailsAutoScalingPolicyTypeThreshold ManagedComputeClusterAutoScalingPolicyDetailsAutoScalingPolicyTypeEnum = "THRESHOLD"
)

var mappingManagedComputeClusterAutoScalingPolicyDetailsAutoScalingPolicyTypeEnum = map[string]ManagedComputeClusterAutoScalingPolicyDetailsAutoScalingPolicyTypeEnum{
	"THRESHOLD": ManagedComputeClusterAutoScalingPolicyDetailsAutoScalingPolicyTypeThreshold,
}

var mappingManagedComputeClusterAutoScalingPolicyDetailsAutoScalingPolicyTypeEnumLowerCase = map[string]ManagedComputeClusterAutoScalingPolicyDetailsAutoScalingPolicyTypeEnum{
	"threshold": ManagedComputeClusterAutoScalingPolicyDetailsAutoScalingPolicyTypeThreshold,
}

// GetManagedComputeClusterAutoScalingPolicyDetailsAutoScalingPolicyTypeEnumValues Enumerates the set of values for ManagedComputeClusterAutoScalingPolicyDetailsAutoScalingPolicyTypeEnum
func GetManagedComputeClusterAutoScalingPolicyDetailsAutoScalingPolicyTypeEnumValues() []ManagedComputeClusterAutoScalingPolicyDetailsAutoScalingPolicyTypeEnum {
	values := make([]ManagedComputeClusterAutoScalingPolicyDetailsAutoScalingPolicyTypeEnum, 0)
	for _, v := range mappingManagedComputeClusterAutoScalingPolicyDetailsAutoScalingPolicyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetManagedComputeClusterAutoScalingPolicyDetailsAutoScalingPolicyTypeEnumStringValues Enumerates the set of values in String for ManagedComputeClusterAutoScalingPolicyDetailsAutoScalingPolicyTypeEnum
func GetManagedComputeClusterAutoScalingPolicyDetailsAutoScalingPolicyTypeEnumStringValues() []string {
	return []string{
		"THRESHOLD",
	}
}

// GetMappingManagedComputeClusterAutoScalingPolicyDetailsAutoScalingPolicyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingManagedComputeClusterAutoScalingPolicyDetailsAutoScalingPolicyTypeEnum(val string) (ManagedComputeClusterAutoScalingPolicyDetailsAutoScalingPolicyTypeEnum, bool) {
	enum, ok := mappingManagedComputeClusterAutoScalingPolicyDetailsAutoScalingPolicyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
