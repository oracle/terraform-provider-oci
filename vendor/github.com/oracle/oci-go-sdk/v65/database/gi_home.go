// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// GiHome Details of the Grid Infrastructure Home.
type GiHome struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Grid Infrastructure Home.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The user-provided name for the Grid Infrastructure Home. The name does not need to be unique and gets dynamically generated if null.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The current state of the Grid Infrastructure Home.
	LifecycleState GiHomeLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The time and date as an RFC3339 formatted string, e.g., 2024-04-11T01:59:07.032Z, when the grid infrastructure home was created
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The path of the Grid Infrastructure Home.
	HomePath *string `mandatory:"true" json:"homePath"`

	// A valid Oracle Grid Infrastructure (GI) software version.
	GiVersion *string `mandatory:"true" json:"giVersion"`

	// The time and date as an RFC3339 formatted string, e.g., 2024-04-11T01:59:07.032Z, when the grid infrastructure home was updated
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Indicates whether the grid infrastructure home is created by default.
	IsDefaultCreated *bool `mandatory:"false" json:"isDefaultCreated"`

	// Indicates whether the Grid Infrastructure Home is currently active for the cluster.
	IsActive *bool `mandatory:"false" json:"isActive"`

	// The possible Grid Infrastructure update actions that can be performed using this Grid Infrastructure Home.
	AvailableActions []GiHomeAvailableActionsEnum `mandatory:"false" json:"availableActions,omitempty"`

	// The type of update that the Grid Infrastructure Home can be used for.
	AvailableType GiHomeAvailableTypeEnum `mandatory:"false" json:"availableType,omitempty"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The grid infrastructure home image OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm)
	GiImageId *string `mandatory:"false" json:"giImageId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DB system.
	DbSystemId *string `mandatory:"false" json:"dbSystemId"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m GiHome) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GiHome) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGiHomeLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetGiHomeLifecycleStateEnumStringValues(), ",")))
	}

	for _, val := range m.AvailableActions {
		if _, ok := GetMappingGiHomeAvailableActionsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AvailableActions: %s. Supported values are: %s.", val, strings.Join(GetGiHomeAvailableActionsEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingGiHomeAvailableTypeEnum(string(m.AvailableType)); !ok && m.AvailableType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AvailableType: %s. Supported values are: %s.", m.AvailableType, strings.Join(GetGiHomeAvailableTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GiHomeLifecycleStateEnum Enum with underlying type: string
type GiHomeLifecycleStateEnum string

// Set of constants representing the allowable values for GiHomeLifecycleStateEnum
const (
	GiHomeLifecycleStateProvisioning GiHomeLifecycleStateEnum = "PROVISIONING"
	GiHomeLifecycleStateAvailable    GiHomeLifecycleStateEnum = "AVAILABLE"
	GiHomeLifecycleStateUpdating     GiHomeLifecycleStateEnum = "UPDATING"
	GiHomeLifecycleStateTerminating  GiHomeLifecycleStateEnum = "TERMINATING"
	GiHomeLifecycleStateTerminated   GiHomeLifecycleStateEnum = "TERMINATED"
	GiHomeLifecycleStateFailed       GiHomeLifecycleStateEnum = "FAILED"
	GiHomeLifecycleStateInactive     GiHomeLifecycleStateEnum = "INACTIVE"
)

var mappingGiHomeLifecycleStateEnum = map[string]GiHomeLifecycleStateEnum{
	"PROVISIONING": GiHomeLifecycleStateProvisioning,
	"AVAILABLE":    GiHomeLifecycleStateAvailable,
	"UPDATING":     GiHomeLifecycleStateUpdating,
	"TERMINATING":  GiHomeLifecycleStateTerminating,
	"TERMINATED":   GiHomeLifecycleStateTerminated,
	"FAILED":       GiHomeLifecycleStateFailed,
	"INACTIVE":     GiHomeLifecycleStateInactive,
}

var mappingGiHomeLifecycleStateEnumLowerCase = map[string]GiHomeLifecycleStateEnum{
	"provisioning": GiHomeLifecycleStateProvisioning,
	"available":    GiHomeLifecycleStateAvailable,
	"updating":     GiHomeLifecycleStateUpdating,
	"terminating":  GiHomeLifecycleStateTerminating,
	"terminated":   GiHomeLifecycleStateTerminated,
	"failed":       GiHomeLifecycleStateFailed,
	"inactive":     GiHomeLifecycleStateInactive,
}

// GetGiHomeLifecycleStateEnumValues Enumerates the set of values for GiHomeLifecycleStateEnum
func GetGiHomeLifecycleStateEnumValues() []GiHomeLifecycleStateEnum {
	values := make([]GiHomeLifecycleStateEnum, 0)
	for _, v := range mappingGiHomeLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetGiHomeLifecycleStateEnumStringValues Enumerates the set of values in String for GiHomeLifecycleStateEnum
func GetGiHomeLifecycleStateEnumStringValues() []string {
	return []string{
		"PROVISIONING",
		"AVAILABLE",
		"UPDATING",
		"TERMINATING",
		"TERMINATED",
		"FAILED",
		"INACTIVE",
	}
}

// GetMappingGiHomeLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGiHomeLifecycleStateEnum(val string) (GiHomeLifecycleStateEnum, bool) {
	enum, ok := mappingGiHomeLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// GiHomeAvailableActionsEnum Enum with underlying type: string
type GiHomeAvailableActionsEnum string

// Set of constants representing the allowable values for GiHomeAvailableActionsEnum
const (
	GiHomeAvailableActionsPrecheck     GiHomeAvailableActionsEnum = "PRECHECK"
	GiHomeAvailableActionsRollingApply GiHomeAvailableActionsEnum = "ROLLING_APPLY"
	GiHomeAvailableActionsRetry        GiHomeAvailableActionsEnum = "RETRY"
	GiHomeAvailableActionsRollback     GiHomeAvailableActionsEnum = "ROLLBACK"
)

var mappingGiHomeAvailableActionsEnum = map[string]GiHomeAvailableActionsEnum{
	"PRECHECK":      GiHomeAvailableActionsPrecheck,
	"ROLLING_APPLY": GiHomeAvailableActionsRollingApply,
	"RETRY":         GiHomeAvailableActionsRetry,
	"ROLLBACK":      GiHomeAvailableActionsRollback,
}

var mappingGiHomeAvailableActionsEnumLowerCase = map[string]GiHomeAvailableActionsEnum{
	"precheck":      GiHomeAvailableActionsPrecheck,
	"rolling_apply": GiHomeAvailableActionsRollingApply,
	"retry":         GiHomeAvailableActionsRetry,
	"rollback":      GiHomeAvailableActionsRollback,
}

// GetGiHomeAvailableActionsEnumValues Enumerates the set of values for GiHomeAvailableActionsEnum
func GetGiHomeAvailableActionsEnumValues() []GiHomeAvailableActionsEnum {
	values := make([]GiHomeAvailableActionsEnum, 0)
	for _, v := range mappingGiHomeAvailableActionsEnum {
		values = append(values, v)
	}
	return values
}

// GetGiHomeAvailableActionsEnumStringValues Enumerates the set of values in String for GiHomeAvailableActionsEnum
func GetGiHomeAvailableActionsEnumStringValues() []string {
	return []string{
		"PRECHECK",
		"ROLLING_APPLY",
		"RETRY",
		"ROLLBACK",
	}
}

// GetMappingGiHomeAvailableActionsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGiHomeAvailableActionsEnum(val string) (GiHomeAvailableActionsEnum, bool) {
	enum, ok := mappingGiHomeAvailableActionsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// GiHomeAvailableTypeEnum Enum with underlying type: string
type GiHomeAvailableTypeEnum string

// Set of constants representing the allowable values for GiHomeAvailableTypeEnum
const (
	GiHomeAvailableTypeUpgrade GiHomeAvailableTypeEnum = "GI_UPGRADE"
	GiHomeAvailableTypePatch   GiHomeAvailableTypeEnum = "GI_PATCH"
)

var mappingGiHomeAvailableTypeEnum = map[string]GiHomeAvailableTypeEnum{
	"GI_UPGRADE": GiHomeAvailableTypeUpgrade,
	"GI_PATCH":   GiHomeAvailableTypePatch,
}

var mappingGiHomeAvailableTypeEnumLowerCase = map[string]GiHomeAvailableTypeEnum{
	"gi_upgrade": GiHomeAvailableTypeUpgrade,
	"gi_patch":   GiHomeAvailableTypePatch,
}

// GetGiHomeAvailableTypeEnumValues Enumerates the set of values for GiHomeAvailableTypeEnum
func GetGiHomeAvailableTypeEnumValues() []GiHomeAvailableTypeEnum {
	values := make([]GiHomeAvailableTypeEnum, 0)
	for _, v := range mappingGiHomeAvailableTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetGiHomeAvailableTypeEnumStringValues Enumerates the set of values in String for GiHomeAvailableTypeEnum
func GetGiHomeAvailableTypeEnumStringValues() []string {
	return []string{
		"GI_UPGRADE",
		"GI_PATCH",
	}
}

// GetMappingGiHomeAvailableTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGiHomeAvailableTypeEnum(val string) (GiHomeAvailableTypeEnum, bool) {
	enum, ok := mappingGiHomeAvailableTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
