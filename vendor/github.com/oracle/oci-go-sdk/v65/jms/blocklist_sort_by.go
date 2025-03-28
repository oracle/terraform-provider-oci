// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Fleets API
//
// The APIs for the Fleet Management (https://docs.oracle.com/en-us/iaas/jms/doc/fleet-management.html) feature of Java Management Service to monitor and manage the usage of Java in your enterprise. Use these APIs to manage fleets, configure managed instances to report to fleets, and gain insights into the Java workloads running on these instances by carrying out basic and advanced features.
//

package jms

import (
	"strings"
)

// BlocklistSortByEnum Enum with underlying type: string
type BlocklistSortByEnum string

// Set of constants representing the allowable values for BlocklistSortByEnum
const (
	BlocklistSortByOperation BlocklistSortByEnum = "operation"
)

var mappingBlocklistSortByEnum = map[string]BlocklistSortByEnum{
	"operation": BlocklistSortByOperation,
}

var mappingBlocklistSortByEnumLowerCase = map[string]BlocklistSortByEnum{
	"operation": BlocklistSortByOperation,
}

// GetBlocklistSortByEnumValues Enumerates the set of values for BlocklistSortByEnum
func GetBlocklistSortByEnumValues() []BlocklistSortByEnum {
	values := make([]BlocklistSortByEnum, 0)
	for _, v := range mappingBlocklistSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetBlocklistSortByEnumStringValues Enumerates the set of values in String for BlocklistSortByEnum
func GetBlocklistSortByEnumStringValues() []string {
	return []string{
		"operation",
	}
}

// GetMappingBlocklistSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBlocklistSortByEnum(val string) (BlocklistSortByEnum, bool) {
	enum, ok := mappingBlocklistSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
