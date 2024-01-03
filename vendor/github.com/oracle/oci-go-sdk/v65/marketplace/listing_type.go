// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Use the Marketplace API to manage applications in Oracle Cloud Infrastructure Marketplace. For more information, see Overview of Marketplace (https://docs.cloud.oracle.com/Content/Marketplace/Concepts/marketoverview.htm)
//

package marketplace

import (
	"strings"
)

// ListingTypeEnum Enum with underlying type: string
type ListingTypeEnum string

// Set of constants representing the allowable values for ListingTypeEnum
const (
	ListingTypeCommunity ListingTypeEnum = "COMMUNITY"
	ListingTypePartner   ListingTypeEnum = "PARTNER"
	ListingTypePrivate   ListingTypeEnum = "PRIVATE"
)

var mappingListingTypeEnum = map[string]ListingTypeEnum{
	"COMMUNITY": ListingTypeCommunity,
	"PARTNER":   ListingTypePartner,
	"PRIVATE":   ListingTypePrivate,
}

var mappingListingTypeEnumLowerCase = map[string]ListingTypeEnum{
	"community": ListingTypeCommunity,
	"partner":   ListingTypePartner,
	"private":   ListingTypePrivate,
}

// GetListingTypeEnumValues Enumerates the set of values for ListingTypeEnum
func GetListingTypeEnumValues() []ListingTypeEnum {
	values := make([]ListingTypeEnum, 0)
	for _, v := range mappingListingTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListingTypeEnumStringValues Enumerates the set of values in String for ListingTypeEnum
func GetListingTypeEnumStringValues() []string {
	return []string{
		"COMMUNITY",
		"PARTNER",
		"PRIVATE",
	}
}

// GetMappingListingTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListingTypeEnum(val string) (ListingTypeEnum, bool) {
	enum, ok := mappingListingTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
