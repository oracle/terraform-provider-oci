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

// CreateAzureDataLakeStorageConnectionDetails The information about a new Azure Data Lake Storage Connection.
type CreateAzureDataLakeStorageConnectionDetails struct {

	// An object's Display Name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment being referenced.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

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

	// Azure storage account key. This property is required when 'authenticationType' is set to 'SHARED_KEY'.
	// e.g.: pa3WbhVATzj56xD4DH1VjOUhApRGEGHvOo58eQJVWIzX+j8j4CUVFcTjpIqDSRaSa1Wo2LbWY5at+AStEgLOIQ==
	// Deprecated: This field is deprecated and replaced by "accountKeySecretId". This field will be removed after February 15 2026.
	AccountKey *string `mandatory:"false" json:"accountKey"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the account key is stored.
	// Note: When provided, 'accountKey' field must not be provided.
	AccountKeySecretId *string `mandatory:"false" json:"accountKeySecretId"`

	// Credential that uses a shared access signature (SAS) to authenticate to an Azure Service. This property is
	// required when 'authenticationType' is set to 'SHARED_ACCESS_SIGNATURE'.
	// e.g.: ?sv=2020-06-08&ss=bfqt&srt=sco&sp=rwdlacupyx&se=2020-09-10T20:27:28Z&st=2022-08-05T12:27:28Z&spr=https&sig=C1IgHsiLBmTSStYkXXGLTP8it0xBrArcgCqOsZbXwIQ%3D
	// Deprecated: This field is deprecated and replaced by "sasTokenSecretId". This field will be removed after February 15 2026.
	SasToken *string `mandatory:"false" json:"sasToken"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the sas token is stored.
	// Note: When provided, 'sasToken' field must not be provided.
	SasTokenSecretId *string `mandatory:"false" json:"sasTokenSecretId"`

	// Azure tenant ID of the application. This property is required when 'authenticationType' is set to 'AZURE_ACTIVE_DIRECTORY'.
	// e.g.: 14593954-d337-4a61-a364-9f758c64f97f
	AzureTenantId *string `mandatory:"false" json:"azureTenantId"`

	// Azure client ID of the application. This property is required when 'authenticationType' is set to 'AZURE_ACTIVE_DIRECTORY'.
	// e.g.: 06ecaabf-8b80-4ec8-a0ec-20cbf463703d
	ClientId *string `mandatory:"false" json:"clientId"`

	// Azure client secret (aka application password) for authentication. This property is required when 'authenticationType' is set to 'AZURE_ACTIVE_DIRECTORY'.
	// e.g.: dO29Q~F5-VwnA.lZdd11xFF_t5NAXCaGwDl9NbT1
	// Deprecated: This field is deprecated and replaced by "clientSecretSecretId". This field will be removed after February 15 2026.
	ClientSecret *string `mandatory:"false" json:"clientSecret"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the client secret is stored.
	// Note: When provided, 'clientSecret' field must not be provided.
	ClientSecretSecretId *string `mandatory:"false" json:"clientSecretSecretId"`

	// Azure Storage service endpoint.
	// e.g: https://test.blob.core.windows.net
	Endpoint *string `mandatory:"false" json:"endpoint"`

	// Controls the network traffic direction to the target:
	// SHARED_SERVICE_ENDPOINT: Traffic flows through the Goldengate Service's network to public hosts. Cannot be used for private targets.
	// SHARED_DEPLOYMENT_ENDPOINT: Network traffic flows from the assigned deployment's private endpoint through the deployment's subnet.
	// DEDICATED_ENDPOINT: A dedicated private endpoint is created in the target VCN subnet for the connection. The subnetId is required when DEDICATED_ENDPOINT networking is selected.
	RoutingMethod RoutingMethodEnum `mandatory:"false" json:"routingMethod,omitempty"`

	// The Azure Data Lake Storage technology type.
	TechnologyType AzureDataLakeStorageConnectionTechnologyTypeEnum `mandatory:"true" json:"technologyType"`

	// Used authentication mechanism to access Azure Data Lake Storage.
	AuthenticationType AzureDataLakeStorageConnectionAuthenticationTypeEnum `mandatory:"true" json:"authenticationType"`
}

// GetDisplayName returns DisplayName
func (m CreateAzureDataLakeStorageConnectionDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m CreateAzureDataLakeStorageConnectionDetails) GetDescription() *string {
	return m.Description
}

// GetCompartmentId returns CompartmentId
func (m CreateAzureDataLakeStorageConnectionDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetFreeformTags returns FreeformTags
func (m CreateAzureDataLakeStorageConnectionDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateAzureDataLakeStorageConnectionDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetLocks returns Locks
func (m CreateAzureDataLakeStorageConnectionDetails) GetLocks() []AddResourceLockDetails {
	return m.Locks
}

// GetVaultId returns VaultId
func (m CreateAzureDataLakeStorageConnectionDetails) GetVaultId() *string {
	return m.VaultId
}

// GetKeyId returns KeyId
func (m CreateAzureDataLakeStorageConnectionDetails) GetKeyId() *string {
	return m.KeyId
}

// GetNsgIds returns NsgIds
func (m CreateAzureDataLakeStorageConnectionDetails) GetNsgIds() []string {
	return m.NsgIds
}

// GetSubnetId returns SubnetId
func (m CreateAzureDataLakeStorageConnectionDetails) GetSubnetId() *string {
	return m.SubnetId
}

// GetRoutingMethod returns RoutingMethod
func (m CreateAzureDataLakeStorageConnectionDetails) GetRoutingMethod() RoutingMethodEnum {
	return m.RoutingMethod
}

// GetDoesUseSecretIds returns DoesUseSecretIds
func (m CreateAzureDataLakeStorageConnectionDetails) GetDoesUseSecretIds() *bool {
	return m.DoesUseSecretIds
}

func (m CreateAzureDataLakeStorageConnectionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateAzureDataLakeStorageConnectionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingRoutingMethodEnum(string(m.RoutingMethod)); !ok && m.RoutingMethod != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RoutingMethod: %s. Supported values are: %s.", m.RoutingMethod, strings.Join(GetRoutingMethodEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAzureDataLakeStorageConnectionTechnologyTypeEnum(string(m.TechnologyType)); !ok && m.TechnologyType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TechnologyType: %s. Supported values are: %s.", m.TechnologyType, strings.Join(GetAzureDataLakeStorageConnectionTechnologyTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAzureDataLakeStorageConnectionAuthenticationTypeEnum(string(m.AuthenticationType)); !ok && m.AuthenticationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AuthenticationType: %s. Supported values are: %s.", m.AuthenticationType, strings.Join(GetAzureDataLakeStorageConnectionAuthenticationTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateAzureDataLakeStorageConnectionDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateAzureDataLakeStorageConnectionDetails CreateAzureDataLakeStorageConnectionDetails
	s := struct {
		DiscriminatorParam string `json:"connectionType"`
		MarshalTypeCreateAzureDataLakeStorageConnectionDetails
	}{
		"AZURE_DATA_LAKE_STORAGE",
		(MarshalTypeCreateAzureDataLakeStorageConnectionDetails)(m),
	}

	return json.Marshal(&s)
}
