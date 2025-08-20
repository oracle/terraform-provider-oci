// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateCloudDbSystemMacsConnectorDetails The details for creating a cloud connector that is used to connect to a cloud DB system component
// using the Management Agent Cloud Service (MACS) (https://docs.oracle.com/iaas/management-agents/index.html).
type CreateCloudDbSystemMacsConnectorDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud DB system.
	CloudDbSystemId *string `mandatory:"true" json:"cloudDbSystemId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the management agent
	// used for the cloud DB system connector.
	AgentId *string `mandatory:"true" json:"agentId"`

	// The user-friendly name for the cloud connector. The name does not have to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	ConnectionInfo CloudDbSystemConnectionInfo `mandatory:"false" json:"connectionInfo"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

// GetDisplayName returns DisplayName
func (m CreateCloudDbSystemMacsConnectorDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetCloudDbSystemId returns CloudDbSystemId
func (m CreateCloudDbSystemMacsConnectorDetails) GetCloudDbSystemId() *string {
	return m.CloudDbSystemId
}

func (m CreateCloudDbSystemMacsConnectorDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateCloudDbSystemMacsConnectorDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateCloudDbSystemMacsConnectorDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateCloudDbSystemMacsConnectorDetails CreateCloudDbSystemMacsConnectorDetails
	s := struct {
		DiscriminatorParam string `json:"connectorType"`
		MarshalTypeCreateCloudDbSystemMacsConnectorDetails
	}{
		"MACS",
		(MarshalTypeCreateCloudDbSystemMacsConnectorDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CreateCloudDbSystemMacsConnectorDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName     *string                           `json:"displayName"`
		ConnectionInfo  clouddbsystemconnectioninfo       `json:"connectionInfo"`
		FreeformTags    map[string]string                 `json:"freeformTags"`
		DefinedTags     map[string]map[string]interface{} `json:"definedTags"`
		CloudDbSystemId *string                           `json:"cloudDbSystemId"`
		AgentId         *string                           `json:"agentId"`
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
		m.ConnectionInfo = nn.(CloudDbSystemConnectionInfo)
	} else {
		m.ConnectionInfo = nil
	}

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.CloudDbSystemId = model.CloudDbSystemId

	m.AgentId = model.AgentId

	return
}
