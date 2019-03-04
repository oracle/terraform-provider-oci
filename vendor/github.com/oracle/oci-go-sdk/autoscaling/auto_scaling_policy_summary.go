// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Auto Scaling API
//
// Auto Scaling API spec
//

package autoscaling

import (
	"github.com/oracle/oci-go-sdk/common"
)

// AutoScalingPolicySummary The representation of AutoScalingPolicySummary
type AutoScalingPolicySummary struct {

	// The ID of the policy that is assigned after creation
	Id *string `mandatory:"true" json:"id"`

	// Indicates type of Policy
	PolicyType *string `mandatory:"true" json:"policyType"`

	// A user-friendly name for the Policy. Does not have to be unique, and it's changeable. Avoid entering
	// confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`
}

func (m AutoScalingPolicySummary) String() string {
	return common.PointerString(m)
}
