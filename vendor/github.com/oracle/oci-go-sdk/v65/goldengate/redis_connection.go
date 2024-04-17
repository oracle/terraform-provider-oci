// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// RedisConnection Represents the metadata of a Redis Database Connection.
type RedisConnection struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the connection being
	// referenced.
	Id *string `mandatory:"true" json:"id"`

	// An object's Display Name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment being referenced.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The time the resource was created. The format is defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the resource was last updated. The format is defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Comma separated list of Redis server addresses, specified as host:port entries, where :port is optional.
	// If port is not specified, it defaults to 6379.
	// Used for establishing the initial connection to the Redis cluster.
	// Example: `"server1.example.com:6379,server2.example.com:6379"`
	Servers *string `mandatory:"true" json:"servers"`

	// Metadata about this specific object.
	Description *string `mandatory:"false" json:"description"`

	// A simple key-value pair that is applied without any predefined name, type, or scope. Exists
	// for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Tags defined for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The system tags associated with this resource, if any. The system tags are set by Oracle
	// Cloud Infrastructure services. Each key is predefined and scoped to namespaces.  For more
	// information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{orcl-cloud: {free-tier-retain: true}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// Describes the object's current state in detail. For example, it can be used to provide
	// actionable information for a resource in a Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Locks associated with this resource.
	Locks []ResourceLock `mandatory:"false" json:"locks"`

	// Refers to the customer's vault OCID.
	// If provided, it references a vault where GoldenGate can manage secrets. Customers must add policies to permit GoldenGate
	// to manage secrets contained within this vault.
	VaultId *string `mandatory:"false" json:"vaultId"`

	// Refers to the customer's master key OCID.
	// If provided, it references a key to manage secrets. Customers must add policies to permit GoldenGate to use this key.
	KeyId *string `mandatory:"false" json:"keyId"`

	// List of ingress IP addresses from where the GoldenGate deployment connects to this connection's privateIp.
	// Customers may optionally set up ingress security rules to restrict traffic from these IP addresses.
	IngressIps []IngressIpDetails `mandatory:"false" json:"ingressIps"`

	// An array of Network Security Group OCIDs used to define network access for either Deployments or Connections.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the target subnet of the dedicated connection.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// The username Oracle GoldenGate uses to connect the associated system of the given technology.
	// This username must already exist and be available by the system/application to be connected to
	// and must conform to the case sensitivty requirments defined in it.
	Username *string `mandatory:"false" json:"username"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Redis cluster.
	RedisClusterId *string `mandatory:"false" json:"redisClusterId"`

	// The Redis technology type.
	TechnologyType RedisConnectionTechnologyTypeEnum `mandatory:"true" json:"technologyType"`

	// Security protocol for Redis
	SecurityProtocol RedisConnectionSecurityProtocolEnum `mandatory:"true" json:"securityProtocol"`

	// Authentication type for Redis.
	AuthenticationType RedisConnectionAuthenticationTypeEnum `mandatory:"true" json:"authenticationType"`

	// Possible lifecycle states for connection.
	LifecycleState ConnectionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Controls the network traffic direction to the target:
	// SHARED_SERVICE_ENDPOINT: Traffic flows through the Goldengate Service's network to public hosts. Cannot be used for private targets.
	// SHARED_DEPLOYMENT_ENDPOINT: Network traffic flows from the assigned deployment's private endpoint through the deployment's subnet.
	// DEDICATED_ENDPOINT: A dedicated private endpoint is created in the target VCN subnet for the connection. The subnetId is required when DEDICATED_ENDPOINT networking is selected.
	RoutingMethod RoutingMethodEnum `mandatory:"false" json:"routingMethod,omitempty"`
}

// GetId returns Id
func (m RedisConnection) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m RedisConnection) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m RedisConnection) GetDescription() *string {
	return m.Description
}

// GetCompartmentId returns CompartmentId
func (m RedisConnection) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetFreeformTags returns FreeformTags
func (m RedisConnection) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m RedisConnection) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m RedisConnection) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetLifecycleState returns LifecycleState
func (m RedisConnection) GetLifecycleState() ConnectionLifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m RedisConnection) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetTimeCreated returns TimeCreated
func (m RedisConnection) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m RedisConnection) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLocks returns Locks
func (m RedisConnection) GetLocks() []ResourceLock {
	return m.Locks
}

// GetVaultId returns VaultId
func (m RedisConnection) GetVaultId() *string {
	return m.VaultId
}

// GetKeyId returns KeyId
func (m RedisConnection) GetKeyId() *string {
	return m.KeyId
}

// GetIngressIps returns IngressIps
func (m RedisConnection) GetIngressIps() []IngressIpDetails {
	return m.IngressIps
}

// GetNsgIds returns NsgIds
func (m RedisConnection) GetNsgIds() []string {
	return m.NsgIds
}

// GetSubnetId returns SubnetId
func (m RedisConnection) GetSubnetId() *string {
	return m.SubnetId
}

// GetRoutingMethod returns RoutingMethod
func (m RedisConnection) GetRoutingMethod() RoutingMethodEnum {
	return m.RoutingMethod
}

func (m RedisConnection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RedisConnection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRedisConnectionTechnologyTypeEnum(string(m.TechnologyType)); !ok && m.TechnologyType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TechnologyType: %s. Supported values are: %s.", m.TechnologyType, strings.Join(GetRedisConnectionTechnologyTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRedisConnectionSecurityProtocolEnum(string(m.SecurityProtocol)); !ok && m.SecurityProtocol != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SecurityProtocol: %s. Supported values are: %s.", m.SecurityProtocol, strings.Join(GetRedisConnectionSecurityProtocolEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRedisConnectionAuthenticationTypeEnum(string(m.AuthenticationType)); !ok && m.AuthenticationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AuthenticationType: %s. Supported values are: %s.", m.AuthenticationType, strings.Join(GetRedisConnectionAuthenticationTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingConnectionLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetConnectionLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRoutingMethodEnum(string(m.RoutingMethod)); !ok && m.RoutingMethod != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RoutingMethod: %s. Supported values are: %s.", m.RoutingMethod, strings.Join(GetRoutingMethodEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m RedisConnection) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeRedisConnection RedisConnection
	s := struct {
		DiscriminatorParam string `json:"connectionType"`
		MarshalTypeRedisConnection
	}{
		"REDIS",
		(MarshalTypeRedisConnection)(m),
	}

	return json.Marshal(&s)
}

// RedisConnectionTechnologyTypeEnum Enum with underlying type: string
type RedisConnectionTechnologyTypeEnum string

// Set of constants representing the allowable values for RedisConnectionTechnologyTypeEnum
const (
	RedisConnectionTechnologyTypeRedis             RedisConnectionTechnologyTypeEnum = "REDIS"
	RedisConnectionTechnologyTypeOciCacheWithRedis RedisConnectionTechnologyTypeEnum = "OCI_CACHE_WITH_REDIS"
)

var mappingRedisConnectionTechnologyTypeEnum = map[string]RedisConnectionTechnologyTypeEnum{
	"REDIS":                RedisConnectionTechnologyTypeRedis,
	"OCI_CACHE_WITH_REDIS": RedisConnectionTechnologyTypeOciCacheWithRedis,
}

var mappingRedisConnectionTechnologyTypeEnumLowerCase = map[string]RedisConnectionTechnologyTypeEnum{
	"redis":                RedisConnectionTechnologyTypeRedis,
	"oci_cache_with_redis": RedisConnectionTechnologyTypeOciCacheWithRedis,
}

// GetRedisConnectionTechnologyTypeEnumValues Enumerates the set of values for RedisConnectionTechnologyTypeEnum
func GetRedisConnectionTechnologyTypeEnumValues() []RedisConnectionTechnologyTypeEnum {
	values := make([]RedisConnectionTechnologyTypeEnum, 0)
	for _, v := range mappingRedisConnectionTechnologyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRedisConnectionTechnologyTypeEnumStringValues Enumerates the set of values in String for RedisConnectionTechnologyTypeEnum
func GetRedisConnectionTechnologyTypeEnumStringValues() []string {
	return []string{
		"REDIS",
		"OCI_CACHE_WITH_REDIS",
	}
}

// GetMappingRedisConnectionTechnologyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRedisConnectionTechnologyTypeEnum(val string) (RedisConnectionTechnologyTypeEnum, bool) {
	enum, ok := mappingRedisConnectionTechnologyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// RedisConnectionSecurityProtocolEnum Enum with underlying type: string
type RedisConnectionSecurityProtocolEnum string

// Set of constants representing the allowable values for RedisConnectionSecurityProtocolEnum
const (
	RedisConnectionSecurityProtocolPlain RedisConnectionSecurityProtocolEnum = "PLAIN"
	RedisConnectionSecurityProtocolTls   RedisConnectionSecurityProtocolEnum = "TLS"
	RedisConnectionSecurityProtocolMtls  RedisConnectionSecurityProtocolEnum = "MTLS"
)

var mappingRedisConnectionSecurityProtocolEnum = map[string]RedisConnectionSecurityProtocolEnum{
	"PLAIN": RedisConnectionSecurityProtocolPlain,
	"TLS":   RedisConnectionSecurityProtocolTls,
	"MTLS":  RedisConnectionSecurityProtocolMtls,
}

var mappingRedisConnectionSecurityProtocolEnumLowerCase = map[string]RedisConnectionSecurityProtocolEnum{
	"plain": RedisConnectionSecurityProtocolPlain,
	"tls":   RedisConnectionSecurityProtocolTls,
	"mtls":  RedisConnectionSecurityProtocolMtls,
}

// GetRedisConnectionSecurityProtocolEnumValues Enumerates the set of values for RedisConnectionSecurityProtocolEnum
func GetRedisConnectionSecurityProtocolEnumValues() []RedisConnectionSecurityProtocolEnum {
	values := make([]RedisConnectionSecurityProtocolEnum, 0)
	for _, v := range mappingRedisConnectionSecurityProtocolEnum {
		values = append(values, v)
	}
	return values
}

// GetRedisConnectionSecurityProtocolEnumStringValues Enumerates the set of values in String for RedisConnectionSecurityProtocolEnum
func GetRedisConnectionSecurityProtocolEnumStringValues() []string {
	return []string{
		"PLAIN",
		"TLS",
		"MTLS",
	}
}

// GetMappingRedisConnectionSecurityProtocolEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRedisConnectionSecurityProtocolEnum(val string) (RedisConnectionSecurityProtocolEnum, bool) {
	enum, ok := mappingRedisConnectionSecurityProtocolEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// RedisConnectionAuthenticationTypeEnum Enum with underlying type: string
type RedisConnectionAuthenticationTypeEnum string

// Set of constants representing the allowable values for RedisConnectionAuthenticationTypeEnum
const (
	RedisConnectionAuthenticationTypeNone  RedisConnectionAuthenticationTypeEnum = "NONE"
	RedisConnectionAuthenticationTypeBasic RedisConnectionAuthenticationTypeEnum = "BASIC"
)

var mappingRedisConnectionAuthenticationTypeEnum = map[string]RedisConnectionAuthenticationTypeEnum{
	"NONE":  RedisConnectionAuthenticationTypeNone,
	"BASIC": RedisConnectionAuthenticationTypeBasic,
}

var mappingRedisConnectionAuthenticationTypeEnumLowerCase = map[string]RedisConnectionAuthenticationTypeEnum{
	"none":  RedisConnectionAuthenticationTypeNone,
	"basic": RedisConnectionAuthenticationTypeBasic,
}

// GetRedisConnectionAuthenticationTypeEnumValues Enumerates the set of values for RedisConnectionAuthenticationTypeEnum
func GetRedisConnectionAuthenticationTypeEnumValues() []RedisConnectionAuthenticationTypeEnum {
	values := make([]RedisConnectionAuthenticationTypeEnum, 0)
	for _, v := range mappingRedisConnectionAuthenticationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRedisConnectionAuthenticationTypeEnumStringValues Enumerates the set of values in String for RedisConnectionAuthenticationTypeEnum
func GetRedisConnectionAuthenticationTypeEnumStringValues() []string {
	return []string{
		"NONE",
		"BASIC",
	}
}

// GetMappingRedisConnectionAuthenticationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRedisConnectionAuthenticationTypeEnum(val string) (RedisConnectionAuthenticationTypeEnum, bool) {
	enum, ok := mappingRedisConnectionAuthenticationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
