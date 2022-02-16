// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

// PricingStrategyEnumEnum Enum with underlying type: string
type PricingStrategyEnumEnum string

// Set of constants representing the allowable values for PricingStrategyEnumEnum
const (
	PricingStrategyEnumPerOcpuLinear               PricingStrategyEnumEnum = "PER_OCPU_LINEAR"
	PricingStrategyEnumPerOcpuMinBilling           PricingStrategyEnumEnum = "PER_OCPU_MIN_BILLING"
	PricingStrategyEnumPerInstance                 PricingStrategyEnumEnum = "PER_INSTANCE"
	PricingStrategyEnumPerInstanceMonthlyInclusive PricingStrategyEnumEnum = "PER_INSTANCE_MONTHLY_INCLUSIVE"
)

var mappingPricingStrategyEnumEnum = map[string]PricingStrategyEnumEnum{
	"PER_OCPU_LINEAR":                PricingStrategyEnumPerOcpuLinear,
	"PER_OCPU_MIN_BILLING":           PricingStrategyEnumPerOcpuMinBilling,
	"PER_INSTANCE":                   PricingStrategyEnumPerInstance,
	"PER_INSTANCE_MONTHLY_INCLUSIVE": PricingStrategyEnumPerInstanceMonthlyInclusive,
}

// GetPricingStrategyEnumEnumValues Enumerates the set of values for PricingStrategyEnumEnum
func GetPricingStrategyEnumEnumValues() []PricingStrategyEnumEnum {
	values := make([]PricingStrategyEnumEnum, 0)
	for _, v := range mappingPricingStrategyEnumEnum {
		values = append(values, v)
	}
	return values
}

// GetPricingStrategyEnumEnumStringValues Enumerates the set of values in String for PricingStrategyEnumEnum
func GetPricingStrategyEnumEnumStringValues() []string {
	return []string{
		"PER_OCPU_LINEAR",
		"PER_OCPU_MIN_BILLING",
		"PER_INSTANCE",
		"PER_INSTANCE_MONTHLY_INCLUSIVE",
	}
}

// GetMappingPricingStrategyEnumEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPricingStrategyEnumEnum(val string) (PricingStrategyEnumEnum, bool) {
	mappingPricingStrategyEnumEnumIgnoreCase := make(map[string]PricingStrategyEnumEnum)
	for k, v := range mappingPricingStrategyEnumEnum {
		mappingPricingStrategyEnumEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingPricingStrategyEnumEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
