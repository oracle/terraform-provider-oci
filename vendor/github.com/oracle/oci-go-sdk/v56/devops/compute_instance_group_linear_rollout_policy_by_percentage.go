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

// ComputeInstanceGroupLinearRolloutPolicyByPercentage Specifies a linear rollout strategy for a compute instance group rolling deployment stage.
type ComputeInstanceGroupLinearRolloutPolicyByPercentage struct {

	// The percentage that will be used to determine how many instances will be deployed concurrently.
	BatchPercentage *int `mandatory:"true" json:"batchPercentage"`

	// The duration of delay between batch rollout. The default delay is 1 minute.
	BatchDelayInSeconds *int `mandatory:"false" json:"batchDelayInSeconds"`
}

//GetBatchDelayInSeconds returns BatchDelayInSeconds
func (m ComputeInstanceGroupLinearRolloutPolicyByPercentage) GetBatchDelayInSeconds() *int {
	return m.BatchDelayInSeconds
}

func (m ComputeInstanceGroupLinearRolloutPolicyByPercentage) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ComputeInstanceGroupLinearRolloutPolicyByPercentage) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeComputeInstanceGroupLinearRolloutPolicyByPercentage ComputeInstanceGroupLinearRolloutPolicyByPercentage
	s := struct {
		DiscriminatorParam string `json:"policyType"`
		MarshalTypeComputeInstanceGroupLinearRolloutPolicyByPercentage
	}{
		"COMPUTE_INSTANCE_GROUP_LINEAR_ROLLOUT_POLICY_BY_PERCENTAGE",
		(MarshalTypeComputeInstanceGroupLinearRolloutPolicyByPercentage)(m),
	}

	return json.Marshal(&s)
}
