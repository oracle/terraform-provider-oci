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

// VmNetworkConsumerTypeEnum Enum with underlying type: string
type VmNetworkConsumerTypeEnum string

// Set of constants representing the allowable values for VmNetworkConsumerTypeEnum
const (
	VmNetworkConsumerTypeInstance VmNetworkConsumerTypeEnum = "INSTANCE"
	VmNetworkConsumerTypeCluster  VmNetworkConsumerTypeEnum = "CLUSTER"
)

var mappingVmNetworkConsumerTypeEnum = map[string]VmNetworkConsumerTypeEnum{
	"INSTANCE": VmNetworkConsumerTypeInstance,
	"CLUSTER":  VmNetworkConsumerTypeCluster,
}

var mappingVmNetworkConsumerTypeEnumLowerCase = map[string]VmNetworkConsumerTypeEnum{
	"instance": VmNetworkConsumerTypeInstance,
	"cluster":  VmNetworkConsumerTypeCluster,
}

// GetVmNetworkConsumerTypeEnumValues Enumerates the set of values for VmNetworkConsumerTypeEnum
func GetVmNetworkConsumerTypeEnumValues() []VmNetworkConsumerTypeEnum {
	values := make([]VmNetworkConsumerTypeEnum, 0)
	for _, v := range mappingVmNetworkConsumerTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetVmNetworkConsumerTypeEnumStringValues Enumerates the set of values in String for VmNetworkConsumerTypeEnum
func GetVmNetworkConsumerTypeEnumStringValues() []string {
	return []string{
		"INSTANCE",
		"CLUSTER",
	}
}

// GetMappingVmNetworkConsumerTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVmNetworkConsumerTypeEnum(val string) (VmNetworkConsumerTypeEnum, bool) {
	enum, ok := mappingVmNetworkConsumerTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
