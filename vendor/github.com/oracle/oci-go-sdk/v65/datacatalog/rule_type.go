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

// RuleTypeEnum Enum with underlying type: string
type RuleTypeEnum string

// Set of constants representing the allowable values for RuleTypeEnum
const (
	RuleTypePrimarykey RuleTypeEnum = "PRIMARYKEY"
	RuleTypeForeignkey RuleTypeEnum = "FOREIGNKEY"
	RuleTypeUniquekey  RuleTypeEnum = "UNIQUEKEY"
)

var mappingRuleTypeEnum = map[string]RuleTypeEnum{
	"PRIMARYKEY": RuleTypePrimarykey,
	"FOREIGNKEY": RuleTypeForeignkey,
	"UNIQUEKEY":  RuleTypeUniquekey,
}

var mappingRuleTypeEnumLowerCase = map[string]RuleTypeEnum{
	"primarykey": RuleTypePrimarykey,
	"foreignkey": RuleTypeForeignkey,
	"uniquekey":  RuleTypeUniquekey,
}

// GetRuleTypeEnumValues Enumerates the set of values for RuleTypeEnum
func GetRuleTypeEnumValues() []RuleTypeEnum {
	values := make([]RuleTypeEnum, 0)
	for _, v := range mappingRuleTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRuleTypeEnumStringValues Enumerates the set of values in String for RuleTypeEnum
func GetRuleTypeEnumStringValues() []string {
	return []string{
		"PRIMARYKEY",
		"FOREIGNKEY",
		"UNIQUEKEY",
	}
}

// GetMappingRuleTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRuleTypeEnum(val string) (RuleTypeEnum, bool) {
	enum, ok := mappingRuleTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
