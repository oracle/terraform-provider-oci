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

// VmClusterNetworkTypeEnum Enum with underlying type: string
type VmClusterNetworkTypeEnum string

// Set of constants representing the allowable values for VmClusterNetworkTypeEnum
const (
	VmClusterNetworkTypeClient VmClusterNetworkTypeEnum = "CLIENT"
	VmClusterNetworkTypeBackup VmClusterNetworkTypeEnum = "BACKUP"
)

var mappingVmClusterNetworkTypeEnum = map[string]VmClusterNetworkTypeEnum{
	"CLIENT": VmClusterNetworkTypeClient,
	"BACKUP": VmClusterNetworkTypeBackup,
}

var mappingVmClusterNetworkTypeEnumLowerCase = map[string]VmClusterNetworkTypeEnum{
	"client": VmClusterNetworkTypeClient,
	"backup": VmClusterNetworkTypeBackup,
}

// GetVmClusterNetworkTypeEnumValues Enumerates the set of values for VmClusterNetworkTypeEnum
func GetVmClusterNetworkTypeEnumValues() []VmClusterNetworkTypeEnum {
	values := make([]VmClusterNetworkTypeEnum, 0)
	for _, v := range mappingVmClusterNetworkTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetVmClusterNetworkTypeEnumStringValues Enumerates the set of values in String for VmClusterNetworkTypeEnum
func GetVmClusterNetworkTypeEnumStringValues() []string {
	return []string{
		"CLIENT",
		"BACKUP",
	}
}

// GetMappingVmClusterNetworkTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVmClusterNetworkTypeEnum(val string) (VmClusterNetworkTypeEnum, bool) {
	enum, ok := mappingVmClusterNetworkTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
