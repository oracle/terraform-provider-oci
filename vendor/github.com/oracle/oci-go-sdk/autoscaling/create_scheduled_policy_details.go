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

// CreateScheduledPolicyDetails Creation details for a schedule-based autoscaling policy.
// In a schedule-based autoscaling policy, an autoscaling action is triggered when execution time is current.
type CreateScheduledPolicyDetails struct {

	// The capacity requirements of the autoscaling policy.
	Capacity *Capacity `mandatory:"true" json:"capacity"`

	ExecutionSchedule ExecutionSchedule `mandatory:"true" json:"executionSchedule"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Boolean field indicating whether this policy is enabled or not.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`
}

//GetCapacity returns Capacity
func (m CreateScheduledPolicyDetails) GetCapacity() *Capacity {
	return m.Capacity
}

//GetDisplayName returns DisplayName
func (m CreateScheduledPolicyDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetIsEnabled returns IsEnabled
func (m CreateScheduledPolicyDetails) GetIsEnabled() *bool {
	return m.IsEnabled
}

func (m CreateScheduledPolicyDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateScheduledPolicyDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateScheduledPolicyDetails CreateScheduledPolicyDetails
	s := struct {
		DiscriminatorParam string `json:"policyType"`
		MarshalTypeCreateScheduledPolicyDetails
	}{
		"scheduled",
		(MarshalTypeCreateScheduledPolicyDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CreateScheduledPolicyDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName       *string           `json:"displayName"`
		IsEnabled         *bool             `json:"isEnabled"`
		Capacity          *Capacity         `json:"capacity"`
		ExecutionSchedule executionschedule `json:"executionSchedule"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.IsEnabled = model.IsEnabled

	m.Capacity = model.Capacity

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
