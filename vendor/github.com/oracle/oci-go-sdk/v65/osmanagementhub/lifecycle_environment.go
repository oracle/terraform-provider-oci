// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for the operating system environments in your private data centers through a single management console. For more information, see Overview of OS Management Hub (https://docs.cloud.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LifecycleEnvironment Contains versioned software source content and lifecycle stages for a managed instance.
type LifecycleEnvironment struct {

	// The OCID of the resource that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the tenancy containing the lifecycle environment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// User specified list of lifecycle stages to be created for the lifecycle environment.
	Stages []LifecycleStage `mandatory:"true" json:"stages"`

	// The current state of the lifecycle environment.
	LifecycleState LifecycleEnvironmentLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The operating system type of the target instances.
	OsFamily OsFamilyEnum `mandatory:"true" json:"osFamily"`

	// The CPU architecture of the target instances.
	ArchType ArchTypeEnum `mandatory:"true" json:"archType"`

	// The software source vendor name.
	VendorName VendorNameEnum `mandatory:"true" json:"vendorName"`

	// The time the lifecycle environment was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// User specified information about the lifecycle environment.
	Description *string `mandatory:"false" json:"description"`

	// The list of managed instance OCIDs specified in the lifecycle stage.
	ManagedInstanceIds []ManagedInstanceDetails `mandatory:"false" json:"managedInstanceIds"`

	// The time the lifecycle environment was last modified. An RFC3339 formatted datetime string.
	TimeModified *common.SDKTime `mandatory:"false" json:"timeModified"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m LifecycleEnvironment) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LifecycleEnvironment) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLifecycleEnvironmentLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleEnvironmentLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOsFamilyEnum(string(m.OsFamily)); !ok && m.OsFamily != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OsFamily: %s. Supported values are: %s.", m.OsFamily, strings.Join(GetOsFamilyEnumStringValues(), ",")))
	}
	if _, ok := GetMappingArchTypeEnum(string(m.ArchType)); !ok && m.ArchType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ArchType: %s. Supported values are: %s.", m.ArchType, strings.Join(GetArchTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingVendorNameEnum(string(m.VendorName)); !ok && m.VendorName != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for VendorName: %s. Supported values are: %s.", m.VendorName, strings.Join(GetVendorNameEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// LifecycleEnvironmentLifecycleStateEnum Enum with underlying type: string
type LifecycleEnvironmentLifecycleStateEnum string

// Set of constants representing the allowable values for LifecycleEnvironmentLifecycleStateEnum
const (
	LifecycleEnvironmentLifecycleStateCreating LifecycleEnvironmentLifecycleStateEnum = "CREATING"
	LifecycleEnvironmentLifecycleStateUpdating LifecycleEnvironmentLifecycleStateEnum = "UPDATING"
	LifecycleEnvironmentLifecycleStateActive   LifecycleEnvironmentLifecycleStateEnum = "ACTIVE"
	LifecycleEnvironmentLifecycleStateDeleting LifecycleEnvironmentLifecycleStateEnum = "DELETING"
	LifecycleEnvironmentLifecycleStateDeleted  LifecycleEnvironmentLifecycleStateEnum = "DELETED"
	LifecycleEnvironmentLifecycleStateFailed   LifecycleEnvironmentLifecycleStateEnum = "FAILED"
)

var mappingLifecycleEnvironmentLifecycleStateEnum = map[string]LifecycleEnvironmentLifecycleStateEnum{
	"CREATING": LifecycleEnvironmentLifecycleStateCreating,
	"UPDATING": LifecycleEnvironmentLifecycleStateUpdating,
	"ACTIVE":   LifecycleEnvironmentLifecycleStateActive,
	"DELETING": LifecycleEnvironmentLifecycleStateDeleting,
	"DELETED":  LifecycleEnvironmentLifecycleStateDeleted,
	"FAILED":   LifecycleEnvironmentLifecycleStateFailed,
}

var mappingLifecycleEnvironmentLifecycleStateEnumLowerCase = map[string]LifecycleEnvironmentLifecycleStateEnum{
	"creating": LifecycleEnvironmentLifecycleStateCreating,
	"updating": LifecycleEnvironmentLifecycleStateUpdating,
	"active":   LifecycleEnvironmentLifecycleStateActive,
	"deleting": LifecycleEnvironmentLifecycleStateDeleting,
	"deleted":  LifecycleEnvironmentLifecycleStateDeleted,
	"failed":   LifecycleEnvironmentLifecycleStateFailed,
}

// GetLifecycleEnvironmentLifecycleStateEnumValues Enumerates the set of values for LifecycleEnvironmentLifecycleStateEnum
func GetLifecycleEnvironmentLifecycleStateEnumValues() []LifecycleEnvironmentLifecycleStateEnum {
	values := make([]LifecycleEnvironmentLifecycleStateEnum, 0)
	for _, v := range mappingLifecycleEnvironmentLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetLifecycleEnvironmentLifecycleStateEnumStringValues Enumerates the set of values in String for LifecycleEnvironmentLifecycleStateEnum
func GetLifecycleEnvironmentLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingLifecycleEnvironmentLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLifecycleEnvironmentLifecycleStateEnum(val string) (LifecycleEnvironmentLifecycleStateEnum, bool) {
	enum, ok := mappingLifecycleEnvironmentLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
