// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Use the Marketplace API to manage applications in Oracle Cloud Infrastructure Marketplace. For more information, see Overview of Marketplace (https://docs.cloud.oracle.com/Content/Marketplace/Concepts/marketoverview.htm)
//

package marketplace

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ThirdPartyPaidListingEligibility Tenant eligibility for using third party paid listings
type ThirdPartyPaidListingEligibility struct {

	// Whether the tenant is permitted to use paid listings
	IsPaidListingEligible *bool `mandatory:"true" json:"isPaidListingEligible"`

	// Whether the tenant is currently prevented from using paid listings because of throttling
	IsPaidListingThrottled *bool `mandatory:"true" json:"isPaidListingThrottled"`

	// Reason the account is ineligible to launch paid listings
	EligibilityReason ThirdPartyPaidListingEligibilityEligibilityReasonEnum `mandatory:"true" json:"eligibilityReason"`
}

func (m ThirdPartyPaidListingEligibility) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ThirdPartyPaidListingEligibility) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingThirdPartyPaidListingEligibilityEligibilityReasonEnum(string(m.EligibilityReason)); !ok && m.EligibilityReason != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EligibilityReason: %s. Supported values are: %s.", m.EligibilityReason, strings.Join(GetThirdPartyPaidListingEligibilityEligibilityReasonEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ThirdPartyPaidListingEligibilityEligibilityReasonEnum Enum with underlying type: string
type ThirdPartyPaidListingEligibilityEligibilityReasonEnum string

// Set of constants representing the allowable values for ThirdPartyPaidListingEligibilityEligibilityReasonEnum
const (
	ThirdPartyPaidListingEligibilityEligibilityReasonEligible                         ThirdPartyPaidListingEligibilityEligibilityReasonEnum = "ELIGIBLE"
	ThirdPartyPaidListingEligibilityEligibilityReasonIneligibleAccountCountry         ThirdPartyPaidListingEligibilityEligibilityReasonEnum = "INELIGIBLE_ACCOUNT_COUNTRY"
	ThirdPartyPaidListingEligibilityEligibilityReasonIneligibleRegion                 ThirdPartyPaidListingEligibilityEligibilityReasonEnum = "INELIGIBLE_REGION"
	ThirdPartyPaidListingEligibilityEligibilityReasonIneligibleAccountBlacklisted     ThirdPartyPaidListingEligibilityEligibilityReasonEnum = "INELIGIBLE_ACCOUNT_BLACKLISTED"
	ThirdPartyPaidListingEligibilityEligibilityReasonIneligibleAccountFeatureDisabled ThirdPartyPaidListingEligibilityEligibilityReasonEnum = "INELIGIBLE_ACCOUNT_FEATURE_DISABLED"
	ThirdPartyPaidListingEligibilityEligibilityReasonIneligibleAccountCurrency        ThirdPartyPaidListingEligibilityEligibilityReasonEnum = "INELIGIBLE_ACCOUNT_CURRENCY"
	ThirdPartyPaidListingEligibilityEligibilityReasonIneligibleAccountNotPaid         ThirdPartyPaidListingEligibilityEligibilityReasonEnum = "INELIGIBLE_ACCOUNT_NOT_PAID"
	ThirdPartyPaidListingEligibilityEligibilityReasonIneligibleAccountInternal        ThirdPartyPaidListingEligibilityEligibilityReasonEnum = "INELIGIBLE_ACCOUNT_INTERNAL"
	ThirdPartyPaidListingEligibilityEligibilityReasonIneligibleAccountGovSubscription ThirdPartyPaidListingEligibilityEligibilityReasonEnum = "INELIGIBLE_ACCOUNT_GOV_SUBSCRIPTION"
	ThirdPartyPaidListingEligibilityEligibilityReasonNotAuthorized                    ThirdPartyPaidListingEligibilityEligibilityReasonEnum = "NOT_AUTHORIZED"
)

var mappingThirdPartyPaidListingEligibilityEligibilityReasonEnum = map[string]ThirdPartyPaidListingEligibilityEligibilityReasonEnum{
	"ELIGIBLE":                            ThirdPartyPaidListingEligibilityEligibilityReasonEligible,
	"INELIGIBLE_ACCOUNT_COUNTRY":          ThirdPartyPaidListingEligibilityEligibilityReasonIneligibleAccountCountry,
	"INELIGIBLE_REGION":                   ThirdPartyPaidListingEligibilityEligibilityReasonIneligibleRegion,
	"INELIGIBLE_ACCOUNT_BLACKLISTED":      ThirdPartyPaidListingEligibilityEligibilityReasonIneligibleAccountBlacklisted,
	"INELIGIBLE_ACCOUNT_FEATURE_DISABLED": ThirdPartyPaidListingEligibilityEligibilityReasonIneligibleAccountFeatureDisabled,
	"INELIGIBLE_ACCOUNT_CURRENCY":         ThirdPartyPaidListingEligibilityEligibilityReasonIneligibleAccountCurrency,
	"INELIGIBLE_ACCOUNT_NOT_PAID":         ThirdPartyPaidListingEligibilityEligibilityReasonIneligibleAccountNotPaid,
	"INELIGIBLE_ACCOUNT_INTERNAL":         ThirdPartyPaidListingEligibilityEligibilityReasonIneligibleAccountInternal,
	"INELIGIBLE_ACCOUNT_GOV_SUBSCRIPTION": ThirdPartyPaidListingEligibilityEligibilityReasonIneligibleAccountGovSubscription,
	"NOT_AUTHORIZED":                      ThirdPartyPaidListingEligibilityEligibilityReasonNotAuthorized,
}

var mappingThirdPartyPaidListingEligibilityEligibilityReasonEnumLowerCase = map[string]ThirdPartyPaidListingEligibilityEligibilityReasonEnum{
	"eligible":                            ThirdPartyPaidListingEligibilityEligibilityReasonEligible,
	"ineligible_account_country":          ThirdPartyPaidListingEligibilityEligibilityReasonIneligibleAccountCountry,
	"ineligible_region":                   ThirdPartyPaidListingEligibilityEligibilityReasonIneligibleRegion,
	"ineligible_account_blacklisted":      ThirdPartyPaidListingEligibilityEligibilityReasonIneligibleAccountBlacklisted,
	"ineligible_account_feature_disabled": ThirdPartyPaidListingEligibilityEligibilityReasonIneligibleAccountFeatureDisabled,
	"ineligible_account_currency":         ThirdPartyPaidListingEligibilityEligibilityReasonIneligibleAccountCurrency,
	"ineligible_account_not_paid":         ThirdPartyPaidListingEligibilityEligibilityReasonIneligibleAccountNotPaid,
	"ineligible_account_internal":         ThirdPartyPaidListingEligibilityEligibilityReasonIneligibleAccountInternal,
	"ineligible_account_gov_subscription": ThirdPartyPaidListingEligibilityEligibilityReasonIneligibleAccountGovSubscription,
	"not_authorized":                      ThirdPartyPaidListingEligibilityEligibilityReasonNotAuthorized,
}

// GetThirdPartyPaidListingEligibilityEligibilityReasonEnumValues Enumerates the set of values for ThirdPartyPaidListingEligibilityEligibilityReasonEnum
func GetThirdPartyPaidListingEligibilityEligibilityReasonEnumValues() []ThirdPartyPaidListingEligibilityEligibilityReasonEnum {
	values := make([]ThirdPartyPaidListingEligibilityEligibilityReasonEnum, 0)
	for _, v := range mappingThirdPartyPaidListingEligibilityEligibilityReasonEnum {
		values = append(values, v)
	}
	return values
}

// GetThirdPartyPaidListingEligibilityEligibilityReasonEnumStringValues Enumerates the set of values in String for ThirdPartyPaidListingEligibilityEligibilityReasonEnum
func GetThirdPartyPaidListingEligibilityEligibilityReasonEnumStringValues() []string {
	return []string{
		"ELIGIBLE",
		"INELIGIBLE_ACCOUNT_COUNTRY",
		"INELIGIBLE_REGION",
		"INELIGIBLE_ACCOUNT_BLACKLISTED",
		"INELIGIBLE_ACCOUNT_FEATURE_DISABLED",
		"INELIGIBLE_ACCOUNT_CURRENCY",
		"INELIGIBLE_ACCOUNT_NOT_PAID",
		"INELIGIBLE_ACCOUNT_INTERNAL",
		"INELIGIBLE_ACCOUNT_GOV_SUBSCRIPTION",
		"NOT_AUTHORIZED",
	}
}

// GetMappingThirdPartyPaidListingEligibilityEligibilityReasonEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingThirdPartyPaidListingEligibilityEligibilityReasonEnum(val string) (ThirdPartyPaidListingEligibilityEligibilityReasonEnum, bool) {
	enum, ok := mappingThirdPartyPaidListingEligibilityEligibilityReasonEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
