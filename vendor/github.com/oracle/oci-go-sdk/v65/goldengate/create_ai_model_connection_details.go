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

// CreateAiModelConnectionDetails The information about a new AI Model Connection.
type CreateAiModelConnectionDetails struct {

	// An object's Display Name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment being referenced.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// AI model identifier.
	ModelKey *string `mandatory:"true" json:"modelKey"`

	AuthDetails CreateAiModelAuthDetails `mandatory:"true" json:"authDetails"`

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

	// Refers to the customer's vault OCID.
	// If provided, it references a vault where GoldenGate can manage secrets. Customers must add policies to permit GoldenGate
	// to manage secrets contained within this vault.
	VaultId *string `mandatory:"false" json:"vaultId"`

	// Refers to the customer's master key OCID.
	// If provided, it references a key to manage secrets. Customers must add policies to permit GoldenGate to use this key.
	KeyId *string `mandatory:"false" json:"keyId"`

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

	// Maximum number of input characters supported by this AI model connection.
	MaxInputChars *int `mandatory:"false" json:"maxInputChars"`

	// AI Provider type used by the AI Model Connection.
	ProviderType CreateAiModelConnectionDetailsProviderTypeEnum `mandatory:"true" json:"providerType"`

	// Controls the network traffic direction to the target:
	// SHARED_DEPLOYMENT_ENDPOINT: Network traffic flows from the assigned deployment's private endpoint through the deployment's subnet.
	// DEDICATED_ENDPOINT: A dedicated private endpoint is created in the target VCN subnet for the connection. The subnetId is required when DEDICATED_ENDPOINT networking is selected.
	// SHARED_SERVICE_ENDPOINT: Traffic flows through the Goldengate Service's network to public hosts. Cannot be used for private targets.
	// Deprecated: SHARED_SERVICE_ENDPOINT is deprecated. Use another supported routingMethod value, or update existing connections to use a supported routing method.
	// This change follows the GoldenGate "Plain Text Fields in Connections" deprecation:
	// https://docs.oracle.com/en-us/iaas/Content/servicechanges.htm#servicechanges_topic-GoldenGate
	RoutingMethod RoutingMethodEnum `mandatory:"false" json:"routingMethod,omitempty"`

	// The AI Model technology type.
	TechnologyType AiModelConnectionTechnologyTypeEnum `mandatory:"true" json:"technologyType"`
}

// GetDisplayName returns DisplayName
func (m CreateAiModelConnectionDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m CreateAiModelConnectionDetails) GetDescription() *string {
	return m.Description
}

// GetCompartmentId returns CompartmentId
func (m CreateAiModelConnectionDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetFreeformTags returns FreeformTags
func (m CreateAiModelConnectionDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateAiModelConnectionDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetLocks returns Locks
func (m CreateAiModelConnectionDetails) GetLocks() []AddResourceLockDetails {
	return m.Locks
}

// GetVaultId returns VaultId
func (m CreateAiModelConnectionDetails) GetVaultId() *string {
	return m.VaultId
}

// GetKeyId returns KeyId
func (m CreateAiModelConnectionDetails) GetKeyId() *string {
	return m.KeyId
}

// GetNsgIds returns NsgIds
func (m CreateAiModelConnectionDetails) GetNsgIds() []string {
	return m.NsgIds
}

// GetSubnetId returns SubnetId
func (m CreateAiModelConnectionDetails) GetSubnetId() *string {
	return m.SubnetId
}

// GetRoutingMethod returns RoutingMethod
func (m CreateAiModelConnectionDetails) GetRoutingMethod() RoutingMethodEnum {
	return m.RoutingMethod
}

// GetDoesUseSecretIds returns DoesUseSecretIds
func (m CreateAiModelConnectionDetails) GetDoesUseSecretIds() *bool {
	return m.DoesUseSecretIds
}

// GetSubscriptionId returns SubscriptionId
func (m CreateAiModelConnectionDetails) GetSubscriptionId() *string {
	return m.SubscriptionId
}

// GetClusterPlacementGroupId returns ClusterPlacementGroupId
func (m CreateAiModelConnectionDetails) GetClusterPlacementGroupId() *string {
	return m.ClusterPlacementGroupId
}

// GetSecurityAttributes returns SecurityAttributes
func (m CreateAiModelConnectionDetails) GetSecurityAttributes() map[string]map[string]interface{} {
	return m.SecurityAttributes
}

func (m CreateAiModelConnectionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateAiModelConnectionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateAiModelConnectionDetailsProviderTypeEnum(string(m.ProviderType)); !ok && m.ProviderType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ProviderType: %s. Supported values are: %s.", m.ProviderType, strings.Join(GetCreateAiModelConnectionDetailsProviderTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingRoutingMethodEnum(string(m.RoutingMethod)); !ok && m.RoutingMethod != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RoutingMethod: %s. Supported values are: %s.", m.RoutingMethod, strings.Join(GetRoutingMethodEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAiModelConnectionTechnologyTypeEnum(string(m.TechnologyType)); !ok && m.TechnologyType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TechnologyType: %s. Supported values are: %s.", m.TechnologyType, strings.Join(GetAiModelConnectionTechnologyTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateAiModelConnectionDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateAiModelConnectionDetails CreateAiModelConnectionDetails
	s := struct {
		DiscriminatorParam string `json:"connectionType"`
		MarshalTypeCreateAiModelConnectionDetails
	}{
		"AI_MODEL",
		(MarshalTypeCreateAiModelConnectionDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CreateAiModelConnectionDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description             *string                                        `json:"description"`
		FreeformTags            map[string]string                              `json:"freeformTags"`
		DefinedTags             map[string]map[string]interface{}              `json:"definedTags"`
		Locks                   []AddResourceLockDetails                       `json:"locks"`
		VaultId                 *string                                        `json:"vaultId"`
		KeyId                   *string                                        `json:"keyId"`
		NsgIds                  []string                                       `json:"nsgIds"`
		SubnetId                *string                                        `json:"subnetId"`
		RoutingMethod           RoutingMethodEnum                              `json:"routingMethod"`
		DoesUseSecretIds        *bool                                          `json:"doesUseSecretIds"`
		SubscriptionId          *string                                        `json:"subscriptionId"`
		ClusterPlacementGroupId *string                                        `json:"clusterPlacementGroupId"`
		SecurityAttributes      map[string]map[string]interface{}              `json:"securityAttributes"`
		MaxInputChars           *int                                           `json:"maxInputChars"`
		DisplayName             *string                                        `json:"displayName"`
		CompartmentId           *string                                        `json:"compartmentId"`
		TechnologyType          AiModelConnectionTechnologyTypeEnum            `json:"technologyType"`
		ProviderType            CreateAiModelConnectionDetailsProviderTypeEnum `json:"providerType"`
		ModelKey                *string                                        `json:"modelKey"`
		AuthDetails             createaimodelauthdetails                       `json:"authDetails"`
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

	m.MaxInputChars = model.MaxInputChars

	m.DisplayName = model.DisplayName

	m.CompartmentId = model.CompartmentId

	m.TechnologyType = model.TechnologyType

	m.ProviderType = model.ProviderType

	m.ModelKey = model.ModelKey

	nn, e = model.AuthDetails.UnmarshalPolymorphicJSON(model.AuthDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.AuthDetails = nn.(CreateAiModelAuthDetails)
	} else {
		m.AuthDetails = nil
	}

	return
}

// CreateAiModelConnectionDetailsProviderTypeEnum Enum with underlying type: string
type CreateAiModelConnectionDetailsProviderTypeEnum string

// Set of constants representing the allowable values for CreateAiModelConnectionDetailsProviderTypeEnum
const (
	CreateAiModelConnectionDetailsProviderTypeOciGenerativeAi CreateAiModelConnectionDetailsProviderTypeEnum = "OCI_GENERATIVE_AI"
	CreateAiModelConnectionDetailsProviderTypeGemini          CreateAiModelConnectionDetailsProviderTypeEnum = "GEMINI"
	CreateAiModelConnectionDetailsProviderTypeOpenAi          CreateAiModelConnectionDetailsProviderTypeEnum = "OPEN_AI"
	CreateAiModelConnectionDetailsProviderTypeVoyageAi        CreateAiModelConnectionDetailsProviderTypeEnum = "VOYAGE_AI"
)

var mappingCreateAiModelConnectionDetailsProviderTypeEnum = map[string]CreateAiModelConnectionDetailsProviderTypeEnum{
	"OCI_GENERATIVE_AI": CreateAiModelConnectionDetailsProviderTypeOciGenerativeAi,
	"GEMINI":            CreateAiModelConnectionDetailsProviderTypeGemini,
	"OPEN_AI":           CreateAiModelConnectionDetailsProviderTypeOpenAi,
	"VOYAGE_AI":         CreateAiModelConnectionDetailsProviderTypeVoyageAi,
}

var mappingCreateAiModelConnectionDetailsProviderTypeEnumLowerCase = map[string]CreateAiModelConnectionDetailsProviderTypeEnum{
	"oci_generative_ai": CreateAiModelConnectionDetailsProviderTypeOciGenerativeAi,
	"gemini":            CreateAiModelConnectionDetailsProviderTypeGemini,
	"open_ai":           CreateAiModelConnectionDetailsProviderTypeOpenAi,
	"voyage_ai":         CreateAiModelConnectionDetailsProviderTypeVoyageAi,
}

// GetCreateAiModelConnectionDetailsProviderTypeEnumValues Enumerates the set of values for CreateAiModelConnectionDetailsProviderTypeEnum
func GetCreateAiModelConnectionDetailsProviderTypeEnumValues() []CreateAiModelConnectionDetailsProviderTypeEnum {
	values := make([]CreateAiModelConnectionDetailsProviderTypeEnum, 0)
	for _, v := range mappingCreateAiModelConnectionDetailsProviderTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateAiModelConnectionDetailsProviderTypeEnumStringValues Enumerates the set of values in String for CreateAiModelConnectionDetailsProviderTypeEnum
func GetCreateAiModelConnectionDetailsProviderTypeEnumStringValues() []string {
	return []string{
		"OCI_GENERATIVE_AI",
		"GEMINI",
		"OPEN_AI",
		"VOYAGE_AI",
	}
}

// GetMappingCreateAiModelConnectionDetailsProviderTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateAiModelConnectionDetailsProviderTypeEnum(val string) (CreateAiModelConnectionDetailsProviderTypeEnum, bool) {
	enum, ok := mappingCreateAiModelConnectionDetailsProviderTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
