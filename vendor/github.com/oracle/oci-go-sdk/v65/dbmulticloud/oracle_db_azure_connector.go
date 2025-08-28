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

// OracleDbAzureConnector Oracle DB Azure Connector Details, this resource is for to create Azure Identity on Database Resource.
type OracleDbAzureConnector struct {

	// The ID of the Oracle DB Azure Connector resource.
	Id *string `mandatory:"true" json:"id"`

	// The ID of the compartment that contains Oracle DB Azure Connector resource.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The ID of the DB Cluster Resource where this Azure Arc Agent identity to configure.
	DbClusterResourceId *string `mandatory:"true" json:"dbClusterResourceId"`

	// Azure Tenant ID.
	AzureTenantId *string `mandatory:"true" json:"azureTenantId"`

	// Azure Subscription ID.
	AzureSubscriptionId *string `mandatory:"true" json:"azureSubscriptionId"`

	// Azure Resource Group Name.
	AzureResourceGroup *string `mandatory:"true" json:"azureResourceGroup"`

	// Oracle DB Azure Connector resource name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// List of All VMs where Arc Agent is Install under VMCluster.
	ArcAgentNodes []ArcAgentNodes `mandatory:"false" json:"arcAgentNodes"`

	// Azure bearer access token. If bearer access token is provided then Service Principal detail is not required.
	AccessToken *string `mandatory:"false" json:"accessToken"`

	// Private endpoint IP.
	PrivateEndpointIpAddress *string `mandatory:"false" json:"privateEndpointIpAddress"`

	// Private endpoint DNS Alias.
	PrivateEndpointDnsAlias *string `mandatory:"false" json:"privateEndpointDnsAlias"`

	// Azure Identity Mechanism.
	AzureIdentityMechanism OracleDbAzureConnectorAzureIdentityMechanismEnum `mandatory:"false" json:"azureIdentityMechanism,omitempty"`

	// The current lifecycle state of the Azure Arc Agent Resource.
	LifecycleState OracleDbAzureConnectorLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Description of the current lifecycle state in more detail.
	LifecycleStateDetails *string `mandatory:"false" json:"lifecycleStateDetails"`

	// Time when the Oracle DB Azure Connector Resource was created expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-22T21:10:29.600Z'
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Time when the Oracle DB Azure Connector Resource was last modified expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-22T21:10:29.600Z'
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Description of the latest modification of the Oracle DB Azure Connector Resource.
	LastModification *string `mandatory:"false" json:"lastModification"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m OracleDbAzureConnector) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OracleDbAzureConnector) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOracleDbAzureConnectorAzureIdentityMechanismEnum(string(m.AzureIdentityMechanism)); !ok && m.AzureIdentityMechanism != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AzureIdentityMechanism: %s. Supported values are: %s.", m.AzureIdentityMechanism, strings.Join(GetOracleDbAzureConnectorAzureIdentityMechanismEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOracleDbAzureConnectorLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetOracleDbAzureConnectorLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OracleDbAzureConnectorAzureIdentityMechanismEnum Enum with underlying type: string
type OracleDbAzureConnectorAzureIdentityMechanismEnum string

// Set of constants representing the allowable values for OracleDbAzureConnectorAzureIdentityMechanismEnum
const (
	OracleDbAzureConnectorAzureIdentityMechanismArcAgent         OracleDbAzureConnectorAzureIdentityMechanismEnum = "ARC_AGENT"
	OracleDbAzureConnectorAzureIdentityMechanismServicePrincipal OracleDbAzureConnectorAzureIdentityMechanismEnum = "SERVICE_PRINCIPAL"
)

var mappingOracleDbAzureConnectorAzureIdentityMechanismEnum = map[string]OracleDbAzureConnectorAzureIdentityMechanismEnum{
	"ARC_AGENT":         OracleDbAzureConnectorAzureIdentityMechanismArcAgent,
	"SERVICE_PRINCIPAL": OracleDbAzureConnectorAzureIdentityMechanismServicePrincipal,
}

var mappingOracleDbAzureConnectorAzureIdentityMechanismEnumLowerCase = map[string]OracleDbAzureConnectorAzureIdentityMechanismEnum{
	"arc_agent":         OracleDbAzureConnectorAzureIdentityMechanismArcAgent,
	"service_principal": OracleDbAzureConnectorAzureIdentityMechanismServicePrincipal,
}

// GetOracleDbAzureConnectorAzureIdentityMechanismEnumValues Enumerates the set of values for OracleDbAzureConnectorAzureIdentityMechanismEnum
func GetOracleDbAzureConnectorAzureIdentityMechanismEnumValues() []OracleDbAzureConnectorAzureIdentityMechanismEnum {
	values := make([]OracleDbAzureConnectorAzureIdentityMechanismEnum, 0)
	for _, v := range mappingOracleDbAzureConnectorAzureIdentityMechanismEnum {
		values = append(values, v)
	}
	return values
}

// GetOracleDbAzureConnectorAzureIdentityMechanismEnumStringValues Enumerates the set of values in String for OracleDbAzureConnectorAzureIdentityMechanismEnum
func GetOracleDbAzureConnectorAzureIdentityMechanismEnumStringValues() []string {
	return []string{
		"ARC_AGENT",
		"SERVICE_PRINCIPAL",
	}
}

// GetMappingOracleDbAzureConnectorAzureIdentityMechanismEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOracleDbAzureConnectorAzureIdentityMechanismEnum(val string) (OracleDbAzureConnectorAzureIdentityMechanismEnum, bool) {
	enum, ok := mappingOracleDbAzureConnectorAzureIdentityMechanismEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// OracleDbAzureConnectorLifecycleStateEnum Enum with underlying type: string
type OracleDbAzureConnectorLifecycleStateEnum string

// Set of constants representing the allowable values for OracleDbAzureConnectorLifecycleStateEnum
const (
	OracleDbAzureConnectorLifecycleStateCreating OracleDbAzureConnectorLifecycleStateEnum = "CREATING"
	OracleDbAzureConnectorLifecycleStateActive   OracleDbAzureConnectorLifecycleStateEnum = "ACTIVE"
	OracleDbAzureConnectorLifecycleStateUpdating OracleDbAzureConnectorLifecycleStateEnum = "UPDATING"
	OracleDbAzureConnectorLifecycleStateDeleting OracleDbAzureConnectorLifecycleStateEnum = "DELETING"
	OracleDbAzureConnectorLifecycleStateDeleted  OracleDbAzureConnectorLifecycleStateEnum = "DELETED"
	OracleDbAzureConnectorLifecycleStateFailed   OracleDbAzureConnectorLifecycleStateEnum = "FAILED"
)

var mappingOracleDbAzureConnectorLifecycleStateEnum = map[string]OracleDbAzureConnectorLifecycleStateEnum{
	"CREATING": OracleDbAzureConnectorLifecycleStateCreating,
	"ACTIVE":   OracleDbAzureConnectorLifecycleStateActive,
	"UPDATING": OracleDbAzureConnectorLifecycleStateUpdating,
	"DELETING": OracleDbAzureConnectorLifecycleStateDeleting,
	"DELETED":  OracleDbAzureConnectorLifecycleStateDeleted,
	"FAILED":   OracleDbAzureConnectorLifecycleStateFailed,
}

var mappingOracleDbAzureConnectorLifecycleStateEnumLowerCase = map[string]OracleDbAzureConnectorLifecycleStateEnum{
	"creating": OracleDbAzureConnectorLifecycleStateCreating,
	"active":   OracleDbAzureConnectorLifecycleStateActive,
	"updating": OracleDbAzureConnectorLifecycleStateUpdating,
	"deleting": OracleDbAzureConnectorLifecycleStateDeleting,
	"deleted":  OracleDbAzureConnectorLifecycleStateDeleted,
	"failed":   OracleDbAzureConnectorLifecycleStateFailed,
}

// GetOracleDbAzureConnectorLifecycleStateEnumValues Enumerates the set of values for OracleDbAzureConnectorLifecycleStateEnum
func GetOracleDbAzureConnectorLifecycleStateEnumValues() []OracleDbAzureConnectorLifecycleStateEnum {
	values := make([]OracleDbAzureConnectorLifecycleStateEnum, 0)
	for _, v := range mappingOracleDbAzureConnectorLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetOracleDbAzureConnectorLifecycleStateEnumStringValues Enumerates the set of values in String for OracleDbAzureConnectorLifecycleStateEnum
func GetOracleDbAzureConnectorLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingOracleDbAzureConnectorLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOracleDbAzureConnectorLifecycleStateEnum(val string) (OracleDbAzureConnectorLifecycleStateEnum, bool) {
	enum, ok := mappingOracleDbAzureConnectorLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
