// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Migrations API
//
// A description of the Oracle Cloud Migrations API.
//

package cloudmigrations

import (
	"strings"
)

// AssetSourceCredentialsTypeEnum Enum with underlying type: string
type AssetSourceCredentialsTypeEnum string

// Set of constants representing the allowable values for AssetSourceCredentialsTypeEnum
const (
	AssetSourceCredentialsTypeBasic AssetSourceCredentialsTypeEnum = "BASIC"
)

var mappingAssetSourceCredentialsTypeEnum = map[string]AssetSourceCredentialsTypeEnum{
	"BASIC": AssetSourceCredentialsTypeBasic,
}

var mappingAssetSourceCredentialsTypeEnumLowerCase = map[string]AssetSourceCredentialsTypeEnum{
	"basic": AssetSourceCredentialsTypeBasic,
}

// GetAssetSourceCredentialsTypeEnumValues Enumerates the set of values for AssetSourceCredentialsTypeEnum
func GetAssetSourceCredentialsTypeEnumValues() []AssetSourceCredentialsTypeEnum {
	values := make([]AssetSourceCredentialsTypeEnum, 0)
	for _, v := range mappingAssetSourceCredentialsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAssetSourceCredentialsTypeEnumStringValues Enumerates the set of values in String for AssetSourceCredentialsTypeEnum
func GetAssetSourceCredentialsTypeEnumStringValues() []string {
	return []string{
		"BASIC",
	}
}

// GetMappingAssetSourceCredentialsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAssetSourceCredentialsTypeEnum(val string) (AssetSourceCredentialsTypeEnum, bool) {
	enum, ok := mappingAssetSourceCredentialsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
