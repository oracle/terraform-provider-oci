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

// PaymentOption Payment option of a subscription.
type PaymentOption interface {

	// Wallet instrument internal id.
	GetWalletInstrumentId() *string

	// Wallet transaction id.
	GetWalletTransactionId() *string
}

type paymentoption struct {
	JsonData            []byte
	WalletInstrumentId  *string `mandatory:"false" json:"walletInstrumentId"`
	WalletTransactionId *string `mandatory:"false" json:"walletTransactionId"`
	PaymentMethod       string  `json:"paymentMethod"`
}

// UnmarshalJSON unmarshals json
func (m *paymentoption) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerpaymentoption paymentoption
	s := struct {
		Model Unmarshalerpaymentoption
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.WalletInstrumentId = s.Model.WalletInstrumentId
	m.WalletTransactionId = s.Model.WalletTransactionId
	m.PaymentMethod = s.Model.PaymentMethod

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *paymentoption) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.PaymentMethod {
	case "CREDIT_CARD":
		mm := CreditCardPaymentOption{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PAYPAL":
		mm := PaypalPaymentOption{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for PaymentOption: %s.", m.PaymentMethod)
		return *m, nil
	}
}

// GetWalletInstrumentId returns WalletInstrumentId
func (m paymentoption) GetWalletInstrumentId() *string {
	return m.WalletInstrumentId
}

// GetWalletTransactionId returns WalletTransactionId
func (m paymentoption) GetWalletTransactionId() *string {
	return m.WalletTransactionId
}

func (m paymentoption) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m paymentoption) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
