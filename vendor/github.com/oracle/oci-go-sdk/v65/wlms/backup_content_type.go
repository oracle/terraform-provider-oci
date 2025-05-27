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

// BackupContentTypeEnum Enum with underlying type: string
type BackupContentTypeEnum string

// Set of constants representing the allowable values for BackupContentTypeEnum
const (
	BackupContentTypeBinary BackupContentTypeEnum = "BINARY"
)

var mappingBackupContentTypeEnum = map[string]BackupContentTypeEnum{
	"BINARY": BackupContentTypeBinary,
}

var mappingBackupContentTypeEnumLowerCase = map[string]BackupContentTypeEnum{
	"binary": BackupContentTypeBinary,
}

// GetBackupContentTypeEnumValues Enumerates the set of values for BackupContentTypeEnum
func GetBackupContentTypeEnumValues() []BackupContentTypeEnum {
	values := make([]BackupContentTypeEnum, 0)
	for _, v := range mappingBackupContentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetBackupContentTypeEnumStringValues Enumerates the set of values in String for BackupContentTypeEnum
func GetBackupContentTypeEnumStringValues() []string {
	return []string{
		"BINARY",
	}
}

// GetMappingBackupContentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBackupContentTypeEnum(val string) (BackupContentTypeEnum, bool) {
	enum, ok := mappingBackupContentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
