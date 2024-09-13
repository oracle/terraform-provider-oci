// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"strings"
)

// ManagedEntityEnum Enum with underlying type: string
type ManagedEntityEnum string

// Set of constants representing the allowable values for ManagedEntityEnum
const (
	ManagedEntityResource ManagedEntityEnum = "RESOURCE"
	ManagedEntityTarget   ManagedEntityEnum = "TARGET"
)

var mappingManagedEntityEnum = map[string]ManagedEntityEnum{
	"RESOURCE": ManagedEntityResource,
	"TARGET":   ManagedEntityTarget,
}

var mappingManagedEntityEnumLowerCase = map[string]ManagedEntityEnum{
	"resource": ManagedEntityResource,
	"target":   ManagedEntityTarget,
}

// GetManagedEntityEnumValues Enumerates the set of values for ManagedEntityEnum
func GetManagedEntityEnumValues() []ManagedEntityEnum {
	values := make([]ManagedEntityEnum, 0)
	for _, v := range mappingManagedEntityEnum {
		values = append(values, v)
	}
	return values
}

// GetManagedEntityEnumStringValues Enumerates the set of values in String for ManagedEntityEnum
func GetManagedEntityEnumStringValues() []string {
	return []string{
		"RESOURCE",
		"TARGET",
	}
}

// GetMappingManagedEntityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingManagedEntityEnum(val string) (ManagedEntityEnum, bool) {
	enum, ok := mappingManagedEntityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
