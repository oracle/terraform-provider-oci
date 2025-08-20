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

// CreateMySqlDatabaseConnectorDetails Create Details of external database connector.
type CreateMySqlDatabaseConnectorDetails struct {

	// External MySQL Database Connector Name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Agent Id of the MACS agent.
	MacsAgentId *string `mandatory:"true" json:"macsAgentId"`

	// Host name for Connector.
	HostName *string `mandatory:"true" json:"hostName"`

	// Port number to connect to External MySQL Database.
	Port *int `mandatory:"true" json:"port"`

	// Protocol to be used to connect to External MySQL Database; TCP, TCP with SSL or Socket.
	NetworkProtocol MySqlNetworkProtocolTypeEnum `mandatory:"true" json:"networkProtocol"`

	// OCID of MySQL Database resource.
	ExternalDatabaseId *string `mandatory:"true" json:"externalDatabaseId"`

	// Type of the credential.
	CredentialType MySqlCredTypeEnum `mandatory:"true" json:"credentialType"`

	// If using existing SSL secret to connect, OCID for the secret resource.
	SslSecretId *string `mandatory:"true" json:"sslSecretId"`
}

func (m CreateMySqlDatabaseConnectorDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateMySqlDatabaseConnectorDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
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
