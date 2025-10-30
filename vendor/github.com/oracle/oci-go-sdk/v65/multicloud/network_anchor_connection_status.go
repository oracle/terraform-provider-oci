// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Multicloud API
//
// Use the Oracle Multicloud API to retrieve resource anchors and network anchors, and the metadata mappings related a Cloud Service Provider. For more information, see <link to docs>.
//

package multicloud

import (
	"strings"
)

// NetworkAnchorConnectionStatusEnum Enum with underlying type: string
type NetworkAnchorConnectionStatusEnum string

// Set of constants representing the allowable values for NetworkAnchorConnectionStatusEnum
const (
	NetworkAnchorConnectionStatusConnected      NetworkAnchorConnectionStatusEnum = "CONNECTED"
	NetworkAnchorConnectionStatusDisconnected   NetworkAnchorConnectionStatusEnum = "DISCONNECTED"
	NetworkAnchorConnectionStatusConnecting     NetworkAnchorConnectionStatusEnum = "CONNECTING"
	NetworkAnchorConnectionStatusActive         NetworkAnchorConnectionStatusEnum = "ACTIVE"
	NetworkAnchorConnectionStatusError          NetworkAnchorConnectionStatusEnum = "ERROR"
	NetworkAnchorConnectionStatusUpdating       NetworkAnchorConnectionStatusEnum = "UPDATING"
	NetworkAnchorConnectionStatusNeedsAttention NetworkAnchorConnectionStatusEnum = "NEEDS_ATTENTION"
	NetworkAnchorConnectionStatusFailed         NetworkAnchorConnectionStatusEnum = "FAILED"
	NetworkAnchorConnectionStatusDeleting       NetworkAnchorConnectionStatusEnum = "DELETING"
	NetworkAnchorConnectionStatusDeleted        NetworkAnchorConnectionStatusEnum = "DELETED"
)

var mappingNetworkAnchorConnectionStatusEnum = map[string]NetworkAnchorConnectionStatusEnum{
	"CONNECTED":       NetworkAnchorConnectionStatusConnected,
	"DISCONNECTED":    NetworkAnchorConnectionStatusDisconnected,
	"CONNECTING":      NetworkAnchorConnectionStatusConnecting,
	"ACTIVE":          NetworkAnchorConnectionStatusActive,
	"ERROR":           NetworkAnchorConnectionStatusError,
	"UPDATING":        NetworkAnchorConnectionStatusUpdating,
	"NEEDS_ATTENTION": NetworkAnchorConnectionStatusNeedsAttention,
	"FAILED":          NetworkAnchorConnectionStatusFailed,
	"DELETING":        NetworkAnchorConnectionStatusDeleting,
	"DELETED":         NetworkAnchorConnectionStatusDeleted,
}

var mappingNetworkAnchorConnectionStatusEnumLowerCase = map[string]NetworkAnchorConnectionStatusEnum{
	"connected":       NetworkAnchorConnectionStatusConnected,
	"disconnected":    NetworkAnchorConnectionStatusDisconnected,
	"connecting":      NetworkAnchorConnectionStatusConnecting,
	"active":          NetworkAnchorConnectionStatusActive,
	"error":           NetworkAnchorConnectionStatusError,
	"updating":        NetworkAnchorConnectionStatusUpdating,
	"needs_attention": NetworkAnchorConnectionStatusNeedsAttention,
	"failed":          NetworkAnchorConnectionStatusFailed,
	"deleting":        NetworkAnchorConnectionStatusDeleting,
	"deleted":         NetworkAnchorConnectionStatusDeleted,
}

// GetNetworkAnchorConnectionStatusEnumValues Enumerates the set of values for NetworkAnchorConnectionStatusEnum
func GetNetworkAnchorConnectionStatusEnumValues() []NetworkAnchorConnectionStatusEnum {
	values := make([]NetworkAnchorConnectionStatusEnum, 0)
	for _, v := range mappingNetworkAnchorConnectionStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetNetworkAnchorConnectionStatusEnumStringValues Enumerates the set of values in String for NetworkAnchorConnectionStatusEnum
func GetNetworkAnchorConnectionStatusEnumStringValues() []string {
	return []string{
		"CONNECTED",
		"DISCONNECTED",
		"CONNECTING",
		"ACTIVE",
		"ERROR",
		"UPDATING",
		"NEEDS_ATTENTION",
		"FAILED",
		"DELETING",
		"DELETED",
	}
}

// GetMappingNetworkAnchorConnectionStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNetworkAnchorConnectionStatusEnum(val string) (NetworkAnchorConnectionStatusEnum, bool) {
	enum, ok := mappingNetworkAnchorConnectionStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
