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

// AssetSourceConnectionTypeEnum Enum with underlying type: string
type AssetSourceConnectionTypeEnum string

// Set of constants representing the allowable values for AssetSourceConnectionTypeEnum
const (
	AssetSourceConnectionTypeDiscovery   AssetSourceConnectionTypeEnum = "DISCOVERY"
	AssetSourceConnectionTypeReplication AssetSourceConnectionTypeEnum = "REPLICATION"
)

var mappingAssetSourceConnectionTypeEnum = map[string]AssetSourceConnectionTypeEnum{
	"DISCOVERY":   AssetSourceConnectionTypeDiscovery,
	"REPLICATION": AssetSourceConnectionTypeReplication,
}

var mappingAssetSourceConnectionTypeEnumLowerCase = map[string]AssetSourceConnectionTypeEnum{
	"discovery":   AssetSourceConnectionTypeDiscovery,
	"replication": AssetSourceConnectionTypeReplication,
}

// GetAssetSourceConnectionTypeEnumValues Enumerates the set of values for AssetSourceConnectionTypeEnum
func GetAssetSourceConnectionTypeEnumValues() []AssetSourceConnectionTypeEnum {
	values := make([]AssetSourceConnectionTypeEnum, 0)
	for _, v := range mappingAssetSourceConnectionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAssetSourceConnectionTypeEnumStringValues Enumerates the set of values in String for AssetSourceConnectionTypeEnum
func GetAssetSourceConnectionTypeEnumStringValues() []string {
	return []string{
		"DISCOVERY",
		"REPLICATION",
	}
}

// GetMappingAssetSourceConnectionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAssetSourceConnectionTypeEnum(val string) (AssetSourceConnectionTypeEnum, bool) {
	enum, ok := mappingAssetSourceConnectionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
