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

// ThresholdPolicy A Policy that defines threshold based rules for an AutoScalingConfiguration
type ThresholdPolicy struct {

	// The capacity requirements of the Policy
	Capacity *Capacity `mandatory:"true" json:"capacity"`

	// The date and time the AutoScalingConfiguration was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	Rules []Condition `mandatory:"true" json:"rules"`

	// The ID of the policy that is assigned after creation
	Id *string `mandatory:"false" json:"id"`

	// A user-friendly name for the Policy. Does not have to be unique, and it's changeable. Avoid entering
	// confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`
}

//GetCapacity returns Capacity
func (m ThresholdPolicy) GetCapacity() *Capacity {
	return m.Capacity
}

//GetId returns Id
func (m ThresholdPolicy) GetId() *string {
	return m.Id
}

//GetDisplayName returns DisplayName
func (m ThresholdPolicy) GetDisplayName() *string {
	return m.DisplayName
}

//GetTimeCreated returns TimeCreated
func (m ThresholdPolicy) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

func (m ThresholdPolicy) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ThresholdPolicy) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeThresholdPolicy ThresholdPolicy
	s := struct {
		DiscriminatorParam string `json:"policyType"`
		MarshalTypeThresholdPolicy
	}{
		"threshold",
		(MarshalTypeThresholdPolicy)(m),
	}

	return json.Marshal(&s)
}
