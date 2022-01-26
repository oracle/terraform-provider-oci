// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Manage applications in Oracle Cloud Infrastructure Marketplace.
//

package marketplace

import (
	"github.com/oracle/oci-go-sdk/v56/common"
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

var mappingThirdPartyPaidListingEligibilityEligibilityReason = map[string]ThirdPartyPaidListingEligibilityEligibilityReasonEnum{
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

// GetThirdPartyPaidListingEligibilityEligibilityReasonEnumValues Enumerates the set of values for ThirdPartyPaidListingEligibilityEligibilityReasonEnum
func GetThirdPartyPaidListingEligibilityEligibilityReasonEnumValues() []ThirdPartyPaidListingEligibilityEligibilityReasonEnum {
	values := make([]ThirdPartyPaidListingEligibilityEligibilityReasonEnum, 0)
	for _, v := range mappingThirdPartyPaidListingEligibilityEligibilityReason {
		values = append(values, v)
	}
	return values
}
