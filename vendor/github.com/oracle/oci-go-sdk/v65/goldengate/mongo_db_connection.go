// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// MongoDbConnection Represents the metadata of a MongoDB Connection.
type MongoDbConnection struct {

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

	// MongoDB connection string.
	// e.g.: 'mongodb://mongodb0.example.com:27017/recordsrecords'
	ConnectionString *string `mandatory:"false" json:"connectionString"`

	// The username Oracle GoldenGate uses to connect to the database.
	// This username must already exist and be available by the database to be connected to.
	Username *string `mandatory:"false" json:"username"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Autonomous Json Database.
	DatabaseId *string `mandatory:"false" json:"databaseId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the password Oracle GoldenGate uses to connect the associated database.
	// Note: When provided, 'password' field must not be provided.
	PasswordSecretId *string `mandatory:"false" json:"passwordSecretId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the certificate key file of the mtls connection.
	// - The content of a .pem file containing the client private key (for 2-way SSL).
	// Note: When provided, 'tlsCertificateKeyFile' field must not be provided.
	TlsCertificateKeyFileSecretId *string `mandatory:"false" json:"tlsCertificateKeyFileSecretId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the password of the tls certificate key file.
	// Note: When provided, 'tlsCertificateKeyFilePassword' field must not be provided.
	TlsCertificateKeyFilePasswordSecretId *string `mandatory:"false" json:"tlsCertificateKeyFilePasswordSecretId"`

	// Database Certificate - The base64 encoded content of a .pem file, containing the server public key (for 1 and 2-way SSL).
	// It is not included in GET responses if the `view=COMPACT` query parameter is specified.
	TlsCaFile *string `mandatory:"false" json:"tlsCaFile"`

	// The MongoDB technology type.
	TechnologyType MongoDbConnectionTechnologyTypeEnum `mandatory:"true" json:"technologyType"`

	// Security Protocol for MongoDB.
	SecurityProtocol MongoDbConnectionSecurityProtocolEnum `mandatory:"false" json:"securityProtocol,omitempty"`

	// Possible lifecycle states for connection.
	LifecycleState ConnectionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Controls the network traffic direction to the target:
	// SHARED_SERVICE_ENDPOINT: Traffic flows through the Goldengate Service's network to public hosts. Cannot be used for private targets.
	// SHARED_DEPLOYMENT_ENDPOINT: Network traffic flows from the assigned deployment's private endpoint through the deployment's subnet.
	// DEDICATED_ENDPOINT: A dedicated private endpoint is created in the target VCN subnet for the connection. The subnetId is required when DEDICATED_ENDPOINT networking is selected.
	RoutingMethod RoutingMethodEnum `mandatory:"false" json:"routingMethod,omitempty"`
}

// GetId returns Id
func (m MongoDbConnection) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m MongoDbConnection) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m MongoDbConnection) GetDescription() *string {
	return m.Description
}

// GetCompartmentId returns CompartmentId
func (m MongoDbConnection) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetFreeformTags returns FreeformTags
func (m MongoDbConnection) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m MongoDbConnection) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m MongoDbConnection) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetLifecycleState returns LifecycleState
func (m MongoDbConnection) GetLifecycleState() ConnectionLifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m MongoDbConnection) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetTimeCreated returns TimeCreated
func (m MongoDbConnection) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m MongoDbConnection) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLocks returns Locks
func (m MongoDbConnection) GetLocks() []ResourceLock {
	return m.Locks
}

// GetVaultId returns VaultId
func (m MongoDbConnection) GetVaultId() *string {
	return m.VaultId
}

// GetKeyId returns KeyId
func (m MongoDbConnection) GetKeyId() *string {
	return m.KeyId
}

// GetIngressIps returns IngressIps
func (m MongoDbConnection) GetIngressIps() []IngressIpDetails {
	return m.IngressIps
}

// GetNsgIds returns NsgIds
func (m MongoDbConnection) GetNsgIds() []string {
	return m.NsgIds
}

// GetSubnetId returns SubnetId
func (m MongoDbConnection) GetSubnetId() *string {
	return m.SubnetId
}

// GetRoutingMethod returns RoutingMethod
func (m MongoDbConnection) GetRoutingMethod() RoutingMethodEnum {
	return m.RoutingMethod
}

// GetDoesUseSecretIds returns DoesUseSecretIds
func (m MongoDbConnection) GetDoesUseSecretIds() *bool {
	return m.DoesUseSecretIds
}

// GetSubscriptionId returns SubscriptionId
func (m MongoDbConnection) GetSubscriptionId() *string {
	return m.SubscriptionId
}

// GetClusterPlacementGroupId returns ClusterPlacementGroupId
func (m MongoDbConnection) GetClusterPlacementGroupId() *string {
	return m.ClusterPlacementGroupId
}

// GetSecurityAttributes returns SecurityAttributes
func (m MongoDbConnection) GetSecurityAttributes() map[string]map[string]interface{} {
	return m.SecurityAttributes
}

func (m MongoDbConnection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MongoDbConnection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMongoDbConnectionTechnologyTypeEnum(string(m.TechnologyType)); !ok && m.TechnologyType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TechnologyType: %s. Supported values are: %s.", m.TechnologyType, strings.Join(GetMongoDbConnectionTechnologyTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMongoDbConnectionSecurityProtocolEnum(string(m.SecurityProtocol)); !ok && m.SecurityProtocol != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SecurityProtocol: %s. Supported values are: %s.", m.SecurityProtocol, strings.Join(GetMongoDbConnectionSecurityProtocolEnumStringValues(), ",")))
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
func (m MongoDbConnection) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeMongoDbConnection MongoDbConnection
	s := struct {
		DiscriminatorParam string `json:"connectionType"`
		MarshalTypeMongoDbConnection
	}{
		"MONGODB",
		(MarshalTypeMongoDbConnection)(m),
	}

	return json.Marshal(&s)
}

// MongoDbConnectionTechnologyTypeEnum Enum with underlying type: string
type MongoDbConnectionTechnologyTypeEnum string

// Set of constants representing the allowable values for MongoDbConnectionTechnologyTypeEnum
const (
	MongoDbConnectionTechnologyTypeMongodb                   MongoDbConnectionTechnologyTypeEnum = "MONGODB"
	MongoDbConnectionTechnologyTypeOciAutonomousJsonDatabase MongoDbConnectionTechnologyTypeEnum = "OCI_AUTONOMOUS_JSON_DATABASE"
	MongoDbConnectionTechnologyTypeAzureCosmosDbForMongodb   MongoDbConnectionTechnologyTypeEnum = "AZURE_COSMOS_DB_FOR_MONGODB"
	MongoDbConnectionTechnologyTypeAmazonDocumentDb          MongoDbConnectionTechnologyTypeEnum = "AMAZON_DOCUMENT_DB"
	MongoDbConnectionTechnologyTypeOracleJsonCollection      MongoDbConnectionTechnologyTypeEnum = "ORACLE_JSON_COLLECTION"
	MongoDbConnectionTechnologyTypeOracleRestDataServices    MongoDbConnectionTechnologyTypeEnum = "ORACLE_REST_DATA_SERVICES"
)

var mappingMongoDbConnectionTechnologyTypeEnum = map[string]MongoDbConnectionTechnologyTypeEnum{
	"MONGODB":                      MongoDbConnectionTechnologyTypeMongodb,
	"OCI_AUTONOMOUS_JSON_DATABASE": MongoDbConnectionTechnologyTypeOciAutonomousJsonDatabase,
	"AZURE_COSMOS_DB_FOR_MONGODB":  MongoDbConnectionTechnologyTypeAzureCosmosDbForMongodb,
	"AMAZON_DOCUMENT_DB":           MongoDbConnectionTechnologyTypeAmazonDocumentDb,
	"ORACLE_JSON_COLLECTION":       MongoDbConnectionTechnologyTypeOracleJsonCollection,
	"ORACLE_REST_DATA_SERVICES":    MongoDbConnectionTechnologyTypeOracleRestDataServices,
}

var mappingMongoDbConnectionTechnologyTypeEnumLowerCase = map[string]MongoDbConnectionTechnologyTypeEnum{
	"mongodb":                      MongoDbConnectionTechnologyTypeMongodb,
	"oci_autonomous_json_database": MongoDbConnectionTechnologyTypeOciAutonomousJsonDatabase,
	"azure_cosmos_db_for_mongodb":  MongoDbConnectionTechnologyTypeAzureCosmosDbForMongodb,
	"amazon_document_db":           MongoDbConnectionTechnologyTypeAmazonDocumentDb,
	"oracle_json_collection":       MongoDbConnectionTechnologyTypeOracleJsonCollection,
	"oracle_rest_data_services":    MongoDbConnectionTechnologyTypeOracleRestDataServices,
}

// GetMongoDbConnectionTechnologyTypeEnumValues Enumerates the set of values for MongoDbConnectionTechnologyTypeEnum
func GetMongoDbConnectionTechnologyTypeEnumValues() []MongoDbConnectionTechnologyTypeEnum {
	values := make([]MongoDbConnectionTechnologyTypeEnum, 0)
	for _, v := range mappingMongoDbConnectionTechnologyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMongoDbConnectionTechnologyTypeEnumStringValues Enumerates the set of values in String for MongoDbConnectionTechnologyTypeEnum
func GetMongoDbConnectionTechnologyTypeEnumStringValues() []string {
	return []string{
		"MONGODB",
		"OCI_AUTONOMOUS_JSON_DATABASE",
		"AZURE_COSMOS_DB_FOR_MONGODB",
		"AMAZON_DOCUMENT_DB",
		"ORACLE_JSON_COLLECTION",
		"ORACLE_REST_DATA_SERVICES",
	}
}

// GetMappingMongoDbConnectionTechnologyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMongoDbConnectionTechnologyTypeEnum(val string) (MongoDbConnectionTechnologyTypeEnum, bool) {
	enum, ok := mappingMongoDbConnectionTechnologyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// MongoDbConnectionSecurityProtocolEnum Enum with underlying type: string
type MongoDbConnectionSecurityProtocolEnum string

// Set of constants representing the allowable values for MongoDbConnectionSecurityProtocolEnum
const (
	MongoDbConnectionSecurityProtocolPlain MongoDbConnectionSecurityProtocolEnum = "PLAIN"
	MongoDbConnectionSecurityProtocolTls   MongoDbConnectionSecurityProtocolEnum = "TLS"
	MongoDbConnectionSecurityProtocolMtls  MongoDbConnectionSecurityProtocolEnum = "MTLS"
)

var mappingMongoDbConnectionSecurityProtocolEnum = map[string]MongoDbConnectionSecurityProtocolEnum{
	"PLAIN": MongoDbConnectionSecurityProtocolPlain,
	"TLS":   MongoDbConnectionSecurityProtocolTls,
	"MTLS":  MongoDbConnectionSecurityProtocolMtls,
}

var mappingMongoDbConnectionSecurityProtocolEnumLowerCase = map[string]MongoDbConnectionSecurityProtocolEnum{
	"plain": MongoDbConnectionSecurityProtocolPlain,
	"tls":   MongoDbConnectionSecurityProtocolTls,
	"mtls":  MongoDbConnectionSecurityProtocolMtls,
}

// GetMongoDbConnectionSecurityProtocolEnumValues Enumerates the set of values for MongoDbConnectionSecurityProtocolEnum
func GetMongoDbConnectionSecurityProtocolEnumValues() []MongoDbConnectionSecurityProtocolEnum {
	values := make([]MongoDbConnectionSecurityProtocolEnum, 0)
	for _, v := range mappingMongoDbConnectionSecurityProtocolEnum {
		values = append(values, v)
	}
	return values
}

// GetMongoDbConnectionSecurityProtocolEnumStringValues Enumerates the set of values in String for MongoDbConnectionSecurityProtocolEnum
func GetMongoDbConnectionSecurityProtocolEnumStringValues() []string {
	return []string{
		"PLAIN",
		"TLS",
		"MTLS",
	}
}

// GetMappingMongoDbConnectionSecurityProtocolEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMongoDbConnectionSecurityProtocolEnum(val string) (MongoDbConnectionSecurityProtocolEnum, bool) {
	enum, ok := mappingMongoDbConnectionSecurityProtocolEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
