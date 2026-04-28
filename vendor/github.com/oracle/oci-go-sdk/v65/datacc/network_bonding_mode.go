// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Infrastructure Cloud@Customer Service API
//
// API for Database Infrastructure Cloud@Customer Service. Use this API to manage Database Infrastructure VM clusters, Application VMs, and related resources.
//

package datacc

import (
	"strings"
)

// NetworkBondingModeEnum Enum with underlying type: string
type NetworkBondingModeEnum string

// Set of constants representing the allowable values for NetworkBondingModeEnum
const (
	NetworkBondingModeActiveBackup NetworkBondingModeEnum = "ACTIVE_BACKUP"
	NetworkBondingModeLacp         NetworkBondingModeEnum = "LACP"
)

var mappingNetworkBondingModeEnum = map[string]NetworkBondingModeEnum{
	"ACTIVE_BACKUP": NetworkBondingModeActiveBackup,
	"LACP":          NetworkBondingModeLacp,
}

var mappingNetworkBondingModeEnumLowerCase = map[string]NetworkBondingModeEnum{
	"active_backup": NetworkBondingModeActiveBackup,
	"lacp":          NetworkBondingModeLacp,
}

// GetNetworkBondingModeEnumValues Enumerates the set of values for NetworkBondingModeEnum
func GetNetworkBondingModeEnumValues() []NetworkBondingModeEnum {
	values := make([]NetworkBondingModeEnum, 0)
	for _, v := range mappingNetworkBondingModeEnum {
		values = append(values, v)
	}
	return values
}

// GetNetworkBondingModeEnumStringValues Enumerates the set of values in String for NetworkBondingModeEnum
func GetNetworkBondingModeEnumStringValues() []string {
	return []string{
		"ACTIVE_BACKUP",
		"LACP",
	}
}

// GetMappingNetworkBondingModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNetworkBondingModeEnum(val string) (NetworkBondingModeEnum, bool) {
	enum, ok := mappingNetworkBondingModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
