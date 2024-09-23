// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management Service API. Use this API to for all FAMS related activities.
// To manage fleets,view complaince report for the Fleet,scedule patches and other lifecycle activities
//

package fleetappsmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// FleetCredentialSummary Summary of the FleetCredential.
type FleetCredentialSummary struct {

	// The unique id of the resource.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	// Example: `My new resource`
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The time this resource was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	EntitySpecifics CredentialEntitySpecificDetails `mandatory:"true" json:"entitySpecifics"`

	// The current state of the FleetCredential.
	LifecycleState FleetCredentialLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Tenancy OCID
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The time this resource was last updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	User CredentialDetails `mandatory:"false" json:"user"`

	Password CredentialDetails `mandatory:"false" json:"password"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m FleetCredentialSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FleetCredentialSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingFleetCredentialLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetFleetCredentialLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *FleetCredentialSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		CompartmentId    *string                           `json:"compartmentId"`
		TimeUpdated      *common.SDKTime                   `json:"timeUpdated"`
		User             credentialdetails                 `json:"user"`
		Password         credentialdetails                 `json:"password"`
		LifecycleDetails *string                           `json:"lifecycleDetails"`
		SystemTags       map[string]map[string]interface{} `json:"systemTags"`
		Id               *string                           `json:"id"`
		DisplayName      *string                           `json:"displayName"`
		TimeCreated      *common.SDKTime                   `json:"timeCreated"`
		EntitySpecifics  credentialentityspecificdetails   `json:"entitySpecifics"`
		LifecycleState   FleetCredentialLifecycleStateEnum `json:"lifecycleState"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.CompartmentId = model.CompartmentId

	m.TimeUpdated = model.TimeUpdated

	nn, e = model.User.UnmarshalPolymorphicJSON(model.User.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.User = nn.(CredentialDetails)
	} else {
		m.User = nil
	}

	nn, e = model.Password.UnmarshalPolymorphicJSON(model.Password.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Password = nn.(CredentialDetails)
	} else {
		m.Password = nil
	}

	m.LifecycleDetails = model.LifecycleDetails

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.DisplayName = model.DisplayName

	m.TimeCreated = model.TimeCreated

	nn, e = model.EntitySpecifics.UnmarshalPolymorphicJSON(model.EntitySpecifics.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.EntitySpecifics = nn.(CredentialEntitySpecificDetails)
	} else {
		m.EntitySpecifics = nil
	}

	m.LifecycleState = model.LifecycleState

	return
}
