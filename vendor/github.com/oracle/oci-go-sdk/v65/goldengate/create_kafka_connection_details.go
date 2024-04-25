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

// CreateKafkaConnectionDetails The information about a new Kafka Connection.
type CreateKafkaConnectionDetails struct {

	// An object's Display Name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment being referenced.
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

	// The password for the cert inside of the KeyStore.
	// In case it differs from the KeyStore password, it should be provided.
	SslKeyPassword *string `mandatory:"false" json:"sslKeyPassword"`

	// The base64 encoded content of the consumer.properties file.
	ConsumerProperties *string `mandatory:"false" json:"consumerProperties"`

	// The base64 encoded content of the producer.properties file.
	ProducerProperties *string `mandatory:"false" json:"producerProperties"`

	// Controls the network traffic direction to the target:
	// SHARED_SERVICE_ENDPOINT: Traffic flows through the Goldengate Service's network to public hosts. Cannot be used for private targets.
	// SHARED_DEPLOYMENT_ENDPOINT: Network traffic flows from the assigned deployment's private endpoint through the deployment's subnet.
	// DEDICATED_ENDPOINT: A dedicated private endpoint is created in the target VCN subnet for the connection. The subnetId is required when DEDICATED_ENDPOINT networking is selected.
	RoutingMethod RoutingMethodEnum `mandatory:"false" json:"routingMethod,omitempty"`

	// The Kafka technology type.
	TechnologyType KafkaConnectionTechnologyTypeEnum `mandatory:"true" json:"technologyType"`

	// Security Type for Kafka.
	SecurityProtocol KafkaConnectionSecurityProtocolEnum `mandatory:"false" json:"securityProtocol,omitempty"`
}

// GetDisplayName returns DisplayName
func (m CreateKafkaConnectionDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m CreateKafkaConnectionDetails) GetDescription() *string {
	return m.Description
}

// GetCompartmentId returns CompartmentId
func (m CreateKafkaConnectionDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetFreeformTags returns FreeformTags
func (m CreateKafkaConnectionDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateKafkaConnectionDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetLocks returns Locks
func (m CreateKafkaConnectionDetails) GetLocks() []AddResourceLockDetails {
	return m.Locks
}

// GetVaultId returns VaultId
func (m CreateKafkaConnectionDetails) GetVaultId() *string {
	return m.VaultId
}

// GetKeyId returns KeyId
func (m CreateKafkaConnectionDetails) GetKeyId() *string {
	return m.KeyId
}

// GetNsgIds returns NsgIds
func (m CreateKafkaConnectionDetails) GetNsgIds() []string {
	return m.NsgIds
}

// GetSubnetId returns SubnetId
func (m CreateKafkaConnectionDetails) GetSubnetId() *string {
	return m.SubnetId
}

// GetRoutingMethod returns RoutingMethod
func (m CreateKafkaConnectionDetails) GetRoutingMethod() RoutingMethodEnum {
	return m.RoutingMethod
}

func (m CreateKafkaConnectionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateKafkaConnectionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingRoutingMethodEnum(string(m.RoutingMethod)); !ok && m.RoutingMethod != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RoutingMethod: %s. Supported values are: %s.", m.RoutingMethod, strings.Join(GetRoutingMethodEnumStringValues(), ",")))
	}
	if _, ok := GetMappingKafkaConnectionTechnologyTypeEnum(string(m.TechnologyType)); !ok && m.TechnologyType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TechnologyType: %s. Supported values are: %s.", m.TechnologyType, strings.Join(GetKafkaConnectionTechnologyTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingKafkaConnectionSecurityProtocolEnum(string(m.SecurityProtocol)); !ok && m.SecurityProtocol != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SecurityProtocol: %s. Supported values are: %s.", m.SecurityProtocol, strings.Join(GetKafkaConnectionSecurityProtocolEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateKafkaConnectionDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateKafkaConnectionDetails CreateKafkaConnectionDetails
	s := struct {
		DiscriminatorParam string `json:"connectionType"`
		MarshalTypeCreateKafkaConnectionDetails
	}{
		"KAFKA",
		(MarshalTypeCreateKafkaConnectionDetails)(m),
	}

	return json.Marshal(&s)
}
