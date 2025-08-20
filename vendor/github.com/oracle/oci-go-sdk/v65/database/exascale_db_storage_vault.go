// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExascaleDbStorageVault Details of the Exadata Database Storage Vault.
type ExascaleDbStorageVault struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata Database Storage Vault.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name of the availability domain in which the Exadata Database Storage Vault is located.
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The current state of the Exadata Database Storage Vault.
	LifecycleState ExascaleDbStorageVaultLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The user-friendly name for the Exadata Database Storage Vault. The name does not need to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	HighCapacityDatabaseStorage *ExascaleDbStorageDetails `mandatory:"true" json:"highCapacityDatabaseStorage"`

	// Exadata Database Storage Vault description.
	Description *string `mandatory:"false" json:"description"`

	// The date and time that the Exadata Database Storage Vault was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The time zone that you want to use for the Exadata Database Storage Vault. For details, see Time Zones (https://docs.oracle.com/iaas/Content/Database/References/timezones.htm).
	TimeZone *string `mandatory:"false" json:"timeZone"`

	// The List of Exadata VM cluster on Exascale Infrastructure OCIDs (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm)
	// **Note:** If Exadata Database Storage Vault is not used for any Exadata VM cluster on Exascale Infrastructure, this list is empty.
	VmClusterIds []string `mandatory:"false" json:"vmClusterIds"`

	// The number of Exadata VM clusters used the Exadata Database Storage Vault.
	VmClusterCount *int `mandatory:"false" json:"vmClusterCount"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata infrastructure.
	ExadataInfrastructureId *string `mandatory:"false" json:"exadataInfrastructureId"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The size of additional Flash Cache in percentage of High Capacity database storage.
	AdditionalFlashCacheInPercent *int `mandatory:"false" json:"additionalFlashCacheInPercent"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cluster placement group of the Exadata Infrastructure or Db System.
	ClusterPlacementGroupId *string `mandatory:"false" json:"clusterPlacementGroupId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
	SubscriptionId *string `mandatory:"false" json:"subscriptionId"`

	// Indicates if autoscale feature is enabled for the Database Storage Vault. The default value is `FALSE`.
	IsAutoscaleEnabled *bool `mandatory:"false" json:"isAutoscaleEnabled"`

	// Maximum limit storage size in gigabytes, that is applicable for the Database Storage Vault.
	AutoscaleLimitInGBs *int `mandatory:"false" json:"autoscaleLimitInGBs"`

	// The shapeAttribute of the Exadata VM cluster(s) associated with the Exadata Database Storage Vault.
	AttachedShapeAttributes []ExascaleDbStorageVaultAttachedShapeAttributesEnum `mandatory:"false" json:"attachedShapeAttributes,omitempty"`
}

func (m ExascaleDbStorageVault) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExascaleDbStorageVault) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExascaleDbStorageVaultLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetExascaleDbStorageVaultLifecycleStateEnumStringValues(), ",")))
	}

	for _, val := range m.AttachedShapeAttributes {
		if _, ok := GetMappingExascaleDbStorageVaultAttachedShapeAttributesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AttachedShapeAttributes: %s. Supported values are: %s.", val, strings.Join(GetExascaleDbStorageVaultAttachedShapeAttributesEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExascaleDbStorageVaultLifecycleStateEnum Enum with underlying type: string
type ExascaleDbStorageVaultLifecycleStateEnum string

// Set of constants representing the allowable values for ExascaleDbStorageVaultLifecycleStateEnum
const (
	ExascaleDbStorageVaultLifecycleStateProvisioning ExascaleDbStorageVaultLifecycleStateEnum = "PROVISIONING"
	ExascaleDbStorageVaultLifecycleStateAvailable    ExascaleDbStorageVaultLifecycleStateEnum = "AVAILABLE"
	ExascaleDbStorageVaultLifecycleStateUpdating     ExascaleDbStorageVaultLifecycleStateEnum = "UPDATING"
	ExascaleDbStorageVaultLifecycleStateTerminating  ExascaleDbStorageVaultLifecycleStateEnum = "TERMINATING"
	ExascaleDbStorageVaultLifecycleStateTerminated   ExascaleDbStorageVaultLifecycleStateEnum = "TERMINATED"
	ExascaleDbStorageVaultLifecycleStateFailed       ExascaleDbStorageVaultLifecycleStateEnum = "FAILED"
)

var mappingExascaleDbStorageVaultLifecycleStateEnum = map[string]ExascaleDbStorageVaultLifecycleStateEnum{
	"PROVISIONING": ExascaleDbStorageVaultLifecycleStateProvisioning,
	"AVAILABLE":    ExascaleDbStorageVaultLifecycleStateAvailable,
	"UPDATING":     ExascaleDbStorageVaultLifecycleStateUpdating,
	"TERMINATING":  ExascaleDbStorageVaultLifecycleStateTerminating,
	"TERMINATED":   ExascaleDbStorageVaultLifecycleStateTerminated,
	"FAILED":       ExascaleDbStorageVaultLifecycleStateFailed,
}

var mappingExascaleDbStorageVaultLifecycleStateEnumLowerCase = map[string]ExascaleDbStorageVaultLifecycleStateEnum{
	"provisioning": ExascaleDbStorageVaultLifecycleStateProvisioning,
	"available":    ExascaleDbStorageVaultLifecycleStateAvailable,
	"updating":     ExascaleDbStorageVaultLifecycleStateUpdating,
	"terminating":  ExascaleDbStorageVaultLifecycleStateTerminating,
	"terminated":   ExascaleDbStorageVaultLifecycleStateTerminated,
	"failed":       ExascaleDbStorageVaultLifecycleStateFailed,
}

// GetExascaleDbStorageVaultLifecycleStateEnumValues Enumerates the set of values for ExascaleDbStorageVaultLifecycleStateEnum
func GetExascaleDbStorageVaultLifecycleStateEnumValues() []ExascaleDbStorageVaultLifecycleStateEnum {
	values := make([]ExascaleDbStorageVaultLifecycleStateEnum, 0)
	for _, v := range mappingExascaleDbStorageVaultLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetExascaleDbStorageVaultLifecycleStateEnumStringValues Enumerates the set of values in String for ExascaleDbStorageVaultLifecycleStateEnum
func GetExascaleDbStorageVaultLifecycleStateEnumStringValues() []string {
	return []string{
		"PROVISIONING",
		"AVAILABLE",
		"UPDATING",
		"TERMINATING",
		"TERMINATED",
		"FAILED",
	}
}

// GetMappingExascaleDbStorageVaultLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExascaleDbStorageVaultLifecycleStateEnum(val string) (ExascaleDbStorageVaultLifecycleStateEnum, bool) {
	enum, ok := mappingExascaleDbStorageVaultLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ExascaleDbStorageVaultAttachedShapeAttributesEnum Enum with underlying type: string
type ExascaleDbStorageVaultAttachedShapeAttributesEnum string

// Set of constants representing the allowable values for ExascaleDbStorageVaultAttachedShapeAttributesEnum
const (
	ExascaleDbStorageVaultAttachedShapeAttributesSmartStorage ExascaleDbStorageVaultAttachedShapeAttributesEnum = "SMART_STORAGE"
	ExascaleDbStorageVaultAttachedShapeAttributesBlockStorage ExascaleDbStorageVaultAttachedShapeAttributesEnum = "BLOCK_STORAGE"
)

var mappingExascaleDbStorageVaultAttachedShapeAttributesEnum = map[string]ExascaleDbStorageVaultAttachedShapeAttributesEnum{
	"SMART_STORAGE": ExascaleDbStorageVaultAttachedShapeAttributesSmartStorage,
	"BLOCK_STORAGE": ExascaleDbStorageVaultAttachedShapeAttributesBlockStorage,
}

var mappingExascaleDbStorageVaultAttachedShapeAttributesEnumLowerCase = map[string]ExascaleDbStorageVaultAttachedShapeAttributesEnum{
	"smart_storage": ExascaleDbStorageVaultAttachedShapeAttributesSmartStorage,
	"block_storage": ExascaleDbStorageVaultAttachedShapeAttributesBlockStorage,
}

// GetExascaleDbStorageVaultAttachedShapeAttributesEnumValues Enumerates the set of values for ExascaleDbStorageVaultAttachedShapeAttributesEnum
func GetExascaleDbStorageVaultAttachedShapeAttributesEnumValues() []ExascaleDbStorageVaultAttachedShapeAttributesEnum {
	values := make([]ExascaleDbStorageVaultAttachedShapeAttributesEnum, 0)
	for _, v := range mappingExascaleDbStorageVaultAttachedShapeAttributesEnum {
		values = append(values, v)
	}
	return values
}

// GetExascaleDbStorageVaultAttachedShapeAttributesEnumStringValues Enumerates the set of values in String for ExascaleDbStorageVaultAttachedShapeAttributesEnum
func GetExascaleDbStorageVaultAttachedShapeAttributesEnumStringValues() []string {
	return []string{
		"SMART_STORAGE",
		"BLOCK_STORAGE",
	}
}

// GetMappingExascaleDbStorageVaultAttachedShapeAttributesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExascaleDbStorageVaultAttachedShapeAttributesEnum(val string) (ExascaleDbStorageVaultAttachedShapeAttributesEnum, bool) {
	enum, ok := mappingExascaleDbStorageVaultAttachedShapeAttributesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
