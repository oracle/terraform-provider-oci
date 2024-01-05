// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OpenSearch Service API
//
// The OpenSearch service API provides access to OCI Search Service with OpenSearch.
//

package opensearch

import (
	"strings"
)

// BackupStateEnum Enum with underlying type: string
type BackupStateEnum string

// Set of constants representing the allowable values for BackupStateEnum
const (
	BackupStateDeleted BackupStateEnum = "DELETED"
	BackupStateSuccess BackupStateEnum = "SUCCESS"
	BackupStateFailed  BackupStateEnum = "FAILED"
)

var mappingBackupStateEnum = map[string]BackupStateEnum{
	"DELETED": BackupStateDeleted,
	"SUCCESS": BackupStateSuccess,
	"FAILED":  BackupStateFailed,
}

var mappingBackupStateEnumLowerCase = map[string]BackupStateEnum{
	"deleted": BackupStateDeleted,
	"success": BackupStateSuccess,
	"failed":  BackupStateFailed,
}

// GetBackupStateEnumValues Enumerates the set of values for BackupStateEnum
func GetBackupStateEnumValues() []BackupStateEnum {
	values := make([]BackupStateEnum, 0)
	for _, v := range mappingBackupStateEnum {
		values = append(values, v)
	}
	return values
}

// GetBackupStateEnumStringValues Enumerates the set of values in String for BackupStateEnum
func GetBackupStateEnumStringValues() []string {
	return []string{
		"DELETED",
		"SUCCESS",
		"FAILED",
	}
}

// GetMappingBackupStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBackupStateEnum(val string) (BackupStateEnum, bool) {
	enum, ok := mappingBackupStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
