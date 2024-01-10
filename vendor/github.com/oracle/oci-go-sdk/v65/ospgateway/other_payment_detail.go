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

// OtherPaymentDetail Other Payment related details
type OtherPaymentDetail struct {

	// Paid the invoice on this day
	TimePaidOn *common.SDKTime `mandatory:"false" json:"timePaidOn"`

	// example
	PaidBy *string `mandatory:"false" json:"paidBy"`

	// Amount that paid
	AmountPaid *float32 `mandatory:"false" json:"amountPaid"`

	// Last four routing digits of the card
	EcheckRouting *string `mandatory:"false" json:"echeckRouting"`

	// Name on the echeck card
	NameOnCard *string `mandatory:"false" json:"nameOnCard"`

	// Last four digits of the card
	LastDigits *string `mandatory:"false" json:"lastDigits"`

	// Expired date of the echeck card
	TimeExpiration *common.SDKTime `mandatory:"false" json:"timeExpiration"`

	// Echeck card type
	CreditCardType OtherPaymentDetailCreditCardTypeEnum `mandatory:"false" json:"creditCardType,omitempty"`
}

// GetTimePaidOn returns TimePaidOn
func (m OtherPaymentDetail) GetTimePaidOn() *common.SDKTime {
	return m.TimePaidOn
}

// GetPaidBy returns PaidBy
func (m OtherPaymentDetail) GetPaidBy() *string {
	return m.PaidBy
}

// GetAmountPaid returns AmountPaid
func (m OtherPaymentDetail) GetAmountPaid() *float32 {
	return m.AmountPaid
}

func (m OtherPaymentDetail) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OtherPaymentDetail) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOtherPaymentDetailCreditCardTypeEnum(string(m.CreditCardType)); !ok && m.CreditCardType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CreditCardType: %s. Supported values are: %s.", m.CreditCardType, strings.Join(GetOtherPaymentDetailCreditCardTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m OtherPaymentDetail) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOtherPaymentDetail OtherPaymentDetail
	s := struct {
		DiscriminatorParam string `json:"paymentMethod"`
		MarshalTypeOtherPaymentDetail
	}{
		"OTHER",
		(MarshalTypeOtherPaymentDetail)(m),
	}

	return json.Marshal(&s)
}

// OtherPaymentDetailCreditCardTypeEnum Enum with underlying type: string
type OtherPaymentDetailCreditCardTypeEnum string

// Set of constants representing the allowable values for OtherPaymentDetailCreditCardTypeEnum
const (
	OtherPaymentDetailCreditCardTypeVisa              OtherPaymentDetailCreditCardTypeEnum = "VISA"
	OtherPaymentDetailCreditCardTypeAmex              OtherPaymentDetailCreditCardTypeEnum = "AMEX"
	OtherPaymentDetailCreditCardTypeMastercard        OtherPaymentDetailCreditCardTypeEnum = "MASTERCARD"
	OtherPaymentDetailCreditCardTypeDiscover          OtherPaymentDetailCreditCardTypeEnum = "DISCOVER"
	OtherPaymentDetailCreditCardTypeJcb               OtherPaymentDetailCreditCardTypeEnum = "JCB"
	OtherPaymentDetailCreditCardTypeDiner             OtherPaymentDetailCreditCardTypeEnum = "DINER"
	OtherPaymentDetailCreditCardTypeElo               OtherPaymentDetailCreditCardTypeEnum = "ELO"
	OtherPaymentDetailCreditCardTypeSaving            OtherPaymentDetailCreditCardTypeEnum = "SAVING"
	OtherPaymentDetailCreditCardTypeChecking          OtherPaymentDetailCreditCardTypeEnum = "CHECKING"
	OtherPaymentDetailCreditCardTypeCorporateChecking OtherPaymentDetailCreditCardTypeEnum = "CORPORATE_CHECKING"
)

var mappingOtherPaymentDetailCreditCardTypeEnum = map[string]OtherPaymentDetailCreditCardTypeEnum{
	"VISA":               OtherPaymentDetailCreditCardTypeVisa,
	"AMEX":               OtherPaymentDetailCreditCardTypeAmex,
	"MASTERCARD":         OtherPaymentDetailCreditCardTypeMastercard,
	"DISCOVER":           OtherPaymentDetailCreditCardTypeDiscover,
	"JCB":                OtherPaymentDetailCreditCardTypeJcb,
	"DINER":              OtherPaymentDetailCreditCardTypeDiner,
	"ELO":                OtherPaymentDetailCreditCardTypeElo,
	"SAVING":             OtherPaymentDetailCreditCardTypeSaving,
	"CHECKING":           OtherPaymentDetailCreditCardTypeChecking,
	"CORPORATE_CHECKING": OtherPaymentDetailCreditCardTypeCorporateChecking,
}

var mappingOtherPaymentDetailCreditCardTypeEnumLowerCase = map[string]OtherPaymentDetailCreditCardTypeEnum{
	"visa":               OtherPaymentDetailCreditCardTypeVisa,
	"amex":               OtherPaymentDetailCreditCardTypeAmex,
	"mastercard":         OtherPaymentDetailCreditCardTypeMastercard,
	"discover":           OtherPaymentDetailCreditCardTypeDiscover,
	"jcb":                OtherPaymentDetailCreditCardTypeJcb,
	"diner":              OtherPaymentDetailCreditCardTypeDiner,
	"elo":                OtherPaymentDetailCreditCardTypeElo,
	"saving":             OtherPaymentDetailCreditCardTypeSaving,
	"checking":           OtherPaymentDetailCreditCardTypeChecking,
	"corporate_checking": OtherPaymentDetailCreditCardTypeCorporateChecking,
}

// GetOtherPaymentDetailCreditCardTypeEnumValues Enumerates the set of values for OtherPaymentDetailCreditCardTypeEnum
func GetOtherPaymentDetailCreditCardTypeEnumValues() []OtherPaymentDetailCreditCardTypeEnum {
	values := make([]OtherPaymentDetailCreditCardTypeEnum, 0)
	for _, v := range mappingOtherPaymentDetailCreditCardTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOtherPaymentDetailCreditCardTypeEnumStringValues Enumerates the set of values in String for OtherPaymentDetailCreditCardTypeEnum
func GetOtherPaymentDetailCreditCardTypeEnumStringValues() []string {
	return []string{
		"VISA",
		"AMEX",
		"MASTERCARD",
		"DISCOVER",
		"JCB",
		"DINER",
		"ELO",
		"SAVING",
		"CHECKING",
		"CORPORATE_CHECKING",
	}
}

// GetMappingOtherPaymentDetailCreditCardTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOtherPaymentDetailCreditCardTypeEnum(val string) (OtherPaymentDetailCreditCardTypeEnum, bool) {
	enum, ok := mappingOtherPaymentDetailCreditCardTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
