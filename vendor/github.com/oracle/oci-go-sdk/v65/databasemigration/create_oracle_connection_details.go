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

// CreateOracleConnectionDetails The information about a new Oracle Database Connection.
type CreateOracleConnectionDetails struct {

	// An object's Display Name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// OCI resource ID.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// OCI resource ID.
	VaultId *string `mandatory:"true" json:"vaultId"`

	// OCI resource ID.
	KeyId *string `mandatory:"true" json:"keyId"`

	// The username.
	Username *string `mandatory:"true" json:"username"`

	// The password.
	Password *string `mandatory:"true" json:"password"`

	// An object's Description.
	Description *string `mandatory:"false" json:"description"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// OCI resource ID.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// An array of Network Security Group OCIDs used to define network access for Connections.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// The username.
	ReplicationUsername *string `mandatory:"false" json:"replicationUsername"`

	// The password.
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

	// The username.
	SshUser *string `mandatory:"false" json:"sshUser"`

	// Sudo location
	SshSudoLocation *string `mandatory:"false" json:"sshSudoLocation"`

	// The Oracle technology type.
	TechnologyType OracleConnectionTechnologyTypeEnum `mandatory:"true" json:"technologyType"`
}

// GetDisplayName returns DisplayName
func (m CreateOracleConnectionDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m CreateOracleConnectionDetails) GetDescription() *string {
	return m.Description
}

// GetCompartmentId returns CompartmentId
func (m CreateOracleConnectionDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetFreeformTags returns FreeformTags
func (m CreateOracleConnectionDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateOracleConnectionDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetVaultId returns VaultId
func (m CreateOracleConnectionDetails) GetVaultId() *string {
	return m.VaultId
}

// GetKeyId returns KeyId
func (m CreateOracleConnectionDetails) GetKeyId() *string {
	return m.KeyId
}

// GetSubnetId returns SubnetId
func (m CreateOracleConnectionDetails) GetSubnetId() *string {
	return m.SubnetId
}

// GetNsgIds returns NsgIds
func (m CreateOracleConnectionDetails) GetNsgIds() []string {
	return m.NsgIds
}

// GetUsername returns Username
func (m CreateOracleConnectionDetails) GetUsername() *string {
	return m.Username
}

// GetPassword returns Password
func (m CreateOracleConnectionDetails) GetPassword() *string {
	return m.Password
}

// GetReplicationUsername returns ReplicationUsername
func (m CreateOracleConnectionDetails) GetReplicationUsername() *string {
	return m.ReplicationUsername
}

// GetReplicationPassword returns ReplicationPassword
func (m CreateOracleConnectionDetails) GetReplicationPassword() *string {
	return m.ReplicationPassword
}

func (m CreateOracleConnectionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateOracleConnectionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOracleConnectionTechnologyTypeEnum(string(m.TechnologyType)); !ok && m.TechnologyType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TechnologyType: %s. Supported values are: %s.", m.TechnologyType, strings.Join(GetOracleConnectionTechnologyTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateOracleConnectionDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateOracleConnectionDetails CreateOracleConnectionDetails
	s := struct {
		DiscriminatorParam string `json:"connectionType"`
		MarshalTypeCreateOracleConnectionDetails
	}{
		"ORACLE",
		(MarshalTypeCreateOracleConnectionDetails)(m),
	}

	return json.Marshal(&s)
}
