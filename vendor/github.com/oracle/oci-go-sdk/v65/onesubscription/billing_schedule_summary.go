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

// BillingScheduleSummary Billing schedule details related to Subscription Id
type BillingScheduleSummary struct {

	// SPM internal Subscribed Service ID
	SubscribedServiceId *string `mandatory:"false" json:"subscribedServiceId"`

	// Billing schedule start date
	TimeStart *common.SDKTime `mandatory:"false" json:"timeStart"`

	// Billing schedule end date
	TimeEnd *common.SDKTime `mandatory:"false" json:"timeEnd"`

	// Billing schedule invoicing date
	TimeInvoicing *common.SDKTime `mandatory:"false" json:"timeInvoicing"`

	// Billing schedule invoice status
	InvoiceStatus BillingScheduleSummaryInvoiceStatusEnum `mandatory:"false" json:"invoiceStatus,omitempty"`

	// Billing schedule quantity
	Quantity *string `mandatory:"false" json:"quantity"`

	// Billing schedule net unit price
	NetUnitPrice *string `mandatory:"false" json:"netUnitPrice"`

	// Billing schedule line net amount
	Amount *string `mandatory:"false" json:"amount"`

	// Billing frequency
	BillingFrequency *string `mandatory:"false" json:"billingFrequency"`

	// Indicates the associated AR Invoice Number
	ArInvoiceNumber *string `mandatory:"false" json:"arInvoiceNumber"`

	// Indicates the associated AR Customer transaction id a unique identifier existing on AR.
	ArCustomerTransactionId *string `mandatory:"false" json:"arCustomerTransactionId"`

	// Order number associated with the Subscribed Service
	OrderNumber *string `mandatory:"false" json:"orderNumber"`

	Product *BillingScheduleProduct `mandatory:"false" json:"product"`
}

func (m BillingScheduleSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BillingScheduleSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingBillingScheduleSummaryInvoiceStatusEnum(string(m.InvoiceStatus)); !ok && m.InvoiceStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InvoiceStatus: %s. Supported values are: %s.", m.InvoiceStatus, strings.Join(GetBillingScheduleSummaryInvoiceStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BillingScheduleSummaryInvoiceStatusEnum Enum with underlying type: string
type BillingScheduleSummaryInvoiceStatusEnum string

// Set of constants representing the allowable values for BillingScheduleSummaryInvoiceStatusEnum
const (
	BillingScheduleSummaryInvoiceStatusInvoiced    BillingScheduleSummaryInvoiceStatusEnum = "INVOICED"
	BillingScheduleSummaryInvoiceStatusNotInvoiced BillingScheduleSummaryInvoiceStatusEnum = "NOT_INVOICED"
)

var mappingBillingScheduleSummaryInvoiceStatusEnum = map[string]BillingScheduleSummaryInvoiceStatusEnum{
	"INVOICED":     BillingScheduleSummaryInvoiceStatusInvoiced,
	"NOT_INVOICED": BillingScheduleSummaryInvoiceStatusNotInvoiced,
}

var mappingBillingScheduleSummaryInvoiceStatusEnumLowerCase = map[string]BillingScheduleSummaryInvoiceStatusEnum{
	"invoiced":     BillingScheduleSummaryInvoiceStatusInvoiced,
	"not_invoiced": BillingScheduleSummaryInvoiceStatusNotInvoiced,
}

// GetBillingScheduleSummaryInvoiceStatusEnumValues Enumerates the set of values for BillingScheduleSummaryInvoiceStatusEnum
func GetBillingScheduleSummaryInvoiceStatusEnumValues() []BillingScheduleSummaryInvoiceStatusEnum {
	values := make([]BillingScheduleSummaryInvoiceStatusEnum, 0)
	for _, v := range mappingBillingScheduleSummaryInvoiceStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetBillingScheduleSummaryInvoiceStatusEnumStringValues Enumerates the set of values in String for BillingScheduleSummaryInvoiceStatusEnum
func GetBillingScheduleSummaryInvoiceStatusEnumStringValues() []string {
	return []string{
		"INVOICED",
		"NOT_INVOICED",
	}
}

// GetMappingBillingScheduleSummaryInvoiceStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBillingScheduleSummaryInvoiceStatusEnum(val string) (BillingScheduleSummaryInvoiceStatusEnum, bool) {
	enum, ok := mappingBillingScheduleSummaryInvoiceStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
