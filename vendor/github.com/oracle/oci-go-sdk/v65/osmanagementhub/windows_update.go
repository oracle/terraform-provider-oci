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

// WindowsUpdate An object that provides information about an update for a Windows instance.
type WindowsUpdate struct {

	// Name of the Windows update.
	Name *string `mandatory:"true" json:"name"`

	// Unique identifier for the Windows update. Note that this is not an OCID, but is a unique identifier assigned by Microsoft.
	// Example: '6981d463-cd91-4a26-b7c4-ea4ded9183ed'
	UpdateId *string `mandatory:"true" json:"updateId"`

	// The type of Windows update.
	UpdateType ClassificationTypesEnum `mandatory:"true" json:"updateType"`

	// Description of the update.
	Description *string `mandatory:"false" json:"description"`

	// size of the package in bytes
	SizeInBytes *int64 `mandatory:"false" json:"sizeInBytes"`

	// Indicates whether the update can be installed using the service.
	Installable WindowsUpdateInstallableEnum `mandatory:"false" json:"installable,omitempty"`

	// List of requirements for installing the update on the managed instance.
	InstallationRequirements []InstallationRequirementsEnum `mandatory:"false" json:"installationRequirements"`

	// Indicates whether a reboot is required to complete the installation of this update.
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
	if _, ok := GetMappingClassificationTypesEnum(string(m.UpdateType)); !ok && m.UpdateType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpdateType: %s. Supported values are: %s.", m.UpdateType, strings.Join(GetClassificationTypesEnumStringValues(), ",")))
	}

	if _, ok := GetMappingWindowsUpdateInstallableEnum(string(m.Installable)); !ok && m.Installable != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Installable: %s. Supported values are: %s.", m.Installable, strings.Join(GetWindowsUpdateInstallableEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// WindowsUpdateInstallableEnum Enum with underlying type: string
type WindowsUpdateInstallableEnum string

// Set of constants representing the allowable values for WindowsUpdateInstallableEnum
const (
	WindowsUpdateInstallableInstallable    WindowsUpdateInstallableEnum = "INSTALLABLE"
	WindowsUpdateInstallableNotInstallable WindowsUpdateInstallableEnum = "NOT_INSTALLABLE"
)

var mappingWindowsUpdateInstallableEnum = map[string]WindowsUpdateInstallableEnum{
	"INSTALLABLE":     WindowsUpdateInstallableInstallable,
	"NOT_INSTALLABLE": WindowsUpdateInstallableNotInstallable,
}

var mappingWindowsUpdateInstallableEnumLowerCase = map[string]WindowsUpdateInstallableEnum{
	"installable":     WindowsUpdateInstallableInstallable,
	"not_installable": WindowsUpdateInstallableNotInstallable,
}

// GetWindowsUpdateInstallableEnumValues Enumerates the set of values for WindowsUpdateInstallableEnum
func GetWindowsUpdateInstallableEnumValues() []WindowsUpdateInstallableEnum {
	values := make([]WindowsUpdateInstallableEnum, 0)
	for _, v := range mappingWindowsUpdateInstallableEnum {
		values = append(values, v)
	}
	return values
}

// GetWindowsUpdateInstallableEnumStringValues Enumerates the set of values in String for WindowsUpdateInstallableEnum
func GetWindowsUpdateInstallableEnumStringValues() []string {
	return []string{
		"INSTALLABLE",
		"NOT_INSTALLABLE",
	}
}

// GetMappingWindowsUpdateInstallableEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWindowsUpdateInstallableEnum(val string) (WindowsUpdateInstallableEnum, bool) {
	enum, ok := mappingWindowsUpdateInstallableEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
