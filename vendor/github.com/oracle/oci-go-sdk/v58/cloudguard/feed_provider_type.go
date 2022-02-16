// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

import (
	"strings"
)

// FeedProviderTypeEnum Enum with underlying type: string
type FeedProviderTypeEnum string

// Set of constants representing the allowable values for FeedProviderTypeEnum
const (
	FeedProviderTypeCustomer FeedProviderTypeEnum = "CUSTOMER"
	FeedProviderTypeOracle   FeedProviderTypeEnum = "ORACLE"
)

var mappingFeedProviderTypeEnum = map[string]FeedProviderTypeEnum{
	"CUSTOMER": FeedProviderTypeCustomer,
	"ORACLE":   FeedProviderTypeOracle,
}

// GetFeedProviderTypeEnumValues Enumerates the set of values for FeedProviderTypeEnum
func GetFeedProviderTypeEnumValues() []FeedProviderTypeEnum {
	values := make([]FeedProviderTypeEnum, 0)
	for _, v := range mappingFeedProviderTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetFeedProviderTypeEnumStringValues Enumerates the set of values in String for FeedProviderTypeEnum
func GetFeedProviderTypeEnumStringValues() []string {
	return []string{
		"CUSTOMER",
		"ORACLE",
	}
}

// GetMappingFeedProviderTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFeedProviderTypeEnum(val string) (FeedProviderTypeEnum, bool) {
	mappingFeedProviderTypeEnumIgnoreCase := make(map[string]FeedProviderTypeEnum)
	for k, v := range mappingFeedProviderTypeEnum {
		mappingFeedProviderTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingFeedProviderTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
