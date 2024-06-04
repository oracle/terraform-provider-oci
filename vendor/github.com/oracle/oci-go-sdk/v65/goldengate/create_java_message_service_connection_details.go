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

// CreateJavaMessageServiceConnectionDetails The information about a new Java Message Service Connection.
type CreateJavaMessageServiceConnectionDetails struct {

	// An object's Display Name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment being referenced.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// If set to true, Java Naming and Directory Interface (JNDI) properties should be provided.
	ShouldUseJndi *bool `mandatory:"true" json:"shouldUseJndi"`

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

	// The Connection Factory can be looked up using this name.
	// e.g.: 'ConnectionFactory'
	JndiConnectionFactory *string `mandatory:"false" json:"jndiConnectionFactory"`

	// The URL that Java Message Service will use to contact the JNDI provider.
	// e.g.: 'tcp://myjms.host.domain:61616?jms.prefetchPolicy.all=1000'
	JndiProviderUrl *string `mandatory:"false" json:"jndiProviderUrl"`

	// The implementation of javax.naming.spi.InitialContextFactory interface
	// that the client uses to obtain initial naming context.
	// e.g.: 'org.apache.activemq.jndi.ActiveMQInitialContextFactory'
	JndiInitialContextFactory *string `mandatory:"false" json:"jndiInitialContextFactory"`

	// Specifies the identity of the principal (user) to be authenticated.
	// e.g.: 'admin2'
	JndiSecurityPrincipal *string `mandatory:"false" json:"jndiSecurityPrincipal"`

	// The password associated to the principal.
	JndiSecurityCredentials *string `mandatory:"false" json:"jndiSecurityCredentials"`

	// Connectin URL of the Java Message Service, specifying the protocol, host, and port.
	// e.g.: 'mq://myjms.host.domain:7676'
	ConnectionUrl *string `mandatory:"false" json:"connectionUrl"`

	// The of Java class implementing javax.jms.ConnectionFactory interface
	// supplied by the Java Message Service provider.
	// e.g.: 'com.stc.jmsjca.core.JConnectionFactoryXA'
	ConnectionFactory *string `mandatory:"false" json:"connectionFactory"`

	// The username Oracle GoldenGate uses to connect to the Java Message Service.
	// This username must already exist and be available by the Java Message Service to be connected to.
	Username *string `mandatory:"false" json:"username"`

	// The password Oracle GoldenGate uses to connect the associated Java Message Service.
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

	// The Java Message Service technology type.
	TechnologyType JavaMessageServiceConnectionTechnologyTypeEnum `mandatory:"true" json:"technologyType"`

	// Security protocol for Java Message Service. If not provided, default is PLAIN.
	// Optional until 2024-06-27, in the release after it will be made required.
	SecurityProtocol JavaMessageServiceConnectionSecurityProtocolEnum `mandatory:"false" json:"securityProtocol,omitempty"`

	// Authentication type for Java Message Service.  If not provided, default is NONE.
	// Optional until 2024-06-27, in the release after it will be made required.
	AuthenticationType JavaMessageServiceConnectionAuthenticationTypeEnum `mandatory:"false" json:"authenticationType,omitempty"`
}

// GetDisplayName returns DisplayName
func (m CreateJavaMessageServiceConnectionDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m CreateJavaMessageServiceConnectionDetails) GetDescription() *string {
	return m.Description
}

// GetCompartmentId returns CompartmentId
func (m CreateJavaMessageServiceConnectionDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetFreeformTags returns FreeformTags
func (m CreateJavaMessageServiceConnectionDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateJavaMessageServiceConnectionDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetLocks returns Locks
func (m CreateJavaMessageServiceConnectionDetails) GetLocks() []AddResourceLockDetails {
	return m.Locks
}

// GetVaultId returns VaultId
func (m CreateJavaMessageServiceConnectionDetails) GetVaultId() *string {
	return m.VaultId
}

// GetKeyId returns KeyId
func (m CreateJavaMessageServiceConnectionDetails) GetKeyId() *string {
	return m.KeyId
}

// GetNsgIds returns NsgIds
func (m CreateJavaMessageServiceConnectionDetails) GetNsgIds() []string {
	return m.NsgIds
}

// GetSubnetId returns SubnetId
func (m CreateJavaMessageServiceConnectionDetails) GetSubnetId() *string {
	return m.SubnetId
}

// GetRoutingMethod returns RoutingMethod
func (m CreateJavaMessageServiceConnectionDetails) GetRoutingMethod() RoutingMethodEnum {
	return m.RoutingMethod
}

func (m CreateJavaMessageServiceConnectionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateJavaMessageServiceConnectionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingRoutingMethodEnum(string(m.RoutingMethod)); !ok && m.RoutingMethod != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RoutingMethod: %s. Supported values are: %s.", m.RoutingMethod, strings.Join(GetRoutingMethodEnumStringValues(), ",")))
	}
	if _, ok := GetMappingJavaMessageServiceConnectionTechnologyTypeEnum(string(m.TechnologyType)); !ok && m.TechnologyType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TechnologyType: %s. Supported values are: %s.", m.TechnologyType, strings.Join(GetJavaMessageServiceConnectionTechnologyTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingJavaMessageServiceConnectionSecurityProtocolEnum(string(m.SecurityProtocol)); !ok && m.SecurityProtocol != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SecurityProtocol: %s. Supported values are: %s.", m.SecurityProtocol, strings.Join(GetJavaMessageServiceConnectionSecurityProtocolEnumStringValues(), ",")))
	}
	if _, ok := GetMappingJavaMessageServiceConnectionAuthenticationTypeEnum(string(m.AuthenticationType)); !ok && m.AuthenticationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AuthenticationType: %s. Supported values are: %s.", m.AuthenticationType, strings.Join(GetJavaMessageServiceConnectionAuthenticationTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateJavaMessageServiceConnectionDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateJavaMessageServiceConnectionDetails CreateJavaMessageServiceConnectionDetails
	s := struct {
		DiscriminatorParam string `json:"connectionType"`
		MarshalTypeCreateJavaMessageServiceConnectionDetails
	}{
		"JAVA_MESSAGE_SERVICE",
		(MarshalTypeCreateJavaMessageServiceConnectionDetails)(m),
	}

	return json.Marshal(&s)
}
