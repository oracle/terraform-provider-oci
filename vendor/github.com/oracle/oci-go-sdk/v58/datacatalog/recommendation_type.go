// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

// RecommendationTypeEnum Enum with underlying type: string
type RecommendationTypeEnum string

// Set of constants representing the allowable values for RecommendationTypeEnum
const (
	RecommendationTypeLinkGlossaryTerm RecommendationTypeEnum = "LINK_GLOSSARY_TERM"
)

var mappingRecommendationTypeEnum = map[string]RecommendationTypeEnum{
	"LINK_GLOSSARY_TERM": RecommendationTypeLinkGlossaryTerm,
}

// GetRecommendationTypeEnumValues Enumerates the set of values for RecommendationTypeEnum
func GetRecommendationTypeEnumValues() []RecommendationTypeEnum {
	values := make([]RecommendationTypeEnum, 0)
	for _, v := range mappingRecommendationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRecommendationTypeEnumStringValues Enumerates the set of values in String for RecommendationTypeEnum
func GetRecommendationTypeEnumStringValues() []string {
	return []string{
		"LINK_GLOSSARY_TERM",
	}
}

// GetMappingRecommendationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRecommendationTypeEnum(val string) (RecommendationTypeEnum, bool) {
	mappingRecommendationTypeEnumIgnoreCase := make(map[string]RecommendationTypeEnum)
	for k, v := range mappingRecommendationTypeEnum {
		mappingRecommendationTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingRecommendationTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
