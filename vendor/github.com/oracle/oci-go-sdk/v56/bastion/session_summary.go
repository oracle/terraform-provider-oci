// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Bastion API
//
// Oracle Cloud Infrastructure Bastion provides restricted and time-limited access to target resources that don't have public endpoints. Through the configuration of a bastion, you can let authorized users connect from specific IP addresses to target resources by way of Secure Shell (SSH) sessions hosted on the bastion.
//

package bastion

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// SessionSummary Summary information for a bastion session resource.
type SessionSummary struct {

	// The unique identifier (OCID) of the session, which can't be changed after creation.
	Id *string `mandatory:"true" json:"id"`

	// The name of the bastion that is hosting this session.
	BastionName *string `mandatory:"true" json:"bastionName"`

	// The unique identifier (OCID) of the bastion that is hosting this session.
	BastionId *string `mandatory:"true" json:"bastionId"`

	TargetResourceDetails TargetResourceDetails `mandatory:"true" json:"targetResourceDetails"`

	// The time the session was created. Format is defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2020-01-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the session.
	LifecycleState SessionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The amount of time the session can remain active.
	SessionTtlInSeconds *int `mandatory:"true" json:"sessionTtlInSeconds"`

	// The name of the session.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The time the session was updated. Format is defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2020-01-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current session state in more detail.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`
}

func (m SessionSummary) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *SessionSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName           *string                   `json:"displayName"`
		TimeUpdated           *common.SDKTime           `json:"timeUpdated"`
		LifecycleDetails      *string                   `json:"lifecycleDetails"`
		Id                    *string                   `json:"id"`
		BastionName           *string                   `json:"bastionName"`
		BastionId             *string                   `json:"bastionId"`
		TargetResourceDetails targetresourcedetails     `json:"targetResourceDetails"`
		TimeCreated           *common.SDKTime           `json:"timeCreated"`
		LifecycleState        SessionLifecycleStateEnum `json:"lifecycleState"`
		SessionTtlInSeconds   *int                      `json:"sessionTtlInSeconds"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleDetails = model.LifecycleDetails

	m.Id = model.Id

	m.BastionName = model.BastionName

	m.BastionId = model.BastionId

	nn, e = model.TargetResourceDetails.UnmarshalPolymorphicJSON(model.TargetResourceDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.TargetResourceDetails = nn.(TargetResourceDetails)
	} else {
		m.TargetResourceDetails = nil
	}

	m.TimeCreated = model.TimeCreated

	m.LifecycleState = model.LifecycleState

	m.SessionTtlInSeconds = model.SessionTtlInSeconds

	return
}
