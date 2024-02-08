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

// AssetTypeEnum Enum with underlying type: string
type AssetTypeEnum string

// Set of constants representing the allowable values for AssetTypeEnum
const (
	AssetTypeVmwareVm AssetTypeEnum = "VMWARE_VM"
	AssetTypeVm       AssetTypeEnum = "VM"
)

var mappingAssetTypeEnum = map[string]AssetTypeEnum{
	"VMWARE_VM": AssetTypeVmwareVm,
	"VM":        AssetTypeVm,
}

var mappingAssetTypeEnumLowerCase = map[string]AssetTypeEnum{
	"vmware_vm": AssetTypeVmwareVm,
	"vm":        AssetTypeVm,
}

// GetAssetTypeEnumValues Enumerates the set of values for AssetTypeEnum
func GetAssetTypeEnumValues() []AssetTypeEnum {
	values := make([]AssetTypeEnum, 0)
	for _, v := range mappingAssetTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAssetTypeEnumStringValues Enumerates the set of values in String for AssetTypeEnum
func GetAssetTypeEnumStringValues() []string {
	return []string{
		"VMWARE_VM",
		"VM",
	}
}

// GetMappingAssetTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAssetTypeEnum(val string) (AssetTypeEnum, bool) {
	enum, ok := mappingAssetTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
