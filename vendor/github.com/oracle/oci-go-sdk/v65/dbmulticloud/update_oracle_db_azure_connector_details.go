// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Database MultiCloud Data plane Integration
//
// 1. Oracle Azure Connector Resource: This is for installing Azure Arc Server in ExaCS VM Cluster.
//   There are two way to install Azure Arc Server (Azure Identity) in ExaCS VMCluster.
//     a. Using Bearer Access Token or
//     b. By providing Authentication token
// 2. Oracle Azure Blob Container Resource: This is for to capture Azure Container details
//    and same will be used in multiple ExaCS VMCluster to mount the Azure Container.
// 3. Oracle Azure Blob Mount Resource: This is for to mount Azure Container in ExaCS VMCluster
//    using Oracle Azure Connector and Oracle Azure Blob Container Resource.
//

package dbmulticloud

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateOracleDbAzureConnectorDetails This object is about to provide input params to update Oracle DB Azure Connector Resource.
type UpdateOracleDbAzureConnectorDetails struct {

	// The ID of the compartment that contains Oracle DB Azure Connector Resource.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Oracle DB Azure Connector Resource name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The ID of the Oracle DB Cluster Resource where this Azure Arc Agent Identity to configure.
	DbClusterResourceId *string `mandatory:"false" json:"dbClusterResourceId"`

	// Azure Identity Mechanism.
	AzureIdentityMechanism OracleDbAzureConnectorAzureIdentityMechanismEnum `mandatory:"false" json:"azureIdentityMechanism,omitempty"`

	// Azure Tenant ID.
	AzureTenantId *string `mandatory:"false" json:"azureTenantId"`

	// Azure Subscription ID.
	AzureSubscriptionId *string `mandatory:"false" json:"azureSubscriptionId"`

	// Azure Resource Group Name.
	AzureResourceGroup *string `mandatory:"false" json:"azureResourceGroup"`

	// Azure bearer access token. If bearer access token is provided then Service Principal detail is not required.
	AccessToken *string `mandatory:"false" json:"accessToken"`

	// Private endpoint IP.
	PrivateEndpointIpAddress *string `mandatory:"false" json:"privateEndpointIpAddress"`

	// Private endpoint DNS Alias.
	PrivateEndpointDnsAlias *string `mandatory:"false" json:"privateEndpointDnsAlias"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateOracleDbAzureConnectorDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateOracleDbAzureConnectorDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOracleDbAzureConnectorAzureIdentityMechanismEnum(string(m.AzureIdentityMechanism)); !ok && m.AzureIdentityMechanism != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AzureIdentityMechanism: %s. Supported values are: %s.", m.AzureIdentityMechanism, strings.Join(GetOracleDbAzureConnectorAzureIdentityMechanismEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
