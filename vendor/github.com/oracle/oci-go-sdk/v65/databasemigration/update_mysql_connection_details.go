// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// UpdateMysqlConnectionDetails The information to update a MySQL Connection.
type UpdateMysqlConnectionDetails struct {

	// An object's Display Name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// An object's Description.
	Description *string `mandatory:"false" json:"description"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// OCI resource ID.
	VaultId *string `mandatory:"false" json:"vaultId"`

	// OCI resource ID.
	KeyId *string `mandatory:"false" json:"keyId"`

	// OCI resource ID.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// An array of Network Security Group OCIDs used to define network access for Connections.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// The username.
	Username *string `mandatory:"false" json:"username"`

	// The password.
	Password *string `mandatory:"false" json:"password"`

	// The username.
	ReplicationUsername *string `mandatory:"false" json:"replicationUsername"`

	// The password.
	ReplicationPassword *string `mandatory:"false" json:"replicationPassword"`

	// The name or address of a host.
	Host *string `mandatory:"false" json:"host"`

	// The port of an endpoint usually specified for a connection.
	Port *int `mandatory:"false" json:"port"`

	// The name of the database.
	DatabaseName *string `mandatory:"false" json:"databaseName"`

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

	// Security Type for MySQL.
	SecurityProtocol MysqlConnectionSecurityProtocolEnum `mandatory:"false" json:"securityProtocol,omitempty"`

	// SSL modes for MySQL.
	SslMode MysqlConnectionSslModeEnum `mandatory:"false" json:"sslMode,omitempty"`
}

// GetDisplayName returns DisplayName
func (m UpdateMysqlConnectionDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m UpdateMysqlConnectionDetails) GetDescription() *string {
	return m.Description
}

// GetFreeformTags returns FreeformTags
func (m UpdateMysqlConnectionDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m UpdateMysqlConnectionDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetVaultId returns VaultId
func (m UpdateMysqlConnectionDetails) GetVaultId() *string {
	return m.VaultId
}

// GetKeyId returns KeyId
func (m UpdateMysqlConnectionDetails) GetKeyId() *string {
	return m.KeyId
}

// GetSubnetId returns SubnetId
func (m UpdateMysqlConnectionDetails) GetSubnetId() *string {
	return m.SubnetId
}

// GetNsgIds returns NsgIds
func (m UpdateMysqlConnectionDetails) GetNsgIds() []string {
	return m.NsgIds
}

// GetUsername returns Username
func (m UpdateMysqlConnectionDetails) GetUsername() *string {
	return m.Username
}

// GetPassword returns Password
func (m UpdateMysqlConnectionDetails) GetPassword() *string {
	return m.Password
}

// GetReplicationUsername returns ReplicationUsername
func (m UpdateMysqlConnectionDetails) GetReplicationUsername() *string {
	return m.ReplicationUsername
}

// GetReplicationPassword returns ReplicationPassword
func (m UpdateMysqlConnectionDetails) GetReplicationPassword() *string {
	return m.ReplicationPassword
}

func (m UpdateMysqlConnectionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateMysqlConnectionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

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
func (m UpdateMysqlConnectionDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateMysqlConnectionDetails UpdateMysqlConnectionDetails
	s := struct {
		DiscriminatorParam string `json:"connectionType"`
		MarshalTypeUpdateMysqlConnectionDetails
	}{
		"MYSQL",
		(MarshalTypeUpdateMysqlConnectionDetails)(m),
	}

	return json.Marshal(&s)
}
