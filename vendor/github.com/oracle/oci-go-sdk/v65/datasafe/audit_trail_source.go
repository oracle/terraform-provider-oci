// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"strings"
)

// AuditTrailSourceEnum Enum with underlying type: string
type AuditTrailSourceEnum string

// Set of constants representing the allowable values for AuditTrailSourceEnum
const (
	AuditTrailSourceTable AuditTrailSourceEnum = "TABLE"
	AuditTrailSourceFile  AuditTrailSourceEnum = "FILE"
)

var mappingAuditTrailSourceEnum = map[string]AuditTrailSourceEnum{
	"TABLE": AuditTrailSourceTable,
	"FILE":  AuditTrailSourceFile,
}

var mappingAuditTrailSourceEnumLowerCase = map[string]AuditTrailSourceEnum{
	"table": AuditTrailSourceTable,
	"file":  AuditTrailSourceFile,
}

// GetAuditTrailSourceEnumValues Enumerates the set of values for AuditTrailSourceEnum
func GetAuditTrailSourceEnumValues() []AuditTrailSourceEnum {
	values := make([]AuditTrailSourceEnum, 0)
	for _, v := range mappingAuditTrailSourceEnum {
		values = append(values, v)
	}
	return values
}

// GetAuditTrailSourceEnumStringValues Enumerates the set of values in String for AuditTrailSourceEnum
func GetAuditTrailSourceEnumStringValues() []string {
	return []string{
		"TABLE",
		"FILE",
	}
}

// GetMappingAuditTrailSourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAuditTrailSourceEnum(val string) (AuditTrailSourceEnum, bool) {
	enum, ok := mappingAuditTrailSourceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
