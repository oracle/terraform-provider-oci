// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for the operating system environments in your private data centers through a single management console. For more information, see Overview of OS Management Hub (https://docs.cloud.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"strings"
)

// OperationTypesEnum Enum with underlying type: string
type OperationTypesEnum string

// Set of constants representing the allowable values for OperationTypesEnum
const (
	OperationTypesInstallPackages             OperationTypesEnum = "INSTALL_PACKAGES"
	OperationTypesUpdatePackages              OperationTypesEnum = "UPDATE_PACKAGES"
	OperationTypesRemovePackages              OperationTypesEnum = "REMOVE_PACKAGES"
	OperationTypesUpdateAll                   OperationTypesEnum = "UPDATE_ALL"
	OperationTypesUpdateSecurity              OperationTypesEnum = "UPDATE_SECURITY"
	OperationTypesUpdateBugfix                OperationTypesEnum = "UPDATE_BUGFIX"
	OperationTypesUpdateEnhancement           OperationTypesEnum = "UPDATE_ENHANCEMENT"
	OperationTypesUpdateOther                 OperationTypesEnum = "UPDATE_OTHER"
	OperationTypesUpdateKspliceUserspace      OperationTypesEnum = "UPDATE_KSPLICE_USERSPACE"
	OperationTypesUpdateKspliceKernel         OperationTypesEnum = "UPDATE_KSPLICE_KERNEL"
	OperationTypesManageModuleStreams         OperationTypesEnum = "MANAGE_MODULE_STREAMS"
	OperationTypesSwitchModuleStream          OperationTypesEnum = "SWITCH_MODULE_STREAM"
	OperationTypesAttachSoftwareSources       OperationTypesEnum = "ATTACH_SOFTWARE_SOURCES"
	OperationTypesDetachSoftwareSources       OperationTypesEnum = "DETACH_SOFTWARE_SOURCES"
	OperationTypesSyncManagementStationMirror OperationTypesEnum = "SYNC_MANAGEMENT_STATION_MIRROR"
	OperationTypesPromoteLifecycle            OperationTypesEnum = "PROMOTE_LIFECYCLE"
)

var mappingOperationTypesEnum = map[string]OperationTypesEnum{
	"INSTALL_PACKAGES":               OperationTypesInstallPackages,
	"UPDATE_PACKAGES":                OperationTypesUpdatePackages,
	"REMOVE_PACKAGES":                OperationTypesRemovePackages,
	"UPDATE_ALL":                     OperationTypesUpdateAll,
	"UPDATE_SECURITY":                OperationTypesUpdateSecurity,
	"UPDATE_BUGFIX":                  OperationTypesUpdateBugfix,
	"UPDATE_ENHANCEMENT":             OperationTypesUpdateEnhancement,
	"UPDATE_OTHER":                   OperationTypesUpdateOther,
	"UPDATE_KSPLICE_USERSPACE":       OperationTypesUpdateKspliceUserspace,
	"UPDATE_KSPLICE_KERNEL":          OperationTypesUpdateKspliceKernel,
	"MANAGE_MODULE_STREAMS":          OperationTypesManageModuleStreams,
	"SWITCH_MODULE_STREAM":           OperationTypesSwitchModuleStream,
	"ATTACH_SOFTWARE_SOURCES":        OperationTypesAttachSoftwareSources,
	"DETACH_SOFTWARE_SOURCES":        OperationTypesDetachSoftwareSources,
	"SYNC_MANAGEMENT_STATION_MIRROR": OperationTypesSyncManagementStationMirror,
	"PROMOTE_LIFECYCLE":              OperationTypesPromoteLifecycle,
}

var mappingOperationTypesEnumLowerCase = map[string]OperationTypesEnum{
	"install_packages":               OperationTypesInstallPackages,
	"update_packages":                OperationTypesUpdatePackages,
	"remove_packages":                OperationTypesRemovePackages,
	"update_all":                     OperationTypesUpdateAll,
	"update_security":                OperationTypesUpdateSecurity,
	"update_bugfix":                  OperationTypesUpdateBugfix,
	"update_enhancement":             OperationTypesUpdateEnhancement,
	"update_other":                   OperationTypesUpdateOther,
	"update_ksplice_userspace":       OperationTypesUpdateKspliceUserspace,
	"update_ksplice_kernel":          OperationTypesUpdateKspliceKernel,
	"manage_module_streams":          OperationTypesManageModuleStreams,
	"switch_module_stream":           OperationTypesSwitchModuleStream,
	"attach_software_sources":        OperationTypesAttachSoftwareSources,
	"detach_software_sources":        OperationTypesDetachSoftwareSources,
	"sync_management_station_mirror": OperationTypesSyncManagementStationMirror,
	"promote_lifecycle":              OperationTypesPromoteLifecycle,
}

// GetOperationTypesEnumValues Enumerates the set of values for OperationTypesEnum
func GetOperationTypesEnumValues() []OperationTypesEnum {
	values := make([]OperationTypesEnum, 0)
	for _, v := range mappingOperationTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetOperationTypesEnumStringValues Enumerates the set of values in String for OperationTypesEnum
func GetOperationTypesEnumStringValues() []string {
	return []string{
		"INSTALL_PACKAGES",
		"UPDATE_PACKAGES",
		"REMOVE_PACKAGES",
		"UPDATE_ALL",
		"UPDATE_SECURITY",
		"UPDATE_BUGFIX",
		"UPDATE_ENHANCEMENT",
		"UPDATE_OTHER",
		"UPDATE_KSPLICE_USERSPACE",
		"UPDATE_KSPLICE_KERNEL",
		"MANAGE_MODULE_STREAMS",
		"SWITCH_MODULE_STREAM",
		"ATTACH_SOFTWARE_SOURCES",
		"DETACH_SOFTWARE_SOURCES",
		"SYNC_MANAGEMENT_STATION_MIRROR",
		"PROMOTE_LIFECYCLE",
	}
}

// GetMappingOperationTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypesEnum(val string) (OperationTypesEnum, bool) {
	enum, ok := mappingOperationTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
