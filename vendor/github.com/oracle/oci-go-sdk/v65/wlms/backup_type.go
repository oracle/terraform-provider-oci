// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// BackupTypeEnum Enum with underlying type: string
type BackupTypeEnum string

// Set of constants representing the allowable values for BackupTypeEnum
const (
	BackupTypeLocalFile BackupTypeEnum = "LOCAL_FILE"
)

var mappingBackupTypeEnum = map[string]BackupTypeEnum{
	"LOCAL_FILE": BackupTypeLocalFile,
}

var mappingBackupTypeEnumLowerCase = map[string]BackupTypeEnum{
	"local_file": BackupTypeLocalFile,
}

// GetBackupTypeEnumValues Enumerates the set of values for BackupTypeEnum
func GetBackupTypeEnumValues() []BackupTypeEnum {
	values := make([]BackupTypeEnum, 0)
	for _, v := range mappingBackupTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetBackupTypeEnumStringValues Enumerates the set of values in String for BackupTypeEnum
func GetBackupTypeEnumStringValues() []string {
	return []string{
		"LOCAL_FILE",
	}
}

// GetMappingBackupTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBackupTypeEnum(val string) (BackupTypeEnum, bool) {
	enum, ok := mappingBackupTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
