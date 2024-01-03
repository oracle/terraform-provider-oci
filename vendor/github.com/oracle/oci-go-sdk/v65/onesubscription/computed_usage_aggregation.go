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

// ComputedUsageAggregation Computed Usage Aggregation object
type ComputedUsageAggregation struct {

	// Total Quantity that was used for computation
	Quantity *string `mandatory:"false" json:"quantity"`

	Product *ComputedUsageProduct `mandatory:"false" json:"product"`

	// Data Center Attribute as sent by MQS to SPM.
	DataCenter *string `mandatory:"false" json:"dataCenter"`

	// Metered Service date , expressed in RFC 3339 timestamp format.
	TimeMeteredOn *common.SDKTime `mandatory:"false" json:"timeMeteredOn"`

	// Net Unit Price for the product in consideration.
	NetUnitPrice *string `mandatory:"false" json:"netUnitPrice"`

	// Sum of Computed Line Amount unrounded
	CostUnrounded *string `mandatory:"false" json:"costUnrounded"`

	// Sum of Computed Line Amount rounded
	Cost *string `mandatory:"false" json:"cost"`

	// Usage compute type in SPM.
	Type ComputedUsageAggregationTypeEnum `mandatory:"false" json:"type,omitempty"`
}

func (m ComputedUsageAggregation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ComputedUsageAggregation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingComputedUsageAggregationTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetComputedUsageAggregationTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ComputedUsageAggregationTypeEnum Enum with underlying type: string
type ComputedUsageAggregationTypeEnum string

// Set of constants representing the allowable values for ComputedUsageAggregationTypeEnum
const (
	ComputedUsageAggregationTypePromotion                     ComputedUsageAggregationTypeEnum = "PROMOTION"
	ComputedUsageAggregationTypeDoNotBill                     ComputedUsageAggregationTypeEnum = "DO_NOT_BILL"
	ComputedUsageAggregationTypeUsage                         ComputedUsageAggregationTypeEnum = "USAGE"
	ComputedUsageAggregationTypeCommit                        ComputedUsageAggregationTypeEnum = "COMMIT"
	ComputedUsageAggregationTypeOverage                       ComputedUsageAggregationTypeEnum = "OVERAGE"
	ComputedUsageAggregationTypePayAsYouGo                    ComputedUsageAggregationTypeEnum = "PAY_AS_YOU_GO"
	ComputedUsageAggregationTypeMonthlyMinimum                ComputedUsageAggregationTypeEnum = "MONTHLY_MINIMUM"
	ComputedUsageAggregationTypeDelayedUsageInvoiceTiming     ComputedUsageAggregationTypeEnum = "DELAYED_USAGE_INVOICE_TIMING"
	ComputedUsageAggregationTypeDelayedUsageCommitmentExp     ComputedUsageAggregationTypeEnum = "DELAYED_USAGE_COMMITMENT_EXP"
	ComputedUsageAggregationTypeOnAccountCredit               ComputedUsageAggregationTypeEnum = "ON_ACCOUNT_CREDIT"
	ComputedUsageAggregationTypeServiceCredit                 ComputedUsageAggregationTypeEnum = "SERVICE_CREDIT"
	ComputedUsageAggregationTypeCommitmentExpiration          ComputedUsageAggregationTypeEnum = "COMMITMENT_EXPIRATION"
	ComputedUsageAggregationTypeFundedAllocation              ComputedUsageAggregationTypeEnum = "FUNDED_ALLOCATION"
	ComputedUsageAggregationTypeDonotBillUsagePostTermination ComputedUsageAggregationTypeEnum = "DONOT_BILL_USAGE_POST_TERMINATION"
	ComputedUsageAggregationTypeDelayedUsagePostTermination   ComputedUsageAggregationTypeEnum = "DELAYED_USAGE_POST_TERMINATION"
)

var mappingComputedUsageAggregationTypeEnum = map[string]ComputedUsageAggregationTypeEnum{
	"PROMOTION":                         ComputedUsageAggregationTypePromotion,
	"DO_NOT_BILL":                       ComputedUsageAggregationTypeDoNotBill,
	"USAGE":                             ComputedUsageAggregationTypeUsage,
	"COMMIT":                            ComputedUsageAggregationTypeCommit,
	"OVERAGE":                           ComputedUsageAggregationTypeOverage,
	"PAY_AS_YOU_GO":                     ComputedUsageAggregationTypePayAsYouGo,
	"MONTHLY_MINIMUM":                   ComputedUsageAggregationTypeMonthlyMinimum,
	"DELAYED_USAGE_INVOICE_TIMING":      ComputedUsageAggregationTypeDelayedUsageInvoiceTiming,
	"DELAYED_USAGE_COMMITMENT_EXP":      ComputedUsageAggregationTypeDelayedUsageCommitmentExp,
	"ON_ACCOUNT_CREDIT":                 ComputedUsageAggregationTypeOnAccountCredit,
	"SERVICE_CREDIT":                    ComputedUsageAggregationTypeServiceCredit,
	"COMMITMENT_EXPIRATION":             ComputedUsageAggregationTypeCommitmentExpiration,
	"FUNDED_ALLOCATION":                 ComputedUsageAggregationTypeFundedAllocation,
	"DONOT_BILL_USAGE_POST_TERMINATION": ComputedUsageAggregationTypeDonotBillUsagePostTermination,
	"DELAYED_USAGE_POST_TERMINATION":    ComputedUsageAggregationTypeDelayedUsagePostTermination,
}

var mappingComputedUsageAggregationTypeEnumLowerCase = map[string]ComputedUsageAggregationTypeEnum{
	"promotion":                         ComputedUsageAggregationTypePromotion,
	"do_not_bill":                       ComputedUsageAggregationTypeDoNotBill,
	"usage":                             ComputedUsageAggregationTypeUsage,
	"commit":                            ComputedUsageAggregationTypeCommit,
	"overage":                           ComputedUsageAggregationTypeOverage,
	"pay_as_you_go":                     ComputedUsageAggregationTypePayAsYouGo,
	"monthly_minimum":                   ComputedUsageAggregationTypeMonthlyMinimum,
	"delayed_usage_invoice_timing":      ComputedUsageAggregationTypeDelayedUsageInvoiceTiming,
	"delayed_usage_commitment_exp":      ComputedUsageAggregationTypeDelayedUsageCommitmentExp,
	"on_account_credit":                 ComputedUsageAggregationTypeOnAccountCredit,
	"service_credit":                    ComputedUsageAggregationTypeServiceCredit,
	"commitment_expiration":             ComputedUsageAggregationTypeCommitmentExpiration,
	"funded_allocation":                 ComputedUsageAggregationTypeFundedAllocation,
	"donot_bill_usage_post_termination": ComputedUsageAggregationTypeDonotBillUsagePostTermination,
	"delayed_usage_post_termination":    ComputedUsageAggregationTypeDelayedUsagePostTermination,
}

// GetComputedUsageAggregationTypeEnumValues Enumerates the set of values for ComputedUsageAggregationTypeEnum
func GetComputedUsageAggregationTypeEnumValues() []ComputedUsageAggregationTypeEnum {
	values := make([]ComputedUsageAggregationTypeEnum, 0)
	for _, v := range mappingComputedUsageAggregationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetComputedUsageAggregationTypeEnumStringValues Enumerates the set of values in String for ComputedUsageAggregationTypeEnum
func GetComputedUsageAggregationTypeEnumStringValues() []string {
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

// GetMappingComputedUsageAggregationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingComputedUsageAggregationTypeEnum(val string) (ComputedUsageAggregationTypeEnum, bool) {
	enum, ok := mappingComputedUsageAggregationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
