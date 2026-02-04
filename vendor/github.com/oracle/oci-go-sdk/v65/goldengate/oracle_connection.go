// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// OracleConnection Represents the metadata of an Oracle Database Connection.
type OracleConnection struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the connection being
	// referenced.
	Id *string `mandatory:"true" json:"id"`

	// An object's Display Name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment being referenced.
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
	// information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
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

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the target subnet of the dedicated connection.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// Indicates that sensitive attributes are provided via Secrets.
	DoesUseSecretIds *bool `mandatory:"false" json:"doesUseSecretIds"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
	SubscriptionId *string `mandatory:"false" json:"subscriptionId"`

	// The OCID(/Content/General/Concepts/identifiers.htm) of the cluster placement group for the resource.
	// Only applicable for multicloud subscriptions. The cluster placement group id must be provided when a multicloud
	// subscription id is provided. Otherwise the cluster placement group must not be provided.
	ClusterPlacementGroupId *string `mandatory:"false" json:"clusterPlacementGroupId"`

	// Security attributes for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Oracle-ZPR": {"MaxEgressCount": {"value": "42", "mode": "enforce"}}}`
	SecurityAttributes map[string]map[string]interface{} `mandatory:"false" json:"securityAttributes"`

	// Connect descriptor or Easy Connect Naming method used to connect to a database.
	ConnectionString *string `mandatory:"false" json:"connectionString"`

	// This property is not available when creating connections. For existing deprecated connections having this value set, the value cannot be updated; set it to empty.
	// For deprecated connections created with this field in the past, either the private IP had to be specified in the connectionString or host field, or the host name had to be resolvable in the target VCN.
	PrivateIp *string `mandatory:"false" json:"privateIp"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database being referenced.
	DatabaseId *string `mandatory:"false" json:"databaseId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the password is stored.
	// The password Oracle GoldenGate uses to connect the associated system of the given technology.
	// It must conform to the specific security requirements including length, case sensitivity, and so on.
	// If secretId is used plaintext field must not be provided.
	// Note: When provided, 'password' field must not be provided.
	PasswordSecretId *string `mandatory:"false" json:"passwordSecretId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the wallet file is stored.
	// The wallet contents Oracle GoldenGate uses to make connections to a database.
	// Note: When provided, 'wallet' field must not be provided.
	WalletSecretId *string `mandatory:"false" json:"walletSecretId"`

	// The Oracle technology type.
	TechnologyType OracleConnectionTechnologyTypeEnum `mandatory:"true" json:"technologyType"`

	// Authentication mode. It can be provided at creation of Oracle Autonomous Database Serverless connections,
	// when a databaseId is provided. The default value is MTLS.
	AuthenticationMode OracleConnectionAuthenticationModeEnum `mandatory:"false" json:"authenticationMode,omitempty"`

	// Specifies the session mode for the database connection.
	// Use REDIRECT only for RAC databases with SCAN listeners that return IP addresses.
	// For RAC databases with SCAN listeners that return FQDNs, and for all other Oracle database technologies, use DIRECT.
	// In RAC deployments, SCAN listeners redirects a connection to a specific database node, identified by either IP address or FQDN.
	// It is recommended to configure RAC with FQDN-based SCAN listeners.
	// The default is DIRECT, except when databaseId is provided and the discovered database relies on the SCAN listener.
	// In this case, the default is REDIRECT.
	// Deprecated: Defaulting to the REDIRECT session mode will be removed after March 1, 2027.
	SessionMode OracleConnectionSessionModeEnum `mandatory:"false" json:"sessionMode,omitempty"`

	// Possible lifecycle states for connection.
	LifecycleState ConnectionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Controls the network traffic direction to the target:
	// SHARED_SERVICE_ENDPOINT: Traffic flows through the Goldengate Service's network to public hosts. Cannot be used for private targets.
	// SHARED_DEPLOYMENT_ENDPOINT: Network traffic flows from the assigned deployment's private endpoint through the deployment's subnet.
	// DEDICATED_ENDPOINT: A dedicated private endpoint is created in the target VCN subnet for the connection. The subnetId is required when DEDICATED_ENDPOINT networking is selected.
	RoutingMethod RoutingMethodEnum `mandatory:"false" json:"routingMethod,omitempty"`
}

// GetId returns Id
func (m OracleConnection) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m OracleConnection) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m OracleConnection) GetDescription() *string {
	return m.Description
}

// GetCompartmentId returns CompartmentId
func (m OracleConnection) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetFreeformTags returns FreeformTags
func (m OracleConnection) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m OracleConnection) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m OracleConnection) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetLifecycleState returns LifecycleState
func (m OracleConnection) GetLifecycleState() ConnectionLifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m OracleConnection) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetTimeCreated returns TimeCreated
func (m OracleConnection) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m OracleConnection) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLocks returns Locks
func (m OracleConnection) GetLocks() []ResourceLock {
	return m.Locks
}

// GetVaultId returns VaultId
func (m OracleConnection) GetVaultId() *string {
	return m.VaultId
}

// GetKeyId returns KeyId
func (m OracleConnection) GetKeyId() *string {
	return m.KeyId
}

// GetIngressIps returns IngressIps
func (m OracleConnection) GetIngressIps() []IngressIpDetails {
	return m.IngressIps
}

// GetNsgIds returns NsgIds
func (m OracleConnection) GetNsgIds() []string {
	return m.NsgIds
}

// GetSubnetId returns SubnetId
func (m OracleConnection) GetSubnetId() *string {
	return m.SubnetId
}

// GetRoutingMethod returns RoutingMethod
func (m OracleConnection) GetRoutingMethod() RoutingMethodEnum {
	return m.RoutingMethod
}

// GetDoesUseSecretIds returns DoesUseSecretIds
func (m OracleConnection) GetDoesUseSecretIds() *bool {
	return m.DoesUseSecretIds
}

// GetSubscriptionId returns SubscriptionId
func (m OracleConnection) GetSubscriptionId() *string {
	return m.SubscriptionId
}

// GetClusterPlacementGroupId returns ClusterPlacementGroupId
func (m OracleConnection) GetClusterPlacementGroupId() *string {
	return m.ClusterPlacementGroupId
}

// GetSecurityAttributes returns SecurityAttributes
func (m OracleConnection) GetSecurityAttributes() map[string]map[string]interface{} {
	return m.SecurityAttributes
}

func (m OracleConnection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OracleConnection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOracleConnectionTechnologyTypeEnum(string(m.TechnologyType)); !ok && m.TechnologyType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TechnologyType: %s. Supported values are: %s.", m.TechnologyType, strings.Join(GetOracleConnectionTechnologyTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOracleConnectionAuthenticationModeEnum(string(m.AuthenticationMode)); !ok && m.AuthenticationMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AuthenticationMode: %s. Supported values are: %s.", m.AuthenticationMode, strings.Join(GetOracleConnectionAuthenticationModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOracleConnectionSessionModeEnum(string(m.SessionMode)); !ok && m.SessionMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SessionMode: %s. Supported values are: %s.", m.SessionMode, strings.Join(GetOracleConnectionSessionModeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingConnectionLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetConnectionLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRoutingMethodEnum(string(m.RoutingMethod)); !ok && m.RoutingMethod != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RoutingMethod: %s. Supported values are: %s.", m.RoutingMethod, strings.Join(GetRoutingMethodEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m OracleConnection) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOracleConnection OracleConnection
	s := struct {
		DiscriminatorParam string `json:"connectionType"`
		MarshalTypeOracleConnection
	}{
		"ORACLE",
		(MarshalTypeOracleConnection)(m),
	}

	return json.Marshal(&s)
}

// OracleConnectionTechnologyTypeEnum Enum with underlying type: string
type OracleConnectionTechnologyTypeEnum string

// Set of constants representing the allowable values for OracleConnectionTechnologyTypeEnum
const (
	OracleConnectionTechnologyTypeAmazonRdsOracle                       OracleConnectionTechnologyTypeEnum = "AMAZON_RDS_ORACLE"
	OracleConnectionTechnologyTypeOciAutonomousDatabase                 OracleConnectionTechnologyTypeEnum = "OCI_AUTONOMOUS_DATABASE"
	OracleConnectionTechnologyTypeOracleDatabase                        OracleConnectionTechnologyTypeEnum = "ORACLE_DATABASE"
	OracleConnectionTechnologyTypeOracleExadata                         OracleConnectionTechnologyTypeEnum = "ORACLE_EXADATA"
	OracleConnectionTechnologyTypeOracleExadataDatabaseAtAzure          OracleConnectionTechnologyTypeEnum = "ORACLE_EXADATA_DATABASE_AT_AZURE"
	OracleConnectionTechnologyTypeOracleAutonomousDatabaseAtAzure       OracleConnectionTechnologyTypeEnum = "ORACLE_AUTONOMOUS_DATABASE_AT_AZURE"
	OracleConnectionTechnologyTypeOracleExadataDatabaseAtGoogleCloud    OracleConnectionTechnologyTypeEnum = "ORACLE_EXADATA_DATABASE_AT_GOOGLE_CLOUD"
	OracleConnectionTechnologyTypeOracleAutonomousDatabaseAtGoogleCloud OracleConnectionTechnologyTypeEnum = "ORACLE_AUTONOMOUS_DATABASE_AT_GOOGLE_CLOUD"
	OracleConnectionTechnologyTypeOracleExadataDatabaseAtAws            OracleConnectionTechnologyTypeEnum = "ORACLE_EXADATA_DATABASE_AT_AWS"
	OracleConnectionTechnologyTypeOracleAutonomousDatabaseAtAws         OracleConnectionTechnologyTypeEnum = "ORACLE_AUTONOMOUS_DATABASE_AT_AWS"
)

var mappingOracleConnectionTechnologyTypeEnum = map[string]OracleConnectionTechnologyTypeEnum{
	"AMAZON_RDS_ORACLE":                          OracleConnectionTechnologyTypeAmazonRdsOracle,
	"OCI_AUTONOMOUS_DATABASE":                    OracleConnectionTechnologyTypeOciAutonomousDatabase,
	"ORACLE_DATABASE":                            OracleConnectionTechnologyTypeOracleDatabase,
	"ORACLE_EXADATA":                             OracleConnectionTechnologyTypeOracleExadata,
	"ORACLE_EXADATA_DATABASE_AT_AZURE":           OracleConnectionTechnologyTypeOracleExadataDatabaseAtAzure,
	"ORACLE_AUTONOMOUS_DATABASE_AT_AZURE":        OracleConnectionTechnologyTypeOracleAutonomousDatabaseAtAzure,
	"ORACLE_EXADATA_DATABASE_AT_GOOGLE_CLOUD":    OracleConnectionTechnologyTypeOracleExadataDatabaseAtGoogleCloud,
	"ORACLE_AUTONOMOUS_DATABASE_AT_GOOGLE_CLOUD": OracleConnectionTechnologyTypeOracleAutonomousDatabaseAtGoogleCloud,
	"ORACLE_EXADATA_DATABASE_AT_AWS":             OracleConnectionTechnologyTypeOracleExadataDatabaseAtAws,
	"ORACLE_AUTONOMOUS_DATABASE_AT_AWS":          OracleConnectionTechnologyTypeOracleAutonomousDatabaseAtAws,
}

var mappingOracleConnectionTechnologyTypeEnumLowerCase = map[string]OracleConnectionTechnologyTypeEnum{
	"amazon_rds_oracle":                          OracleConnectionTechnologyTypeAmazonRdsOracle,
	"oci_autonomous_database":                    OracleConnectionTechnologyTypeOciAutonomousDatabase,
	"oracle_database":                            OracleConnectionTechnologyTypeOracleDatabase,
	"oracle_exadata":                             OracleConnectionTechnologyTypeOracleExadata,
	"oracle_exadata_database_at_azure":           OracleConnectionTechnologyTypeOracleExadataDatabaseAtAzure,
	"oracle_autonomous_database_at_azure":        OracleConnectionTechnologyTypeOracleAutonomousDatabaseAtAzure,
	"oracle_exadata_database_at_google_cloud":    OracleConnectionTechnologyTypeOracleExadataDatabaseAtGoogleCloud,
	"oracle_autonomous_database_at_google_cloud": OracleConnectionTechnologyTypeOracleAutonomousDatabaseAtGoogleCloud,
	"oracle_exadata_database_at_aws":             OracleConnectionTechnologyTypeOracleExadataDatabaseAtAws,
	"oracle_autonomous_database_at_aws":          OracleConnectionTechnologyTypeOracleAutonomousDatabaseAtAws,
}

// GetOracleConnectionTechnologyTypeEnumValues Enumerates the set of values for OracleConnectionTechnologyTypeEnum
func GetOracleConnectionTechnologyTypeEnumValues() []OracleConnectionTechnologyTypeEnum {
	values := make([]OracleConnectionTechnologyTypeEnum, 0)
	for _, v := range mappingOracleConnectionTechnologyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOracleConnectionTechnologyTypeEnumStringValues Enumerates the set of values in String for OracleConnectionTechnologyTypeEnum
func GetOracleConnectionTechnologyTypeEnumStringValues() []string {
	return []string{
		"AMAZON_RDS_ORACLE",
		"OCI_AUTONOMOUS_DATABASE",
		"ORACLE_DATABASE",
		"ORACLE_EXADATA",
		"ORACLE_EXADATA_DATABASE_AT_AZURE",
		"ORACLE_AUTONOMOUS_DATABASE_AT_AZURE",
		"ORACLE_EXADATA_DATABASE_AT_GOOGLE_CLOUD",
		"ORACLE_AUTONOMOUS_DATABASE_AT_GOOGLE_CLOUD",
		"ORACLE_EXADATA_DATABASE_AT_AWS",
		"ORACLE_AUTONOMOUS_DATABASE_AT_AWS",
	}
}

// GetMappingOracleConnectionTechnologyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOracleConnectionTechnologyTypeEnum(val string) (OracleConnectionTechnologyTypeEnum, bool) {
	enum, ok := mappingOracleConnectionTechnologyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// OracleConnectionAuthenticationModeEnum Enum with underlying type: string
type OracleConnectionAuthenticationModeEnum string

// Set of constants representing the allowable values for OracleConnectionAuthenticationModeEnum
const (
	OracleConnectionAuthenticationModeTls  OracleConnectionAuthenticationModeEnum = "TLS"
	OracleConnectionAuthenticationModeMtls OracleConnectionAuthenticationModeEnum = "MTLS"
)

var mappingOracleConnectionAuthenticationModeEnum = map[string]OracleConnectionAuthenticationModeEnum{
	"TLS":  OracleConnectionAuthenticationModeTls,
	"MTLS": OracleConnectionAuthenticationModeMtls,
}

var mappingOracleConnectionAuthenticationModeEnumLowerCase = map[string]OracleConnectionAuthenticationModeEnum{
	"tls":  OracleConnectionAuthenticationModeTls,
	"mtls": OracleConnectionAuthenticationModeMtls,
}

// GetOracleConnectionAuthenticationModeEnumValues Enumerates the set of values for OracleConnectionAuthenticationModeEnum
func GetOracleConnectionAuthenticationModeEnumValues() []OracleConnectionAuthenticationModeEnum {
	values := make([]OracleConnectionAuthenticationModeEnum, 0)
	for _, v := range mappingOracleConnectionAuthenticationModeEnum {
		values = append(values, v)
	}
	return values
}

// GetOracleConnectionAuthenticationModeEnumStringValues Enumerates the set of values in String for OracleConnectionAuthenticationModeEnum
func GetOracleConnectionAuthenticationModeEnumStringValues() []string {
	return []string{
		"TLS",
		"MTLS",
	}
}

// GetMappingOracleConnectionAuthenticationModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOracleConnectionAuthenticationModeEnum(val string) (OracleConnectionAuthenticationModeEnum, bool) {
	enum, ok := mappingOracleConnectionAuthenticationModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// OracleConnectionSessionModeEnum Enum with underlying type: string
type OracleConnectionSessionModeEnum string

// Set of constants representing the allowable values for OracleConnectionSessionModeEnum
const (
	OracleConnectionSessionModeDirect   OracleConnectionSessionModeEnum = "DIRECT"
	OracleConnectionSessionModeRedirect OracleConnectionSessionModeEnum = "REDIRECT"
)

var mappingOracleConnectionSessionModeEnum = map[string]OracleConnectionSessionModeEnum{
	"DIRECT":   OracleConnectionSessionModeDirect,
	"REDIRECT": OracleConnectionSessionModeRedirect,
}

var mappingOracleConnectionSessionModeEnumLowerCase = map[string]OracleConnectionSessionModeEnum{
	"direct":   OracleConnectionSessionModeDirect,
	"redirect": OracleConnectionSessionModeRedirect,
}

// GetOracleConnectionSessionModeEnumValues Enumerates the set of values for OracleConnectionSessionModeEnum
func GetOracleConnectionSessionModeEnumValues() []OracleConnectionSessionModeEnum {
	values := make([]OracleConnectionSessionModeEnum, 0)
	for _, v := range mappingOracleConnectionSessionModeEnum {
		values = append(values, v)
	}
	return values
}

// GetOracleConnectionSessionModeEnumStringValues Enumerates the set of values in String for OracleConnectionSessionModeEnum
func GetOracleConnectionSessionModeEnumStringValues() []string {
	return []string{
		"DIRECT",
		"REDIRECT",
	}
}

// GetMappingOracleConnectionSessionModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOracleConnectionSessionModeEnum(val string) (OracleConnectionSessionModeEnum, bool) {
	enum, ok := mappingOracleConnectionSessionModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
