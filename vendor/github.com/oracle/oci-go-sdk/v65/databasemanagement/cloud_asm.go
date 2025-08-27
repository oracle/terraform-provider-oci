// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CloudAsm The details of a cloud ASM.
type CloudAsm struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud ASM.
	Id *string `mandatory:"true" json:"id"`

	// The user-friendly name for the cloud ASM. The name does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The name of the cloud ASM.
	ComponentName *string `mandatory:"true" json:"componentName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud DB system that the ASM is a part of.
	CloudDbSystemId *string `mandatory:"true" json:"cloudDbSystemId"`

	// The current lifecycle state of the cloud ASM.
	LifecycleState CloudAsmLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the cloud ASM was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the cloud ASM was last updated.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) in DBaas service.
	DbaasId *string `mandatory:"false" json:"dbaasId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud connector.
	CloudConnectorId *string `mandatory:"false" json:"cloudConnectorId"`

	// The directory in which ASM is installed. This is the same directory in which Oracle Grid Infrastructure is installed.
	GridHome *string `mandatory:"false" json:"gridHome"`

	// Indicates whether the ASM is a cluster ASM or not.
	IsCluster *bool `mandatory:"false" json:"isCluster"`

	// Indicates whether Oracle Flex ASM is enabled or not.
	IsFlexEnabled *bool `mandatory:"false" json:"isFlexEnabled"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The list of databases that are serviced by the ASM.
	ServicedDatabases []CloudAsmServicedDatabase `mandatory:"false" json:"servicedDatabases"`

	// The additional details of the cloud ASM defined in `{"key": "value"}` format.
	// Example: `{"bar-key": "value"}`
	AdditionalDetails map[string]string `mandatory:"false" json:"additionalDetails"`

	// The ASM version.
	Version *string `mandatory:"false" json:"version"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// System tags can be viewed by users, but can only be created by the system.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m CloudAsm) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CloudAsm) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCloudAsmLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetCloudAsmLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CloudAsmLifecycleStateEnum Enum with underlying type: string
type CloudAsmLifecycleStateEnum string

// Set of constants representing the allowable values for CloudAsmLifecycleStateEnum
const (
	CloudAsmLifecycleStateCreating     CloudAsmLifecycleStateEnum = "CREATING"
	CloudAsmLifecycleStateNotConnected CloudAsmLifecycleStateEnum = "NOT_CONNECTED"
	CloudAsmLifecycleStateActive       CloudAsmLifecycleStateEnum = "ACTIVE"
	CloudAsmLifecycleStateInactive     CloudAsmLifecycleStateEnum = "INACTIVE"
	CloudAsmLifecycleStateUpdating     CloudAsmLifecycleStateEnum = "UPDATING"
	CloudAsmLifecycleStateDeleting     CloudAsmLifecycleStateEnum = "DELETING"
	CloudAsmLifecycleStateDeleted      CloudAsmLifecycleStateEnum = "DELETED"
	CloudAsmLifecycleStateFailed       CloudAsmLifecycleStateEnum = "FAILED"
)

var mappingCloudAsmLifecycleStateEnum = map[string]CloudAsmLifecycleStateEnum{
	"CREATING":      CloudAsmLifecycleStateCreating,
	"NOT_CONNECTED": CloudAsmLifecycleStateNotConnected,
	"ACTIVE":        CloudAsmLifecycleStateActive,
	"INACTIVE":      CloudAsmLifecycleStateInactive,
	"UPDATING":      CloudAsmLifecycleStateUpdating,
	"DELETING":      CloudAsmLifecycleStateDeleting,
	"DELETED":       CloudAsmLifecycleStateDeleted,
	"FAILED":        CloudAsmLifecycleStateFailed,
}

var mappingCloudAsmLifecycleStateEnumLowerCase = map[string]CloudAsmLifecycleStateEnum{
	"creating":      CloudAsmLifecycleStateCreating,
	"not_connected": CloudAsmLifecycleStateNotConnected,
	"active":        CloudAsmLifecycleStateActive,
	"inactive":      CloudAsmLifecycleStateInactive,
	"updating":      CloudAsmLifecycleStateUpdating,
	"deleting":      CloudAsmLifecycleStateDeleting,
	"deleted":       CloudAsmLifecycleStateDeleted,
	"failed":        CloudAsmLifecycleStateFailed,
}

// GetCloudAsmLifecycleStateEnumValues Enumerates the set of values for CloudAsmLifecycleStateEnum
func GetCloudAsmLifecycleStateEnumValues() []CloudAsmLifecycleStateEnum {
	values := make([]CloudAsmLifecycleStateEnum, 0)
	for _, v := range mappingCloudAsmLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetCloudAsmLifecycleStateEnumStringValues Enumerates the set of values in String for CloudAsmLifecycleStateEnum
func GetCloudAsmLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"NOT_CONNECTED",
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingCloudAsmLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCloudAsmLifecycleStateEnum(val string) (CloudAsmLifecycleStateEnum, bool) {
	enum, ok := mappingCloudAsmLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
