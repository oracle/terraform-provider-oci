// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Globally Distributed Database
//
// Use the Globally Distributed Database service APIs to create and manage the Globally distributed databases.
//

package distributeddatabase

import (
	"strings"
)

// BaseDbDiskRedundancyEnum Enum with underlying type: string
type BaseDbDiskRedundancyEnum string

// Set of constants representing the allowable values for BaseDbDiskRedundancyEnum
const (
	BaseDbDiskRedundancyHigh   BaseDbDiskRedundancyEnum = "HIGH"
	BaseDbDiskRedundancyNormal BaseDbDiskRedundancyEnum = "NORMAL"
)

var mappingBaseDbDiskRedundancyEnum = map[string]BaseDbDiskRedundancyEnum{
	"HIGH":   BaseDbDiskRedundancyHigh,
	"NORMAL": BaseDbDiskRedundancyNormal,
}

var mappingBaseDbDiskRedundancyEnumLowerCase = map[string]BaseDbDiskRedundancyEnum{
	"high":   BaseDbDiskRedundancyHigh,
	"normal": BaseDbDiskRedundancyNormal,
}

// GetBaseDbDiskRedundancyEnumValues Enumerates the set of values for BaseDbDiskRedundancyEnum
func GetBaseDbDiskRedundancyEnumValues() []BaseDbDiskRedundancyEnum {
	values := make([]BaseDbDiskRedundancyEnum, 0)
	for _, v := range mappingBaseDbDiskRedundancyEnum {
		values = append(values, v)
	}
	return values
}

// GetBaseDbDiskRedundancyEnumStringValues Enumerates the set of values in String for BaseDbDiskRedundancyEnum
func GetBaseDbDiskRedundancyEnumStringValues() []string {
	return []string{
		"HIGH",
		"NORMAL",
	}
}

// GetMappingBaseDbDiskRedundancyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBaseDbDiskRedundancyEnum(val string) (BaseDbDiskRedundancyEnum, bool) {
	enum, ok := mappingBaseDbDiskRedundancyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
