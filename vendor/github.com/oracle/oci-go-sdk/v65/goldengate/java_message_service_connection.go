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

// JavaMessageServiceConnection Represents the metadata of a Java Message Service Connection.
type JavaMessageServiceConnection struct {

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

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the subnet being referenced.
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

	// The private IP address of the connection's endpoint in the customer's VCN, typically a
	// database endpoint or a big data endpoint (e.g. Kafka bootstrap server).
	// In case the privateIp is provided, the subnetId must also be provided.
	// In case the privateIp (and the subnetId) is not provided it is assumed the datasource is publicly accessible.
	// In case the connection is accessible only privately, the lack of privateIp will result in not being able to access the connection.
	PrivateIp *string `mandatory:"false" json:"privateIp"`

	// The Java Message Service technology type.
	TechnologyType JavaMessageServiceConnectionTechnologyTypeEnum `mandatory:"true" json:"technologyType"`

	// Security protocol for Java Message Service. If not provided, default is PLAIN.
	// Optional until 2024-06-27, in the release after it will be made required.
	SecurityProtocol JavaMessageServiceConnectionSecurityProtocolEnum `mandatory:"false" json:"securityProtocol,omitempty"`

	// Authentication type for Java Message Service.  If not provided, default is NONE.
	// Optional until 2024-06-27, in the release after it will be made required.
	AuthenticationType JavaMessageServiceConnectionAuthenticationTypeEnum `mandatory:"false" json:"authenticationType,omitempty"`

	// Possible lifecycle states for connection.
	LifecycleState ConnectionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

// GetId returns Id
func (m JavaMessageServiceConnection) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m JavaMessageServiceConnection) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m JavaMessageServiceConnection) GetDescription() *string {
	return m.Description
}

// GetCompartmentId returns CompartmentId
func (m JavaMessageServiceConnection) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetFreeformTags returns FreeformTags
func (m JavaMessageServiceConnection) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m JavaMessageServiceConnection) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m JavaMessageServiceConnection) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetLifecycleState returns LifecycleState
func (m JavaMessageServiceConnection) GetLifecycleState() ConnectionLifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m JavaMessageServiceConnection) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetTimeCreated returns TimeCreated
func (m JavaMessageServiceConnection) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m JavaMessageServiceConnection) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetVaultId returns VaultId
func (m JavaMessageServiceConnection) GetVaultId() *string {
	return m.VaultId
}

// GetKeyId returns KeyId
func (m JavaMessageServiceConnection) GetKeyId() *string {
	return m.KeyId
}

// GetIngressIps returns IngressIps
func (m JavaMessageServiceConnection) GetIngressIps() []IngressIpDetails {
	return m.IngressIps
}

// GetNsgIds returns NsgIds
func (m JavaMessageServiceConnection) GetNsgIds() []string {
	return m.NsgIds
}

// GetSubnetId returns SubnetId
func (m JavaMessageServiceConnection) GetSubnetId() *string {
	return m.SubnetId
}

func (m JavaMessageServiceConnection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m JavaMessageServiceConnection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingJavaMessageServiceConnectionTechnologyTypeEnum(string(m.TechnologyType)); !ok && m.TechnologyType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TechnologyType: %s. Supported values are: %s.", m.TechnologyType, strings.Join(GetJavaMessageServiceConnectionTechnologyTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingJavaMessageServiceConnectionSecurityProtocolEnum(string(m.SecurityProtocol)); !ok && m.SecurityProtocol != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SecurityProtocol: %s. Supported values are: %s.", m.SecurityProtocol, strings.Join(GetJavaMessageServiceConnectionSecurityProtocolEnumStringValues(), ",")))
	}
	if _, ok := GetMappingJavaMessageServiceConnectionAuthenticationTypeEnum(string(m.AuthenticationType)); !ok && m.AuthenticationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AuthenticationType: %s. Supported values are: %s.", m.AuthenticationType, strings.Join(GetJavaMessageServiceConnectionAuthenticationTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingConnectionLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetConnectionLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m JavaMessageServiceConnection) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeJavaMessageServiceConnection JavaMessageServiceConnection
	s := struct {
		DiscriminatorParam string `json:"connectionType"`
		MarshalTypeJavaMessageServiceConnection
	}{
		"JAVA_MESSAGE_SERVICE",
		(MarshalTypeJavaMessageServiceConnection)(m),
	}

	return json.Marshal(&s)
}

// JavaMessageServiceConnectionTechnologyTypeEnum Enum with underlying type: string
type JavaMessageServiceConnectionTechnologyTypeEnum string

// Set of constants representing the allowable values for JavaMessageServiceConnectionTechnologyTypeEnum
const (
	JavaMessageServiceConnectionTechnologyTypeOracleWeblogicJms JavaMessageServiceConnectionTechnologyTypeEnum = "ORACLE_WEBLOGIC_JMS"
)

var mappingJavaMessageServiceConnectionTechnologyTypeEnum = map[string]JavaMessageServiceConnectionTechnologyTypeEnum{
	"ORACLE_WEBLOGIC_JMS": JavaMessageServiceConnectionTechnologyTypeOracleWeblogicJms,
}

var mappingJavaMessageServiceConnectionTechnologyTypeEnumLowerCase = map[string]JavaMessageServiceConnectionTechnologyTypeEnum{
	"oracle_weblogic_jms": JavaMessageServiceConnectionTechnologyTypeOracleWeblogicJms,
}

// GetJavaMessageServiceConnectionTechnologyTypeEnumValues Enumerates the set of values for JavaMessageServiceConnectionTechnologyTypeEnum
func GetJavaMessageServiceConnectionTechnologyTypeEnumValues() []JavaMessageServiceConnectionTechnologyTypeEnum {
	values := make([]JavaMessageServiceConnectionTechnologyTypeEnum, 0)
	for _, v := range mappingJavaMessageServiceConnectionTechnologyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetJavaMessageServiceConnectionTechnologyTypeEnumStringValues Enumerates the set of values in String for JavaMessageServiceConnectionTechnologyTypeEnum
func GetJavaMessageServiceConnectionTechnologyTypeEnumStringValues() []string {
	return []string{
		"ORACLE_WEBLOGIC_JMS",
	}
}

// GetMappingJavaMessageServiceConnectionTechnologyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJavaMessageServiceConnectionTechnologyTypeEnum(val string) (JavaMessageServiceConnectionTechnologyTypeEnum, bool) {
	enum, ok := mappingJavaMessageServiceConnectionTechnologyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// JavaMessageServiceConnectionSecurityProtocolEnum Enum with underlying type: string
type JavaMessageServiceConnectionSecurityProtocolEnum string

// Set of constants representing the allowable values for JavaMessageServiceConnectionSecurityProtocolEnum
const (
	JavaMessageServiceConnectionSecurityProtocolPlain JavaMessageServiceConnectionSecurityProtocolEnum = "PLAIN"
	JavaMessageServiceConnectionSecurityProtocolTls   JavaMessageServiceConnectionSecurityProtocolEnum = "TLS"
	JavaMessageServiceConnectionSecurityProtocolMtls  JavaMessageServiceConnectionSecurityProtocolEnum = "MTLS"
)

var mappingJavaMessageServiceConnectionSecurityProtocolEnum = map[string]JavaMessageServiceConnectionSecurityProtocolEnum{
	"PLAIN": JavaMessageServiceConnectionSecurityProtocolPlain,
	"TLS":   JavaMessageServiceConnectionSecurityProtocolTls,
	"MTLS":  JavaMessageServiceConnectionSecurityProtocolMtls,
}

var mappingJavaMessageServiceConnectionSecurityProtocolEnumLowerCase = map[string]JavaMessageServiceConnectionSecurityProtocolEnum{
	"plain": JavaMessageServiceConnectionSecurityProtocolPlain,
	"tls":   JavaMessageServiceConnectionSecurityProtocolTls,
	"mtls":  JavaMessageServiceConnectionSecurityProtocolMtls,
}

// GetJavaMessageServiceConnectionSecurityProtocolEnumValues Enumerates the set of values for JavaMessageServiceConnectionSecurityProtocolEnum
func GetJavaMessageServiceConnectionSecurityProtocolEnumValues() []JavaMessageServiceConnectionSecurityProtocolEnum {
	values := make([]JavaMessageServiceConnectionSecurityProtocolEnum, 0)
	for _, v := range mappingJavaMessageServiceConnectionSecurityProtocolEnum {
		values = append(values, v)
	}
	return values
}

// GetJavaMessageServiceConnectionSecurityProtocolEnumStringValues Enumerates the set of values in String for JavaMessageServiceConnectionSecurityProtocolEnum
func GetJavaMessageServiceConnectionSecurityProtocolEnumStringValues() []string {
	return []string{
		"PLAIN",
		"TLS",
		"MTLS",
	}
}

// GetMappingJavaMessageServiceConnectionSecurityProtocolEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJavaMessageServiceConnectionSecurityProtocolEnum(val string) (JavaMessageServiceConnectionSecurityProtocolEnum, bool) {
	enum, ok := mappingJavaMessageServiceConnectionSecurityProtocolEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// JavaMessageServiceConnectionAuthenticationTypeEnum Enum with underlying type: string
type JavaMessageServiceConnectionAuthenticationTypeEnum string

// Set of constants representing the allowable values for JavaMessageServiceConnectionAuthenticationTypeEnum
const (
	JavaMessageServiceConnectionAuthenticationTypeNone  JavaMessageServiceConnectionAuthenticationTypeEnum = "NONE"
	JavaMessageServiceConnectionAuthenticationTypeBasic JavaMessageServiceConnectionAuthenticationTypeEnum = "BASIC"
)

var mappingJavaMessageServiceConnectionAuthenticationTypeEnum = map[string]JavaMessageServiceConnectionAuthenticationTypeEnum{
	"NONE":  JavaMessageServiceConnectionAuthenticationTypeNone,
	"BASIC": JavaMessageServiceConnectionAuthenticationTypeBasic,
}

var mappingJavaMessageServiceConnectionAuthenticationTypeEnumLowerCase = map[string]JavaMessageServiceConnectionAuthenticationTypeEnum{
	"none":  JavaMessageServiceConnectionAuthenticationTypeNone,
	"basic": JavaMessageServiceConnectionAuthenticationTypeBasic,
}

// GetJavaMessageServiceConnectionAuthenticationTypeEnumValues Enumerates the set of values for JavaMessageServiceConnectionAuthenticationTypeEnum
func GetJavaMessageServiceConnectionAuthenticationTypeEnumValues() []JavaMessageServiceConnectionAuthenticationTypeEnum {
	values := make([]JavaMessageServiceConnectionAuthenticationTypeEnum, 0)
	for _, v := range mappingJavaMessageServiceConnectionAuthenticationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetJavaMessageServiceConnectionAuthenticationTypeEnumStringValues Enumerates the set of values in String for JavaMessageServiceConnectionAuthenticationTypeEnum
func GetJavaMessageServiceConnectionAuthenticationTypeEnumStringValues() []string {
	return []string{
		"NONE",
		"BASIC",
	}
}

// GetMappingJavaMessageServiceConnectionAuthenticationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJavaMessageServiceConnectionAuthenticationTypeEnum(val string) (JavaMessageServiceConnectionAuthenticationTypeEnum, bool) {
	enum, ok := mappingJavaMessageServiceConnectionAuthenticationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
