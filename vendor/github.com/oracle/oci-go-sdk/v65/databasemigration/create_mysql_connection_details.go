// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateMysqlConnectionDetails The information about a new MySQL Connection.
type CreateMysqlConnectionDetails struct {

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// OCI resource ID.
	VaultId *string `mandatory:"true" json:"vaultId"`

	// The OCID of the key used in cryptographic operations.
	KeyId *string `mandatory:"true" json:"keyId"`

	// The username (credential) used when creating or updating this resource.
	Username *string `mandatory:"true" json:"username"`

	// The password (credential) used when creating or updating this resource.
	Password *string `mandatory:"true" json:"password"`

	// The name of the database being referenced.
	DatabaseName *string `mandatory:"true" json:"databaseName"`

	// A user-friendly description. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags. Example: {"Department": "Finance"}
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// OCI resource ID.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// An array of Network Security Group OCIDs used to define network access for Connections.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// The username (credential) used when creating or updating this resource.
	ReplicationUsername *string `mandatory:"false" json:"replicationUsername"`

	// The password (credential) used when creating or updating this resource.
	ReplicationPassword *string `mandatory:"false" json:"replicationPassword"`

	// The IP Address of the host.
	Host *string `mandatory:"false" json:"host"`

	// The port to be used for the connection.
	Port *int `mandatory:"false" json:"port"`

	// Database Certificate - The base64 encoded content of mysql.pem file
	// containing the server public key (for 1 and 2-way SSL).
	SslCa *string `mandatory:"false" json:"sslCa"`

	// Certificates revoked by certificate authorities (CA).
	// Server certificate must not be on this list (for 1 and 2-way SSL).
	// Note: This is an optional and that too only applicable if TLS/MTLS option is selected.
	SslCrl *string `mandatory:"false" json:"sslCrl"`

	// Client Certificate - The base64 encoded content of client-cert.pem file
	// containing the client public key (for 2-way SSL).
	SslCert *string `mandatory:"false" json:"sslCert"`

	// Client Key - The client-key.pem containing the client private key (for 2-way SSL).
	SslKey *string `mandatory:"false" json:"sslKey"`

	// An array of name-value pair attribute entries.
	AdditionalAttributes []NameValuePair `mandatory:"false" json:"additionalAttributes"`

	// The OCID of the database system being referenced.
	DbSystemId *string `mandatory:"false" json:"dbSystemId"`

	// The type of MySQL source or target connection.
	// Example: OCI_MYSQL represents OCI MySQL HeatWave Database Service
	TechnologyType MysqlConnectionTechnologyTypeEnum `mandatory:"true" json:"technologyType"`

	// Security Type for MySQL.
	SecurityProtocol MysqlConnectionSecurityProtocolEnum `mandatory:"true" json:"securityProtocol"`

	// SSL modes for MySQL.
	SslMode MysqlConnectionSslModeEnum `mandatory:"false" json:"sslMode,omitempty"`
}

// GetDisplayName returns DisplayName
func (m CreateMysqlConnectionDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m CreateMysqlConnectionDetails) GetDescription() *string {
	return m.Description
}

// GetCompartmentId returns CompartmentId
func (m CreateMysqlConnectionDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetFreeformTags returns FreeformTags
func (m CreateMysqlConnectionDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateMysqlConnectionDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetVaultId returns VaultId
func (m CreateMysqlConnectionDetails) GetVaultId() *string {
	return m.VaultId
}

// GetKeyId returns KeyId
func (m CreateMysqlConnectionDetails) GetKeyId() *string {
	return m.KeyId
}

// GetSubnetId returns SubnetId
func (m CreateMysqlConnectionDetails) GetSubnetId() *string {
	return m.SubnetId
}

// GetNsgIds returns NsgIds
func (m CreateMysqlConnectionDetails) GetNsgIds() []string {
	return m.NsgIds
}

// GetUsername returns Username
func (m CreateMysqlConnectionDetails) GetUsername() *string {
	return m.Username
}

// GetPassword returns Password
func (m CreateMysqlConnectionDetails) GetPassword() *string {
	return m.Password
}

// GetReplicationUsername returns ReplicationUsername
func (m CreateMysqlConnectionDetails) GetReplicationUsername() *string {
	return m.ReplicationUsername
}

// GetReplicationPassword returns ReplicationPassword
func (m CreateMysqlConnectionDetails) GetReplicationPassword() *string {
	return m.ReplicationPassword
}

func (m CreateMysqlConnectionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateMysqlConnectionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingMysqlConnectionTechnologyTypeEnum(string(m.TechnologyType)); !ok && m.TechnologyType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TechnologyType: %s. Supported values are: %s.", m.TechnologyType, strings.Join(GetMysqlConnectionTechnologyTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMysqlConnectionSecurityProtocolEnum(string(m.SecurityProtocol)); !ok && m.SecurityProtocol != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SecurityProtocol: %s. Supported values are: %s.", m.SecurityProtocol, strings.Join(GetMysqlConnectionSecurityProtocolEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMysqlConnectionSslModeEnum(string(m.SslMode)); !ok && m.SslMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SslMode: %s. Supported values are: %s.", m.SslMode, strings.Join(GetMysqlConnectionSslModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateMysqlConnectionDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateMysqlConnectionDetails CreateMysqlConnectionDetails
	s := struct {
		DiscriminatorParam string `json:"connectionType"`
		MarshalTypeCreateMysqlConnectionDetails
	}{
		"MYSQL",
		(MarshalTypeCreateMysqlConnectionDetails)(m),
	}

	return json.Marshal(&s)
}
