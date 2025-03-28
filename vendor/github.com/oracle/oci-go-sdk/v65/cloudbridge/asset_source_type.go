// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// AssetSourceTypeEnum Enum with underlying type: string
type AssetSourceTypeEnum string

// Set of constants representing the allowable values for AssetSourceTypeEnum
const (
	AssetSourceTypeVmware AssetSourceTypeEnum = "VMWARE"
	AssetSourceTypeAws    AssetSourceTypeEnum = "AWS"
)

var mappingAssetSourceTypeEnum = map[string]AssetSourceTypeEnum{
	"VMWARE": AssetSourceTypeVmware,
	"AWS":    AssetSourceTypeAws,
}

var mappingAssetSourceTypeEnumLowerCase = map[string]AssetSourceTypeEnum{
	"vmware": AssetSourceTypeVmware,
	"aws":    AssetSourceTypeAws,
}

// GetAssetSourceTypeEnumValues Enumerates the set of values for AssetSourceTypeEnum
func GetAssetSourceTypeEnumValues() []AssetSourceTypeEnum {
	values := make([]AssetSourceTypeEnum, 0)
	for _, v := range mappingAssetSourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAssetSourceTypeEnumStringValues Enumerates the set of values in String for AssetSourceTypeEnum
func GetAssetSourceTypeEnumStringValues() []string {
	return []string{
		"VMWARE",
		"AWS",
	}
}

// GetMappingAssetSourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAssetSourceTypeEnum(val string) (AssetSourceTypeEnum, bool) {
	enum, ok := mappingAssetSourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
