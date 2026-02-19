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

// GiHomeSummary A Grid Infrastructure Home summary.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an
// administrator. If you're an administrator who needs to write policies to give users access,
// see Getting Started with Policies (https://docs.oracle.com/iaas/Content/Identity/Concepts/policygetstarted.htm).
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type GiHomeSummary struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Grid Infrastructure Home.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The user-provided name for the Grid Infrastructure Home. The name does not need to be unique and gets dynamically generated if null.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The current state of the Grid Infrastructure Home.
	LifecycleState GiHomeSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

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
	AvailableActions []GiHomeSummaryAvailableActionsEnum `mandatory:"false" json:"availableActions,omitempty"`

	// The type of update that the Grid Infrastructure Home can be used for.
	AvailableType GiHomeSummaryAvailableTypeEnum `mandatory:"false" json:"availableType,omitempty"`

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

func (m GiHomeSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GiHomeSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGiHomeSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetGiHomeSummaryLifecycleStateEnumStringValues(), ",")))
	}

	for _, val := range m.AvailableActions {
		if _, ok := GetMappingGiHomeSummaryAvailableActionsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AvailableActions: %s. Supported values are: %s.", val, strings.Join(GetGiHomeSummaryAvailableActionsEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingGiHomeSummaryAvailableTypeEnum(string(m.AvailableType)); !ok && m.AvailableType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AvailableType: %s. Supported values are: %s.", m.AvailableType, strings.Join(GetGiHomeSummaryAvailableTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GiHomeSummaryLifecycleStateEnum Enum with underlying type: string
type GiHomeSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for GiHomeSummaryLifecycleStateEnum
const (
	GiHomeSummaryLifecycleStateProvisioning GiHomeSummaryLifecycleStateEnum = "PROVISIONING"
	GiHomeSummaryLifecycleStateAvailable    GiHomeSummaryLifecycleStateEnum = "AVAILABLE"
	GiHomeSummaryLifecycleStateUpdating     GiHomeSummaryLifecycleStateEnum = "UPDATING"
	GiHomeSummaryLifecycleStateTerminating  GiHomeSummaryLifecycleStateEnum = "TERMINATING"
	GiHomeSummaryLifecycleStateTerminated   GiHomeSummaryLifecycleStateEnum = "TERMINATED"
	GiHomeSummaryLifecycleStateFailed       GiHomeSummaryLifecycleStateEnum = "FAILED"
	GiHomeSummaryLifecycleStateInactive     GiHomeSummaryLifecycleStateEnum = "INACTIVE"
)

var mappingGiHomeSummaryLifecycleStateEnum = map[string]GiHomeSummaryLifecycleStateEnum{
	"PROVISIONING": GiHomeSummaryLifecycleStateProvisioning,
	"AVAILABLE":    GiHomeSummaryLifecycleStateAvailable,
	"UPDATING":     GiHomeSummaryLifecycleStateUpdating,
	"TERMINATING":  GiHomeSummaryLifecycleStateTerminating,
	"TERMINATED":   GiHomeSummaryLifecycleStateTerminated,
	"FAILED":       GiHomeSummaryLifecycleStateFailed,
	"INACTIVE":     GiHomeSummaryLifecycleStateInactive,
}

var mappingGiHomeSummaryLifecycleStateEnumLowerCase = map[string]GiHomeSummaryLifecycleStateEnum{
	"provisioning": GiHomeSummaryLifecycleStateProvisioning,
	"available":    GiHomeSummaryLifecycleStateAvailable,
	"updating":     GiHomeSummaryLifecycleStateUpdating,
	"terminating":  GiHomeSummaryLifecycleStateTerminating,
	"terminated":   GiHomeSummaryLifecycleStateTerminated,
	"failed":       GiHomeSummaryLifecycleStateFailed,
	"inactive":     GiHomeSummaryLifecycleStateInactive,
}

// GetGiHomeSummaryLifecycleStateEnumValues Enumerates the set of values for GiHomeSummaryLifecycleStateEnum
func GetGiHomeSummaryLifecycleStateEnumValues() []GiHomeSummaryLifecycleStateEnum {
	values := make([]GiHomeSummaryLifecycleStateEnum, 0)
	for _, v := range mappingGiHomeSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetGiHomeSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for GiHomeSummaryLifecycleStateEnum
func GetGiHomeSummaryLifecycleStateEnumStringValues() []string {
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

// GetMappingGiHomeSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGiHomeSummaryLifecycleStateEnum(val string) (GiHomeSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingGiHomeSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// GiHomeSummaryAvailableActionsEnum Enum with underlying type: string
type GiHomeSummaryAvailableActionsEnum string

// Set of constants representing the allowable values for GiHomeSummaryAvailableActionsEnum
const (
	GiHomeSummaryAvailableActionsPrecheck     GiHomeSummaryAvailableActionsEnum = "PRECHECK"
	GiHomeSummaryAvailableActionsRollingApply GiHomeSummaryAvailableActionsEnum = "ROLLING_APPLY"
	GiHomeSummaryAvailableActionsRetry        GiHomeSummaryAvailableActionsEnum = "RETRY"
	GiHomeSummaryAvailableActionsRollback     GiHomeSummaryAvailableActionsEnum = "ROLLBACK"
)

var mappingGiHomeSummaryAvailableActionsEnum = map[string]GiHomeSummaryAvailableActionsEnum{
	"PRECHECK":      GiHomeSummaryAvailableActionsPrecheck,
	"ROLLING_APPLY": GiHomeSummaryAvailableActionsRollingApply,
	"RETRY":         GiHomeSummaryAvailableActionsRetry,
	"ROLLBACK":      GiHomeSummaryAvailableActionsRollback,
}

var mappingGiHomeSummaryAvailableActionsEnumLowerCase = map[string]GiHomeSummaryAvailableActionsEnum{
	"precheck":      GiHomeSummaryAvailableActionsPrecheck,
	"rolling_apply": GiHomeSummaryAvailableActionsRollingApply,
	"retry":         GiHomeSummaryAvailableActionsRetry,
	"rollback":      GiHomeSummaryAvailableActionsRollback,
}

// GetGiHomeSummaryAvailableActionsEnumValues Enumerates the set of values for GiHomeSummaryAvailableActionsEnum
func GetGiHomeSummaryAvailableActionsEnumValues() []GiHomeSummaryAvailableActionsEnum {
	values := make([]GiHomeSummaryAvailableActionsEnum, 0)
	for _, v := range mappingGiHomeSummaryAvailableActionsEnum {
		values = append(values, v)
	}
	return values
}

// GetGiHomeSummaryAvailableActionsEnumStringValues Enumerates the set of values in String for GiHomeSummaryAvailableActionsEnum
func GetGiHomeSummaryAvailableActionsEnumStringValues() []string {
	return []string{
		"PRECHECK",
		"ROLLING_APPLY",
		"RETRY",
		"ROLLBACK",
	}
}

// GetMappingGiHomeSummaryAvailableActionsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGiHomeSummaryAvailableActionsEnum(val string) (GiHomeSummaryAvailableActionsEnum, bool) {
	enum, ok := mappingGiHomeSummaryAvailableActionsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// GiHomeSummaryAvailableTypeEnum Enum with underlying type: string
type GiHomeSummaryAvailableTypeEnum string

// Set of constants representing the allowable values for GiHomeSummaryAvailableTypeEnum
const (
	GiHomeSummaryAvailableTypeUpgrade GiHomeSummaryAvailableTypeEnum = "GI_UPGRADE"
	GiHomeSummaryAvailableTypePatch   GiHomeSummaryAvailableTypeEnum = "GI_PATCH"
)

var mappingGiHomeSummaryAvailableTypeEnum = map[string]GiHomeSummaryAvailableTypeEnum{
	"GI_UPGRADE": GiHomeSummaryAvailableTypeUpgrade,
	"GI_PATCH":   GiHomeSummaryAvailableTypePatch,
}

var mappingGiHomeSummaryAvailableTypeEnumLowerCase = map[string]GiHomeSummaryAvailableTypeEnum{
	"gi_upgrade": GiHomeSummaryAvailableTypeUpgrade,
	"gi_patch":   GiHomeSummaryAvailableTypePatch,
}

// GetGiHomeSummaryAvailableTypeEnumValues Enumerates the set of values for GiHomeSummaryAvailableTypeEnum
func GetGiHomeSummaryAvailableTypeEnumValues() []GiHomeSummaryAvailableTypeEnum {
	values := make([]GiHomeSummaryAvailableTypeEnum, 0)
	for _, v := range mappingGiHomeSummaryAvailableTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetGiHomeSummaryAvailableTypeEnumStringValues Enumerates the set of values in String for GiHomeSummaryAvailableTypeEnum
func GetGiHomeSummaryAvailableTypeEnumStringValues() []string {
	return []string{
		"GI_UPGRADE",
		"GI_PATCH",
	}
}

// GetMappingGiHomeSummaryAvailableTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGiHomeSummaryAvailableTypeEnum(val string) (GiHomeSummaryAvailableTypeEnum, bool) {
	enum, ok := mappingGiHomeSummaryAvailableTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
