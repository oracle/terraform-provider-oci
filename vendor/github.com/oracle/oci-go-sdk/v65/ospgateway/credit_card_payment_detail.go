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

// CreditCardPaymentDetail Credit card Payment related details
type CreditCardPaymentDetail struct {

	// Paid the invoice on this day
	TimePaidOn *common.SDKTime `mandatory:"false" json:"timePaidOn"`

	// example
	PaidBy *string `mandatory:"false" json:"paidBy"`

	// Amount that paid
	AmountPaid *float32 `mandatory:"false" json:"amountPaid"`

	// Name on the credit card
	NameOnCard *string `mandatory:"false" json:"nameOnCard"`

	// Last four digits of the card
	LastDigits *string `mandatory:"false" json:"lastDigits"`

	// Expired date of the credit card
	TimeExpiration *common.SDKTime `mandatory:"false" json:"timeExpiration"`

	// Credit card type
	CreditCardType CreditCardPaymentDetailCreditCardTypeEnum `mandatory:"false" json:"creditCardType,omitempty"`
}

// GetTimePaidOn returns TimePaidOn
func (m CreditCardPaymentDetail) GetTimePaidOn() *common.SDKTime {
	return m.TimePaidOn
}

// GetPaidBy returns PaidBy
func (m CreditCardPaymentDetail) GetPaidBy() *string {
	return m.PaidBy
}

// GetAmountPaid returns AmountPaid
func (m CreditCardPaymentDetail) GetAmountPaid() *float32 {
	return m.AmountPaid
}

func (m CreditCardPaymentDetail) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreditCardPaymentDetail) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreditCardPaymentDetailCreditCardTypeEnum(string(m.CreditCardType)); !ok && m.CreditCardType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CreditCardType: %s. Supported values are: %s.", m.CreditCardType, strings.Join(GetCreditCardPaymentDetailCreditCardTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreditCardPaymentDetail) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreditCardPaymentDetail CreditCardPaymentDetail
	s := struct {
		DiscriminatorParam string `json:"paymentMethod"`
		MarshalTypeCreditCardPaymentDetail
	}{
		"CREDIT_CARD",
		(MarshalTypeCreditCardPaymentDetail)(m),
	}

	return json.Marshal(&s)
}

// CreditCardPaymentDetailCreditCardTypeEnum Enum with underlying type: string
type CreditCardPaymentDetailCreditCardTypeEnum string

// Set of constants representing the allowable values for CreditCardPaymentDetailCreditCardTypeEnum
const (
	CreditCardPaymentDetailCreditCardTypeVisa       CreditCardPaymentDetailCreditCardTypeEnum = "VISA"
	CreditCardPaymentDetailCreditCardTypeAmex       CreditCardPaymentDetailCreditCardTypeEnum = "AMEX"
	CreditCardPaymentDetailCreditCardTypeMastercard CreditCardPaymentDetailCreditCardTypeEnum = "MASTERCARD"
	CreditCardPaymentDetailCreditCardTypeDiscover   CreditCardPaymentDetailCreditCardTypeEnum = "DISCOVER"
	CreditCardPaymentDetailCreditCardTypeJcb        CreditCardPaymentDetailCreditCardTypeEnum = "JCB"
	CreditCardPaymentDetailCreditCardTypeDiner      CreditCardPaymentDetailCreditCardTypeEnum = "DINER"
	CreditCardPaymentDetailCreditCardTypeElo        CreditCardPaymentDetailCreditCardTypeEnum = "ELO"
)

var mappingCreditCardPaymentDetailCreditCardTypeEnum = map[string]CreditCardPaymentDetailCreditCardTypeEnum{
	"VISA":       CreditCardPaymentDetailCreditCardTypeVisa,
	"AMEX":       CreditCardPaymentDetailCreditCardTypeAmex,
	"MASTERCARD": CreditCardPaymentDetailCreditCardTypeMastercard,
	"DISCOVER":   CreditCardPaymentDetailCreditCardTypeDiscover,
	"JCB":        CreditCardPaymentDetailCreditCardTypeJcb,
	"DINER":      CreditCardPaymentDetailCreditCardTypeDiner,
	"ELO":        CreditCardPaymentDetailCreditCardTypeElo,
}

var mappingCreditCardPaymentDetailCreditCardTypeEnumLowerCase = map[string]CreditCardPaymentDetailCreditCardTypeEnum{
	"visa":       CreditCardPaymentDetailCreditCardTypeVisa,
	"amex":       CreditCardPaymentDetailCreditCardTypeAmex,
	"mastercard": CreditCardPaymentDetailCreditCardTypeMastercard,
	"discover":   CreditCardPaymentDetailCreditCardTypeDiscover,
	"jcb":        CreditCardPaymentDetailCreditCardTypeJcb,
	"diner":      CreditCardPaymentDetailCreditCardTypeDiner,
	"elo":        CreditCardPaymentDetailCreditCardTypeElo,
}

// GetCreditCardPaymentDetailCreditCardTypeEnumValues Enumerates the set of values for CreditCardPaymentDetailCreditCardTypeEnum
func GetCreditCardPaymentDetailCreditCardTypeEnumValues() []CreditCardPaymentDetailCreditCardTypeEnum {
	values := make([]CreditCardPaymentDetailCreditCardTypeEnum, 0)
	for _, v := range mappingCreditCardPaymentDetailCreditCardTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreditCardPaymentDetailCreditCardTypeEnumStringValues Enumerates the set of values in String for CreditCardPaymentDetailCreditCardTypeEnum
func GetCreditCardPaymentDetailCreditCardTypeEnumStringValues() []string {
	return []string{
		"VISA",
		"AMEX",
		"MASTERCARD",
		"DISCOVER",
		"JCB",
		"DINER",
		"ELO",
	}
}

// GetMappingCreditCardPaymentDetailCreditCardTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreditCardPaymentDetailCreditCardTypeEnum(val string) (CreditCardPaymentDetailCreditCardTypeEnum, bool) {
	enum, ok := mappingCreditCardPaymentDetailCreditCardTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
