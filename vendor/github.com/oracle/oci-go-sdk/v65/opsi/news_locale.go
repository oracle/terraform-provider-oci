// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"strings"
)

// NewsLocaleEnum Enum with underlying type: string
type NewsLocaleEnum string

// Set of constants representing the allowable values for NewsLocaleEnum
const (
	NewsLocaleEn NewsLocaleEnum = "EN"
)

var mappingNewsLocaleEnum = map[string]NewsLocaleEnum{
	"EN": NewsLocaleEn,
}

var mappingNewsLocaleEnumLowerCase = map[string]NewsLocaleEnum{
	"en": NewsLocaleEn,
}

// GetNewsLocaleEnumValues Enumerates the set of values for NewsLocaleEnum
func GetNewsLocaleEnumValues() []NewsLocaleEnum {
	values := make([]NewsLocaleEnum, 0)
	for _, v := range mappingNewsLocaleEnum {
		values = append(values, v)
	}
	return values
}

// GetNewsLocaleEnumStringValues Enumerates the set of values in String for NewsLocaleEnum
func GetNewsLocaleEnumStringValues() []string {
	return []string{
		"EN",
	}
}

// GetMappingNewsLocaleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNewsLocaleEnum(val string) (NewsLocaleEnum, bool) {
	enum, ok := mappingNewsLocaleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
