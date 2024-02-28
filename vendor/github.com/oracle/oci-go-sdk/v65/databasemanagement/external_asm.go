// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExternalAsm The details of an external ASM.
type ExternalAsm struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the external ASM.
	Id *string `mandatory:"true" json:"id"`

	// The user-friendly name for the external ASM. The name does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The name of the external ASM.
	ComponentName *string `mandatory:"true" json:"componentName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the external DB system that the ASM is a part of.
	ExternalDbSystemId *string `mandatory:"true" json:"externalDbSystemId"`

	// The current lifecycle state of the external ASM.
	LifecycleState ExternalAsmLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the external ASM was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the external ASM was last updated.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the external connector.
	ExternalConnectorId *string `mandatory:"false" json:"externalConnectorId"`

	// The directory in which ASM is installed. This is the same directory in which Oracle Grid Infrastructure is installed.
	GridHome *string `mandatory:"false" json:"gridHome"`

	// Indicates whether the ASM is a cluster ASM or not.
	IsCluster *bool `mandatory:"false" json:"isCluster"`

	// Indicates whether Oracle Flex ASM is enabled or not.
	IsFlexEnabled *bool `mandatory:"false" json:"isFlexEnabled"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The list of databases that are serviced by the ASM.
	ServicedDatabases []ExternalAsmServicedDatabase `mandatory:"false" json:"servicedDatabases"`

	// The additional details of the external ASM defined in `{"key": "value"}` format.
	// Example: `{"bar-key": "value"}`
	AdditionalDetails map[string]string `mandatory:"false" json:"additionalDetails"`

	// The ASM version.
	Version *string `mandatory:"false" json:"version"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m ExternalAsm) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExternalAsm) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExternalAsmLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetExternalAsmLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExternalAsmLifecycleStateEnum Enum with underlying type: string
type ExternalAsmLifecycleStateEnum string

// Set of constants representing the allowable values for ExternalAsmLifecycleStateEnum
const (
	ExternalAsmLifecycleStateCreating     ExternalAsmLifecycleStateEnum = "CREATING"
	ExternalAsmLifecycleStateNotConnected ExternalAsmLifecycleStateEnum = "NOT_CONNECTED"
	ExternalAsmLifecycleStateActive       ExternalAsmLifecycleStateEnum = "ACTIVE"
	ExternalAsmLifecycleStateInactive     ExternalAsmLifecycleStateEnum = "INACTIVE"
	ExternalAsmLifecycleStateUpdating     ExternalAsmLifecycleStateEnum = "UPDATING"
	ExternalAsmLifecycleStateDeleting     ExternalAsmLifecycleStateEnum = "DELETING"
	ExternalAsmLifecycleStateDeleted      ExternalAsmLifecycleStateEnum = "DELETED"
	ExternalAsmLifecycleStateFailed       ExternalAsmLifecycleStateEnum = "FAILED"
)

var mappingExternalAsmLifecycleStateEnum = map[string]ExternalAsmLifecycleStateEnum{
	"CREATING":      ExternalAsmLifecycleStateCreating,
	"NOT_CONNECTED": ExternalAsmLifecycleStateNotConnected,
	"ACTIVE":        ExternalAsmLifecycleStateActive,
	"INACTIVE":      ExternalAsmLifecycleStateInactive,
	"UPDATING":      ExternalAsmLifecycleStateUpdating,
	"DELETING":      ExternalAsmLifecycleStateDeleting,
	"DELETED":       ExternalAsmLifecycleStateDeleted,
	"FAILED":        ExternalAsmLifecycleStateFailed,
}

var mappingExternalAsmLifecycleStateEnumLowerCase = map[string]ExternalAsmLifecycleStateEnum{
	"creating":      ExternalAsmLifecycleStateCreating,
	"not_connected": ExternalAsmLifecycleStateNotConnected,
	"active":        ExternalAsmLifecycleStateActive,
	"inactive":      ExternalAsmLifecycleStateInactive,
	"updating":      ExternalAsmLifecycleStateUpdating,
	"deleting":      ExternalAsmLifecycleStateDeleting,
	"deleted":       ExternalAsmLifecycleStateDeleted,
	"failed":        ExternalAsmLifecycleStateFailed,
}

// GetExternalAsmLifecycleStateEnumValues Enumerates the set of values for ExternalAsmLifecycleStateEnum
func GetExternalAsmLifecycleStateEnumValues() []ExternalAsmLifecycleStateEnum {
	values := make([]ExternalAsmLifecycleStateEnum, 0)
	for _, v := range mappingExternalAsmLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetExternalAsmLifecycleStateEnumStringValues Enumerates the set of values in String for ExternalAsmLifecycleStateEnum
func GetExternalAsmLifecycleStateEnumStringValues() []string {
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

// GetMappingExternalAsmLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExternalAsmLifecycleStateEnum(val string) (ExternalAsmLifecycleStateEnum, bool) {
	enum, ok := mappingExternalAsmLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
