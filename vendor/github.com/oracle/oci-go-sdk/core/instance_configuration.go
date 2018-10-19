// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// InstanceConfiguration Instance Configuration
type InstanceConfiguration struct {

	// The OCID of the compartment containing the instance configuration.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the instance configuration
	Id *string `mandatory:"true" json:"id"`

	// The date and time the instance configuration was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name for the instance configuration
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see
	// Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	InstanceDetails InstanceConfigurationInstanceDetails `mandatory:"false" json:"instanceDetails"`

	DeferredFields []string `mandatory:"false" json:"deferredFields"`
}

func (m InstanceConfiguration) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *InstanceConfiguration) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DefinedTags     map[string]map[string]interface{}    `json:"definedTags"`
		DisplayName     *string                              `json:"displayName"`
		FreeformTags    map[string]string                    `json:"freeformTags"`
		InstanceDetails instanceconfigurationinstancedetails `json:"instanceDetails"`
		DeferredFields  []string                             `json:"deferredFields"`
		CompartmentId   *string                              `json:"compartmentId"`
		Id              *string                              `json:"id"`
		TimeCreated     *common.SDKTime                      `json:"timeCreated"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	m.DefinedTags = model.DefinedTags
	m.DisplayName = model.DisplayName
	m.FreeformTags = model.FreeformTags
	nn, e := model.InstanceDetails.UnmarshalPolymorphicJSON(model.InstanceDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.InstanceDetails = nn.(InstanceConfigurationInstanceDetails)
	} else {
		m.InstanceDetails = nil
	}
	m.DeferredFields = make([]string, len(model.DeferredFields))
	for i, n := range model.DeferredFields {
		m.DeferredFields[i] = n
	}
	m.CompartmentId = model.CompartmentId
	m.Id = model.Id
	m.TimeCreated = model.TimeCreated
	return
}
