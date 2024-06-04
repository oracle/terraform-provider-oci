// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// LibraryTypeEnum Enum with underlying type: string
type LibraryTypeEnum string

// Set of constants representing the allowable values for LibraryTypeEnum
const (
	LibraryTypeLogReaderComponent LibraryTypeEnum = "LOG_READER_COMPONENT"
)

var mappingLibraryTypeEnum = map[string]LibraryTypeEnum{
	"LOG_READER_COMPONENT": LibraryTypeLogReaderComponent,
}

var mappingLibraryTypeEnumLowerCase = map[string]LibraryTypeEnum{
	"log_reader_component": LibraryTypeLogReaderComponent,
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
		"LOG_READER_COMPONENT",
	}
}

// GetMappingLibraryTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLibraryTypeEnum(val string) (LibraryTypeEnum, bool) {
	enum, ok := mappingLibraryTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
