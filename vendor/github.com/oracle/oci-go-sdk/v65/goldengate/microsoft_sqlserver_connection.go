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

// MicrosoftSqlserverConnection Represents the metadata of a Microsoft SQL Server Connection.
type MicrosoftSqlserverConnection struct {

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

	// The username Oracle GoldenGate uses to connect to the Microsoft SQL Server.
	// This username must already exist and be available by the Microsoft SQL Server to be connected to.
	Username *string `mandatory:"true" json:"username"`

	// The name or address of a host.
	Host *string `mandatory:"true" json:"host"`

	// The port of an endpoint usually specified for a connection.
	Port *int `mandatory:"true" json:"port"`

	// The name of the database.
	DatabaseName *string `mandatory:"true" json:"databaseName"`

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

	// An array of name-value pair attribute entries.
	// Used as additional parameters in connection string.
	AdditionalAttributes []NameValuePair `mandatory:"false" json:"additionalAttributes"`

	// Database Certificate - The base64 encoded content of a .pem or .crt file.
	// containing the server public key (for 1-way SSL).
	SslCa *string `mandatory:"false" json:"sslCa"`

	// If set to true, the driver validates the certificate that is sent by the database server.
	ShouldValidateServerCertificate *bool `mandatory:"false" json:"shouldValidateServerCertificate"`

	// Deprecated: this field will be removed in future versions. Either specify the private IP in the connectionString or host
	// field, or make sure the host name is resolvable in the target VCN.
	// The private IP address of the connection's endpoint in the customer's VCN, typically a
	// database endpoint or a big data endpoint (e.g. Kafka bootstrap server).
	// In case the privateIp is provided, the subnetId must also be provided.
	// In case the privateIp (and the subnetId) is not provided it is assumed the datasource is publicly accessible.
	// In case the connection is accessible only privately, the lack of privateIp will result in not being able to access the connection.
	PrivateIp *string `mandatory:"false" json:"privateIp"`

	// The Microsoft SQL Server technology type.
	TechnologyType MicrosoftSqlserverConnectionTechnologyTypeEnum `mandatory:"true" json:"technologyType"`

	// Security Protocol for Microsoft SQL Server.
	SecurityProtocol MicrosoftSqlserverConnectionSecurityProtocolEnum `mandatory:"true" json:"securityProtocol"`

	// Possible lifecycle states for connection.
	LifecycleState ConnectionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Controls the network traffic direction to the target:
	// SHARED_SERVICE_ENDPOINT: Traffic flows through the Goldengate Service's network to public hosts. Cannot be used for private targets.
	// SHARED_DEPLOYMENT_ENDPOINT: Network traffic flows from the assigned deployment's private endpoint through the deployment's subnet.
	// DEDICATED_ENDPOINT: A dedicated private endpoint is created in the target VCN subnet for the connection. The subnetId is required when DEDICATED_ENDPOINT networking is selected.
	RoutingMethod RoutingMethodEnum `mandatory:"false" json:"routingMethod,omitempty"`
}

// GetId returns Id
func (m MicrosoftSqlserverConnection) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m MicrosoftSqlserverConnection) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m MicrosoftSqlserverConnection) GetDescription() *string {
	return m.Description
}

// GetCompartmentId returns CompartmentId
func (m MicrosoftSqlserverConnection) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetFreeformTags returns FreeformTags
func (m MicrosoftSqlserverConnection) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m MicrosoftSqlserverConnection) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m MicrosoftSqlserverConnection) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetLifecycleState returns LifecycleState
func (m MicrosoftSqlserverConnection) GetLifecycleState() ConnectionLifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m MicrosoftSqlserverConnection) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetTimeCreated returns TimeCreated
func (m MicrosoftSqlserverConnection) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m MicrosoftSqlserverConnection) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLocks returns Locks
func (m MicrosoftSqlserverConnection) GetLocks() []ResourceLock {
	return m.Locks
}

// GetVaultId returns VaultId
func (m MicrosoftSqlserverConnection) GetVaultId() *string {
	return m.VaultId
}

// GetKeyId returns KeyId
func (m MicrosoftSqlserverConnection) GetKeyId() *string {
	return m.KeyId
}

// GetIngressIps returns IngressIps
func (m MicrosoftSqlserverConnection) GetIngressIps() []IngressIpDetails {
	return m.IngressIps
}

// GetNsgIds returns NsgIds
func (m MicrosoftSqlserverConnection) GetNsgIds() []string {
	return m.NsgIds
}

// GetSubnetId returns SubnetId
func (m MicrosoftSqlserverConnection) GetSubnetId() *string {
	return m.SubnetId
}

// GetRoutingMethod returns RoutingMethod
func (m MicrosoftSqlserverConnection) GetRoutingMethod() RoutingMethodEnum {
	return m.RoutingMethod
}

func (m MicrosoftSqlserverConnection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MicrosoftSqlserverConnection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMicrosoftSqlserverConnectionTechnologyTypeEnum(string(m.TechnologyType)); !ok && m.TechnologyType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TechnologyType: %s. Supported values are: %s.", m.TechnologyType, strings.Join(GetMicrosoftSqlserverConnectionTechnologyTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMicrosoftSqlserverConnectionSecurityProtocolEnum(string(m.SecurityProtocol)); !ok && m.SecurityProtocol != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SecurityProtocol: %s. Supported values are: %s.", m.SecurityProtocol, strings.Join(GetMicrosoftSqlserverConnectionSecurityProtocolEnumStringValues(), ",")))
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
func (m MicrosoftSqlserverConnection) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeMicrosoftSqlserverConnection MicrosoftSqlserverConnection
	s := struct {
		DiscriminatorParam string `json:"connectionType"`
		MarshalTypeMicrosoftSqlserverConnection
	}{
		"MICROSOFT_SQLSERVER",
		(MarshalTypeMicrosoftSqlserverConnection)(m),
	}

	return json.Marshal(&s)
}

// MicrosoftSqlserverConnectionTechnologyTypeEnum Enum with underlying type: string
type MicrosoftSqlserverConnectionTechnologyTypeEnum string

// Set of constants representing the allowable values for MicrosoftSqlserverConnectionTechnologyTypeEnum
const (
	MicrosoftSqlserverConnectionTechnologyTypeAmazonRdsSqlserver               MicrosoftSqlserverConnectionTechnologyTypeEnum = "AMAZON_RDS_SQLSERVER"
	MicrosoftSqlserverConnectionTechnologyTypeAzureSqlserverManagedInstance    MicrosoftSqlserverConnectionTechnologyTypeEnum = "AZURE_SQLSERVER_MANAGED_INSTANCE"
	MicrosoftSqlserverConnectionTechnologyTypeAzureSqlserverNonManagedInstance MicrosoftSqlserverConnectionTechnologyTypeEnum = "AZURE_SQLSERVER_NON_MANAGED_INSTANCE"
	MicrosoftSqlserverConnectionTechnologyTypeGoogleCloudSqlSqlserver          MicrosoftSqlserverConnectionTechnologyTypeEnum = "GOOGLE_CLOUD_SQL_SQLSERVER"
	MicrosoftSqlserverConnectionTechnologyTypeMicrosoftSqlserver               MicrosoftSqlserverConnectionTechnologyTypeEnum = "MICROSOFT_SQLSERVER"
)

var mappingMicrosoftSqlserverConnectionTechnologyTypeEnum = map[string]MicrosoftSqlserverConnectionTechnologyTypeEnum{
	"AMAZON_RDS_SQLSERVER":                 MicrosoftSqlserverConnectionTechnologyTypeAmazonRdsSqlserver,
	"AZURE_SQLSERVER_MANAGED_INSTANCE":     MicrosoftSqlserverConnectionTechnologyTypeAzureSqlserverManagedInstance,
	"AZURE_SQLSERVER_NON_MANAGED_INSTANCE": MicrosoftSqlserverConnectionTechnologyTypeAzureSqlserverNonManagedInstance,
	"GOOGLE_CLOUD_SQL_SQLSERVER":           MicrosoftSqlserverConnectionTechnologyTypeGoogleCloudSqlSqlserver,
	"MICROSOFT_SQLSERVER":                  MicrosoftSqlserverConnectionTechnologyTypeMicrosoftSqlserver,
}

var mappingMicrosoftSqlserverConnectionTechnologyTypeEnumLowerCase = map[string]MicrosoftSqlserverConnectionTechnologyTypeEnum{
	"amazon_rds_sqlserver":                 MicrosoftSqlserverConnectionTechnologyTypeAmazonRdsSqlserver,
	"azure_sqlserver_managed_instance":     MicrosoftSqlserverConnectionTechnologyTypeAzureSqlserverManagedInstance,
	"azure_sqlserver_non_managed_instance": MicrosoftSqlserverConnectionTechnologyTypeAzureSqlserverNonManagedInstance,
	"google_cloud_sql_sqlserver":           MicrosoftSqlserverConnectionTechnologyTypeGoogleCloudSqlSqlserver,
	"microsoft_sqlserver":                  MicrosoftSqlserverConnectionTechnologyTypeMicrosoftSqlserver,
}

// GetMicrosoftSqlserverConnectionTechnologyTypeEnumValues Enumerates the set of values for MicrosoftSqlserverConnectionTechnologyTypeEnum
func GetMicrosoftSqlserverConnectionTechnologyTypeEnumValues() []MicrosoftSqlserverConnectionTechnologyTypeEnum {
	values := make([]MicrosoftSqlserverConnectionTechnologyTypeEnum, 0)
	for _, v := range mappingMicrosoftSqlserverConnectionTechnologyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMicrosoftSqlserverConnectionTechnologyTypeEnumStringValues Enumerates the set of values in String for MicrosoftSqlserverConnectionTechnologyTypeEnum
func GetMicrosoftSqlserverConnectionTechnologyTypeEnumStringValues() []string {
	return []string{
		"AMAZON_RDS_SQLSERVER",
		"AZURE_SQLSERVER_MANAGED_INSTANCE",
		"AZURE_SQLSERVER_NON_MANAGED_INSTANCE",
		"GOOGLE_CLOUD_SQL_SQLSERVER",
		"MICROSOFT_SQLSERVER",
	}
}

// GetMappingMicrosoftSqlserverConnectionTechnologyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMicrosoftSqlserverConnectionTechnologyTypeEnum(val string) (MicrosoftSqlserverConnectionTechnologyTypeEnum, bool) {
	enum, ok := mappingMicrosoftSqlserverConnectionTechnologyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// MicrosoftSqlserverConnectionSecurityProtocolEnum Enum with underlying type: string
type MicrosoftSqlserverConnectionSecurityProtocolEnum string

// Set of constants representing the allowable values for MicrosoftSqlserverConnectionSecurityProtocolEnum
const (
	MicrosoftSqlserverConnectionSecurityProtocolPlain MicrosoftSqlserverConnectionSecurityProtocolEnum = "PLAIN"
	MicrosoftSqlserverConnectionSecurityProtocolTls   MicrosoftSqlserverConnectionSecurityProtocolEnum = "TLS"
)

var mappingMicrosoftSqlserverConnectionSecurityProtocolEnum = map[string]MicrosoftSqlserverConnectionSecurityProtocolEnum{
	"PLAIN": MicrosoftSqlserverConnectionSecurityProtocolPlain,
	"TLS":   MicrosoftSqlserverConnectionSecurityProtocolTls,
}

var mappingMicrosoftSqlserverConnectionSecurityProtocolEnumLowerCase = map[string]MicrosoftSqlserverConnectionSecurityProtocolEnum{
	"plain": MicrosoftSqlserverConnectionSecurityProtocolPlain,
	"tls":   MicrosoftSqlserverConnectionSecurityProtocolTls,
}

// GetMicrosoftSqlserverConnectionSecurityProtocolEnumValues Enumerates the set of values for MicrosoftSqlserverConnectionSecurityProtocolEnum
func GetMicrosoftSqlserverConnectionSecurityProtocolEnumValues() []MicrosoftSqlserverConnectionSecurityProtocolEnum {
	values := make([]MicrosoftSqlserverConnectionSecurityProtocolEnum, 0)
	for _, v := range mappingMicrosoftSqlserverConnectionSecurityProtocolEnum {
		values = append(values, v)
	}
	return values
}

// GetMicrosoftSqlserverConnectionSecurityProtocolEnumStringValues Enumerates the set of values in String for MicrosoftSqlserverConnectionSecurityProtocolEnum
func GetMicrosoftSqlserverConnectionSecurityProtocolEnumStringValues() []string {
	return []string{
		"PLAIN",
		"TLS",
	}
}

// GetMappingMicrosoftSqlserverConnectionSecurityProtocolEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMicrosoftSqlserverConnectionSecurityProtocolEnum(val string) (MicrosoftSqlserverConnectionSecurityProtocolEnum, bool) {
	enum, ok := mappingMicrosoftSqlserverConnectionSecurityProtocolEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
