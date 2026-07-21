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

// OracleAiDataCatalogStorageConnection Represents the metadata of an Oracle AI Data Catalog Storage Connection.
type OracleAiDataCatalogStorageConnection struct {

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

	// The name of the warehouse configured in the Oracle AI Data Catalog deployment.
	// This value corresponds to the target bucket or container name.
	WarehouseName *string `mandatory:"true" json:"warehouseName"`

	WarehouseStorage WarehouseStorage `mandatory:"true" json:"warehouseStorage"`

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

	// The Oracle AI Data Catalog Storage technology type.
	TechnologyType OracleAiDataCatalogStorageConnectionTechnologyTypeEnum `mandatory:"true" json:"technologyType"`

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
func (m OracleAiDataCatalogStorageConnection) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m OracleAiDataCatalogStorageConnection) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m OracleAiDataCatalogStorageConnection) GetDescription() *string {
	return m.Description
}

// GetCompartmentId returns CompartmentId
func (m OracleAiDataCatalogStorageConnection) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetFreeformTags returns FreeformTags
func (m OracleAiDataCatalogStorageConnection) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m OracleAiDataCatalogStorageConnection) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m OracleAiDataCatalogStorageConnection) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetLifecycleState returns LifecycleState
func (m OracleAiDataCatalogStorageConnection) GetLifecycleState() ConnectionLifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m OracleAiDataCatalogStorageConnection) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetTimeCreated returns TimeCreated
func (m OracleAiDataCatalogStorageConnection) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m OracleAiDataCatalogStorageConnection) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLocks returns Locks
func (m OracleAiDataCatalogStorageConnection) GetLocks() []ResourceLock {
	return m.Locks
}

// GetVaultId returns VaultId
func (m OracleAiDataCatalogStorageConnection) GetVaultId() *string {
	return m.VaultId
}

// GetKeyId returns KeyId
func (m OracleAiDataCatalogStorageConnection) GetKeyId() *string {
	return m.KeyId
}

// GetIngressIps returns IngressIps
func (m OracleAiDataCatalogStorageConnection) GetIngressIps() []IngressIpDetails {
	return m.IngressIps
}

// GetNsgIds returns NsgIds
func (m OracleAiDataCatalogStorageConnection) GetNsgIds() []string {
	return m.NsgIds
}

// GetSubnetId returns SubnetId
func (m OracleAiDataCatalogStorageConnection) GetSubnetId() *string {
	return m.SubnetId
}

// GetRoutingMethod returns RoutingMethod
func (m OracleAiDataCatalogStorageConnection) GetRoutingMethod() RoutingMethodEnum {
	return m.RoutingMethod
}

// GetDoesUseSecretIds returns DoesUseSecretIds
func (m OracleAiDataCatalogStorageConnection) GetDoesUseSecretIds() *bool {
	return m.DoesUseSecretIds
}

// GetSubscriptionId returns SubscriptionId
func (m OracleAiDataCatalogStorageConnection) GetSubscriptionId() *string {
	return m.SubscriptionId
}

// GetClusterPlacementGroupId returns ClusterPlacementGroupId
func (m OracleAiDataCatalogStorageConnection) GetClusterPlacementGroupId() *string {
	return m.ClusterPlacementGroupId
}

// GetSecurityAttributes returns SecurityAttributes
func (m OracleAiDataCatalogStorageConnection) GetSecurityAttributes() map[string]map[string]interface{} {
	return m.SecurityAttributes
}

func (m OracleAiDataCatalogStorageConnection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OracleAiDataCatalogStorageConnection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOracleAiDataCatalogStorageConnectionTechnologyTypeEnum(string(m.TechnologyType)); !ok && m.TechnologyType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TechnologyType: %s. Supported values are: %s.", m.TechnologyType, strings.Join(GetOracleAiDataCatalogStorageConnectionTechnologyTypeEnumStringValues(), ",")))
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
func (m OracleAiDataCatalogStorageConnection) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOracleAiDataCatalogStorageConnection OracleAiDataCatalogStorageConnection
	s := struct {
		DiscriminatorParam string `json:"connectionType"`
		MarshalTypeOracleAiDataCatalogStorageConnection
	}{
		"ORACLE_AI_DATA_CATALOG_STORAGE",
		(MarshalTypeOracleAiDataCatalogStorageConnection)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *OracleAiDataCatalogStorageConnection) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description             *string                                                `json:"description"`
		FreeformTags            map[string]string                                      `json:"freeformTags"`
		DefinedTags             map[string]map[string]interface{}                      `json:"definedTags"`
		SystemTags              map[string]map[string]interface{}                      `json:"systemTags"`
		LifecycleDetails        *string                                                `json:"lifecycleDetails"`
		Locks                   []ResourceLock                                         `json:"locks"`
		VaultId                 *string                                                `json:"vaultId"`
		KeyId                   *string                                                `json:"keyId"`
		IngressIps              []IngressIpDetails                                     `json:"ingressIps"`
		NsgIds                  []string                                               `json:"nsgIds"`
		SubnetId                *string                                                `json:"subnetId"`
		RoutingMethod           RoutingMethodEnum                                      `json:"routingMethod"`
		DoesUseSecretIds        *bool                                                  `json:"doesUseSecretIds"`
		SubscriptionId          *string                                                `json:"subscriptionId"`
		ClusterPlacementGroupId *string                                                `json:"clusterPlacementGroupId"`
		SecurityAttributes      map[string]map[string]interface{}                      `json:"securityAttributes"`
		Id                      *string                                                `json:"id"`
		DisplayName             *string                                                `json:"displayName"`
		CompartmentId           *string                                                `json:"compartmentId"`
		LifecycleState          ConnectionLifecycleStateEnum                           `json:"lifecycleState"`
		TimeCreated             *common.SDKTime                                        `json:"timeCreated"`
		TimeUpdated             *common.SDKTime                                        `json:"timeUpdated"`
		TechnologyType          OracleAiDataCatalogStorageConnectionTechnologyTypeEnum `json:"technologyType"`
		WarehouseName           *string                                                `json:"warehouseName"`
		WarehouseStorage        warehousestorage                                       `json:"warehouseStorage"`
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

	m.Id = model.Id

	m.DisplayName = model.DisplayName

	m.CompartmentId = model.CompartmentId

	m.LifecycleState = model.LifecycleState

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.TechnologyType = model.TechnologyType

	m.WarehouseName = model.WarehouseName

	nn, e = model.WarehouseStorage.UnmarshalPolymorphicJSON(model.WarehouseStorage.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.WarehouseStorage = nn.(WarehouseStorage)
	} else {
		m.WarehouseStorage = nil
	}

	return
}

// OracleAiDataCatalogStorageConnectionTechnologyTypeEnum Enum with underlying type: string
type OracleAiDataCatalogStorageConnectionTechnologyTypeEnum string

// Set of constants representing the allowable values for OracleAiDataCatalogStorageConnectionTechnologyTypeEnum
const (
	OracleAiDataCatalogStorageConnectionTechnologyTypeOracleAiDataCatalogStorage OracleAiDataCatalogStorageConnectionTechnologyTypeEnum = "ORACLE_AI_DATA_CATALOG_STORAGE"
)

var mappingOracleAiDataCatalogStorageConnectionTechnologyTypeEnum = map[string]OracleAiDataCatalogStorageConnectionTechnologyTypeEnum{
	"ORACLE_AI_DATA_CATALOG_STORAGE": OracleAiDataCatalogStorageConnectionTechnologyTypeOracleAiDataCatalogStorage,
}

var mappingOracleAiDataCatalogStorageConnectionTechnologyTypeEnumLowerCase = map[string]OracleAiDataCatalogStorageConnectionTechnologyTypeEnum{
	"oracle_ai_data_catalog_storage": OracleAiDataCatalogStorageConnectionTechnologyTypeOracleAiDataCatalogStorage,
}

// GetOracleAiDataCatalogStorageConnectionTechnologyTypeEnumValues Enumerates the set of values for OracleAiDataCatalogStorageConnectionTechnologyTypeEnum
func GetOracleAiDataCatalogStorageConnectionTechnologyTypeEnumValues() []OracleAiDataCatalogStorageConnectionTechnologyTypeEnum {
	values := make([]OracleAiDataCatalogStorageConnectionTechnologyTypeEnum, 0)
	for _, v := range mappingOracleAiDataCatalogStorageConnectionTechnologyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOracleAiDataCatalogStorageConnectionTechnologyTypeEnumStringValues Enumerates the set of values in String for OracleAiDataCatalogStorageConnectionTechnologyTypeEnum
func GetOracleAiDataCatalogStorageConnectionTechnologyTypeEnumStringValues() []string {
	return []string{
		"ORACLE_AI_DATA_CATALOG_STORAGE",
	}
}

// GetMappingOracleAiDataCatalogStorageConnectionTechnologyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOracleAiDataCatalogStorageConnectionTechnologyTypeEnum(val string) (OracleAiDataCatalogStorageConnectionTechnologyTypeEnum, bool) {
	enum, ok := mappingOracleAiDataCatalogStorageConnectionTechnologyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
