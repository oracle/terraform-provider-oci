// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Database MultiCloud Data Plane Integration
//
// <b>Microsoft Azure:</b> <br>
// <b>Oracle Azure Connector Resource:</b>:&nbsp;&nbsp;The Oracle Azure Connector Resource is used to install the Azure Arc Server on an Exadata VM cluster in Oracle Exadata Database Service on Dedicated Infrastructure (ExaDB-D).
//  The supported method to install the Azure Arc Server (Azure Identity) on the Exadata VM cluster:
// <ul>
//  <li>Using a Bearer Access Token</li>
// </ul>
// <b>Oracle Azure Blob Container Resource:</b>&nbsp;&nbsp;The Oracle Azure Blob Container Resource is used to capture the details of an Azure Blob Container.
// This resource can then be reused across multiple Exadata VM clusters in Oracle Exadata Database Service on Dedicated Infrastructure (ExaDB-D) to mount the Azure container.
// <b>Oracle Azure Blob Mount Resource:</b>&nbsp;&nbsp;The Oracle Azure Blob Mount Resource is used to mount an Azure Blob Container on an Exadata VM cluster in Oracle Exadata Database Service on Dedicated Infrastructure (ExaDB-D).
// It relies on both the Oracle Azure Connector and the Oracle Azure Blob Container Resource to perform the mount operation.
// <b>Discover Azure Vaults and Keys Resource:</b>&nbsp;&nbsp;The Discover Oracle Azure Vaults and Azure Keys Resource is used to discover Azure Vaults and the associated encryption keys available in your Azure project.
// <b>Oracle Azure Vault:</b>&nbsp;&nbsp;The Oracle Azure Vault Resource is used to manage Azure Vaults within Oracle Cloud Infrastructure (OCI) for use with services such as Oracle Exadata Database Service on Dedicated Infrastructure.
// <b>Oracle Azure Key:</b>&nbsp;&nbsp;Oracle Azure Key Resource is used to register and manage a Oracle Azure Key Key within Oracle Cloud Infrastructure (OCI) under an associated Azure Vault.
// <br>
// <b>Google Cloud:</b><br>
// <b>Oracle Google Cloud Connector Resource:</b>&nbsp;&nbsp;The Oracle Google Cloud Connector Resource is used to install the Google Cloud Identity Connector on an Exadata VM cluster in Oracle Exadata Database Service on Dedicated Infrastructure (ExaDB-D).
// <b>Discover Google Key Rings and Keys Resource:</b>&nbsp;&nbsp;The Discover Google Key Rings and Keys Resource is used to discover Google Cloud Key Rings and the associated encryption keys available in your Google Cloud project.
// <b>Google Key Rings Resource:</b>&nbsp;&nbsp;The Google Key Rings Resource is used to register and manage Google Cloud Key Rings within Oracle Cloud Infrastructure (OCI) for use with services such as Oracle Exadata Database Service on Dedicated Infrastructure.
// <b>Google Key Resource:</b>&nbsp;&nbsp;The Google Key Resource is used to register and manage a Google Cloud Key within Oracle Cloud Infrastructure (OCI) under an associated Google Key Ring.
// <br>
// <b>AWS</b>:<br>
// <b>Oracle AWS Connector Resource:</b>&nbsp;&nbsp;The Oracle AWS Connector Resource is used to install the AWS Identity Connector on an Exadata VM cluster in Oracle Exadata Database Service on Dedicated Infrastructure (ExaDB-D).
// <b>Google AWS Key Resource:</b>&nbsp;&nbsp;The Oracle AWS Key Resource is used to register and manage a AWS Key within Oracle Cloud Infrastructure (OCI).
//

package dbmulticloud

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OracleDbAzureConnector Oracle DB Azure Connector resource.
type OracleDbAzureConnector struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle DB Azure Connector resource.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains Oracle DB Azure Connector resource.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle DB Cloud VM Cluster resource where this Azure Arc Agent identity to configure.
	DbClusterResourceId *string `mandatory:"true" json:"dbClusterResourceId"`

	// Azure Tenant ID.
	AzureTenantId *string `mandatory:"true" json:"azureTenantId"`

	// Azure Subscription ID.
	AzureSubscriptionId *string `mandatory:"true" json:"azureSubscriptionId"`

	// Azure Resource group name.
	AzureResourceGroup *string `mandatory:"true" json:"azureResourceGroup"`

	// Oracle DB Azure Connector resource name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// List of all VMs where Arc Agent is installed under Cloud VM Cluster.
	ArcAgentNodes []ArcAgentNodes `mandatory:"false" json:"arcAgentNodes"`

	// Azure bearer access token.
	AccessToken *string `mandatory:"false" json:"accessToken"`

	// Private endpoint IP.
	PrivateEndpointIpAddress *string `mandatory:"false" json:"privateEndpointIpAddress"`

	// Private endpoint's DNS alias.
	PrivateEndpointDnsAlias *string `mandatory:"false" json:"privateEndpointDnsAlias"`

	// Azure Identity mechanism.
	AzureIdentityMechanism OracleDbAzureConnectorAzureIdentityMechanismEnum `mandatory:"false" json:"azureIdentityMechanism,omitempty"`

	// The current lifecycle state of the Azure Arc Agent resource.
	LifecycleState OracleDbAzureConnectorLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Description of the current lifecycle state in more detail.
	LifecycleStateDetails *string `mandatory:"false" json:"lifecycleStateDetails"`

	// The current Connectivity status of Azure Identity Connector resource.
	AzureIdentityConnectivityStatus OracleDbAzureConnectorAzureIdentityConnectivityStatusEnum `mandatory:"false" json:"azureIdentityConnectivityStatus,omitempty"`

	// Time when the Oracle DB Azure Connector resource was created expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-22T21:10:29.600Z'
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Time when the Oracle DB Azure Connector resource was last modified expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-22T21:10:29.600Z'
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Description of the latest modification of the Oracle DB Azure Connector resource.
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
	if _, ok := GetMappingOracleDbAzureConnectorAzureIdentityConnectivityStatusEnum(string(m.AzureIdentityConnectivityStatus)); !ok && m.AzureIdentityConnectivityStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AzureIdentityConnectivityStatus: %s. Supported values are: %s.", m.AzureIdentityConnectivityStatus, strings.Join(GetOracleDbAzureConnectorAzureIdentityConnectivityStatusEnumStringValues(), ",")))
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

// OracleDbAzureConnectorAzureIdentityConnectivityStatusEnum Enum with underlying type: string
type OracleDbAzureConnectorAzureIdentityConnectivityStatusEnum string

// Set of constants representing the allowable values for OracleDbAzureConnectorAzureIdentityConnectivityStatusEnum
const (
	OracleDbAzureConnectorAzureIdentityConnectivityStatusConnected          OracleDbAzureConnectorAzureIdentityConnectivityStatusEnum = "CONNECTED"
	OracleDbAzureConnectorAzureIdentityConnectivityStatusDisconnected       OracleDbAzureConnectorAzureIdentityConnectivityStatusEnum = "DISCONNECTED"
	OracleDbAzureConnectorAzureIdentityConnectivityStatusPartiallyConnected OracleDbAzureConnectorAzureIdentityConnectivityStatusEnum = "PARTIALLY_CONNECTED"
	OracleDbAzureConnectorAzureIdentityConnectivityStatusUnknown            OracleDbAzureConnectorAzureIdentityConnectivityStatusEnum = "UNKNOWN"
)

var mappingOracleDbAzureConnectorAzureIdentityConnectivityStatusEnum = map[string]OracleDbAzureConnectorAzureIdentityConnectivityStatusEnum{
	"CONNECTED":           OracleDbAzureConnectorAzureIdentityConnectivityStatusConnected,
	"DISCONNECTED":        OracleDbAzureConnectorAzureIdentityConnectivityStatusDisconnected,
	"PARTIALLY_CONNECTED": OracleDbAzureConnectorAzureIdentityConnectivityStatusPartiallyConnected,
	"UNKNOWN":             OracleDbAzureConnectorAzureIdentityConnectivityStatusUnknown,
}

var mappingOracleDbAzureConnectorAzureIdentityConnectivityStatusEnumLowerCase = map[string]OracleDbAzureConnectorAzureIdentityConnectivityStatusEnum{
	"connected":           OracleDbAzureConnectorAzureIdentityConnectivityStatusConnected,
	"disconnected":        OracleDbAzureConnectorAzureIdentityConnectivityStatusDisconnected,
	"partially_connected": OracleDbAzureConnectorAzureIdentityConnectivityStatusPartiallyConnected,
	"unknown":             OracleDbAzureConnectorAzureIdentityConnectivityStatusUnknown,
}

// GetOracleDbAzureConnectorAzureIdentityConnectivityStatusEnumValues Enumerates the set of values for OracleDbAzureConnectorAzureIdentityConnectivityStatusEnum
func GetOracleDbAzureConnectorAzureIdentityConnectivityStatusEnumValues() []OracleDbAzureConnectorAzureIdentityConnectivityStatusEnum {
	values := make([]OracleDbAzureConnectorAzureIdentityConnectivityStatusEnum, 0)
	for _, v := range mappingOracleDbAzureConnectorAzureIdentityConnectivityStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetOracleDbAzureConnectorAzureIdentityConnectivityStatusEnumStringValues Enumerates the set of values in String for OracleDbAzureConnectorAzureIdentityConnectivityStatusEnum
func GetOracleDbAzureConnectorAzureIdentityConnectivityStatusEnumStringValues() []string {
	return []string{
		"CONNECTED",
		"DISCONNECTED",
		"PARTIALLY_CONNECTED",
		"UNKNOWN",
	}
}

// GetMappingOracleDbAzureConnectorAzureIdentityConnectivityStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOracleDbAzureConnectorAzureIdentityConnectivityStatusEnum(val string) (OracleDbAzureConnectorAzureIdentityConnectivityStatusEnum, bool) {
	enum, ok := mappingOracleDbAzureConnectorAzureIdentityConnectivityStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
