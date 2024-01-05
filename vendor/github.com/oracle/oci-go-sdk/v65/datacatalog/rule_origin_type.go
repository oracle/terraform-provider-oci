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

// RuleOriginTypeEnum Enum with underlying type: string
type RuleOriginTypeEnum string

// Set of constants representing the allowable values for RuleOriginTypeEnum
const (
	RuleOriginTypeSource    RuleOriginTypeEnum = "SOURCE"
	RuleOriginTypeUser      RuleOriginTypeEnum = "USER"
	RuleOriginTypeProfiling RuleOriginTypeEnum = "PROFILING"
)

var mappingRuleOriginTypeEnum = map[string]RuleOriginTypeEnum{
	"SOURCE":    RuleOriginTypeSource,
	"USER":      RuleOriginTypeUser,
	"PROFILING": RuleOriginTypeProfiling,
}

var mappingRuleOriginTypeEnumLowerCase = map[string]RuleOriginTypeEnum{
	"source":    RuleOriginTypeSource,
	"user":      RuleOriginTypeUser,
	"profiling": RuleOriginTypeProfiling,
}

// GetRuleOriginTypeEnumValues Enumerates the set of values for RuleOriginTypeEnum
func GetRuleOriginTypeEnumValues() []RuleOriginTypeEnum {
	values := make([]RuleOriginTypeEnum, 0)
	for _, v := range mappingRuleOriginTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRuleOriginTypeEnumStringValues Enumerates the set of values in String for RuleOriginTypeEnum
func GetRuleOriginTypeEnumStringValues() []string {
	return []string{
		"SOURCE",
		"USER",
		"PROFILING",
	}
}

// GetMappingRuleOriginTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRuleOriginTypeEnum(val string) (RuleOriginTypeEnum, bool) {
	enum, ok := mappingRuleOriginTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
