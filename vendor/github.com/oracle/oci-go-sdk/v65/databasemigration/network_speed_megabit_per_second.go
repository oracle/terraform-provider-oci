// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"strings"
)

// NetworkSpeedMegabitPerSecondEnum Enum with underlying type: string
type NetworkSpeedMegabitPerSecondEnum string

// Set of constants representing the allowable values for NetworkSpeedMegabitPerSecondEnum
const (
	NetworkSpeedMegabitPerSecondMbps10    NetworkSpeedMegabitPerSecondEnum = "MBPS_10"
	NetworkSpeedMegabitPerSecondMbps100   NetworkSpeedMegabitPerSecondEnum = "MBPS_100"
	NetworkSpeedMegabitPerSecondMbps1000  NetworkSpeedMegabitPerSecondEnum = "MBPS_1000"
	NetworkSpeedMegabitPerSecondMbps2500  NetworkSpeedMegabitPerSecondEnum = "MBPS_2500"
	NetworkSpeedMegabitPerSecondMbps5000  NetworkSpeedMegabitPerSecondEnum = "MBPS_5000"
	NetworkSpeedMegabitPerSecondMbps10000 NetworkSpeedMegabitPerSecondEnum = "MBPS_10000"
)

var mappingNetworkSpeedMegabitPerSecondEnum = map[string]NetworkSpeedMegabitPerSecondEnum{
	"MBPS_10":    NetworkSpeedMegabitPerSecondMbps10,
	"MBPS_100":   NetworkSpeedMegabitPerSecondMbps100,
	"MBPS_1000":  NetworkSpeedMegabitPerSecondMbps1000,
	"MBPS_2500":  NetworkSpeedMegabitPerSecondMbps2500,
	"MBPS_5000":  NetworkSpeedMegabitPerSecondMbps5000,
	"MBPS_10000": NetworkSpeedMegabitPerSecondMbps10000,
}

var mappingNetworkSpeedMegabitPerSecondEnumLowerCase = map[string]NetworkSpeedMegabitPerSecondEnum{
	"mbps_10":    NetworkSpeedMegabitPerSecondMbps10,
	"mbps_100":   NetworkSpeedMegabitPerSecondMbps100,
	"mbps_1000":  NetworkSpeedMegabitPerSecondMbps1000,
	"mbps_2500":  NetworkSpeedMegabitPerSecondMbps2500,
	"mbps_5000":  NetworkSpeedMegabitPerSecondMbps5000,
	"mbps_10000": NetworkSpeedMegabitPerSecondMbps10000,
}

// GetNetworkSpeedMegabitPerSecondEnumValues Enumerates the set of values for NetworkSpeedMegabitPerSecondEnum
func GetNetworkSpeedMegabitPerSecondEnumValues() []NetworkSpeedMegabitPerSecondEnum {
	values := make([]NetworkSpeedMegabitPerSecondEnum, 0)
	for _, v := range mappingNetworkSpeedMegabitPerSecondEnum {
		values = append(values, v)
	}
	return values
}

// GetNetworkSpeedMegabitPerSecondEnumStringValues Enumerates the set of values in String for NetworkSpeedMegabitPerSecondEnum
func GetNetworkSpeedMegabitPerSecondEnumStringValues() []string {
	return []string{
		"MBPS_10",
		"MBPS_100",
		"MBPS_1000",
		"MBPS_2500",
		"MBPS_5000",
		"MBPS_10000",
	}
}

// GetMappingNetworkSpeedMegabitPerSecondEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNetworkSpeedMegabitPerSecondEnum(val string) (NetworkSpeedMegabitPerSecondEnum, bool) {
	enum, ok := mappingNetworkSpeedMegabitPerSecondEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
