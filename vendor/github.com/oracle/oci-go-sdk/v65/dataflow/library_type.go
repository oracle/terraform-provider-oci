// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// LibraryTypeEnum Enum with underlying type: string
type LibraryTypeEnum string

// Set of constants representing the allowable values for LibraryTypeEnum
const (
	LibraryTypeWorkspaceFile LibraryTypeEnum = "WORKSPACE_FILE"
	LibraryTypePythonPackage LibraryTypeEnum = "PYTHON_PACKAGE"
	LibraryTypeVolumeFile    LibraryTypeEnum = "VOLUME_FILE"
	LibraryTypeMavenModule   LibraryTypeEnum = "MAVEN_MODULE"
)

var mappingLibraryTypeEnum = map[string]LibraryTypeEnum{
	"WORKSPACE_FILE": LibraryTypeWorkspaceFile,
	"PYTHON_PACKAGE": LibraryTypePythonPackage,
	"VOLUME_FILE":    LibraryTypeVolumeFile,
	"MAVEN_MODULE":   LibraryTypeMavenModule,
}

var mappingLibraryTypeEnumLowerCase = map[string]LibraryTypeEnum{
	"workspace_file": LibraryTypeWorkspaceFile,
	"python_package": LibraryTypePythonPackage,
	"volume_file":    LibraryTypeVolumeFile,
	"maven_module":   LibraryTypeMavenModule,
}

// GetLibraryTypeEnumValues Enumerates the set of values for LibraryTypeEnum
func GetLibraryTypeEnumValues() []LibraryTypeEnum {
	values := make([]LibraryTypeEnum, 0)
	for _, v := range mappingLibraryTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetLibraryTypeEnumStringValues Enumerates the set of values in String for LibraryTypeEnum
func GetLibraryTypeEnumStringValues() []string {
	return []string{
		"WORKSPACE_FILE",
		"PYTHON_PACKAGE",
		"VOLUME_FILE",
		"MAVEN_MODULE",
	}
}

// GetMappingLibraryTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLibraryTypeEnum(val string) (LibraryTypeEnum, bool) {
	enum, ok := mappingLibraryTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
