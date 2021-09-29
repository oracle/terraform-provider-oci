// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps APIs to create a DevOps project to group the pipelines,  add reference to target deployment environments, add artifacts to deploy,  and create deployment pipelines needed to deploy your software.
//

package devops

import (
	"github.com/oracle/oci-go-sdk/v48/common"
)

// LoadBalancerTrafficShiftRolloutPolicy Description of rollout policy for load balancer traffic shift stage.
type LoadBalancerTrafficShiftRolloutPolicy struct {

	// Specifies number of batches for this stage.
	BatchCount *int `mandatory:"true" json:"batchCount"`

	// Specifies delay in seconds between batches. The default delay is 1 minute.
	BatchDelayInSeconds *int `mandatory:"false" json:"batchDelayInSeconds"`

	// Indicates the criteria to stop.
	RampLimitPercent *float32 `mandatory:"false" json:"rampLimitPercent"`
}

func (m LoadBalancerTrafficShiftRolloutPolicy) String() string {
	return common.PointerString(m)
}
