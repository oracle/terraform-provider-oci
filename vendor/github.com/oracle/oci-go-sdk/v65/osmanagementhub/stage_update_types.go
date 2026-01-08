// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for instances in OCI, your private data center, or 3rd-party clouds.
// For more information, see Overview of OS Management Hub (https://docs.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"strings"
)

// StageUpdateTypesEnum Enum with underlying type: string
type StageUpdateTypesEnum string

// Set of constants representing the allowable values for StageUpdateTypesEnum
const (
	StageUpdateTypesUpdatePackages                   StageUpdateTypesEnum = "UPDATE_PACKAGES"
	StageUpdateTypesUpdateAll                        StageUpdateTypesEnum = "UPDATE_ALL"
	StageUpdateTypesUpdateSecurity                   StageUpdateTypesEnum = "UPDATE_SECURITY"
	StageUpdateTypesUpdateBugfix                     StageUpdateTypesEnum = "UPDATE_BUGFIX"
	StageUpdateTypesUpdateEnhancement                StageUpdateTypesEnum = "UPDATE_ENHANCEMENT"
	StageUpdateTypesUpdateOther                      StageUpdateTypesEnum = "UPDATE_OTHER"
	StageUpdateTypesUpdateKspliceUserspace           StageUpdateTypesEnum = "UPDATE_KSPLICE_USERSPACE"
	StageUpdateTypesUpdateKspliceKernel              StageUpdateTypesEnum = "UPDATE_KSPLICE_KERNEL"
	StageUpdateTypesInstallWindowsUpdates            StageUpdateTypesEnum = "INSTALL_WINDOWS_UPDATES"
	StageUpdateTypesInstallAllWindowsUpdates         StageUpdateTypesEnum = "INSTALL_ALL_WINDOWS_UPDATES"
	StageUpdateTypesInstallSecurityWindowsUpdates    StageUpdateTypesEnum = "INSTALL_SECURITY_WINDOWS_UPDATES"
	StageUpdateTypesInstallBugfixWindowsUpdates      StageUpdateTypesEnum = "INSTALL_BUGFIX_WINDOWS_UPDATES"
	StageUpdateTypesInstallEnhancementWindowsUpdates StageUpdateTypesEnum = "INSTALL_ENHANCEMENT_WINDOWS_UPDATES"
	StageUpdateTypesInstallOtherWindowsUpdates       StageUpdateTypesEnum = "INSTALL_OTHER_WINDOWS_UPDATES"
)

var mappingStageUpdateTypesEnum = map[string]StageUpdateTypesEnum{
	"UPDATE_PACKAGES":                     StageUpdateTypesUpdatePackages,
	"UPDATE_ALL":                          StageUpdateTypesUpdateAll,
	"UPDATE_SECURITY":                     StageUpdateTypesUpdateSecurity,
	"UPDATE_BUGFIX":                       StageUpdateTypesUpdateBugfix,
	"UPDATE_ENHANCEMENT":                  StageUpdateTypesUpdateEnhancement,
	"UPDATE_OTHER":                        StageUpdateTypesUpdateOther,
	"UPDATE_KSPLICE_USERSPACE":            StageUpdateTypesUpdateKspliceUserspace,
	"UPDATE_KSPLICE_KERNEL":               StageUpdateTypesUpdateKspliceKernel,
	"INSTALL_WINDOWS_UPDATES":             StageUpdateTypesInstallWindowsUpdates,
	"INSTALL_ALL_WINDOWS_UPDATES":         StageUpdateTypesInstallAllWindowsUpdates,
	"INSTALL_SECURITY_WINDOWS_UPDATES":    StageUpdateTypesInstallSecurityWindowsUpdates,
	"INSTALL_BUGFIX_WINDOWS_UPDATES":      StageUpdateTypesInstallBugfixWindowsUpdates,
	"INSTALL_ENHANCEMENT_WINDOWS_UPDATES": StageUpdateTypesInstallEnhancementWindowsUpdates,
	"INSTALL_OTHER_WINDOWS_UPDATES":       StageUpdateTypesInstallOtherWindowsUpdates,
}

var mappingStageUpdateTypesEnumLowerCase = map[string]StageUpdateTypesEnum{
	"update_packages":                     StageUpdateTypesUpdatePackages,
	"update_all":                          StageUpdateTypesUpdateAll,
	"update_security":                     StageUpdateTypesUpdateSecurity,
	"update_bugfix":                       StageUpdateTypesUpdateBugfix,
	"update_enhancement":                  StageUpdateTypesUpdateEnhancement,
	"update_other":                        StageUpdateTypesUpdateOther,
	"update_ksplice_userspace":            StageUpdateTypesUpdateKspliceUserspace,
	"update_ksplice_kernel":               StageUpdateTypesUpdateKspliceKernel,
	"install_windows_updates":             StageUpdateTypesInstallWindowsUpdates,
	"install_all_windows_updates":         StageUpdateTypesInstallAllWindowsUpdates,
	"install_security_windows_updates":    StageUpdateTypesInstallSecurityWindowsUpdates,
	"install_bugfix_windows_updates":      StageUpdateTypesInstallBugfixWindowsUpdates,
	"install_enhancement_windows_updates": StageUpdateTypesInstallEnhancementWindowsUpdates,
	"install_other_windows_updates":       StageUpdateTypesInstallOtherWindowsUpdates,
}

// GetStageUpdateTypesEnumValues Enumerates the set of values for StageUpdateTypesEnum
func GetStageUpdateTypesEnumValues() []StageUpdateTypesEnum {
	values := make([]StageUpdateTypesEnum, 0)
	for _, v := range mappingStageUpdateTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetStageUpdateTypesEnumStringValues Enumerates the set of values in String for StageUpdateTypesEnum
func GetStageUpdateTypesEnumStringValues() []string {
	return []string{
		"UPDATE_PACKAGES",
		"UPDATE_ALL",
		"UPDATE_SECURITY",
		"UPDATE_BUGFIX",
		"UPDATE_ENHANCEMENT",
		"UPDATE_OTHER",
		"UPDATE_KSPLICE_USERSPACE",
		"UPDATE_KSPLICE_KERNEL",
		"INSTALL_WINDOWS_UPDATES",
		"INSTALL_ALL_WINDOWS_UPDATES",
		"INSTALL_SECURITY_WINDOWS_UPDATES",
		"INSTALL_BUGFIX_WINDOWS_UPDATES",
		"INSTALL_ENHANCEMENT_WINDOWS_UPDATES",
		"INSTALL_OTHER_WINDOWS_UPDATES",
	}
}

// GetMappingStageUpdateTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingStageUpdateTypesEnum(val string) (StageUpdateTypesEnum, bool) {
	enum, ok := mappingStageUpdateTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
