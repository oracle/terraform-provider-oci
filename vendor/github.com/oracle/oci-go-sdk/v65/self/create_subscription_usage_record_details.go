// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// SELF Service API
//
// Use the SELF Service API to manage Subscriptions in Oracle Cloud Infrastructure Marketplace. For more information, see Overview of Marketplace (https://docs.oracle.com/iaas/Content/Marketplace/Concepts/marketoverview.htm)
//

package self

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateSubscriptionUsageRecordDetails A single usage record to submit for a marketplace offer. The usage window must have
// `timeUsageStarted` before `timeUsageEnded`.
type CreateSubscriptionUsageRecordDetails struct {

	// Partner-provided unique identifier for this usage record. This identifier must be unique per partner.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the marketplace offer being billed. A submit request can contain records for multiple marketplace offers.
	MarketplaceOfferId *string `mandatory:"true" json:"marketplaceOfferId"`

	// The pricing dimension against which usage is being reported.
	UsageDimensionName *string `mandatory:"true" json:"usageDimensionName"`

	// The non-negative usage cost computed by the partner for the submitted usage quantity.
	Amount *float64 `mandatory:"true" json:"amount"`

	// The ISO-4217 currency submitted for the computed usage cost.
	CurrencyCode *string `mandatory:"true" json:"currencyCode"`

	// The inclusive start timestamp for the usage window. This value must be before `timeUsageEnded`.
	TimeUsageStarted *common.SDKTime `mandatory:"true" json:"timeUsageStarted"`

	// The exclusive end timestamp for the usage window. This value must be after `timeUsageStarted`.
	TimeUsageEnded *common.SDKTime `mandatory:"true" json:"timeUsageEnded"`

	// The billing type this usage record applies to.
	BillingType CreateSubscriptionUsageRecordDetailsBillingTypeEnum `mandatory:"false" json:"billingType,omitempty"`

	// An optional non-negative usage quantity being reported.
	ConsumedQuantity *float64 `mandatory:"false" json:"consumedQuantity"`

	// The billing period associated with this usage record in `YYYY-MM` format.
	BillingPeriod *string `mandatory:"false" json:"billingPeriod"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the customer tenancy associated with this usage record.
	CustomerTenancyId *string `mandatory:"false" json:"customerTenancyId"`

	// The partner billing identifier associated with this usage record.
	BillingIdentifier *string `mandatory:"false" json:"billingIdentifier"`

	// The unit of measure associated with the reported usage quantity.
	UnitOfMeasure *string `mandatory:"false" json:"unitOfMeasure"`

	// The unit price associated with this usage record.
	UnitPrice *float64 `mandatory:"false" json:"unitPrice"`

	// The product SKU associated with this usage record.
	ProductSku *string `mandatory:"false" json:"productSku"`

	// The contract duration associated with this usage record.
	ContractDuration BillingFrequencyEnum `mandatory:"false" json:"contractDuration,omitempty"`

	// Additional key/value metadata associated with this usage record for extensibility.
	AdditionalMetadata []ExtendedMetadata `mandatory:"false" json:"additionalMetadata"`

	// Partner-provided usage record identifier for traceability.
	UsageRecordId *string `mandatory:"false" json:"usageRecordId"`
}

func (m CreateSubscriptionUsageRecordDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateSubscriptionUsageRecordDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreateSubscriptionUsageRecordDetailsBillingTypeEnum(string(m.BillingType)); !ok && m.BillingType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BillingType: %s. Supported values are: %s.", m.BillingType, strings.Join(GetCreateSubscriptionUsageRecordDetailsBillingTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBillingFrequencyEnum(string(m.ContractDuration)); !ok && m.ContractDuration != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ContractDuration: %s. Supported values are: %s.", m.ContractDuration, strings.Join(GetBillingFrequencyEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateSubscriptionUsageRecordDetailsBillingTypeEnum Enum with underlying type: string
type CreateSubscriptionUsageRecordDetailsBillingTypeEnum string

// Set of constants representing the allowable values for CreateSubscriptionUsageRecordDetailsBillingTypeEnum
const (
	CreateSubscriptionUsageRecordDetailsBillingTypeFlatRate   CreateSubscriptionUsageRecordDetailsBillingTypeEnum = "FLAT_RATE"
	CreateSubscriptionUsageRecordDetailsBillingTypeUsageBased CreateSubscriptionUsageRecordDetailsBillingTypeEnum = "USAGE_BASED"
)

var mappingCreateSubscriptionUsageRecordDetailsBillingTypeEnum = map[string]CreateSubscriptionUsageRecordDetailsBillingTypeEnum{
	"FLAT_RATE":   CreateSubscriptionUsageRecordDetailsBillingTypeFlatRate,
	"USAGE_BASED": CreateSubscriptionUsageRecordDetailsBillingTypeUsageBased,
}

var mappingCreateSubscriptionUsageRecordDetailsBillingTypeEnumLowerCase = map[string]CreateSubscriptionUsageRecordDetailsBillingTypeEnum{
	"flat_rate":   CreateSubscriptionUsageRecordDetailsBillingTypeFlatRate,
	"usage_based": CreateSubscriptionUsageRecordDetailsBillingTypeUsageBased,
}

// GetCreateSubscriptionUsageRecordDetailsBillingTypeEnumValues Enumerates the set of values for CreateSubscriptionUsageRecordDetailsBillingTypeEnum
func GetCreateSubscriptionUsageRecordDetailsBillingTypeEnumValues() []CreateSubscriptionUsageRecordDetailsBillingTypeEnum {
	values := make([]CreateSubscriptionUsageRecordDetailsBillingTypeEnum, 0)
	for _, v := range mappingCreateSubscriptionUsageRecordDetailsBillingTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateSubscriptionUsageRecordDetailsBillingTypeEnumStringValues Enumerates the set of values in String for CreateSubscriptionUsageRecordDetailsBillingTypeEnum
func GetCreateSubscriptionUsageRecordDetailsBillingTypeEnumStringValues() []string {
	return []string{
		"FLAT_RATE",
		"USAGE_BASED",
	}
}

// GetMappingCreateSubscriptionUsageRecordDetailsBillingTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateSubscriptionUsageRecordDetailsBillingTypeEnum(val string) (CreateSubscriptionUsageRecordDetailsBillingTypeEnum, bool) {
	enum, ok := mappingCreateSubscriptionUsageRecordDetailsBillingTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
