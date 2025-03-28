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

// AutoScalingConfiguration An autoscaling configuration lets you dynamically scale the resources in a Compute instance pool.
// For more information, see Autoscaling (https://docs.oracle.com/iaas/Content/Compute/Tasks/autoscalinginstancepools.htm).
type AutoScalingConfiguration struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the autoscaling configuration.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the autoscaling configuration.
	Id *string `mandatory:"true" json:"id"`

	Resource Resource `mandatory:"true" json:"resource"`

	// Autoscaling policy definitions for the autoscaling configuration. An autoscaling policy defines the criteria that
	// trigger autoscaling actions and the actions to take.
	Policies []AutoScalingPolicy `mandatory:"true" json:"policies"`

	// The date and time the autoscaling configuration was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// For threshold-based autoscaling policies, this value is the minimum period of time to wait between scaling actions.
	// The cooldown period gives the system time to stabilize before rescaling. The minimum value is 300 seconds, which
	// is also the default. The cooldown period starts when the instance pool reaches the running state.
	// For schedule-based autoscaling policies, this value is not used.
	CoolDownInSeconds *int `mandatory:"false" json:"coolDownInSeconds"`

	// Whether the autoscaling configuration is enabled.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// The maximum number of resources to scale out to.
	MaxResourceCount *int `mandatory:"false" json:"maxResourceCount"`

	// The minimum number of resources to scale in to.
	MinResourceCount *int `mandatory:"false" json:"minResourceCount"`
}

func (m AutoScalingConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutoScalingConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *AutoScalingConfiguration) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DefinedTags       map[string]map[string]interface{} `json:"definedTags"`
		DisplayName       *string                           `json:"displayName"`
		FreeformTags      map[string]string                 `json:"freeformTags"`
		CoolDownInSeconds *int                              `json:"coolDownInSeconds"`
		IsEnabled         *bool                             `json:"isEnabled"`
		MaxResourceCount  *int                              `json:"maxResourceCount"`
		MinResourceCount  *int                              `json:"minResourceCount"`
		CompartmentId     *string                           `json:"compartmentId"`
		Id                *string                           `json:"id"`
		Resource          resource                          `json:"resource"`
		Policies          []autoscalingpolicy               `json:"policies"`
		TimeCreated       *common.SDKTime                   `json:"timeCreated"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DefinedTags = model.DefinedTags

	m.DisplayName = model.DisplayName

	m.FreeformTags = model.FreeformTags

	m.CoolDownInSeconds = model.CoolDownInSeconds

	m.IsEnabled = model.IsEnabled

	m.MaxResourceCount = model.MaxResourceCount

	m.MinResourceCount = model.MinResourceCount

	m.CompartmentId = model.CompartmentId

	m.Id = model.Id

	nn, e = model.Resource.UnmarshalPolymorphicJSON(model.Resource.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Resource = nn.(Resource)
	} else {
		m.Resource = nil
	}

	m.Policies = make([]AutoScalingPolicy, len(model.Policies))
	for i, n := range model.Policies {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Policies[i] = nn.(AutoScalingPolicy)
		} else {
			m.Policies[i] = nil
		}
	}
	m.TimeCreated = model.TimeCreated

	return
}
