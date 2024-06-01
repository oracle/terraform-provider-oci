// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// AzureDataLakeStorageConnection Represents the metadata of a Azure Data Lake Storage Connection.
type AzureDataLakeStorageConnection struct {

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

	// Sets the Azure storage account name.
	AccountName *string `mandatory:"true" json:"accountName"`

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

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the target subnet of the dedicated connection.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// Azure tenant ID of the application. This property is required when 'authenticationType' is set to 'AZURE_ACTIVE_DIRECTORY'.
	// e.g.: 14593954-d337-4a61-a364-9f758c64f97f
	AzureTenantId *string `mandatory:"false" json:"azureTenantId"`

	// Azure client ID of the application. This property is required when 'authenticationType' is set to 'AZURE_ACTIVE_DIRECTORY'.
	// e.g.: 06ecaabf-8b80-4ec8-a0ec-20cbf463703d
	ClientId *string `mandatory:"false" json:"clientId"`

	// Azure Storage service endpoint.
	// e.g: https://test.blob.core.windows.net
	Endpoint *string `mandatory:"false" json:"endpoint"`

	// The Azure Data Lake Storage technology type.
	TechnologyType AzureDataLakeStorageConnectionTechnologyTypeEnum `mandatory:"true" json:"technologyType"`

	// Used authentication mechanism to access Azure Data Lake Storage.
	AuthenticationType AzureDataLakeStorageConnectionAuthenticationTypeEnum `mandatory:"true" json:"authenticationType"`

	// Possible lifecycle states for connection.
	LifecycleState ConnectionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Controls the network traffic direction to the target:
	// SHARED_SERVICE_ENDPOINT: Traffic flows through the Goldengate Service's network to public hosts. Cannot be used for private targets.
	// SHARED_DEPLOYMENT_ENDPOINT: Network traffic flows from the assigned deployment's private endpoint through the deployment's subnet.
	// DEDICATED_ENDPOINT: A dedicated private endpoint is created in the target VCN subnet for the connection. The subnetId is required when DEDICATED_ENDPOINT networking is selected.
	RoutingMethod RoutingMethodEnum `mandatory:"false" json:"routingMethod,omitempty"`
}

// GetId returns Id
func (m AzureDataLakeStorageConnection) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m AzureDataLakeStorageConnection) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m AzureDataLakeStorageConnection) GetDescription() *string {
	return m.Description
}

// GetCompartmentId returns CompartmentId
func (m AzureDataLakeStorageConnection) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetFreeformTags returns FreeformTags
func (m AzureDataLakeStorageConnection) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m AzureDataLakeStorageConnection) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m AzureDataLakeStorageConnection) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetLifecycleState returns LifecycleState
func (m AzureDataLakeStorageConnection) GetLifecycleState() ConnectionLifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m AzureDataLakeStorageConnection) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetTimeCreated returns TimeCreated
func (m AzureDataLakeStorageConnection) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m AzureDataLakeStorageConnection) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLocks returns Locks
func (m AzureDataLakeStorageConnection) GetLocks() []ResourceLock {
	return m.Locks
}

// GetVaultId returns VaultId
func (m AzureDataLakeStorageConnection) GetVaultId() *string {
	return m.VaultId
}

// GetKeyId returns KeyId
func (m AzureDataLakeStorageConnection) GetKeyId() *string {
	return m.KeyId
}

// GetIngressIps returns IngressIps
func (m AzureDataLakeStorageConnection) GetIngressIps() []IngressIpDetails {
	return m.IngressIps
}

// GetNsgIds returns NsgIds
func (m AzureDataLakeStorageConnection) GetNsgIds() []string {
	return m.NsgIds
}

// GetSubnetId returns SubnetId
func (m AzureDataLakeStorageConnection) GetSubnetId() *string {
	return m.SubnetId
}

// GetRoutingMethod returns RoutingMethod
func (m AzureDataLakeStorageConnection) GetRoutingMethod() RoutingMethodEnum {
	return m.RoutingMethod
}

func (m AzureDataLakeStorageConnection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AzureDataLakeStorageConnection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAzureDataLakeStorageConnectionTechnologyTypeEnum(string(m.TechnologyType)); !ok && m.TechnologyType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TechnologyType: %s. Supported values are: %s.", m.TechnologyType, strings.Join(GetAzureDataLakeStorageConnectionTechnologyTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAzureDataLakeStorageConnectionAuthenticationTypeEnum(string(m.AuthenticationType)); !ok && m.AuthenticationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AuthenticationType: %s. Supported values are: %s.", m.AuthenticationType, strings.Join(GetAzureDataLakeStorageConnectionAuthenticationTypeEnumStringValues(), ",")))
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
func (m AzureDataLakeStorageConnection) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAzureDataLakeStorageConnection AzureDataLakeStorageConnection
	s := struct {
		DiscriminatorParam string `json:"connectionType"`
		MarshalTypeAzureDataLakeStorageConnection
	}{
		"AZURE_DATA_LAKE_STORAGE",
		(MarshalTypeAzureDataLakeStorageConnection)(m),
	}

	return json.Marshal(&s)
}

// AzureDataLakeStorageConnectionTechnologyTypeEnum Enum with underlying type: string
type AzureDataLakeStorageConnectionTechnologyTypeEnum string

// Set of constants representing the allowable values for AzureDataLakeStorageConnectionTechnologyTypeEnum
const (
	AzureDataLakeStorageConnectionTechnologyTypeAzureDataLakeStorage AzureDataLakeStorageConnectionTechnologyTypeEnum = "AZURE_DATA_LAKE_STORAGE"
)

var mappingAzureDataLakeStorageConnectionTechnologyTypeEnum = map[string]AzureDataLakeStorageConnectionTechnologyTypeEnum{
	"AZURE_DATA_LAKE_STORAGE": AzureDataLakeStorageConnectionTechnologyTypeAzureDataLakeStorage,
}

var mappingAzureDataLakeStorageConnectionTechnologyTypeEnumLowerCase = map[string]AzureDataLakeStorageConnectionTechnologyTypeEnum{
	"azure_data_lake_storage": AzureDataLakeStorageConnectionTechnologyTypeAzureDataLakeStorage,
}

// GetAzureDataLakeStorageConnectionTechnologyTypeEnumValues Enumerates the set of values for AzureDataLakeStorageConnectionTechnologyTypeEnum
func GetAzureDataLakeStorageConnectionTechnologyTypeEnumValues() []AzureDataLakeStorageConnectionTechnologyTypeEnum {
	values := make([]AzureDataLakeStorageConnectionTechnologyTypeEnum, 0)
	for _, v := range mappingAzureDataLakeStorageConnectionTechnologyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAzureDataLakeStorageConnectionTechnologyTypeEnumStringValues Enumerates the set of values in String for AzureDataLakeStorageConnectionTechnologyTypeEnum
func GetAzureDataLakeStorageConnectionTechnologyTypeEnumStringValues() []string {
	return []string{
		"AZURE_DATA_LAKE_STORAGE",
	}
}

// GetMappingAzureDataLakeStorageConnectionTechnologyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAzureDataLakeStorageConnectionTechnologyTypeEnum(val string) (AzureDataLakeStorageConnectionTechnologyTypeEnum, bool) {
	enum, ok := mappingAzureDataLakeStorageConnectionTechnologyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AzureDataLakeStorageConnectionAuthenticationTypeEnum Enum with underlying type: string
type AzureDataLakeStorageConnectionAuthenticationTypeEnum string

// Set of constants representing the allowable values for AzureDataLakeStorageConnectionAuthenticationTypeEnum
const (
	AzureDataLakeStorageConnectionAuthenticationTypeSharedKey             AzureDataLakeStorageConnectionAuthenticationTypeEnum = "SHARED_KEY"
	AzureDataLakeStorageConnectionAuthenticationTypeSharedAccessSignature AzureDataLakeStorageConnectionAuthenticationTypeEnum = "SHARED_ACCESS_SIGNATURE"
	AzureDataLakeStorageConnectionAuthenticationTypeAzureActiveDirectory  AzureDataLakeStorageConnectionAuthenticationTypeEnum = "AZURE_ACTIVE_DIRECTORY"
)

var mappingAzureDataLakeStorageConnectionAuthenticationTypeEnum = map[string]AzureDataLakeStorageConnectionAuthenticationTypeEnum{
	"SHARED_KEY":              AzureDataLakeStorageConnectionAuthenticationTypeSharedKey,
	"SHARED_ACCESS_SIGNATURE": AzureDataLakeStorageConnectionAuthenticationTypeSharedAccessSignature,
	"AZURE_ACTIVE_DIRECTORY":  AzureDataLakeStorageConnectionAuthenticationTypeAzureActiveDirectory,
}

var mappingAzureDataLakeStorageConnectionAuthenticationTypeEnumLowerCase = map[string]AzureDataLakeStorageConnectionAuthenticationTypeEnum{
	"shared_key":              AzureDataLakeStorageConnectionAuthenticationTypeSharedKey,
	"shared_access_signature": AzureDataLakeStorageConnectionAuthenticationTypeSharedAccessSignature,
	"azure_active_directory":  AzureDataLakeStorageConnectionAuthenticationTypeAzureActiveDirectory,
}

// GetAzureDataLakeStorageConnectionAuthenticationTypeEnumValues Enumerates the set of values for AzureDataLakeStorageConnectionAuthenticationTypeEnum
func GetAzureDataLakeStorageConnectionAuthenticationTypeEnumValues() []AzureDataLakeStorageConnectionAuthenticationTypeEnum {
	values := make([]AzureDataLakeStorageConnectionAuthenticationTypeEnum, 0)
	for _, v := range mappingAzureDataLakeStorageConnectionAuthenticationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAzureDataLakeStorageConnectionAuthenticationTypeEnumStringValues Enumerates the set of values in String for AzureDataLakeStorageConnectionAuthenticationTypeEnum
func GetAzureDataLakeStorageConnectionAuthenticationTypeEnumStringValues() []string {
	return []string{
		"SHARED_KEY",
		"SHARED_ACCESS_SIGNATURE",
		"AZURE_ACTIVE_DIRECTORY",
	}
}

// GetMappingAzureDataLakeStorageConnectionAuthenticationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAzureDataLakeStorageConnectionAuthenticationTypeEnum(val string) (AzureDataLakeStorageConnectionAuthenticationTypeEnum, bool) {
	enum, ok := mappingAzureDataLakeStorageConnectionAuthenticationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
