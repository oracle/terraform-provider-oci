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

// AutoScalingConfigurationSummary The representation of AutoScalingConfigurationSummary
type AutoScalingConfigurationSummary struct {

	// The OCID of the compartment containing the AutoScalingConfiguration.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the AutoScalingConfiguration
	Id *string `mandatory:"true" json:"id"`

	// The date and time the AutoScalingConfiguration was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// A user-friendly name for the AutoScalingConfiguration. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The minimum period of time between scaling actions. The default is 300 seconds.
	CoolDownInSeconds *int `mandatory:"false" json:"coolDownInSeconds"`

	// If the AutoScalingConfiguration is enabled
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	Resource Resource `mandatory:"false" json:"resource"`
}

func (m AutoScalingConfigurationSummary) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *AutoScalingConfigurationSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName       *string         `json:"displayName"`
		CoolDownInSeconds *int            `json:"coolDownInSeconds"`
		IsEnabled         *bool           `json:"isEnabled"`
		Resource          resource        `json:"resource"`
		CompartmentId     *string         `json:"compartmentId"`
		Id                *string         `json:"id"`
		TimeCreated       *common.SDKTime `json:"timeCreated"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	m.DisplayName = model.DisplayName
	m.CoolDownInSeconds = model.CoolDownInSeconds
	m.IsEnabled = model.IsEnabled
	nn, e := model.Resource.UnmarshalPolymorphicJSON(model.Resource.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Resource = nn.(Resource)
	} else {
		m.Resource = nil
	}
	m.CompartmentId = model.CompartmentId
	m.Id = model.Id
	m.TimeCreated = model.TimeCreated
	return
}
