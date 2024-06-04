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

// ElasticsearchConnection Represents the metadata of a Elasticsearch Connection.
type ElasticsearchConnection struct {

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

	// Comma separated list of Elasticsearch server addresses, specified as host:port entries, where :port is optional.
	// If port is not specified, it defaults to 9200.
	// Used for establishing the initial connection to the Elasticsearch cluster.
	// Example: `"server1.example.com:4000,server2.example.com:4000"`
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

	// The Elasticsearch technology type.
	TechnologyType ElasticsearchConnectionTechnologyTypeEnum `mandatory:"true" json:"technologyType"`

	// Security protocol for Elasticsearch
	SecurityProtocol ElasticsearchConnectionSecurityProtocolEnum `mandatory:"true" json:"securityProtocol"`

	// Authentication type for Elasticsearch.
	AuthenticationType ElasticsearchConnectionAuthenticationTypeEnum `mandatory:"true" json:"authenticationType"`

	// Possible lifecycle states for connection.
	LifecycleState ConnectionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Controls the network traffic direction to the target:
	// SHARED_SERVICE_ENDPOINT: Traffic flows through the Goldengate Service's network to public hosts. Cannot be used for private targets.
	// SHARED_DEPLOYMENT_ENDPOINT: Network traffic flows from the assigned deployment's private endpoint through the deployment's subnet.
	// DEDICATED_ENDPOINT: A dedicated private endpoint is created in the target VCN subnet for the connection. The subnetId is required when DEDICATED_ENDPOINT networking is selected.
	RoutingMethod RoutingMethodEnum `mandatory:"false" json:"routingMethod,omitempty"`
}

// GetId returns Id
func (m ElasticsearchConnection) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m ElasticsearchConnection) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m ElasticsearchConnection) GetDescription() *string {
	return m.Description
}

// GetCompartmentId returns CompartmentId
func (m ElasticsearchConnection) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetFreeformTags returns FreeformTags
func (m ElasticsearchConnection) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m ElasticsearchConnection) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m ElasticsearchConnection) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetLifecycleState returns LifecycleState
func (m ElasticsearchConnection) GetLifecycleState() ConnectionLifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m ElasticsearchConnection) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetTimeCreated returns TimeCreated
func (m ElasticsearchConnection) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m ElasticsearchConnection) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLocks returns Locks
func (m ElasticsearchConnection) GetLocks() []ResourceLock {
	return m.Locks
}

// GetVaultId returns VaultId
func (m ElasticsearchConnection) GetVaultId() *string {
	return m.VaultId
}

// GetKeyId returns KeyId
func (m ElasticsearchConnection) GetKeyId() *string {
	return m.KeyId
}

// GetIngressIps returns IngressIps
func (m ElasticsearchConnection) GetIngressIps() []IngressIpDetails {
	return m.IngressIps
}

// GetNsgIds returns NsgIds
func (m ElasticsearchConnection) GetNsgIds() []string {
	return m.NsgIds
}

// GetSubnetId returns SubnetId
func (m ElasticsearchConnection) GetSubnetId() *string {
	return m.SubnetId
}

// GetRoutingMethod returns RoutingMethod
func (m ElasticsearchConnection) GetRoutingMethod() RoutingMethodEnum {
	return m.RoutingMethod
}

func (m ElasticsearchConnection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ElasticsearchConnection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingElasticsearchConnectionTechnologyTypeEnum(string(m.TechnologyType)); !ok && m.TechnologyType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TechnologyType: %s. Supported values are: %s.", m.TechnologyType, strings.Join(GetElasticsearchConnectionTechnologyTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingElasticsearchConnectionSecurityProtocolEnum(string(m.SecurityProtocol)); !ok && m.SecurityProtocol != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SecurityProtocol: %s. Supported values are: %s.", m.SecurityProtocol, strings.Join(GetElasticsearchConnectionSecurityProtocolEnumStringValues(), ",")))
	}
	if _, ok := GetMappingElasticsearchConnectionAuthenticationTypeEnum(string(m.AuthenticationType)); !ok && m.AuthenticationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AuthenticationType: %s. Supported values are: %s.", m.AuthenticationType, strings.Join(GetElasticsearchConnectionAuthenticationTypeEnumStringValues(), ",")))
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
func (m ElasticsearchConnection) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeElasticsearchConnection ElasticsearchConnection
	s := struct {
		DiscriminatorParam string `json:"connectionType"`
		MarshalTypeElasticsearchConnection
	}{
		"ELASTICSEARCH",
		(MarshalTypeElasticsearchConnection)(m),
	}

	return json.Marshal(&s)
}

// ElasticsearchConnectionTechnologyTypeEnum Enum with underlying type: string
type ElasticsearchConnectionTechnologyTypeEnum string

// Set of constants representing the allowable values for ElasticsearchConnectionTechnologyTypeEnum
const (
	ElasticsearchConnectionTechnologyTypeElasticsearch ElasticsearchConnectionTechnologyTypeEnum = "ELASTICSEARCH"
)

var mappingElasticsearchConnectionTechnologyTypeEnum = map[string]ElasticsearchConnectionTechnologyTypeEnum{
	"ELASTICSEARCH": ElasticsearchConnectionTechnologyTypeElasticsearch,
}

var mappingElasticsearchConnectionTechnologyTypeEnumLowerCase = map[string]ElasticsearchConnectionTechnologyTypeEnum{
	"elasticsearch": ElasticsearchConnectionTechnologyTypeElasticsearch,
}

// GetElasticsearchConnectionTechnologyTypeEnumValues Enumerates the set of values for ElasticsearchConnectionTechnologyTypeEnum
func GetElasticsearchConnectionTechnologyTypeEnumValues() []ElasticsearchConnectionTechnologyTypeEnum {
	values := make([]ElasticsearchConnectionTechnologyTypeEnum, 0)
	for _, v := range mappingElasticsearchConnectionTechnologyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetElasticsearchConnectionTechnologyTypeEnumStringValues Enumerates the set of values in String for ElasticsearchConnectionTechnologyTypeEnum
func GetElasticsearchConnectionTechnologyTypeEnumStringValues() []string {
	return []string{
		"ELASTICSEARCH",
	}
}

// GetMappingElasticsearchConnectionTechnologyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingElasticsearchConnectionTechnologyTypeEnum(val string) (ElasticsearchConnectionTechnologyTypeEnum, bool) {
	enum, ok := mappingElasticsearchConnectionTechnologyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ElasticsearchConnectionSecurityProtocolEnum Enum with underlying type: string
type ElasticsearchConnectionSecurityProtocolEnum string

// Set of constants representing the allowable values for ElasticsearchConnectionSecurityProtocolEnum
const (
	ElasticsearchConnectionSecurityProtocolPlain ElasticsearchConnectionSecurityProtocolEnum = "PLAIN"
	ElasticsearchConnectionSecurityProtocolTls   ElasticsearchConnectionSecurityProtocolEnum = "TLS"
)

var mappingElasticsearchConnectionSecurityProtocolEnum = map[string]ElasticsearchConnectionSecurityProtocolEnum{
	"PLAIN": ElasticsearchConnectionSecurityProtocolPlain,
	"TLS":   ElasticsearchConnectionSecurityProtocolTls,
}

var mappingElasticsearchConnectionSecurityProtocolEnumLowerCase = map[string]ElasticsearchConnectionSecurityProtocolEnum{
	"plain": ElasticsearchConnectionSecurityProtocolPlain,
	"tls":   ElasticsearchConnectionSecurityProtocolTls,
}

// GetElasticsearchConnectionSecurityProtocolEnumValues Enumerates the set of values for ElasticsearchConnectionSecurityProtocolEnum
func GetElasticsearchConnectionSecurityProtocolEnumValues() []ElasticsearchConnectionSecurityProtocolEnum {
	values := make([]ElasticsearchConnectionSecurityProtocolEnum, 0)
	for _, v := range mappingElasticsearchConnectionSecurityProtocolEnum {
		values = append(values, v)
	}
	return values
}

// GetElasticsearchConnectionSecurityProtocolEnumStringValues Enumerates the set of values in String for ElasticsearchConnectionSecurityProtocolEnum
func GetElasticsearchConnectionSecurityProtocolEnumStringValues() []string {
	return []string{
		"PLAIN",
		"TLS",
	}
}

// GetMappingElasticsearchConnectionSecurityProtocolEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingElasticsearchConnectionSecurityProtocolEnum(val string) (ElasticsearchConnectionSecurityProtocolEnum, bool) {
	enum, ok := mappingElasticsearchConnectionSecurityProtocolEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ElasticsearchConnectionAuthenticationTypeEnum Enum with underlying type: string
type ElasticsearchConnectionAuthenticationTypeEnum string

// Set of constants representing the allowable values for ElasticsearchConnectionAuthenticationTypeEnum
const (
	ElasticsearchConnectionAuthenticationTypeNone  ElasticsearchConnectionAuthenticationTypeEnum = "NONE"
	ElasticsearchConnectionAuthenticationTypeBasic ElasticsearchConnectionAuthenticationTypeEnum = "BASIC"
)

var mappingElasticsearchConnectionAuthenticationTypeEnum = map[string]ElasticsearchConnectionAuthenticationTypeEnum{
	"NONE":  ElasticsearchConnectionAuthenticationTypeNone,
	"BASIC": ElasticsearchConnectionAuthenticationTypeBasic,
}

var mappingElasticsearchConnectionAuthenticationTypeEnumLowerCase = map[string]ElasticsearchConnectionAuthenticationTypeEnum{
	"none":  ElasticsearchConnectionAuthenticationTypeNone,
	"basic": ElasticsearchConnectionAuthenticationTypeBasic,
}

// GetElasticsearchConnectionAuthenticationTypeEnumValues Enumerates the set of values for ElasticsearchConnectionAuthenticationTypeEnum
func GetElasticsearchConnectionAuthenticationTypeEnumValues() []ElasticsearchConnectionAuthenticationTypeEnum {
	values := make([]ElasticsearchConnectionAuthenticationTypeEnum, 0)
	for _, v := range mappingElasticsearchConnectionAuthenticationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetElasticsearchConnectionAuthenticationTypeEnumStringValues Enumerates the set of values in String for ElasticsearchConnectionAuthenticationTypeEnum
func GetElasticsearchConnectionAuthenticationTypeEnumStringValues() []string {
	return []string{
		"NONE",
		"BASIC",
	}
}

// GetMappingElasticsearchConnectionAuthenticationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingElasticsearchConnectionAuthenticationTypeEnum(val string) (ElasticsearchConnectionAuthenticationTypeEnum, bool) {
	enum, ok := mappingElasticsearchConnectionAuthenticationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
