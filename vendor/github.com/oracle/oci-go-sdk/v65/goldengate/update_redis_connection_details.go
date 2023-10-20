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

// UpdateRedisConnectionDetails The information to update a Redis Database Connection.
type UpdateRedisConnectionDetails struct {

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

	// Comma separated list of Redis server addresses, specified as host:port entries, where :port is optional.
	// If port is not specified, it defaults to 6379.
	// Used for establishing the initial connection to the Redis cluster.
	// Example: `"server1.example.com:6379,server2.example.com:6379"`
	Servers *string `mandatory:"false" json:"servers"`

	// The username Oracle GoldenGate uses to connect the associated system of the given technology.
	// This username must already exist and be available by the system/application to be connected to
	// and must conform to the case sensitivty requirments defined in it.
	Username *string `mandatory:"false" json:"username"`

	// The password Oracle GoldenGate uses to connect the associated system of the given technology.
	// It must conform to the specific security requirements including length, case sensitivity, and so on.
	Password *string `mandatory:"false" json:"password"`

	// The base64 encoded content of the TrustStore file.
	TrustStore *string `mandatory:"false" json:"trustStore"`

	// The TrustStore password.
	TrustStorePassword *string `mandatory:"false" json:"trustStorePassword"`

	// The base64 encoded content of the KeyStore file.
	KeyStore *string `mandatory:"false" json:"keyStore"`

	// The KeyStore password.
	KeyStorePassword *string `mandatory:"false" json:"keyStorePassword"`

	// Security protocol for Redis.
	SecurityProtocol RedisConnectionSecurityProtocolEnum `mandatory:"false" json:"securityProtocol,omitempty"`

	// Authenticationentication type for the Redis database.
	AuthenticationType RedisConnectionAuthenticationTypeEnum `mandatory:"false" json:"authenticationType,omitempty"`
}

// GetDisplayName returns DisplayName
func (m UpdateRedisConnectionDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m UpdateRedisConnectionDetails) GetDescription() *string {
	return m.Description
}

// GetFreeformTags returns FreeformTags
func (m UpdateRedisConnectionDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m UpdateRedisConnectionDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetVaultId returns VaultId
func (m UpdateRedisConnectionDetails) GetVaultId() *string {
	return m.VaultId
}

// GetKeyId returns KeyId
func (m UpdateRedisConnectionDetails) GetKeyId() *string {
	return m.KeyId
}

// GetNsgIds returns NsgIds
func (m UpdateRedisConnectionDetails) GetNsgIds() []string {
	return m.NsgIds
}

func (m UpdateRedisConnectionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateRedisConnectionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingRedisConnectionSecurityProtocolEnum(string(m.SecurityProtocol)); !ok && m.SecurityProtocol != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SecurityProtocol: %s. Supported values are: %s.", m.SecurityProtocol, strings.Join(GetRedisConnectionSecurityProtocolEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRedisConnectionAuthenticationTypeEnum(string(m.AuthenticationType)); !ok && m.AuthenticationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AuthenticationType: %s. Supported values are: %s.", m.AuthenticationType, strings.Join(GetRedisConnectionAuthenticationTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateRedisConnectionDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateRedisConnectionDetails UpdateRedisConnectionDetails
	s := struct {
		DiscriminatorParam string `json:"connectionType"`
		MarshalTypeUpdateRedisConnectionDetails
	}{
		"REDIS",
		(MarshalTypeUpdateRedisConnectionDetails)(m),
	}

	return json.Marshal(&s)
}
