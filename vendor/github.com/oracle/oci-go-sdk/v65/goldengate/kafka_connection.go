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

// KafkaConnection Represents the metadata of a Kafka Connection.
type KafkaConnection struct {

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

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the stream pool being referenced.
	StreamPoolId *string `mandatory:"false" json:"streamPoolId"`

	// Kafka bootstrap. Equivalent of bootstrap.servers configuration property in Kafka:
	// list of KafkaBootstrapServer objects specified by host/port.
	// Used for establishing the initial connection to the Kafka cluster.
	// Example: `"server1.example.com:9092,server2.example.com:9092"`
	BootstrapServers []KafkaBootstrapServer `mandatory:"false" json:"bootstrapServers"`

	// The username Oracle GoldenGate uses to connect the associated system of the given technology.
	// This username must already exist and be available by the system/application to be connected to
	// and must conform to the case sensitivty requirments defined in it.
	Username *string `mandatory:"false" json:"username"`

	// The Kafka technology type.
	TechnologyType KafkaConnectionTechnologyTypeEnum `mandatory:"true" json:"technologyType"`

	// Kafka security protocol.
	SecurityProtocol KafkaConnectionSecurityProtocolEnum `mandatory:"false" json:"securityProtocol,omitempty"`

	// Possible lifecycle states for connection.
	LifecycleState ConnectionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Controls the network traffic direction to the target:
	// SHARED_SERVICE_ENDPOINT: Traffic flows through the Goldengate Service's network to public hosts. Cannot be used for private targets.
	// SHARED_DEPLOYMENT_ENDPOINT: Network traffic flows from the assigned deployment's private endpoint through the deployment's subnet.
	// DEDICATED_ENDPOINT: A dedicated private endpoint is created in the target VCN subnet for the connection. The subnetId is required when DEDICATED_ENDPOINT networking is selected.
	RoutingMethod RoutingMethodEnum `mandatory:"false" json:"routingMethod,omitempty"`
}

// GetId returns Id
func (m KafkaConnection) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m KafkaConnection) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m KafkaConnection) GetDescription() *string {
	return m.Description
}

// GetCompartmentId returns CompartmentId
func (m KafkaConnection) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetFreeformTags returns FreeformTags
func (m KafkaConnection) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m KafkaConnection) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m KafkaConnection) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetLifecycleState returns LifecycleState
func (m KafkaConnection) GetLifecycleState() ConnectionLifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m KafkaConnection) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetTimeCreated returns TimeCreated
func (m KafkaConnection) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m KafkaConnection) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetVaultId returns VaultId
func (m KafkaConnection) GetVaultId() *string {
	return m.VaultId
}

// GetKeyId returns KeyId
func (m KafkaConnection) GetKeyId() *string {
	return m.KeyId
}

// GetIngressIps returns IngressIps
func (m KafkaConnection) GetIngressIps() []IngressIpDetails {
	return m.IngressIps
}

// GetNsgIds returns NsgIds
func (m KafkaConnection) GetNsgIds() []string {
	return m.NsgIds
}

// GetSubnetId returns SubnetId
func (m KafkaConnection) GetSubnetId() *string {
	return m.SubnetId
}

// GetRoutingMethod returns RoutingMethod
func (m KafkaConnection) GetRoutingMethod() RoutingMethodEnum {
	return m.RoutingMethod
}

func (m KafkaConnection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m KafkaConnection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingKafkaConnectionTechnologyTypeEnum(string(m.TechnologyType)); !ok && m.TechnologyType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TechnologyType: %s. Supported values are: %s.", m.TechnologyType, strings.Join(GetKafkaConnectionTechnologyTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingKafkaConnectionSecurityProtocolEnum(string(m.SecurityProtocol)); !ok && m.SecurityProtocol != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SecurityProtocol: %s. Supported values are: %s.", m.SecurityProtocol, strings.Join(GetKafkaConnectionSecurityProtocolEnumStringValues(), ",")))
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
func (m KafkaConnection) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeKafkaConnection KafkaConnection
	s := struct {
		DiscriminatorParam string `json:"connectionType"`
		MarshalTypeKafkaConnection
	}{
		"KAFKA",
		(MarshalTypeKafkaConnection)(m),
	}

	return json.Marshal(&s)
}

// KafkaConnectionTechnologyTypeEnum Enum with underlying type: string
type KafkaConnectionTechnologyTypeEnum string

// Set of constants representing the allowable values for KafkaConnectionTechnologyTypeEnum
const (
	KafkaConnectionTechnologyTypeApacheKafka    KafkaConnectionTechnologyTypeEnum = "APACHE_KAFKA"
	KafkaConnectionTechnologyTypeAzureEventHubs KafkaConnectionTechnologyTypeEnum = "AZURE_EVENT_HUBS"
	KafkaConnectionTechnologyTypeConfluentKafka KafkaConnectionTechnologyTypeEnum = "CONFLUENT_KAFKA"
	KafkaConnectionTechnologyTypeOciStreaming   KafkaConnectionTechnologyTypeEnum = "OCI_STREAMING"
)

var mappingKafkaConnectionTechnologyTypeEnum = map[string]KafkaConnectionTechnologyTypeEnum{
	"APACHE_KAFKA":     KafkaConnectionTechnologyTypeApacheKafka,
	"AZURE_EVENT_HUBS": KafkaConnectionTechnologyTypeAzureEventHubs,
	"CONFLUENT_KAFKA":  KafkaConnectionTechnologyTypeConfluentKafka,
	"OCI_STREAMING":    KafkaConnectionTechnologyTypeOciStreaming,
}

var mappingKafkaConnectionTechnologyTypeEnumLowerCase = map[string]KafkaConnectionTechnologyTypeEnum{
	"apache_kafka":     KafkaConnectionTechnologyTypeApacheKafka,
	"azure_event_hubs": KafkaConnectionTechnologyTypeAzureEventHubs,
	"confluent_kafka":  KafkaConnectionTechnologyTypeConfluentKafka,
	"oci_streaming":    KafkaConnectionTechnologyTypeOciStreaming,
}

// GetKafkaConnectionTechnologyTypeEnumValues Enumerates the set of values for KafkaConnectionTechnologyTypeEnum
func GetKafkaConnectionTechnologyTypeEnumValues() []KafkaConnectionTechnologyTypeEnum {
	values := make([]KafkaConnectionTechnologyTypeEnum, 0)
	for _, v := range mappingKafkaConnectionTechnologyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetKafkaConnectionTechnologyTypeEnumStringValues Enumerates the set of values in String for KafkaConnectionTechnologyTypeEnum
func GetKafkaConnectionTechnologyTypeEnumStringValues() []string {
	return []string{
		"APACHE_KAFKA",
		"AZURE_EVENT_HUBS",
		"CONFLUENT_KAFKA",
		"OCI_STREAMING",
	}
}

// GetMappingKafkaConnectionTechnologyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingKafkaConnectionTechnologyTypeEnum(val string) (KafkaConnectionTechnologyTypeEnum, bool) {
	enum, ok := mappingKafkaConnectionTechnologyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// KafkaConnectionSecurityProtocolEnum Enum with underlying type: string
type KafkaConnectionSecurityProtocolEnum string

// Set of constants representing the allowable values for KafkaConnectionSecurityProtocolEnum
const (
	KafkaConnectionSecurityProtocolSsl           KafkaConnectionSecurityProtocolEnum = "SSL"
	KafkaConnectionSecurityProtocolSaslSsl       KafkaConnectionSecurityProtocolEnum = "SASL_SSL"
	KafkaConnectionSecurityProtocolPlaintext     KafkaConnectionSecurityProtocolEnum = "PLAINTEXT"
	KafkaConnectionSecurityProtocolSaslPlaintext KafkaConnectionSecurityProtocolEnum = "SASL_PLAINTEXT"
)

var mappingKafkaConnectionSecurityProtocolEnum = map[string]KafkaConnectionSecurityProtocolEnum{
	"SSL":            KafkaConnectionSecurityProtocolSsl,
	"SASL_SSL":       KafkaConnectionSecurityProtocolSaslSsl,
	"PLAINTEXT":      KafkaConnectionSecurityProtocolPlaintext,
	"SASL_PLAINTEXT": KafkaConnectionSecurityProtocolSaslPlaintext,
}

var mappingKafkaConnectionSecurityProtocolEnumLowerCase = map[string]KafkaConnectionSecurityProtocolEnum{
	"ssl":            KafkaConnectionSecurityProtocolSsl,
	"sasl_ssl":       KafkaConnectionSecurityProtocolSaslSsl,
	"plaintext":      KafkaConnectionSecurityProtocolPlaintext,
	"sasl_plaintext": KafkaConnectionSecurityProtocolSaslPlaintext,
}

// GetKafkaConnectionSecurityProtocolEnumValues Enumerates the set of values for KafkaConnectionSecurityProtocolEnum
func GetKafkaConnectionSecurityProtocolEnumValues() []KafkaConnectionSecurityProtocolEnum {
	values := make([]KafkaConnectionSecurityProtocolEnum, 0)
	for _, v := range mappingKafkaConnectionSecurityProtocolEnum {
		values = append(values, v)
	}
	return values
}

// GetKafkaConnectionSecurityProtocolEnumStringValues Enumerates the set of values in String for KafkaConnectionSecurityProtocolEnum
func GetKafkaConnectionSecurityProtocolEnumStringValues() []string {
	return []string{
		"SSL",
		"SASL_SSL",
		"PLAINTEXT",
		"SASL_PLAINTEXT",
	}
}

// GetMappingKafkaConnectionSecurityProtocolEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingKafkaConnectionSecurityProtocolEnum(val string) (KafkaConnectionSecurityProtocolEnum, bool) {
	enum, ok := mappingKafkaConnectionSecurityProtocolEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
