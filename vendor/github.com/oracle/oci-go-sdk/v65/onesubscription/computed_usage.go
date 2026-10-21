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

// ComputedUsage Computed usage summary object.
type ComputedUsage struct {

	// SPM internal computed usage ID, which is a 32-character string.
	Id *string `mandatory:"true" json:"id"`

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

	// Ratecard ID at subscribed service level.
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

	// SPM internal original usage line number identifier in SPM, originating from the metered services entity.
	OriginalUsageNumber *string `mandatory:"false" json:"originalUsageNumber"`

	// Subscribed service commitment ID.
	CommitmentServiceId *string `mandatory:"false" json:"commitmentServiceId"`

	// Invoicing status for the aggregated compute usage.
	IsInvoiced *bool `mandatory:"false" json:"isInvoiced"`

	// Usage compute type in SPM.
	Type ComputedUsageTypeEnum `mandatory:"false" json:"type,omitempty"`

	// Usage computation date, expressed in RFC 3339 timestamp format.
	TimeOfArrival *common.SDKTime `mandatory:"false" json:"timeOfArrival"`

	// Metered service date, expressed in RFC 3339 timestamp format.
	TimeMeteredOn *common.SDKTime `mandatory:"false" json:"timeMeteredOn"`

	// Net unit price for the product in consideration, price actual.
	NetUnitPrice *string `mandatory:"false" json:"netUnitPrice"`

	// Computed rounded line amount.
	CostRounded *string `mandatory:"false" json:"costRounded"`

	// Computed unrounded line Amount.
	Cost *string `mandatory:"false" json:"cost"`

	Product *ComputedUsageProduct `mandatory:"false" json:"product"`

	// Unit of measure.
	UnitOfMeasure *string `mandatory:"false" json:"unitOfMeasure"`
}

func (m ComputedUsage) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ComputedUsage) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingComputedUsageTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetComputedUsageTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ComputedUsageTypeEnum Enum with underlying type: string
type ComputedUsageTypeEnum string

// Set of constants representing the allowable values for ComputedUsageTypeEnum
const (
	ComputedUsageTypePromotion                     ComputedUsageTypeEnum = "PROMOTION"
	ComputedUsageTypeDoNotBill                     ComputedUsageTypeEnum = "DO_NOT_BILL"
	ComputedUsageTypeUsage                         ComputedUsageTypeEnum = "USAGE"
	ComputedUsageTypeCommit                        ComputedUsageTypeEnum = "COMMIT"
	ComputedUsageTypeOverage                       ComputedUsageTypeEnum = "OVERAGE"
	ComputedUsageTypePayAsYouGo                    ComputedUsageTypeEnum = "PAY_AS_YOU_GO"
	ComputedUsageTypeMonthlyMinimum                ComputedUsageTypeEnum = "MONTHLY_MINIMUM"
	ComputedUsageTypeDelayedUsageInvoiceTiming     ComputedUsageTypeEnum = "DELAYED_USAGE_INVOICE_TIMING"
	ComputedUsageTypeDelayedUsageCommitmentExp     ComputedUsageTypeEnum = "DELAYED_USAGE_COMMITMENT_EXP"
	ComputedUsageTypeOnAccountCredit               ComputedUsageTypeEnum = "ON_ACCOUNT_CREDIT"
	ComputedUsageTypeServiceCredit                 ComputedUsageTypeEnum = "SERVICE_CREDIT"
	ComputedUsageTypeCommitmentExpiration          ComputedUsageTypeEnum = "COMMITMENT_EXPIRATION"
	ComputedUsageTypeFundedAllocation              ComputedUsageTypeEnum = "FUNDED_ALLOCATION"
	ComputedUsageTypeDonotBillUsagePostTermination ComputedUsageTypeEnum = "DONOT_BILL_USAGE_POST_TERMINATION"
	ComputedUsageTypeDelayedUsagePostTermination   ComputedUsageTypeEnum = "DELAYED_USAGE_POST_TERMINATION"
)

var mappingComputedUsageTypeEnum = map[string]ComputedUsageTypeEnum{
	"PROMOTION":                         ComputedUsageTypePromotion,
	"DO_NOT_BILL":                       ComputedUsageTypeDoNotBill,
	"USAGE":                             ComputedUsageTypeUsage,
	"COMMIT":                            ComputedUsageTypeCommit,
	"OVERAGE":                           ComputedUsageTypeOverage,
	"PAY_AS_YOU_GO":                     ComputedUsageTypePayAsYouGo,
	"MONTHLY_MINIMUM":                   ComputedUsageTypeMonthlyMinimum,
	"DELAYED_USAGE_INVOICE_TIMING":      ComputedUsageTypeDelayedUsageInvoiceTiming,
	"DELAYED_USAGE_COMMITMENT_EXP":      ComputedUsageTypeDelayedUsageCommitmentExp,
	"ON_ACCOUNT_CREDIT":                 ComputedUsageTypeOnAccountCredit,
	"SERVICE_CREDIT":                    ComputedUsageTypeServiceCredit,
	"COMMITMENT_EXPIRATION":             ComputedUsageTypeCommitmentExpiration,
	"FUNDED_ALLOCATION":                 ComputedUsageTypeFundedAllocation,
	"DONOT_BILL_USAGE_POST_TERMINATION": ComputedUsageTypeDonotBillUsagePostTermination,
	"DELAYED_USAGE_POST_TERMINATION":    ComputedUsageTypeDelayedUsagePostTermination,
}

var mappingComputedUsageTypeEnumLowerCase = map[string]ComputedUsageTypeEnum{
	"promotion":                         ComputedUsageTypePromotion,
	"do_not_bill":                       ComputedUsageTypeDoNotBill,
	"usage":                             ComputedUsageTypeUsage,
	"commit":                            ComputedUsageTypeCommit,
	"overage":                           ComputedUsageTypeOverage,
	"pay_as_you_go":                     ComputedUsageTypePayAsYouGo,
	"monthly_minimum":                   ComputedUsageTypeMonthlyMinimum,
	"delayed_usage_invoice_timing":      ComputedUsageTypeDelayedUsageInvoiceTiming,
	"delayed_usage_commitment_exp":      ComputedUsageTypeDelayedUsageCommitmentExp,
	"on_account_credit":                 ComputedUsageTypeOnAccountCredit,
	"service_credit":                    ComputedUsageTypeServiceCredit,
	"commitment_expiration":             ComputedUsageTypeCommitmentExpiration,
	"funded_allocation":                 ComputedUsageTypeFundedAllocation,
	"donot_bill_usage_post_termination": ComputedUsageTypeDonotBillUsagePostTermination,
	"delayed_usage_post_termination":    ComputedUsageTypeDelayedUsagePostTermination,
}

// GetComputedUsageTypeEnumValues Enumerates the set of values for ComputedUsageTypeEnum
func GetComputedUsageTypeEnumValues() []ComputedUsageTypeEnum {
	values := make([]ComputedUsageTypeEnum, 0)
	for _, v := range mappingComputedUsageTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetComputedUsageTypeEnumStringValues Enumerates the set of values in String for ComputedUsageTypeEnum
func GetComputedUsageTypeEnumStringValues() []string {
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

// GetMappingComputedUsageTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingComputedUsageTypeEnum(val string) (ComputedUsageTypeEnum, bool) {
	enum, ok := mappingComputedUsageTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
