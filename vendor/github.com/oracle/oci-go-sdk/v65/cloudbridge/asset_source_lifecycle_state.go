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

// AssetSourceLifecycleStateEnum Enum with underlying type: string
type AssetSourceLifecycleStateEnum string

// Set of constants representing the allowable values for AssetSourceLifecycleStateEnum
const (
	AssetSourceLifecycleStateCreating       AssetSourceLifecycleStateEnum = "CREATING"
	AssetSourceLifecycleStateActive         AssetSourceLifecycleStateEnum = "ACTIVE"
	AssetSourceLifecycleStateDeleting       AssetSourceLifecycleStateEnum = "DELETING"
	AssetSourceLifecycleStateDeleted        AssetSourceLifecycleStateEnum = "DELETED"
	AssetSourceLifecycleStateFailed         AssetSourceLifecycleStateEnum = "FAILED"
	AssetSourceLifecycleStateUpdating       AssetSourceLifecycleStateEnum = "UPDATING"
	AssetSourceLifecycleStateNeedsAttention AssetSourceLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingAssetSourceLifecycleStateEnum = map[string]AssetSourceLifecycleStateEnum{
	"CREATING":        AssetSourceLifecycleStateCreating,
	"ACTIVE":          AssetSourceLifecycleStateActive,
	"DELETING":        AssetSourceLifecycleStateDeleting,
	"DELETED":         AssetSourceLifecycleStateDeleted,
	"FAILED":          AssetSourceLifecycleStateFailed,
	"UPDATING":        AssetSourceLifecycleStateUpdating,
	"NEEDS_ATTENTION": AssetSourceLifecycleStateNeedsAttention,
}

var mappingAssetSourceLifecycleStateEnumLowerCase = map[string]AssetSourceLifecycleStateEnum{
	"creating":        AssetSourceLifecycleStateCreating,
	"active":          AssetSourceLifecycleStateActive,
	"deleting":        AssetSourceLifecycleStateDeleting,
	"deleted":         AssetSourceLifecycleStateDeleted,
	"failed":          AssetSourceLifecycleStateFailed,
	"updating":        AssetSourceLifecycleStateUpdating,
	"needs_attention": AssetSourceLifecycleStateNeedsAttention,
}

// GetAssetSourceLifecycleStateEnumValues Enumerates the set of values for AssetSourceLifecycleStateEnum
func GetAssetSourceLifecycleStateEnumValues() []AssetSourceLifecycleStateEnum {
	values := make([]AssetSourceLifecycleStateEnum, 0)
	for _, v := range mappingAssetSourceLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAssetSourceLifecycleStateEnumStringValues Enumerates the set of values in String for AssetSourceLifecycleStateEnum
func GetAssetSourceLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
		"UPDATING",
		"NEEDS_ATTENTION",
	}
}

// GetMappingAssetSourceLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAssetSourceLifecycleStateEnum(val string) (AssetSourceLifecycleStateEnum, bool) {
	enum, ok := mappingAssetSourceLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
