// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OSP Gateway API
//
// This site describes all the Rest endpoints of OSP Gateway.
//

package ospgateway

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InvoiceSummary Invoice list elements
type InvoiceSummary struct {

	// Invoice identifier
	InvoiceId *string `mandatory:"true" json:"invoiceId"`

	// Invoice external reference
	InvoiceNumber *string `mandatory:"false" json:"invoiceNumber"`

	// PC invoice identifier
	InternalInvoiceId *string `mandatory:"false" json:"internalInvoiceId"`

	// Is credit card payment eligible
	IsCreditCardPayable *bool `mandatory:"false" json:"isCreditCardPayable"`

	// Invoice status
	InvoiceStatus InvoiceSummaryInvoiceStatusEnum `mandatory:"false" json:"invoiceStatus,omitempty"`

	// Type of invoice
	InvoiceType InvoiceSummaryInvoiceTypeEnum `mandatory:"false" json:"invoiceType,omitempty"`

	// Is the invoice has been already payed
	IsPaid *bool `mandatory:"false" json:"isPaid"`

	// Whether invoice can be payed
	IsPayable *bool `mandatory:"false" json:"isPayable"`

	// Invoice amount
	InvoiceAmount *float32 `mandatory:"false" json:"invoiceAmount"`

	// Invoice amount due
	InvoiceAmountDue *float32 `mandatory:"false" json:"invoiceAmountDue"`

	// Invoice amount credit
	InvoiceAmountCredited *float32 `mandatory:"false" json:"invoiceAmountCredited"`

	// Invoice amount adjust
	InvoiceAmountAdjusted *float32 `mandatory:"false" json:"invoiceAmountAdjusted"`

	// Invoice amount applied
	InvoiceAmountApplied *float32 `mandatory:"false" json:"invoiceAmountApplied"`

	// Due date of invoice amount
	TimeInvoiceDue *common.SDKTime `mandatory:"false" json:"timeInvoiceDue"`

	// Is the last payment failed
	IsPaymentFailed *bool `mandatory:"false" json:"isPaymentFailed"`

	// Invoice amount in dispute
	InvoiceAmountInDispute *float32 `mandatory:"false" json:"invoiceAmountInDispute"`

	// Invoice reference number
	InvoiceRefNumber *string `mandatory:"false" json:"invoiceRefNumber"`

	// Invoice PO number
	InvoicePoNumber *string `mandatory:"false" json:"invoicePoNumber"`

	// Date of invoice
	TimeInvoice *common.SDKTime `mandatory:"false" json:"timeInvoice"`

	Currency *Currency `mandatory:"false" json:"currency"`

	// Is emailing pdf allowed
	IsPdfEmailAvailable *bool `mandatory:"false" json:"isPdfEmailAvailable"`

	// Is view access allowed
	IsDisplayViewPdf *bool `mandatory:"false" json:"isDisplayViewPdf"`

	// Is pdf download access allowed
	IsDisplayDownloadPdf *bool `mandatory:"false" json:"isDisplayDownloadPdf"`

	LastPaymentDetail PaymentDetail `mandatory:"false" json:"lastPaymentDetail"`

	// Name of the bill to customer
	PartyName *string `mandatory:"false" json:"partyName"`

	// List of subscription identifiers
	SubscriptionIds []string `mandatory:"false" json:"subscriptionIds"`
}

func (m InvoiceSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InvoiceSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingInvoiceSummaryInvoiceStatusEnum(string(m.InvoiceStatus)); !ok && m.InvoiceStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InvoiceStatus: %s. Supported values are: %s.", m.InvoiceStatus, strings.Join(GetInvoiceSummaryInvoiceStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingInvoiceSummaryInvoiceTypeEnum(string(m.InvoiceType)); !ok && m.InvoiceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InvoiceType: %s. Supported values are: %s.", m.InvoiceType, strings.Join(GetInvoiceSummaryInvoiceTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *InvoiceSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		InvoiceNumber          *string                         `json:"invoiceNumber"`
		InternalInvoiceId      *string                         `json:"internalInvoiceId"`
		IsCreditCardPayable    *bool                           `json:"isCreditCardPayable"`
		InvoiceStatus          InvoiceSummaryInvoiceStatusEnum `json:"invoiceStatus"`
		InvoiceType            InvoiceSummaryInvoiceTypeEnum   `json:"invoiceType"`
		IsPaid                 *bool                           `json:"isPaid"`
		IsPayable              *bool                           `json:"isPayable"`
		InvoiceAmount          *float32                        `json:"invoiceAmount"`
		InvoiceAmountDue       *float32                        `json:"invoiceAmountDue"`
		InvoiceAmountCredited  *float32                        `json:"invoiceAmountCredited"`
		InvoiceAmountAdjusted  *float32                        `json:"invoiceAmountAdjusted"`
		InvoiceAmountApplied   *float32                        `json:"invoiceAmountApplied"`
		TimeInvoiceDue         *common.SDKTime                 `json:"timeInvoiceDue"`
		IsPaymentFailed        *bool                           `json:"isPaymentFailed"`
		InvoiceAmountInDispute *float32                        `json:"invoiceAmountInDispute"`
		InvoiceRefNumber       *string                         `json:"invoiceRefNumber"`
		InvoicePoNumber        *string                         `json:"invoicePoNumber"`
		TimeInvoice            *common.SDKTime                 `json:"timeInvoice"`
		Currency               *Currency                       `json:"currency"`
		IsPdfEmailAvailable    *bool                           `json:"isPdfEmailAvailable"`
		IsDisplayViewPdf       *bool                           `json:"isDisplayViewPdf"`
		IsDisplayDownloadPdf   *bool                           `json:"isDisplayDownloadPdf"`
		LastPaymentDetail      paymentdetail                   `json:"lastPaymentDetail"`
		PartyName              *string                         `json:"partyName"`
		SubscriptionIds        []string                        `json:"subscriptionIds"`
		InvoiceId              *string                         `json:"invoiceId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.InvoiceNumber = model.InvoiceNumber

	m.InternalInvoiceId = model.InternalInvoiceId

	m.IsCreditCardPayable = model.IsCreditCardPayable

	m.InvoiceStatus = model.InvoiceStatus

	m.InvoiceType = model.InvoiceType

	m.IsPaid = model.IsPaid

	m.IsPayable = model.IsPayable

	m.InvoiceAmount = model.InvoiceAmount

	m.InvoiceAmountDue = model.InvoiceAmountDue

	m.InvoiceAmountCredited = model.InvoiceAmountCredited

	m.InvoiceAmountAdjusted = model.InvoiceAmountAdjusted

	m.InvoiceAmountApplied = model.InvoiceAmountApplied

	m.TimeInvoiceDue = model.TimeInvoiceDue

	m.IsPaymentFailed = model.IsPaymentFailed

	m.InvoiceAmountInDispute = model.InvoiceAmountInDispute

	m.InvoiceRefNumber = model.InvoiceRefNumber

	m.InvoicePoNumber = model.InvoicePoNumber

	m.TimeInvoice = model.TimeInvoice

	m.Currency = model.Currency

	m.IsPdfEmailAvailable = model.IsPdfEmailAvailable

	m.IsDisplayViewPdf = model.IsDisplayViewPdf

	m.IsDisplayDownloadPdf = model.IsDisplayDownloadPdf

	nn, e = model.LastPaymentDetail.UnmarshalPolymorphicJSON(model.LastPaymentDetail.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.LastPaymentDetail = nn.(PaymentDetail)
	} else {
		m.LastPaymentDetail = nil
	}

	m.PartyName = model.PartyName

	m.SubscriptionIds = make([]string, len(model.SubscriptionIds))
	copy(m.SubscriptionIds, model.SubscriptionIds)
	m.InvoiceId = model.InvoiceId

	return
}

// InvoiceSummaryInvoiceStatusEnum Enum with underlying type: string
type InvoiceSummaryInvoiceStatusEnum string

// Set of constants representing the allowable values for InvoiceSummaryInvoiceStatusEnum
const (
	InvoiceSummaryInvoiceStatusOpen             InvoiceSummaryInvoiceStatusEnum = "OPEN"
	InvoiceSummaryInvoiceStatusPastDue          InvoiceSummaryInvoiceStatusEnum = "PAST_DUE"
	InvoiceSummaryInvoiceStatusPaymentSubmitted InvoiceSummaryInvoiceStatusEnum = "PAYMENT_SUBMITTED"
	InvoiceSummaryInvoiceStatusClosed           InvoiceSummaryInvoiceStatusEnum = "CLOSED"
)

var mappingInvoiceSummaryInvoiceStatusEnum = map[string]InvoiceSummaryInvoiceStatusEnum{
	"OPEN":              InvoiceSummaryInvoiceStatusOpen,
	"PAST_DUE":          InvoiceSummaryInvoiceStatusPastDue,
	"PAYMENT_SUBMITTED": InvoiceSummaryInvoiceStatusPaymentSubmitted,
	"CLOSED":            InvoiceSummaryInvoiceStatusClosed,
}

var mappingInvoiceSummaryInvoiceStatusEnumLowerCase = map[string]InvoiceSummaryInvoiceStatusEnum{
	"open":              InvoiceSummaryInvoiceStatusOpen,
	"past_due":          InvoiceSummaryInvoiceStatusPastDue,
	"payment_submitted": InvoiceSummaryInvoiceStatusPaymentSubmitted,
	"closed":            InvoiceSummaryInvoiceStatusClosed,
}

// GetInvoiceSummaryInvoiceStatusEnumValues Enumerates the set of values for InvoiceSummaryInvoiceStatusEnum
func GetInvoiceSummaryInvoiceStatusEnumValues() []InvoiceSummaryInvoiceStatusEnum {
	values := make([]InvoiceSummaryInvoiceStatusEnum, 0)
	for _, v := range mappingInvoiceSummaryInvoiceStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetInvoiceSummaryInvoiceStatusEnumStringValues Enumerates the set of values in String for InvoiceSummaryInvoiceStatusEnum
func GetInvoiceSummaryInvoiceStatusEnumStringValues() []string {
	return []string{
		"OPEN",
		"PAST_DUE",
		"PAYMENT_SUBMITTED",
		"CLOSED",
	}
}

// GetMappingInvoiceSummaryInvoiceStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInvoiceSummaryInvoiceStatusEnum(val string) (InvoiceSummaryInvoiceStatusEnum, bool) {
	enum, ok := mappingInvoiceSummaryInvoiceStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// InvoiceSummaryInvoiceTypeEnum Enum with underlying type: string
type InvoiceSummaryInvoiceTypeEnum string

// Set of constants representing the allowable values for InvoiceSummaryInvoiceTypeEnum
const (
	InvoiceSummaryInvoiceTypeHardware     InvoiceSummaryInvoiceTypeEnum = "HARDWARE"
	InvoiceSummaryInvoiceTypeSubscription InvoiceSummaryInvoiceTypeEnum = "SUBSCRIPTION"
	InvoiceSummaryInvoiceTypeSupport      InvoiceSummaryInvoiceTypeEnum = "SUPPORT"
	InvoiceSummaryInvoiceTypeLicense      InvoiceSummaryInvoiceTypeEnum = "LICENSE"
	InvoiceSummaryInvoiceTypeEducation    InvoiceSummaryInvoiceTypeEnum = "EDUCATION"
	InvoiceSummaryInvoiceTypeConsulting   InvoiceSummaryInvoiceTypeEnum = "CONSULTING"
	InvoiceSummaryInvoiceTypeService      InvoiceSummaryInvoiceTypeEnum = "SERVICE"
	InvoiceSummaryInvoiceTypeUsage        InvoiceSummaryInvoiceTypeEnum = "USAGE"
)

var mappingInvoiceSummaryInvoiceTypeEnum = map[string]InvoiceSummaryInvoiceTypeEnum{
	"HARDWARE":     InvoiceSummaryInvoiceTypeHardware,
	"SUBSCRIPTION": InvoiceSummaryInvoiceTypeSubscription,
	"SUPPORT":      InvoiceSummaryInvoiceTypeSupport,
	"LICENSE":      InvoiceSummaryInvoiceTypeLicense,
	"EDUCATION":    InvoiceSummaryInvoiceTypeEducation,
	"CONSULTING":   InvoiceSummaryInvoiceTypeConsulting,
	"SERVICE":      InvoiceSummaryInvoiceTypeService,
	"USAGE":        InvoiceSummaryInvoiceTypeUsage,
}

var mappingInvoiceSummaryInvoiceTypeEnumLowerCase = map[string]InvoiceSummaryInvoiceTypeEnum{
	"hardware":     InvoiceSummaryInvoiceTypeHardware,
	"subscription": InvoiceSummaryInvoiceTypeSubscription,
	"support":      InvoiceSummaryInvoiceTypeSupport,
	"license":      InvoiceSummaryInvoiceTypeLicense,
	"education":    InvoiceSummaryInvoiceTypeEducation,
	"consulting":   InvoiceSummaryInvoiceTypeConsulting,
	"service":      InvoiceSummaryInvoiceTypeService,
	"usage":        InvoiceSummaryInvoiceTypeUsage,
}

// GetInvoiceSummaryInvoiceTypeEnumValues Enumerates the set of values for InvoiceSummaryInvoiceTypeEnum
func GetInvoiceSummaryInvoiceTypeEnumValues() []InvoiceSummaryInvoiceTypeEnum {
	values := make([]InvoiceSummaryInvoiceTypeEnum, 0)
	for _, v := range mappingInvoiceSummaryInvoiceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetInvoiceSummaryInvoiceTypeEnumStringValues Enumerates the set of values in String for InvoiceSummaryInvoiceTypeEnum
func GetInvoiceSummaryInvoiceTypeEnumStringValues() []string {
	return []string{
		"HARDWARE",
		"SUBSCRIPTION",
		"SUPPORT",
		"LICENSE",
		"EDUCATION",
		"CONSULTING",
		"SERVICE",
		"USAGE",
	}
}

// GetMappingInvoiceSummaryInvoiceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInvoiceSummaryInvoiceTypeEnum(val string) (InvoiceSummaryInvoiceTypeEnum, bool) {
	enum, ok := mappingInvoiceSummaryInvoiceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
