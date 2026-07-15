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

// NetworkBondingInterfaceEnumEnum Enum with underlying type: string
type NetworkBondingInterfaceEnumEnum string

// Set of constants representing the allowable values for NetworkBondingInterfaceEnumEnum
const (
	NetworkBondingInterfaceEnumBtbond1 NetworkBondingInterfaceEnumEnum = "BTBOND1"
	NetworkBondingInterfaceEnumBtbond2 NetworkBondingInterfaceEnumEnum = "BTBOND2"
	NetworkBondingInterfaceEnumBtbond3 NetworkBondingInterfaceEnumEnum = "BTBOND3"
	NetworkBondingInterfaceEnumBtbond4 NetworkBondingInterfaceEnumEnum = "BTBOND4"
	NetworkBondingInterfaceEnumBtbond5 NetworkBondingInterfaceEnumEnum = "BTBOND5"
	NetworkBondingInterfaceEnumBtbond6 NetworkBondingInterfaceEnumEnum = "BTBOND6"
)

var mappingNetworkBondingInterfaceEnumEnum = map[string]NetworkBondingInterfaceEnumEnum{
	"BTBOND1": NetworkBondingInterfaceEnumBtbond1,
	"BTBOND2": NetworkBondingInterfaceEnumBtbond2,
	"BTBOND3": NetworkBondingInterfaceEnumBtbond3,
	"BTBOND4": NetworkBondingInterfaceEnumBtbond4,
	"BTBOND5": NetworkBondingInterfaceEnumBtbond5,
	"BTBOND6": NetworkBondingInterfaceEnumBtbond6,
}

var mappingNetworkBondingInterfaceEnumEnumLowerCase = map[string]NetworkBondingInterfaceEnumEnum{
	"btbond1": NetworkBondingInterfaceEnumBtbond1,
	"btbond2": NetworkBondingInterfaceEnumBtbond2,
	"btbond3": NetworkBondingInterfaceEnumBtbond3,
	"btbond4": NetworkBondingInterfaceEnumBtbond4,
	"btbond5": NetworkBondingInterfaceEnumBtbond5,
	"btbond6": NetworkBondingInterfaceEnumBtbond6,
}

// GetNetworkBondingInterfaceEnumEnumValues Enumerates the set of values for NetworkBondingInterfaceEnumEnum
func GetNetworkBondingInterfaceEnumEnumValues() []NetworkBondingInterfaceEnumEnum {
	values := make([]NetworkBondingInterfaceEnumEnum, 0)
	for _, v := range mappingNetworkBondingInterfaceEnumEnum {
		values = append(values, v)
	}
	return values
}

// GetNetworkBondingInterfaceEnumEnumStringValues Enumerates the set of values in String for NetworkBondingInterfaceEnumEnum
func GetNetworkBondingInterfaceEnumEnumStringValues() []string {
	return []string{
		"BTBOND1",
		"BTBOND2",
		"BTBOND3",
		"BTBOND4",
		"BTBOND5",
		"BTBOND6",
	}
}

// GetMappingNetworkBondingInterfaceEnumEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNetworkBondingInterfaceEnumEnum(val string) (NetworkBondingInterfaceEnumEnum, bool) {
	enum, ok := mappingNetworkBondingInterfaceEnumEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
