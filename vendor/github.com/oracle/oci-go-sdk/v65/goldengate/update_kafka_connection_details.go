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

// UpdateKafkaConnectionDetails The information to update an Kafka Connection.
type UpdateKafkaConnectionDetails struct {

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

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the stream pool being referenced.
	StreamPoolId *string `mandatory:"false" json:"streamPoolId"`

	// Kafka bootstrap. Equivalent of bootstrap.servers configuration property in Kafka:
	// list of KafkaBootstrapServer objects specified by host/port.
	// Used for establishing the initial connection to the Kafka cluster.
	// Example: `"server1.example.com:9092,server2.example.com:9092"`
	BootstrapServers []KafkaBootstrapServer `mandatory:"false" json:"bootstrapServers"`

	// The username Oracle GoldenGate uses to connect the associated RDBMS.  This username must
	// already exist and be available for use by the database.  It must conform to the security
	// requirements implemented by the database including length, case sensitivity, and so on.
	Username *string `mandatory:"false" json:"username"`

	// The password Oracle GoldenGate uses to connect the associated RDBMS.  It must conform to the
	// specific security requirements implemented by the database including length, case
	// sensitivity, and so on.
	Password *string `mandatory:"false" json:"password"`

	// The base64 encoded content of the TrustStore file.
	TrustStore *string `mandatory:"false" json:"trustStore"`

	// The TrustStore password.
	TrustStorePassword *string `mandatory:"false" json:"trustStorePassword"`

	// The base64 encoded content of the KeyStore file.
	KeyStore *string `mandatory:"false" json:"keyStore"`

	// The KeyStore password.
	KeyStorePassword *string `mandatory:"false" json:"keyStorePassword"`

	// The base64 encoded content of the consumer.properties file.
	ConsumerProperties *string `mandatory:"false" json:"consumerProperties"`

	// The base64 encoded content of the producer.properties file.
	ProducerProperties *string `mandatory:"false" json:"producerProperties"`

	// Security Type for Kafka.
	SecurityProtocol KafkaConnectionSecurityProtocolEnum `mandatory:"false" json:"securityProtocol,omitempty"`
}

//GetDisplayName returns DisplayName
func (m UpdateKafkaConnectionDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetDescription returns Description
func (m UpdateKafkaConnectionDetails) GetDescription() *string {
	return m.Description
}

//GetFreeformTags returns FreeformTags
func (m UpdateKafkaConnectionDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m UpdateKafkaConnectionDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetVaultId returns VaultId
func (m UpdateKafkaConnectionDetails) GetVaultId() *string {
	return m.VaultId
}

//GetKeyId returns KeyId
func (m UpdateKafkaConnectionDetails) GetKeyId() *string {
	return m.KeyId
}

func (m UpdateKafkaConnectionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateKafkaConnectionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingKafkaConnectionSecurityProtocolEnum(string(m.SecurityProtocol)); !ok && m.SecurityProtocol != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SecurityProtocol: %s. Supported values are: %s.", m.SecurityProtocol, strings.Join(GetKafkaConnectionSecurityProtocolEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateKafkaConnectionDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateKafkaConnectionDetails UpdateKafkaConnectionDetails
	s := struct {
		DiscriminatorParam string `json:"connectionType"`
		MarshalTypeUpdateKafkaConnectionDetails
	}{
		"KAFKA",
		(MarshalTypeUpdateKafkaConnectionDetails)(m),
	}

	return json.Marshal(&s)
}
