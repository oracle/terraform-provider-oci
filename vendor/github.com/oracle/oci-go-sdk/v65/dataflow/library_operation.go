// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// LibraryOperationEnum Enum with underlying type: string
type LibraryOperationEnum string

// Set of constants representing the allowable values for LibraryOperationEnum
const (
	LibraryOperationInstall   LibraryOperationEnum = "INSTALL"
	LibraryOperationUninstall LibraryOperationEnum = "UNINSTALL"
)

var mappingLibraryOperationEnum = map[string]LibraryOperationEnum{
	"INSTALL":   LibraryOperationInstall,
	"UNINSTALL": LibraryOperationUninstall,
}

var mappingLibraryOperationEnumLowerCase = map[string]LibraryOperationEnum{
	"install":   LibraryOperationInstall,
	"uninstall": LibraryOperationUninstall,
}

// GetLibraryOperationEnumValues Enumerates the set of values for LibraryOperationEnum
func GetLibraryOperationEnumValues() []LibraryOperationEnum {
	values := make([]LibraryOperationEnum, 0)
	for _, v := range mappingLibraryOperationEnum {
		values = append(values, v)
	}
	return values
}

// GetLibraryOperationEnumStringValues Enumerates the set of values in String for LibraryOperationEnum
func GetLibraryOperationEnumStringValues() []string {
	return []string{
		"INSTALL",
		"UNINSTALL",
	}
}

// GetMappingLibraryOperationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLibraryOperationEnum(val string) (LibraryOperationEnum, bool) {
	enum, ok := mappingLibraryOperationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
