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

// CreateThresholdPolicyDetails Creation details for a threshold-based autoscaling policy.
// In a threshold-based autoscaling policy, an autoscaling action is triggered when a performance metric meets
// or exceeds a threshold.
type CreateThresholdPolicyDetails struct {

	// The capacity requirements of the autoscaling policy.
	Capacity *Capacity `mandatory:"true" json:"capacity"`

	Rules []CreateConditionDetails `mandatory:"true" json:"rules"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Boolean field indicating whether this policy is enabled or not.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`
}

//GetCapacity returns Capacity
func (m CreateThresholdPolicyDetails) GetCapacity() *Capacity {
	return m.Capacity
}

//GetDisplayName returns DisplayName
func (m CreateThresholdPolicyDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetIsEnabled returns IsEnabled
func (m CreateThresholdPolicyDetails) GetIsEnabled() *bool {
	return m.IsEnabled
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
