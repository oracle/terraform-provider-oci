// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateExternalDbSystemMacsConnectorDetails The details for creating an external connector that is used to connect to an external DB system component
// using the Management Agent Cloud Service (MACS) (https://docs.cloud.oracle.com/iaas/management-agents/index.html).
type CreateExternalDbSystemMacsConnectorDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the external DB system.
	ExternalDbSystemId *string `mandatory:"true" json:"externalDbSystemId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the management agent
	// used for the external DB system connector.
	AgentId *string `mandatory:"true" json:"agentId"`

	// The user-friendly name for the external connector. The name does not have to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	ConnectionInfo ExternalDbSystemConnectionInfo `mandatory:"false" json:"connectionInfo"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

// GetDisplayName returns DisplayName
func (m CreateExternalDbSystemMacsConnectorDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetExternalDbSystemId returns ExternalDbSystemId
func (m CreateExternalDbSystemMacsConnectorDetails) GetExternalDbSystemId() *string {
	return m.ExternalDbSystemId
}

func (m CreateExternalDbSystemMacsConnectorDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateExternalDbSystemMacsConnectorDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateExternalDbSystemMacsConnectorDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateExternalDbSystemMacsConnectorDetails CreateExternalDbSystemMacsConnectorDetails
	s := struct {
		DiscriminatorParam string `json:"connectorType"`
		MarshalTypeCreateExternalDbSystemMacsConnectorDetails
	}{
		"MACS",
		(MarshalTypeCreateExternalDbSystemMacsConnectorDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CreateExternalDbSystemMacsConnectorDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName        *string                           `json:"displayName"`
		ConnectionInfo     externaldbsystemconnectioninfo    `json:"connectionInfo"`
		FreeformTags       map[string]string                 `json:"freeformTags"`
		DefinedTags        map[string]map[string]interface{} `json:"definedTags"`
		ExternalDbSystemId *string                           `json:"externalDbSystemId"`
		AgentId            *string                           `json:"agentId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	nn, e = model.ConnectionInfo.UnmarshalPolymorphicJSON(model.ConnectionInfo.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ConnectionInfo = nn.(ExternalDbSystemConnectionInfo)
	} else {
		m.ConnectionInfo = nil
	}

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.ExternalDbSystemId = model.ExternalDbSystemId

	m.AgentId = model.AgentId

	return
}
