// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps APIs to create a DevOps project to group the pipelines,  add reference to target deployment environments, add artifacts to deploy,  and create deployment pipelines needed to deploy your software.
//

package devops

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v44/common"
)

// ComputeInstanceGroupFailurePolicyByPercentage Specifies a failure policy by percentage for a compute instance group rolling deployment stage.
type ComputeInstanceGroupFailurePolicyByPercentage struct {

	// The failure percentage threshold, which when reached or exceeded sets the stage as FAILED. Percentage is computed as the ceiling value of the number of failed instances over the total count of the instances in the group.
	FailurePercentage *int `mandatory:"true" json:"failurePercentage"`
}

func (m ComputeInstanceGroupFailurePolicyByPercentage) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ComputeInstanceGroupFailurePolicyByPercentage) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeComputeInstanceGroupFailurePolicyByPercentage ComputeInstanceGroupFailurePolicyByPercentage
	s := struct {
		DiscriminatorParam string `json:"policyType"`
		MarshalTypeComputeInstanceGroupFailurePolicyByPercentage
	}{
		"COMPUTE_INSTANCE_GROUP_FAILURE_POLICY_BY_PERCENTAGE",
		(MarshalTypeComputeInstanceGroupFailurePolicyByPercentage)(m),
	}

	return json.Marshal(&s)
}
