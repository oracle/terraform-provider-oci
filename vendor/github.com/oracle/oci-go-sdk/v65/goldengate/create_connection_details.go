// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateConnectionDetails The information about a new Connection.
type CreateConnectionDetails interface {

	// An object's Display Name.
	GetDisplayName() *string

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment being referenced.
	GetCompartmentId() *string

	// Metadata about this specific object.
	GetDescription() *string

	// A simple key-value pair that is applied without any predefined name, type, or scope. Exists
	// for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Tags defined for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the customer vault being
	// referenced.
	// If provided, this will reference a vault which the customer will be required to ensure
	// the policies are established to permit the GoldenGate Service to manage secrets contained
	// within this vault.
	GetVaultId() *string

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the customer "Master" key being
	// referenced.
	// If provided, this will reference a key which the customer will be required to ensure
	// the policies are established to permit the GoldenGate Service to utilize this key to
	// manage secrets.
	GetKeyId() *string

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the subnet being referenced.
	GetSubnetId() *string

	// An array of Network Security Group OCIDs used to define network access for either Deployments or Connections.
	GetNsgIds() []string
}

type createconnectiondetails struct {
	JsonData       []byte
	DisplayName    *string                           `mandatory:"true" json:"displayName"`
	CompartmentId  *string                           `mandatory:"true" json:"compartmentId"`
	Description    *string                           `mandatory:"false" json:"description"`
	FreeformTags   map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags    map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	VaultId        *string                           `mandatory:"false" json:"vaultId"`
	KeyId          *string                           `mandatory:"false" json:"keyId"`
	SubnetId       *string                           `mandatory:"false" json:"subnetId"`
	NsgIds         []string                          `mandatory:"false" json:"nsgIds"`
	ConnectionType string                            `json:"connectionType"`
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
	m.Description = s.Model.Description
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.VaultId = s.Model.VaultId
	m.KeyId = s.Model.KeyId
	m.SubnetId = s.Model.SubnetId
	m.NsgIds = s.Model.NsgIds
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
	case "POSTGRESQL":
		mm := CreatePostgresqlConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "KAFKA_SCHEMA_REGISTRY":
		mm := CreateKafkaSchemaRegistryConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AZURE_SYNAPSE_ANALYTICS":
		mm := CreateAzureSynapseConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AZURE_DATA_LAKE_STORAGE":
		mm := CreateAzureDataLakeStorageConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MYSQL":
		mm := CreateMysqlConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_OBJECT_STORAGE":
		mm := CreateOciObjectStorageConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "KAFKA":
		mm := CreateKafkaConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE":
		mm := CreateOracleConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GOLDENGATE":
		mm := CreateGoldenGateConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for CreateConnectionDetails: %s.", m.ConnectionType)
		return *m, nil
	}
}

//GetDisplayName returns DisplayName
func (m createconnectiondetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetCompartmentId returns CompartmentId
func (m createconnectiondetails) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetDescription returns Description
func (m createconnectiondetails) GetDescription() *string {
	return m.Description
}

//GetFreeformTags returns FreeformTags
func (m createconnectiondetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m createconnectiondetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetVaultId returns VaultId
func (m createconnectiondetails) GetVaultId() *string {
	return m.VaultId
}

//GetKeyId returns KeyId
func (m createconnectiondetails) GetKeyId() *string {
	return m.KeyId
}

//GetSubnetId returns SubnetId
func (m createconnectiondetails) GetSubnetId() *string {
	return m.SubnetId
}

//GetNsgIds returns NsgIds
func (m createconnectiondetails) GetNsgIds() []string {
	return m.NsgIds
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
