// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.oracle.com/iaas/Content/devops/using/home.htm).
//

package devops

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ComputeInstanceGroupLinearRolloutPolicyByCount Specifies a linear rollout strategy for a compute instance group rolling deployment stage.
type ComputeInstanceGroupLinearRolloutPolicyByCount struct {

	// The number that will be used to determine how many instances will be deployed concurrently.
	BatchCount *int `mandatory:"true" json:"batchCount"`

	// The duration of delay between batch rollout. The default delay is 1 minute.
	BatchDelayInSeconds *int `mandatory:"false" json:"batchDelayInSeconds"`
}

// GetBatchDelayInSeconds returns BatchDelayInSeconds
func (m ComputeInstanceGroupLinearRolloutPolicyByCount) GetBatchDelayInSeconds() *int {
	return m.BatchDelayInSeconds
}

func (m ComputeInstanceGroupLinearRolloutPolicyByCount) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ComputeInstanceGroupLinearRolloutPolicyByCount) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ComputeInstanceGroupLinearRolloutPolicyByCount) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeComputeInstanceGroupLinearRolloutPolicyByCount ComputeInstanceGroupLinearRolloutPolicyByCount
	s := struct {
		DiscriminatorParam string `json:"policyType"`
		MarshalTypeComputeInstanceGroupLinearRolloutPolicyByCount
	}{
		"COMPUTE_INSTANCE_GROUP_LINEAR_ROLLOUT_POLICY_BY_COUNT",
		(MarshalTypeComputeInstanceGroupLinearRolloutPolicyByCount)(m),
	}

	return json.Marshal(&s)
}
