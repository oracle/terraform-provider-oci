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
	"strings"
)

// WorkRequestOperationTypeEnum Enum with underlying type: string
type WorkRequestOperationTypeEnum string

// Set of constants representing the allowable values for WorkRequestOperationTypeEnum
const (
	WorkRequestOperationTypeInstallPackages                  WorkRequestOperationTypeEnum = "INSTALL_PACKAGES"
	WorkRequestOperationTypeRemovePackages                   WorkRequestOperationTypeEnum = "REMOVE_PACKAGES"
	WorkRequestOperationTypeUpdatePackages                   WorkRequestOperationTypeEnum = "UPDATE_PACKAGES"
	WorkRequestOperationTypeUpdateAllPackages                WorkRequestOperationTypeEnum = "UPDATE_ALL_PACKAGES"
	WorkRequestOperationTypeUpdateSecurity                   WorkRequestOperationTypeEnum = "UPDATE_SECURITY"
	WorkRequestOperationTypeUpdateBugfix                     WorkRequestOperationTypeEnum = "UPDATE_BUGFIX"
	WorkRequestOperationTypeUpdateEnhancement                WorkRequestOperationTypeEnum = "UPDATE_ENHANCEMENT"
	WorkRequestOperationTypeUpdateOther                      WorkRequestOperationTypeEnum = "UPDATE_OTHER"
	WorkRequestOperationTypeUpdateKspliceKernel              WorkRequestOperationTypeEnum = "UPDATE_KSPLICE_KERNEL"
	WorkRequestOperationTypeUpdateKspliceUserspace           WorkRequestOperationTypeEnum = "UPDATE_KSPLICE_USERSPACE"
	WorkRequestOperationTypeEnableModuleStreams              WorkRequestOperationTypeEnum = "ENABLE_MODULE_STREAMS"
	WorkRequestOperationTypeDisableModuleStreams             WorkRequestOperationTypeEnum = "DISABLE_MODULE_STREAMS"
	WorkRequestOperationTypeSwitchModuleStream               WorkRequestOperationTypeEnum = "SWITCH_MODULE_STREAM"
	WorkRequestOperationTypeInstallModuleProfiles            WorkRequestOperationTypeEnum = "INSTALL_MODULE_PROFILES"
	WorkRequestOperationTypeRemoveModuleProfiles             WorkRequestOperationTypeEnum = "REMOVE_MODULE_PROFILES"
	WorkRequestOperationTypeSetSoftwareSources               WorkRequestOperationTypeEnum = "SET_SOFTWARE_SOURCES"
	WorkRequestOperationTypeListPackages                     WorkRequestOperationTypeEnum = "LIST_PACKAGES"
	WorkRequestOperationTypeSetManagementStationConfig       WorkRequestOperationTypeEnum = "SET_MANAGEMENT_STATION_CONFIG"
	WorkRequestOperationTypeSyncManagementStationMirror      WorkRequestOperationTypeEnum = "SYNC_MANAGEMENT_STATION_MIRROR"
	WorkRequestOperationTypeUpdateManagementStationSoftware  WorkRequestOperationTypeEnum = "UPDATE_MANAGEMENT_STATION_SOFTWARE"
	WorkRequestOperationTypeUpdate                           WorkRequestOperationTypeEnum = "UPDATE"
	WorkRequestOperationTypeModuleActions                    WorkRequestOperationTypeEnum = "MODULE_ACTIONS"
	WorkRequestOperationTypeLifecyclePromotion               WorkRequestOperationTypeEnum = "LIFECYCLE_PROMOTION"
	WorkRequestOperationTypeCreateSoftwareSource             WorkRequestOperationTypeEnum = "CREATE_SOFTWARE_SOURCE"
	WorkRequestOperationTypeUpdateSoftwareSource             WorkRequestOperationTypeEnum = "UPDATE_SOFTWARE_SOURCE"
	WorkRequestOperationTypeImportContent                    WorkRequestOperationTypeEnum = "IMPORT_CONTENT"
	WorkRequestOperationTypeSyncAgentConfig                  WorkRequestOperationTypeEnum = "SYNC_AGENT_CONFIG"
	WorkRequestOperationTypeInstallWindowsUpdates            WorkRequestOperationTypeEnum = "INSTALL_WINDOWS_UPDATES"
	WorkRequestOperationTypeListWindowsUpdate                WorkRequestOperationTypeEnum = "LIST_WINDOWS_UPDATE"
	WorkRequestOperationTypeGetWindowsUpdateDetails          WorkRequestOperationTypeEnum = "GET_WINDOWS_UPDATE_DETAILS"
	WorkRequestOperationTypeInstallAllWindowsUpdates         WorkRequestOperationTypeEnum = "INSTALL_ALL_WINDOWS_UPDATES"
	WorkRequestOperationTypeInstallSecurityWindowsUpdates    WorkRequestOperationTypeEnum = "INSTALL_SECURITY_WINDOWS_UPDATES"
	WorkRequestOperationTypeInstallBugfixWindowsUpdates      WorkRequestOperationTypeEnum = "INSTALL_BUGFIX_WINDOWS_UPDATES"
	WorkRequestOperationTypeInstallEnhancementWindowsUpdates WorkRequestOperationTypeEnum = "INSTALL_ENHANCEMENT_WINDOWS_UPDATES"
	WorkRequestOperationTypeInstallOtherWindowsUpdates       WorkRequestOperationTypeEnum = "INSTALL_OTHER_WINDOWS_UPDATES"
	WorkRequestOperationTypeRemoveContent                    WorkRequestOperationTypeEnum = "REMOVE_CONTENT"
	WorkRequestOperationTypeUnregisterManagedInstance        WorkRequestOperationTypeEnum = "UNREGISTER_MANAGED_INSTANCE"
)

var mappingWorkRequestOperationTypeEnum = map[string]WorkRequestOperationTypeEnum{
	"INSTALL_PACKAGES":                    WorkRequestOperationTypeInstallPackages,
	"REMOVE_PACKAGES":                     WorkRequestOperationTypeRemovePackages,
	"UPDATE_PACKAGES":                     WorkRequestOperationTypeUpdatePackages,
	"UPDATE_ALL_PACKAGES":                 WorkRequestOperationTypeUpdateAllPackages,
	"UPDATE_SECURITY":                     WorkRequestOperationTypeUpdateSecurity,
	"UPDATE_BUGFIX":                       WorkRequestOperationTypeUpdateBugfix,
	"UPDATE_ENHANCEMENT":                  WorkRequestOperationTypeUpdateEnhancement,
	"UPDATE_OTHER":                        WorkRequestOperationTypeUpdateOther,
	"UPDATE_KSPLICE_KERNEL":               WorkRequestOperationTypeUpdateKspliceKernel,
	"UPDATE_KSPLICE_USERSPACE":            WorkRequestOperationTypeUpdateKspliceUserspace,
	"ENABLE_MODULE_STREAMS":               WorkRequestOperationTypeEnableModuleStreams,
	"DISABLE_MODULE_STREAMS":              WorkRequestOperationTypeDisableModuleStreams,
	"SWITCH_MODULE_STREAM":                WorkRequestOperationTypeSwitchModuleStream,
	"INSTALL_MODULE_PROFILES":             WorkRequestOperationTypeInstallModuleProfiles,
	"REMOVE_MODULE_PROFILES":              WorkRequestOperationTypeRemoveModuleProfiles,
	"SET_SOFTWARE_SOURCES":                WorkRequestOperationTypeSetSoftwareSources,
	"LIST_PACKAGES":                       WorkRequestOperationTypeListPackages,
	"SET_MANAGEMENT_STATION_CONFIG":       WorkRequestOperationTypeSetManagementStationConfig,
	"SYNC_MANAGEMENT_STATION_MIRROR":      WorkRequestOperationTypeSyncManagementStationMirror,
	"UPDATE_MANAGEMENT_STATION_SOFTWARE":  WorkRequestOperationTypeUpdateManagementStationSoftware,
	"UPDATE":                              WorkRequestOperationTypeUpdate,
	"MODULE_ACTIONS":                      WorkRequestOperationTypeModuleActions,
	"LIFECYCLE_PROMOTION":                 WorkRequestOperationTypeLifecyclePromotion,
	"CREATE_SOFTWARE_SOURCE":              WorkRequestOperationTypeCreateSoftwareSource,
	"UPDATE_SOFTWARE_SOURCE":              WorkRequestOperationTypeUpdateSoftwareSource,
	"IMPORT_CONTENT":                      WorkRequestOperationTypeImportContent,
	"SYNC_AGENT_CONFIG":                   WorkRequestOperationTypeSyncAgentConfig,
	"INSTALL_WINDOWS_UPDATES":             WorkRequestOperationTypeInstallWindowsUpdates,
	"LIST_WINDOWS_UPDATE":                 WorkRequestOperationTypeListWindowsUpdate,
	"GET_WINDOWS_UPDATE_DETAILS":          WorkRequestOperationTypeGetWindowsUpdateDetails,
	"INSTALL_ALL_WINDOWS_UPDATES":         WorkRequestOperationTypeInstallAllWindowsUpdates,
	"INSTALL_SECURITY_WINDOWS_UPDATES":    WorkRequestOperationTypeInstallSecurityWindowsUpdates,
	"INSTALL_BUGFIX_WINDOWS_UPDATES":      WorkRequestOperationTypeInstallBugfixWindowsUpdates,
	"INSTALL_ENHANCEMENT_WINDOWS_UPDATES": WorkRequestOperationTypeInstallEnhancementWindowsUpdates,
	"INSTALL_OTHER_WINDOWS_UPDATES":       WorkRequestOperationTypeInstallOtherWindowsUpdates,
	"REMOVE_CONTENT":                      WorkRequestOperationTypeRemoveContent,
	"UNREGISTER_MANAGED_INSTANCE":         WorkRequestOperationTypeUnregisterManagedInstance,
}

var mappingWorkRequestOperationTypeEnumLowerCase = map[string]WorkRequestOperationTypeEnum{
	"install_packages":                    WorkRequestOperationTypeInstallPackages,
	"remove_packages":                     WorkRequestOperationTypeRemovePackages,
	"update_packages":                     WorkRequestOperationTypeUpdatePackages,
	"update_all_packages":                 WorkRequestOperationTypeUpdateAllPackages,
	"update_security":                     WorkRequestOperationTypeUpdateSecurity,
	"update_bugfix":                       WorkRequestOperationTypeUpdateBugfix,
	"update_enhancement":                  WorkRequestOperationTypeUpdateEnhancement,
	"update_other":                        WorkRequestOperationTypeUpdateOther,
	"update_ksplice_kernel":               WorkRequestOperationTypeUpdateKspliceKernel,
	"update_ksplice_userspace":            WorkRequestOperationTypeUpdateKspliceUserspace,
	"enable_module_streams":               WorkRequestOperationTypeEnableModuleStreams,
	"disable_module_streams":              WorkRequestOperationTypeDisableModuleStreams,
	"switch_module_stream":                WorkRequestOperationTypeSwitchModuleStream,
	"install_module_profiles":             WorkRequestOperationTypeInstallModuleProfiles,
	"remove_module_profiles":              WorkRequestOperationTypeRemoveModuleProfiles,
	"set_software_sources":                WorkRequestOperationTypeSetSoftwareSources,
	"list_packages":                       WorkRequestOperationTypeListPackages,
	"set_management_station_config":       WorkRequestOperationTypeSetManagementStationConfig,
	"sync_management_station_mirror":      WorkRequestOperationTypeSyncManagementStationMirror,
	"update_management_station_software":  WorkRequestOperationTypeUpdateManagementStationSoftware,
	"update":                              WorkRequestOperationTypeUpdate,
	"module_actions":                      WorkRequestOperationTypeModuleActions,
	"lifecycle_promotion":                 WorkRequestOperationTypeLifecyclePromotion,
	"create_software_source":              WorkRequestOperationTypeCreateSoftwareSource,
	"update_software_source":              WorkRequestOperationTypeUpdateSoftwareSource,
	"import_content":                      WorkRequestOperationTypeImportContent,
	"sync_agent_config":                   WorkRequestOperationTypeSyncAgentConfig,
	"install_windows_updates":             WorkRequestOperationTypeInstallWindowsUpdates,
	"list_windows_update":                 WorkRequestOperationTypeListWindowsUpdate,
	"get_windows_update_details":          WorkRequestOperationTypeGetWindowsUpdateDetails,
	"install_all_windows_updates":         WorkRequestOperationTypeInstallAllWindowsUpdates,
	"install_security_windows_updates":    WorkRequestOperationTypeInstallSecurityWindowsUpdates,
	"install_bugfix_windows_updates":      WorkRequestOperationTypeInstallBugfixWindowsUpdates,
	"install_enhancement_windows_updates": WorkRequestOperationTypeInstallEnhancementWindowsUpdates,
	"install_other_windows_updates":       WorkRequestOperationTypeInstallOtherWindowsUpdates,
	"remove_content":                      WorkRequestOperationTypeRemoveContent,
	"unregister_managed_instance":         WorkRequestOperationTypeUnregisterManagedInstance,
}

// GetWorkRequestOperationTypeEnumValues Enumerates the set of values for WorkRequestOperationTypeEnum
func GetWorkRequestOperationTypeEnumValues() []WorkRequestOperationTypeEnum {
	values := make([]WorkRequestOperationTypeEnum, 0)
	for _, v := range mappingWorkRequestOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetWorkRequestOperationTypeEnumStringValues Enumerates the set of values in String for WorkRequestOperationTypeEnum
func GetWorkRequestOperationTypeEnumStringValues() []string {
	return []string{
		"INSTALL_PACKAGES",
		"REMOVE_PACKAGES",
		"UPDATE_PACKAGES",
		"UPDATE_ALL_PACKAGES",
		"UPDATE_SECURITY",
		"UPDATE_BUGFIX",
		"UPDATE_ENHANCEMENT",
		"UPDATE_OTHER",
		"UPDATE_KSPLICE_KERNEL",
		"UPDATE_KSPLICE_USERSPACE",
		"ENABLE_MODULE_STREAMS",
		"DISABLE_MODULE_STREAMS",
		"SWITCH_MODULE_STREAM",
		"INSTALL_MODULE_PROFILES",
		"REMOVE_MODULE_PROFILES",
		"SET_SOFTWARE_SOURCES",
		"LIST_PACKAGES",
		"SET_MANAGEMENT_STATION_CONFIG",
		"SYNC_MANAGEMENT_STATION_MIRROR",
		"UPDATE_MANAGEMENT_STATION_SOFTWARE",
		"UPDATE",
		"MODULE_ACTIONS",
		"LIFECYCLE_PROMOTION",
		"CREATE_SOFTWARE_SOURCE",
		"UPDATE_SOFTWARE_SOURCE",
		"IMPORT_CONTENT",
		"SYNC_AGENT_CONFIG",
		"INSTALL_WINDOWS_UPDATES",
		"LIST_WINDOWS_UPDATE",
		"GET_WINDOWS_UPDATE_DETAILS",
		"INSTALL_ALL_WINDOWS_UPDATES",
		"INSTALL_SECURITY_WINDOWS_UPDATES",
		"INSTALL_BUGFIX_WINDOWS_UPDATES",
		"INSTALL_ENHANCEMENT_WINDOWS_UPDATES",
		"INSTALL_OTHER_WINDOWS_UPDATES",
		"REMOVE_CONTENT",
		"UNREGISTER_MANAGED_INSTANCE",
	}
}

// GetMappingWorkRequestOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkRequestOperationTypeEnum(val string) (WorkRequestOperationTypeEnum, bool) {
	enum, ok := mappingWorkRequestOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
