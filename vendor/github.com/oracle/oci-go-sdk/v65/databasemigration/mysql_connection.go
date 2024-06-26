// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MysqlConnection Represents the metadata of a MySQL Connection.
type MysqlConnection struct {

	// The OCID of the connection being referenced.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The time when this resource was created.
	// An RFC3339 formatted datetime string such as `2016-08-25T21:10:29.600Z`.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time when this resource was updated.
	// An RFC3339 formatted datetime string such as `2016-08-25T21:10:29.600Z`.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The username (credential) used when creating or updating this resource.
	Username *string `mandatory:"true" json:"username"`

	// A user-friendly description. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags. Example: {"Department": "Finance"}
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The message describing the current state of the connection's lifecycle in detail.
	// For example, can be used to provide actionable information for a connection in a Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// OCI resource ID.
	VaultId *string `mandatory:"false" json:"vaultId"`

	// The OCID of the key used in cryptographic operations.
	KeyId *string `mandatory:"false" json:"keyId"`

	// OCI resource ID.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// List of ingress IP addresses from where to connect to this connection's privateIp.
	IngressIps []IngressIpDetails `mandatory:"false" json:"ingressIps"`

	// An array of Network Security Group OCIDs used to define network access for Connections.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// The password (credential) used when creating or updating this resource.
	Password *string `mandatory:"false" json:"password"`

	// The username (credential) used when creating or updating this resource.
	ReplicationUsername *string `mandatory:"false" json:"replicationUsername"`

	// The password (credential) used when creating or updating this resource.
	ReplicationPassword *string `mandatory:"false" json:"replicationPassword"`

	// The OCID of the resource being referenced.
	SecretId *string `mandatory:"false" json:"secretId"`

	// The OCID of the resource being referenced.
	PrivateEndpointId *string `mandatory:"false" json:"privateEndpointId"`

	// The IP Address of the host.
	Host *string `mandatory:"false" json:"host"`

	// The port to be used for the connection.
	Port *int `mandatory:"false" json:"port"`

	// The name of the database being referenced.
	DatabaseName *string `mandatory:"false" json:"databaseName"`

	// An array of name-value pair attribute entries.
	AdditionalAttributes []NameValuePair `mandatory:"false" json:"additionalAttributes"`

	// The OCID of the database system being referenced.
	DbSystemId *string `mandatory:"false" json:"dbSystemId"`

	// The type of MySQL source or target connection.
	// Example: OCI_MYSQL represents OCI MySQL HeatWave Database Service
	TechnologyType MysqlConnectionTechnologyTypeEnum `mandatory:"true" json:"technologyType"`

	// Security Protocol to be used for the connection.
	SecurityProtocol MysqlConnectionSecurityProtocolEnum `mandatory:"true" json:"securityProtocol"`

	// SSL mode to be used for the connection.
	SslMode MysqlConnectionSslModeEnum `mandatory:"false" json:"sslMode,omitempty"`

	// The Connection's current lifecycle state.
	LifecycleState ConnectionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
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

// GetSubnetId returns SubnetId
func (m MysqlConnection) GetSubnetId() *string {
	return m.SubnetId
}

// GetIngressIps returns IngressIps
func (m MysqlConnection) GetIngressIps() []IngressIpDetails {
	return m.IngressIps
}

// GetNsgIds returns NsgIds
func (m MysqlConnection) GetNsgIds() []string {
	return m.NsgIds
}

// GetUsername returns Username
func (m MysqlConnection) GetUsername() *string {
	return m.Username
}

// GetPassword returns Password
func (m MysqlConnection) GetPassword() *string {
	return m.Password
}

// GetReplicationUsername returns ReplicationUsername
func (m MysqlConnection) GetReplicationUsername() *string {
	return m.ReplicationUsername
}

// GetReplicationPassword returns ReplicationPassword
func (m MysqlConnection) GetReplicationPassword() *string {
	return m.ReplicationPassword
}

// GetSecretId returns SecretId
func (m MysqlConnection) GetSecretId() *string {
	return m.SecretId
}

// GetPrivateEndpointId returns PrivateEndpointId
func (m MysqlConnection) GetPrivateEndpointId() *string {
	return m.PrivateEndpointId
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
	MysqlConnectionTechnologyTypeAmazonRdsMysql      MysqlConnectionTechnologyTypeEnum = "AMAZON_RDS_MYSQL"
	MysqlConnectionTechnologyTypeAzureMysql          MysqlConnectionTechnologyTypeEnum = "AZURE_MYSQL"
	MysqlConnectionTechnologyTypeGoogleCloudSqlMysql MysqlConnectionTechnologyTypeEnum = "GOOGLE_CLOUD_SQL_MYSQL"
	MysqlConnectionTechnologyTypeMysqlServer         MysqlConnectionTechnologyTypeEnum = "MYSQL_SERVER"
	MysqlConnectionTechnologyTypeOciMysql            MysqlConnectionTechnologyTypeEnum = "OCI_MYSQL"
)

var mappingMysqlConnectionTechnologyTypeEnum = map[string]MysqlConnectionTechnologyTypeEnum{
	"AMAZON_AURORA_MYSQL":    MysqlConnectionTechnologyTypeAmazonAuroraMysql,
	"AMAZON_RDS_MYSQL":       MysqlConnectionTechnologyTypeAmazonRdsMysql,
	"AZURE_MYSQL":            MysqlConnectionTechnologyTypeAzureMysql,
	"GOOGLE_CLOUD_SQL_MYSQL": MysqlConnectionTechnologyTypeGoogleCloudSqlMysql,
	"MYSQL_SERVER":           MysqlConnectionTechnologyTypeMysqlServer,
	"OCI_MYSQL":              MysqlConnectionTechnologyTypeOciMysql,
}

var mappingMysqlConnectionTechnologyTypeEnumLowerCase = map[string]MysqlConnectionTechnologyTypeEnum{
	"amazon_aurora_mysql":    MysqlConnectionTechnologyTypeAmazonAuroraMysql,
	"amazon_rds_mysql":       MysqlConnectionTechnologyTypeAmazonRdsMysql,
	"azure_mysql":            MysqlConnectionTechnologyTypeAzureMysql,
	"google_cloud_sql_mysql": MysqlConnectionTechnologyTypeGoogleCloudSqlMysql,
	"mysql_server":           MysqlConnectionTechnologyTypeMysqlServer,
	"oci_mysql":              MysqlConnectionTechnologyTypeOciMysql,
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
		"AMAZON_RDS_MYSQL",
		"AZURE_MYSQL",
		"GOOGLE_CLOUD_SQL_MYSQL",
		"MYSQL_SERVER",
		"OCI_MYSQL",
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
