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

// UpdateOracleConnectionDetails The information to update an Oracle Database Connection.
type UpdateOracleConnectionDetails struct {

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

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
	VaultId *string `mandatory:"false" json:"vaultId"`

	// The OCID of the key used in cryptographic operations.
	KeyId *string `mandatory:"false" json:"keyId"`

	// OCI resource ID.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// An array of Network Security Group OCIDs used to define network access for Connections.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// The username (credential) used when creating or updating this resource.
	Username *string `mandatory:"false" json:"username"`

	// The password (credential) used when creating or updating this resource.
	Password *string `mandatory:"false" json:"password"`

	// The username (credential) used when creating or updating this resource.
	ReplicationUsername *string `mandatory:"false" json:"replicationUsername"`

	// The password (credential) used when creating or updating this resource.
	ReplicationPassword *string `mandatory:"false" json:"replicationPassword"`

	// Connect descriptor or Easy Connect Naming method used to connect to a database.
	ConnectionString *string `mandatory:"false" json:"connectionString"`

	// The wallet contents used to make connections to a database.  This
	// attribute is expected to be base64 encoded.
	Wallet *string `mandatory:"false" json:"wallet"`

	// The OCID of the database being referenced.
	DatabaseId *string `mandatory:"false" json:"databaseId"`

	// Name of the host the SSH key is valid for.
	SshHost *string `mandatory:"false" json:"sshHost"`

	// Private SSH key string.
	SshKey *string `mandatory:"false" json:"sshKey"`

	// The username (credential) used when creating or updating this resource.
	SshUser *string `mandatory:"false" json:"sshUser"`

	// Sudo location
	SshSudoLocation *string `mandatory:"false" json:"sshSudoLocation"`
}

// GetDisplayName returns DisplayName
func (m UpdateOracleConnectionDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m UpdateOracleConnectionDetails) GetDescription() *string {
	return m.Description
}

// GetFreeformTags returns FreeformTags
func (m UpdateOracleConnectionDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m UpdateOracleConnectionDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetVaultId returns VaultId
func (m UpdateOracleConnectionDetails) GetVaultId() *string {
	return m.VaultId
}

// GetKeyId returns KeyId
func (m UpdateOracleConnectionDetails) GetKeyId() *string {
	return m.KeyId
}

// GetSubnetId returns SubnetId
func (m UpdateOracleConnectionDetails) GetSubnetId() *string {
	return m.SubnetId
}

// GetNsgIds returns NsgIds
func (m UpdateOracleConnectionDetails) GetNsgIds() []string {
	return m.NsgIds
}

// GetUsername returns Username
func (m UpdateOracleConnectionDetails) GetUsername() *string {
	return m.Username
}

// GetPassword returns Password
func (m UpdateOracleConnectionDetails) GetPassword() *string {
	return m.Password
}

// GetReplicationUsername returns ReplicationUsername
func (m UpdateOracleConnectionDetails) GetReplicationUsername() *string {
	return m.ReplicationUsername
}

// GetReplicationPassword returns ReplicationPassword
func (m UpdateOracleConnectionDetails) GetReplicationPassword() *string {
	return m.ReplicationPassword
}

func (m UpdateOracleConnectionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateOracleConnectionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateOracleConnectionDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateOracleConnectionDetails UpdateOracleConnectionDetails
	s := struct {
		DiscriminatorParam string `json:"connectionType"`
		MarshalTypeUpdateOracleConnectionDetails
	}{
		"ORACLE",
		(MarshalTypeUpdateOracleConnectionDetails)(m),
	}

	return json.Marshal(&s)
}
