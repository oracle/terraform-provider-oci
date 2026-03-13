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

// BaseDbEditionEnum Enum with underlying type: string
type BaseDbEditionEnum string

// Set of constants representing the allowable values for BaseDbEditionEnum
const (
	BaseDbEditionStandardEdition                     BaseDbEditionEnum = "STANDARD_EDITION"
	BaseDbEditionEnterpriseEdition                   BaseDbEditionEnum = "ENTERPRISE_EDITION"
	BaseDbEditionEnterpriseEditionHighPerformance    BaseDbEditionEnum = "ENTERPRISE_EDITION_HIGH_PERFORMANCE"
	BaseDbEditionEnterpriseEditionExtremePerformance BaseDbEditionEnum = "ENTERPRISE_EDITION_EXTREME_PERFORMANCE"
)

var mappingBaseDbEditionEnum = map[string]BaseDbEditionEnum{
	"STANDARD_EDITION":                       BaseDbEditionStandardEdition,
	"ENTERPRISE_EDITION":                     BaseDbEditionEnterpriseEdition,
	"ENTERPRISE_EDITION_HIGH_PERFORMANCE":    BaseDbEditionEnterpriseEditionHighPerformance,
	"ENTERPRISE_EDITION_EXTREME_PERFORMANCE": BaseDbEditionEnterpriseEditionExtremePerformance,
}

var mappingBaseDbEditionEnumLowerCase = map[string]BaseDbEditionEnum{
	"standard_edition":                       BaseDbEditionStandardEdition,
	"enterprise_edition":                     BaseDbEditionEnterpriseEdition,
	"enterprise_edition_high_performance":    BaseDbEditionEnterpriseEditionHighPerformance,
	"enterprise_edition_extreme_performance": BaseDbEditionEnterpriseEditionExtremePerformance,
}

// GetBaseDbEditionEnumValues Enumerates the set of values for BaseDbEditionEnum
func GetBaseDbEditionEnumValues() []BaseDbEditionEnum {
	values := make([]BaseDbEditionEnum, 0)
	for _, v := range mappingBaseDbEditionEnum {
		values = append(values, v)
	}
	return values
}

// GetBaseDbEditionEnumStringValues Enumerates the set of values in String for BaseDbEditionEnum
func GetBaseDbEditionEnumStringValues() []string {
	return []string{
		"STANDARD_EDITION",
		"ENTERPRISE_EDITION",
		"ENTERPRISE_EDITION_HIGH_PERFORMANCE",
		"ENTERPRISE_EDITION_EXTREME_PERFORMANCE",
	}
}

// GetMappingBaseDbEditionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBaseDbEditionEnum(val string) (BaseDbEditionEnum, bool) {
	enum, ok := mappingBaseDbEditionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
