// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Visual Builder API
//
// Oracle Visual Builder enables developers to quickly build web and mobile applications. With a visual development environment that makes it easy to connect to Oracle data and third-party REST services, developers can build modern, consumer-grade applications in a fraction of the time it would take in other tools.
// The Visual Builder Instance Management API allows users to create and manage a Visual Builder instance.
//

package visualbuilder

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateVbInstanceDetails Information about updating a VbInstance.
type UpdateVbInstanceDetails struct {

	// Vb Instance Identifier.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name,
	// type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to
	// namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Encrypted IDCS Open ID token. This is required for pre-UCPIS cloud accounts, but not UCPIS, hence not a required parameter
	IdcsOpenId *string `mandatory:"false" json:"idcsOpenId"`

	// The number of Nodes
	NodeCount *int `mandatory:"false" json:"nodeCount"`

	// Enable Visual Builder. If Visual Builder is enabled alredy, then it cannot be disabled.
	IsVisualBuilderEnabled *bool `mandatory:"false" json:"isVisualBuilderEnabled"`

	CustomEndpoint *UpdateCustomEndpointDetails `mandatory:"false" json:"customEndpoint"`

	// A list of alternate custom endpoints to be used for the vb instance URL
	// (contact Oracle for alternateCustomEndpoints availability for a specific instance).
	AlternateCustomEndpoints []UpdateCustomEndpointDetails `mandatory:"false" json:"alternateCustomEndpoints"`

	NetworkEndpointDetails UpdateNetworkEndpointDetails `mandatory:"false" json:"networkEndpointDetails"`
}

func (m UpdateVbInstanceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateVbInstanceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *UpdateVbInstanceDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName              *string                           `json:"displayName"`
		FreeformTags             map[string]string                 `json:"freeformTags"`
		DefinedTags              map[string]map[string]interface{} `json:"definedTags"`
		IdcsOpenId               *string                           `json:"idcsOpenId"`
		NodeCount                *int                              `json:"nodeCount"`
		IsVisualBuilderEnabled   *bool                             `json:"isVisualBuilderEnabled"`
		CustomEndpoint           *UpdateCustomEndpointDetails      `json:"customEndpoint"`
		AlternateCustomEndpoints []UpdateCustomEndpointDetails     `json:"alternateCustomEndpoints"`
		NetworkEndpointDetails   updatenetworkendpointdetails      `json:"networkEndpointDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.IdcsOpenId = model.IdcsOpenId

	m.NodeCount = model.NodeCount

	m.IsVisualBuilderEnabled = model.IsVisualBuilderEnabled

	m.CustomEndpoint = model.CustomEndpoint

	m.AlternateCustomEndpoints = make([]UpdateCustomEndpointDetails, len(model.AlternateCustomEndpoints))
	copy(m.AlternateCustomEndpoints, model.AlternateCustomEndpoints)
	nn, e = model.NetworkEndpointDetails.UnmarshalPolymorphicJSON(model.NetworkEndpointDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.NetworkEndpointDetails = nn.(UpdateNetworkEndpointDetails)
	} else {
		m.NetworkEndpointDetails = nil
	}

	return
}
