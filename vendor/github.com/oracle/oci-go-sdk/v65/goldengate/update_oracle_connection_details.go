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

// UpdateOracleConnectionDetails The information to update an Oracle Database Connection.
type UpdateOracleConnectionDetails struct {

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

	// The username Oracle GoldenGate uses to connect the associated RDBMS.  This username must
	// already exist and be available for use by the database.  It must conform to the security
	// requirements implemented by the database including length, case sensitivity, and so on.
	Username *string `mandatory:"false" json:"username"`

	// The password Oracle GoldenGate uses to connect the associated RDBMS.  It must conform to the
	// specific security requirements implemented by the database including length, case
	// sensitivity, and so on.
	Password *string `mandatory:"false" json:"password"`

	// Connect descriptor or Easy Connect Naming method that Oracle GoldenGate uses to connect to a
	// database.
	ConnectionString *string `mandatory:"false" json:"connectionString"`

	// The wallet contents Oracle GoldenGate uses to make connections to a database.  This
	// attribute is expected to be base64 encoded.
	Wallet *string `mandatory:"false" json:"wallet"`

	// The private IP address of the connection's endpoint in the customer's VCN, typically a
	// database endpoint or a big data endpoint (e.g. Kafka bootstrap server).
	// In case the privateIp is provided, the subnetId must also be provided.
	// In case the privateIp (and the subnetId) is not provided it is assumed the datasource is publicly accessible.
	// In case the connection is accessible only privately, the lack of privateIp will result in not being able to access the connection.
	PrivateIp *string `mandatory:"false" json:"privateIp"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the database being referenced.
	DatabaseId *string `mandatory:"false" json:"databaseId"`

	// The mode of the database connection session to be established by the data client.
	// 'REDIRECT' - for a RAC database, 'DIRECT' - for a non-RAC database.
	// Connection to a RAC database involves a redirection received from the SCAN listeners
	// to the database node to connect to. By default the mode would be DIRECT.
	SessionMode OracleConnectionSessionModeEnum `mandatory:"false" json:"sessionMode,omitempty"`
}

//GetDisplayName returns DisplayName
func (m UpdateOracleConnectionDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetDescription returns Description
func (m UpdateOracleConnectionDetails) GetDescription() *string {
	return m.Description
}

//GetFreeformTags returns FreeformTags
func (m UpdateOracleConnectionDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m UpdateOracleConnectionDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetVaultId returns VaultId
func (m UpdateOracleConnectionDetails) GetVaultId() *string {
	return m.VaultId
}

//GetKeyId returns KeyId
func (m UpdateOracleConnectionDetails) GetKeyId() *string {
	return m.KeyId
}

//GetNsgIds returns NsgIds
func (m UpdateOracleConnectionDetails) GetNsgIds() []string {
	return m.NsgIds
}

func (m UpdateOracleConnectionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateOracleConnectionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOracleConnectionSessionModeEnum(string(m.SessionMode)); !ok && m.SessionMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SessionMode: %s. Supported values are: %s.", m.SessionMode, strings.Join(GetOracleConnectionSessionModeEnumStringValues(), ",")))
	}
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
