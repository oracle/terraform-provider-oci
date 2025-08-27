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

// ProfileVersion Represents a specific version of a registration profile.
type ProfileVersion struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the registration profile.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the registration profile.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name for the profile.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The vendor of the operating system for the instance.
	VendorName VendorNameEnum `mandatory:"true" json:"vendorName"`

	// The operating system family.
	OsFamily OsFamilyEnum `mandatory:"true" json:"osFamily"`

	// The architecture type.
	ArchType ArchTypeEnum `mandatory:"true" json:"archType"`

	// The description of the registration profile.
	Description *string `mandatory:"false" json:"description"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the management station to associate with an
	// instance once registered. Management stations are only used with non-OCI instances.
	ManagementStationId *string `mandatory:"false" json:"managementStationId"`

	// The list of software sources that the registration profile will use.
	SoftwareSources []SoftwareSourceDetails `mandatory:"false" json:"softwareSources"`

	ManagedInstanceGroup *ManagedInstanceGroupDetails `mandatory:"false" json:"managedInstanceGroup"`

	LifecycleEnvironment *LifecycleEnvironmentDetails `mandatory:"false" json:"lifecycleEnvironment"`

	LifecycleStage *LifecycleStageDetails `mandatory:"false" json:"lifecycleStage"`

	// The type of profile.
	ProfileType ProfileTypeEnum `mandatory:"false" json:"profileType,omitempty"`

	// The time the registration profile was created (in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) format).
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the registration profile was last modified (in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) format).
	TimeModified *common.SDKTime `mandatory:"false" json:"timeModified"`

	// The version of the profile.
	ProfileVersion *string `mandatory:"false" json:"profileVersion"`

	// The current state of the registration profile.
	LifecycleState ProfileLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The type of instance to register.
	RegistrationType ProfileVersionRegistrationTypeEnum `mandatory:"false" json:"registrationType,omitempty"`

	// Indicates if the profile is set as the default. There is exactly one default profile for a specified architecture, OS family, registration type, and vendor. When registering an instance with the corresonding characteristics, the default profile is used, unless another profile is specified.
	IsDefaultProfile *bool `mandatory:"false" json:"isDefaultProfile"`

	// Indicates if the profile was created by the service. OS Management Hub provides a limited set of standardized profiles that can be used to register Autonomous Linux or Windows instances.
	IsServiceProvidedProfile *bool `mandatory:"false" json:"isServiceProvidedProfile"`
}

func (m ProfileVersion) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ProfileVersion) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingVendorNameEnum(string(m.VendorName)); !ok && m.VendorName != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for VendorName: %s. Supported values are: %s.", m.VendorName, strings.Join(GetVendorNameEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOsFamilyEnum(string(m.OsFamily)); !ok && m.OsFamily != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OsFamily: %s. Supported values are: %s.", m.OsFamily, strings.Join(GetOsFamilyEnumStringValues(), ",")))
	}
	if _, ok := GetMappingArchTypeEnum(string(m.ArchType)); !ok && m.ArchType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ArchType: %s. Supported values are: %s.", m.ArchType, strings.Join(GetArchTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingProfileTypeEnum(string(m.ProfileType)); !ok && m.ProfileType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ProfileType: %s. Supported values are: %s.", m.ProfileType, strings.Join(GetProfileTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingProfileLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetProfileLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingProfileVersionRegistrationTypeEnum(string(m.RegistrationType)); !ok && m.RegistrationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RegistrationType: %s. Supported values are: %s.", m.RegistrationType, strings.Join(GetProfileVersionRegistrationTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ProfileVersionRegistrationTypeEnum Enum with underlying type: string
type ProfileVersionRegistrationTypeEnum string

// Set of constants representing the allowable values for ProfileVersionRegistrationTypeEnum
const (
	ProfileVersionRegistrationTypeOciLinux        ProfileVersionRegistrationTypeEnum = "OCI_LINUX"
	ProfileVersionRegistrationTypeNonOciLinux     ProfileVersionRegistrationTypeEnum = "NON_OCI_LINUX"
	ProfileVersionRegistrationTypeOciWindows      ProfileVersionRegistrationTypeEnum = "OCI_WINDOWS"
	ProfileVersionRegistrationTypeAutonomousLinux ProfileVersionRegistrationTypeEnum = "AUTONOMOUS_LINUX"
)

var mappingProfileVersionRegistrationTypeEnum = map[string]ProfileVersionRegistrationTypeEnum{
	"OCI_LINUX":        ProfileVersionRegistrationTypeOciLinux,
	"NON_OCI_LINUX":    ProfileVersionRegistrationTypeNonOciLinux,
	"OCI_WINDOWS":      ProfileVersionRegistrationTypeOciWindows,
	"AUTONOMOUS_LINUX": ProfileVersionRegistrationTypeAutonomousLinux,
}

var mappingProfileVersionRegistrationTypeEnumLowerCase = map[string]ProfileVersionRegistrationTypeEnum{
	"oci_linux":        ProfileVersionRegistrationTypeOciLinux,
	"non_oci_linux":    ProfileVersionRegistrationTypeNonOciLinux,
	"oci_windows":      ProfileVersionRegistrationTypeOciWindows,
	"autonomous_linux": ProfileVersionRegistrationTypeAutonomousLinux,
}

// GetProfileVersionRegistrationTypeEnumValues Enumerates the set of values for ProfileVersionRegistrationTypeEnum
func GetProfileVersionRegistrationTypeEnumValues() []ProfileVersionRegistrationTypeEnum {
	values := make([]ProfileVersionRegistrationTypeEnum, 0)
	for _, v := range mappingProfileVersionRegistrationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetProfileVersionRegistrationTypeEnumStringValues Enumerates the set of values in String for ProfileVersionRegistrationTypeEnum
func GetProfileVersionRegistrationTypeEnumStringValues() []string {
	return []string{
		"OCI_LINUX",
		"NON_OCI_LINUX",
		"OCI_WINDOWS",
		"AUTONOMOUS_LINUX",
	}
}

// GetMappingProfileVersionRegistrationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingProfileVersionRegistrationTypeEnum(val string) (ProfileVersionRegistrationTypeEnum, bool) {
	enum, ok := mappingProfileVersionRegistrationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
