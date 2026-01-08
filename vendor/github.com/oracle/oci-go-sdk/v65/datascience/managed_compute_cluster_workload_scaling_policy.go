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

// ManagedComputeClusterWorkloadScalingPolicy The scaling policy to apply to workloads on managed compute cluster type compute target.
type ManagedComputeClusterWorkloadScalingPolicy interface {
}

type managedcomputeclusterworkloadscalingpolicy struct {
	JsonData   []byte
	PolicyType string `json:"policyType"`
}

// UnmarshalJSON unmarshals json
func (m *managedcomputeclusterworkloadscalingpolicy) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalermanagedcomputeclusterworkloadscalingpolicy managedcomputeclusterworkloadscalingpolicy
	s := struct {
		Model Unmarshalermanagedcomputeclusterworkloadscalingpolicy
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.PolicyType = s.Model.PolicyType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *managedcomputeclusterworkloadscalingpolicy) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.PolicyType {
	case "AUTOSCALING":
		mm := ManagedComputeClusterWorkloadAutoScalingPolicy{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FIXED_SIZE":
		mm := ManagedComputeClusterWorkloadFixedSizeScalingPolicy{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for ManagedComputeClusterWorkloadScalingPolicy: %s.", m.PolicyType)
		return *m, nil
	}
}

func (m managedcomputeclusterworkloadscalingpolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m managedcomputeclusterworkloadscalingpolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ManagedComputeClusterWorkloadScalingPolicyPolicyTypeEnum Enum with underlying type: string
type ManagedComputeClusterWorkloadScalingPolicyPolicyTypeEnum string

// Set of constants representing the allowable values for ManagedComputeClusterWorkloadScalingPolicyPolicyTypeEnum
const (
	ManagedComputeClusterWorkloadScalingPolicyPolicyTypeFixedSize   ManagedComputeClusterWorkloadScalingPolicyPolicyTypeEnum = "FIXED_SIZE"
	ManagedComputeClusterWorkloadScalingPolicyPolicyTypeAutoscaling ManagedComputeClusterWorkloadScalingPolicyPolicyTypeEnum = "AUTOSCALING"
)

var mappingManagedComputeClusterWorkloadScalingPolicyPolicyTypeEnum = map[string]ManagedComputeClusterWorkloadScalingPolicyPolicyTypeEnum{
	"FIXED_SIZE":  ManagedComputeClusterWorkloadScalingPolicyPolicyTypeFixedSize,
	"AUTOSCALING": ManagedComputeClusterWorkloadScalingPolicyPolicyTypeAutoscaling,
}

var mappingManagedComputeClusterWorkloadScalingPolicyPolicyTypeEnumLowerCase = map[string]ManagedComputeClusterWorkloadScalingPolicyPolicyTypeEnum{
	"fixed_size":  ManagedComputeClusterWorkloadScalingPolicyPolicyTypeFixedSize,
	"autoscaling": ManagedComputeClusterWorkloadScalingPolicyPolicyTypeAutoscaling,
}

// GetManagedComputeClusterWorkloadScalingPolicyPolicyTypeEnumValues Enumerates the set of values for ManagedComputeClusterWorkloadScalingPolicyPolicyTypeEnum
func GetManagedComputeClusterWorkloadScalingPolicyPolicyTypeEnumValues() []ManagedComputeClusterWorkloadScalingPolicyPolicyTypeEnum {
	values := make([]ManagedComputeClusterWorkloadScalingPolicyPolicyTypeEnum, 0)
	for _, v := range mappingManagedComputeClusterWorkloadScalingPolicyPolicyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetManagedComputeClusterWorkloadScalingPolicyPolicyTypeEnumStringValues Enumerates the set of values in String for ManagedComputeClusterWorkloadScalingPolicyPolicyTypeEnum
func GetManagedComputeClusterWorkloadScalingPolicyPolicyTypeEnumStringValues() []string {
	return []string{
		"FIXED_SIZE",
		"AUTOSCALING",
	}
}

// GetMappingManagedComputeClusterWorkloadScalingPolicyPolicyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingManagedComputeClusterWorkloadScalingPolicyPolicyTypeEnum(val string) (ManagedComputeClusterWorkloadScalingPolicyPolicyTypeEnum, bool) {
	enum, ok := mappingManagedComputeClusterWorkloadScalingPolicyPolicyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
