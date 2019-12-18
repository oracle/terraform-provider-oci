// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Manage applications in Oracle Cloud Infrastructure Marketplace.
//

package marketplace

// PricingCurrencyEnumEnum Enum with underlying type: string
type PricingCurrencyEnumEnum string

// Set of constants representing the allowable values for PricingCurrencyEnumEnum
const (
	PricingCurrencyEnumUsd PricingCurrencyEnumEnum = "USD"
)

var mappingPricingCurrencyEnum = map[string]PricingCurrencyEnumEnum{
	"USD": PricingCurrencyEnumUsd,
}

// GetPricingCurrencyEnumEnumValues Enumerates the set of values for PricingCurrencyEnumEnum
func GetPricingCurrencyEnumEnumValues() []PricingCurrencyEnumEnum {
	values := make([]PricingCurrencyEnumEnum, 0)
	for _, v := range mappingPricingCurrencyEnum {
		values = append(values, v)
	}
	return values
}
