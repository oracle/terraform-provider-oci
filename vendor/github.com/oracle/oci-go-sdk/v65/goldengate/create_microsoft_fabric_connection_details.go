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

// CreateMicrosoftFabricConnectionDetails The information about a new Microsoft Fabric Connection.
type CreateMicrosoftFabricConnectionDetails struct {

	// An object's Display Name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment being referenced.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Azure tenant ID of the application.
	// e.g.: 14593954-d337-4a61-a364-9f758c64f97f
	TenantId *string `mandatory:"true" json:"tenantId"`

	// Azure client ID of the application.
	// e.g.: 06ecaabf-8b80-4ec8-a0ec-20cbf463703d
	ClientId *string `mandatory:"true" json:"clientId"`

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

	// Client secret associated with the client id.
	// Deprecated: This field is deprecated and replaced by "clientSecretSecretId". This field will be removed after February 15 2026.
	ClientSecret *string `mandatory:"false" json:"clientSecret"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the client secret is stored.
	// Note: When provided, 'clientSecret' field must not be provided.
	ClientSecretSecretId *string `mandatory:"false" json:"clientSecretSecretId"`

	// Optional Microsoft Fabric service endpoint.
	// Default value: https://onelake.dfs.fabric.microsoft.com
	Endpoint *string `mandatory:"false" json:"endpoint"`

	// Controls the network traffic direction to the target:
	// SHARED_SERVICE_ENDPOINT: Traffic flows through the Goldengate Service's network to public hosts. Cannot be used for private targets.
	// SHARED_DEPLOYMENT_ENDPOINT: Network traffic flows from the assigned deployment's private endpoint through the deployment's subnet.
	// DEDICATED_ENDPOINT: A dedicated private endpoint is created in the target VCN subnet for the connection. The subnetId is required when DEDICATED_ENDPOINT networking is selected.
	RoutingMethod RoutingMethodEnum `mandatory:"false" json:"routingMethod,omitempty"`

	// The Microsoft Fabric technology type.
	TechnologyType MicrosoftFabricConnectionTechnologyTypeEnum `mandatory:"true" json:"technologyType"`
}

// GetDisplayName returns DisplayName
func (m CreateMicrosoftFabricConnectionDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m CreateMicrosoftFabricConnectionDetails) GetDescription() *string {
	return m.Description
}

// GetCompartmentId returns CompartmentId
func (m CreateMicrosoftFabricConnectionDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetFreeformTags returns FreeformTags
func (m CreateMicrosoftFabricConnectionDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateMicrosoftFabricConnectionDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetLocks returns Locks
func (m CreateMicrosoftFabricConnectionDetails) GetLocks() []AddResourceLockDetails {
	return m.Locks
}

// GetVaultId returns VaultId
func (m CreateMicrosoftFabricConnectionDetails) GetVaultId() *string {
	return m.VaultId
}

// GetKeyId returns KeyId
func (m CreateMicrosoftFabricConnectionDetails) GetKeyId() *string {
	return m.KeyId
}

// GetNsgIds returns NsgIds
func (m CreateMicrosoftFabricConnectionDetails) GetNsgIds() []string {
	return m.NsgIds
}

// GetSubnetId returns SubnetId
func (m CreateMicrosoftFabricConnectionDetails) GetSubnetId() *string {
	return m.SubnetId
}

// GetRoutingMethod returns RoutingMethod
func (m CreateMicrosoftFabricConnectionDetails) GetRoutingMethod() RoutingMethodEnum {
	return m.RoutingMethod
}

// GetDoesUseSecretIds returns DoesUseSecretIds
func (m CreateMicrosoftFabricConnectionDetails) GetDoesUseSecretIds() *bool {
	return m.DoesUseSecretIds
}

// GetSubscriptionId returns SubscriptionId
func (m CreateMicrosoftFabricConnectionDetails) GetSubscriptionId() *string {
	return m.SubscriptionId
}

// GetClusterPlacementGroupId returns ClusterPlacementGroupId
func (m CreateMicrosoftFabricConnectionDetails) GetClusterPlacementGroupId() *string {
	return m.ClusterPlacementGroupId
}

// GetSecurityAttributes returns SecurityAttributes
func (m CreateMicrosoftFabricConnectionDetails) GetSecurityAttributes() map[string]map[string]interface{} {
	return m.SecurityAttributes
}

func (m CreateMicrosoftFabricConnectionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateMicrosoftFabricConnectionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingRoutingMethodEnum(string(m.RoutingMethod)); !ok && m.RoutingMethod != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RoutingMethod: %s. Supported values are: %s.", m.RoutingMethod, strings.Join(GetRoutingMethodEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMicrosoftFabricConnectionTechnologyTypeEnum(string(m.TechnologyType)); !ok && m.TechnologyType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TechnologyType: %s. Supported values are: %s.", m.TechnologyType, strings.Join(GetMicrosoftFabricConnectionTechnologyTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateMicrosoftFabricConnectionDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateMicrosoftFabricConnectionDetails CreateMicrosoftFabricConnectionDetails
	s := struct {
		DiscriminatorParam string `json:"connectionType"`
		MarshalTypeCreateMicrosoftFabricConnectionDetails
	}{
		"MICROSOFT_FABRIC",
		(MarshalTypeCreateMicrosoftFabricConnectionDetails)(m),
	}

	return json.Marshal(&s)
}
