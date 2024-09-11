// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Flow API
//
// Use the Data Flow APIs to run any Apache Spark application at any scale without deploying or managing any infrastructure.
//

package dataflow

import (
	"strings"
)

// LibraryStatusEnum Enum with underlying type: string
type LibraryStatusEnum string

// Set of constants representing the allowable values for LibraryStatusEnum
const (
	LibraryStatusPending            LibraryStatusEnum = "PENDING"
	LibraryStatusDownloading        LibraryStatusEnum = "DOWNLOADING"
	LibraryStatusInstalling         LibraryStatusEnum = "INSTALLING"
	LibraryStatusInstalled          LibraryStatusEnum = "INSTALLED"
	LibraryStatusFailed             LibraryStatusEnum = "FAILED"
	LibraryStatusSkipped            LibraryStatusEnum = "SKIPPED"
	LibraryStatusUninstallOnRestart LibraryStatusEnum = "UNINSTALL_ON_RESTART"
	LibraryStatusDeleted            LibraryStatusEnum = "DELETED"
)

var mappingLibraryStatusEnum = map[string]LibraryStatusEnum{
	"PENDING":              LibraryStatusPending,
	"DOWNLOADING":          LibraryStatusDownloading,
	"INSTALLING":           LibraryStatusInstalling,
	"INSTALLED":            LibraryStatusInstalled,
	"FAILED":               LibraryStatusFailed,
	"SKIPPED":              LibraryStatusSkipped,
	"UNINSTALL_ON_RESTART": LibraryStatusUninstallOnRestart,
	"DELETED":              LibraryStatusDeleted,
}

var mappingLibraryStatusEnumLowerCase = map[string]LibraryStatusEnum{
	"pending":              LibraryStatusPending,
	"downloading":          LibraryStatusDownloading,
	"installing":           LibraryStatusInstalling,
	"installed":            LibraryStatusInstalled,
	"failed":               LibraryStatusFailed,
	"skipped":              LibraryStatusSkipped,
	"uninstall_on_restart": LibraryStatusUninstallOnRestart,
	"deleted":              LibraryStatusDeleted,
}

// GetLibraryStatusEnumValues Enumerates the set of values for LibraryStatusEnum
func GetLibraryStatusEnumValues() []LibraryStatusEnum {
	values := make([]LibraryStatusEnum, 0)
	for _, v := range mappingLibraryStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetLibraryStatusEnumStringValues Enumerates the set of values in String for LibraryStatusEnum
func GetLibraryStatusEnumStringValues() []string {
	return []string{
		"PENDING",
		"DOWNLOADING",
		"INSTALLING",
		"INSTALLED",
		"FAILED",
		"SKIPPED",
		"UNINSTALL_ON_RESTART",
		"DELETED",
	}
}

// GetMappingLibraryStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLibraryStatusEnum(val string) (LibraryStatusEnum, bool) {
	enum, ok := mappingLibraryStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
