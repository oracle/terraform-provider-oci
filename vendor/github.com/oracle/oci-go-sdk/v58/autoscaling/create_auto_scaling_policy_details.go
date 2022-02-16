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

// CreateAutoScalingPolicyDetails Creation details for an autoscaling policy. You can create the following types of autoscaling policies:
// - **Schedule-based:** Autoscaling events take place at the specific times that you schedule.
// - **Threshold-based:** An autoscaling action is triggered when a performance metric meets
// or exceeds a threshold.
// An autoscaling configuration can either have multiple schedule-based autoscaling policies, or one
// threshold-based autoscaling policy.
type CreateAutoScalingPolicyDetails interface {

	// The capacity requirements of the autoscaling policy.
	GetCapacity() *Capacity

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	GetDisplayName() *string

	// Whether the autoscaling policy is enabled.
	GetIsEnabled() *bool
}

type createautoscalingpolicydetails struct {
	JsonData    []byte
	Capacity    *Capacity `mandatory:"false" json:"capacity"`
	DisplayName *string   `mandatory:"false" json:"displayName"`
	IsEnabled   *bool     `mandatory:"false" json:"isEnabled"`
	PolicyType  string    `json:"policyType"`
}

// UnmarshalJSON unmarshals json
func (m *createautoscalingpolicydetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreateautoscalingpolicydetails createautoscalingpolicydetails
	s := struct {
		Model Unmarshalercreateautoscalingpolicydetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Capacity = s.Model.Capacity
	m.DisplayName = s.Model.DisplayName
	m.IsEnabled = s.Model.IsEnabled
	m.PolicyType = s.Model.PolicyType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createautoscalingpolicydetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.PolicyType {
	case "scheduled":
		mm := CreateScheduledPolicyDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "threshold":
		mm := CreateThresholdPolicyDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetCapacity returns Capacity
func (m createautoscalingpolicydetails) GetCapacity() *Capacity {
	return m.Capacity
}

//GetDisplayName returns DisplayName
func (m createautoscalingpolicydetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetIsEnabled returns IsEnabled
func (m createautoscalingpolicydetails) GetIsEnabled() *bool {
	return m.IsEnabled
}

func (m createautoscalingpolicydetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createautoscalingpolicydetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
