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

// CreateManagementCellDetails The data required for creating a ManagementCell.
type CreateManagementCellDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the ManagementCell.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The availability domain the Management Cell is in. May be unset
	// as a blank or NULL value.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The current state of the ManagementCell.
	LifecycleState CreateManagementCellDetailsLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// ManagementCell can be categorized based on the customer filesystems it is hosting.
	// Example: `RESTRICTED` category cell is restricted for use by only one customer
	Category CreateManagementCellDetailsCategoryEnum `mandatory:"true" json:"category"`

	// Provisional cell capacity available for creating new filesystems on the cell. Measured in GB.
	AvailableCapacityInGBs *int64 `mandatory:"true" json:"availableCapacityInGBs"`

	Details *Details `mandatory:"true" json:"details"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateManagementCellDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateManagementCellDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateManagementCellDetailsLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetCreateManagementCellDetailsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCreateManagementCellDetailsCategoryEnum(string(m.Category)); !ok && m.Category != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Category: %s. Supported values are: %s.", m.Category, strings.Join(GetCreateManagementCellDetailsCategoryEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateManagementCellDetailsLifecycleStateEnum Enum with underlying type: string
type CreateManagementCellDetailsLifecycleStateEnum string

// Set of constants representing the allowable values for CreateManagementCellDetailsLifecycleStateEnum
const (
	CreateManagementCellDetailsLifecycleStateCreating CreateManagementCellDetailsLifecycleStateEnum = "CREATING"
	CreateManagementCellDetailsLifecycleStateActive   CreateManagementCellDetailsLifecycleStateEnum = "ACTIVE"
	CreateManagementCellDetailsLifecycleStateInactive CreateManagementCellDetailsLifecycleStateEnum = "INACTIVE"
	CreateManagementCellDetailsLifecycleStateDeleted  CreateManagementCellDetailsLifecycleStateEnum = "DELETED"
)

var mappingCreateManagementCellDetailsLifecycleStateEnum = map[string]CreateManagementCellDetailsLifecycleStateEnum{
	"CREATING": CreateManagementCellDetailsLifecycleStateCreating,
	"ACTIVE":   CreateManagementCellDetailsLifecycleStateActive,
	"INACTIVE": CreateManagementCellDetailsLifecycleStateInactive,
	"DELETED":  CreateManagementCellDetailsLifecycleStateDeleted,
}

var mappingCreateManagementCellDetailsLifecycleStateEnumLowerCase = map[string]CreateManagementCellDetailsLifecycleStateEnum{
	"creating": CreateManagementCellDetailsLifecycleStateCreating,
	"active":   CreateManagementCellDetailsLifecycleStateActive,
	"inactive": CreateManagementCellDetailsLifecycleStateInactive,
	"deleted":  CreateManagementCellDetailsLifecycleStateDeleted,
}

// GetCreateManagementCellDetailsLifecycleStateEnumValues Enumerates the set of values for CreateManagementCellDetailsLifecycleStateEnum
func GetCreateManagementCellDetailsLifecycleStateEnumValues() []CreateManagementCellDetailsLifecycleStateEnum {
	values := make([]CreateManagementCellDetailsLifecycleStateEnum, 0)
	for _, v := range mappingCreateManagementCellDetailsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateManagementCellDetailsLifecycleStateEnumStringValues Enumerates the set of values in String for CreateManagementCellDetailsLifecycleStateEnum
func GetCreateManagementCellDetailsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"DELETED",
	}
}

// GetMappingCreateManagementCellDetailsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateManagementCellDetailsLifecycleStateEnum(val string) (CreateManagementCellDetailsLifecycleStateEnum, bool) {
	enum, ok := mappingCreateManagementCellDetailsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateManagementCellDetailsCategoryEnum Enum with underlying type: string
type CreateManagementCellDetailsCategoryEnum string

// Set of constants representing the allowable values for CreateManagementCellDetailsCategoryEnum
const (
	CreateManagementCellDetailsCategoryGeneral    CreateManagementCellDetailsCategoryEnum = "GENERAL"
	CreateManagementCellDetailsCategoryInternal   CreateManagementCellDetailsCategoryEnum = "INTERNAL"
	CreateManagementCellDetailsCategoryRestricted CreateManagementCellDetailsCategoryEnum = "RESTRICTED"
)

var mappingCreateManagementCellDetailsCategoryEnum = map[string]CreateManagementCellDetailsCategoryEnum{
	"GENERAL":    CreateManagementCellDetailsCategoryGeneral,
	"INTERNAL":   CreateManagementCellDetailsCategoryInternal,
	"RESTRICTED": CreateManagementCellDetailsCategoryRestricted,
}

var mappingCreateManagementCellDetailsCategoryEnumLowerCase = map[string]CreateManagementCellDetailsCategoryEnum{
	"general":    CreateManagementCellDetailsCategoryGeneral,
	"internal":   CreateManagementCellDetailsCategoryInternal,
	"restricted": CreateManagementCellDetailsCategoryRestricted,
}

// GetCreateManagementCellDetailsCategoryEnumValues Enumerates the set of values for CreateManagementCellDetailsCategoryEnum
func GetCreateManagementCellDetailsCategoryEnumValues() []CreateManagementCellDetailsCategoryEnum {
	values := make([]CreateManagementCellDetailsCategoryEnum, 0)
	for _, v := range mappingCreateManagementCellDetailsCategoryEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateManagementCellDetailsCategoryEnumStringValues Enumerates the set of values in String for CreateManagementCellDetailsCategoryEnum
func GetCreateManagementCellDetailsCategoryEnumStringValues() []string {
	return []string{
		"GENERAL",
		"INTERNAL",
		"RESTRICTED",
	}
}

// GetMappingCreateManagementCellDetailsCategoryEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateManagementCellDetailsCategoryEnum(val string) (CreateManagementCellDetailsCategoryEnum, bool) {
	enum, ok := mappingCreateManagementCellDetailsCategoryEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
