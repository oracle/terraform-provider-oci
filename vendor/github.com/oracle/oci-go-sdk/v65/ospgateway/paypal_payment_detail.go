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

// PaypalPaymentDetail PayPal Payment related details
type PaypalPaymentDetail struct {

	// Paid the invoice on this day
	TimePaidOn *common.SDKTime `mandatory:"false" json:"timePaidOn"`

	// example
	PaidBy *string `mandatory:"false" json:"paidBy"`

	// Amount that paid
	AmountPaid *float32 `mandatory:"false" json:"amountPaid"`

	// The id (email address) of the paypal payment
	PaypalId *string `mandatory:"false" json:"paypalId"`

	// paypal payment reference
	PaypalReference *string `mandatory:"false" json:"paypalReference"`
}

// GetTimePaidOn returns TimePaidOn
func (m PaypalPaymentDetail) GetTimePaidOn() *common.SDKTime {
	return m.TimePaidOn
}

// GetPaidBy returns PaidBy
func (m PaypalPaymentDetail) GetPaidBy() *string {
	return m.PaidBy
}

// GetAmountPaid returns AmountPaid
func (m PaypalPaymentDetail) GetAmountPaid() *float32 {
	return m.AmountPaid
}

func (m PaypalPaymentDetail) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PaypalPaymentDetail) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PaypalPaymentDetail) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePaypalPaymentDetail PaypalPaymentDetail
	s := struct {
		DiscriminatorParam string `json:"paymentMethod"`
		MarshalTypePaypalPaymentDetail
	}{
		"PAYPAL",
		(MarshalTypePaypalPaymentDetail)(m),
	}

	return json.Marshal(&s)
}
