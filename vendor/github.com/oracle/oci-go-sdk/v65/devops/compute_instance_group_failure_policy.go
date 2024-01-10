// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ComputeInstanceGroupFailurePolicy Specifies a failure policy for a compute instance group rolling deployment stage.
type ComputeInstanceGroupFailurePolicy interface {
}

type computeinstancegroupfailurepolicy struct {
	JsonData   []byte
	PolicyType string `json:"policyType"`
}

// UnmarshalJSON unmarshals json
func (m *computeinstancegroupfailurepolicy) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercomputeinstancegroupfailurepolicy computeinstancegroupfailurepolicy
	s := struct {
		Model Unmarshalercomputeinstancegroupfailurepolicy
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.PolicyType = s.Model.PolicyType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *computeinstancegroupfailurepolicy) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.PolicyType {
	case "COMPUTE_INSTANCE_GROUP_FAILURE_POLICY_BY_PERCENTAGE":
		mm := ComputeInstanceGroupFailurePolicyByPercentage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "COMPUTE_INSTANCE_GROUP_FAILURE_POLICY_BY_COUNT":
		mm := ComputeInstanceGroupFailurePolicyByCount{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ComputeInstanceGroupFailurePolicy: %s.", m.PolicyType)
		return *m, nil
	}
}

func (m computeinstancegroupfailurepolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m computeinstancegroupfailurepolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ComputeInstanceGroupFailurePolicyPolicyTypeEnum Enum with underlying type: string
type ComputeInstanceGroupFailurePolicyPolicyTypeEnum string

// Set of constants representing the allowable values for ComputeInstanceGroupFailurePolicyPolicyTypeEnum
const (
	ComputeInstanceGroupFailurePolicyPolicyTypeCount      ComputeInstanceGroupFailurePolicyPolicyTypeEnum = "COMPUTE_INSTANCE_GROUP_FAILURE_POLICY_BY_COUNT"
	ComputeInstanceGroupFailurePolicyPolicyTypePercentage ComputeInstanceGroupFailurePolicyPolicyTypeEnum = "COMPUTE_INSTANCE_GROUP_FAILURE_POLICY_BY_PERCENTAGE"
)

var mappingComputeInstanceGroupFailurePolicyPolicyTypeEnum = map[string]ComputeInstanceGroupFailurePolicyPolicyTypeEnum{
	"COMPUTE_INSTANCE_GROUP_FAILURE_POLICY_BY_COUNT":      ComputeInstanceGroupFailurePolicyPolicyTypeCount,
	"COMPUTE_INSTANCE_GROUP_FAILURE_POLICY_BY_PERCENTAGE": ComputeInstanceGroupFailurePolicyPolicyTypePercentage,
}

var mappingComputeInstanceGroupFailurePolicyPolicyTypeEnumLowerCase = map[string]ComputeInstanceGroupFailurePolicyPolicyTypeEnum{
	"compute_instance_group_failure_policy_by_count":      ComputeInstanceGroupFailurePolicyPolicyTypeCount,
	"compute_instance_group_failure_policy_by_percentage": ComputeInstanceGroupFailurePolicyPolicyTypePercentage,
}

// GetComputeInstanceGroupFailurePolicyPolicyTypeEnumValues Enumerates the set of values for ComputeInstanceGroupFailurePolicyPolicyTypeEnum
func GetComputeInstanceGroupFailurePolicyPolicyTypeEnumValues() []ComputeInstanceGroupFailurePolicyPolicyTypeEnum {
	values := make([]ComputeInstanceGroupFailurePolicyPolicyTypeEnum, 0)
	for _, v := range mappingComputeInstanceGroupFailurePolicyPolicyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetComputeInstanceGroupFailurePolicyPolicyTypeEnumStringValues Enumerates the set of values in String for ComputeInstanceGroupFailurePolicyPolicyTypeEnum
func GetComputeInstanceGroupFailurePolicyPolicyTypeEnumStringValues() []string {
	return []string{
		"COMPUTE_INSTANCE_GROUP_FAILURE_POLICY_BY_COUNT",
		"COMPUTE_INSTANCE_GROUP_FAILURE_POLICY_BY_PERCENTAGE",
	}
}

// GetMappingComputeInstanceGroupFailurePolicyPolicyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingComputeInstanceGroupFailurePolicyPolicyTypeEnum(val string) (ComputeInstanceGroupFailurePolicyPolicyTypeEnum, bool) {
	enum, ok := mappingComputeInstanceGroupFailurePolicyPolicyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
