// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
// For more information, see Data Catalog (https://docs.oracle.com/iaas/data-catalog/home.htm).
//

package datacatalog

import (
	"strings"
)

// LineageDirectionEnum Enum with underlying type: string
type LineageDirectionEnum string

// Set of constants representing the allowable values for LineageDirectionEnum
const (
	LineageDirectionUpstream   LineageDirectionEnum = "UPSTREAM"
	LineageDirectionBoth       LineageDirectionEnum = "BOTH"
	LineageDirectionDownstream LineageDirectionEnum = "DOWNSTREAM"
)

var mappingLineageDirectionEnum = map[string]LineageDirectionEnum{
	"UPSTREAM":   LineageDirectionUpstream,
	"BOTH":       LineageDirectionBoth,
	"DOWNSTREAM": LineageDirectionDownstream,
}

var mappingLineageDirectionEnumLowerCase = map[string]LineageDirectionEnum{
	"upstream":   LineageDirectionUpstream,
	"both":       LineageDirectionBoth,
	"downstream": LineageDirectionDownstream,
}

// GetLineageDirectionEnumValues Enumerates the set of values for LineageDirectionEnum
func GetLineageDirectionEnumValues() []LineageDirectionEnum {
	values := make([]LineageDirectionEnum, 0)
	for _, v := range mappingLineageDirectionEnum {
		values = append(values, v)
	}
	return values
}

// GetLineageDirectionEnumStringValues Enumerates the set of values in String for LineageDirectionEnum
func GetLineageDirectionEnumStringValues() []string {
	return []string{
		"UPSTREAM",
		"BOTH",
		"DOWNSTREAM",
	}
}

// GetMappingLineageDirectionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLineageDirectionEnum(val string) (LineageDirectionEnum, bool) {
	enum, ok := mappingLineageDirectionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
