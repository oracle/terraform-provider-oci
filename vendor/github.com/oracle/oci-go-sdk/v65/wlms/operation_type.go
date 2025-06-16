// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// WebLogic Management Service API
//
// WebLogic Management Service is an OCI service that enables a unified view and management of WebLogic domains
// in Oracle Cloud Infrastructure. Features include on-demand patching of WebLogic domains, rollback of the
// last applied patch, discovery and management of WebLogic instances on a compute host.
//

package wlms

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeStartWlsdomain      OperationTypeEnum = "START_WLSDOMAIN"
	OperationTypeStopWlsdomain       OperationTypeEnum = "STOP_WLSDOMAIN"
	OperationTypeRestartWlsdomain    OperationTypeEnum = "RESTART_WLSDOMAIN"
	OperationTypeScanWlsdomain       OperationTypeEnum = "SCAN_WLSDOMAIN"
	OperationTypeApplyPatch          OperationTypeEnum = "APPLY_PATCH"
	OperationTypeScanManagedInstance OperationTypeEnum = "SCAN_MANAGED_INSTANCE"
	OperationTypeRestoreBackup       OperationTypeEnum = "RESTORE_BACKUP"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"START_WLSDOMAIN":       OperationTypeStartWlsdomain,
	"STOP_WLSDOMAIN":        OperationTypeStopWlsdomain,
	"RESTART_WLSDOMAIN":     OperationTypeRestartWlsdomain,
	"SCAN_WLSDOMAIN":        OperationTypeScanWlsdomain,
	"APPLY_PATCH":           OperationTypeApplyPatch,
	"SCAN_MANAGED_INSTANCE": OperationTypeScanManagedInstance,
	"RESTORE_BACKUP":        OperationTypeRestoreBackup,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"start_wlsdomain":       OperationTypeStartWlsdomain,
	"stop_wlsdomain":        OperationTypeStopWlsdomain,
	"restart_wlsdomain":     OperationTypeRestartWlsdomain,
	"scan_wlsdomain":        OperationTypeScanWlsdomain,
	"apply_patch":           OperationTypeApplyPatch,
	"scan_managed_instance": OperationTypeScanManagedInstance,
	"restore_backup":        OperationTypeRestoreBackup,
}

// GetOperationTypeEnumValues Enumerates the set of values for OperationTypeEnum
func GetOperationTypeEnumValues() []OperationTypeEnum {
	values := make([]OperationTypeEnum, 0)
	for _, v := range mappingOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOperationTypeEnumStringValues Enumerates the set of values in String for OperationTypeEnum
func GetOperationTypeEnumStringValues() []string {
	return []string{
		"START_WLSDOMAIN",
		"STOP_WLSDOMAIN",
		"RESTART_WLSDOMAIN",
		"SCAN_WLSDOMAIN",
		"APPLY_PATCH",
		"SCAN_MANAGED_INSTANCE",
		"RESTORE_BACKUP",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
