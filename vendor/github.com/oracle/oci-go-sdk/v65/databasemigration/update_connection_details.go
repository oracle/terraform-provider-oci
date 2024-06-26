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

// UpdateConnectionDetails The information to update a Connection.
type UpdateConnectionDetails interface {

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	GetDisplayName() *string

	// A user-friendly description. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	GetDescription() *string

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags. Example: {"Department": "Finance"}
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// OCI resource ID.
	GetVaultId() *string

	// The OCID of the key used in cryptographic operations.
	GetKeyId() *string

	// OCI resource ID.
	GetSubnetId() *string

	// An array of Network Security Group OCIDs used to define network access for Connections.
	GetNsgIds() []string

	// The username (credential) used when creating or updating this resource.
	GetUsername() *string

	// The password (credential) used when creating or updating this resource.
	GetPassword() *string

	// The username (credential) used when creating or updating this resource.
	GetReplicationUsername() *string

	// The password (credential) used when creating or updating this resource.
	GetReplicationPassword() *string
}

type updateconnectiondetails struct {
	JsonData            []byte
	DisplayName         *string                           `mandatory:"false" json:"displayName"`
	Description         *string                           `mandatory:"false" json:"description"`
	FreeformTags        map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags         map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	VaultId             *string                           `mandatory:"false" json:"vaultId"`
	KeyId               *string                           `mandatory:"false" json:"keyId"`
	SubnetId            *string                           `mandatory:"false" json:"subnetId"`
	NsgIds              []string                          `mandatory:"false" json:"nsgIds"`
	Username            *string                           `mandatory:"false" json:"username"`
	Password            *string                           `mandatory:"false" json:"password"`
	ReplicationUsername *string                           `mandatory:"false" json:"replicationUsername"`
	ReplicationPassword *string                           `mandatory:"false" json:"replicationPassword"`
	ConnectionType      string                            `json:"connectionType"`
}

// UnmarshalJSON unmarshals json
func (m *updateconnectiondetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdateconnectiondetails updateconnectiondetails
	s := struct {
		Model Unmarshalerupdateconnectiondetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DisplayName = s.Model.DisplayName
	m.Description = s.Model.Description
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.VaultId = s.Model.VaultId
	m.KeyId = s.Model.KeyId
	m.SubnetId = s.Model.SubnetId
	m.NsgIds = s.Model.NsgIds
	m.Username = s.Model.Username
	m.Password = s.Model.Password
	m.ReplicationUsername = s.Model.ReplicationUsername
	m.ReplicationPassword = s.Model.ReplicationPassword
	m.ConnectionType = s.Model.ConnectionType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updateconnectiondetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ConnectionType {
	case "ORACLE":
		mm := UpdateOracleConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MYSQL":
		mm := UpdateMysqlConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for UpdateConnectionDetails: %s.", m.ConnectionType)
		return *m, nil
	}
}

// GetDisplayName returns DisplayName
func (m updateconnectiondetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m updateconnectiondetails) GetDescription() *string {
	return m.Description
}

// GetFreeformTags returns FreeformTags
func (m updateconnectiondetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m updateconnectiondetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetVaultId returns VaultId
func (m updateconnectiondetails) GetVaultId() *string {
	return m.VaultId
}

// GetKeyId returns KeyId
func (m updateconnectiondetails) GetKeyId() *string {
	return m.KeyId
}

// GetSubnetId returns SubnetId
func (m updateconnectiondetails) GetSubnetId() *string {
	return m.SubnetId
}

// GetNsgIds returns NsgIds
func (m updateconnectiondetails) GetNsgIds() []string {
	return m.NsgIds
}

// GetUsername returns Username
func (m updateconnectiondetails) GetUsername() *string {
	return m.Username
}

// GetPassword returns Password
func (m updateconnectiondetails) GetPassword() *string {
	return m.Password
}

// GetReplicationUsername returns ReplicationUsername
func (m updateconnectiondetails) GetReplicationUsername() *string {
	return m.ReplicationUsername
}

// GetReplicationPassword returns ReplicationPassword
func (m updateconnectiondetails) GetReplicationPassword() *string {
	return m.ReplicationPassword
}

func (m updateconnectiondetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updateconnectiondetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
