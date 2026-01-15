// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"strings"
)

// SupportedCapabilitiesEnum Enum with underlying type: string
type SupportedCapabilitiesEnum string

// Set of constants representing the allowable values for SupportedCapabilitiesEnum
const (
	SupportedCapabilitiesGroupToRole          SupportedCapabilitiesEnum = "GROUP_TO_ROLE"
	SupportedCapabilitiesBackupRestore        SupportedCapabilitiesEnum = "BACKUP_RESTORE"
	SupportedCapabilitiesCopyBackup           SupportedCapabilitiesEnum = "COPY_BACKUP"
	SupportedCapabilitiesManualBackup         SupportedCapabilitiesEnum = "MANUAL_BACKUP"
	SupportedCapabilitiesPlacement            SupportedCapabilitiesEnum = "PLACEMENT"
	SupportedCapabilitiesDisasterRecovery     SupportedCapabilitiesEnum = "DISASTER_RECOVERY"
	SupportedCapabilitiesScheduleManualBackup SupportedCapabilitiesEnum = "SCHEDULE_MANUAL_BACKUP"
	SupportedCapabilitiesMulticloud           SupportedCapabilitiesEnum = "MULTICLOUD"
)

var mappingSupportedCapabilitiesEnum = map[string]SupportedCapabilitiesEnum{
	"GROUP_TO_ROLE":          SupportedCapabilitiesGroupToRole,
	"BACKUP_RESTORE":         SupportedCapabilitiesBackupRestore,
	"COPY_BACKUP":            SupportedCapabilitiesCopyBackup,
	"MANUAL_BACKUP":          SupportedCapabilitiesManualBackup,
	"PLACEMENT":              SupportedCapabilitiesPlacement,
	"DISASTER_RECOVERY":      SupportedCapabilitiesDisasterRecovery,
	"SCHEDULE_MANUAL_BACKUP": SupportedCapabilitiesScheduleManualBackup,
	"MULTICLOUD":             SupportedCapabilitiesMulticloud,
}

var mappingSupportedCapabilitiesEnumLowerCase = map[string]SupportedCapabilitiesEnum{
	"group_to_role":          SupportedCapabilitiesGroupToRole,
	"backup_restore":         SupportedCapabilitiesBackupRestore,
	"copy_backup":            SupportedCapabilitiesCopyBackup,
	"manual_backup":          SupportedCapabilitiesManualBackup,
	"placement":              SupportedCapabilitiesPlacement,
	"disaster_recovery":      SupportedCapabilitiesDisasterRecovery,
	"schedule_manual_backup": SupportedCapabilitiesScheduleManualBackup,
	"multicloud":             SupportedCapabilitiesMulticloud,
}

// GetSupportedCapabilitiesEnumValues Enumerates the set of values for SupportedCapabilitiesEnum
func GetSupportedCapabilitiesEnumValues() []SupportedCapabilitiesEnum {
	values := make([]SupportedCapabilitiesEnum, 0)
	for _, v := range mappingSupportedCapabilitiesEnum {
		values = append(values, v)
	}
	return values
}

// GetSupportedCapabilitiesEnumStringValues Enumerates the set of values in String for SupportedCapabilitiesEnum
func GetSupportedCapabilitiesEnumStringValues() []string {
	return []string{
		"GROUP_TO_ROLE",
		"BACKUP_RESTORE",
		"COPY_BACKUP",
		"MANUAL_BACKUP",
		"PLACEMENT",
		"DISASTER_RECOVERY",
		"SCHEDULE_MANUAL_BACKUP",
		"MULTICLOUD",
	}
}

// GetMappingSupportedCapabilitiesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSupportedCapabilitiesEnum(val string) (SupportedCapabilitiesEnum, bool) {
	enum, ok := mappingSupportedCapabilitiesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
