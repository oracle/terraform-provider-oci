// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"strings"
)

// ProductSortByEnum Enum with underlying type: string
type ProductSortByEnum string

// Set of constants representing the allowable values for ProductSortByEnum
const (
	ProductSortByDisplayName         ProductSortByEnum = "displayName"
	ProductSortByResourceDisplayName ProductSortByEnum = "resourceDisplayName"
)

var mappingProductSortByEnum = map[string]ProductSortByEnum{
	"displayName":         ProductSortByDisplayName,
	"resourceDisplayName": ProductSortByResourceDisplayName,
}

var mappingProductSortByEnumLowerCase = map[string]ProductSortByEnum{
	"displayname":         ProductSortByDisplayName,
	"resourcedisplayname": ProductSortByResourceDisplayName,
}

// GetProductSortByEnumValues Enumerates the set of values for ProductSortByEnum
func GetProductSortByEnumValues() []ProductSortByEnum {
	values := make([]ProductSortByEnum, 0)
	for _, v := range mappingProductSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetProductSortByEnumStringValues Enumerates the set of values in String for ProductSortByEnum
func GetProductSortByEnumStringValues() []string {
	return []string{
		"displayName",
		"resourceDisplayName",
	}
}

// GetMappingProductSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingProductSortByEnum(val string) (ProductSortByEnum, bool) {
	enum, ok := mappingProductSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
