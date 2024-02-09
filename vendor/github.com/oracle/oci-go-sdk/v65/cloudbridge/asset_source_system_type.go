// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Bridge API
//
// API for Oracle Cloud Bridge service.
//

package cloudbridge

import (
	"strings"
)

// AssetSourceSystemTypeEnum Enum with underlying type: string
type AssetSourceSystemTypeEnum string

// Set of constants representing the allowable values for AssetSourceSystemTypeEnum
const (
	AssetSourceSystemTypeVcenter      AssetSourceSystemTypeEnum = "VCENTER"
	AssetSourceSystemTypeAws          AssetSourceSystemTypeEnum = "AWS"
	AssetSourceSystemTypeOracleCdb    AssetSourceSystemTypeEnum = "ORACLE_CDB"
	AssetSourceSystemTypeOraclePdb    AssetSourceSystemTypeEnum = "ORACLE_PDB"
	AssetSourceSystemTypeOracleNonCdb AssetSourceSystemTypeEnum = "ORACLE_NON_CDB"
)

var mappingAssetSourceSystemTypeEnum = map[string]AssetSourceSystemTypeEnum{
	"VCENTER":        AssetSourceSystemTypeVcenter,
	"AWS":            AssetSourceSystemTypeAws,
	"ORACLE_CDB":     AssetSourceSystemTypeOracleCdb,
	"ORACLE_PDB":     AssetSourceSystemTypeOraclePdb,
	"ORACLE_NON_CDB": AssetSourceSystemTypeOracleNonCdb,
}

var mappingAssetSourceSystemTypeEnumLowerCase = map[string]AssetSourceSystemTypeEnum{
	"vcenter":        AssetSourceSystemTypeVcenter,
	"aws":            AssetSourceSystemTypeAws,
	"oracle_cdb":     AssetSourceSystemTypeOracleCdb,
	"oracle_pdb":     AssetSourceSystemTypeOraclePdb,
	"oracle_non_cdb": AssetSourceSystemTypeOracleNonCdb,
}

// GetAssetSourceSystemTypeEnumValues Enumerates the set of values for AssetSourceSystemTypeEnum
func GetAssetSourceSystemTypeEnumValues() []AssetSourceSystemTypeEnum {
	values := make([]AssetSourceSystemTypeEnum, 0)
	for _, v := range mappingAssetSourceSystemTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAssetSourceSystemTypeEnumStringValues Enumerates the set of values in String for AssetSourceSystemTypeEnum
func GetAssetSourceSystemTypeEnumStringValues() []string {
	return []string{
		"VCENTER",
		"AWS",
		"ORACLE_CDB",
		"ORACLE_PDB",
		"ORACLE_NON_CDB",
	}
}

// GetMappingAssetSourceSystemTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAssetSourceSystemTypeEnum(val string) (AssetSourceSystemTypeEnum, bool) {
	enum, ok := mappingAssetSourceSystemTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
