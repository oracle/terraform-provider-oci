// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Auto Scaling API
//
// Auto Scaling API spec
//

package autoscaling

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// CreateThresholdPolicyDetails An AutoScalingConfiguration ThresholdPolicy creation details
type CreateThresholdPolicyDetails struct {

	// The capacity requirements of the Policy
	Capacity *Capacity `mandatory:"true" json:"capacity"`

	Rules []CreateConditionDetails `mandatory:"true" json:"rules"`

	// A user-friendly name for the Policy. Does not have to be unique, and it's changeable. Avoid entering
	// confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`
}

//GetCapacity returns Capacity
func (m CreateThresholdPolicyDetails) GetCapacity() *Capacity {
	return m.Capacity
}

//GetDisplayName returns DisplayName
func (m CreateThresholdPolicyDetails) GetDisplayName() *string {
	return m.DisplayName
}

func (m CreateThresholdPolicyDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateThresholdPolicyDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateThresholdPolicyDetails CreateThresholdPolicyDetails
	s := struct {
		DiscriminatorParam string `json:"policyType"`
		MarshalTypeCreateThresholdPolicyDetails
	}{
		"threshold",
		(MarshalTypeCreateThresholdPolicyDetails)(m),
	}

	return json.Marshal(&s)
}
