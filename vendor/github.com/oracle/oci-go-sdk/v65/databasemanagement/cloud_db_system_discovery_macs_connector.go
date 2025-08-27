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

// CloudDbSystemDiscoveryMacsConnector The details of a cloud DB system connector that uses the
// Management Agent Cloud Service (MACS) (https://docs.oracle.com/iaas/management-agents/index.html)
// to connect to a cloud DB system component.
type CloudDbSystemDiscoveryMacsConnector struct {

	// The user-friendly name for the cloud connector. The name does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the management agent
	// used for the cloud DB system connector.
	AgentId *string `mandatory:"true" json:"agentId"`

	// The status of connectivity to the cloud DB system component.
	ConnectionStatus *string `mandatory:"false" json:"connectionStatus"`

	// The error message indicating the reason for connection failure or `null` if
	// the connection was successful.
	ConnectionFailureMessage *string `mandatory:"false" json:"connectionFailureMessage"`

	// The date and time the connectionStatus of the cloud DB system connector was last updated.
	TimeConnectionStatusLastUpdated *common.SDKTime `mandatory:"false" json:"timeConnectionStatusLastUpdated"`

	ConnectionInfo CloudDbSystemConnectionInfo `mandatory:"false" json:"connectionInfo"`
}

// GetDisplayName returns DisplayName
func (m CloudDbSystemDiscoveryMacsConnector) GetDisplayName() *string {
	return m.DisplayName
}

// GetConnectionStatus returns ConnectionStatus
func (m CloudDbSystemDiscoveryMacsConnector) GetConnectionStatus() *string {
	return m.ConnectionStatus
}

// GetConnectionFailureMessage returns ConnectionFailureMessage
func (m CloudDbSystemDiscoveryMacsConnector) GetConnectionFailureMessage() *string {
	return m.ConnectionFailureMessage
}

// GetTimeConnectionStatusLastUpdated returns TimeConnectionStatusLastUpdated
func (m CloudDbSystemDiscoveryMacsConnector) GetTimeConnectionStatusLastUpdated() *common.SDKTime {
	return m.TimeConnectionStatusLastUpdated
}

func (m CloudDbSystemDiscoveryMacsConnector) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CloudDbSystemDiscoveryMacsConnector) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CloudDbSystemDiscoveryMacsConnector) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCloudDbSystemDiscoveryMacsConnector CloudDbSystemDiscoveryMacsConnector
	s := struct {
		DiscriminatorParam string `json:"connectorType"`
		MarshalTypeCloudDbSystemDiscoveryMacsConnector
	}{
		"MACS",
		(MarshalTypeCloudDbSystemDiscoveryMacsConnector)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CloudDbSystemDiscoveryMacsConnector) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ConnectionStatus                *string                     `json:"connectionStatus"`
		ConnectionFailureMessage        *string                     `json:"connectionFailureMessage"`
		TimeConnectionStatusLastUpdated *common.SDKTime             `json:"timeConnectionStatusLastUpdated"`
		ConnectionInfo                  clouddbsystemconnectioninfo `json:"connectionInfo"`
		DisplayName                     *string                     `json:"displayName"`
		AgentId                         *string                     `json:"agentId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.ConnectionStatus = model.ConnectionStatus

	m.ConnectionFailureMessage = model.ConnectionFailureMessage

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

	m.DisplayName = model.DisplayName

	m.AgentId = model.AgentId

	return
}
