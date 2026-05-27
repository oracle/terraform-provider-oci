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

// ManagedComputeClusterWorkloadAutoScalingPolicyDetails Details for an autoscaling policy to enable autoscaling of workload on the managed compute cluster type compute target. Each autoscaling configuration can have one autoscaling policy.
type ManagedComputeClusterWorkloadAutoScalingPolicyDetails interface {
}

type managedcomputeclusterworkloadautoscalingpolicydetails struct {
	JsonData              []byte
	AutoScalingPolicyType string `json:"autoScalingPolicyType"`
}

// UnmarshalJSON unmarshals json
func (m *managedcomputeclusterworkloadautoscalingpolicydetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalermanagedcomputeclusterworkloadautoscalingpolicydetails managedcomputeclusterworkloadautoscalingpolicydetails
	s := struct {
		Model Unmarshalermanagedcomputeclusterworkloadautoscalingpolicydetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.AutoScalingPolicyType = s.Model.AutoScalingPolicyType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *managedcomputeclusterworkloadautoscalingpolicydetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.AutoScalingPolicyType {
	case "THRESHOLD":
		mm := ManagedComputeClusterWorkloadThresholdBasedPolicyDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for ManagedComputeClusterWorkloadAutoScalingPolicyDetails: %s.", m.AutoScalingPolicyType)
		return *m, nil
	}
}

func (m managedcomputeclusterworkloadautoscalingpolicydetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m managedcomputeclusterworkloadautoscalingpolicydetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ManagedComputeClusterWorkloadAutoScalingPolicyDetailsAutoScalingPolicyTypeEnum Enum with underlying type: string
type ManagedComputeClusterWorkloadAutoScalingPolicyDetailsAutoScalingPolicyTypeEnum string

// Set of constants representing the allowable values for ManagedComputeClusterWorkloadAutoScalingPolicyDetailsAutoScalingPolicyTypeEnum
const (
	ManagedComputeClusterWorkloadAutoScalingPolicyDetailsAutoScalingPolicyTypeThreshold ManagedComputeClusterWorkloadAutoScalingPolicyDetailsAutoScalingPolicyTypeEnum = "THRESHOLD"
)

var mappingManagedComputeClusterWorkloadAutoScalingPolicyDetailsAutoScalingPolicyTypeEnum = map[string]ManagedComputeClusterWorkloadAutoScalingPolicyDetailsAutoScalingPolicyTypeEnum{
	"THRESHOLD": ManagedComputeClusterWorkloadAutoScalingPolicyDetailsAutoScalingPolicyTypeThreshold,
}

var mappingManagedComputeClusterWorkloadAutoScalingPolicyDetailsAutoScalingPolicyTypeEnumLowerCase = map[string]ManagedComputeClusterWorkloadAutoScalingPolicyDetailsAutoScalingPolicyTypeEnum{
	"threshold": ManagedComputeClusterWorkloadAutoScalingPolicyDetailsAutoScalingPolicyTypeThreshold,
}

// GetManagedComputeClusterWorkloadAutoScalingPolicyDetailsAutoScalingPolicyTypeEnumValues Enumerates the set of values for ManagedComputeClusterWorkloadAutoScalingPolicyDetailsAutoScalingPolicyTypeEnum
func GetManagedComputeClusterWorkloadAutoScalingPolicyDetailsAutoScalingPolicyTypeEnumValues() []ManagedComputeClusterWorkloadAutoScalingPolicyDetailsAutoScalingPolicyTypeEnum {
	values := make([]ManagedComputeClusterWorkloadAutoScalingPolicyDetailsAutoScalingPolicyTypeEnum, 0)
	for _, v := range mappingManagedComputeClusterWorkloadAutoScalingPolicyDetailsAutoScalingPolicyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetManagedComputeClusterWorkloadAutoScalingPolicyDetailsAutoScalingPolicyTypeEnumStringValues Enumerates the set of values in String for ManagedComputeClusterWorkloadAutoScalingPolicyDetailsAutoScalingPolicyTypeEnum
func GetManagedComputeClusterWorkloadAutoScalingPolicyDetailsAutoScalingPolicyTypeEnumStringValues() []string {
	return []string{
		"THRESHOLD",
	}
}

// GetMappingManagedComputeClusterWorkloadAutoScalingPolicyDetailsAutoScalingPolicyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingManagedComputeClusterWorkloadAutoScalingPolicyDetailsAutoScalingPolicyTypeEnum(val string) (ManagedComputeClusterWorkloadAutoScalingPolicyDetailsAutoScalingPolicyTypeEnum, bool) {
	enum, ok := mappingManagedComputeClusterWorkloadAutoScalingPolicyDetailsAutoScalingPolicyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
