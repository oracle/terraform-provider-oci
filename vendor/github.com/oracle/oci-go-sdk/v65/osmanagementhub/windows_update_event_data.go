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

// WindowsUpdateEventData Provides additional information for a windows update event.
type WindowsUpdateEventData struct {

	// Windows update for the current event
	WindowsUpdateStatus WindowsUpdateEventDataWindowsUpdateStatusEnum `mandatory:"true" json:"windowsUpdateStatus"`

	// Status of the software source operation.
	Status EventStatusEnum `mandatory:"false" json:"status,omitempty"`

	// Message from work request
	Message *string `mandatory:"false" json:"message"`

	AdditionalDetails *WorkRequestEventDataAdditionalDetails `mandatory:"false" json:"additionalDetails"`
}

func (m WindowsUpdateEventData) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m WindowsUpdateEventData) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingWindowsUpdateEventDataWindowsUpdateStatusEnum(string(m.WindowsUpdateStatus)); !ok && m.WindowsUpdateStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for WindowsUpdateStatus: %s. Supported values are: %s.", m.WindowsUpdateStatus, strings.Join(GetWindowsUpdateEventDataWindowsUpdateStatusEnumStringValues(), ",")))
	}

	if _, ok := GetMappingEventStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetEventStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// WindowsUpdateEventDataWindowsUpdateStatusEnum Enum with underlying type: string
type WindowsUpdateEventDataWindowsUpdateStatusEnum string

// Set of constants representing the allowable values for WindowsUpdateEventDataWindowsUpdateStatusEnum
const (
	WindowsUpdateEventDataWindowsUpdateStatusListWindowsUpdate                WindowsUpdateEventDataWindowsUpdateStatusEnum = "LIST_WINDOWS_UPDATE"
	WindowsUpdateEventDataWindowsUpdateStatusInstallWindowsUpdates            WindowsUpdateEventDataWindowsUpdateStatusEnum = "INSTALL_WINDOWS_UPDATES"
	WindowsUpdateEventDataWindowsUpdateStatusInstallAllWindowsUpdates         WindowsUpdateEventDataWindowsUpdateStatusEnum = "INSTALL_ALL_WINDOWS_UPDATES"
	WindowsUpdateEventDataWindowsUpdateStatusInstallSecurityWindowsUpdates    WindowsUpdateEventDataWindowsUpdateStatusEnum = "INSTALL_SECURITY_WINDOWS_UPDATES"
	WindowsUpdateEventDataWindowsUpdateStatusInstallBugfixWindowsUpdates      WindowsUpdateEventDataWindowsUpdateStatusEnum = "INSTALL_BUGFIX_WINDOWS_UPDATES"
	WindowsUpdateEventDataWindowsUpdateStatusInstallEnhancementWindowsUpdates WindowsUpdateEventDataWindowsUpdateStatusEnum = "INSTALL_ENHANCEMENT_WINDOWS_UPDATES"
	WindowsUpdateEventDataWindowsUpdateStatusInstallOtherWindowsUpdates       WindowsUpdateEventDataWindowsUpdateStatusEnum = "INSTALL_OTHER_WINDOWS_UPDATES"
)

var mappingWindowsUpdateEventDataWindowsUpdateStatusEnum = map[string]WindowsUpdateEventDataWindowsUpdateStatusEnum{
	"LIST_WINDOWS_UPDATE":                 WindowsUpdateEventDataWindowsUpdateStatusListWindowsUpdate,
	"INSTALL_WINDOWS_UPDATES":             WindowsUpdateEventDataWindowsUpdateStatusInstallWindowsUpdates,
	"INSTALL_ALL_WINDOWS_UPDATES":         WindowsUpdateEventDataWindowsUpdateStatusInstallAllWindowsUpdates,
	"INSTALL_SECURITY_WINDOWS_UPDATES":    WindowsUpdateEventDataWindowsUpdateStatusInstallSecurityWindowsUpdates,
	"INSTALL_BUGFIX_WINDOWS_UPDATES":      WindowsUpdateEventDataWindowsUpdateStatusInstallBugfixWindowsUpdates,
	"INSTALL_ENHANCEMENT_WINDOWS_UPDATES": WindowsUpdateEventDataWindowsUpdateStatusInstallEnhancementWindowsUpdates,
	"INSTALL_OTHER_WINDOWS_UPDATES":       WindowsUpdateEventDataWindowsUpdateStatusInstallOtherWindowsUpdates,
}

var mappingWindowsUpdateEventDataWindowsUpdateStatusEnumLowerCase = map[string]WindowsUpdateEventDataWindowsUpdateStatusEnum{
	"list_windows_update":                 WindowsUpdateEventDataWindowsUpdateStatusListWindowsUpdate,
	"install_windows_updates":             WindowsUpdateEventDataWindowsUpdateStatusInstallWindowsUpdates,
	"install_all_windows_updates":         WindowsUpdateEventDataWindowsUpdateStatusInstallAllWindowsUpdates,
	"install_security_windows_updates":    WindowsUpdateEventDataWindowsUpdateStatusInstallSecurityWindowsUpdates,
	"install_bugfix_windows_updates":      WindowsUpdateEventDataWindowsUpdateStatusInstallBugfixWindowsUpdates,
	"install_enhancement_windows_updates": WindowsUpdateEventDataWindowsUpdateStatusInstallEnhancementWindowsUpdates,
	"install_other_windows_updates":       WindowsUpdateEventDataWindowsUpdateStatusInstallOtherWindowsUpdates,
}

// GetWindowsUpdateEventDataWindowsUpdateStatusEnumValues Enumerates the set of values for WindowsUpdateEventDataWindowsUpdateStatusEnum
func GetWindowsUpdateEventDataWindowsUpdateStatusEnumValues() []WindowsUpdateEventDataWindowsUpdateStatusEnum {
	values := make([]WindowsUpdateEventDataWindowsUpdateStatusEnum, 0)
	for _, v := range mappingWindowsUpdateEventDataWindowsUpdateStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetWindowsUpdateEventDataWindowsUpdateStatusEnumStringValues Enumerates the set of values in String for WindowsUpdateEventDataWindowsUpdateStatusEnum
func GetWindowsUpdateEventDataWindowsUpdateStatusEnumStringValues() []string {
	return []string{
		"LIST_WINDOWS_UPDATE",
		"INSTALL_WINDOWS_UPDATES",
		"INSTALL_ALL_WINDOWS_UPDATES",
		"INSTALL_SECURITY_WINDOWS_UPDATES",
		"INSTALL_BUGFIX_WINDOWS_UPDATES",
		"INSTALL_ENHANCEMENT_WINDOWS_UPDATES",
		"INSTALL_OTHER_WINDOWS_UPDATES",
	}
}

// GetMappingWindowsUpdateEventDataWindowsUpdateStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWindowsUpdateEventDataWindowsUpdateStatusEnum(val string) (WindowsUpdateEventDataWindowsUpdateStatusEnum, bool) {
	enum, ok := mappingWindowsUpdateEventDataWindowsUpdateStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
