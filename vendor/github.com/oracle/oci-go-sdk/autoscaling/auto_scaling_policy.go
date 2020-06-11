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

// AutoScalingPolicy Autoscaling policies define the criteria that trigger autoscaling actions and the actions to take.
// An autoscaling policy is part of an autoscaling configuration. For more information, see
// Autoscaling (https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/autoscalinginstancepools.htm).
type AutoScalingPolicy interface {

	// The capacity requirements of the autoscaling policy.
	GetCapacity() *Capacity

	// The date and time the autoscaling configuration was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	GetTimeCreated() *common.SDKTime

	// The ID of the autoscaling policy that is assigned after creation.
	GetId() *string

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	GetDisplayName() *string

	// Boolean field indicating whether this policy is enabled or not.
	GetIsEnabled() *bool
}

type autoscalingpolicy struct {
	JsonData    []byte
	Capacity    *Capacity       `mandatory:"true" json:"capacity"`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`
	Id          *string         `mandatory:"false" json:"id"`
	DisplayName *string         `mandatory:"false" json:"displayName"`
	IsEnabled   *bool           `mandatory:"false" json:"isEnabled"`
	PolicyType  string          `json:"policyType"`
}

// UnmarshalJSON unmarshals json
func (m *autoscalingpolicy) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerautoscalingpolicy autoscalingpolicy
	s := struct {
		Model Unmarshalerautoscalingpolicy
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Capacity = s.Model.Capacity
	m.TimeCreated = s.Model.TimeCreated
	m.Id = s.Model.Id
	m.DisplayName = s.Model.DisplayName
	m.IsEnabled = s.Model.IsEnabled
	m.PolicyType = s.Model.PolicyType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *autoscalingpolicy) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.PolicyType {
	case "scheduled":
		mm := ScheduledPolicy{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "threshold":
		mm := ThresholdPolicy{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetCapacity returns Capacity
func (m autoscalingpolicy) GetCapacity() *Capacity {
	return m.Capacity
}

//GetTimeCreated returns TimeCreated
func (m autoscalingpolicy) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetId returns Id
func (m autoscalingpolicy) GetId() *string {
	return m.Id
}

//GetDisplayName returns DisplayName
func (m autoscalingpolicy) GetDisplayName() *string {
	return m.DisplayName
}

//GetIsEnabled returns IsEnabled
func (m autoscalingpolicy) GetIsEnabled() *bool {
	return m.IsEnabled
}

func (m autoscalingpolicy) String() string {
	return common.PointerString(m)
}
