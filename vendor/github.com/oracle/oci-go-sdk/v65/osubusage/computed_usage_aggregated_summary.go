// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OneSubscription API Usage Computation
//
// OneSubscription API Common set of Subscription Plan Management (SPM) Usage Computation resources
//

package osubusage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ComputedUsageAggregatedSummary Subscribed Service Contract details
type ComputedUsageAggregatedSummary struct {

	// Subscription Id is an identifier associated to the service used for filter the Computed Usage in SPM
	SubscriptionId *string `mandatory:"true" json:"subscriptionId"`

	// Subscribed service line parent id
	ParentSubscribedServiceId *string `mandatory:"false" json:"parentSubscribedServiceId"`

	ParentProduct *Product `mandatory:"false" json:"parentProduct"`

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
	PricingModel ComputedUsageAggregatedSummaryPricingModelEnum `mandatory:"false" json:"pricingModel,omitempty"`

	// Aggregation of computed usages for the subscribed service.
	AggregatedComputedUsages []ComputedUsageAggregation `mandatory:"false" json:"aggregatedComputedUsages"`
}

func (m ComputedUsageAggregatedSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ComputedUsageAggregatedSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingComputedUsageAggregatedSummaryPricingModelEnum(string(m.PricingModel)); !ok && m.PricingModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PricingModel: %s. Supported values are: %s.", m.PricingModel, strings.Join(GetComputedUsageAggregatedSummaryPricingModelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ComputedUsageAggregatedSummaryPricingModelEnum Enum with underlying type: string
type ComputedUsageAggregatedSummaryPricingModelEnum string

// Set of constants representing the allowable values for ComputedUsageAggregatedSummaryPricingModelEnum
const (
	ComputedUsageAggregatedSummaryPricingModelPayAsYouGo       ComputedUsageAggregatedSummaryPricingModelEnum = "PAY_AS_YOU_GO"
	ComputedUsageAggregatedSummaryPricingModelMonthly          ComputedUsageAggregatedSummaryPricingModelEnum = "MONTHLY"
	ComputedUsageAggregatedSummaryPricingModelAnnual           ComputedUsageAggregatedSummaryPricingModelEnum = "ANNUAL"
	ComputedUsageAggregatedSummaryPricingModelPrepaid          ComputedUsageAggregatedSummaryPricingModelEnum = "PREPAID"
	ComputedUsageAggregatedSummaryPricingModelFundedAllocation ComputedUsageAggregatedSummaryPricingModelEnum = "FUNDED_ALLOCATION"
)

var mappingComputedUsageAggregatedSummaryPricingModelEnum = map[string]ComputedUsageAggregatedSummaryPricingModelEnum{
	"PAY_AS_YOU_GO":     ComputedUsageAggregatedSummaryPricingModelPayAsYouGo,
	"MONTHLY":           ComputedUsageAggregatedSummaryPricingModelMonthly,
	"ANNUAL":            ComputedUsageAggregatedSummaryPricingModelAnnual,
	"PREPAID":           ComputedUsageAggregatedSummaryPricingModelPrepaid,
	"FUNDED_ALLOCATION": ComputedUsageAggregatedSummaryPricingModelFundedAllocation,
}

var mappingComputedUsageAggregatedSummaryPricingModelEnumLowerCase = map[string]ComputedUsageAggregatedSummaryPricingModelEnum{
	"pay_as_you_go":     ComputedUsageAggregatedSummaryPricingModelPayAsYouGo,
	"monthly":           ComputedUsageAggregatedSummaryPricingModelMonthly,
	"annual":            ComputedUsageAggregatedSummaryPricingModelAnnual,
	"prepaid":           ComputedUsageAggregatedSummaryPricingModelPrepaid,
	"funded_allocation": ComputedUsageAggregatedSummaryPricingModelFundedAllocation,
}

// GetComputedUsageAggregatedSummaryPricingModelEnumValues Enumerates the set of values for ComputedUsageAggregatedSummaryPricingModelEnum
func GetComputedUsageAggregatedSummaryPricingModelEnumValues() []ComputedUsageAggregatedSummaryPricingModelEnum {
	values := make([]ComputedUsageAggregatedSummaryPricingModelEnum, 0)
	for _, v := range mappingComputedUsageAggregatedSummaryPricingModelEnum {
		values = append(values, v)
	}
	return values
}

// GetComputedUsageAggregatedSummaryPricingModelEnumStringValues Enumerates the set of values in String for ComputedUsageAggregatedSummaryPricingModelEnum
func GetComputedUsageAggregatedSummaryPricingModelEnumStringValues() []string {
	return []string{
		"PAY_AS_YOU_GO",
		"MONTHLY",
		"ANNUAL",
		"PREPAID",
		"FUNDED_ALLOCATION",
	}
}

// GetMappingComputedUsageAggregatedSummaryPricingModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingComputedUsageAggregatedSummaryPricingModelEnum(val string) (ComputedUsageAggregatedSummaryPricingModelEnum, bool) {
	enum, ok := mappingComputedUsageAggregatedSummaryPricingModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
