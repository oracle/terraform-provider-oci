// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// ComputedUsageSummary Computed usage summary object.
type ComputedUsageSummary struct {

	// SPM internal computed usage ID, which is a 32-character string.
	ComputedUsageId *string `mandatory:"true" json:"computedUsageId"`

	// Computed usage created time, expressed in RFC 3339 timestamp format.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Computed usage updated time, expressed in RFC 3339 timestamp format.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Subscribed service line parent ID.
	ParentSubscribedServiceId *string `mandatory:"false" json:"parentSubscribedServiceId"`

	ParentProduct *ComputedUsageProduct `mandatory:"false" json:"parentProduct"`

	// Subscription plan number.
	PlanNumber *string `mandatory:"false" json:"planNumber"`

	// Currency code.
	CurrencyCode *string `mandatory:"false" json:"currencyCode"`

	// References the tier in the rate card for the associated usage (OCI uses the same reference to cross-reference for correctness in CSV Usage Reports), obtained from the OBSCNTR_IPT_PRODUCTTIER entity.
	RateCardTierdId *string `mandatory:"false" json:"rateCardTierdId"`

	// Rate card ID at the subscribed service level.
	RateCardId *string `mandatory:"false" json:"rateCardId"`

	// SPM (Subscription Plan Management) internal compute records source.
	ComputeSource *string `mandatory:"false" json:"computeSource"`

	// Data center attribute as sent by the meterering system to SPM.
	DataCenter *string `mandatory:"false" json:"dataCenter"`

	// The meterering system identfier sent to SPM. SPM does not transform this attribute and is received as-is.
	MqsMessageId *string `mandatory:"false" json:"mqsMessageId"`

	// Total quantity that was used for computation.
	Quantity *string `mandatory:"false" json:"quantity"`

	// SPM internal usage line number identifier in SPM, originating from the metered services entity.
	UsageNumber *string `mandatory:"false" json:"usageNumber"`

	// SPM internal original usage Line number identifier in SPM, originating from the metered services entity.
	OriginalUsageNumber *string `mandatory:"false" json:"originalUsageNumber"`

	// Subscribed service commitment ID.
	CommitmentServiceId *string `mandatory:"false" json:"commitmentServiceId"`

	// Invoicing status for the aggregated compute usage.
	IsInvoiced *bool `mandatory:"false" json:"isInvoiced"`

	// Usage compute type in SPM.
	Type ComputedUsageSummaryTypeEnum `mandatory:"false" json:"type,omitempty"`

	// Usage computation date, expressed in RFC 3339 timestamp format.
	TimeOfArrival *common.SDKTime `mandatory:"false" json:"timeOfArrival"`

	// Metered service date, expressed in RFC 3339 timestamp format.
	TimeMeteredOn *common.SDKTime `mandatory:"false" json:"timeMeteredOn"`

	// Net unit price for the product in consideration, price actual.
	NetUnitPrice *string `mandatory:"false" json:"netUnitPrice"`

	// Computed rounded line amount.
	CostRounded *string `mandatory:"false" json:"costRounded"`

	// Computed unrounded line amount.
	Cost *string `mandatory:"false" json:"cost"`

	Product *ComputedUsageProduct `mandatory:"false" json:"product"`

	// Unit of measure.
	UnitOfMeasure *string `mandatory:"false" json:"unitOfMeasure"`
}

func (m ComputedUsageSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ComputedUsageSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingComputedUsageSummaryTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetComputedUsageSummaryTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ComputedUsageSummaryTypeEnum Enum with underlying type: string
type ComputedUsageSummaryTypeEnum string

// Set of constants representing the allowable values for ComputedUsageSummaryTypeEnum
const (
	ComputedUsageSummaryTypePromotion                     ComputedUsageSummaryTypeEnum = "PROMOTION"
	ComputedUsageSummaryTypeDoNotBill                     ComputedUsageSummaryTypeEnum = "DO_NOT_BILL"
	ComputedUsageSummaryTypeUsage                         ComputedUsageSummaryTypeEnum = "USAGE"
	ComputedUsageSummaryTypeCommit                        ComputedUsageSummaryTypeEnum = "COMMIT"
	ComputedUsageSummaryTypeOverage                       ComputedUsageSummaryTypeEnum = "OVERAGE"
	ComputedUsageSummaryTypePayAsYouGo                    ComputedUsageSummaryTypeEnum = "PAY_AS_YOU_GO"
	ComputedUsageSummaryTypeMonthlyMinimum                ComputedUsageSummaryTypeEnum = "MONTHLY_MINIMUM"
	ComputedUsageSummaryTypeDelayedUsageInvoiceTiming     ComputedUsageSummaryTypeEnum = "DELAYED_USAGE_INVOICE_TIMING"
	ComputedUsageSummaryTypeDelayedUsageCommitmentExp     ComputedUsageSummaryTypeEnum = "DELAYED_USAGE_COMMITMENT_EXP"
	ComputedUsageSummaryTypeOnAccountCredit               ComputedUsageSummaryTypeEnum = "ON_ACCOUNT_CREDIT"
	ComputedUsageSummaryTypeServiceCredit                 ComputedUsageSummaryTypeEnum = "SERVICE_CREDIT"
	ComputedUsageSummaryTypeCommitmentExpiration          ComputedUsageSummaryTypeEnum = "COMMITMENT_EXPIRATION"
	ComputedUsageSummaryTypeFundedAllocation              ComputedUsageSummaryTypeEnum = "FUNDED_ALLOCATION"
	ComputedUsageSummaryTypeDonotBillUsagePostTermination ComputedUsageSummaryTypeEnum = "DONOT_BILL_USAGE_POST_TERMINATION"
	ComputedUsageSummaryTypeDelayedUsagePostTermination   ComputedUsageSummaryTypeEnum = "DELAYED_USAGE_POST_TERMINATION"
)

var mappingComputedUsageSummaryTypeEnum = map[string]ComputedUsageSummaryTypeEnum{
	"PROMOTION":                         ComputedUsageSummaryTypePromotion,
	"DO_NOT_BILL":                       ComputedUsageSummaryTypeDoNotBill,
	"USAGE":                             ComputedUsageSummaryTypeUsage,
	"COMMIT":                            ComputedUsageSummaryTypeCommit,
	"OVERAGE":                           ComputedUsageSummaryTypeOverage,
	"PAY_AS_YOU_GO":                     ComputedUsageSummaryTypePayAsYouGo,
	"MONTHLY_MINIMUM":                   ComputedUsageSummaryTypeMonthlyMinimum,
	"DELAYED_USAGE_INVOICE_TIMING":      ComputedUsageSummaryTypeDelayedUsageInvoiceTiming,
	"DELAYED_USAGE_COMMITMENT_EXP":      ComputedUsageSummaryTypeDelayedUsageCommitmentExp,
	"ON_ACCOUNT_CREDIT":                 ComputedUsageSummaryTypeOnAccountCredit,
	"SERVICE_CREDIT":                    ComputedUsageSummaryTypeServiceCredit,
	"COMMITMENT_EXPIRATION":             ComputedUsageSummaryTypeCommitmentExpiration,
	"FUNDED_ALLOCATION":                 ComputedUsageSummaryTypeFundedAllocation,
	"DONOT_BILL_USAGE_POST_TERMINATION": ComputedUsageSummaryTypeDonotBillUsagePostTermination,
	"DELAYED_USAGE_POST_TERMINATION":    ComputedUsageSummaryTypeDelayedUsagePostTermination,
}

var mappingComputedUsageSummaryTypeEnumLowerCase = map[string]ComputedUsageSummaryTypeEnum{
	"promotion":                         ComputedUsageSummaryTypePromotion,
	"do_not_bill":                       ComputedUsageSummaryTypeDoNotBill,
	"usage":                             ComputedUsageSummaryTypeUsage,
	"commit":                            ComputedUsageSummaryTypeCommit,
	"overage":                           ComputedUsageSummaryTypeOverage,
	"pay_as_you_go":                     ComputedUsageSummaryTypePayAsYouGo,
	"monthly_minimum":                   ComputedUsageSummaryTypeMonthlyMinimum,
	"delayed_usage_invoice_timing":      ComputedUsageSummaryTypeDelayedUsageInvoiceTiming,
	"delayed_usage_commitment_exp":      ComputedUsageSummaryTypeDelayedUsageCommitmentExp,
	"on_account_credit":                 ComputedUsageSummaryTypeOnAccountCredit,
	"service_credit":                    ComputedUsageSummaryTypeServiceCredit,
	"commitment_expiration":             ComputedUsageSummaryTypeCommitmentExpiration,
	"funded_allocation":                 ComputedUsageSummaryTypeFundedAllocation,
	"donot_bill_usage_post_termination": ComputedUsageSummaryTypeDonotBillUsagePostTermination,
	"delayed_usage_post_termination":    ComputedUsageSummaryTypeDelayedUsagePostTermination,
}

// GetComputedUsageSummaryTypeEnumValues Enumerates the set of values for ComputedUsageSummaryTypeEnum
func GetComputedUsageSummaryTypeEnumValues() []ComputedUsageSummaryTypeEnum {
	values := make([]ComputedUsageSummaryTypeEnum, 0)
	for _, v := range mappingComputedUsageSummaryTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetComputedUsageSummaryTypeEnumStringValues Enumerates the set of values in String for ComputedUsageSummaryTypeEnum
func GetComputedUsageSummaryTypeEnumStringValues() []string {
	return []string{
		"PROMOTION",
		"DO_NOT_BILL",
		"USAGE",
		"COMMIT",
		"OVERAGE",
		"PAY_AS_YOU_GO",
		"MONTHLY_MINIMUM",
		"DELAYED_USAGE_INVOICE_TIMING",
		"DELAYED_USAGE_COMMITMENT_EXP",
		"ON_ACCOUNT_CREDIT",
		"SERVICE_CREDIT",
		"COMMITMENT_EXPIRATION",
		"FUNDED_ALLOCATION",
		"DONOT_BILL_USAGE_POST_TERMINATION",
		"DELAYED_USAGE_POST_TERMINATION",
	}
}

// GetMappingComputedUsageSummaryTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingComputedUsageSummaryTypeEnum(val string) (ComputedUsageSummaryTypeEnum, bool) {
	enum, ok := mappingComputedUsageSummaryTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
