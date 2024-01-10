// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OneSubscription APIs
//
// OneSubscription APIs
//

package onesubscription

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AggregatedComputedUsageSummary Subscribed Service Contract details
type AggregatedComputedUsageSummary struct {

	// Subscription Id is an identifier associated to the service used for filter the Computed Usage in SPM
	SubscriptionId *string `mandatory:"true" json:"subscriptionId"`

	// Subscribed service line parent id
	ParentSubscribedServiceId *string `mandatory:"false" json:"parentSubscribedServiceId"`

	ParentProduct *ComputedUsageProduct `mandatory:"false" json:"parentProduct"`

	// Subscribed services contract line start date, expressed in RFC 3339 timestamp format.
	TimeStart *common.SDKTime `mandatory:"false" json:"timeStart"`

	// Subscribed services contract line end date, expressed in RFC 3339 timestamp format.
	TimeEnd *common.SDKTime `mandatory:"false" json:"timeEnd"`

	// Subscribed service asociated subscription plan number.
	PlanNumber *string `mandatory:"false" json:"planNumber"`

	// Currency code
	CurrencyCode *string `mandatory:"false" json:"currencyCode"`

	// Inernal SPM Ratecard Id at line level
	RateCardId *string `mandatory:"false" json:"rateCardId"`

	// Subscribed services pricing model
	PricingModel AggregatedComputedUsageSummaryPricingModelEnum `mandatory:"false" json:"pricingModel,omitempty"`

	// Aggregation of computed usages for the subscribed service.
	AggregatedComputedUsages []ComputedUsageAggregation `mandatory:"false" json:"aggregatedComputedUsages"`
}

func (m AggregatedComputedUsageSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AggregatedComputedUsageSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAggregatedComputedUsageSummaryPricingModelEnum(string(m.PricingModel)); !ok && m.PricingModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PricingModel: %s. Supported values are: %s.", m.PricingModel, strings.Join(GetAggregatedComputedUsageSummaryPricingModelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AggregatedComputedUsageSummaryPricingModelEnum Enum with underlying type: string
type AggregatedComputedUsageSummaryPricingModelEnum string

// Set of constants representing the allowable values for AggregatedComputedUsageSummaryPricingModelEnum
const (
	AggregatedComputedUsageSummaryPricingModelPayAsYouGo       AggregatedComputedUsageSummaryPricingModelEnum = "PAY_AS_YOU_GO"
	AggregatedComputedUsageSummaryPricingModelMonthly          AggregatedComputedUsageSummaryPricingModelEnum = "MONTHLY"
	AggregatedComputedUsageSummaryPricingModelAnnual           AggregatedComputedUsageSummaryPricingModelEnum = "ANNUAL"
	AggregatedComputedUsageSummaryPricingModelPrepaid          AggregatedComputedUsageSummaryPricingModelEnum = "PREPAID"
	AggregatedComputedUsageSummaryPricingModelFundedAllocation AggregatedComputedUsageSummaryPricingModelEnum = "FUNDED_ALLOCATION"
)

var mappingAggregatedComputedUsageSummaryPricingModelEnum = map[string]AggregatedComputedUsageSummaryPricingModelEnum{
	"PAY_AS_YOU_GO":     AggregatedComputedUsageSummaryPricingModelPayAsYouGo,
	"MONTHLY":           AggregatedComputedUsageSummaryPricingModelMonthly,
	"ANNUAL":            AggregatedComputedUsageSummaryPricingModelAnnual,
	"PREPAID":           AggregatedComputedUsageSummaryPricingModelPrepaid,
	"FUNDED_ALLOCATION": AggregatedComputedUsageSummaryPricingModelFundedAllocation,
}

var mappingAggregatedComputedUsageSummaryPricingModelEnumLowerCase = map[string]AggregatedComputedUsageSummaryPricingModelEnum{
	"pay_as_you_go":     AggregatedComputedUsageSummaryPricingModelPayAsYouGo,
	"monthly":           AggregatedComputedUsageSummaryPricingModelMonthly,
	"annual":            AggregatedComputedUsageSummaryPricingModelAnnual,
	"prepaid":           AggregatedComputedUsageSummaryPricingModelPrepaid,
	"funded_allocation": AggregatedComputedUsageSummaryPricingModelFundedAllocation,
}

// GetAggregatedComputedUsageSummaryPricingModelEnumValues Enumerates the set of values for AggregatedComputedUsageSummaryPricingModelEnum
func GetAggregatedComputedUsageSummaryPricingModelEnumValues() []AggregatedComputedUsageSummaryPricingModelEnum {
	values := make([]AggregatedComputedUsageSummaryPricingModelEnum, 0)
	for _, v := range mappingAggregatedComputedUsageSummaryPricingModelEnum {
		values = append(values, v)
	}
	return values
}

// GetAggregatedComputedUsageSummaryPricingModelEnumStringValues Enumerates the set of values in String for AggregatedComputedUsageSummaryPricingModelEnum
func GetAggregatedComputedUsageSummaryPricingModelEnumStringValues() []string {
	return []string{
		"PAY_AS_YOU_GO",
		"MONTHLY",
		"ANNUAL",
		"PREPAID",
		"FUNDED_ALLOCATION",
	}
}

// GetMappingAggregatedComputedUsageSummaryPricingModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAggregatedComputedUsageSummaryPricingModelEnum(val string) (AggregatedComputedUsageSummaryPricingModelEnum, bool) {
	enum, ok := mappingAggregatedComputedUsageSummaryPricingModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
