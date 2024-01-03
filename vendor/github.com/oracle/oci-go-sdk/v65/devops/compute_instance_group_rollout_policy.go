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

// ComputeInstanceGroupRolloutPolicy Specifies the rollout policy for compute instance group stages.
type ComputeInstanceGroupRolloutPolicy interface {

	// The duration of delay between batch rollout. The default delay is 1 minute.
	GetBatchDelayInSeconds() *int
}

type computeinstancegrouprolloutpolicy struct {
	JsonData            []byte
	BatchDelayInSeconds *int   `mandatory:"false" json:"batchDelayInSeconds"`
	PolicyType          string `json:"policyType"`
}

// UnmarshalJSON unmarshals json
func (m *computeinstancegrouprolloutpolicy) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercomputeinstancegrouprolloutpolicy computeinstancegrouprolloutpolicy
	s := struct {
		Model Unmarshalercomputeinstancegrouprolloutpolicy
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.BatchDelayInSeconds = s.Model.BatchDelayInSeconds
	m.PolicyType = s.Model.PolicyType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *computeinstancegrouprolloutpolicy) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.PolicyType {
	case "COMPUTE_INSTANCE_GROUP_LINEAR_ROLLOUT_POLICY_BY_PERCENTAGE":
		mm := ComputeInstanceGroupLinearRolloutPolicyByPercentage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "COMPUTE_INSTANCE_GROUP_LINEAR_ROLLOUT_POLICY_BY_COUNT":
		mm := ComputeInstanceGroupLinearRolloutPolicyByCount{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ComputeInstanceGroupRolloutPolicy: %s.", m.PolicyType)
		return *m, nil
	}
}

// GetBatchDelayInSeconds returns BatchDelayInSeconds
func (m computeinstancegrouprolloutpolicy) GetBatchDelayInSeconds() *int {
	return m.BatchDelayInSeconds
}

func (m computeinstancegrouprolloutpolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m computeinstancegrouprolloutpolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ComputeInstanceGroupRolloutPolicyPolicyTypeEnum Enum with underlying type: string
type ComputeInstanceGroupRolloutPolicyPolicyTypeEnum string

// Set of constants representing the allowable values for ComputeInstanceGroupRolloutPolicyPolicyTypeEnum
const (
	ComputeInstanceGroupRolloutPolicyPolicyTypeCount      ComputeInstanceGroupRolloutPolicyPolicyTypeEnum = "COMPUTE_INSTANCE_GROUP_LINEAR_ROLLOUT_POLICY_BY_COUNT"
	ComputeInstanceGroupRolloutPolicyPolicyTypePercentage ComputeInstanceGroupRolloutPolicyPolicyTypeEnum = "COMPUTE_INSTANCE_GROUP_LINEAR_ROLLOUT_POLICY_BY_PERCENTAGE"
)

var mappingComputeInstanceGroupRolloutPolicyPolicyTypeEnum = map[string]ComputeInstanceGroupRolloutPolicyPolicyTypeEnum{
	"COMPUTE_INSTANCE_GROUP_LINEAR_ROLLOUT_POLICY_BY_COUNT":      ComputeInstanceGroupRolloutPolicyPolicyTypeCount,
	"COMPUTE_INSTANCE_GROUP_LINEAR_ROLLOUT_POLICY_BY_PERCENTAGE": ComputeInstanceGroupRolloutPolicyPolicyTypePercentage,
}

var mappingComputeInstanceGroupRolloutPolicyPolicyTypeEnumLowerCase = map[string]ComputeInstanceGroupRolloutPolicyPolicyTypeEnum{
	"compute_instance_group_linear_rollout_policy_by_count":      ComputeInstanceGroupRolloutPolicyPolicyTypeCount,
	"compute_instance_group_linear_rollout_policy_by_percentage": ComputeInstanceGroupRolloutPolicyPolicyTypePercentage,
}

// GetComputeInstanceGroupRolloutPolicyPolicyTypeEnumValues Enumerates the set of values for ComputeInstanceGroupRolloutPolicyPolicyTypeEnum
func GetComputeInstanceGroupRolloutPolicyPolicyTypeEnumValues() []ComputeInstanceGroupRolloutPolicyPolicyTypeEnum {
	values := make([]ComputeInstanceGroupRolloutPolicyPolicyTypeEnum, 0)
	for _, v := range mappingComputeInstanceGroupRolloutPolicyPolicyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetComputeInstanceGroupRolloutPolicyPolicyTypeEnumStringValues Enumerates the set of values in String for ComputeInstanceGroupRolloutPolicyPolicyTypeEnum
func GetComputeInstanceGroupRolloutPolicyPolicyTypeEnumStringValues() []string {
	return []string{
		"COMPUTE_INSTANCE_GROUP_LINEAR_ROLLOUT_POLICY_BY_COUNT",
		"COMPUTE_INSTANCE_GROUP_LINEAR_ROLLOUT_POLICY_BY_PERCENTAGE",
	}
}

// GetMappingComputeInstanceGroupRolloutPolicyPolicyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingComputeInstanceGroupRolloutPolicyPolicyTypeEnum(val string) (ComputeInstanceGroupRolloutPolicyPolicyTypeEnum, bool) {
	enum, ok := mappingComputeInstanceGroupRolloutPolicyPolicyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
