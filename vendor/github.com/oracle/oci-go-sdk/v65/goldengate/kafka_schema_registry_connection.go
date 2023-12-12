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

// KafkaSchemaRegistryConnection Represents the metadata of a Kafka (e.g. Confluent) Schema Registry Connection.
type KafkaSchemaRegistryConnection struct {

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

	// Kafka Schema Registry URL.
	// e.g.: 'https://server1.us.oracle.com:8081'
	Url *string `mandatory:"true" json:"url"`

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

	// The username to access Schema Registry using basic authentation.
	// This value is injected into 'schema.registry.basic.auth.user.info=user:password' configuration property.
	Username *string `mandatory:"false" json:"username"`

	// Deprecated: this field will be removed in future versions. Either specify the private IP in the connectionString or host
	// field, or make sure the host name is resolvable in the target VCN.
	// The private IP address of the connection's endpoint in the customer's VCN, typically a
	// database endpoint or a big data endpoint (e.g. Kafka bootstrap server).
	// In case the privateIp is provided, the subnetId must also be provided.
	// In case the privateIp (and the subnetId) is not provided it is assumed the datasource is publicly accessible.
	// In case the connection is accessible only privately, the lack of privateIp will result in not being able to access the connection.
	PrivateIp *string `mandatory:"false" json:"privateIp"`

	// The Kafka (e.g. Confluent) Schema Registry technology type.
	TechnologyType KafkaSchemaRegistryConnectionTechnologyTypeEnum `mandatory:"true" json:"technologyType"`

	// Used authentication mechanism to access Schema Registry.
	AuthenticationType KafkaSchemaRegistryConnectionAuthenticationTypeEnum `mandatory:"true" json:"authenticationType"`

	// Possible lifecycle states for connection.
	LifecycleState ConnectionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Controls the network traffic direction to the target:
	// SHARED_SERVICE_ENDPOINT: Traffic flows through the Goldengate Service's network to public hosts. Cannot be used for private targets.
	// SHARED_DEPLOYMENT_ENDPOINT: Network traffic flows from the assigned deployment's private endpoint through the deployment's subnet.
	// DEDICATED_ENDPOINT: A dedicated private endpoint is created in the target VCN subnet for the connection. The subnetId is required when DEDICATED_ENDPOINT networking is selected.
	RoutingMethod RoutingMethodEnum `mandatory:"false" json:"routingMethod,omitempty"`
}

// GetId returns Id
func (m KafkaSchemaRegistryConnection) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m KafkaSchemaRegistryConnection) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m KafkaSchemaRegistryConnection) GetDescription() *string {
	return m.Description
}

// GetCompartmentId returns CompartmentId
func (m KafkaSchemaRegistryConnection) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetFreeformTags returns FreeformTags
func (m KafkaSchemaRegistryConnection) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m KafkaSchemaRegistryConnection) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m KafkaSchemaRegistryConnection) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetLifecycleState returns LifecycleState
func (m KafkaSchemaRegistryConnection) GetLifecycleState() ConnectionLifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m KafkaSchemaRegistryConnection) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetTimeCreated returns TimeCreated
func (m KafkaSchemaRegistryConnection) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m KafkaSchemaRegistryConnection) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetVaultId returns VaultId
func (m KafkaSchemaRegistryConnection) GetVaultId() *string {
	return m.VaultId
}

// GetKeyId returns KeyId
func (m KafkaSchemaRegistryConnection) GetKeyId() *string {
	return m.KeyId
}

// GetIngressIps returns IngressIps
func (m KafkaSchemaRegistryConnection) GetIngressIps() []IngressIpDetails {
	return m.IngressIps
}

// GetNsgIds returns NsgIds
func (m KafkaSchemaRegistryConnection) GetNsgIds() []string {
	return m.NsgIds
}

// GetSubnetId returns SubnetId
func (m KafkaSchemaRegistryConnection) GetSubnetId() *string {
	return m.SubnetId
}

// GetRoutingMethod returns RoutingMethod
func (m KafkaSchemaRegistryConnection) GetRoutingMethod() RoutingMethodEnum {
	return m.RoutingMethod
}

func (m KafkaSchemaRegistryConnection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m KafkaSchemaRegistryConnection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingKafkaSchemaRegistryConnectionTechnologyTypeEnum(string(m.TechnologyType)); !ok && m.TechnologyType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TechnologyType: %s. Supported values are: %s.", m.TechnologyType, strings.Join(GetKafkaSchemaRegistryConnectionTechnologyTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingKafkaSchemaRegistryConnectionAuthenticationTypeEnum(string(m.AuthenticationType)); !ok && m.AuthenticationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AuthenticationType: %s. Supported values are: %s.", m.AuthenticationType, strings.Join(GetKafkaSchemaRegistryConnectionAuthenticationTypeEnumStringValues(), ",")))
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
func (m KafkaSchemaRegistryConnection) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeKafkaSchemaRegistryConnection KafkaSchemaRegistryConnection
	s := struct {
		DiscriminatorParam string `json:"connectionType"`
		MarshalTypeKafkaSchemaRegistryConnection
	}{
		"KAFKA_SCHEMA_REGISTRY",
		(MarshalTypeKafkaSchemaRegistryConnection)(m),
	}

	return json.Marshal(&s)
}

// KafkaSchemaRegistryConnectionTechnologyTypeEnum Enum with underlying type: string
type KafkaSchemaRegistryConnectionTechnologyTypeEnum string

// Set of constants representing the allowable values for KafkaSchemaRegistryConnectionTechnologyTypeEnum
const (
	KafkaSchemaRegistryConnectionTechnologyTypeConfluentSchemaRegistry KafkaSchemaRegistryConnectionTechnologyTypeEnum = "CONFLUENT_SCHEMA_REGISTRY"
)

var mappingKafkaSchemaRegistryConnectionTechnologyTypeEnum = map[string]KafkaSchemaRegistryConnectionTechnologyTypeEnum{
	"CONFLUENT_SCHEMA_REGISTRY": KafkaSchemaRegistryConnectionTechnologyTypeConfluentSchemaRegistry,
}

var mappingKafkaSchemaRegistryConnectionTechnologyTypeEnumLowerCase = map[string]KafkaSchemaRegistryConnectionTechnologyTypeEnum{
	"confluent_schema_registry": KafkaSchemaRegistryConnectionTechnologyTypeConfluentSchemaRegistry,
}

// GetKafkaSchemaRegistryConnectionTechnologyTypeEnumValues Enumerates the set of values for KafkaSchemaRegistryConnectionTechnologyTypeEnum
func GetKafkaSchemaRegistryConnectionTechnologyTypeEnumValues() []KafkaSchemaRegistryConnectionTechnologyTypeEnum {
	values := make([]KafkaSchemaRegistryConnectionTechnologyTypeEnum, 0)
	for _, v := range mappingKafkaSchemaRegistryConnectionTechnologyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetKafkaSchemaRegistryConnectionTechnologyTypeEnumStringValues Enumerates the set of values in String for KafkaSchemaRegistryConnectionTechnologyTypeEnum
func GetKafkaSchemaRegistryConnectionTechnologyTypeEnumStringValues() []string {
	return []string{
		"CONFLUENT_SCHEMA_REGISTRY",
	}
}

// GetMappingKafkaSchemaRegistryConnectionTechnologyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingKafkaSchemaRegistryConnectionTechnologyTypeEnum(val string) (KafkaSchemaRegistryConnectionTechnologyTypeEnum, bool) {
	enum, ok := mappingKafkaSchemaRegistryConnectionTechnologyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// KafkaSchemaRegistryConnectionAuthenticationTypeEnum Enum with underlying type: string
type KafkaSchemaRegistryConnectionAuthenticationTypeEnum string

// Set of constants representing the allowable values for KafkaSchemaRegistryConnectionAuthenticationTypeEnum
const (
	KafkaSchemaRegistryConnectionAuthenticationTypeNone   KafkaSchemaRegistryConnectionAuthenticationTypeEnum = "NONE"
	KafkaSchemaRegistryConnectionAuthenticationTypeBasic  KafkaSchemaRegistryConnectionAuthenticationTypeEnum = "BASIC"
	KafkaSchemaRegistryConnectionAuthenticationTypeMutual KafkaSchemaRegistryConnectionAuthenticationTypeEnum = "MUTUAL"
)

var mappingKafkaSchemaRegistryConnectionAuthenticationTypeEnum = map[string]KafkaSchemaRegistryConnectionAuthenticationTypeEnum{
	"NONE":   KafkaSchemaRegistryConnectionAuthenticationTypeNone,
	"BASIC":  KafkaSchemaRegistryConnectionAuthenticationTypeBasic,
	"MUTUAL": KafkaSchemaRegistryConnectionAuthenticationTypeMutual,
}

var mappingKafkaSchemaRegistryConnectionAuthenticationTypeEnumLowerCase = map[string]KafkaSchemaRegistryConnectionAuthenticationTypeEnum{
	"none":   KafkaSchemaRegistryConnectionAuthenticationTypeNone,
	"basic":  KafkaSchemaRegistryConnectionAuthenticationTypeBasic,
	"mutual": KafkaSchemaRegistryConnectionAuthenticationTypeMutual,
}

// GetKafkaSchemaRegistryConnectionAuthenticationTypeEnumValues Enumerates the set of values for KafkaSchemaRegistryConnectionAuthenticationTypeEnum
func GetKafkaSchemaRegistryConnectionAuthenticationTypeEnumValues() []KafkaSchemaRegistryConnectionAuthenticationTypeEnum {
	values := make([]KafkaSchemaRegistryConnectionAuthenticationTypeEnum, 0)
	for _, v := range mappingKafkaSchemaRegistryConnectionAuthenticationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetKafkaSchemaRegistryConnectionAuthenticationTypeEnumStringValues Enumerates the set of values in String for KafkaSchemaRegistryConnectionAuthenticationTypeEnum
func GetKafkaSchemaRegistryConnectionAuthenticationTypeEnumStringValues() []string {
	return []string{
		"NONE",
		"BASIC",
		"MUTUAL",
	}
}

// GetMappingKafkaSchemaRegistryConnectionAuthenticationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingKafkaSchemaRegistryConnectionAuthenticationTypeEnum(val string) (KafkaSchemaRegistryConnectionAuthenticationTypeEnum, bool) {
	enum, ok := mappingKafkaSchemaRegistryConnectionAuthenticationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
