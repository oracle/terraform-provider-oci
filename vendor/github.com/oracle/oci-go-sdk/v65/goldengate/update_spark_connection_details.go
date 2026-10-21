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

// UpdateSparkConnectionDetails The information to update a Spark Connection.
type UpdateSparkConnectionDetails struct {

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

	// Security attributes for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Oracle-ZPR": {"MaxEgressCount": {"value": "42", "mode": "enforce"}}}`
	SecurityAttributes map[string]map[string]interface{} `mandatory:"false" json:"securityAttributes"`

	// Spark master URL used for job submission.
	// Required for both submitMode SUBMIT and REST.
	// e.g.: 'spark://spark-master.example.com:7077'
	// High availability can be configured with a comma-separated list, for example:
	// 'spark://h1.example.com:7077,h2.example.com:7077'
	ClusterEndpoint *string `mandatory:"false" json:"clusterEndpoint"`

	// Spark standalone REST submission base URL.
	// e.g.: 'https://spark-master.example.com:6066'
	// Applicable only and required when submitMode is REST.
	RestEndpoint *string `mandatory:"false" json:"restEndpoint"`

	// The shared secret used for Spark authentication.
	// Applicable when authenticationType is SHARED_SECRET.
	// Deprecated: This field is deprecated and replaced by "authSecretSecretId".
	// This change follows the GoldenGate "Plain Text Fields in Connections" deprecation:
	// https://docs.oracle.com/en-us/iaas/Content/servicechanges.htm#servicechanges_topic-GoldenGate
	AuthSecret *string `mandatory:"false" json:"authSecret" sensitive:"true"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the Spark authentication shared secret.
	// Applicable when authenticationType is SHARED_SECRET.
	// Note: When provided, 'authSecret' field must not be provided.
	AuthSecretSecretId *string `mandatory:"false" json:"authSecretSecretId"`

	// The BASE64URL-encoded HMAC secret key used to sign REST submission requests when the Spark master's REST endpoint is protected by the JWS filter.
	// Applicable only when submitMode is REST and the cluster enforces REST request signing.
	// This setting is independent of authenticationType.
	// Deprecated: This field is deprecated and replaced by "restAuthSecretKeySecretId".
	// This change follows the GoldenGate "Plain Text Fields in Connections" deprecation:
	// https://docs.oracle.com/en-us/iaas/Content/servicechanges.htm#servicechanges_topic-GoldenGate
	RestAuthSecretKey *string `mandatory:"false" json:"restAuthSecretKey" sensitive:"true"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the REST submission signing key.
	// Applicable only when submitMode is REST and the cluster enforces REST request signing.
	// This setting is independent of authenticationType.
	// Note: When provided, 'restAuthSecretKey' field must not be provided.
	RestAuthSecretKeySecretId *string `mandatory:"false" json:"restAuthSecretKeySecretId"`

	// The base64 encoded content of the trust store used to validate the Spark server certificate.
	// Applicable when securityProtocol is TLS.
	// Provide this when the server certificate is not already trusted by the JVM.
	// Deprecated: This field is deprecated and replaced by "trustStoreSecretId".
	// This change follows the GoldenGate "Plain Text Fields in Connections" deprecation:
	// https://docs.oracle.com/en-us/iaas/Content/servicechanges.htm#servicechanges_topic-GoldenGate
	TrustStore *string `mandatory:"false" json:"trustStore" sensitive:"true"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the trust store content.
	// Applicable when securityProtocol is TLS.
	// Provide this when the server certificate is not already trusted by the JVM.
	// Note: When provided, 'trustStore' field must not be provided.
	TrustStoreSecretId *string `mandatory:"false" json:"trustStoreSecretId"`

	// Password for the trust store.
	// Applicable when securityProtocol is TLS and a trust store is provided.
	// Required when 'trustStore' is provided.
	// Deprecated: This field is deprecated and replaced by "trustStorePasswordSecretId".
	// This change follows the GoldenGate "Plain Text Fields in Connections" deprecation:
	// https://docs.oracle.com/en-us/iaas/Content/servicechanges.htm#servicechanges_topic-GoldenGate
	TrustStorePassword *string `mandatory:"false" json:"trustStorePassword" sensitive:"true"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the trust store password.
	// Applicable when securityProtocol is TLS and a trust store is provided.
	// Required when 'trustStoreSecretId' is provided.
	// Note: When provided, 'trustStorePassword' field must not be provided.
	TrustStorePasswordSecretId *string `mandatory:"false" json:"trustStorePasswordSecretId"`

	// The base64 encoded content of the client key store used for mutual TLS.
	// Provide this only when the target cluster requires client certificate authentication.
	// Applicable when securityProtocol is TLS.
	// Deprecated: This field is deprecated and replaced by "keyStoreSecretId".
	// This change follows the GoldenGate "Plain Text Fields in Connections" deprecation:
	// https://docs.oracle.com/en-us/iaas/Content/servicechanges.htm#servicechanges_topic-GoldenGate
	KeyStore *string `mandatory:"false" json:"keyStore" sensitive:"true"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the client key store content.
	// Applicable when securityProtocol is TLS.
	// Note: When provided, 'keyStore' field must not be provided.
	KeyStoreSecretId *string `mandatory:"false" json:"keyStoreSecretId"`

	// Password for the client key store.
	// Applicable when securityProtocol is TLS and a client key store is provided.
	// Required when 'keyStore' is provided.
	// Deprecated: This field is deprecated and replaced by "keyStorePasswordSecretId".
	// This change follows the GoldenGate "Plain Text Fields in Connections" deprecation:
	// https://docs.oracle.com/en-us/iaas/Content/servicechanges.htm#servicechanges_topic-GoldenGate
	KeyStorePassword *string `mandatory:"false" json:"keyStorePassword" sensitive:"true"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the client key store password.
	// Applicable when securityProtocol is TLS and a client key store is provided.
	// Required when 'keyStoreSecretId' is provided.
	// Note: When provided, 'keyStorePassword' field must not be provided.
	KeyStorePasswordSecretId *string `mandatory:"false" json:"keyStorePasswordSecretId"`

	// Password for the certificate private key inside the client key store.
	// Provide this only when it differs from the key store password.
	// Applicable when securityProtocol is TLS and a client key store is provided.
	// Deprecated: This field is deprecated and replaced by "sslKeyPasswordSecretId".
	// This change follows the GoldenGate "Plain Text Fields in Connections" deprecation:
	// https://docs.oracle.com/en-us/iaas/Content/servicechanges.htm#servicechanges_topic-GoldenGate
	SslKeyPassword *string `mandatory:"false" json:"sslKeyPassword" sensitive:"true"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the private key password.
	// Applicable when securityProtocol is TLS and a client key store is provided.
	// Note: When provided, 'sslKeyPassword' field must not be provided.
	SslKeyPasswordSecretId *string `mandatory:"false" json:"sslKeyPasswordSecretId"`

	ArtifactLocation UpdateSparkArtifactLocationDetails `mandatory:"false" json:"artifactLocation"`

	// Controls the network traffic direction to the target:
	// SHARED_DEPLOYMENT_ENDPOINT: Network traffic flows from the assigned deployment's private endpoint through the deployment's subnet.
	// DEDICATED_ENDPOINT: A dedicated private endpoint is created in the target VCN subnet for the connection. The subnetId is required when DEDICATED_ENDPOINT networking is selected.
	// SHARED_SERVICE_ENDPOINT: Traffic flows through the Goldengate Service's network to public hosts. Cannot be used for private targets.
	// Deprecated: SHARED_SERVICE_ENDPOINT is deprecated. Use another supported routingMethod value, or update existing connections to use a supported routing method.
	// This change follows the GoldenGate "Plain Text Fields in Connections" deprecation:
	// https://docs.oracle.com/en-us/iaas/Content/servicechanges.htm#servicechanges_topic-GoldenGate
	RoutingMethod RoutingMethodEnum `mandatory:"false" json:"routingMethod,omitempty"`

	// Selects the submission channel.
	// SUBMIT uses the Spark master RPC endpoint.
	// REST uses the Spark standalone REST submission gateway.
	SubmitMode SparkConnectionSubmitModeEnum `mandatory:"false" json:"submitMode,omitempty"`

	// Authentication type for Spark.
	// AES_RPC and SASL_ENCRYPTED securityProtocol values require authenticationType to be SHARED_SECRET.
	// Use SHARED_SECRET only when the target cluster is configured with spark.authenticate=true.
	AuthenticationType SparkConnectionAuthenticationTypeEnum `mandatory:"false" json:"authenticationType,omitempty"`

	// Security protocol for Spark submission traffic.
	// PLAINTEXT uses an unencrypted channel, TLS uses TLS for the selected submission endpoint,
	// AES_RPC uses Spark RPC AES encryption, and SASL_ENCRYPTED uses legacy SASL encryption.
	// AES_RPC and SASL_ENCRYPTED require authenticationType to be SHARED_SECRET.
	SecurityProtocol SparkConnectionSecurityProtocolEnum `mandatory:"false" json:"securityProtocol,omitempty"`

	// Cipher transform used for Spark RPC encryption.
	// Applicable when securityProtocol is AES_RPC. Defaults to AES_CTR_NOPADDING when not provided.
	RpcCipher SparkConnectionRpcCipherEnum `mandatory:"false" json:"rpcCipher,omitempty"`

	// Trust store format used by the Spark client.
	// Applicable when securityProtocol is TLS and a trust store is provided. Defaults to JKS when not provided.
	// Provide this only when a trust store is provided.
	TrustStoreType SparkConnectionTrustStoreTypeEnum `mandatory:"false" json:"trustStoreType,omitempty"`

	// Client key store format used for mutual TLS.
	// Applicable when securityProtocol is TLS and a client key store is provided. Defaults to JKS when not provided.
	// Provide this only when a client key store is provided.
	KeyStoreType SparkConnectionKeyStoreTypeEnum `mandatory:"false" json:"keyStoreType,omitempty"`

	// TLS protocol negotiated by the Spark client.
	// Applicable when securityProtocol is TLS. Defaults to TLS_V1_3 when not provided.
	// Must be supported by the target cluster.
	TlsProtocol SparkConnectionTlsProtocolEnum `mandatory:"false" json:"tlsProtocol,omitempty"`
}

// GetDisplayName returns DisplayName
func (m UpdateSparkConnectionDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m UpdateSparkConnectionDetails) GetDescription() *string {
	return m.Description
}

// GetFreeformTags returns FreeformTags
func (m UpdateSparkConnectionDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m UpdateSparkConnectionDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetVaultId returns VaultId
func (m UpdateSparkConnectionDetails) GetVaultId() *string {
	return m.VaultId
}

// GetKeyId returns KeyId
func (m UpdateSparkConnectionDetails) GetKeyId() *string {
	return m.KeyId
}

// GetNsgIds returns NsgIds
func (m UpdateSparkConnectionDetails) GetNsgIds() []string {
	return m.NsgIds
}

// GetSubnetId returns SubnetId
func (m UpdateSparkConnectionDetails) GetSubnetId() *string {
	return m.SubnetId
}

// GetRoutingMethod returns RoutingMethod
func (m UpdateSparkConnectionDetails) GetRoutingMethod() RoutingMethodEnum {
	return m.RoutingMethod
}

// GetDoesUseSecretIds returns DoesUseSecretIds
func (m UpdateSparkConnectionDetails) GetDoesUseSecretIds() *bool {
	return m.DoesUseSecretIds
}

// GetSecurityAttributes returns SecurityAttributes
func (m UpdateSparkConnectionDetails) GetSecurityAttributes() map[string]map[string]interface{} {
	return m.SecurityAttributes
}

func (m UpdateSparkConnectionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateSparkConnectionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingRoutingMethodEnum(string(m.RoutingMethod)); !ok && m.RoutingMethod != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RoutingMethod: %s. Supported values are: %s.", m.RoutingMethod, strings.Join(GetRoutingMethodEnumStringValues(), ",")))
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
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateSparkConnectionDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateSparkConnectionDetails UpdateSparkConnectionDetails
	s := struct {
		DiscriminatorParam string `json:"connectionType"`
		MarshalTypeUpdateSparkConnectionDetails
	}{
		"SPARK",
		(MarshalTypeUpdateSparkConnectionDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *UpdateSparkConnectionDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName                *string                               `json:"displayName"`
		Description                *string                               `json:"description"`
		FreeformTags               map[string]string                     `json:"freeformTags"`
		DefinedTags                map[string]map[string]interface{}     `json:"definedTags"`
		VaultId                    *string                               `json:"vaultId"`
		KeyId                      *string                               `json:"keyId"`
		NsgIds                     []string                              `json:"nsgIds"`
		SubnetId                   *string                               `json:"subnetId"`
		RoutingMethod              RoutingMethodEnum                     `json:"routingMethod"`
		DoesUseSecretIds           *bool                                 `json:"doesUseSecretIds"`
		SecurityAttributes         map[string]map[string]interface{}     `json:"securityAttributes"`
		SubmitMode                 SparkConnectionSubmitModeEnum         `json:"submitMode"`
		AuthenticationType         SparkConnectionAuthenticationTypeEnum `json:"authenticationType"`
		SecurityProtocol           SparkConnectionSecurityProtocolEnum   `json:"securityProtocol"`
		ClusterEndpoint            *string                               `json:"clusterEndpoint"`
		RestEndpoint               *string                               `json:"restEndpoint"`
		AuthSecret                 *string                               `json:"authSecret"`
		AuthSecretSecretId         *string                               `json:"authSecretSecretId"`
		RpcCipher                  SparkConnectionRpcCipherEnum          `json:"rpcCipher"`
		RestAuthSecretKey          *string                               `json:"restAuthSecretKey"`
		RestAuthSecretKeySecretId  *string                               `json:"restAuthSecretKeySecretId"`
		TrustStore                 *string                               `json:"trustStore"`
		TrustStoreSecretId         *string                               `json:"trustStoreSecretId"`
		TrustStorePassword         *string                               `json:"trustStorePassword"`
		TrustStorePasswordSecretId *string                               `json:"trustStorePasswordSecretId"`
		TrustStoreType             SparkConnectionTrustStoreTypeEnum     `json:"trustStoreType"`
		KeyStore                   *string                               `json:"keyStore"`
		KeyStoreSecretId           *string                               `json:"keyStoreSecretId"`
		KeyStorePassword           *string                               `json:"keyStorePassword"`
		KeyStorePasswordSecretId   *string                               `json:"keyStorePasswordSecretId"`
		KeyStoreType               SparkConnectionKeyStoreTypeEnum       `json:"keyStoreType"`
		SslKeyPassword             *string                               `json:"sslKeyPassword"`
		SslKeyPasswordSecretId     *string                               `json:"sslKeyPasswordSecretId"`
		TlsProtocol                SparkConnectionTlsProtocolEnum        `json:"tlsProtocol"`
		ArtifactLocation           updatesparkartifactlocationdetails    `json:"artifactLocation"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.VaultId = model.VaultId

	m.KeyId = model.KeyId

	m.NsgIds = make([]string, len(model.NsgIds))
	copy(m.NsgIds, model.NsgIds)
	m.SubnetId = model.SubnetId

	m.RoutingMethod = model.RoutingMethod

	m.DoesUseSecretIds = model.DoesUseSecretIds

	m.SecurityAttributes = model.SecurityAttributes

	m.SubmitMode = model.SubmitMode

	m.AuthenticationType = model.AuthenticationType

	m.SecurityProtocol = model.SecurityProtocol

	m.ClusterEndpoint = model.ClusterEndpoint

	m.RestEndpoint = model.RestEndpoint

	m.AuthSecret = model.AuthSecret

	m.AuthSecretSecretId = model.AuthSecretSecretId

	m.RpcCipher = model.RpcCipher

	m.RestAuthSecretKey = model.RestAuthSecretKey

	m.RestAuthSecretKeySecretId = model.RestAuthSecretKeySecretId

	m.TrustStore = model.TrustStore

	m.TrustStoreSecretId = model.TrustStoreSecretId

	m.TrustStorePassword = model.TrustStorePassword

	m.TrustStorePasswordSecretId = model.TrustStorePasswordSecretId

	m.TrustStoreType = model.TrustStoreType

	m.KeyStore = model.KeyStore

	m.KeyStoreSecretId = model.KeyStoreSecretId

	m.KeyStorePassword = model.KeyStorePassword

	m.KeyStorePasswordSecretId = model.KeyStorePasswordSecretId

	m.KeyStoreType = model.KeyStoreType

	m.SslKeyPassword = model.SslKeyPassword

	m.SslKeyPasswordSecretId = model.SslKeyPasswordSecretId

	m.TlsProtocol = model.TlsProtocol

	nn, e = model.ArtifactLocation.UnmarshalPolymorphicJSON(model.ArtifactLocation.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ArtifactLocation = nn.(UpdateSparkArtifactLocationDetails)
	} else {
		m.ArtifactLocation = nil
	}

	return
}
