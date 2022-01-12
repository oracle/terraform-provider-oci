// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
// For more information, see Data Catalog (https://docs.oracle.com/iaas/data-catalog/home.htm).
//

package datacatalog

// RuleOriginTypeEnum Enum with underlying type: string
type RuleOriginTypeEnum string

// Set of constants representing the allowable values for RuleOriginTypeEnum
const (
	RuleOriginTypeSource    RuleOriginTypeEnum = "SOURCE"
	RuleOriginTypeUser      RuleOriginTypeEnum = "USER"
	RuleOriginTypeProfiling RuleOriginTypeEnum = "PROFILING"
)

var mappingRuleOriginType = map[string]RuleOriginTypeEnum{
	"SOURCE":    RuleOriginTypeSource,
	"USER":      RuleOriginTypeUser,
	"PROFILING": RuleOriginTypeProfiling,
}

// GetRuleOriginTypeEnumValues Enumerates the set of values for RuleOriginTypeEnum
func GetRuleOriginTypeEnumValues() []RuleOriginTypeEnum {
	values := make([]RuleOriginTypeEnum, 0)
	for _, v := range mappingRuleOriginType {
		values = append(values, v)
	}
	return values
}
