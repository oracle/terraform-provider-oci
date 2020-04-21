// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
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
