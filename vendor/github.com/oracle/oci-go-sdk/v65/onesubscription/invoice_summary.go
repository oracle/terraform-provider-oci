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

// InvoiceSummary Invoice details
type InvoiceSummary struct {

	// SPM Document Number is an functional identifier for invoice in SPM
	SpmInvoiceNumber *string `mandatory:"true" json:"spmInvoiceNumber"`

	BillToCustomer *InvoicingBusinessPartner `mandatory:"true" json:"billToCustomer"`

	BillToContact *InvoicingUser `mandatory:"true" json:"billToContact"`

	BillToAddress *InvoicingAddress `mandatory:"true" json:"billToAddress"`

	// Payment Method
	PaymentMethod *string `mandatory:"true" json:"paymentMethod"`

	PaymentTerm *InvoicingPaymentTerm `mandatory:"true" json:"paymentTerm"`

	Currency *InvoicingCurrency `mandatory:"true" json:"currency"`

	Organization *InvoicingOrganization `mandatory:"true" json:"organization"`

	// Document Type in SPM like SPM Invoice,SPM Credit Memo etc.,
	Type *string `mandatory:"true" json:"type"`

	// Document Status in SPM which depicts current state of invoice
	Status *string `mandatory:"true" json:"status"`

	// Invoice associated subscription plan number.
	SubscriptionNumber *string `mandatory:"true" json:"subscriptionNumber"`

	// Invoice Date
	TimeInvoiceDate *common.SDKTime `mandatory:"true" json:"timeInvoiceDate"`

	// AR Invoice Numbers comma separated under one invoice
	ArInvoices *string `mandatory:"false" json:"arInvoices"`

	// Receipt Method of Payment Mode
	ReceiptMethod *string `mandatory:"false" json:"receiptMethod"`

	// SPM Invocie creation date
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// User that executed SPM Invoice process
	CreatedBy *string `mandatory:"false" json:"createdBy"`

	// SPM Invoice updated date
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// User that updated SPM Invoice
	UpdatedBy *string `mandatory:"false" json:"updatedBy"`

	// Invoice Lines under particular invoice.
	InvoiceLines []InvoiceLineSummary `mandatory:"false" json:"invoiceLines"`
}

func (m InvoiceSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InvoiceSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
