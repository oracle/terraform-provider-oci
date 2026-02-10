// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"strings"
)

// DatabaseDataSizeEnum Enum with underlying type: string
type DatabaseDataSizeEnum string

// Set of constants representing the allowable values for DatabaseDataSizeEnum
const (
	DatabaseDataSizeLessThan1Gb     DatabaseDataSizeEnum = "LESS_THAN_1GB"
	DatabaseDataSizeGb110           DatabaseDataSizeEnum = "GB_1_10"
	DatabaseDataSizeGb1050          DatabaseDataSizeEnum = "GB_10_50"
	DatabaseDataSizeGb50100         DatabaseDataSizeEnum = "GB_50_100"
	DatabaseDataSizeGb100500        DatabaseDataSizeEnum = "GB_100_500"
	DatabaseDataSizeGb500Tb1        DatabaseDataSizeEnum = "GB_500_TB_1"
	DatabaseDataSizeTb13            DatabaseDataSizeEnum = "TB_1_3"
	DatabaseDataSizeTb310           DatabaseDataSizeEnum = "TB_3_10"
	DatabaseDataSizeTb1050          DatabaseDataSizeEnum = "TB_10_50"
	DatabaseDataSizeGreaterThan50Tb DatabaseDataSizeEnum = "GREATER_THAN_50TB"
)

var mappingDatabaseDataSizeEnum = map[string]DatabaseDataSizeEnum{
	"LESS_THAN_1GB":     DatabaseDataSizeLessThan1Gb,
	"GB_1_10":           DatabaseDataSizeGb110,
	"GB_10_50":          DatabaseDataSizeGb1050,
	"GB_50_100":         DatabaseDataSizeGb50100,
	"GB_100_500":        DatabaseDataSizeGb100500,
	"GB_500_TB_1":       DatabaseDataSizeGb500Tb1,
	"TB_1_3":            DatabaseDataSizeTb13,
	"TB_3_10":           DatabaseDataSizeTb310,
	"TB_10_50":          DatabaseDataSizeTb1050,
	"GREATER_THAN_50TB": DatabaseDataSizeGreaterThan50Tb,
}

var mappingDatabaseDataSizeEnumLowerCase = map[string]DatabaseDataSizeEnum{
	"less_than_1gb":     DatabaseDataSizeLessThan1Gb,
	"gb_1_10":           DatabaseDataSizeGb110,
	"gb_10_50":          DatabaseDataSizeGb1050,
	"gb_50_100":         DatabaseDataSizeGb50100,
	"gb_100_500":        DatabaseDataSizeGb100500,
	"gb_500_tb_1":       DatabaseDataSizeGb500Tb1,
	"tb_1_3":            DatabaseDataSizeTb13,
	"tb_3_10":           DatabaseDataSizeTb310,
	"tb_10_50":          DatabaseDataSizeTb1050,
	"greater_than_50tb": DatabaseDataSizeGreaterThan50Tb,
}

// GetDatabaseDataSizeEnumValues Enumerates the set of values for DatabaseDataSizeEnum
func GetDatabaseDataSizeEnumValues() []DatabaseDataSizeEnum {
	values := make([]DatabaseDataSizeEnum, 0)
	for _, v := range mappingDatabaseDataSizeEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseDataSizeEnumStringValues Enumerates the set of values in String for DatabaseDataSizeEnum
func GetDatabaseDataSizeEnumStringValues() []string {
	return []string{
		"LESS_THAN_1GB",
		"GB_1_10",
		"GB_10_50",
		"GB_50_100",
		"GB_100_500",
		"GB_500_TB_1",
		"TB_1_3",
		"TB_3_10",
		"TB_10_50",
		"GREATER_THAN_50TB",
	}
}

// GetMappingDatabaseDataSizeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseDataSizeEnum(val string) (DatabaseDataSizeEnum, bool) {
	enum, ok := mappingDatabaseDataSizeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
