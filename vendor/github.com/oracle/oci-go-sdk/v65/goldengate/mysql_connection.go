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

// MysqlConnection Represents the metadata of a MySQL Connection.
type MysqlConnection struct {

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

	// The name or address of a host.
	Host *string `mandatory:"false" json:"host"`

	// The port of an endpoint usually specified for a connection.
	Port *int `mandatory:"false" json:"port"`

	// The name of the database.
	DatabaseName *string `mandatory:"false" json:"databaseName"`

	// Deprecated: this field will be removed in future versions. Either specify the private IP in the connectionString or host
	// field, or make sure the host name is resolvable in the target VCN.
	// The private IP address of the connection's endpoint in the customer's VCN, typically a
	// database endpoint or a big data endpoint (e.g. Kafka bootstrap server).
	// In case the privateIp is provided, the subnetId must also be provided.
	// In case the privateIp (and the subnetId) is not provided it is assumed the datasource is publicly accessible.
	// In case the connection is accessible only privately, the lack of privateIp will result in not being able to access the connection.
	PrivateIp *string `mandatory:"false" json:"privateIp"`

	// An array of name-value pair attribute entries.
	// Used as additional parameters in connection string.
	AdditionalAttributes []NameValuePair `mandatory:"false" json:"additionalAttributes"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the database system being referenced.
	DbSystemId *string `mandatory:"false" json:"dbSystemId"`

	// The MySQL technology type.
	TechnologyType MysqlConnectionTechnologyTypeEnum `mandatory:"true" json:"technologyType"`

	// Security Protocol for MySQL.
	SecurityProtocol MysqlConnectionSecurityProtocolEnum `mandatory:"true" json:"securityProtocol"`

	// SSL modes for MySQL.
	SslMode MysqlConnectionSslModeEnum `mandatory:"false" json:"sslMode,omitempty"`

	// Possible lifecycle states for connection.
	LifecycleState ConnectionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Controls the network traffic direction to the target:
	// SHARED_SERVICE_ENDPOINT: Traffic flows through the Goldengate Service's network to public hosts. Cannot be used for private targets.
	// SHARED_DEPLOYMENT_ENDPOINT: Network traffic flows from the assigned deployment's private endpoint through the deployment's subnet.
	// DEDICATED_ENDPOINT: A dedicated private endpoint is created in the target VCN subnet for the connection. The subnetId is required when DEDICATED_ENDPOINT networking is selected.
	RoutingMethod RoutingMethodEnum `mandatory:"false" json:"routingMethod,omitempty"`
}

// GetId returns Id
func (m MysqlConnection) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m MysqlConnection) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m MysqlConnection) GetDescription() *string {
	return m.Description
}

// GetCompartmentId returns CompartmentId
func (m MysqlConnection) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetFreeformTags returns FreeformTags
func (m MysqlConnection) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m MysqlConnection) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m MysqlConnection) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetLifecycleState returns LifecycleState
func (m MysqlConnection) GetLifecycleState() ConnectionLifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m MysqlConnection) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetTimeCreated returns TimeCreated
func (m MysqlConnection) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m MysqlConnection) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetVaultId returns VaultId
func (m MysqlConnection) GetVaultId() *string {
	return m.VaultId
}

// GetKeyId returns KeyId
func (m MysqlConnection) GetKeyId() *string {
	return m.KeyId
}

// GetIngressIps returns IngressIps
func (m MysqlConnection) GetIngressIps() []IngressIpDetails {
	return m.IngressIps
}

// GetNsgIds returns NsgIds
func (m MysqlConnection) GetNsgIds() []string {
	return m.NsgIds
}

// GetSubnetId returns SubnetId
func (m MysqlConnection) GetSubnetId() *string {
	return m.SubnetId
}

// GetRoutingMethod returns RoutingMethod
func (m MysqlConnection) GetRoutingMethod() RoutingMethodEnum {
	return m.RoutingMethod
}

func (m MysqlConnection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MysqlConnection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMysqlConnectionTechnologyTypeEnum(string(m.TechnologyType)); !ok && m.TechnologyType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TechnologyType: %s. Supported values are: %s.", m.TechnologyType, strings.Join(GetMysqlConnectionTechnologyTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMysqlConnectionSecurityProtocolEnum(string(m.SecurityProtocol)); !ok && m.SecurityProtocol != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SecurityProtocol: %s. Supported values are: %s.", m.SecurityProtocol, strings.Join(GetMysqlConnectionSecurityProtocolEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMysqlConnectionSslModeEnum(string(m.SslMode)); !ok && m.SslMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SslMode: %s. Supported values are: %s.", m.SslMode, strings.Join(GetMysqlConnectionSslModeEnumStringValues(), ",")))
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
func (m MysqlConnection) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeMysqlConnection MysqlConnection
	s := struct {
		DiscriminatorParam string `json:"connectionType"`
		MarshalTypeMysqlConnection
	}{
		"MYSQL",
		(MarshalTypeMysqlConnection)(m),
	}

	return json.Marshal(&s)
}

// MysqlConnectionTechnologyTypeEnum Enum with underlying type: string
type MysqlConnectionTechnologyTypeEnum string

// Set of constants representing the allowable values for MysqlConnectionTechnologyTypeEnum
const (
	MysqlConnectionTechnologyTypeAmazonAuroraMysql   MysqlConnectionTechnologyTypeEnum = "AMAZON_AURORA_MYSQL"
	MysqlConnectionTechnologyTypeAmazonRdsMariadb    MysqlConnectionTechnologyTypeEnum = "AMAZON_RDS_MARIADB"
	MysqlConnectionTechnologyTypeAmazonRdsMysql      MysqlConnectionTechnologyTypeEnum = "AMAZON_RDS_MYSQL"
	MysqlConnectionTechnologyTypeAzureMysql          MysqlConnectionTechnologyTypeEnum = "AZURE_MYSQL"
	MysqlConnectionTechnologyTypeGoogleCloudSqlMysql MysqlConnectionTechnologyTypeEnum = "GOOGLE_CLOUD_SQL_MYSQL"
	MysqlConnectionTechnologyTypeMariadb             MysqlConnectionTechnologyTypeEnum = "MARIADB"
	MysqlConnectionTechnologyTypeMysqlServer         MysqlConnectionTechnologyTypeEnum = "MYSQL_SERVER"
	MysqlConnectionTechnologyTypeOciMysql            MysqlConnectionTechnologyTypeEnum = "OCI_MYSQL"
	MysqlConnectionTechnologyTypeSinglestoredb       MysqlConnectionTechnologyTypeEnum = "SINGLESTOREDB"
	MysqlConnectionTechnologyTypeSinglestoredbCloud  MysqlConnectionTechnologyTypeEnum = "SINGLESTOREDB_CLOUD"
)

var mappingMysqlConnectionTechnologyTypeEnum = map[string]MysqlConnectionTechnologyTypeEnum{
	"AMAZON_AURORA_MYSQL":    MysqlConnectionTechnologyTypeAmazonAuroraMysql,
	"AMAZON_RDS_MARIADB":     MysqlConnectionTechnologyTypeAmazonRdsMariadb,
	"AMAZON_RDS_MYSQL":       MysqlConnectionTechnologyTypeAmazonRdsMysql,
	"AZURE_MYSQL":            MysqlConnectionTechnologyTypeAzureMysql,
	"GOOGLE_CLOUD_SQL_MYSQL": MysqlConnectionTechnologyTypeGoogleCloudSqlMysql,
	"MARIADB":                MysqlConnectionTechnologyTypeMariadb,
	"MYSQL_SERVER":           MysqlConnectionTechnologyTypeMysqlServer,
	"OCI_MYSQL":              MysqlConnectionTechnologyTypeOciMysql,
	"SINGLESTOREDB":          MysqlConnectionTechnologyTypeSinglestoredb,
	"SINGLESTOREDB_CLOUD":    MysqlConnectionTechnologyTypeSinglestoredbCloud,
}

var mappingMysqlConnectionTechnologyTypeEnumLowerCase = map[string]MysqlConnectionTechnologyTypeEnum{
	"amazon_aurora_mysql":    MysqlConnectionTechnologyTypeAmazonAuroraMysql,
	"amazon_rds_mariadb":     MysqlConnectionTechnologyTypeAmazonRdsMariadb,
	"amazon_rds_mysql":       MysqlConnectionTechnologyTypeAmazonRdsMysql,
	"azure_mysql":            MysqlConnectionTechnologyTypeAzureMysql,
	"google_cloud_sql_mysql": MysqlConnectionTechnologyTypeGoogleCloudSqlMysql,
	"mariadb":                MysqlConnectionTechnologyTypeMariadb,
	"mysql_server":           MysqlConnectionTechnologyTypeMysqlServer,
	"oci_mysql":              MysqlConnectionTechnologyTypeOciMysql,
	"singlestoredb":          MysqlConnectionTechnologyTypeSinglestoredb,
	"singlestoredb_cloud":    MysqlConnectionTechnologyTypeSinglestoredbCloud,
}

// GetMysqlConnectionTechnologyTypeEnumValues Enumerates the set of values for MysqlConnectionTechnologyTypeEnum
func GetMysqlConnectionTechnologyTypeEnumValues() []MysqlConnectionTechnologyTypeEnum {
	values := make([]MysqlConnectionTechnologyTypeEnum, 0)
	for _, v := range mappingMysqlConnectionTechnologyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMysqlConnectionTechnologyTypeEnumStringValues Enumerates the set of values in String for MysqlConnectionTechnologyTypeEnum
func GetMysqlConnectionTechnologyTypeEnumStringValues() []string {
	return []string{
		"AMAZON_AURORA_MYSQL",
		"AMAZON_RDS_MARIADB",
		"AMAZON_RDS_MYSQL",
		"AZURE_MYSQL",
		"GOOGLE_CLOUD_SQL_MYSQL",
		"MARIADB",
		"MYSQL_SERVER",
		"OCI_MYSQL",
		"SINGLESTOREDB",
		"SINGLESTOREDB_CLOUD",
	}
}

// GetMappingMysqlConnectionTechnologyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMysqlConnectionTechnologyTypeEnum(val string) (MysqlConnectionTechnologyTypeEnum, bool) {
	enum, ok := mappingMysqlConnectionTechnologyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// MysqlConnectionSecurityProtocolEnum Enum with underlying type: string
type MysqlConnectionSecurityProtocolEnum string

// Set of constants representing the allowable values for MysqlConnectionSecurityProtocolEnum
const (
	MysqlConnectionSecurityProtocolPlain MysqlConnectionSecurityProtocolEnum = "PLAIN"
	MysqlConnectionSecurityProtocolTls   MysqlConnectionSecurityProtocolEnum = "TLS"
	MysqlConnectionSecurityProtocolMtls  MysqlConnectionSecurityProtocolEnum = "MTLS"
)

var mappingMysqlConnectionSecurityProtocolEnum = map[string]MysqlConnectionSecurityProtocolEnum{
	"PLAIN": MysqlConnectionSecurityProtocolPlain,
	"TLS":   MysqlConnectionSecurityProtocolTls,
	"MTLS":  MysqlConnectionSecurityProtocolMtls,
}

var mappingMysqlConnectionSecurityProtocolEnumLowerCase = map[string]MysqlConnectionSecurityProtocolEnum{
	"plain": MysqlConnectionSecurityProtocolPlain,
	"tls":   MysqlConnectionSecurityProtocolTls,
	"mtls":  MysqlConnectionSecurityProtocolMtls,
}

// GetMysqlConnectionSecurityProtocolEnumValues Enumerates the set of values for MysqlConnectionSecurityProtocolEnum
func GetMysqlConnectionSecurityProtocolEnumValues() []MysqlConnectionSecurityProtocolEnum {
	values := make([]MysqlConnectionSecurityProtocolEnum, 0)
	for _, v := range mappingMysqlConnectionSecurityProtocolEnum {
		values = append(values, v)
	}
	return values
}

// GetMysqlConnectionSecurityProtocolEnumStringValues Enumerates the set of values in String for MysqlConnectionSecurityProtocolEnum
func GetMysqlConnectionSecurityProtocolEnumStringValues() []string {
	return []string{
		"PLAIN",
		"TLS",
		"MTLS",
	}
}

// GetMappingMysqlConnectionSecurityProtocolEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMysqlConnectionSecurityProtocolEnum(val string) (MysqlConnectionSecurityProtocolEnum, bool) {
	enum, ok := mappingMysqlConnectionSecurityProtocolEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// MysqlConnectionSslModeEnum Enum with underlying type: string
type MysqlConnectionSslModeEnum string

// Set of constants representing the allowable values for MysqlConnectionSslModeEnum
const (
	MysqlConnectionSslModeDisabled       MysqlConnectionSslModeEnum = "DISABLED"
	MysqlConnectionSslModePreferred      MysqlConnectionSslModeEnum = "PREFERRED"
	MysqlConnectionSslModeRequired       MysqlConnectionSslModeEnum = "REQUIRED"
	MysqlConnectionSslModeVerifyCa       MysqlConnectionSslModeEnum = "VERIFY_CA"
	MysqlConnectionSslModeVerifyIdentity MysqlConnectionSslModeEnum = "VERIFY_IDENTITY"
)

var mappingMysqlConnectionSslModeEnum = map[string]MysqlConnectionSslModeEnum{
	"DISABLED":        MysqlConnectionSslModeDisabled,
	"PREFERRED":       MysqlConnectionSslModePreferred,
	"REQUIRED":        MysqlConnectionSslModeRequired,
	"VERIFY_CA":       MysqlConnectionSslModeVerifyCa,
	"VERIFY_IDENTITY": MysqlConnectionSslModeVerifyIdentity,
}

var mappingMysqlConnectionSslModeEnumLowerCase = map[string]MysqlConnectionSslModeEnum{
	"disabled":        MysqlConnectionSslModeDisabled,
	"preferred":       MysqlConnectionSslModePreferred,
	"required":        MysqlConnectionSslModeRequired,
	"verify_ca":       MysqlConnectionSslModeVerifyCa,
	"verify_identity": MysqlConnectionSslModeVerifyIdentity,
}

// GetMysqlConnectionSslModeEnumValues Enumerates the set of values for MysqlConnectionSslModeEnum
func GetMysqlConnectionSslModeEnumValues() []MysqlConnectionSslModeEnum {
	values := make([]MysqlConnectionSslModeEnum, 0)
	for _, v := range mappingMysqlConnectionSslModeEnum {
		values = append(values, v)
	}
	return values
}

// GetMysqlConnectionSslModeEnumStringValues Enumerates the set of values in String for MysqlConnectionSslModeEnum
func GetMysqlConnectionSslModeEnumStringValues() []string {
	return []string{
		"DISABLED",
		"PREFERRED",
		"REQUIRED",
		"VERIFY_CA",
		"VERIFY_IDENTITY",
	}
}

// GetMappingMysqlConnectionSslModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMysqlConnectionSslModeEnum(val string) (MysqlConnectionSslModeEnum, bool) {
	enum, ok := mappingMysqlConnectionSslModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
