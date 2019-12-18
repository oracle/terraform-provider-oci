// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Manage applications in Oracle Cloud Infrastructure Marketplace.
//

package marketplace

// PricingTypeEnumEnum Enum with underlying type: string
type PricingTypeEnumEnum string

// Set of constants representing the allowable values for PricingTypeEnumEnum
const (
	PricingTypeEnumFree  PricingTypeEnumEnum = "FREE"
	PricingTypeEnumByol  PricingTypeEnumEnum = "BYOL"
	PricingTypeEnumPaygo PricingTypeEnumEnum = "PAYGO"
)

var mappingPricingTypeEnum = map[string]PricingTypeEnumEnum{
	"FREE":  PricingTypeEnumFree,
	"BYOL":  PricingTypeEnumByol,
	"PAYGO": PricingTypeEnumPaygo,
}

// GetPricingTypeEnumEnumValues Enumerates the set of values for PricingTypeEnumEnum
func GetPricingTypeEnumEnumValues() []PricingTypeEnumEnum {
	values := make([]PricingTypeEnumEnum, 0)
	for _, v := range mappingPricingTypeEnum {
		values = append(values, v)
	}
	return values
}
