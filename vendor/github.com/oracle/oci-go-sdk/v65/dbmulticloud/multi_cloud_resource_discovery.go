// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Database MultiCloud Data plane Integration
//
// <b>Microsoft Azure</b>:<br>
// 1. Oracle Azure Connector Resource: This is for installing Azure Arc Server in ExaCS VM Cluster.
//   There are two way to install Azure Arc Server (Azure Identity) in ExaCS VMCluster.
//     a. Using Bearer Access Token or
//     b. By providing Authentication token
// 2. Oracle Azure Blob Container Resource: This is for to capture Azure Container details
//    and same will be used in multiple ExaCS VMCluster to mount the Azure Container.
// 3. Oracle Azure Blob Mount Resource: This is for to mount Azure Container in ExaCS VMCluster
//    using Oracle Azure Connector and Oracle Azure Blob Container Resource.
// <b>Google Cloud</b>:<br>
// 1. Oracle Google Cloud Connector Resource: This is for installing Google Identity in ExaCS VM Cluster.<br>
// 2. Discover Google Key-Rings and Keys Resource: This is for to discover Google Key-Rings.<br>
// 3. Google Key-Rings Resource: This is for to maintain Google Key-Rings in Oracle Cloud.<br>
// 4. Google Key Resource: This is for to maintain Google Key in Oracle Cloud for a Google Key-Ring.<br>
//

package dbmulticloud

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MultiCloudResourceDiscovery Multi Cloud Resource Discovery Object.
type MultiCloudResourceDiscovery struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Multi Cloud Discovery Resource.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains Multi Cloud Discovery Resource.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Display name of Multi Cloud Discovery Resource.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle DB Connector Resource.
	OracleDbConnectorId *string `mandatory:"true" json:"oracleDbConnectorId"`

	// Resource Type to discover.
	ResourceType MultiCloudResourceDiscoveryResourceTypeEnum `mandatory:"true" json:"resourceType"`

	// Discover resource using attributes as key-value pair.
	// For GCP supported attributes (keyRing)
	// For Azure supported attributes (keyVault)
	// GCP Example
	// `{"keyRing": "projects/db-mc-dataplane/locations/global/keyRings/dbmci-keyring"}` or
	// `{"keyRing": "dbmci-keyring"}`
	// Azure Example
	// `{"keyVault": "/subscriptions/fd42b73d-5f28-4a23-ae7c-ca08c625fe07/resourceGroups/yumfei0808Test/providers/Microsoft.KeyVault/managedHSMs/orp7HSM001"}` or
	// `{"keyVault": "orp7HSM001"}`
	ResourcesFilter map[string]string `mandatory:"false" json:"resourcesFilter"`

	// List of All Discovered resources.
	Resources []Resources `mandatory:"false" json:"resources"`

	// The current lifecycle state of the discovered resource.
	LifecycleState MultiCloudResourceDiscoveryLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Description of the current lifecycle state in more detail.
	LifecycleStateDetails *string `mandatory:"false" json:"lifecycleStateDetails"`

	// Time when the Multi Cloud Discovery Resource was created in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-22T21:10:29.600Z'
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Time when the Multi Cloud Discovery Resource was last modified, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-22T21:10:29.600Z'
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Description of the latest modification of the Multi Cloud Discovery Resource.
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

func (m MultiCloudResourceDiscovery) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MultiCloudResourceDiscovery) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMultiCloudResourceDiscoveryResourceTypeEnum(string(m.ResourceType)); !ok && m.ResourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResourceType: %s. Supported values are: %s.", m.ResourceType, strings.Join(GetMultiCloudResourceDiscoveryResourceTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingMultiCloudResourceDiscoveryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetMultiCloudResourceDiscoveryLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MultiCloudResourceDiscoveryResourceTypeEnum Enum with underlying type: string
type MultiCloudResourceDiscoveryResourceTypeEnum string

// Set of constants representing the allowable values for MultiCloudResourceDiscoveryResourceTypeEnum
const (
	MultiCloudResourceDiscoveryResourceTypeVaults      MultiCloudResourceDiscoveryResourceTypeEnum = "VAULTS"
	MultiCloudResourceDiscoveryResourceTypeStorage     MultiCloudResourceDiscoveryResourceTypeEnum = "STORAGE"
	MultiCloudResourceDiscoveryResourceTypeGcpKeyRings MultiCloudResourceDiscoveryResourceTypeEnum = "GCP_KEY_RINGS"
)

var mappingMultiCloudResourceDiscoveryResourceTypeEnum = map[string]MultiCloudResourceDiscoveryResourceTypeEnum{
	"VAULTS":        MultiCloudResourceDiscoveryResourceTypeVaults,
	"STORAGE":       MultiCloudResourceDiscoveryResourceTypeStorage,
	"GCP_KEY_RINGS": MultiCloudResourceDiscoveryResourceTypeGcpKeyRings,
}

var mappingMultiCloudResourceDiscoveryResourceTypeEnumLowerCase = map[string]MultiCloudResourceDiscoveryResourceTypeEnum{
	"vaults":        MultiCloudResourceDiscoveryResourceTypeVaults,
	"storage":       MultiCloudResourceDiscoveryResourceTypeStorage,
	"gcp_key_rings": MultiCloudResourceDiscoveryResourceTypeGcpKeyRings,
}

// GetMultiCloudResourceDiscoveryResourceTypeEnumValues Enumerates the set of values for MultiCloudResourceDiscoveryResourceTypeEnum
func GetMultiCloudResourceDiscoveryResourceTypeEnumValues() []MultiCloudResourceDiscoveryResourceTypeEnum {
	values := make([]MultiCloudResourceDiscoveryResourceTypeEnum, 0)
	for _, v := range mappingMultiCloudResourceDiscoveryResourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMultiCloudResourceDiscoveryResourceTypeEnumStringValues Enumerates the set of values in String for MultiCloudResourceDiscoveryResourceTypeEnum
func GetMultiCloudResourceDiscoveryResourceTypeEnumStringValues() []string {
	return []string{
		"VAULTS",
		"STORAGE",
		"GCP_KEY_RINGS",
	}
}

// GetMappingMultiCloudResourceDiscoveryResourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMultiCloudResourceDiscoveryResourceTypeEnum(val string) (MultiCloudResourceDiscoveryResourceTypeEnum, bool) {
	enum, ok := mappingMultiCloudResourceDiscoveryResourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// MultiCloudResourceDiscoveryLifecycleStateEnum Enum with underlying type: string
type MultiCloudResourceDiscoveryLifecycleStateEnum string

// Set of constants representing the allowable values for MultiCloudResourceDiscoveryLifecycleStateEnum
const (
	MultiCloudResourceDiscoveryLifecycleStateAccepted       MultiCloudResourceDiscoveryLifecycleStateEnum = "ACCEPTED"
	MultiCloudResourceDiscoveryLifecycleStateInProgress     MultiCloudResourceDiscoveryLifecycleStateEnum = "IN_PROGRESS"
	MultiCloudResourceDiscoveryLifecycleStateWaiting        MultiCloudResourceDiscoveryLifecycleStateEnum = "WAITING"
	MultiCloudResourceDiscoveryLifecycleStateSucceeded      MultiCloudResourceDiscoveryLifecycleStateEnum = "SUCCEEDED"
	MultiCloudResourceDiscoveryLifecycleStateUpdating       MultiCloudResourceDiscoveryLifecycleStateEnum = "UPDATING"
	MultiCloudResourceDiscoveryLifecycleStateCanceling      MultiCloudResourceDiscoveryLifecycleStateEnum = "CANCELING"
	MultiCloudResourceDiscoveryLifecycleStateCanceled       MultiCloudResourceDiscoveryLifecycleStateEnum = "CANCELED"
	MultiCloudResourceDiscoveryLifecycleStateFailed         MultiCloudResourceDiscoveryLifecycleStateEnum = "FAILED"
	MultiCloudResourceDiscoveryLifecycleStateNeedsAttention MultiCloudResourceDiscoveryLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingMultiCloudResourceDiscoveryLifecycleStateEnum = map[string]MultiCloudResourceDiscoveryLifecycleStateEnum{
	"ACCEPTED":        MultiCloudResourceDiscoveryLifecycleStateAccepted,
	"IN_PROGRESS":     MultiCloudResourceDiscoveryLifecycleStateInProgress,
	"WAITING":         MultiCloudResourceDiscoveryLifecycleStateWaiting,
	"SUCCEEDED":       MultiCloudResourceDiscoveryLifecycleStateSucceeded,
	"UPDATING":        MultiCloudResourceDiscoveryLifecycleStateUpdating,
	"CANCELING":       MultiCloudResourceDiscoveryLifecycleStateCanceling,
	"CANCELED":        MultiCloudResourceDiscoveryLifecycleStateCanceled,
	"FAILED":          MultiCloudResourceDiscoveryLifecycleStateFailed,
	"NEEDS_ATTENTION": MultiCloudResourceDiscoveryLifecycleStateNeedsAttention,
}

var mappingMultiCloudResourceDiscoveryLifecycleStateEnumLowerCase = map[string]MultiCloudResourceDiscoveryLifecycleStateEnum{
	"accepted":        MultiCloudResourceDiscoveryLifecycleStateAccepted,
	"in_progress":     MultiCloudResourceDiscoveryLifecycleStateInProgress,
	"waiting":         MultiCloudResourceDiscoveryLifecycleStateWaiting,
	"succeeded":       MultiCloudResourceDiscoveryLifecycleStateSucceeded,
	"updating":        MultiCloudResourceDiscoveryLifecycleStateUpdating,
	"canceling":       MultiCloudResourceDiscoveryLifecycleStateCanceling,
	"canceled":        MultiCloudResourceDiscoveryLifecycleStateCanceled,
	"failed":          MultiCloudResourceDiscoveryLifecycleStateFailed,
	"needs_attention": MultiCloudResourceDiscoveryLifecycleStateNeedsAttention,
}

// GetMultiCloudResourceDiscoveryLifecycleStateEnumValues Enumerates the set of values for MultiCloudResourceDiscoveryLifecycleStateEnum
func GetMultiCloudResourceDiscoveryLifecycleStateEnumValues() []MultiCloudResourceDiscoveryLifecycleStateEnum {
	values := make([]MultiCloudResourceDiscoveryLifecycleStateEnum, 0)
	for _, v := range mappingMultiCloudResourceDiscoveryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetMultiCloudResourceDiscoveryLifecycleStateEnumStringValues Enumerates the set of values in String for MultiCloudResourceDiscoveryLifecycleStateEnum
func GetMultiCloudResourceDiscoveryLifecycleStateEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"WAITING",
		"SUCCEEDED",
		"UPDATING",
		"CANCELING",
		"CANCELED",
		"FAILED",
		"NEEDS_ATTENTION",
	}
}

// GetMappingMultiCloudResourceDiscoveryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMultiCloudResourceDiscoveryLifecycleStateEnum(val string) (MultiCloudResourceDiscoveryLifecycleStateEnum, bool) {
	enum, ok := mappingMultiCloudResourceDiscoveryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
