// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// File Storage with Lustre API
//
// Use the File Storage with Lustre API to manage Lustre file systems and related resources. For more information, see File Storage with Lustre (https://docs.oracle.com/iaas/Content/lustre/home.htm).
//

package lustrefilestorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateManagementCellDetails The data required for updating a ManagementCell.
type UpdateManagementCellDetails struct {

	// The current state of the Management cell.
	LifecycleState UpdateManagementCellDetailsLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// ManagementCell can be categorized based on the customer filesystems it is hosting.
	// Example: `RESTRICTED` category cell is restricted for use by only one customer
	Category UpdateManagementCellDetailsCategoryEnum `mandatory:"false" json:"category,omitempty"`

	// Provisional cell capacity available for creating new filesystems on the cell. Measured in GB.
	AvailableCapacityInGBs *int64 `mandatory:"false" json:"availableCapacityInGBs"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	Details *Details `mandatory:"false" json:"details"`
}

func (m UpdateManagementCellDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateManagementCellDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateManagementCellDetailsLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetUpdateManagementCellDetailsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUpdateManagementCellDetailsCategoryEnum(string(m.Category)); !ok && m.Category != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Category: %s. Supported values are: %s.", m.Category, strings.Join(GetUpdateManagementCellDetailsCategoryEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateManagementCellDetailsLifecycleStateEnum Enum with underlying type: string
type UpdateManagementCellDetailsLifecycleStateEnum string

// Set of constants representing the allowable values for UpdateManagementCellDetailsLifecycleStateEnum
const (
	UpdateManagementCellDetailsLifecycleStateCreating UpdateManagementCellDetailsLifecycleStateEnum = "CREATING"
	UpdateManagementCellDetailsLifecycleStateActive   UpdateManagementCellDetailsLifecycleStateEnum = "ACTIVE"
	UpdateManagementCellDetailsLifecycleStateInactive UpdateManagementCellDetailsLifecycleStateEnum = "INACTIVE"
	UpdateManagementCellDetailsLifecycleStateDeleted  UpdateManagementCellDetailsLifecycleStateEnum = "DELETED"
)

var mappingUpdateManagementCellDetailsLifecycleStateEnum = map[string]UpdateManagementCellDetailsLifecycleStateEnum{
	"CREATING": UpdateManagementCellDetailsLifecycleStateCreating,
	"ACTIVE":   UpdateManagementCellDetailsLifecycleStateActive,
	"INACTIVE": UpdateManagementCellDetailsLifecycleStateInactive,
	"DELETED":  UpdateManagementCellDetailsLifecycleStateDeleted,
}

var mappingUpdateManagementCellDetailsLifecycleStateEnumLowerCase = map[string]UpdateManagementCellDetailsLifecycleStateEnum{
	"creating": UpdateManagementCellDetailsLifecycleStateCreating,
	"active":   UpdateManagementCellDetailsLifecycleStateActive,
	"inactive": UpdateManagementCellDetailsLifecycleStateInactive,
	"deleted":  UpdateManagementCellDetailsLifecycleStateDeleted,
}

// GetUpdateManagementCellDetailsLifecycleStateEnumValues Enumerates the set of values for UpdateManagementCellDetailsLifecycleStateEnum
func GetUpdateManagementCellDetailsLifecycleStateEnumValues() []UpdateManagementCellDetailsLifecycleStateEnum {
	values := make([]UpdateManagementCellDetailsLifecycleStateEnum, 0)
	for _, v := range mappingUpdateManagementCellDetailsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateManagementCellDetailsLifecycleStateEnumStringValues Enumerates the set of values in String for UpdateManagementCellDetailsLifecycleStateEnum
func GetUpdateManagementCellDetailsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"DELETED",
	}
}

// GetMappingUpdateManagementCellDetailsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateManagementCellDetailsLifecycleStateEnum(val string) (UpdateManagementCellDetailsLifecycleStateEnum, bool) {
	enum, ok := mappingUpdateManagementCellDetailsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// UpdateManagementCellDetailsCategoryEnum Enum with underlying type: string
type UpdateManagementCellDetailsCategoryEnum string

// Set of constants representing the allowable values for UpdateManagementCellDetailsCategoryEnum
const (
	UpdateManagementCellDetailsCategoryGeneral    UpdateManagementCellDetailsCategoryEnum = "GENERAL"
	UpdateManagementCellDetailsCategoryInternal   UpdateManagementCellDetailsCategoryEnum = "INTERNAL"
	UpdateManagementCellDetailsCategoryRestricted UpdateManagementCellDetailsCategoryEnum = "RESTRICTED"
)

var mappingUpdateManagementCellDetailsCategoryEnum = map[string]UpdateManagementCellDetailsCategoryEnum{
	"GENERAL":    UpdateManagementCellDetailsCategoryGeneral,
	"INTERNAL":   UpdateManagementCellDetailsCategoryInternal,
	"RESTRICTED": UpdateManagementCellDetailsCategoryRestricted,
}

var mappingUpdateManagementCellDetailsCategoryEnumLowerCase = map[string]UpdateManagementCellDetailsCategoryEnum{
	"general":    UpdateManagementCellDetailsCategoryGeneral,
	"internal":   UpdateManagementCellDetailsCategoryInternal,
	"restricted": UpdateManagementCellDetailsCategoryRestricted,
}

// GetUpdateManagementCellDetailsCategoryEnumValues Enumerates the set of values for UpdateManagementCellDetailsCategoryEnum
func GetUpdateManagementCellDetailsCategoryEnumValues() []UpdateManagementCellDetailsCategoryEnum {
	values := make([]UpdateManagementCellDetailsCategoryEnum, 0)
	for _, v := range mappingUpdateManagementCellDetailsCategoryEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateManagementCellDetailsCategoryEnumStringValues Enumerates the set of values in String for UpdateManagementCellDetailsCategoryEnum
func GetUpdateManagementCellDetailsCategoryEnumStringValues() []string {
	return []string{
		"GENERAL",
		"INTERNAL",
		"RESTRICTED",
	}
}

// GetMappingUpdateManagementCellDetailsCategoryEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateManagementCellDetailsCategoryEnum(val string) (UpdateManagementCellDetailsCategoryEnum, bool) {
	enum, ok := mappingUpdateManagementCellDetailsCategoryEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
