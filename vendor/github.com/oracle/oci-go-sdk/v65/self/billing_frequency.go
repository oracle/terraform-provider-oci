// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// SELF Service API
//
// Use the SELF Service API to manage Subscriptions in Oracle Cloud Infrastructure Marketplace. For more information, see Overview of Marketplace (https://docs.oracle.com/iaas/Content/Marketplace/Concepts/marketoverview.htm)
//

package self

import (
	"strings"
)

// BillingFrequencyEnum Enum with underlying type: string
type BillingFrequencyEnum string

// Set of constants representing the allowable values for BillingFrequencyEnum
const (
	BillingFrequencyMonthly    BillingFrequencyEnum = "MONTHLY"
	BillingFrequencyQuarterly  BillingFrequencyEnum = "QUARTERLY"
	BillingFrequencySemiAnnual BillingFrequencyEnum = "SEMI_ANNUAL"
	BillingFrequencyAnnual     BillingFrequencyEnum = "ANNUAL"
	BillingFrequencyBiennial   BillingFrequencyEnum = "BIENNIAL"
	BillingFrequencyTriennial  BillingFrequencyEnum = "TRIENNIAL"
	BillingFrequencyYearly     BillingFrequencyEnum = "YEARLY"
)

var mappingBillingFrequencyEnum = map[string]BillingFrequencyEnum{
	"MONTHLY":     BillingFrequencyMonthly,
	"QUARTERLY":   BillingFrequencyQuarterly,
	"SEMI_ANNUAL": BillingFrequencySemiAnnual,
	"ANNUAL":      BillingFrequencyAnnual,
	"BIENNIAL":    BillingFrequencyBiennial,
	"TRIENNIAL":   BillingFrequencyTriennial,
	"YEARLY":      BillingFrequencyYearly,
}

var mappingBillingFrequencyEnumLowerCase = map[string]BillingFrequencyEnum{
	"monthly":     BillingFrequencyMonthly,
	"quarterly":   BillingFrequencyQuarterly,
	"semi_annual": BillingFrequencySemiAnnual,
	"annual":      BillingFrequencyAnnual,
	"biennial":    BillingFrequencyBiennial,
	"triennial":   BillingFrequencyTriennial,
	"yearly":      BillingFrequencyYearly,
}

// GetBillingFrequencyEnumValues Enumerates the set of values for BillingFrequencyEnum
func GetBillingFrequencyEnumValues() []BillingFrequencyEnum {
	values := make([]BillingFrequencyEnum, 0)
	for _, v := range mappingBillingFrequencyEnum {
		values = append(values, v)
	}
	return values
}

// GetBillingFrequencyEnumStringValues Enumerates the set of values in String for BillingFrequencyEnum
func GetBillingFrequencyEnumStringValues() []string {
	return []string{
		"MONTHLY",
		"QUARTERLY",
		"SEMI_ANNUAL",
		"ANNUAL",
		"BIENNIAL",
		"TRIENNIAL",
		"YEARLY",
	}
}

// GetMappingBillingFrequencyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBillingFrequencyEnum(val string) (BillingFrequencyEnum, bool) {
	enum, ok := mappingBillingFrequencyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
