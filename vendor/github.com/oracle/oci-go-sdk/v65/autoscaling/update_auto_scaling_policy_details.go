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

// UpdateAutoScalingPolicyDetails The representation of UpdateAutoScalingPolicyDetails
type UpdateAutoScalingPolicyDetails interface {

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	GetDisplayName() *string

	// The capacity requirements of the autoscaling policy.
	GetCapacity() *Capacity

	// Whether the autoscaling policy is enabled.
	GetIsEnabled() *bool
}

type updateautoscalingpolicydetails struct {
	JsonData    []byte
	DisplayName *string   `mandatory:"false" json:"displayName"`
	Capacity    *Capacity `mandatory:"false" json:"capacity"`
	IsEnabled   *bool     `mandatory:"false" json:"isEnabled"`
	PolicyType  string    `json:"policyType"`
}

// UnmarshalJSON unmarshals json
func (m *updateautoscalingpolicydetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdateautoscalingpolicydetails updateautoscalingpolicydetails
	s := struct {
		Model Unmarshalerupdateautoscalingpolicydetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DisplayName = s.Model.DisplayName
	m.Capacity = s.Model.Capacity
	m.IsEnabled = s.Model.IsEnabled
	m.PolicyType = s.Model.PolicyType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updateautoscalingpolicydetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.PolicyType {
	case "threshold":
		mm := UpdateThresholdPolicyDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "scheduled":
		mm := UpdateScheduledPolicyDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for UpdateAutoScalingPolicyDetails: %s.", m.PolicyType)
		return *m, nil
	}
}

// GetDisplayName returns DisplayName
func (m updateautoscalingpolicydetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetCapacity returns Capacity
func (m updateautoscalingpolicydetails) GetCapacity() *Capacity {
	return m.Capacity
}

// GetIsEnabled returns IsEnabled
func (m updateautoscalingpolicydetails) GetIsEnabled() *bool {
	return m.IsEnabled
}

func (m updateautoscalingpolicydetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updateautoscalingpolicydetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
