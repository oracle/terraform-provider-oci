// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v41/common"
)

// ExternalMacsConnector An Oracle Cloud Infrastructure resource that uses the Management Agent cloud service (MACS) (https://docs.cloud.oracle.com/iaas/management-agents/index.html) to connect to an external Oracle Database.
type ExternalMacsConnector struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The user-friendly name for the
	// CreateExternalDatabaseConnectorDetails.
	// The name does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the
	// CreateExternalDatabaseConnectorDetails.
	Id *string `mandatory:"true" json:"id"`

	// The date and time the external connector was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the external database resource.
	ExternalDatabaseId *string `mandatory:"true" json:"externalDatabaseId"`

	// The status of connectivity to the external database.
	ConnectionStatus *string `mandatory:"true" json:"connectionStatus"`

	// The date and time the connectionStatus of this external connector was last updated.
	TimeConnectionStatusLastUpdated *common.SDKTime `mandatory:"true" json:"timeConnectionStatusLastUpdated"`

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

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The current lifecycle state of the external database connector resource.
	LifecycleState ExternalDatabaseConnectorLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

//GetCompartmentId returns CompartmentId
func (m ExternalMacsConnector) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetFreeformTags returns FreeformTags
func (m ExternalMacsConnector) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m ExternalMacsConnector) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetDisplayName returns DisplayName
func (m ExternalMacsConnector) GetDisplayName() *string {
	return m.DisplayName
}

//GetId returns Id
func (m ExternalMacsConnector) GetId() *string {
	return m.Id
}

//GetLifecycleState returns LifecycleState
func (m ExternalMacsConnector) GetLifecycleState() ExternalDatabaseConnectorLifecycleStateEnum {
	return m.LifecycleState
}

//GetLifecycleDetails returns LifecycleDetails
func (m ExternalMacsConnector) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

//GetTimeCreated returns TimeCreated
func (m ExternalMacsConnector) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetExternalDatabaseId returns ExternalDatabaseId
func (m ExternalMacsConnector) GetExternalDatabaseId() *string {
	return m.ExternalDatabaseId
}

//GetConnectionStatus returns ConnectionStatus
func (m ExternalMacsConnector) GetConnectionStatus() *string {
	return m.ConnectionStatus
}

//GetTimeConnectionStatusLastUpdated returns TimeConnectionStatusLastUpdated
func (m ExternalMacsConnector) GetTimeConnectionStatusLastUpdated() *common.SDKTime {
	return m.TimeConnectionStatusLastUpdated
}

func (m ExternalMacsConnector) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ExternalMacsConnector) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeExternalMacsConnector ExternalMacsConnector
	s := struct {
		DiscriminatorParam string `json:"connectorType"`
		MarshalTypeExternalMacsConnector
	}{
		"MACS",
		(MarshalTypeExternalMacsConnector)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *ExternalMacsConnector) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		FreeformTags                    map[string]string                           `json:"freeformTags"`
		DefinedTags                     map[string]map[string]interface{}           `json:"definedTags"`
		LifecycleDetails                *string                                     `json:"lifecycleDetails"`
		CompartmentId                   *string                                     `json:"compartmentId"`
		DisplayName                     *string                                     `json:"displayName"`
		Id                              *string                                     `json:"id"`
		LifecycleState                  ExternalDatabaseConnectorLifecycleStateEnum `json:"lifecycleState"`
		TimeCreated                     *common.SDKTime                             `json:"timeCreated"`
		ExternalDatabaseId              *string                                     `json:"externalDatabaseId"`
		ConnectionStatus                *string                                     `json:"connectionStatus"`
		TimeConnectionStatusLastUpdated *common.SDKTime                             `json:"timeConnectionStatusLastUpdated"`
		ConnectionString                *DatabaseConnectionString                   `json:"connectionString"`
		ConnectionCredentials           databaseconnectioncredentials               `json:"connectionCredentials"`
		ConnectorAgentId                *string                                     `json:"connectorAgentId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.LifecycleDetails = model.LifecycleDetails

	m.CompartmentId = model.CompartmentId

	m.DisplayName = model.DisplayName

	m.Id = model.Id

	m.LifecycleState = model.LifecycleState

	m.TimeCreated = model.TimeCreated

	m.ExternalDatabaseId = model.ExternalDatabaseId

	m.ConnectionStatus = model.ConnectionStatus

	m.TimeConnectionStatusLastUpdated = model.TimeConnectionStatusLastUpdated

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
