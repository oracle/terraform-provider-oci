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

// ManagedInstance An object that defines the instance being managed by the service.
type ManagedInstance struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance.
	Id *string `mandatory:"true" json:"id"`

	// User-friendly name for the managed instance.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the tenancy that the managed instance resides in.
	TenancyId *string `mandatory:"true" json:"tenancyId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the managed instance.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Current status of the managed instance.
	Status ManagedInstanceStatusEnum `mandatory:"true" json:"status"`

	// User-specified description for the managed instance.
	Description *string `mandatory:"false" json:"description"`

	// The location of the managed instance.
	Location ManagedInstanceLocationEnum `mandatory:"false" json:"location,omitempty"`

	// Time that the instance last checked in with the service (in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) format).
	TimeLastCheckin *common.SDKTime `mandatory:"false" json:"timeLastCheckin"`

	// Time that the instance last booted (in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) format).
	TimeLastBoot *common.SDKTime `mandatory:"false" json:"timeLastBoot"`

	// Operating system name.
	OsName *string `mandatory:"false" json:"osName"`

	// Operating system version.
	OsVersion *string `mandatory:"false" json:"osVersion"`

	// Operating system kernel version.
	OsKernelVersion *string `mandatory:"false" json:"osKernelVersion"`

	// The ksplice effective kernel version.
	KspliceEffectiveKernelVersion *string `mandatory:"false" json:"kspliceEffectiveKernelVersion"`

	// The CPU architecture type of the managed instance.
	Architecture ArchTypeEnum `mandatory:"false" json:"architecture,omitempty"`

	// The operating system type of the managed instance.
	OsFamily OsFamilyEnum `mandatory:"false" json:"osFamily,omitempty"`

	// The profile that was used to register this instance with the service.
	Profile *string `mandatory:"false" json:"profile"`

	// Indicates whether this managed instance is acting as an on-premises management station.
	IsManagementStation *bool `mandatory:"false" json:"isManagementStation"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the management station for the instance to use as primary management station.
	PrimaryManagementStationId *string `mandatory:"false" json:"primaryManagementStationId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the management station for the instance to use as secondary managment station.
	SecondaryManagementStationId *string `mandatory:"false" json:"secondaryManagementStationId"`

	// The list of software sources currently attached to the managed instance.
	SoftwareSources []SoftwareSourceDetails `mandatory:"false" json:"softwareSources"`

	ManagedInstanceGroup *Id `mandatory:"false" json:"managedInstanceGroup"`

	LifecycleEnvironment *Id `mandatory:"false" json:"lifecycleEnvironment"`

	LifecycleStage *Id `mandatory:"false" json:"lifecycleStage"`

	// Indicates whether a reboot is required to complete installation of updates.
	IsRebootRequired *bool `mandatory:"false" json:"isRebootRequired"`

	// Number of packages installed on the instance.
	InstalledPackages *int `mandatory:"false" json:"installedPackages"`

	// Number of Windows updates installed on the instance.
	InstalledWindowsUpdates *int `mandatory:"false" json:"installedWindowsUpdates"`

	// Number of updates available for installation.
	UpdatesAvailable *int `mandatory:"false" json:"updatesAvailable"`

	// Number of security type updates available for installation.
	SecurityUpdatesAvailable *int `mandatory:"false" json:"securityUpdatesAvailable"`

	// Number of bug fix type updates available for installation.
	BugUpdatesAvailable *int `mandatory:"false" json:"bugUpdatesAvailable"`

	// Number of enhancement type updates available for installation.
	EnhancementUpdatesAvailable *int `mandatory:"false" json:"enhancementUpdatesAvailable"`

	// Number of non-classified (other) updates available for installation.
	OtherUpdatesAvailable *int `mandatory:"false" json:"otherUpdatesAvailable"`

	// Number of scheduled jobs associated with this instance.
	ScheduledJobCount *int `mandatory:"false" json:"scheduledJobCount"`

	// Number of work requests associated with this instance.
	WorkRequestCount *int `mandatory:"false" json:"workRequestCount"`

	// The date and time the instance was created (in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) format).
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the instance was last updated (in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) format).
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the Oracle Notifications service (ONS) topic. ONS is the channel used to send notifications to the customer.
	NotificationTopicId *string `mandatory:"false" json:"notificationTopicId"`

	AutonomousSettings *AutonomousSettings `mandatory:"false" json:"autonomousSettings"`

	// Indicates whether the Autonomous Linux service manages the instance.
	IsManagedByAutonomousLinux *bool `mandatory:"false" json:"isManagedByAutonomousLinux"`
}

func (m ManagedInstance) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ManagedInstance) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingManagedInstanceStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetManagedInstanceStatusEnumStringValues(), ",")))
	}

	if _, ok := GetMappingManagedInstanceLocationEnum(string(m.Location)); !ok && m.Location != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Location: %s. Supported values are: %s.", m.Location, strings.Join(GetManagedInstanceLocationEnumStringValues(), ",")))
	}
	if _, ok := GetMappingArchTypeEnum(string(m.Architecture)); !ok && m.Architecture != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Architecture: %s. Supported values are: %s.", m.Architecture, strings.Join(GetArchTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOsFamilyEnum(string(m.OsFamily)); !ok && m.OsFamily != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OsFamily: %s. Supported values are: %s.", m.OsFamily, strings.Join(GetOsFamilyEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
