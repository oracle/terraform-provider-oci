// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// UpdateScheduledPolicyDetails The representation of UpdateScheduledPolicyDetails
type UpdateScheduledPolicyDetails struct {

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The capacity requirements of the autoscaling policy.
	Capacity *Capacity `mandatory:"false" json:"capacity"`

	// Whether the autoscaling policy is enabled.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// The schedule for executing the autoscaling policy.
	ExecutionSchedule ExecutionSchedule `mandatory:"false" json:"executionSchedule"`

	ResourceAction ResourceAction `mandatory:"false" json:"resourceAction"`
}

//GetDisplayName returns DisplayName
func (m UpdateScheduledPolicyDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetCapacity returns Capacity
func (m UpdateScheduledPolicyDetails) GetCapacity() *Capacity {
	return m.Capacity
}

//GetIsEnabled returns IsEnabled
func (m UpdateScheduledPolicyDetails) GetIsEnabled() *bool {
	return m.IsEnabled
}

func (m UpdateScheduledPolicyDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateScheduledPolicyDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateScheduledPolicyDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateScheduledPolicyDetails UpdateScheduledPolicyDetails
	s := struct {
		DiscriminatorParam string `json:"policyType"`
		MarshalTypeUpdateScheduledPolicyDetails
	}{
		"scheduled",
		(MarshalTypeUpdateScheduledPolicyDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *UpdateScheduledPolicyDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName       *string           `json:"displayName"`
		Capacity          *Capacity         `json:"capacity"`
		IsEnabled         *bool             `json:"isEnabled"`
		ExecutionSchedule executionschedule `json:"executionSchedule"`
		ResourceAction    resourceaction    `json:"resourceAction"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Capacity = model.Capacity

	m.IsEnabled = model.IsEnabled

	nn, e = model.ExecutionSchedule.UnmarshalPolymorphicJSON(model.ExecutionSchedule.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ExecutionSchedule = nn.(ExecutionSchedule)
	} else {
		m.ExecutionSchedule = nil
	}

	nn, e = model.ResourceAction.UnmarshalPolymorphicJSON(model.ResourceAction.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ResourceAction = nn.(ResourceAction)
	} else {
		m.ResourceAction = nil
	}

	return
}
