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

// PaypalPaymentOption PayPal Payment related details
type PaypalPaymentOption struct {

	// Wallet instrument internal id.
	WalletInstrumentId *string `mandatory:"false" json:"walletInstrumentId"`

	// Wallet transaction id.
	WalletTransactionId *string `mandatory:"false" json:"walletTransactionId"`

	// The email address of the paypal user.
	EmailAddress *string `mandatory:"false" json:"emailAddress"`

	// First name of the paypal user.
	FirstName *string `mandatory:"false" json:"firstName"`

	// Last name of the paypal user.
	LastName *string `mandatory:"false" json:"lastName"`

	// Agreement id for the paypal account.
	ExtBillingAgreementId *string `mandatory:"false" json:"extBillingAgreementId"`
}

// GetWalletInstrumentId returns WalletInstrumentId
func (m PaypalPaymentOption) GetWalletInstrumentId() *string {
	return m.WalletInstrumentId
}

// GetWalletTransactionId returns WalletTransactionId
func (m PaypalPaymentOption) GetWalletTransactionId() *string {
	return m.WalletTransactionId
}

func (m PaypalPaymentOption) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PaypalPaymentOption) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PaypalPaymentOption) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePaypalPaymentOption PaypalPaymentOption
	s := struct {
		DiscriminatorParam string `json:"paymentMethod"`
		MarshalTypePaypalPaymentOption
	}{
		"PAYPAL",
		(MarshalTypePaypalPaymentOption)(m),
	}

	return json.Marshal(&s)
}
