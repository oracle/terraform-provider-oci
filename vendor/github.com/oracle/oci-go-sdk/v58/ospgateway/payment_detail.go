// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// PaymentDetail Payment related details
type PaymentDetail interface {

	// Paid the invoice on this day
	GetTimePaidOn() *common.SDKTime

	// example
	GetPaidBy() *string

	// Amount that paid
	GetAmountPaid() *float32
}

type paymentdetail struct {
	JsonData      []byte
	TimePaidOn    *common.SDKTime `mandatory:"false" json:"timePaidOn"`
	PaidBy        *string         `mandatory:"false" json:"paidBy"`
	AmountPaid    *float32        `mandatory:"false" json:"amountPaid"`
	PaymentMethod string          `json:"paymentMethod"`
}

// UnmarshalJSON unmarshals json
func (m *paymentdetail) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerpaymentdetail paymentdetail
	s := struct {
		Model Unmarshalerpaymentdetail
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.TimePaidOn = s.Model.TimePaidOn
	m.PaidBy = s.Model.PaidBy
	m.AmountPaid = s.Model.AmountPaid
	m.PaymentMethod = s.Model.PaymentMethod

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *paymentdetail) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.PaymentMethod {
	case "OTHER":
		mm := OtherPaymentDetail{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PAYPAL":
		mm := PaypalPaymentDetail{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CREDIT_CARD":
		mm := CreditCardPaymentDetail{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetTimePaidOn returns TimePaidOn
func (m paymentdetail) GetTimePaidOn() *common.SDKTime {
	return m.TimePaidOn
}

//GetPaidBy returns PaidBy
func (m paymentdetail) GetPaidBy() *string {
	return m.PaidBy
}

//GetAmountPaid returns AmountPaid
func (m paymentdetail) GetAmountPaid() *float32 {
	return m.AmountPaid
}

func (m paymentdetail) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m paymentdetail) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PaymentDetailPaymentMethodEnum Enum with underlying type: string
type PaymentDetailPaymentMethodEnum string

// Set of constants representing the allowable values for PaymentDetailPaymentMethodEnum
const (
	PaymentDetailPaymentMethodCreditCard PaymentDetailPaymentMethodEnum = "CREDIT_CARD"
	PaymentDetailPaymentMethodPaypal     PaymentDetailPaymentMethodEnum = "PAYPAL"
	PaymentDetailPaymentMethodOther      PaymentDetailPaymentMethodEnum = "OTHER"
)

var mappingPaymentDetailPaymentMethodEnum = map[string]PaymentDetailPaymentMethodEnum{
	"CREDIT_CARD": PaymentDetailPaymentMethodCreditCard,
	"PAYPAL":      PaymentDetailPaymentMethodPaypal,
	"OTHER":       PaymentDetailPaymentMethodOther,
}

// GetPaymentDetailPaymentMethodEnumValues Enumerates the set of values for PaymentDetailPaymentMethodEnum
func GetPaymentDetailPaymentMethodEnumValues() []PaymentDetailPaymentMethodEnum {
	values := make([]PaymentDetailPaymentMethodEnum, 0)
	for _, v := range mappingPaymentDetailPaymentMethodEnum {
		values = append(values, v)
	}
	return values
}

// GetPaymentDetailPaymentMethodEnumStringValues Enumerates the set of values in String for PaymentDetailPaymentMethodEnum
func GetPaymentDetailPaymentMethodEnumStringValues() []string {
	return []string{
		"CREDIT_CARD",
		"PAYPAL",
		"OTHER",
	}
}

// GetMappingPaymentDetailPaymentMethodEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPaymentDetailPaymentMethodEnum(val string) (PaymentDetailPaymentMethodEnum, bool) {
	mappingPaymentDetailPaymentMethodEnumIgnoreCase := make(map[string]PaymentDetailPaymentMethodEnum)
	for k, v := range mappingPaymentDetailPaymentMethodEnum {
		mappingPaymentDetailPaymentMethodEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingPaymentDetailPaymentMethodEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
