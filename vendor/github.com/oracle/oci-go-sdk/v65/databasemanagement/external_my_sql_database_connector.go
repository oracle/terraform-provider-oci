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
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExternalMySqlDatabaseConnector Details of external database connector.
type ExternalMySqlDatabaseConnector struct {

	// OCID of MySQL Database Connector.
	Id *string `mandatory:"true" json:"id"`

	// External MySQL Database Connector Name.
	Name *string `mandatory:"false" json:"name"`

	// OCID of compartment for the External MySQL connector.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// OCI Services associated with this connector.
	AssociatedServices *string `mandatory:"false" json:"associatedServices"`

	// OCID of MySQL Database resource
	ExternalDatabaseId *string `mandatory:"false" json:"externalDatabaseId"`

	// Connector update time.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Connector creation time.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Indicates lifecycle  state of the resource.
	LifecycleState LifecycleStatesEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Name of MySQL Database.
	SourceDatabase *string `mandatory:"false" json:"sourceDatabase"`

	// Type of MySQL Database.
	SourceDatabaseType MySqlTypeEnum `mandatory:"false" json:"sourceDatabaseType,omitempty"`

	// Agent Id of the MACS agent.
	MacsAgentId *string `mandatory:"false" json:"macsAgentId"`

	// Connection Status
	ConnectionStatus *string `mandatory:"false" json:"connectionStatus"`

	// Time when connection status was last updated.
	TimeConnectionStatusUpdated *common.SDKTime `mandatory:"false" json:"timeConnectionStatusUpdated"`

	// Host name for Connector.
	HostName *string `mandatory:"false" json:"hostName"`

	// Connector port.
	Port *int `mandatory:"false" json:"port"`

	// Connector Type.
	ConnectorType MySqlConnectorTypeEnum `mandatory:"false" json:"connectorType,omitempty"`

	// Network Protocol.
	NetworkProtocol MySqlNetworkProtocolTypeEnum `mandatory:"false" json:"networkProtocol,omitempty"`

	// Credential type used to connect to database.
	CredentialType MySqlCredTypeEnum `mandatory:"false" json:"credentialType,omitempty"`

	// OCID of the SSL secret, if TCPS with SSL is used to connect to database.
	SslSecretId *string `mandatory:"false" json:"sslSecretId"`

	// Name of the SSL secret, if TCPS with SSL is used to connect to database.
	SslSecretName *string `mandatory:"false" json:"sslSecretName"`
}

func (m ExternalMySqlDatabaseConnector) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExternalMySqlDatabaseConnector) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingLifecycleStatesEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStatesEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMySqlTypeEnum(string(m.SourceDatabaseType)); !ok && m.SourceDatabaseType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SourceDatabaseType: %s. Supported values are: %s.", m.SourceDatabaseType, strings.Join(GetMySqlTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMySqlConnectorTypeEnum(string(m.ConnectorType)); !ok && m.ConnectorType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConnectorType: %s. Supported values are: %s.", m.ConnectorType, strings.Join(GetMySqlConnectorTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMySqlNetworkProtocolTypeEnum(string(m.NetworkProtocol)); !ok && m.NetworkProtocol != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NetworkProtocol: %s. Supported values are: %s.", m.NetworkProtocol, strings.Join(GetMySqlNetworkProtocolTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMySqlCredTypeEnum(string(m.CredentialType)); !ok && m.CredentialType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CredentialType: %s. Supported values are: %s.", m.CredentialType, strings.Join(GetMySqlCredTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
