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

// CreateDatabricksEngineConnectionDetails The information about a new Databricks Engine Connection.
type CreateDatabricksEngineConnectionDetails struct {

	// An object's Display Name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment being referenced.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Databricks workspace base URL for the REST API.
	// e.g.: 'https://<workspace>.azuredatabricks.net/api/2.0/'
	// e.g.: 'https://<workspace>.cloud.databricks.com/api/2.0/'
	WorkspaceUrl *string `mandatory:"true" json:"workspaceUrl"`

	ArtifactLocation CreateDatabricksEngineArtifactLocationDetails `mandatory:"true" json:"artifactLocation"`

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

	// Personal access token used to connect to the Databricks Engine REST API.
	// Only applicable for authenticationType == PERSONAL_ACCESS_TOKEN.
	// Deprecated: This field is deprecated and replaced by "authTokenSecretId".
	// This change follows the GoldenGate "Plain Text Fields in Connections" deprecation:
	// https://docs.oracle.com/en-us/iaas/Content/servicechanges.htm#servicechanges_topic-GoldenGate
	AuthToken *string `mandatory:"false" json:"authToken" sensitive:"true"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the personal access token is stored.
	// Only applicable for authenticationType == PERSONAL_ACCESS_TOKEN.
	// Note: When provided, 'authToken' field must not be provided.
	AuthTokenSecretId *string `mandatory:"false" json:"authTokenSecretId"`

	// OAuth client id, only applicable for authenticationType == OAUTH_M2M
	ClientId *string `mandatory:"false" json:"clientId"`

	// OAuth client secret, only applicable for authenticationType == OAUTH_M2M
	// Deprecated: This field is deprecated and replaced by "clientSecretSecretId".
	// This change follows the GoldenGate "Plain Text Fields in Connections" deprecation:
	// https://docs.oracle.com/en-us/iaas/Content/servicechanges.htm#servicechanges_topic-GoldenGate
	ClientSecret *string `mandatory:"false" json:"clientSecret" sensitive:"true"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the client secret is stored.
	// Only applicable for authenticationType == OAUTH_M2M.
	// Note: When provided, 'clientSecret' field must not be provided.
	ClientSecretSecretId *string `mandatory:"false" json:"clientSecretSecretId"`

	// Controls the network traffic direction to the target:
	// SHARED_DEPLOYMENT_ENDPOINT: Network traffic flows from the assigned deployment's private endpoint through the deployment's subnet.
	// DEDICATED_ENDPOINT: A dedicated private endpoint is created in the target VCN subnet for the connection. The subnetId is required when DEDICATED_ENDPOINT networking is selected.
	// SHARED_SERVICE_ENDPOINT: Traffic flows through the Goldengate Service's network to public hosts. Cannot be used for private targets.
	// Deprecated: SHARED_SERVICE_ENDPOINT is deprecated. Use another supported routingMethod value, or update existing connections to use a supported routing method.
	// This change follows the GoldenGate "Plain Text Fields in Connections" deprecation:
	// https://docs.oracle.com/en-us/iaas/Content/servicechanges.htm#servicechanges_topic-GoldenGate
	RoutingMethod RoutingMethodEnum `mandatory:"false" json:"routingMethod,omitempty"`

	// The Databricks Engine technology type.
	TechnologyType DatabricksEngineConnectionTechnologyTypeEnum `mandatory:"true" json:"technologyType"`

	// Used authentication mechanism to access the Databricks Engine REST API.
	AuthenticationType DatabricksEngineConnectionAuthenticationTypeEnum `mandatory:"true" json:"authenticationType"`
}

// GetDisplayName returns DisplayName
func (m CreateDatabricksEngineConnectionDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m CreateDatabricksEngineConnectionDetails) GetDescription() *string {
	return m.Description
}

// GetCompartmentId returns CompartmentId
func (m CreateDatabricksEngineConnectionDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetFreeformTags returns FreeformTags
func (m CreateDatabricksEngineConnectionDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateDatabricksEngineConnectionDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetLocks returns Locks
func (m CreateDatabricksEngineConnectionDetails) GetLocks() []AddResourceLockDetails {
	return m.Locks
}

// GetVaultId returns VaultId
func (m CreateDatabricksEngineConnectionDetails) GetVaultId() *string {
	return m.VaultId
}

// GetKeyId returns KeyId
func (m CreateDatabricksEngineConnectionDetails) GetKeyId() *string {
	return m.KeyId
}

// GetNsgIds returns NsgIds
func (m CreateDatabricksEngineConnectionDetails) GetNsgIds() []string {
	return m.NsgIds
}

// GetSubnetId returns SubnetId
func (m CreateDatabricksEngineConnectionDetails) GetSubnetId() *string {
	return m.SubnetId
}

// GetRoutingMethod returns RoutingMethod
func (m CreateDatabricksEngineConnectionDetails) GetRoutingMethod() RoutingMethodEnum {
	return m.RoutingMethod
}

// GetDoesUseSecretIds returns DoesUseSecretIds
func (m CreateDatabricksEngineConnectionDetails) GetDoesUseSecretIds() *bool {
	return m.DoesUseSecretIds
}

// GetSubscriptionId returns SubscriptionId
func (m CreateDatabricksEngineConnectionDetails) GetSubscriptionId() *string {
	return m.SubscriptionId
}

// GetClusterPlacementGroupId returns ClusterPlacementGroupId
func (m CreateDatabricksEngineConnectionDetails) GetClusterPlacementGroupId() *string {
	return m.ClusterPlacementGroupId
}

// GetSecurityAttributes returns SecurityAttributes
func (m CreateDatabricksEngineConnectionDetails) GetSecurityAttributes() map[string]map[string]interface{} {
	return m.SecurityAttributes
}

func (m CreateDatabricksEngineConnectionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDatabricksEngineConnectionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingRoutingMethodEnum(string(m.RoutingMethod)); !ok && m.RoutingMethod != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RoutingMethod: %s. Supported values are: %s.", m.RoutingMethod, strings.Join(GetRoutingMethodEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDatabricksEngineConnectionTechnologyTypeEnum(string(m.TechnologyType)); !ok && m.TechnologyType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TechnologyType: %s. Supported values are: %s.", m.TechnologyType, strings.Join(GetDatabricksEngineConnectionTechnologyTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDatabricksEngineConnectionAuthenticationTypeEnum(string(m.AuthenticationType)); !ok && m.AuthenticationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AuthenticationType: %s. Supported values are: %s.", m.AuthenticationType, strings.Join(GetDatabricksEngineConnectionAuthenticationTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateDatabricksEngineConnectionDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateDatabricksEngineConnectionDetails CreateDatabricksEngineConnectionDetails
	s := struct {
		DiscriminatorParam string `json:"connectionType"`
		MarshalTypeCreateDatabricksEngineConnectionDetails
	}{
		"DATABRICKS_ENGINE",
		(MarshalTypeCreateDatabricksEngineConnectionDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CreateDatabricksEngineConnectionDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description             *string                                          `json:"description"`
		FreeformTags            map[string]string                                `json:"freeformTags"`
		DefinedTags             map[string]map[string]interface{}                `json:"definedTags"`
		Locks                   []AddResourceLockDetails                         `json:"locks"`
		VaultId                 *string                                          `json:"vaultId"`
		KeyId                   *string                                          `json:"keyId"`
		NsgIds                  []string                                         `json:"nsgIds"`
		SubnetId                *string                                          `json:"subnetId"`
		RoutingMethod           RoutingMethodEnum                                `json:"routingMethod"`
		DoesUseSecretIds        *bool                                            `json:"doesUseSecretIds"`
		SubscriptionId          *string                                          `json:"subscriptionId"`
		ClusterPlacementGroupId *string                                          `json:"clusterPlacementGroupId"`
		SecurityAttributes      map[string]map[string]interface{}                `json:"securityAttributes"`
		AuthToken               *string                                          `json:"authToken"`
		AuthTokenSecretId       *string                                          `json:"authTokenSecretId"`
		ClientId                *string                                          `json:"clientId"`
		ClientSecret            *string                                          `json:"clientSecret"`
		ClientSecretSecretId    *string                                          `json:"clientSecretSecretId"`
		DisplayName             *string                                          `json:"displayName"`
		CompartmentId           *string                                          `json:"compartmentId"`
		TechnologyType          DatabricksEngineConnectionTechnologyTypeEnum     `json:"technologyType"`
		AuthenticationType      DatabricksEngineConnectionAuthenticationTypeEnum `json:"authenticationType"`
		WorkspaceUrl            *string                                          `json:"workspaceUrl"`
		ArtifactLocation        createdatabricksengineartifactlocationdetails    `json:"artifactLocation"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.Locks = make([]AddResourceLockDetails, len(model.Locks))
	copy(m.Locks, model.Locks)
	m.VaultId = model.VaultId

	m.KeyId = model.KeyId

	m.NsgIds = make([]string, len(model.NsgIds))
	copy(m.NsgIds, model.NsgIds)
	m.SubnetId = model.SubnetId

	m.RoutingMethod = model.RoutingMethod

	m.DoesUseSecretIds = model.DoesUseSecretIds

	m.SubscriptionId = model.SubscriptionId

	m.ClusterPlacementGroupId = model.ClusterPlacementGroupId

	m.SecurityAttributes = model.SecurityAttributes

	m.AuthToken = model.AuthToken

	m.AuthTokenSecretId = model.AuthTokenSecretId

	m.ClientId = model.ClientId

	m.ClientSecret = model.ClientSecret

	m.ClientSecretSecretId = model.ClientSecretSecretId

	m.DisplayName = model.DisplayName

	m.CompartmentId = model.CompartmentId

	m.TechnologyType = model.TechnologyType

	m.AuthenticationType = model.AuthenticationType

	m.WorkspaceUrl = model.WorkspaceUrl

	nn, e = model.ArtifactLocation.UnmarshalPolymorphicJSON(model.ArtifactLocation.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ArtifactLocation = nn.(CreateDatabricksEngineArtifactLocationDetails)
	} else {
		m.ArtifactLocation = nil
	}

	return
}
