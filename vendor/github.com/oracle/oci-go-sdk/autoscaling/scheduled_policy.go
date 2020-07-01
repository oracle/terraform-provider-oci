// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Autoscaling API
//
// APIs for dynamically scaling Compute resources to meet application requirements. For more information about
// autoscaling, see Autoscaling (https://docs.cloud.oracle.com/Content/Compute/Tasks/autoscalinginstancepools.htm). For information about the
// Compute service, see Overview of the Compute Service (https://docs.cloud.oracle.com/Content/Compute/Concepts/computeoverview.htm).
// **Note:** Autoscaling is not available in US Government Cloud tenancies. For more information, see
// Oracle Cloud Infrastructure US Government Cloud (https://docs.cloud.oracle.com/Content/General/Concepts/govoverview.htm).
//

package autoscaling

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// ScheduledPolicy An autoscaling policy that defines execution schedules for an autoscaling configuration.
type ScheduledPolicy struct {

	// The capacity requirements of the autoscaling policy.
	Capacity *Capacity `mandatory:"true" json:"capacity"`

	// The date and time the autoscaling configuration was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	ExecutionSchedule ExecutionSchedule `mandatory:"true" json:"executionSchedule"`

	// The ID of the autoscaling policy that is assigned after creation.
	Id *string `mandatory:"false" json:"id"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Boolean field indicating whether this policy is enabled or not.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`
}

//GetCapacity returns Capacity
func (m ScheduledPolicy) GetCapacity() *Capacity {
	return m.Capacity
}

//GetId returns Id
func (m ScheduledPolicy) GetId() *string {
	return m.Id
}

//GetDisplayName returns DisplayName
func (m ScheduledPolicy) GetDisplayName() *string {
	return m.DisplayName
}

//GetTimeCreated returns TimeCreated
func (m ScheduledPolicy) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetIsEnabled returns IsEnabled
func (m ScheduledPolicy) GetIsEnabled() *bool {
	return m.IsEnabled
}

func (m ScheduledPolicy) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ScheduledPolicy) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeScheduledPolicy ScheduledPolicy
	s := struct {
		DiscriminatorParam string `json:"policyType"`
		MarshalTypeScheduledPolicy
	}{
		"scheduled",
		(MarshalTypeScheduledPolicy)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *ScheduledPolicy) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Id                *string           `json:"id"`
		DisplayName       *string           `json:"displayName"`
		IsEnabled         *bool             `json:"isEnabled"`
		Capacity          *Capacity         `json:"capacity"`
		TimeCreated       *common.SDKTime   `json:"timeCreated"`
		ExecutionSchedule executionschedule `json:"executionSchedule"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Id = model.Id

	m.DisplayName = model.DisplayName

	m.IsEnabled = model.IsEnabled

	m.Capacity = model.Capacity

	m.TimeCreated = model.TimeCreated

	nn, e = model.ExecutionSchedule.UnmarshalPolymorphicJSON(model.ExecutionSchedule.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ExecutionSchedule = nn.(ExecutionSchedule)
	} else {
		m.ExecutionSchedule = nil
	}

	return
}
