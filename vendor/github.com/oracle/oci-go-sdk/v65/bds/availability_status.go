// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// REST API for Oracle Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service clusters. Build on Hadoop, Spark and Data Science distributions, which can be fully integrated with existing enterprise data in Oracle Database and Oracle applications.
//

package bds

import (
	"strings"
)

// AvailabilityStatusEnum Enum with underlying type: string
type AvailabilityStatusEnum string

// Set of constants representing the allowable values for AvailabilityStatusEnum
const (
	AvailabilityStatusOutOfHostCapacity    AvailabilityStatusEnum = "OUT_OF_HOST_CAPACITY"
	AvailabilityStatusHardwareNotSupported AvailabilityStatusEnum = "HARDWARE_NOT_SUPPORTED"
	AvailabilityStatusAvailable            AvailabilityStatusEnum = "AVAILABLE"
)

var mappingAvailabilityStatusEnum = map[string]AvailabilityStatusEnum{
	"OUT_OF_HOST_CAPACITY":   AvailabilityStatusOutOfHostCapacity,
	"HARDWARE_NOT_SUPPORTED": AvailabilityStatusHardwareNotSupported,
	"AVAILABLE":              AvailabilityStatusAvailable,
}

var mappingAvailabilityStatusEnumLowerCase = map[string]AvailabilityStatusEnum{
	"out_of_host_capacity":   AvailabilityStatusOutOfHostCapacity,
	"hardware_not_supported": AvailabilityStatusHardwareNotSupported,
	"available":              AvailabilityStatusAvailable,
}

// GetAvailabilityStatusEnumValues Enumerates the set of values for AvailabilityStatusEnum
func GetAvailabilityStatusEnumValues() []AvailabilityStatusEnum {
	values := make([]AvailabilityStatusEnum, 0)
	for _, v := range mappingAvailabilityStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetAvailabilityStatusEnumStringValues Enumerates the set of values in String for AvailabilityStatusEnum
func GetAvailabilityStatusEnumStringValues() []string {
	return []string{
		"OUT_OF_HOST_CAPACITY",
		"HARDWARE_NOT_SUPPORTED",
		"AVAILABLE",
	}
}

// GetMappingAvailabilityStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAvailabilityStatusEnum(val string) (AvailabilityStatusEnum, bool) {
	enum, ok := mappingAvailabilityStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
