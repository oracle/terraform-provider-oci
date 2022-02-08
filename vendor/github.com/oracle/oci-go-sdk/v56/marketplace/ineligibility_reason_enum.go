// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Use the Marketplace API to manage applications in Oracle Cloud Infrastructure Marketplace. For more information, see Overview of Marketplace (https://docs.cloud.oracle.com/Content/Marketplace/Concepts/marketoverview.htm)
//

package marketplace

// IneligibilityReasonEnumEnum Enum with underlying type: string
type IneligibilityReasonEnumEnum string

// Set of constants representing the allowable values for IneligibilityReasonEnumEnum
const (
	IneligibilityReasonEnumIneligibleAccountCountry         IneligibilityReasonEnumEnum = "INELIGIBLE_ACCOUNT_COUNTRY"
	IneligibilityReasonEnumIneligibleRegion                 IneligibilityReasonEnumEnum = "INELIGIBLE_REGION"
	IneligibilityReasonEnumIneligibleAccountBlacklisted     IneligibilityReasonEnumEnum = "INELIGIBLE_ACCOUNT_BLACKLISTED"
	IneligibilityReasonEnumIneligibleAccountFeatureDisabled IneligibilityReasonEnumEnum = "INELIGIBLE_ACCOUNT_FEATURE_DISABLED"
	IneligibilityReasonEnumIneligibleAccountCurrency        IneligibilityReasonEnumEnum = "INELIGIBLE_ACCOUNT_CURRENCY"
	IneligibilityReasonEnumIneligibleAccountNotPaid         IneligibilityReasonEnumEnum = "INELIGIBLE_ACCOUNT_NOT_PAID"
	IneligibilityReasonEnumIneligibleAccountInternal        IneligibilityReasonEnumEnum = "INELIGIBLE_ACCOUNT_INTERNAL"
	IneligibilityReasonEnumIneligibleAccountGovSubscription IneligibilityReasonEnumEnum = "INELIGIBLE_ACCOUNT_GOV_SUBSCRIPTION"
	IneligibilityReasonEnumIneligiblePaidListingThrottled   IneligibilityReasonEnumEnum = "INELIGIBLE_PAID_LISTING_THROTTLED"
	IneligibilityReasonEnumNotAuthorized                    IneligibilityReasonEnumEnum = "NOT_AUTHORIZED"
)

var mappingIneligibilityReasonEnumEnum = map[string]IneligibilityReasonEnumEnum{
	"INELIGIBLE_ACCOUNT_COUNTRY":          IneligibilityReasonEnumIneligibleAccountCountry,
	"INELIGIBLE_REGION":                   IneligibilityReasonEnumIneligibleRegion,
	"INELIGIBLE_ACCOUNT_BLACKLISTED":      IneligibilityReasonEnumIneligibleAccountBlacklisted,
	"INELIGIBLE_ACCOUNT_FEATURE_DISABLED": IneligibilityReasonEnumIneligibleAccountFeatureDisabled,
	"INELIGIBLE_ACCOUNT_CURRENCY":         IneligibilityReasonEnumIneligibleAccountCurrency,
	"INELIGIBLE_ACCOUNT_NOT_PAID":         IneligibilityReasonEnumIneligibleAccountNotPaid,
	"INELIGIBLE_ACCOUNT_INTERNAL":         IneligibilityReasonEnumIneligibleAccountInternal,
	"INELIGIBLE_ACCOUNT_GOV_SUBSCRIPTION": IneligibilityReasonEnumIneligibleAccountGovSubscription,
	"INELIGIBLE_PAID_LISTING_THROTTLED":   IneligibilityReasonEnumIneligiblePaidListingThrottled,
	"NOT_AUTHORIZED":                      IneligibilityReasonEnumNotAuthorized,
}

// GetIneligibilityReasonEnumEnumValues Enumerates the set of values for IneligibilityReasonEnumEnum
func GetIneligibilityReasonEnumEnumValues() []IneligibilityReasonEnumEnum {
	values := make([]IneligibilityReasonEnumEnum, 0)
	for _, v := range mappingIneligibilityReasonEnumEnum {
		values = append(values, v)
	}
	return values
}

// GetIneligibilityReasonEnumEnumStringValues Enumerates the set of values in String for IneligibilityReasonEnumEnum
func GetIneligibilityReasonEnumEnumStringValues() []string {
	return []string{
		"INELIGIBLE_ACCOUNT_COUNTRY",
		"INELIGIBLE_REGION",
		"INELIGIBLE_ACCOUNT_BLACKLISTED",
		"INELIGIBLE_ACCOUNT_FEATURE_DISABLED",
		"INELIGIBLE_ACCOUNT_CURRENCY",
		"INELIGIBLE_ACCOUNT_NOT_PAID",
		"INELIGIBLE_ACCOUNT_INTERNAL",
		"INELIGIBLE_ACCOUNT_GOV_SUBSCRIPTION",
		"INELIGIBLE_PAID_LISTING_THROTTLED",
		"NOT_AUTHORIZED",
	}
}
