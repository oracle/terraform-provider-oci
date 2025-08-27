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

// CloudDbSystemMacsConnector The details of a cloud DB system connector that uses the
// Management Agent Cloud Service (MACS) (https://docs.oracle.com/iaas/management-agents/index.html)
// to connect to a cloud DB system component.
type CloudDbSystemMacsConnector struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud DB system connector.
	Id *string `mandatory:"true" json:"id"`

	// The user-friendly name for the cloud connector. The name does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud DB system that the connector is a part of.
	CloudDbSystemId *string `mandatory:"true" json:"cloudDbSystemId"`

	// The date and time the cloud DB system connector was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the cloud DB system connector was last updated.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the management agent
	// used for the cloud DB system connector.
	AgentId *string `mandatory:"true" json:"agentId"`

	// The status of connectivity to the cloud DB system component.
	ConnectionStatus *string `mandatory:"false" json:"connectionStatus"`

	// The error message indicating the reason for connection failure or `null` if
	// the connection was successful.
	ConnectionFailureMessage *string `mandatory:"false" json:"connectionFailureMessage"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The date and time the connectionStatus of the cloud DB system connector was last updated.
	TimeConnectionStatusLastUpdated *common.SDKTime `mandatory:"false" json:"timeConnectionStatusLastUpdated"`

	ConnectionInfo CloudDbSystemConnectionInfo `mandatory:"false" json:"connectionInfo"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// System tags can be viewed by users, but can only be created by the system.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The current lifecycle state of the cloud DB system connector.
	LifecycleState CloudDbSystemConnectorLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

// GetId returns Id
func (m CloudDbSystemMacsConnector) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m CloudDbSystemMacsConnector) GetDisplayName() *string {
	return m.DisplayName
}

// GetCompartmentId returns CompartmentId
func (m CloudDbSystemMacsConnector) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetCloudDbSystemId returns CloudDbSystemId
func (m CloudDbSystemMacsConnector) GetCloudDbSystemId() *string {
	return m.CloudDbSystemId
}

// GetConnectionStatus returns ConnectionStatus
func (m CloudDbSystemMacsConnector) GetConnectionStatus() *string {
	return m.ConnectionStatus
}

// GetConnectionFailureMessage returns ConnectionFailureMessage
func (m CloudDbSystemMacsConnector) GetConnectionFailureMessage() *string {
	return m.ConnectionFailureMessage
}

// GetLifecycleState returns LifecycleState
func (m CloudDbSystemMacsConnector) GetLifecycleState() CloudDbSystemConnectorLifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m CloudDbSystemMacsConnector) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetTimeConnectionStatusLastUpdated returns TimeConnectionStatusLastUpdated
func (m CloudDbSystemMacsConnector) GetTimeConnectionStatusLastUpdated() *common.SDKTime {
	return m.TimeConnectionStatusLastUpdated
}

// GetTimeCreated returns TimeCreated
func (m CloudDbSystemMacsConnector) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m CloudDbSystemMacsConnector) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

func (m CloudDbSystemMacsConnector) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CloudDbSystemMacsConnector) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCloudDbSystemConnectorLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetCloudDbSystemConnectorLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CloudDbSystemMacsConnector) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCloudDbSystemMacsConnector CloudDbSystemMacsConnector
	s := struct {
		DiscriminatorParam string `json:"connectorType"`
		MarshalTypeCloudDbSystemMacsConnector
	}{
		"MACS",
		(MarshalTypeCloudDbSystemMacsConnector)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CloudDbSystemMacsConnector) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ConnectionStatus                *string                                  `json:"connectionStatus"`
		ConnectionFailureMessage        *string                                  `json:"connectionFailureMessage"`
		LifecycleDetails                *string                                  `json:"lifecycleDetails"`
		TimeConnectionStatusLastUpdated *common.SDKTime                          `json:"timeConnectionStatusLastUpdated"`
		ConnectionInfo                  clouddbsystemconnectioninfo              `json:"connectionInfo"`
		FreeformTags                    map[string]string                        `json:"freeformTags"`
		DefinedTags                     map[string]map[string]interface{}        `json:"definedTags"`
		SystemTags                      map[string]map[string]interface{}        `json:"systemTags"`
		Id                              *string                                  `json:"id"`
		DisplayName                     *string                                  `json:"displayName"`
		CompartmentId                   *string                                  `json:"compartmentId"`
		CloudDbSystemId                 *string                                  `json:"cloudDbSystemId"`
		LifecycleState                  CloudDbSystemConnectorLifecycleStateEnum `json:"lifecycleState"`
		TimeCreated                     *common.SDKTime                          `json:"timeCreated"`
		TimeUpdated                     *common.SDKTime                          `json:"timeUpdated"`
		AgentId                         *string                                  `json:"agentId"`
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
		m.ConnectionInfo = nn.(CloudDbSystemConnectionInfo)
	} else {
		m.ConnectionInfo = nil
	}

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.DisplayName = model.DisplayName

	m.CompartmentId = model.CompartmentId

	m.CloudDbSystemId = model.CloudDbSystemId

	m.LifecycleState = model.LifecycleState

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.AgentId = model.AgentId

	return
}
