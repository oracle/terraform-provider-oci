// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
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
		return *m, nil
	}
}

//GetBatchDelayInSeconds returns BatchDelayInSeconds
func (m computeinstancegrouprolloutpolicy) GetBatchDelayInSeconds() *int {
	return m.BatchDelayInSeconds
}

func (m computeinstancegrouprolloutpolicy) String() string {
	return common.PointerString(m)
}

// ComputeInstanceGroupRolloutPolicyPolicyTypeEnum Enum with underlying type: string
type ComputeInstanceGroupRolloutPolicyPolicyTypeEnum string

// Set of constants representing the allowable values for ComputeInstanceGroupRolloutPolicyPolicyTypeEnum
const (
	ComputeInstanceGroupRolloutPolicyPolicyTypeCount      ComputeInstanceGroupRolloutPolicyPolicyTypeEnum = "COMPUTE_INSTANCE_GROUP_LINEAR_ROLLOUT_POLICY_BY_COUNT"
	ComputeInstanceGroupRolloutPolicyPolicyTypePercentage ComputeInstanceGroupRolloutPolicyPolicyTypeEnum = "COMPUTE_INSTANCE_GROUP_LINEAR_ROLLOUT_POLICY_BY_PERCENTAGE"
)

var mappingComputeInstanceGroupRolloutPolicyPolicyType = map[string]ComputeInstanceGroupRolloutPolicyPolicyTypeEnum{
	"COMPUTE_INSTANCE_GROUP_LINEAR_ROLLOUT_POLICY_BY_COUNT":      ComputeInstanceGroupRolloutPolicyPolicyTypeCount,
	"COMPUTE_INSTANCE_GROUP_LINEAR_ROLLOUT_POLICY_BY_PERCENTAGE": ComputeInstanceGroupRolloutPolicyPolicyTypePercentage,
}

// GetComputeInstanceGroupRolloutPolicyPolicyTypeEnumValues Enumerates the set of values for ComputeInstanceGroupRolloutPolicyPolicyTypeEnum
func GetComputeInstanceGroupRolloutPolicyPolicyTypeEnumValues() []ComputeInstanceGroupRolloutPolicyPolicyTypeEnum {
	values := make([]ComputeInstanceGroupRolloutPolicyPolicyTypeEnum, 0)
	for _, v := range mappingComputeInstanceGroupRolloutPolicyPolicyType {
		values = append(values, v)
	}
	return values
}
