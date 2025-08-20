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

// MySqlDatabaseConnectorSummary Details of external database connector.
type MySqlDatabaseConnectorSummary struct {

	// OCID of MySQL Database Connector.
	Id *string `mandatory:"true" json:"id"`

	// External MySQL Database Connector Name
	Name *string `mandatory:"false" json:"name"`

	// OCID of compartment for the External MySQL connector.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// OCI Services associated with this connector.
	AssociatedServices *string `mandatory:"false" json:"associatedServices"`

	// Connector creation time.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Connector update time.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Name of MySQL Database.
	SourceDatabase *string `mandatory:"false" json:"sourceDatabase"`

	// Type of MySQL Database.
	SourceDatabaseType MySqlTypeEnum `mandatory:"false" json:"sourceDatabaseType,omitempty"`

	// Connection Status.
	ConnectionStatus *string `mandatory:"false" json:"connectionStatus"`

	// Time when connection status was last updated.
	TimeConnectionStatusUpdated *common.SDKTime `mandatory:"false" json:"timeConnectionStatusUpdated"`

	// Host name for Connector.
	HostName *string `mandatory:"false" json:"hostName"`

	// Agent Id of the MACS agent.
	MacsAgentId *string `mandatory:"false" json:"macsAgentId"`

	// Connector port.
	Port *int `mandatory:"false" json:"port"`

	// Connector Type.
	ConnectorType MySqlConnectorTypeEnum `mandatory:"false" json:"connectorType,omitempty"`

	// Network Protocol.
	NetworkProtocol MySqlNetworkProtocolTypeEnum `mandatory:"false" json:"networkProtocol,omitempty"`

	// Credential type used to connect to database.
	CredentialType MySqlCredTypeEnum `mandatory:"false" json:"credentialType,omitempty"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// System tags can be viewed by users, but can only be created by the system.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// Indicates lifecycle  state of the resource.
	LifecycleState LifecycleStatesEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m MySqlDatabaseConnectorSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MySqlDatabaseConnectorSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

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
	if _, ok := GetMappingLifecycleStatesEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStatesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
