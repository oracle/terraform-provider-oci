// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Autoscaling API
//
// APIs for dynamically scaling Compute resources to meet application requirements. For more information about
// autoscaling, see Autoscaling (https://docs.oracle.com/iaas/Content/Compute/Tasks/autoscalinginstancepools.htm). For information about the
// Compute service, see Overview of the Compute Service (https://docs.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm).
// **Note:** Autoscaling is not available in US Government Cloud tenancies. For more information, see
// Oracle Cloud Infrastructure US Government Cloud (https://docs.oracle.com/iaas/Content/General/Concepts/govoverview.htm).
//

package autoscaling

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ThresholdPolicy An autoscaling policy that defines threshold-based rules for an autoscaling configuration.
type ThresholdPolicy struct {

	// The date and time the autoscaling configuration was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	Rules []Condition `mandatory:"true" json:"rules"`

	// The capacity requirements of the autoscaling policy.
	Capacity *Capacity `mandatory:"false" json:"capacity"`

	// The ID of the autoscaling policy that is assigned after creation.
	Id *string `mandatory:"false" json:"id"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Whether the autoscaling policy is enabled.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`
}

// GetCapacity returns Capacity
func (m ThresholdPolicy) GetCapacity() *Capacity {
	return m.Capacity
}

// GetId returns Id
func (m ThresholdPolicy) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m ThresholdPolicy) GetDisplayName() *string {
	return m.DisplayName
}

// GetTimeCreated returns TimeCreated
func (m ThresholdPolicy) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetIsEnabled returns IsEnabled
func (m ThresholdPolicy) GetIsEnabled() *bool {
	return m.IsEnabled
}

func (m ThresholdPolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ThresholdPolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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
