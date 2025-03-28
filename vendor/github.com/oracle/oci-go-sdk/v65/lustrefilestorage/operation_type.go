// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// File Storage with Lustre API
//
// Use the File Storage with Lustre API to manage Lustre file systems and related resources. For more information, see File Storage with Lustre (https://docs.oracle.com/iaas/Content/lustre/home.htm).
//

package lustrefilestorage

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateLustreFileSystem OperationTypeEnum = "CREATE_LUSTRE_FILE_SYSTEM"
	OperationTypeUpdateLustreFileSystem OperationTypeEnum = "UPDATE_LUSTRE_FILE_SYSTEM"
	OperationTypeDeleteLustreFileSystem OperationTypeEnum = "DELETE_LUSTRE_FILE_SYSTEM"
	OperationTypeMoveLustreFileSystem   OperationTypeEnum = "MOVE_LUSTRE_FILE_SYSTEM"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_LUSTRE_FILE_SYSTEM": OperationTypeCreateLustreFileSystem,
	"UPDATE_LUSTRE_FILE_SYSTEM": OperationTypeUpdateLustreFileSystem,
	"DELETE_LUSTRE_FILE_SYSTEM": OperationTypeDeleteLustreFileSystem,
	"MOVE_LUSTRE_FILE_SYSTEM":   OperationTypeMoveLustreFileSystem,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_lustre_file_system": OperationTypeCreateLustreFileSystem,
	"update_lustre_file_system": OperationTypeUpdateLustreFileSystem,
	"delete_lustre_file_system": OperationTypeDeleteLustreFileSystem,
	"move_lustre_file_system":   OperationTypeMoveLustreFileSystem,
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
		"CREATE_LUSTRE_FILE_SYSTEM",
		"UPDATE_LUSTRE_FILE_SYSTEM",
		"DELETE_LUSTRE_FILE_SYSTEM",
		"MOVE_LUSTRE_FILE_SYSTEM",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
