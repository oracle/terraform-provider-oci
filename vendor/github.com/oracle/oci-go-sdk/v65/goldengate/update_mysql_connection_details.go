// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// UpdateMysqlConnectionDetails The information to update a MySQL Connection.
type UpdateMysqlConnectionDetails struct {

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

	// Refers to the customer's vault OCID.
	// If provided, it references a vault where GoldenGate can manage secrets. Customers must add policies to permit GoldenGate
	// to manage secrets contained within this vault.
	VaultId *string `mandatory:"false" json:"vaultId"`

	// Refers to the customer's master key OCID.
	// If provided, it references a key to manage secrets. Customers must add policies to permit GoldenGate to use this key.
	KeyId *string `mandatory:"false" json:"keyId"`

	// An array of Network Security Group OCIDs used to define network access for either Deployments or Connections.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the target subnet of the dedicated connection.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// Indicates that sensitive attributes are provided via Secrets.
	DoesUseSecretIds *bool `mandatory:"false" json:"doesUseSecretIds"`

	// The username Oracle GoldenGate uses to connect the associated system of the given technology.
	// This username must already exist and be available by the system/application to be connected to
	// and must conform to the case sensitivty requirments defined in it.
	Username *string `mandatory:"false" json:"username"`

	// The password Oracle GoldenGate uses to connect the associated system of the given technology.
	// It must conform to the specific security requirements including length, case sensitivity, and so on.
	// Deprecated: This field is deprecated and replaced by "passwordSecretId". This field will be removed after February 15 2026.
	Password *string `mandatory:"false" json:"password"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Secret where the password is stored.
	// The password Oracle GoldenGate uses to connect the associated system of the given technology.
	// It must conform to the specific security requirements including length, case sensitivity, and so on.
	// If secretId is used plaintext field must not be provided.
	// Note: When provided, 'password' field must not be provided.
	PasswordSecretId *string `mandatory:"false" json:"passwordSecretId"`

	// The name or address of a host.
	Host *string `mandatory:"false" json:"host"`

	// The port of an endpoint usually specified for a connection.
	Port *int `mandatory:"false" json:"port"`

	// The name of the database.
	DatabaseName *string `mandatory:"false" json:"databaseName"`

	// Database Certificate - The base64 encoded content of a .pem or .crt file.
	// containing the server public key (for 1 and 2-way SSL).
	SslCa *string `mandatory:"false" json:"sslCa"`

	// The base64 encoded list of certificates revoked by the trusted certificate authorities (Trusted CA).
	// Note: This is an optional property and only applicable if TLS/MTLS option is selected.
	SslCrl *string `mandatory:"false" json:"sslCrl"`

	// Client Certificate - The base64 encoded content of a .pem or .crt file.
	// containing the client public key (for 2-way SSL).
	SslCert *string `mandatory:"false" json:"sslCert"`

	// Client Key - The base64 encoded content of a .pem or .crt file containing the client private key (for 2-way SSL).
	// Deprecated: This field is deprecated and replaced by "sslKeySecretId". This field will be removed after February 15 2026.
	SslKey *string `mandatory:"false" json:"sslKey"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Secret that stores the Client Key
	// - The content of a .pem or .crt file containing the client private key (for 2-way SSL).
	// Note: When provided, 'sslKey' field must not be provided.
	SslKeySecretId *string `mandatory:"false" json:"sslKeySecretId"`

	// Deprecated: this field will be removed in future versions. Either specify the private IP in the connectionString or host
	// field, or make sure the host name is resolvable in the target VCN.
	// The private IP address of the connection's endpoint in the customer's VCN, typically a
	// database endpoint or a big data endpoint (e.g. Kafka bootstrap server).
	// In case the privateIp is provided, the subnetId must also be provided.
	// In case the privateIp (and the subnetId) is not provided it is assumed the datasource is publicly accessible.
	// In case the connection is accessible only privately, the lack of privateIp will result in not being able to access the connection.
	PrivateIp *string `mandatory:"false" json:"privateIp"`

	// An array of name-value pair attribute entries.
	// Used as additional parameters in connection string.
	AdditionalAttributes []NameValuePair `mandatory:"false" json:"additionalAttributes"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the database system being referenced.
	DbSystemId *string `mandatory:"false" json:"dbSystemId"`

	// Controls the network traffic direction to the target:
	// SHARED_SERVICE_ENDPOINT: Traffic flows through the Goldengate Service's network to public hosts. Cannot be used for private targets.
	// SHARED_DEPLOYMENT_ENDPOINT: Network traffic flows from the assigned deployment's private endpoint through the deployment's subnet.
	// DEDICATED_ENDPOINT: A dedicated private endpoint is created in the target VCN subnet for the connection. The subnetId is required when DEDICATED_ENDPOINT networking is selected.
	RoutingMethod RoutingMethodEnum `mandatory:"false" json:"routingMethod,omitempty"`

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

// GetNsgIds returns NsgIds
func (m UpdateMysqlConnectionDetails) GetNsgIds() []string {
	return m.NsgIds
}

// GetSubnetId returns SubnetId
func (m UpdateMysqlConnectionDetails) GetSubnetId() *string {
	return m.SubnetId
}

// GetRoutingMethod returns RoutingMethod
func (m UpdateMysqlConnectionDetails) GetRoutingMethod() RoutingMethodEnum {
	return m.RoutingMethod
}

// GetDoesUseSecretIds returns DoesUseSecretIds
func (m UpdateMysqlConnectionDetails) GetDoesUseSecretIds() *bool {
	return m.DoesUseSecretIds
}

func (m UpdateMysqlConnectionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateMysqlConnectionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingRoutingMethodEnum(string(m.RoutingMethod)); !ok && m.RoutingMethod != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RoutingMethod: %s. Supported values are: %s.", m.RoutingMethod, strings.Join(GetRoutingMethodEnumStringValues(), ",")))
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
