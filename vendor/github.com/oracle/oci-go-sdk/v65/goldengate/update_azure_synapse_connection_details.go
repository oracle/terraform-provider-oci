// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

// UpdateAzureSynapseConnectionDetails The information to update a Azure Synapse Analytics Connection.
type UpdateAzureSynapseConnectionDetails struct {

	// An object's Display Name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Metadata about this specific object.
	Description *string `mandatory:"false" json:"description"`

	// A simple key-value pair that is applied without any predefined name, type, or scope. Exists
	// for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Tags defined for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the customer vault being
	// referenced.
	// If provided, this will reference a vault which the customer will be required to ensure
	// the policies are established to permit the GoldenGate Service to manage secrets contained
	// within this vault.
	VaultId *string `mandatory:"false" json:"vaultId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the customer "Master" key being
	// referenced.
	// If provided, this will reference a key which the customer will be required to ensure
	// the policies are established to permit the GoldenGate Service to utilize this key to
	// manage secrets.
	KeyId *string `mandatory:"false" json:"keyId"`

	// An array of Network Security Group OCIDs used to define network access for either Deployments or Connections.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// JDBC connection string.
	// e.g.: 'jdbc:sqlserver://<synapse-workspace>.sql.azuresynapse.net:1433;database=<db-name>;encrypt=true;trustServerCertificate=false;hostNameInCertificate=*.sql.azuresynapse.net;loginTimeout=300;'
	ConnectionString *string `mandatory:"false" json:"connectionString"`

	// The username Oracle GoldenGate uses to connect the associated RDBMS.  This username must
	// already exist and be available for use by the database.  It must conform to the security
	// requirements implemented by the database including length, case sensitivity, and so on.
	Username *string `mandatory:"false" json:"username"`

	// The password Oracle GoldenGate uses to connect the associated RDBMS.  It must conform to the
	// specific security requirements implemented by the database including length, case
	// sensitivity, and so on.
	Password *string `mandatory:"false" json:"password"`
}

//GetDisplayName returns DisplayName
func (m UpdateAzureSynapseConnectionDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetDescription returns Description
func (m UpdateAzureSynapseConnectionDetails) GetDescription() *string {
	return m.Description
}

//GetFreeformTags returns FreeformTags
func (m UpdateAzureSynapseConnectionDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m UpdateAzureSynapseConnectionDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetVaultId returns VaultId
func (m UpdateAzureSynapseConnectionDetails) GetVaultId() *string {
	return m.VaultId
}

//GetKeyId returns KeyId
func (m UpdateAzureSynapseConnectionDetails) GetKeyId() *string {
	return m.KeyId
}

//GetNsgIds returns NsgIds
func (m UpdateAzureSynapseConnectionDetails) GetNsgIds() []string {
	return m.NsgIds
}

func (m UpdateAzureSynapseConnectionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateAzureSynapseConnectionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateAzureSynapseConnectionDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateAzureSynapseConnectionDetails UpdateAzureSynapseConnectionDetails
	s := struct {
		DiscriminatorParam string `json:"connectionType"`
		MarshalTypeUpdateAzureSynapseConnectionDetails
	}{
		"AZURE_SYNAPSE_ANALYTICS",
		(MarshalTypeUpdateAzureSynapseConnectionDetails)(m),
	}

	return json.Marshal(&s)
}
