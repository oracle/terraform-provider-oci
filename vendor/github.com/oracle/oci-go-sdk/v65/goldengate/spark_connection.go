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

// SparkConnection Represents the metadata of a Spark Connection.
type SparkConnection struct {

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

	// Spark master URL used for job submission.
	// Required for both submitMode SUBMIT and REST.
	// e.g.: 'spark://spark-master.example.com:7077'
	// High availability can be configured with a comma-separated list, for example:
	// 'spark://h1.example.com:7077,h2.example.com:7077'
	ClusterEndpoint *string `mandatory:"true" json:"clusterEndpoint"`

	ArtifactLocation SparkArtifactLocation `mandatory:"true" json:"artifactLocation"`

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

	// References the OCI Vault that contains the customer-managed encryption key identified by `keyId`.
	// Deprecated: This field is deprecated for GoldenGate connections. Sensitive attributes should be provided using the
	// corresponding Secret OCID attributes of the connection (for example, `passwordSecretId`) instead of plain-text
	// attributes encrypted with `vaultId` and `keyId`. This change follows the GoldenGate "Plain Text Fields in Connections" deprecation:
	// https://docs.oracle.com/en-us/iaas/Content/servicechanges.htm#servicechanges_topic-GoldenGate
	// This field is applicable only when `doesUseSecretIds` is set to `false`.
	// If `vaultId` is provided, `keyId` must also be provided.
	VaultId *string `mandatory:"false" json:"vaultId"`

	// References the OCI Vault key in the OCI Vault identified by `vaultId`.
	// Deprecated: This field is deprecated for GoldenGate connections. Sensitive attributes should be provided using the
	// corresponding Secret OCID attributes of the connection (for example, `passwordSecretId`) instead of plain-text
	// attributes encrypted with `vaultId` and `keyId`. This change follows the GoldenGate "Plain Text Fields in Connections" deprecation:
	// https://docs.oracle.com/en-us/iaas/Content/servicechanges.htm#servicechanges_topic-GoldenGate
	// The GoldenGate service uses this key to encrypt sensitive information (for example, `password`) that is provided in plain-text connection attributes through the API.
	// This field is applicable only when `doesUseSecretIds` is set to `false`. If both `vaultId` and `keyId` are provided,
	// the GoldenGate service uses the specified customer-managed key to encrypt the sensitive data.
	// If neither `vaultId` nor `keyId` is provided, the GoldenGate service uses Oracle-managed encryption keys.
	KeyId *string `mandatory:"false" json:"keyId"`

	// List of ingress IP addresses from where the GoldenGate deployment connects to this connection's privateIp.
	// Customers may optionally set up ingress security rules to restrict traffic from these IP addresses.
	IngressIps []IngressIpDetails `mandatory:"false" json:"ingressIps"`

	// An array of Network Security Group OCIDs used to define network access for either Deployments or Connections.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the target subnet of the dedicated connection.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// Indicates that sensitive attributes are provided via Secrets.
	// Deprecated: This field is deprecated. Sensitive attributes should be provided using the corresponding Secret OCID
	// attributes of the connection (for example, `passwordSecretId`) instead of plain-text attributes. This change follows
	// the GoldenGate "Plain Text Fields in Connections" deprecation:
	// https://docs.oracle.com/en-us/iaas/Content/servicechanges.htm#servicechanges_topic-GoldenGate
	// When set to `true`, all sensitive information must be provided as OCI Vault secrets using the corresponding
	// `*SecretId` attributes of the connection (for example, `passwordSecretId`). Plain-text sensitive attributes (for example, `password`) must not be used.
	// This ensures that sensitive information remains stored and managed in the customer's OCI Vault rather than by the GoldenGate service.
	// When set to false, sensitive information must be provided in the corresponding plain-text attributes (for example, `password`) rather than in secret OCID attributes.
	// In this mode, the sensitive information is stored by the GoldenGate service. If `vaultId` and `keyId` are not specified,
	// the GoldenGate service uses Oracle-managed encryption keys to encrypt the stored data.
	// If `vaultId` and `keyId` are provided, the specified customer-managed key is used.
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

	// Spark standalone REST submission base URL.
	// e.g.: 'https://spark-master.example.com:6066'
	// Applicable only and required when submitMode is REST.
	RestEndpoint *string `mandatory:"false" json:"restEndpoint"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the Spark authentication shared secret.
	// Applicable when authenticationType is SHARED_SECRET.
	// Note: When provided, 'authSecret' field must not be provided.
	AuthSecretSecretId *string `mandatory:"false" json:"authSecretSecretId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the REST submission signing key.
	// Applicable only when submitMode is REST and the cluster enforces REST request signing.
	// This setting is independent of authenticationType.
	// Note: When provided, 'restAuthSecretKey' field must not be provided.
	RestAuthSecretKeySecretId *string `mandatory:"false" json:"restAuthSecretKeySecretId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the trust store content.
	// Applicable when securityProtocol is TLS.
	// Provide this when the server certificate is not already trusted by the JVM.
	// Note: When provided, 'trustStore' field must not be provided.
	TrustStoreSecretId *string `mandatory:"false" json:"trustStoreSecretId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the trust store password.
	// Applicable when securityProtocol is TLS and a trust store is provided.
	// Required when 'trustStoreSecretId' is provided.
	// Note: When provided, 'trustStorePassword' field must not be provided.
	TrustStorePasswordSecretId *string `mandatory:"false" json:"trustStorePasswordSecretId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the client key store content.
	// Applicable when securityProtocol is TLS.
	// Note: When provided, 'keyStore' field must not be provided.
	KeyStoreSecretId *string `mandatory:"false" json:"keyStoreSecretId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the client key store password.
	// Applicable when securityProtocol is TLS and a client key store is provided.
	// Required when 'keyStoreSecretId' is provided.
	// Note: When provided, 'keyStorePassword' field must not be provided.
	KeyStorePasswordSecretId *string `mandatory:"false" json:"keyStorePasswordSecretId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the private key password.
	// Applicable when securityProtocol is TLS and a client key store is provided.
	// Note: When provided, 'sslKeyPassword' field must not be provided.
	SslKeyPasswordSecretId *string `mandatory:"false" json:"sslKeyPasswordSecretId"`

	// The Spark technology type.
	TechnologyType SparkConnectionTechnologyTypeEnum `mandatory:"true" json:"technologyType"`

	// Submission channel used to connect to Spark.
	SubmitMode SparkConnectionSubmitModeEnum `mandatory:"true" json:"submitMode"`

	// Authentication type for Spark.
	AuthenticationType SparkConnectionAuthenticationTypeEnum `mandatory:"true" json:"authenticationType"`

	// Security protocol for Spark submission traffic.
	SecurityProtocol SparkConnectionSecurityProtocolEnum `mandatory:"true" json:"securityProtocol"`

	// Cipher transform used for Spark RPC encryption.
	// Applicable when securityProtocol is AES_RPC.
	RpcCipher SparkConnectionRpcCipherEnum `mandatory:"false" json:"rpcCipher,omitempty"`

	// Trust store format used by the Spark client.
	TrustStoreType SparkConnectionTrustStoreTypeEnum `mandatory:"false" json:"trustStoreType,omitempty"`

	// Client key store format used for mutual TLS.
	KeyStoreType SparkConnectionKeyStoreTypeEnum `mandatory:"false" json:"keyStoreType,omitempty"`

	// TLS protocol negotiated by the Spark client.
	TlsProtocol SparkConnectionTlsProtocolEnum `mandatory:"false" json:"tlsProtocol,omitempty"`

	// Possible lifecycle states for connection.
	LifecycleState ConnectionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Controls the network traffic direction to the target:
	// SHARED_DEPLOYMENT_ENDPOINT: Network traffic flows from the assigned deployment's private endpoint through the deployment's subnet.
	// DEDICATED_ENDPOINT: A dedicated private endpoint is created in the target VCN subnet for the connection. The subnetId is required when DEDICATED_ENDPOINT networking is selected.
	// SHARED_SERVICE_ENDPOINT: Traffic flows through the Goldengate Service's network to public hosts. Cannot be used for private targets.
	// Deprecated: SHARED_SERVICE_ENDPOINT is deprecated. Use another supported routingMethod value, or update existing connections to use a supported routing method.
	// This change follows the GoldenGate "Plain Text Fields in Connections" deprecation:
	// https://docs.oracle.com/en-us/iaas/Content/servicechanges.htm#servicechanges_topic-GoldenGate
	RoutingMethod RoutingMethodEnum `mandatory:"false" json:"routingMethod,omitempty"`
}

// GetId returns Id
func (m SparkConnection) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m SparkConnection) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m SparkConnection) GetDescription() *string {
	return m.Description
}

// GetCompartmentId returns CompartmentId
func (m SparkConnection) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetFreeformTags returns FreeformTags
func (m SparkConnection) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m SparkConnection) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m SparkConnection) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetLifecycleState returns LifecycleState
func (m SparkConnection) GetLifecycleState() ConnectionLifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m SparkConnection) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetTimeCreated returns TimeCreated
func (m SparkConnection) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m SparkConnection) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLocks returns Locks
func (m SparkConnection) GetLocks() []ResourceLock {
	return m.Locks
}

// GetVaultId returns VaultId
func (m SparkConnection) GetVaultId() *string {
	return m.VaultId
}

// GetKeyId returns KeyId
func (m SparkConnection) GetKeyId() *string {
	return m.KeyId
}

// GetIngressIps returns IngressIps
func (m SparkConnection) GetIngressIps() []IngressIpDetails {
	return m.IngressIps
}

// GetNsgIds returns NsgIds
func (m SparkConnection) GetNsgIds() []string {
	return m.NsgIds
}

// GetSubnetId returns SubnetId
func (m SparkConnection) GetSubnetId() *string {
	return m.SubnetId
}

// GetRoutingMethod returns RoutingMethod
func (m SparkConnection) GetRoutingMethod() RoutingMethodEnum {
	return m.RoutingMethod
}

// GetDoesUseSecretIds returns DoesUseSecretIds
func (m SparkConnection) GetDoesUseSecretIds() *bool {
	return m.DoesUseSecretIds
}

// GetSubscriptionId returns SubscriptionId
func (m SparkConnection) GetSubscriptionId() *string {
	return m.SubscriptionId
}

// GetClusterPlacementGroupId returns ClusterPlacementGroupId
func (m SparkConnection) GetClusterPlacementGroupId() *string {
	return m.ClusterPlacementGroupId
}

// GetSecurityAttributes returns SecurityAttributes
func (m SparkConnection) GetSecurityAttributes() map[string]map[string]interface{} {
	return m.SecurityAttributes
}

func (m SparkConnection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SparkConnection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSparkConnectionTechnologyTypeEnum(string(m.TechnologyType)); !ok && m.TechnologyType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TechnologyType: %s. Supported values are: %s.", m.TechnologyType, strings.Join(GetSparkConnectionTechnologyTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSparkConnectionSubmitModeEnum(string(m.SubmitMode)); !ok && m.SubmitMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SubmitMode: %s. Supported values are: %s.", m.SubmitMode, strings.Join(GetSparkConnectionSubmitModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSparkConnectionAuthenticationTypeEnum(string(m.AuthenticationType)); !ok && m.AuthenticationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AuthenticationType: %s. Supported values are: %s.", m.AuthenticationType, strings.Join(GetSparkConnectionAuthenticationTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSparkConnectionSecurityProtocolEnum(string(m.SecurityProtocol)); !ok && m.SecurityProtocol != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SecurityProtocol: %s. Supported values are: %s.", m.SecurityProtocol, strings.Join(GetSparkConnectionSecurityProtocolEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSparkConnectionRpcCipherEnum(string(m.RpcCipher)); !ok && m.RpcCipher != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RpcCipher: %s. Supported values are: %s.", m.RpcCipher, strings.Join(GetSparkConnectionRpcCipherEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSparkConnectionTrustStoreTypeEnum(string(m.TrustStoreType)); !ok && m.TrustStoreType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TrustStoreType: %s. Supported values are: %s.", m.TrustStoreType, strings.Join(GetSparkConnectionTrustStoreTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSparkConnectionKeyStoreTypeEnum(string(m.KeyStoreType)); !ok && m.KeyStoreType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for KeyStoreType: %s. Supported values are: %s.", m.KeyStoreType, strings.Join(GetSparkConnectionKeyStoreTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSparkConnectionTlsProtocolEnum(string(m.TlsProtocol)); !ok && m.TlsProtocol != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TlsProtocol: %s. Supported values are: %s.", m.TlsProtocol, strings.Join(GetSparkConnectionTlsProtocolEnumStringValues(), ",")))
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
func (m SparkConnection) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeSparkConnection SparkConnection
	s := struct {
		DiscriminatorParam string `json:"connectionType"`
		MarshalTypeSparkConnection
	}{
		"SPARK",
		(MarshalTypeSparkConnection)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *SparkConnection) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description                *string                               `json:"description"`
		FreeformTags               map[string]string                     `json:"freeformTags"`
		DefinedTags                map[string]map[string]interface{}     `json:"definedTags"`
		SystemTags                 map[string]map[string]interface{}     `json:"systemTags"`
		LifecycleDetails           *string                               `json:"lifecycleDetails"`
		Locks                      []ResourceLock                        `json:"locks"`
		VaultId                    *string                               `json:"vaultId"`
		KeyId                      *string                               `json:"keyId"`
		IngressIps                 []IngressIpDetails                    `json:"ingressIps"`
		NsgIds                     []string                              `json:"nsgIds"`
		SubnetId                   *string                               `json:"subnetId"`
		RoutingMethod              RoutingMethodEnum                     `json:"routingMethod"`
		DoesUseSecretIds           *bool                                 `json:"doesUseSecretIds"`
		SubscriptionId             *string                               `json:"subscriptionId"`
		ClusterPlacementGroupId    *string                               `json:"clusterPlacementGroupId"`
		SecurityAttributes         map[string]map[string]interface{}     `json:"securityAttributes"`
		RestEndpoint               *string                               `json:"restEndpoint"`
		AuthSecretSecretId         *string                               `json:"authSecretSecretId"`
		RpcCipher                  SparkConnectionRpcCipherEnum          `json:"rpcCipher"`
		RestAuthSecretKeySecretId  *string                               `json:"restAuthSecretKeySecretId"`
		TrustStoreSecretId         *string                               `json:"trustStoreSecretId"`
		TrustStorePasswordSecretId *string                               `json:"trustStorePasswordSecretId"`
		TrustStoreType             SparkConnectionTrustStoreTypeEnum     `json:"trustStoreType"`
		KeyStoreSecretId           *string                               `json:"keyStoreSecretId"`
		KeyStorePasswordSecretId   *string                               `json:"keyStorePasswordSecretId"`
		KeyStoreType               SparkConnectionKeyStoreTypeEnum       `json:"keyStoreType"`
		SslKeyPasswordSecretId     *string                               `json:"sslKeyPasswordSecretId"`
		TlsProtocol                SparkConnectionTlsProtocolEnum        `json:"tlsProtocol"`
		Id                         *string                               `json:"id"`
		DisplayName                *string                               `json:"displayName"`
		CompartmentId              *string                               `json:"compartmentId"`
		LifecycleState             ConnectionLifecycleStateEnum          `json:"lifecycleState"`
		TimeCreated                *common.SDKTime                       `json:"timeCreated"`
		TimeUpdated                *common.SDKTime                       `json:"timeUpdated"`
		TechnologyType             SparkConnectionTechnologyTypeEnum     `json:"technologyType"`
		SubmitMode                 SparkConnectionSubmitModeEnum         `json:"submitMode"`
		AuthenticationType         SparkConnectionAuthenticationTypeEnum `json:"authenticationType"`
		SecurityProtocol           SparkConnectionSecurityProtocolEnum   `json:"securityProtocol"`
		ClusterEndpoint            *string                               `json:"clusterEndpoint"`
		ArtifactLocation           sparkartifactlocation                 `json:"artifactLocation"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.LifecycleDetails = model.LifecycleDetails

	m.Locks = make([]ResourceLock, len(model.Locks))
	copy(m.Locks, model.Locks)
	m.VaultId = model.VaultId

	m.KeyId = model.KeyId

	m.IngressIps = make([]IngressIpDetails, len(model.IngressIps))
	copy(m.IngressIps, model.IngressIps)
	m.NsgIds = make([]string, len(model.NsgIds))
	copy(m.NsgIds, model.NsgIds)
	m.SubnetId = model.SubnetId

	m.RoutingMethod = model.RoutingMethod

	m.DoesUseSecretIds = model.DoesUseSecretIds

	m.SubscriptionId = model.SubscriptionId

	m.ClusterPlacementGroupId = model.ClusterPlacementGroupId

	m.SecurityAttributes = model.SecurityAttributes

	m.RestEndpoint = model.RestEndpoint

	m.AuthSecretSecretId = model.AuthSecretSecretId

	m.RpcCipher = model.RpcCipher

	m.RestAuthSecretKeySecretId = model.RestAuthSecretKeySecretId

	m.TrustStoreSecretId = model.TrustStoreSecretId

	m.TrustStorePasswordSecretId = model.TrustStorePasswordSecretId

	m.TrustStoreType = model.TrustStoreType

	m.KeyStoreSecretId = model.KeyStoreSecretId

	m.KeyStorePasswordSecretId = model.KeyStorePasswordSecretId

	m.KeyStoreType = model.KeyStoreType

	m.SslKeyPasswordSecretId = model.SslKeyPasswordSecretId

	m.TlsProtocol = model.TlsProtocol

	m.Id = model.Id

	m.DisplayName = model.DisplayName

	m.CompartmentId = model.CompartmentId

	m.LifecycleState = model.LifecycleState

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.TechnologyType = model.TechnologyType

	m.SubmitMode = model.SubmitMode

	m.AuthenticationType = model.AuthenticationType

	m.SecurityProtocol = model.SecurityProtocol

	m.ClusterEndpoint = model.ClusterEndpoint

	nn, e = model.ArtifactLocation.UnmarshalPolymorphicJSON(model.ArtifactLocation.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ArtifactLocation = nn.(SparkArtifactLocation)
	} else {
		m.ArtifactLocation = nil
	}

	return
}

// SparkConnectionTechnologyTypeEnum Enum with underlying type: string
type SparkConnectionTechnologyTypeEnum string

// Set of constants representing the allowable values for SparkConnectionTechnologyTypeEnum
const (
	SparkConnectionTechnologyTypeApacheSpark SparkConnectionTechnologyTypeEnum = "APACHE_SPARK"
)

var mappingSparkConnectionTechnologyTypeEnum = map[string]SparkConnectionTechnologyTypeEnum{
	"APACHE_SPARK": SparkConnectionTechnologyTypeApacheSpark,
}

var mappingSparkConnectionTechnologyTypeEnumLowerCase = map[string]SparkConnectionTechnologyTypeEnum{
	"apache_spark": SparkConnectionTechnologyTypeApacheSpark,
}

// GetSparkConnectionTechnologyTypeEnumValues Enumerates the set of values for SparkConnectionTechnologyTypeEnum
func GetSparkConnectionTechnologyTypeEnumValues() []SparkConnectionTechnologyTypeEnum {
	values := make([]SparkConnectionTechnologyTypeEnum, 0)
	for _, v := range mappingSparkConnectionTechnologyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSparkConnectionTechnologyTypeEnumStringValues Enumerates the set of values in String for SparkConnectionTechnologyTypeEnum
func GetSparkConnectionTechnologyTypeEnumStringValues() []string {
	return []string{
		"APACHE_SPARK",
	}
}

// GetMappingSparkConnectionTechnologyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSparkConnectionTechnologyTypeEnum(val string) (SparkConnectionTechnologyTypeEnum, bool) {
	enum, ok := mappingSparkConnectionTechnologyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SparkConnectionSubmitModeEnum Enum with underlying type: string
type SparkConnectionSubmitModeEnum string

// Set of constants representing the allowable values for SparkConnectionSubmitModeEnum
const (
	SparkConnectionSubmitModeSubmit SparkConnectionSubmitModeEnum = "SUBMIT"
	SparkConnectionSubmitModeRest   SparkConnectionSubmitModeEnum = "REST"
)

var mappingSparkConnectionSubmitModeEnum = map[string]SparkConnectionSubmitModeEnum{
	"SUBMIT": SparkConnectionSubmitModeSubmit,
	"REST":   SparkConnectionSubmitModeRest,
}

var mappingSparkConnectionSubmitModeEnumLowerCase = map[string]SparkConnectionSubmitModeEnum{
	"submit": SparkConnectionSubmitModeSubmit,
	"rest":   SparkConnectionSubmitModeRest,
}

// GetSparkConnectionSubmitModeEnumValues Enumerates the set of values for SparkConnectionSubmitModeEnum
func GetSparkConnectionSubmitModeEnumValues() []SparkConnectionSubmitModeEnum {
	values := make([]SparkConnectionSubmitModeEnum, 0)
	for _, v := range mappingSparkConnectionSubmitModeEnum {
		values = append(values, v)
	}
	return values
}

// GetSparkConnectionSubmitModeEnumStringValues Enumerates the set of values in String for SparkConnectionSubmitModeEnum
func GetSparkConnectionSubmitModeEnumStringValues() []string {
	return []string{
		"SUBMIT",
		"REST",
	}
}

// GetMappingSparkConnectionSubmitModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSparkConnectionSubmitModeEnum(val string) (SparkConnectionSubmitModeEnum, bool) {
	enum, ok := mappingSparkConnectionSubmitModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SparkConnectionAuthenticationTypeEnum Enum with underlying type: string
type SparkConnectionAuthenticationTypeEnum string

// Set of constants representing the allowable values for SparkConnectionAuthenticationTypeEnum
const (
	SparkConnectionAuthenticationTypeNone         SparkConnectionAuthenticationTypeEnum = "NONE"
	SparkConnectionAuthenticationTypeSharedSecret SparkConnectionAuthenticationTypeEnum = "SHARED_SECRET"
)

var mappingSparkConnectionAuthenticationTypeEnum = map[string]SparkConnectionAuthenticationTypeEnum{
	"NONE":          SparkConnectionAuthenticationTypeNone,
	"SHARED_SECRET": SparkConnectionAuthenticationTypeSharedSecret,
}

var mappingSparkConnectionAuthenticationTypeEnumLowerCase = map[string]SparkConnectionAuthenticationTypeEnum{
	"none":          SparkConnectionAuthenticationTypeNone,
	"shared_secret": SparkConnectionAuthenticationTypeSharedSecret,
}

// GetSparkConnectionAuthenticationTypeEnumValues Enumerates the set of values for SparkConnectionAuthenticationTypeEnum
func GetSparkConnectionAuthenticationTypeEnumValues() []SparkConnectionAuthenticationTypeEnum {
	values := make([]SparkConnectionAuthenticationTypeEnum, 0)
	for _, v := range mappingSparkConnectionAuthenticationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSparkConnectionAuthenticationTypeEnumStringValues Enumerates the set of values in String for SparkConnectionAuthenticationTypeEnum
func GetSparkConnectionAuthenticationTypeEnumStringValues() []string {
	return []string{
		"NONE",
		"SHARED_SECRET",
	}
}

// GetMappingSparkConnectionAuthenticationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSparkConnectionAuthenticationTypeEnum(val string) (SparkConnectionAuthenticationTypeEnum, bool) {
	enum, ok := mappingSparkConnectionAuthenticationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SparkConnectionSecurityProtocolEnum Enum with underlying type: string
type SparkConnectionSecurityProtocolEnum string

// Set of constants representing the allowable values for SparkConnectionSecurityProtocolEnum
const (
	SparkConnectionSecurityProtocolPlaintext     SparkConnectionSecurityProtocolEnum = "PLAINTEXT"
	SparkConnectionSecurityProtocolTls           SparkConnectionSecurityProtocolEnum = "TLS"
	SparkConnectionSecurityProtocolAesRpc        SparkConnectionSecurityProtocolEnum = "AES_RPC"
	SparkConnectionSecurityProtocolSaslEncrypted SparkConnectionSecurityProtocolEnum = "SASL_ENCRYPTED"
)

var mappingSparkConnectionSecurityProtocolEnum = map[string]SparkConnectionSecurityProtocolEnum{
	"PLAINTEXT":      SparkConnectionSecurityProtocolPlaintext,
	"TLS":            SparkConnectionSecurityProtocolTls,
	"AES_RPC":        SparkConnectionSecurityProtocolAesRpc,
	"SASL_ENCRYPTED": SparkConnectionSecurityProtocolSaslEncrypted,
}

var mappingSparkConnectionSecurityProtocolEnumLowerCase = map[string]SparkConnectionSecurityProtocolEnum{
	"plaintext":      SparkConnectionSecurityProtocolPlaintext,
	"tls":            SparkConnectionSecurityProtocolTls,
	"aes_rpc":        SparkConnectionSecurityProtocolAesRpc,
	"sasl_encrypted": SparkConnectionSecurityProtocolSaslEncrypted,
}

// GetSparkConnectionSecurityProtocolEnumValues Enumerates the set of values for SparkConnectionSecurityProtocolEnum
func GetSparkConnectionSecurityProtocolEnumValues() []SparkConnectionSecurityProtocolEnum {
	values := make([]SparkConnectionSecurityProtocolEnum, 0)
	for _, v := range mappingSparkConnectionSecurityProtocolEnum {
		values = append(values, v)
	}
	return values
}

// GetSparkConnectionSecurityProtocolEnumStringValues Enumerates the set of values in String for SparkConnectionSecurityProtocolEnum
func GetSparkConnectionSecurityProtocolEnumStringValues() []string {
	return []string{
		"PLAINTEXT",
		"TLS",
		"AES_RPC",
		"SASL_ENCRYPTED",
	}
}

// GetMappingSparkConnectionSecurityProtocolEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSparkConnectionSecurityProtocolEnum(val string) (SparkConnectionSecurityProtocolEnum, bool) {
	enum, ok := mappingSparkConnectionSecurityProtocolEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SparkConnectionRpcCipherEnum Enum with underlying type: string
type SparkConnectionRpcCipherEnum string

// Set of constants representing the allowable values for SparkConnectionRpcCipherEnum
const (
	SparkConnectionRpcCipherCtrNopadding SparkConnectionRpcCipherEnum = "AES_CTR_NOPADDING"
	SparkConnectionRpcCipherGcmNopadding SparkConnectionRpcCipherEnum = "AES_GCM_NOPADDING"
)

var mappingSparkConnectionRpcCipherEnum = map[string]SparkConnectionRpcCipherEnum{
	"AES_CTR_NOPADDING": SparkConnectionRpcCipherCtrNopadding,
	"AES_GCM_NOPADDING": SparkConnectionRpcCipherGcmNopadding,
}

var mappingSparkConnectionRpcCipherEnumLowerCase = map[string]SparkConnectionRpcCipherEnum{
	"aes_ctr_nopadding": SparkConnectionRpcCipherCtrNopadding,
	"aes_gcm_nopadding": SparkConnectionRpcCipherGcmNopadding,
}

// GetSparkConnectionRpcCipherEnumValues Enumerates the set of values for SparkConnectionRpcCipherEnum
func GetSparkConnectionRpcCipherEnumValues() []SparkConnectionRpcCipherEnum {
	values := make([]SparkConnectionRpcCipherEnum, 0)
	for _, v := range mappingSparkConnectionRpcCipherEnum {
		values = append(values, v)
	}
	return values
}

// GetSparkConnectionRpcCipherEnumStringValues Enumerates the set of values in String for SparkConnectionRpcCipherEnum
func GetSparkConnectionRpcCipherEnumStringValues() []string {
	return []string{
		"AES_CTR_NOPADDING",
		"AES_GCM_NOPADDING",
	}
}

// GetMappingSparkConnectionRpcCipherEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSparkConnectionRpcCipherEnum(val string) (SparkConnectionRpcCipherEnum, bool) {
	enum, ok := mappingSparkConnectionRpcCipherEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SparkConnectionTrustStoreTypeEnum Enum with underlying type: string
type SparkConnectionTrustStoreTypeEnum string

// Set of constants representing the allowable values for SparkConnectionTrustStoreTypeEnum
const (
	SparkConnectionTrustStoreTypeJks    SparkConnectionTrustStoreTypeEnum = "JKS"
	SparkConnectionTrustStoreTypePkcs12 SparkConnectionTrustStoreTypeEnum = "PKCS12"
	SparkConnectionTrustStoreTypeJceks  SparkConnectionTrustStoreTypeEnum = "JCEKS"
)

var mappingSparkConnectionTrustStoreTypeEnum = map[string]SparkConnectionTrustStoreTypeEnum{
	"JKS":    SparkConnectionTrustStoreTypeJks,
	"PKCS12": SparkConnectionTrustStoreTypePkcs12,
	"JCEKS":  SparkConnectionTrustStoreTypeJceks,
}

var mappingSparkConnectionTrustStoreTypeEnumLowerCase = map[string]SparkConnectionTrustStoreTypeEnum{
	"jks":    SparkConnectionTrustStoreTypeJks,
	"pkcs12": SparkConnectionTrustStoreTypePkcs12,
	"jceks":  SparkConnectionTrustStoreTypeJceks,
}

// GetSparkConnectionTrustStoreTypeEnumValues Enumerates the set of values for SparkConnectionTrustStoreTypeEnum
func GetSparkConnectionTrustStoreTypeEnumValues() []SparkConnectionTrustStoreTypeEnum {
	values := make([]SparkConnectionTrustStoreTypeEnum, 0)
	for _, v := range mappingSparkConnectionTrustStoreTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSparkConnectionTrustStoreTypeEnumStringValues Enumerates the set of values in String for SparkConnectionTrustStoreTypeEnum
func GetSparkConnectionTrustStoreTypeEnumStringValues() []string {
	return []string{
		"JKS",
		"PKCS12",
		"JCEKS",
	}
}

// GetMappingSparkConnectionTrustStoreTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSparkConnectionTrustStoreTypeEnum(val string) (SparkConnectionTrustStoreTypeEnum, bool) {
	enum, ok := mappingSparkConnectionTrustStoreTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SparkConnectionKeyStoreTypeEnum Enum with underlying type: string
type SparkConnectionKeyStoreTypeEnum string

// Set of constants representing the allowable values for SparkConnectionKeyStoreTypeEnum
const (
	SparkConnectionKeyStoreTypeJks    SparkConnectionKeyStoreTypeEnum = "JKS"
	SparkConnectionKeyStoreTypePkcs12 SparkConnectionKeyStoreTypeEnum = "PKCS12"
	SparkConnectionKeyStoreTypeJceks  SparkConnectionKeyStoreTypeEnum = "JCEKS"
)

var mappingSparkConnectionKeyStoreTypeEnum = map[string]SparkConnectionKeyStoreTypeEnum{
	"JKS":    SparkConnectionKeyStoreTypeJks,
	"PKCS12": SparkConnectionKeyStoreTypePkcs12,
	"JCEKS":  SparkConnectionKeyStoreTypeJceks,
}

var mappingSparkConnectionKeyStoreTypeEnumLowerCase = map[string]SparkConnectionKeyStoreTypeEnum{
	"jks":    SparkConnectionKeyStoreTypeJks,
	"pkcs12": SparkConnectionKeyStoreTypePkcs12,
	"jceks":  SparkConnectionKeyStoreTypeJceks,
}

// GetSparkConnectionKeyStoreTypeEnumValues Enumerates the set of values for SparkConnectionKeyStoreTypeEnum
func GetSparkConnectionKeyStoreTypeEnumValues() []SparkConnectionKeyStoreTypeEnum {
	values := make([]SparkConnectionKeyStoreTypeEnum, 0)
	for _, v := range mappingSparkConnectionKeyStoreTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSparkConnectionKeyStoreTypeEnumStringValues Enumerates the set of values in String for SparkConnectionKeyStoreTypeEnum
func GetSparkConnectionKeyStoreTypeEnumStringValues() []string {
	return []string{
		"JKS",
		"PKCS12",
		"JCEKS",
	}
}

// GetMappingSparkConnectionKeyStoreTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSparkConnectionKeyStoreTypeEnum(val string) (SparkConnectionKeyStoreTypeEnum, bool) {
	enum, ok := mappingSparkConnectionKeyStoreTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SparkConnectionTlsProtocolEnum Enum with underlying type: string
type SparkConnectionTlsProtocolEnum string

// Set of constants representing the allowable values for SparkConnectionTlsProtocolEnum
const (
	SparkConnectionTlsProtocol2 SparkConnectionTlsProtocolEnum = "TLS_V1_2"
	SparkConnectionTlsProtocol3 SparkConnectionTlsProtocolEnum = "TLS_V1_3"
)

var mappingSparkConnectionTlsProtocolEnum = map[string]SparkConnectionTlsProtocolEnum{
	"TLS_V1_2": SparkConnectionTlsProtocol2,
	"TLS_V1_3": SparkConnectionTlsProtocol3,
}

var mappingSparkConnectionTlsProtocolEnumLowerCase = map[string]SparkConnectionTlsProtocolEnum{
	"tls_v1_2": SparkConnectionTlsProtocol2,
	"tls_v1_3": SparkConnectionTlsProtocol3,
}

// GetSparkConnectionTlsProtocolEnumValues Enumerates the set of values for SparkConnectionTlsProtocolEnum
func GetSparkConnectionTlsProtocolEnumValues() []SparkConnectionTlsProtocolEnum {
	values := make([]SparkConnectionTlsProtocolEnum, 0)
	for _, v := range mappingSparkConnectionTlsProtocolEnum {
		values = append(values, v)
	}
	return values
}

// GetSparkConnectionTlsProtocolEnumStringValues Enumerates the set of values in String for SparkConnectionTlsProtocolEnum
func GetSparkConnectionTlsProtocolEnumStringValues() []string {
	return []string{
		"TLS_V1_2",
		"TLS_V1_3",
	}
}

// GetMappingSparkConnectionTlsProtocolEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSparkConnectionTlsProtocolEnum(val string) (SparkConnectionTlsProtocolEnum, bool) {
	enum, ok := mappingSparkConnectionTlsProtocolEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
