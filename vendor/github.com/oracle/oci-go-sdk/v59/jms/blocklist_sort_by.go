// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service API
//
// API for the Java Management Service. Use this API to view, create, and manage Fleets.
//

package jms

// BlocklistSortByEnum Enum with underlying type: string
type BlocklistSortByEnum string

// Set of constants representing the allowable values for BlocklistSortByEnum
const (
	BlocklistSortByOperation BlocklistSortByEnum = "operation"
)

var mappingBlocklistSortByEnum = map[string]BlocklistSortByEnum{
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
