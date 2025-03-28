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

// CreateMicrosoftSqlserverConnectionDetails The information about a new Microsoft SQL Server Connection.
type CreateMicrosoftSqlserverConnectionDetails struct {

	// An object's Display Name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment being referenced.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name of the database.
	DatabaseName *string `mandatory:"true" json:"databaseName"`

	// The name or address of a host.
	Host *string `mandatory:"true" json:"host"`

	// The port of an endpoint usually specified for a connection.
	Port *int `mandatory:"true" json:"port"`

	// The username Oracle GoldenGate uses to connect to the Microsoft SQL Server.
	// This username must already exist and be available by the Microsoft SQL Server to be connected to.
	Username *string `mandatory:"true" json:"username"`

	// Metadata about this specific object.
	Description *string `mandatory:"false" json:"description"`

	// A simple key-value pair that is applied without any predefined name, type, or scope. Exists
	// for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Tags defined for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Locks associated with this resource.
	Locks []AddResourceLockDetails `mandatory:"false" json:"locks"`

	// Refers to the customer's vault OCID.
	// If provided, it references a vault where GoldenGate can manage secrets. Customers must add policies to permit GoldenGate
	// to manage secrets contained within this vault.
	VaultId *string `mandatory:"false" json:"vaultId"`

	// Refers to the customer's master key OCID.
	// If provided, it references a key to manage secrets. Customers must add policies to permit GoldenGate to use this key.
	KeyId *string `mandatory:"false" json:"keyId"`

	// An array of Network Security Group OCIDs used to define network access for either Deployments or Connections.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the target subnet of the dedicated connection.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// Indicates that sensitive attributes are provided via Secrets.
	DoesUseSecretIds *bool `mandatory:"false" json:"doesUseSecretIds"`

	// The password Oracle GoldenGate uses to connect the associated Microsoft SQL Server.
	// Deprecated: This field is deprecated and replaced by "passwordSecretId". This field will be removed after February 15 2026.
	Password *string `mandatory:"false" json:"password"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the password Oracle GoldenGate uses to connect the associated Microsoft SQL Server.
	// Note: When provided, 'password' field must not be provided.
	PasswordSecretId *string `mandatory:"false" json:"passwordSecretId"`

	// An array of name-value pair attribute entries.
	// Used as additional parameters in connection string.
	AdditionalAttributes []NameValuePair `mandatory:"false" json:"additionalAttributes"`

	// Database Certificate - The base64 encoded content of a .pem or .crt file.
	// containing the server public key (for 1-way SSL).
	SslCa *string `mandatory:"false" json:"sslCa"`

	// If set to true, the driver validates the certificate that is sent by the database server.
	ShouldValidateServerCertificate *bool `mandatory:"false" json:"shouldValidateServerCertificate"`

	// Deprecated: this field will be removed in future versions. Either specify the private IP in the connectionString or host
	// field, or make sure the host name is resolvable in the target VCN.
	// The private IP address of the connection's endpoint in the customer's VCN, typically a
	// database endpoint or a big data endpoint (e.g. Kafka bootstrap server).
	// In case the privateIp is provided, the subnetId must also be provided.
	// In case the privateIp (and the subnetId) is not provided it is assumed the datasource is publicly accessible.
	// In case the connection is accessible only privately, the lack of privateIp will result in not being able to access the connection.
	PrivateIp *string `mandatory:"false" json:"privateIp"`

	// Controls the network traffic direction to the target:
	// SHARED_SERVICE_ENDPOINT: Traffic flows through the Goldengate Service's network to public hosts. Cannot be used for private targets.
	// SHARED_DEPLOYMENT_ENDPOINT: Network traffic flows from the assigned deployment's private endpoint through the deployment's subnet.
	// DEDICATED_ENDPOINT: A dedicated private endpoint is created in the target VCN subnet for the connection. The subnetId is required when DEDICATED_ENDPOINT networking is selected.
	RoutingMethod RoutingMethodEnum `mandatory:"false" json:"routingMethod,omitempty"`

	// The Microsoft SQL Server technology type.
	TechnologyType MicrosoftSqlserverConnectionTechnologyTypeEnum `mandatory:"true" json:"technologyType"`

	// Security Type for Microsoft SQL Server.
	SecurityProtocol MicrosoftSqlserverConnectionSecurityProtocolEnum `mandatory:"true" json:"securityProtocol"`
}

// GetDisplayName returns DisplayName
func (m CreateMicrosoftSqlserverConnectionDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m CreateMicrosoftSqlserverConnectionDetails) GetDescription() *string {
	return m.Description
}

// GetCompartmentId returns CompartmentId
func (m CreateMicrosoftSqlserverConnectionDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetFreeformTags returns FreeformTags
func (m CreateMicrosoftSqlserverConnectionDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateMicrosoftSqlserverConnectionDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetLocks returns Locks
func (m CreateMicrosoftSqlserverConnectionDetails) GetLocks() []AddResourceLockDetails {
	return m.Locks
}

// GetVaultId returns VaultId
func (m CreateMicrosoftSqlserverConnectionDetails) GetVaultId() *string {
	return m.VaultId
}

// GetKeyId returns KeyId
func (m CreateMicrosoftSqlserverConnectionDetails) GetKeyId() *string {
	return m.KeyId
}

// GetNsgIds returns NsgIds
func (m CreateMicrosoftSqlserverConnectionDetails) GetNsgIds() []string {
	return m.NsgIds
}

// GetSubnetId returns SubnetId
func (m CreateMicrosoftSqlserverConnectionDetails) GetSubnetId() *string {
	return m.SubnetId
}

// GetRoutingMethod returns RoutingMethod
func (m CreateMicrosoftSqlserverConnectionDetails) GetRoutingMethod() RoutingMethodEnum {
	return m.RoutingMethod
}

// GetDoesUseSecretIds returns DoesUseSecretIds
func (m CreateMicrosoftSqlserverConnectionDetails) GetDoesUseSecretIds() *bool {
	return m.DoesUseSecretIds
}

func (m CreateMicrosoftSqlserverConnectionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateMicrosoftSqlserverConnectionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingRoutingMethodEnum(string(m.RoutingMethod)); !ok && m.RoutingMethod != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RoutingMethod: %s. Supported values are: %s.", m.RoutingMethod, strings.Join(GetRoutingMethodEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMicrosoftSqlserverConnectionTechnologyTypeEnum(string(m.TechnologyType)); !ok && m.TechnologyType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TechnologyType: %s. Supported values are: %s.", m.TechnologyType, strings.Join(GetMicrosoftSqlserverConnectionTechnologyTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMicrosoftSqlserverConnectionSecurityProtocolEnum(string(m.SecurityProtocol)); !ok && m.SecurityProtocol != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SecurityProtocol: %s. Supported values are: %s.", m.SecurityProtocol, strings.Join(GetMicrosoftSqlserverConnectionSecurityProtocolEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateMicrosoftSqlserverConnectionDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateMicrosoftSqlserverConnectionDetails CreateMicrosoftSqlserverConnectionDetails
	s := struct {
		DiscriminatorParam string `json:"connectionType"`
		MarshalTypeCreateMicrosoftSqlserverConnectionDetails
	}{
		"MICROSOFT_SQLSERVER",
		(MarshalTypeCreateMicrosoftSqlserverConnectionDetails)(m),
	}

	return json.Marshal(&s)
}
