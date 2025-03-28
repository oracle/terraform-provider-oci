// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management API
//
// API for the OS Management service. Use these API operations for working
// with Managed instances and Managed instance groups.
//

package osmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// WindowsUpdate An update available for a Windows managed instance.
type WindowsUpdate struct {

	// Windows Update name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Unique identifier for the Windows update. NOTE - This is not an OCID,
	// but is a unique identifier assigned by Microsoft.
	// Example: `6981d463-cd91-4a26-b7c4-ea4ded9183ed`
	Name *string `mandatory:"true" json:"name"`

	// The purpose of this update.
	UpdateType UpdateTypesEnum `mandatory:"true" json:"updateType"`

	// Information about the Windows Update.
	Description *string `mandatory:"false" json:"description"`

	// size of the package in bytes
	SizeInBytes *int64 `mandatory:"false" json:"sizeInBytes"`

	// Indicates whether the update can be installed using OSMS.
	IsEligibleForInstallation IsEligibleForInstallationEnum `mandatory:"false" json:"isEligibleForInstallation,omitempty"`

	// List of requirements forinstalling on a managed instances
	InstallationRequirements []WindowsUpdateInstallationRequirementsEnum `mandatory:"false" json:"installationRequirements,omitempty"`

	// Indicates whether a reboot may be required to complete installation of this update.
	IsRebootRequiredForInstallation *bool `mandatory:"false" json:"isRebootRequiredForInstallation"`

	// List of the Microsoft Knowledge Base Article Ids related to this Windows Update.
	KbArticleIds []string `mandatory:"false" json:"kbArticleIds"`
}

func (m WindowsUpdate) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m WindowsUpdate) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingUpdateTypesEnum(string(m.UpdateType)); !ok && m.UpdateType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpdateType: %s. Supported values are: %s.", m.UpdateType, strings.Join(GetUpdateTypesEnumStringValues(), ",")))
	}

	if _, ok := GetMappingIsEligibleForInstallationEnum(string(m.IsEligibleForInstallation)); !ok && m.IsEligibleForInstallation != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IsEligibleForInstallation: %s. Supported values are: %s.", m.IsEligibleForInstallation, strings.Join(GetIsEligibleForInstallationEnumStringValues(), ",")))
	}
	for _, val := range m.InstallationRequirements {
		if _, ok := GetMappingWindowsUpdateInstallationRequirementsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InstallationRequirements: %s. Supported values are: %s.", val, strings.Join(GetWindowsUpdateInstallationRequirementsEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// WindowsUpdateInstallationRequirementsEnum Enum with underlying type: string
type WindowsUpdateInstallationRequirementsEnum string

// Set of constants representing the allowable values for WindowsUpdateInstallationRequirementsEnum
const (
	WindowsUpdateInstallationRequirementsEulaAcceptanceRequired  WindowsUpdateInstallationRequirementsEnum = "EULA_ACCEPTANCE_REQUIRED"
	WindowsUpdateInstallationRequirementsSoftwareMediaRequired   WindowsUpdateInstallationRequirementsEnum = "SOFTWARE_MEDIA_REQUIRED"
	WindowsUpdateInstallationRequirementsUserInteractionRequired WindowsUpdateInstallationRequirementsEnum = "USER_INTERACTION_REQUIRED"
)

var mappingWindowsUpdateInstallationRequirementsEnum = map[string]WindowsUpdateInstallationRequirementsEnum{
	"EULA_ACCEPTANCE_REQUIRED":  WindowsUpdateInstallationRequirementsEulaAcceptanceRequired,
	"SOFTWARE_MEDIA_REQUIRED":   WindowsUpdateInstallationRequirementsSoftwareMediaRequired,
	"USER_INTERACTION_REQUIRED": WindowsUpdateInstallationRequirementsUserInteractionRequired,
}

var mappingWindowsUpdateInstallationRequirementsEnumLowerCase = map[string]WindowsUpdateInstallationRequirementsEnum{
	"eula_acceptance_required":  WindowsUpdateInstallationRequirementsEulaAcceptanceRequired,
	"software_media_required":   WindowsUpdateInstallationRequirementsSoftwareMediaRequired,
	"user_interaction_required": WindowsUpdateInstallationRequirementsUserInteractionRequired,
}

// GetWindowsUpdateInstallationRequirementsEnumValues Enumerates the set of values for WindowsUpdateInstallationRequirementsEnum
func GetWindowsUpdateInstallationRequirementsEnumValues() []WindowsUpdateInstallationRequirementsEnum {
	values := make([]WindowsUpdateInstallationRequirementsEnum, 0)
	for _, v := range mappingWindowsUpdateInstallationRequirementsEnum {
		values = append(values, v)
	}
	return values
}

// GetWindowsUpdateInstallationRequirementsEnumStringValues Enumerates the set of values in String for WindowsUpdateInstallationRequirementsEnum
func GetWindowsUpdateInstallationRequirementsEnumStringValues() []string {
	return []string{
		"EULA_ACCEPTANCE_REQUIRED",
		"SOFTWARE_MEDIA_REQUIRED",
		"USER_INTERACTION_REQUIRED",
	}
}

// GetMappingWindowsUpdateInstallationRequirementsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWindowsUpdateInstallationRequirementsEnum(val string) (WindowsUpdateInstallationRequirementsEnum, bool) {
	enum, ok := mappingWindowsUpdateInstallationRequirementsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
