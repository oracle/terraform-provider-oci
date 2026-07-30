// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// TargetDatabaseTypeEnum Enum with underlying type: string
type TargetDatabaseTypeEnum string

// Set of constants representing the allowable values for TargetDatabaseTypeEnum
const (
	TargetDatabaseTypeAdbs  TargetDatabaseTypeEnum = "ADBS"
	TargetDatabaseTypeAdbd  TargetDatabaseTypeEnum = "ADBD"
	TargetDatabaseTypeExacs TargetDatabaseTypeEnum = "EXACS"
	TargetDatabaseTypeDbcs  TargetDatabaseTypeEnum = "DBCS"
	TargetDatabaseTypeNa    TargetDatabaseTypeEnum = "NA"
)

var mappingTargetDatabaseTypeEnum = map[string]TargetDatabaseTypeEnum{
	"ADBS":  TargetDatabaseTypeAdbs,
	"ADBD":  TargetDatabaseTypeAdbd,
	"EXACS": TargetDatabaseTypeExacs,
	"DBCS":  TargetDatabaseTypeDbcs,
	"NA":    TargetDatabaseTypeNa,
}

var mappingTargetDatabaseTypeEnumLowerCase = map[string]TargetDatabaseTypeEnum{
	"adbs":  TargetDatabaseTypeAdbs,
	"adbd":  TargetDatabaseTypeAdbd,
	"exacs": TargetDatabaseTypeExacs,
	"dbcs":  TargetDatabaseTypeDbcs,
	"na":    TargetDatabaseTypeNa,
}

// GetTargetDatabaseTypeEnumValues Enumerates the set of values for TargetDatabaseTypeEnum
func GetTargetDatabaseTypeEnumValues() []TargetDatabaseTypeEnum {
	values := make([]TargetDatabaseTypeEnum, 0)
	for _, v := range mappingTargetDatabaseTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTargetDatabaseTypeEnumStringValues Enumerates the set of values in String for TargetDatabaseTypeEnum
func GetTargetDatabaseTypeEnumStringValues() []string {
	return []string{
		"ADBS",
		"ADBD",
		"EXACS",
		"DBCS",
		"NA",
	}
}

// GetMappingTargetDatabaseTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTargetDatabaseTypeEnum(val string) (TargetDatabaseTypeEnum, bool) {
	enum, ok := mappingTargetDatabaseTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
