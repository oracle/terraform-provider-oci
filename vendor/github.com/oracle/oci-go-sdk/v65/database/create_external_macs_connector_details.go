// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateExternalMacsConnectorDetails Details for creating a resource used to connect to an external Oracle Database using
// the Management Agent cloud service (MACS) (https://docs.cloud.oracle.com/iaas/management-agents/index.html).
type CreateExternalMacsConnectorDetails struct {

	// The user-friendly name for the
	// CreateExternalDatabaseConnectorDetails.
	// The name does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the external database resource.
	ExternalDatabaseId *string `mandatory:"true" json:"externalDatabaseId"`

	ConnectionString *DatabaseConnectionString `mandatory:"true" json:"connectionString"`

	ConnectionCredentials DatabaseConnectionCredentials `mandatory:"true" json:"connectionCredentials"`

	// The ID of the agent used for the
	// CreateExternalDatabaseConnectorDetails.
	ConnectorAgentId *string `mandatory:"true" json:"connectorAgentId"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

//GetFreeformTags returns FreeformTags
func (m CreateExternalMacsConnectorDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m CreateExternalMacsConnectorDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetDisplayName returns DisplayName
func (m CreateExternalMacsConnectorDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetExternalDatabaseId returns ExternalDatabaseId
func (m CreateExternalMacsConnectorDetails) GetExternalDatabaseId() *string {
	return m.ExternalDatabaseId
}

func (m CreateExternalMacsConnectorDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateExternalMacsConnectorDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateExternalMacsConnectorDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateExternalMacsConnectorDetails CreateExternalMacsConnectorDetails
	s := struct {
		DiscriminatorParam string `json:"connectorType"`
		MarshalTypeCreateExternalMacsConnectorDetails
	}{
		"MACS",
		(MarshalTypeCreateExternalMacsConnectorDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CreateExternalMacsConnectorDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		FreeformTags          map[string]string                 `json:"freeformTags"`
		DefinedTags           map[string]map[string]interface{} `json:"definedTags"`
		DisplayName           *string                           `json:"displayName"`
		ExternalDatabaseId    *string                           `json:"externalDatabaseId"`
		ConnectionString      *DatabaseConnectionString         `json:"connectionString"`
		ConnectionCredentials databaseconnectioncredentials     `json:"connectionCredentials"`
		ConnectorAgentId      *string                           `json:"connectorAgentId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.DisplayName = model.DisplayName

	m.ExternalDatabaseId = model.ExternalDatabaseId

	m.ConnectionString = model.ConnectionString

	nn, e = model.ConnectionCredentials.UnmarshalPolymorphicJSON(model.ConnectionCredentials.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ConnectionCredentials = nn.(DatabaseConnectionCredentials)
	} else {
		m.ConnectionCredentials = nil
	}

	m.ConnectorAgentId = model.ConnectorAgentId

	return
}
