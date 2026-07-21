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

// OracleAiDataCatalogIcebergCatalogConfigTypeEnum Enum with underlying type: string
type OracleAiDataCatalogIcebergCatalogConfigTypeEnum string

// Set of constants representing the allowable values for OracleAiDataCatalogIcebergCatalogConfigTypeEnum
const (
	OracleAiDataCatalogIcebergCatalogConfigTypeOciGoldengate OracleAiDataCatalogIcebergCatalogConfigTypeEnum = "OCI_GOLDENGATE"
	OracleAiDataCatalogIcebergCatalogConfigTypeDbCredentials OracleAiDataCatalogIcebergCatalogConfigTypeEnum = "DB_CREDENTIALS"
)

var mappingOracleAiDataCatalogIcebergCatalogConfigTypeEnum = map[string]OracleAiDataCatalogIcebergCatalogConfigTypeEnum{
	"OCI_GOLDENGATE": OracleAiDataCatalogIcebergCatalogConfigTypeOciGoldengate,
	"DB_CREDENTIALS": OracleAiDataCatalogIcebergCatalogConfigTypeDbCredentials,
}

var mappingOracleAiDataCatalogIcebergCatalogConfigTypeEnumLowerCase = map[string]OracleAiDataCatalogIcebergCatalogConfigTypeEnum{
	"oci_goldengate": OracleAiDataCatalogIcebergCatalogConfigTypeOciGoldengate,
	"db_credentials": OracleAiDataCatalogIcebergCatalogConfigTypeDbCredentials,
}

// GetOracleAiDataCatalogIcebergCatalogConfigTypeEnumValues Enumerates the set of values for OracleAiDataCatalogIcebergCatalogConfigTypeEnum
func GetOracleAiDataCatalogIcebergCatalogConfigTypeEnumValues() []OracleAiDataCatalogIcebergCatalogConfigTypeEnum {
	values := make([]OracleAiDataCatalogIcebergCatalogConfigTypeEnum, 0)
	for _, v := range mappingOracleAiDataCatalogIcebergCatalogConfigTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOracleAiDataCatalogIcebergCatalogConfigTypeEnumStringValues Enumerates the set of values in String for OracleAiDataCatalogIcebergCatalogConfigTypeEnum
func GetOracleAiDataCatalogIcebergCatalogConfigTypeEnumStringValues() []string {
	return []string{
		"OCI_GOLDENGATE",
		"DB_CREDENTIALS",
	}
}

// GetMappingOracleAiDataCatalogIcebergCatalogConfigTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOracleAiDataCatalogIcebergCatalogConfigTypeEnum(val string) (OracleAiDataCatalogIcebergCatalogConfigTypeEnum, bool) {
	enum, ok := mappingOracleAiDataCatalogIcebergCatalogConfigTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
