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

// PostgresqlConnection Represents the metadata of a PostgreSQL Database Connection.
type PostgresqlConnection struct {

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

	// The name of the database.
	DatabaseName *string `mandatory:"true" json:"databaseName"`

	// The name or address of a host.
	Host *string `mandatory:"true" json:"host"`

	// The port of an endpoint usually specified for a connection.
	Port *int `mandatory:"true" json:"port"`

	// The username Oracle GoldenGate uses to connect the associated system of the given technology.
	// This username must already exist and be available by the system/application to be connected to
	// and must conform to the case sensitivty requirments defined in it.
	Username *string `mandatory:"true" json:"username"`

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

	// Deprecated: this field will be removed in future versions. Either specify the private IP in the connectionString or host
	// field, or make sure the host name is resolvable in the target VCN.
	// The private IP address of the connection's endpoint in the customer's VCN, typically a
	// database endpoint or a big data endpoint (e.g. Kafka bootstrap server).
	// In case the privateIp is provided, the subnetId must also be provided.
	// In case the privateIp (and the subnetId) is not provided it is assumed the datasource is publicly accessible.
	// In case the connection is accessible only privately, the lack of privateIp will result in not being able to access the connection.
	PrivateIp *string `mandatory:"false" json:"privateIp"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the database system being referenced.
	DbSystemId *string `mandatory:"false" json:"dbSystemId"`

	// The PostgreSQL technology type.
	TechnologyType PostgresqlConnectionTechnologyTypeEnum `mandatory:"true" json:"technologyType"`

	// Security protocol for PostgreSQL.
	SecurityProtocol PostgresqlConnectionSecurityProtocolEnum `mandatory:"true" json:"securityProtocol"`

	// SSL mode for PostgreSQL.
	SslMode PostgresqlConnectionSslModeEnum `mandatory:"false" json:"sslMode,omitempty"`

	// Possible lifecycle states for connection.
	LifecycleState ConnectionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Controls the network traffic direction to the target:
	// SHARED_SERVICE_ENDPOINT: Traffic flows through the Goldengate Service's network to public hosts. Cannot be used for private targets.
	// SHARED_DEPLOYMENT_ENDPOINT: Network traffic flows from the assigned deployment's private endpoint through the deployment's subnet.
	// DEDICATED_ENDPOINT: A dedicated private endpoint is created in the target VCN subnet for the connection. The subnetId is required when DEDICATED_ENDPOINT networking is selected.
	RoutingMethod RoutingMethodEnum `mandatory:"false" json:"routingMethod,omitempty"`
}

// GetId returns Id
func (m PostgresqlConnection) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m PostgresqlConnection) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m PostgresqlConnection) GetDescription() *string {
	return m.Description
}

// GetCompartmentId returns CompartmentId
func (m PostgresqlConnection) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetFreeformTags returns FreeformTags
func (m PostgresqlConnection) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m PostgresqlConnection) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m PostgresqlConnection) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetLifecycleState returns LifecycleState
func (m PostgresqlConnection) GetLifecycleState() ConnectionLifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m PostgresqlConnection) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetTimeCreated returns TimeCreated
func (m PostgresqlConnection) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m PostgresqlConnection) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLocks returns Locks
func (m PostgresqlConnection) GetLocks() []ResourceLock {
	return m.Locks
}

// GetVaultId returns VaultId
func (m PostgresqlConnection) GetVaultId() *string {
	return m.VaultId
}

// GetKeyId returns KeyId
func (m PostgresqlConnection) GetKeyId() *string {
	return m.KeyId
}

// GetIngressIps returns IngressIps
func (m PostgresqlConnection) GetIngressIps() []IngressIpDetails {
	return m.IngressIps
}

// GetNsgIds returns NsgIds
func (m PostgresqlConnection) GetNsgIds() []string {
	return m.NsgIds
}

// GetSubnetId returns SubnetId
func (m PostgresqlConnection) GetSubnetId() *string {
	return m.SubnetId
}

// GetRoutingMethod returns RoutingMethod
func (m PostgresqlConnection) GetRoutingMethod() RoutingMethodEnum {
	return m.RoutingMethod
}

func (m PostgresqlConnection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PostgresqlConnection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPostgresqlConnectionTechnologyTypeEnum(string(m.TechnologyType)); !ok && m.TechnologyType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TechnologyType: %s. Supported values are: %s.", m.TechnologyType, strings.Join(GetPostgresqlConnectionTechnologyTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPostgresqlConnectionSecurityProtocolEnum(string(m.SecurityProtocol)); !ok && m.SecurityProtocol != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SecurityProtocol: %s. Supported values are: %s.", m.SecurityProtocol, strings.Join(GetPostgresqlConnectionSecurityProtocolEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPostgresqlConnectionSslModeEnum(string(m.SslMode)); !ok && m.SslMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SslMode: %s. Supported values are: %s.", m.SslMode, strings.Join(GetPostgresqlConnectionSslModeEnumStringValues(), ",")))
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
func (m PostgresqlConnection) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePostgresqlConnection PostgresqlConnection
	s := struct {
		DiscriminatorParam string `json:"connectionType"`
		MarshalTypePostgresqlConnection
	}{
		"POSTGRESQL",
		(MarshalTypePostgresqlConnection)(m),
	}

	return json.Marshal(&s)
}

// PostgresqlConnectionTechnologyTypeEnum Enum with underlying type: string
type PostgresqlConnectionTechnologyTypeEnum string

// Set of constants representing the allowable values for PostgresqlConnectionTechnologyTypeEnum
const (
	PostgresqlConnectionTechnologyTypeOciPostgresql              PostgresqlConnectionTechnologyTypeEnum = "OCI_POSTGRESQL"
	PostgresqlConnectionTechnologyTypePostgresqlServer           PostgresqlConnectionTechnologyTypeEnum = "POSTGRESQL_SERVER"
	PostgresqlConnectionTechnologyTypeAmazonAuroraPostgresql     PostgresqlConnectionTechnologyTypeEnum = "AMAZON_AURORA_POSTGRESQL"
	PostgresqlConnectionTechnologyTypeAmazonRdsPostgresql        PostgresqlConnectionTechnologyTypeEnum = "AMAZON_RDS_POSTGRESQL"
	PostgresqlConnectionTechnologyTypeAzurePostgresql            PostgresqlConnectionTechnologyTypeEnum = "AZURE_POSTGRESQL"
	PostgresqlConnectionTechnologyTypeAzureCosmosDbForPostgresql PostgresqlConnectionTechnologyTypeEnum = "AZURE_COSMOS_DB_FOR_POSTGRESQL"
	PostgresqlConnectionTechnologyTypeGoogleCloudSqlPostgresql   PostgresqlConnectionTechnologyTypeEnum = "GOOGLE_CLOUD_SQL_POSTGRESQL"
)

var mappingPostgresqlConnectionTechnologyTypeEnum = map[string]PostgresqlConnectionTechnologyTypeEnum{
	"OCI_POSTGRESQL":                 PostgresqlConnectionTechnologyTypeOciPostgresql,
	"POSTGRESQL_SERVER":              PostgresqlConnectionTechnologyTypePostgresqlServer,
	"AMAZON_AURORA_POSTGRESQL":       PostgresqlConnectionTechnologyTypeAmazonAuroraPostgresql,
	"AMAZON_RDS_POSTGRESQL":          PostgresqlConnectionTechnologyTypeAmazonRdsPostgresql,
	"AZURE_POSTGRESQL":               PostgresqlConnectionTechnologyTypeAzurePostgresql,
	"AZURE_COSMOS_DB_FOR_POSTGRESQL": PostgresqlConnectionTechnologyTypeAzureCosmosDbForPostgresql,
	"GOOGLE_CLOUD_SQL_POSTGRESQL":    PostgresqlConnectionTechnologyTypeGoogleCloudSqlPostgresql,
}

var mappingPostgresqlConnectionTechnologyTypeEnumLowerCase = map[string]PostgresqlConnectionTechnologyTypeEnum{
	"oci_postgresql":                 PostgresqlConnectionTechnologyTypeOciPostgresql,
	"postgresql_server":              PostgresqlConnectionTechnologyTypePostgresqlServer,
	"amazon_aurora_postgresql":       PostgresqlConnectionTechnologyTypeAmazonAuroraPostgresql,
	"amazon_rds_postgresql":          PostgresqlConnectionTechnologyTypeAmazonRdsPostgresql,
	"azure_postgresql":               PostgresqlConnectionTechnologyTypeAzurePostgresql,
	"azure_cosmos_db_for_postgresql": PostgresqlConnectionTechnologyTypeAzureCosmosDbForPostgresql,
	"google_cloud_sql_postgresql":    PostgresqlConnectionTechnologyTypeGoogleCloudSqlPostgresql,
}

// GetPostgresqlConnectionTechnologyTypeEnumValues Enumerates the set of values for PostgresqlConnectionTechnologyTypeEnum
func GetPostgresqlConnectionTechnologyTypeEnumValues() []PostgresqlConnectionTechnologyTypeEnum {
	values := make([]PostgresqlConnectionTechnologyTypeEnum, 0)
	for _, v := range mappingPostgresqlConnectionTechnologyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPostgresqlConnectionTechnologyTypeEnumStringValues Enumerates the set of values in String for PostgresqlConnectionTechnologyTypeEnum
func GetPostgresqlConnectionTechnologyTypeEnumStringValues() []string {
	return []string{
		"OCI_POSTGRESQL",
		"POSTGRESQL_SERVER",
		"AMAZON_AURORA_POSTGRESQL",
		"AMAZON_RDS_POSTGRESQL",
		"AZURE_POSTGRESQL",
		"AZURE_COSMOS_DB_FOR_POSTGRESQL",
		"GOOGLE_CLOUD_SQL_POSTGRESQL",
	}
}

// GetMappingPostgresqlConnectionTechnologyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPostgresqlConnectionTechnologyTypeEnum(val string) (PostgresqlConnectionTechnologyTypeEnum, bool) {
	enum, ok := mappingPostgresqlConnectionTechnologyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// PostgresqlConnectionSecurityProtocolEnum Enum with underlying type: string
type PostgresqlConnectionSecurityProtocolEnum string

// Set of constants representing the allowable values for PostgresqlConnectionSecurityProtocolEnum
const (
	PostgresqlConnectionSecurityProtocolPlain PostgresqlConnectionSecurityProtocolEnum = "PLAIN"
	PostgresqlConnectionSecurityProtocolTls   PostgresqlConnectionSecurityProtocolEnum = "TLS"
	PostgresqlConnectionSecurityProtocolMtls  PostgresqlConnectionSecurityProtocolEnum = "MTLS"
)

var mappingPostgresqlConnectionSecurityProtocolEnum = map[string]PostgresqlConnectionSecurityProtocolEnum{
	"PLAIN": PostgresqlConnectionSecurityProtocolPlain,
	"TLS":   PostgresqlConnectionSecurityProtocolTls,
	"MTLS":  PostgresqlConnectionSecurityProtocolMtls,
}

var mappingPostgresqlConnectionSecurityProtocolEnumLowerCase = map[string]PostgresqlConnectionSecurityProtocolEnum{
	"plain": PostgresqlConnectionSecurityProtocolPlain,
	"tls":   PostgresqlConnectionSecurityProtocolTls,
	"mtls":  PostgresqlConnectionSecurityProtocolMtls,
}

// GetPostgresqlConnectionSecurityProtocolEnumValues Enumerates the set of values for PostgresqlConnectionSecurityProtocolEnum
func GetPostgresqlConnectionSecurityProtocolEnumValues() []PostgresqlConnectionSecurityProtocolEnum {
	values := make([]PostgresqlConnectionSecurityProtocolEnum, 0)
	for _, v := range mappingPostgresqlConnectionSecurityProtocolEnum {
		values = append(values, v)
	}
	return values
}

// GetPostgresqlConnectionSecurityProtocolEnumStringValues Enumerates the set of values in String for PostgresqlConnectionSecurityProtocolEnum
func GetPostgresqlConnectionSecurityProtocolEnumStringValues() []string {
	return []string{
		"PLAIN",
		"TLS",
		"MTLS",
	}
}

// GetMappingPostgresqlConnectionSecurityProtocolEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPostgresqlConnectionSecurityProtocolEnum(val string) (PostgresqlConnectionSecurityProtocolEnum, bool) {
	enum, ok := mappingPostgresqlConnectionSecurityProtocolEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// PostgresqlConnectionSslModeEnum Enum with underlying type: string
type PostgresqlConnectionSslModeEnum string

// Set of constants representing the allowable values for PostgresqlConnectionSslModeEnum
const (
	PostgresqlConnectionSslModePrefer     PostgresqlConnectionSslModeEnum = "PREFER"
	PostgresqlConnectionSslModeRequire    PostgresqlConnectionSslModeEnum = "REQUIRE"
	PostgresqlConnectionSslModeVerifyCa   PostgresqlConnectionSslModeEnum = "VERIFY_CA"
	PostgresqlConnectionSslModeVerifyFull PostgresqlConnectionSslModeEnum = "VERIFY_FULL"
)

var mappingPostgresqlConnectionSslModeEnum = map[string]PostgresqlConnectionSslModeEnum{
	"PREFER":      PostgresqlConnectionSslModePrefer,
	"REQUIRE":     PostgresqlConnectionSslModeRequire,
	"VERIFY_CA":   PostgresqlConnectionSslModeVerifyCa,
	"VERIFY_FULL": PostgresqlConnectionSslModeVerifyFull,
}

var mappingPostgresqlConnectionSslModeEnumLowerCase = map[string]PostgresqlConnectionSslModeEnum{
	"prefer":      PostgresqlConnectionSslModePrefer,
	"require":     PostgresqlConnectionSslModeRequire,
	"verify_ca":   PostgresqlConnectionSslModeVerifyCa,
	"verify_full": PostgresqlConnectionSslModeVerifyFull,
}

// GetPostgresqlConnectionSslModeEnumValues Enumerates the set of values for PostgresqlConnectionSslModeEnum
func GetPostgresqlConnectionSslModeEnumValues() []PostgresqlConnectionSslModeEnum {
	values := make([]PostgresqlConnectionSslModeEnum, 0)
	for _, v := range mappingPostgresqlConnectionSslModeEnum {
		values = append(values, v)
	}
	return values
}

// GetPostgresqlConnectionSslModeEnumStringValues Enumerates the set of values in String for PostgresqlConnectionSslModeEnum
func GetPostgresqlConnectionSslModeEnumStringValues() []string {
	return []string{
		"PREFER",
		"REQUIRE",
		"VERIFY_CA",
		"VERIFY_FULL",
	}
}

// GetMappingPostgresqlConnectionSslModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPostgresqlConnectionSslModeEnum(val string) (PostgresqlConnectionSslModeEnum, bool) {
	enum, ok := mappingPostgresqlConnectionSslModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
