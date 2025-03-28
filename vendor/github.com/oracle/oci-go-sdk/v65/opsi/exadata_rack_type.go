// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.oracle.com/iaas/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"strings"
)

// ExadataRackTypeEnum Enum with underlying type: string
type ExadataRackTypeEnum string

// Set of constants representing the allowable values for ExadataRackTypeEnum
const (
	ExadataRackTypeFull         ExadataRackTypeEnum = "FULL"
	ExadataRackTypeHalf         ExadataRackTypeEnum = "HALF"
	ExadataRackTypeQuarter      ExadataRackTypeEnum = "QUARTER"
	ExadataRackTypeEighth       ExadataRackTypeEnum = "EIGHTH"
	ExadataRackTypeFlex         ExadataRackTypeEnum = "FLEX"
	ExadataRackTypeBase         ExadataRackTypeEnum = "BASE"
	ExadataRackTypeElastic      ExadataRackTypeEnum = "ELASTIC"
	ExadataRackTypeElasticBase  ExadataRackTypeEnum = "ELASTIC_BASE"
	ExadataRackTypeElasticLarge ExadataRackTypeEnum = "ELASTIC_LARGE"
)

var mappingExadataRackTypeEnum = map[string]ExadataRackTypeEnum{
	"FULL":          ExadataRackTypeFull,
	"HALF":          ExadataRackTypeHalf,
	"QUARTER":       ExadataRackTypeQuarter,
	"EIGHTH":        ExadataRackTypeEighth,
	"FLEX":          ExadataRackTypeFlex,
	"BASE":          ExadataRackTypeBase,
	"ELASTIC":       ExadataRackTypeElastic,
	"ELASTIC_BASE":  ExadataRackTypeElasticBase,
	"ELASTIC_LARGE": ExadataRackTypeElasticLarge,
}

var mappingExadataRackTypeEnumLowerCase = map[string]ExadataRackTypeEnum{
	"full":          ExadataRackTypeFull,
	"half":          ExadataRackTypeHalf,
	"quarter":       ExadataRackTypeQuarter,
	"eighth":        ExadataRackTypeEighth,
	"flex":          ExadataRackTypeFlex,
	"base":          ExadataRackTypeBase,
	"elastic":       ExadataRackTypeElastic,
	"elastic_base":  ExadataRackTypeElasticBase,
	"elastic_large": ExadataRackTypeElasticLarge,
}

// GetExadataRackTypeEnumValues Enumerates the set of values for ExadataRackTypeEnum
func GetExadataRackTypeEnumValues() []ExadataRackTypeEnum {
	values := make([]ExadataRackTypeEnum, 0)
	for _, v := range mappingExadataRackTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetExadataRackTypeEnumStringValues Enumerates the set of values in String for ExadataRackTypeEnum
func GetExadataRackTypeEnumStringValues() []string {
	return []string{
		"FULL",
		"HALF",
		"QUARTER",
		"EIGHTH",
		"FLEX",
		"BASE",
		"ELASTIC",
		"ELASTIC_BASE",
		"ELASTIC_LARGE",
	}
}

// GetMappingExadataRackTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExadataRackTypeEnum(val string) (ExadataRackTypeEnum, bool) {
	enum, ok := mappingExadataRackTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
