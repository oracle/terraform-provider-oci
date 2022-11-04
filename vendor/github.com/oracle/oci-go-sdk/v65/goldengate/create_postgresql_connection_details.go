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

// CreatePostgresqlConnectionDetails The information about a new PostgreSQL Connection.
type CreatePostgresqlConnectionDetails struct {

	// An object's Display Name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment being referenced.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name of the database.
	DatabaseName *string `mandatory:"true" json:"databaseName"`

	// The name or address of a host.
	Host *string `mandatory:"true" json:"host"`

	// The port of an endpoint usually specified for a connection.
	Port *int `mandatory:"true" json:"port"`

	// The username Oracle GoldenGate uses to connect the associated RDBMS.  This username must
	// already exist and be available for use by the database.  It must conform to the security
	// requirements implemented by the database including length, case sensitivity, and so on.
	Username *string `mandatory:"true" json:"username"`

	// The password Oracle GoldenGate uses to connect the associated RDBMS.  It must conform to the
	// specific security requirements implemented by the database including length, case
	// sensitivity, and so on.
	Password *string `mandatory:"true" json:"password"`

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

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the subnet being referenced.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// An array of Network Security Group OCIDs used to define network access for either Deployments or Connections.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// An array of name-value pair attribute entries.
	// Used as additional parameters in connection string.
	AdditionalAttributes []NameValuePair `mandatory:"false" json:"additionalAttributes"`

	// The base64 encoded certificate of the trusted certificate authorities (Trusted CA) for PostgreSQL.
	SslCa *string `mandatory:"false" json:"sslCa"`

	// The base64 encoded list of certificates revoked by the trusted certificate authorities (Trusted CA) for PostgreSQL.
	SslCrl *string `mandatory:"false" json:"sslCrl"`

	// The base64 encoded certificate of the PostgreSQL server.
	SslCert *string `mandatory:"false" json:"sslCert"`

	// The base64 encoded private key of the PostgreSQL server.
	SslKey *string `mandatory:"false" json:"sslKey"`

	// The private IP address of the connection's endpoint in the customer's VCN, typically a
	// database endpoint or a big data endpoint (e.g. Kafka bootstrap server).
	// In case the privateIp is provided, the subnetId must also be provided.
	// In case the privateIp (and the subnetId) is not provided it is assumed the datasource is publicly accessible.
	// In case the connection is accessible only privately, the lack of privateIp will result in not being able to access the connection.
	PrivateIp *string `mandatory:"false" json:"privateIp"`

	// The PostgreSQL technology type.
	TechnologyType PostgresqlConnectionTechnologyTypeEnum `mandatory:"true" json:"technologyType"`

	// Security protocol for PostgreSQL.
	SecurityProtocol PostgresqlConnectionSecurityProtocolEnum `mandatory:"true" json:"securityProtocol"`

	// SSL modes for PostgreSQL.
	SslMode PostgresqlConnectionSslModeEnum `mandatory:"false" json:"sslMode,omitempty"`
}

//GetDisplayName returns DisplayName
func (m CreatePostgresqlConnectionDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetDescription returns Description
func (m CreatePostgresqlConnectionDetails) GetDescription() *string {
	return m.Description
}

//GetCompartmentId returns CompartmentId
func (m CreatePostgresqlConnectionDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetFreeformTags returns FreeformTags
func (m CreatePostgresqlConnectionDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m CreatePostgresqlConnectionDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetVaultId returns VaultId
func (m CreatePostgresqlConnectionDetails) GetVaultId() *string {
	return m.VaultId
}

//GetKeyId returns KeyId
func (m CreatePostgresqlConnectionDetails) GetKeyId() *string {
	return m.KeyId
}

//GetSubnetId returns SubnetId
func (m CreatePostgresqlConnectionDetails) GetSubnetId() *string {
	return m.SubnetId
}

//GetNsgIds returns NsgIds
func (m CreatePostgresqlConnectionDetails) GetNsgIds() []string {
	return m.NsgIds
}

func (m CreatePostgresqlConnectionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreatePostgresqlConnectionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingPostgresqlConnectionTechnologyTypeEnum(string(m.TechnologyType)); !ok && m.TechnologyType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TechnologyType: %s. Supported values are: %s.", m.TechnologyType, strings.Join(GetPostgresqlConnectionTechnologyTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPostgresqlConnectionSecurityProtocolEnum(string(m.SecurityProtocol)); !ok && m.SecurityProtocol != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SecurityProtocol: %s. Supported values are: %s.", m.SecurityProtocol, strings.Join(GetPostgresqlConnectionSecurityProtocolEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPostgresqlConnectionSslModeEnum(string(m.SslMode)); !ok && m.SslMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SslMode: %s. Supported values are: %s.", m.SslMode, strings.Join(GetPostgresqlConnectionSslModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreatePostgresqlConnectionDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreatePostgresqlConnectionDetails CreatePostgresqlConnectionDetails
	s := struct {
		DiscriminatorParam string `json:"connectionType"`
		MarshalTypeCreatePostgresqlConnectionDetails
	}{
		"POSTGRESQL",
		(MarshalTypeCreatePostgresqlConnectionDetails)(m),
	}

	return json.Marshal(&s)
}
