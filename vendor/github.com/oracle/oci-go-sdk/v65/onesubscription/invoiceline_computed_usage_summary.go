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

// InvoicelineComputedUsageSummary Computed Usage Summary object
type InvoicelineComputedUsageSummary struct {
	ParentProduct *InvoicingProduct `mandatory:"true" json:"parentProduct"`

	// Total Quantity that was used for computation
	Quantity *float64 `mandatory:"true" json:"quantity"`

	// Net Unit Price for the product in consideration, price actual.
	NetUnitPrice *float64 `mandatory:"true" json:"netUnitPrice"`

	// Metered Service date.
	TimeMeteredOn *common.SDKTime `mandatory:"true" json:"timeMeteredOn"`

	// Usage compute type in SPM.
	Type InvoicelineComputedUsageSummaryTypeEnum `mandatory:"true" json:"type"`

	// Computed Line Amount rounded.
	CostRounded *float64 `mandatory:"true" json:"costRounded"`

	Product *InvoicingProduct `mandatory:"false" json:"product"`

	// Sum of Usage/Service Billing Line net Amount
	Cost *float64 `mandatory:"false" json:"cost"`
}

func (m InvoicelineComputedUsageSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InvoicelineComputedUsageSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingInvoicelineComputedUsageSummaryTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetInvoicelineComputedUsageSummaryTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// InvoicelineComputedUsageSummaryTypeEnum Enum with underlying type: string
type InvoicelineComputedUsageSummaryTypeEnum string

// Set of constants representing the allowable values for InvoicelineComputedUsageSummaryTypeEnum
const (
	InvoicelineComputedUsageSummaryTypePromotion                 InvoicelineComputedUsageSummaryTypeEnum = "PROMOTION"
	InvoicelineComputedUsageSummaryTypeDoNotBill                 InvoicelineComputedUsageSummaryTypeEnum = "DO_NOT_BILL"
	InvoicelineComputedUsageSummaryTypeUsage                     InvoicelineComputedUsageSummaryTypeEnum = "USAGE"
	InvoicelineComputedUsageSummaryTypeCommit                    InvoicelineComputedUsageSummaryTypeEnum = "COMMIT"
	InvoicelineComputedUsageSummaryTypeOverage                   InvoicelineComputedUsageSummaryTypeEnum = "OVERAGE"
	InvoicelineComputedUsageSummaryTypePayAsYouGo                InvoicelineComputedUsageSummaryTypeEnum = "PAY_AS_YOU_GO"
	InvoicelineComputedUsageSummaryTypeMonthlyMinimum            InvoicelineComputedUsageSummaryTypeEnum = "MONTHLY_MINIMUM"
	InvoicelineComputedUsageSummaryTypeDelayedUsageInvoiceTiming InvoicelineComputedUsageSummaryTypeEnum = "DELAYED_USAGE_INVOICE_TIMING"
	InvoicelineComputedUsageSummaryTypeDelayedUsageCommitmentExp InvoicelineComputedUsageSummaryTypeEnum = "DELAYED_USAGE_COMMITMENT_EXP"
	InvoicelineComputedUsageSummaryTypeOnAccountCredit           InvoicelineComputedUsageSummaryTypeEnum = "ON_ACCOUNT_CREDIT"
	InvoicelineComputedUsageSummaryTypeServiceCredit             InvoicelineComputedUsageSummaryTypeEnum = "SERVICE_CREDIT"
)

var mappingInvoicelineComputedUsageSummaryTypeEnum = map[string]InvoicelineComputedUsageSummaryTypeEnum{
	"PROMOTION":                    InvoicelineComputedUsageSummaryTypePromotion,
	"DO_NOT_BILL":                  InvoicelineComputedUsageSummaryTypeDoNotBill,
	"USAGE":                        InvoicelineComputedUsageSummaryTypeUsage,
	"COMMIT":                       InvoicelineComputedUsageSummaryTypeCommit,
	"OVERAGE":                      InvoicelineComputedUsageSummaryTypeOverage,
	"PAY_AS_YOU_GO":                InvoicelineComputedUsageSummaryTypePayAsYouGo,
	"MONTHLY_MINIMUM":              InvoicelineComputedUsageSummaryTypeMonthlyMinimum,
	"DELAYED_USAGE_INVOICE_TIMING": InvoicelineComputedUsageSummaryTypeDelayedUsageInvoiceTiming,
	"DELAYED_USAGE_COMMITMENT_EXP": InvoicelineComputedUsageSummaryTypeDelayedUsageCommitmentExp,
	"ON_ACCOUNT_CREDIT":            InvoicelineComputedUsageSummaryTypeOnAccountCredit,
	"SERVICE_CREDIT":               InvoicelineComputedUsageSummaryTypeServiceCredit,
}

var mappingInvoicelineComputedUsageSummaryTypeEnumLowerCase = map[string]InvoicelineComputedUsageSummaryTypeEnum{
	"promotion":                    InvoicelineComputedUsageSummaryTypePromotion,
	"do_not_bill":                  InvoicelineComputedUsageSummaryTypeDoNotBill,
	"usage":                        InvoicelineComputedUsageSummaryTypeUsage,
	"commit":                       InvoicelineComputedUsageSummaryTypeCommit,
	"overage":                      InvoicelineComputedUsageSummaryTypeOverage,
	"pay_as_you_go":                InvoicelineComputedUsageSummaryTypePayAsYouGo,
	"monthly_minimum":              InvoicelineComputedUsageSummaryTypeMonthlyMinimum,
	"delayed_usage_invoice_timing": InvoicelineComputedUsageSummaryTypeDelayedUsageInvoiceTiming,
	"delayed_usage_commitment_exp": InvoicelineComputedUsageSummaryTypeDelayedUsageCommitmentExp,
	"on_account_credit":            InvoicelineComputedUsageSummaryTypeOnAccountCredit,
	"service_credit":               InvoicelineComputedUsageSummaryTypeServiceCredit,
}

// GetInvoicelineComputedUsageSummaryTypeEnumValues Enumerates the set of values for InvoicelineComputedUsageSummaryTypeEnum
func GetInvoicelineComputedUsageSummaryTypeEnumValues() []InvoicelineComputedUsageSummaryTypeEnum {
	values := make([]InvoicelineComputedUsageSummaryTypeEnum, 0)
	for _, v := range mappingInvoicelineComputedUsageSummaryTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetInvoicelineComputedUsageSummaryTypeEnumStringValues Enumerates the set of values in String for InvoicelineComputedUsageSummaryTypeEnum
func GetInvoicelineComputedUsageSummaryTypeEnumStringValues() []string {
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
	}
}

// GetMappingInvoicelineComputedUsageSummaryTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInvoicelineComputedUsageSummaryTypeEnum(val string) (InvoicelineComputedUsageSummaryTypeEnum, bool) {
	enum, ok := mappingInvoicelineComputedUsageSummaryTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
