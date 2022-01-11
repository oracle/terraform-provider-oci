// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Manage applications in Oracle Cloud Infrastructure Marketplace.
//

package marketplace

// ListingTypeEnum Enum with underlying type: string
type ListingTypeEnum string

// Set of constants representing the allowable values for ListingTypeEnum
const (
	ListingTypeCommunity ListingTypeEnum = "COMMUNITY"
	ListingTypePartner   ListingTypeEnum = "PARTNER"
	ListingTypePrivate   ListingTypeEnum = "PRIVATE"
)

var mappingListingType = map[string]ListingTypeEnum{
	"COMMUNITY": ListingTypeCommunity,
	"PARTNER":   ListingTypePartner,
	"PRIVATE":   ListingTypePrivate,
}

// GetListingTypeEnumValues Enumerates the set of values for ListingTypeEnum
func GetListingTypeEnumValues() []ListingTypeEnum {
	values := make([]ListingTypeEnum, 0)
	for _, v := range mappingListingType {
		values = append(values, v)
	}
	return values
}
