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

// ExternalDbSystemMacsConnector The details of an external DB system connector that uses the
// Management Agent Cloud Service (MACS) (https://docs.cloud.oracle.com/iaas/management-agents/index.html)
// to connect to an external DB system component.
type ExternalDbSystemMacsConnector struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the external DB system connector.
	Id *string `mandatory:"true" json:"id"`

	// The user-friendly name for the external connector. The name does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the external DB system that the connector is a part of.
	ExternalDbSystemId *string `mandatory:"true" json:"externalDbSystemId"`

	// The date and time the external DB system connector was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the external DB system connector was last updated.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the management agent
	// used for the external DB system connector.
	AgentId *string `mandatory:"true" json:"agentId"`

	// The status of connectivity to the external DB system component.
	ConnectionStatus *string `mandatory:"false" json:"connectionStatus"`

	// The error message indicating the reason for connection failure or `null` if
	// the connection was successful.
	ConnectionFailureMessage *string `mandatory:"false" json:"connectionFailureMessage"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The date and time the connectionStatus of the external DB system connector was last updated.
	TimeConnectionStatusLastUpdated *common.SDKTime `mandatory:"false" json:"timeConnectionStatusLastUpdated"`

	ConnectionInfo ExternalDbSystemConnectionInfo `mandatory:"false" json:"connectionInfo"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The current lifecycle state of the external DB system connector.
	LifecycleState ExternalDbSystemConnectorLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

// GetId returns Id
func (m ExternalDbSystemMacsConnector) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m ExternalDbSystemMacsConnector) GetDisplayName() *string {
	return m.DisplayName
}

// GetCompartmentId returns CompartmentId
func (m ExternalDbSystemMacsConnector) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetExternalDbSystemId returns ExternalDbSystemId
func (m ExternalDbSystemMacsConnector) GetExternalDbSystemId() *string {
	return m.ExternalDbSystemId
}

// GetConnectionStatus returns ConnectionStatus
func (m ExternalDbSystemMacsConnector) GetConnectionStatus() *string {
	return m.ConnectionStatus
}

// GetConnectionFailureMessage returns ConnectionFailureMessage
func (m ExternalDbSystemMacsConnector) GetConnectionFailureMessage() *string {
	return m.ConnectionFailureMessage
}

// GetLifecycleState returns LifecycleState
func (m ExternalDbSystemMacsConnector) GetLifecycleState() ExternalDbSystemConnectorLifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m ExternalDbSystemMacsConnector) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetTimeConnectionStatusLastUpdated returns TimeConnectionStatusLastUpdated
func (m ExternalDbSystemMacsConnector) GetTimeConnectionStatusLastUpdated() *common.SDKTime {
	return m.TimeConnectionStatusLastUpdated
}

// GetTimeCreated returns TimeCreated
func (m ExternalDbSystemMacsConnector) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m ExternalDbSystemMacsConnector) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

func (m ExternalDbSystemMacsConnector) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExternalDbSystemMacsConnector) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingExternalDbSystemConnectorLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetExternalDbSystemConnectorLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ExternalDbSystemMacsConnector) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeExternalDbSystemMacsConnector ExternalDbSystemMacsConnector
	s := struct {
		DiscriminatorParam string `json:"connectorType"`
		MarshalTypeExternalDbSystemMacsConnector
	}{
		"MACS",
		(MarshalTypeExternalDbSystemMacsConnector)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *ExternalDbSystemMacsConnector) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ConnectionStatus                *string                                     `json:"connectionStatus"`
		ConnectionFailureMessage        *string                                     `json:"connectionFailureMessage"`
		LifecycleDetails                *string                                     `json:"lifecycleDetails"`
		TimeConnectionStatusLastUpdated *common.SDKTime                             `json:"timeConnectionStatusLastUpdated"`
		ConnectionInfo                  externaldbsystemconnectioninfo              `json:"connectionInfo"`
		FreeformTags                    map[string]string                           `json:"freeformTags"`
		DefinedTags                     map[string]map[string]interface{}           `json:"definedTags"`
		Id                              *string                                     `json:"id"`
		DisplayName                     *string                                     `json:"displayName"`
		CompartmentId                   *string                                     `json:"compartmentId"`
		ExternalDbSystemId              *string                                     `json:"externalDbSystemId"`
		LifecycleState                  ExternalDbSystemConnectorLifecycleStateEnum `json:"lifecycleState"`
		TimeCreated                     *common.SDKTime                             `json:"timeCreated"`
		TimeUpdated                     *common.SDKTime                             `json:"timeUpdated"`
		AgentId                         *string                                     `json:"agentId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.ConnectionStatus = model.ConnectionStatus

	m.ConnectionFailureMessage = model.ConnectionFailureMessage

	m.LifecycleDetails = model.LifecycleDetails

	m.TimeConnectionStatusLastUpdated = model.TimeConnectionStatusLastUpdated

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

	m.Id = model.Id

	m.DisplayName = model.DisplayName

	m.CompartmentId = model.CompartmentId

	m.ExternalDbSystemId = model.ExternalDbSystemId

	m.LifecycleState = model.LifecycleState

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.AgentId = model.AgentId

	return
}
