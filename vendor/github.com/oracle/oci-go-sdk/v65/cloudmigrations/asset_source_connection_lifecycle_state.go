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

// AssetSourceConnectionLifecycleStateEnum Enum with underlying type: string
type AssetSourceConnectionLifecycleStateEnum string

// Set of constants representing the allowable values for AssetSourceConnectionLifecycleStateEnum
const (
	AssetSourceConnectionLifecycleStateActive         AssetSourceConnectionLifecycleStateEnum = "ACTIVE"
	AssetSourceConnectionLifecycleStateUpdating       AssetSourceConnectionLifecycleStateEnum = "UPDATING"
	AssetSourceConnectionLifecycleStateNeedsAttention AssetSourceConnectionLifecycleStateEnum = "NEEDS_ATTENTION"
	AssetSourceConnectionLifecycleStateDeleted        AssetSourceConnectionLifecycleStateEnum = "DELETED"
	AssetSourceConnectionLifecycleStateCreating       AssetSourceConnectionLifecycleStateEnum = "CREATING"
)

var mappingAssetSourceConnectionLifecycleStateEnum = map[string]AssetSourceConnectionLifecycleStateEnum{
	"ACTIVE":          AssetSourceConnectionLifecycleStateActive,
	"UPDATING":        AssetSourceConnectionLifecycleStateUpdating,
	"NEEDS_ATTENTION": AssetSourceConnectionLifecycleStateNeedsAttention,
	"DELETED":         AssetSourceConnectionLifecycleStateDeleted,
	"CREATING":        AssetSourceConnectionLifecycleStateCreating,
}

var mappingAssetSourceConnectionLifecycleStateEnumLowerCase = map[string]AssetSourceConnectionLifecycleStateEnum{
	"active":          AssetSourceConnectionLifecycleStateActive,
	"updating":        AssetSourceConnectionLifecycleStateUpdating,
	"needs_attention": AssetSourceConnectionLifecycleStateNeedsAttention,
	"deleted":         AssetSourceConnectionLifecycleStateDeleted,
	"creating":        AssetSourceConnectionLifecycleStateCreating,
}

// GetAssetSourceConnectionLifecycleStateEnumValues Enumerates the set of values for AssetSourceConnectionLifecycleStateEnum
func GetAssetSourceConnectionLifecycleStateEnumValues() []AssetSourceConnectionLifecycleStateEnum {
	values := make([]AssetSourceConnectionLifecycleStateEnum, 0)
	for _, v := range mappingAssetSourceConnectionLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAssetSourceConnectionLifecycleStateEnumStringValues Enumerates the set of values in String for AssetSourceConnectionLifecycleStateEnum
func GetAssetSourceConnectionLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"UPDATING",
		"NEEDS_ATTENTION",
		"DELETED",
		"CREATING",
	}
}

// GetMappingAssetSourceConnectionLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAssetSourceConnectionLifecycleStateEnum(val string) (AssetSourceConnectionLifecycleStateEnum, bool) {
	enum, ok := mappingAssetSourceConnectionLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
