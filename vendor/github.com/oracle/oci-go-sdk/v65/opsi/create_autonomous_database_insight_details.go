// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateAutonomousDatabaseInsightDetails The information about database to be analyzed. When isAdvancedFeaturesEnabled is set to false, parameters connectionDetails, credentialDetails and opsiPrivateEndpoint are optional. Otherwise, connectionDetails and crendetialDetails are required to enable full OPSI service features. If the Autonomouse Database is configured with private, restricted or dedicated access, opsiPrivateEndpoint parameter is required.
type CreateAutonomousDatabaseInsightDetails struct {

	// Compartment Identifier of database
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database.
	DatabaseId *string `mandatory:"true" json:"databaseId"`

	// OCI database resource type
	DatabaseResourceType *string `mandatory:"true" json:"databaseResourceType"`

	// Flag is to identify if advanced features for autonomous database is enabled or not
	IsAdvancedFeaturesEnabled *bool `mandatory:"true" json:"isAdvancedFeaturesEnabled"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	ConnectionDetails *ConnectionDetails `mandatory:"false" json:"connectionDetails"`

	CredentialDetails CredentialDetails `mandatory:"false" json:"credentialDetails"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the OPSI private endpoint
	OpsiPrivateEndpointId *string `mandatory:"false" json:"opsiPrivateEndpointId"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

// GetCompartmentId returns CompartmentId
func (m CreateAutonomousDatabaseInsightDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetFreeformTags returns FreeformTags
func (m CreateAutonomousDatabaseInsightDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateAutonomousDatabaseInsightDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateAutonomousDatabaseInsightDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateAutonomousDatabaseInsightDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateAutonomousDatabaseInsightDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateAutonomousDatabaseInsightDetails CreateAutonomousDatabaseInsightDetails
	s := struct {
		DiscriminatorParam string `json:"entitySource"`
		MarshalTypeCreateAutonomousDatabaseInsightDetails
	}{
		"AUTONOMOUS_DATABASE",
		(MarshalTypeCreateAutonomousDatabaseInsightDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CreateAutonomousDatabaseInsightDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		FreeformTags              map[string]string                 `json:"freeformTags"`
		DefinedTags               map[string]map[string]interface{} `json:"definedTags"`
		ConnectionDetails         *ConnectionDetails                `json:"connectionDetails"`
		CredentialDetails         credentialdetails                 `json:"credentialDetails"`
		OpsiPrivateEndpointId     *string                           `json:"opsiPrivateEndpointId"`
		SystemTags                map[string]map[string]interface{} `json:"systemTags"`
		CompartmentId             *string                           `json:"compartmentId"`
		DatabaseId                *string                           `json:"databaseId"`
		DatabaseResourceType      *string                           `json:"databaseResourceType"`
		IsAdvancedFeaturesEnabled *bool                             `json:"isAdvancedFeaturesEnabled"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.ConnectionDetails = model.ConnectionDetails

	nn, e = model.CredentialDetails.UnmarshalPolymorphicJSON(model.CredentialDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.CredentialDetails = nn.(CredentialDetails)
	} else {
		m.CredentialDetails = nil
	}

	m.OpsiPrivateEndpointId = model.OpsiPrivateEndpointId

	m.SystemTags = model.SystemTags

	m.CompartmentId = model.CompartmentId

	m.DatabaseId = model.DatabaseId

	m.DatabaseResourceType = model.DatabaseResourceType

	m.IsAdvancedFeaturesEnabled = model.IsAdvancedFeaturesEnabled

	return
}
