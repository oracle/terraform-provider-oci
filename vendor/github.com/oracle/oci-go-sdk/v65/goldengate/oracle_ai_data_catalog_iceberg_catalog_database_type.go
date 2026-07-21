// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// OracleAiDataCatalogIcebergCatalogDatabaseTypeEnum Enum with underlying type: string
type OracleAiDataCatalogIcebergCatalogDatabaseTypeEnum string

// Set of constants representing the allowable values for OracleAiDataCatalogIcebergCatalogDatabaseTypeEnum
const (
	OracleAiDataCatalogIcebergCatalogDatabaseTypeAdb OracleAiDataCatalogIcebergCatalogDatabaseTypeEnum = "ADB"
)

var mappingOracleAiDataCatalogIcebergCatalogDatabaseTypeEnum = map[string]OracleAiDataCatalogIcebergCatalogDatabaseTypeEnum{
	"ADB": OracleAiDataCatalogIcebergCatalogDatabaseTypeAdb,
}

var mappingOracleAiDataCatalogIcebergCatalogDatabaseTypeEnumLowerCase = map[string]OracleAiDataCatalogIcebergCatalogDatabaseTypeEnum{
	"adb": OracleAiDataCatalogIcebergCatalogDatabaseTypeAdb,
}

// GetOracleAiDataCatalogIcebergCatalogDatabaseTypeEnumValues Enumerates the set of values for OracleAiDataCatalogIcebergCatalogDatabaseTypeEnum
func GetOracleAiDataCatalogIcebergCatalogDatabaseTypeEnumValues() []OracleAiDataCatalogIcebergCatalogDatabaseTypeEnum {
	values := make([]OracleAiDataCatalogIcebergCatalogDatabaseTypeEnum, 0)
	for _, v := range mappingOracleAiDataCatalogIcebergCatalogDatabaseTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOracleAiDataCatalogIcebergCatalogDatabaseTypeEnumStringValues Enumerates the set of values in String for OracleAiDataCatalogIcebergCatalogDatabaseTypeEnum
func GetOracleAiDataCatalogIcebergCatalogDatabaseTypeEnumStringValues() []string {
	return []string{
		"ADB",
	}
}

// GetMappingOracleAiDataCatalogIcebergCatalogDatabaseTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOracleAiDataCatalogIcebergCatalogDatabaseTypeEnum(val string) (OracleAiDataCatalogIcebergCatalogDatabaseTypeEnum, bool) {
	enum, ok := mappingOracleAiDataCatalogIcebergCatalogDatabaseTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
