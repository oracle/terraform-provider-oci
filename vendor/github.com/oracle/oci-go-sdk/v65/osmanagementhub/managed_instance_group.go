// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for instances in OCI, your private data center, or 3rd-party clouds.
// For more information, see Overview of OS Management Hub (https://docs.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ManagedInstanceGroup An object that defines the managed instance group.
type ManagedInstanceGroup struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance group.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the managed instance group.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The current state of the managed instance group.
	LifecycleState ManagedInstanceGroupLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// A user-friendly name for the managed instance group.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// User-specified information about the managed instance group.
	Description *string `mandatory:"false" json:"description"`

	// The time the managed instance group was created (in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) format).
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the managed instance group was last modified (in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) format).
	TimeModified *common.SDKTime `mandatory:"false" json:"timeModified"`

	// The operating system type of the instances in the managed instance group.
	OsFamily OsFamilyEnum `mandatory:"false" json:"osFamily,omitempty"`

	// The CPU architecture of the instances in the managed instance group.
	ArchType ArchTypeEnum `mandatory:"false" json:"archType,omitempty"`

	// The vendor of the operating system used by the managed instances in the group.
	VendorName VendorNameEnum `mandatory:"false" json:"vendorName,omitempty"`

	// The list of software source OCIDs (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) that the managed instance group will use.
	SoftwareSourceIds []SoftwareSourceDetails `mandatory:"false" json:"softwareSourceIds"`

	// The list of software sources that the managed instance group will use.
	SoftwareSources []SoftwareSourceDetails `mandatory:"false" json:"softwareSources"`

	// The list of managed instance OCIDs (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) attached to the managed instance group.
	ManagedInstanceIds []string `mandatory:"false" json:"managedInstanceIds"`

	// The number of managed instances in the group.
	ManagedInstanceCount *int `mandatory:"false" json:"managedInstanceCount"`

	// The location of managed instances attached to the group.
	Location ManagedInstanceLocationEnum `mandatory:"false" json:"location,omitempty"`

	// The number of scheduled jobs pending against the managed instance group.
	PendingJobCount *int `mandatory:"false" json:"pendingJobCount"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the Oracle Notifications service (ONS) topic. ONS is the channel used to send notifications to the customer.
	NotificationTopicId *string `mandatory:"false" json:"notificationTopicId"`

	AutonomousSettings *AutonomousSettings `mandatory:"false" json:"autonomousSettings"`

	// Indicates whether the Autonomous Linux service manages the group.
	IsManagedByAutonomousLinux *bool `mandatory:"false" json:"isManagedByAutonomousLinux"`

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

func (m ManagedInstanceGroup) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ManagedInstanceGroup) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingManagedInstanceGroupLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetManagedInstanceGroupLifecycleStateEnumStringValues(), ",")))
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
	if _, ok := GetMappingManagedInstanceLocationEnum(string(m.Location)); !ok && m.Location != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Location: %s. Supported values are: %s.", m.Location, strings.Join(GetManagedInstanceLocationEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ManagedInstanceGroupLifecycleStateEnum Enum with underlying type: string
type ManagedInstanceGroupLifecycleStateEnum string

// Set of constants representing the allowable values for ManagedInstanceGroupLifecycleStateEnum
const (
	ManagedInstanceGroupLifecycleStateCreating ManagedInstanceGroupLifecycleStateEnum = "CREATING"
	ManagedInstanceGroupLifecycleStateUpdating ManagedInstanceGroupLifecycleStateEnum = "UPDATING"
	ManagedInstanceGroupLifecycleStateActive   ManagedInstanceGroupLifecycleStateEnum = "ACTIVE"
	ManagedInstanceGroupLifecycleStateDeleting ManagedInstanceGroupLifecycleStateEnum = "DELETING"
	ManagedInstanceGroupLifecycleStateDeleted  ManagedInstanceGroupLifecycleStateEnum = "DELETED"
	ManagedInstanceGroupLifecycleStateFailed   ManagedInstanceGroupLifecycleStateEnum = "FAILED"
)

var mappingManagedInstanceGroupLifecycleStateEnum = map[string]ManagedInstanceGroupLifecycleStateEnum{
	"CREATING": ManagedInstanceGroupLifecycleStateCreating,
	"UPDATING": ManagedInstanceGroupLifecycleStateUpdating,
	"ACTIVE":   ManagedInstanceGroupLifecycleStateActive,
	"DELETING": ManagedInstanceGroupLifecycleStateDeleting,
	"DELETED":  ManagedInstanceGroupLifecycleStateDeleted,
	"FAILED":   ManagedInstanceGroupLifecycleStateFailed,
}

var mappingManagedInstanceGroupLifecycleStateEnumLowerCase = map[string]ManagedInstanceGroupLifecycleStateEnum{
	"creating": ManagedInstanceGroupLifecycleStateCreating,
	"updating": ManagedInstanceGroupLifecycleStateUpdating,
	"active":   ManagedInstanceGroupLifecycleStateActive,
	"deleting": ManagedInstanceGroupLifecycleStateDeleting,
	"deleted":  ManagedInstanceGroupLifecycleStateDeleted,
	"failed":   ManagedInstanceGroupLifecycleStateFailed,
}

// GetManagedInstanceGroupLifecycleStateEnumValues Enumerates the set of values for ManagedInstanceGroupLifecycleStateEnum
func GetManagedInstanceGroupLifecycleStateEnumValues() []ManagedInstanceGroupLifecycleStateEnum {
	values := make([]ManagedInstanceGroupLifecycleStateEnum, 0)
	for _, v := range mappingManagedInstanceGroupLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetManagedInstanceGroupLifecycleStateEnumStringValues Enumerates the set of values in String for ManagedInstanceGroupLifecycleStateEnum
func GetManagedInstanceGroupLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingManagedInstanceGroupLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingManagedInstanceGroupLifecycleStateEnum(val string) (ManagedInstanceGroupLifecycleStateEnum, bool) {
	enum, ok := mappingManagedInstanceGroupLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
