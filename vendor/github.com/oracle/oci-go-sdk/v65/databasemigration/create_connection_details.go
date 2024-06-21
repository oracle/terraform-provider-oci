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

// CreateConnectionDetails The information about a new Connection.
type CreateConnectionDetails interface {

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	GetDisplayName() *string

	// The OCID of the compartment.
	GetCompartmentId() *string

	// OCI resource ID.
	GetVaultId() *string

	// The OCID of the key used in cryptographic operations.
	GetKeyId() *string

	// The username (credential) used when creating or updating this resource.
	GetUsername() *string

	// The password (credential) used when creating or updating this resource.
	GetPassword() *string

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
	GetSubnetId() *string

	// An array of Network Security Group OCIDs used to define network access for Connections.
	GetNsgIds() []string

	// The username (credential) used when creating or updating this resource.
	GetReplicationUsername() *string

	// The password (credential) used when creating or updating this resource.
	GetReplicationPassword() *string
}

type createconnectiondetails struct {
	JsonData            []byte
	Description         *string                           `mandatory:"false" json:"description"`
	FreeformTags        map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags         map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	SubnetId            *string                           `mandatory:"false" json:"subnetId"`
	NsgIds              []string                          `mandatory:"false" json:"nsgIds"`
	ReplicationUsername *string                           `mandatory:"false" json:"replicationUsername"`
	ReplicationPassword *string                           `mandatory:"false" json:"replicationPassword"`
	DisplayName         *string                           `mandatory:"true" json:"displayName"`
	CompartmentId       *string                           `mandatory:"true" json:"compartmentId"`
	VaultId             *string                           `mandatory:"true" json:"vaultId"`
	KeyId               *string                           `mandatory:"true" json:"keyId"`
	Username            *string                           `mandatory:"true" json:"username"`
	Password            *string                           `mandatory:"true" json:"password"`
	ConnectionType      string                            `json:"connectionType"`
}

// UnmarshalJSON unmarshals json
func (m *createconnectiondetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreateconnectiondetails createconnectiondetails
	s := struct {
		Model Unmarshalercreateconnectiondetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DisplayName = s.Model.DisplayName
	m.CompartmentId = s.Model.CompartmentId
	m.VaultId = s.Model.VaultId
	m.KeyId = s.Model.KeyId
	m.Username = s.Model.Username
	m.Password = s.Model.Password
	m.Description = s.Model.Description
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SubnetId = s.Model.SubnetId
	m.NsgIds = s.Model.NsgIds
	m.ReplicationUsername = s.Model.ReplicationUsername
	m.ReplicationPassword = s.Model.ReplicationPassword
	m.ConnectionType = s.Model.ConnectionType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createconnectiondetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ConnectionType {
	case "MYSQL":
		mm := CreateMysqlConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE":
		mm := CreateOracleConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for CreateConnectionDetails: %s.", m.ConnectionType)
		return *m, nil
	}
}

// GetDescription returns Description
func (m createconnectiondetails) GetDescription() *string {
	return m.Description
}

// GetFreeformTags returns FreeformTags
func (m createconnectiondetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m createconnectiondetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSubnetId returns SubnetId
func (m createconnectiondetails) GetSubnetId() *string {
	return m.SubnetId
}

// GetNsgIds returns NsgIds
func (m createconnectiondetails) GetNsgIds() []string {
	return m.NsgIds
}

// GetReplicationUsername returns ReplicationUsername
func (m createconnectiondetails) GetReplicationUsername() *string {
	return m.ReplicationUsername
}

// GetReplicationPassword returns ReplicationPassword
func (m createconnectiondetails) GetReplicationPassword() *string {
	return m.ReplicationPassword
}

// GetDisplayName returns DisplayName
func (m createconnectiondetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetCompartmentId returns CompartmentId
func (m createconnectiondetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetVaultId returns VaultId
func (m createconnectiondetails) GetVaultId() *string {
	return m.VaultId
}

// GetKeyId returns KeyId
func (m createconnectiondetails) GetKeyId() *string {
	return m.KeyId
}

// GetUsername returns Username
func (m createconnectiondetails) GetUsername() *string {
	return m.Username
}

// GetPassword returns Password
func (m createconnectiondetails) GetPassword() *string {
	return m.Password
}

func (m createconnectiondetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createconnectiondetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
