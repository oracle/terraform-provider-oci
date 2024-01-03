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

// RecommendationResourceTypeEnum Enum with underlying type: string
type RecommendationResourceTypeEnum string

// Set of constants representing the allowable values for RecommendationResourceTypeEnum
const (
	RecommendationResourceTypeDataEntity RecommendationResourceTypeEnum = "DATA_ENTITY"
	RecommendationResourceTypeAttribute  RecommendationResourceTypeEnum = "ATTRIBUTE"
	RecommendationResourceTypeTerm       RecommendationResourceTypeEnum = "TERM"
	RecommendationResourceTypeCategory   RecommendationResourceTypeEnum = "CATEGORY"
)

var mappingRecommendationResourceTypeEnum = map[string]RecommendationResourceTypeEnum{
	"DATA_ENTITY": RecommendationResourceTypeDataEntity,
	"ATTRIBUTE":   RecommendationResourceTypeAttribute,
	"TERM":        RecommendationResourceTypeTerm,
	"CATEGORY":    RecommendationResourceTypeCategory,
}

var mappingRecommendationResourceTypeEnumLowerCase = map[string]RecommendationResourceTypeEnum{
	"data_entity": RecommendationResourceTypeDataEntity,
	"attribute":   RecommendationResourceTypeAttribute,
	"term":        RecommendationResourceTypeTerm,
	"category":    RecommendationResourceTypeCategory,
}

// GetRecommendationResourceTypeEnumValues Enumerates the set of values for RecommendationResourceTypeEnum
func GetRecommendationResourceTypeEnumValues() []RecommendationResourceTypeEnum {
	values := make([]RecommendationResourceTypeEnum, 0)
	for _, v := range mappingRecommendationResourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRecommendationResourceTypeEnumStringValues Enumerates the set of values in String for RecommendationResourceTypeEnum
func GetRecommendationResourceTypeEnumStringValues() []string {
	return []string{
		"DATA_ENTITY",
		"ATTRIBUTE",
		"TERM",
		"CATEGORY",
	}
}

// GetMappingRecommendationResourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRecommendationResourceTypeEnum(val string) (RecommendationResourceTypeEnum, bool) {
	enum, ok := mappingRecommendationResourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
