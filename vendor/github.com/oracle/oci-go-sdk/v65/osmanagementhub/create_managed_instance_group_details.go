// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for instances in OCI, your private data center, or 3rd-party clouds.
// For more information, see Overview of OS Management Hub (https://docs.cloud.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateManagedInstanceGroupDetails Provides the information used to create a new managed instance group.
type CreateManagedInstanceGroupDetails struct {

	// A user-friendly name for the managed instance group. Does not have to be unique and you can change the name later. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the managed instance group.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The operating system type of the managed instances that will be attached to this group.
	OsFamily OsFamilyEnum `mandatory:"true" json:"osFamily"`

	// The CPU architecture type of the managed instances that will be attached to this group.
	ArchType ArchTypeEnum `mandatory:"true" json:"archType"`

	// The vendor of the operating system that will be used by the managed instances in the group.
	VendorName VendorNameEnum `mandatory:"true" json:"vendorName"`

	// User-specified description of the managed instance group. Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// The location of managed instances attached to the group. If no location is provided, the default is on premises.
	Location ManagedInstanceLocationEnum `mandatory:"false" json:"location,omitempty"`

	// The list of software source OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) available to the managed instances in the group.
	SoftwareSourceIds []string `mandatory:"false" json:"softwareSourceIds"`

	// The list of managed instance OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) to be added to the group.
	ManagedInstanceIds []string `mandatory:"false" json:"managedInstanceIds"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the Oracle Notifications service (ONS) topic. ONS is the channel used to send notifications to the customer.
	NotificationTopicId *string `mandatory:"false" json:"notificationTopicId"`

	AutonomousSettings *UpdatableAutonomousSettings `mandatory:"false" json:"autonomousSettings"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateManagedInstanceGroupDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateManagedInstanceGroupDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
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
