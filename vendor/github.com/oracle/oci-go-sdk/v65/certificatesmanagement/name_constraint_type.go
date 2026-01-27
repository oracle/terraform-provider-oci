// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Certificates Service Management API
//
// API for managing certificates.
//

package certificatesmanagement

import (
	"strings"
)

// NameConstraintTypeEnum Enum with underlying type: string
type NameConstraintTypeEnum string

// Set of constants representing the allowable values for NameConstraintTypeEnum
const (
	NameConstraintTypeDirectoryName NameConstraintTypeEnum = "DIRECTORY_NAME"
	NameConstraintTypeDns           NameConstraintTypeEnum = "DNS"
	NameConstraintTypeIp            NameConstraintTypeEnum = "IP"
)

var mappingNameConstraintTypeEnum = map[string]NameConstraintTypeEnum{
	"DIRECTORY_NAME": NameConstraintTypeDirectoryName,
	"DNS":            NameConstraintTypeDns,
	"IP":             NameConstraintTypeIp,
}

var mappingNameConstraintTypeEnumLowerCase = map[string]NameConstraintTypeEnum{
	"directory_name": NameConstraintTypeDirectoryName,
	"dns":            NameConstraintTypeDns,
	"ip":             NameConstraintTypeIp,
}

// GetNameConstraintTypeEnumValues Enumerates the set of values for NameConstraintTypeEnum
func GetNameConstraintTypeEnumValues() []NameConstraintTypeEnum {
	values := make([]NameConstraintTypeEnum, 0)
	for _, v := range mappingNameConstraintTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetNameConstraintTypeEnumStringValues Enumerates the set of values in String for NameConstraintTypeEnum
func GetNameConstraintTypeEnumStringValues() []string {
	return []string{
		"DIRECTORY_NAME",
		"DNS",
		"IP",
	}
}

// GetMappingNameConstraintTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNameConstraintTypeEnum(val string) (NameConstraintTypeEnum, bool) {
	enum, ok := mappingNameConstraintTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
