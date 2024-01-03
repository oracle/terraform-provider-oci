// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Catalog API
//
// Manage solutions in Oracle Cloud Infrastructure Service Catalog.
//

package servicecatalog

import (
	"strings"
)

// PricingTypeEnumEnum Enum with underlying type: string
type PricingTypeEnumEnum string

// Set of constants representing the allowable values for PricingTypeEnumEnum
const (
	PricingTypeEnumFree  PricingTypeEnumEnum = "FREE"
	PricingTypeEnumByol  PricingTypeEnumEnum = "BYOL"
	PricingTypeEnumPaygo PricingTypeEnumEnum = "PAYGO"
)

var mappingPricingTypeEnumEnum = map[string]PricingTypeEnumEnum{
	"FREE":  PricingTypeEnumFree,
	"BYOL":  PricingTypeEnumByol,
	"PAYGO": PricingTypeEnumPaygo,
}

var mappingPricingTypeEnumEnumLowerCase = map[string]PricingTypeEnumEnum{
	"free":  PricingTypeEnumFree,
	"byol":  PricingTypeEnumByol,
	"paygo": PricingTypeEnumPaygo,
}

// GetPricingTypeEnumEnumValues Enumerates the set of values for PricingTypeEnumEnum
func GetPricingTypeEnumEnumValues() []PricingTypeEnumEnum {
	values := make([]PricingTypeEnumEnum, 0)
	for _, v := range mappingPricingTypeEnumEnum {
		values = append(values, v)
	}
	return values
}

// GetPricingTypeEnumEnumStringValues Enumerates the set of values in String for PricingTypeEnumEnum
func GetPricingTypeEnumEnumStringValues() []string {
	return []string{
		"FREE",
		"BYOL",
		"PAYGO",
	}
}

// GetMappingPricingTypeEnumEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPricingTypeEnumEnum(val string) (PricingTypeEnumEnum, bool) {
	enum, ok := mappingPricingTypeEnumEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
