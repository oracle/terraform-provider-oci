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

// CreateRedisConnectionDetails The information about a new Redis Connection.
type CreateRedisConnectionDetails struct {

	// An object's Display Name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment being referenced.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

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

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
	SubscriptionId *string `mandatory:"false" json:"subscriptionId"`

	// The OCID(/Content/General/Concepts/identifiers.htm) of the cluster placement group for the resource.
	// Only applicable for multicloud subscriptions. The cluster placement group id must be provided when a multicloud
	// subscription id is provided. Otherwise the cluster placement group must not be provided.
	ClusterPlacementGroupId *string `mandatory:"false" json:"clusterPlacementGroupId"`

	// Security attributes for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Oracle-ZPR": {"MaxEgressCount": {"value": "42", "mode": "enforce"}}}`
	SecurityAttributes map[string]map[string]interface{} `mandatory:"false" json:"securityAttributes"`

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
	// Deprecated: This field is deprecated and replaced by "passwordSecretId". This field will be removed after February 15 2026.
	Password *string `mandatory:"false" json:"password"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the password is stored.
	// The password Oracle GoldenGate uses to connect the associated system of the given technology.
	// It must conform to the specific security requirements including length, case sensitivity, and so on.
	// If secretId is used plaintext field must not be provided.
	// Note: When provided, 'password' field must not be provided.
	PasswordSecretId *string `mandatory:"false" json:"passwordSecretId"`

	// The base64 encoded content of the TrustStore file.
	// Deprecated: This field is deprecated and replaced by "trustStoreSecretId". This field will be removed after February 15 2026.
	TrustStore *string `mandatory:"false" json:"trustStore"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the content of the TrustStore file.
	// Note: When provided, 'trustStore' field must not be provided.
	TrustStoreSecretId *string `mandatory:"false" json:"trustStoreSecretId"`

	// The TrustStore password.
	// Deprecated: This field is deprecated and replaced by "trustStorePasswordSecretId". This field will be removed after February 15 2026.
	TrustStorePassword *string `mandatory:"false" json:"trustStorePassword"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the Redis TrustStore password is stored.
	// Note: When provided, 'trustStorePassword' field must not be provided.
	TrustStorePasswordSecretId *string `mandatory:"false" json:"trustStorePasswordSecretId"`

	// The base64 encoded content of the KeyStore file.
	// Deprecated: This field is deprecated and replaced by "keyStoreSecretId". This field will be removed after February 15 2026.
	KeyStore *string `mandatory:"false" json:"keyStore"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the content of the KeyStore file.
	// Note: When provided, 'keyStore' field must not be provided.
	KeyStoreSecretId *string `mandatory:"false" json:"keyStoreSecretId"`

	// The KeyStore password.
	// Deprecated: This field is deprecated and replaced by "keyStorePasswordSecretId". This field will be removed after February 15 2026.
	KeyStorePassword *string `mandatory:"false" json:"keyStorePassword"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the Redis KeyStore password is stored.
	// Note: When provided, 'keyStorePassword' field must not be provided.
	KeyStorePasswordSecretId *string `mandatory:"false" json:"keyStorePasswordSecretId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Redis cluster.
	RedisClusterId *string `mandatory:"false" json:"redisClusterId"`

	// Controls the network traffic direction to the target:
	// SHARED_SERVICE_ENDPOINT: Traffic flows through the Goldengate Service's network to public hosts. Cannot be used for private targets.
	// SHARED_DEPLOYMENT_ENDPOINT: Network traffic flows from the assigned deployment's private endpoint through the deployment's subnet.
	// DEDICATED_ENDPOINT: A dedicated private endpoint is created in the target VCN subnet for the connection. The subnetId is required when DEDICATED_ENDPOINT networking is selected.
	RoutingMethod RoutingMethodEnum `mandatory:"false" json:"routingMethod,omitempty"`

	// The Redis technology type.
	TechnologyType RedisConnectionTechnologyTypeEnum `mandatory:"true" json:"technologyType"`

	// Security protocol for Redis.
	SecurityProtocol RedisConnectionSecurityProtocolEnum `mandatory:"true" json:"securityProtocol"`

	// Authenticationentication type for the Redis database.
	AuthenticationType RedisConnectionAuthenticationTypeEnum `mandatory:"true" json:"authenticationType"`
}

// GetDisplayName returns DisplayName
func (m CreateRedisConnectionDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m CreateRedisConnectionDetails) GetDescription() *string {
	return m.Description
}

// GetCompartmentId returns CompartmentId
func (m CreateRedisConnectionDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetFreeformTags returns FreeformTags
func (m CreateRedisConnectionDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateRedisConnectionDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetLocks returns Locks
func (m CreateRedisConnectionDetails) GetLocks() []AddResourceLockDetails {
	return m.Locks
}

// GetVaultId returns VaultId
func (m CreateRedisConnectionDetails) GetVaultId() *string {
	return m.VaultId
}

// GetKeyId returns KeyId
func (m CreateRedisConnectionDetails) GetKeyId() *string {
	return m.KeyId
}

// GetNsgIds returns NsgIds
func (m CreateRedisConnectionDetails) GetNsgIds() []string {
	return m.NsgIds
}

// GetSubnetId returns SubnetId
func (m CreateRedisConnectionDetails) GetSubnetId() *string {
	return m.SubnetId
}

// GetRoutingMethod returns RoutingMethod
func (m CreateRedisConnectionDetails) GetRoutingMethod() RoutingMethodEnum {
	return m.RoutingMethod
}

// GetDoesUseSecretIds returns DoesUseSecretIds
func (m CreateRedisConnectionDetails) GetDoesUseSecretIds() *bool {
	return m.DoesUseSecretIds
}

// GetSubscriptionId returns SubscriptionId
func (m CreateRedisConnectionDetails) GetSubscriptionId() *string {
	return m.SubscriptionId
}

// GetClusterPlacementGroupId returns ClusterPlacementGroupId
func (m CreateRedisConnectionDetails) GetClusterPlacementGroupId() *string {
	return m.ClusterPlacementGroupId
}

// GetSecurityAttributes returns SecurityAttributes
func (m CreateRedisConnectionDetails) GetSecurityAttributes() map[string]map[string]interface{} {
	return m.SecurityAttributes
}

func (m CreateRedisConnectionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateRedisConnectionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingRoutingMethodEnum(string(m.RoutingMethod)); !ok && m.RoutingMethod != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RoutingMethod: %s. Supported values are: %s.", m.RoutingMethod, strings.Join(GetRoutingMethodEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRedisConnectionTechnologyTypeEnum(string(m.TechnologyType)); !ok && m.TechnologyType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TechnologyType: %s. Supported values are: %s.", m.TechnologyType, strings.Join(GetRedisConnectionTechnologyTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRedisConnectionSecurityProtocolEnum(string(m.SecurityProtocol)); !ok && m.SecurityProtocol != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SecurityProtocol: %s. Supported values are: %s.", m.SecurityProtocol, strings.Join(GetRedisConnectionSecurityProtocolEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRedisConnectionAuthenticationTypeEnum(string(m.AuthenticationType)); !ok && m.AuthenticationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AuthenticationType: %s. Supported values are: %s.", m.AuthenticationType, strings.Join(GetRedisConnectionAuthenticationTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateRedisConnectionDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateRedisConnectionDetails CreateRedisConnectionDetails
	s := struct {
		DiscriminatorParam string `json:"connectionType"`
		MarshalTypeCreateRedisConnectionDetails
	}{
		"REDIS",
		(MarshalTypeCreateRedisConnectionDetails)(m),
	}

	return json.Marshal(&s)
}
