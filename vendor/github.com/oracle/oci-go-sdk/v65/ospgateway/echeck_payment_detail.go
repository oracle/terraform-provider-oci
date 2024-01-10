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

// EcheckPaymentDetail Echeck Payment related details
type EcheckPaymentDetail struct {

	// Paid the invoice on this day
	TimePaidOn *common.SDKTime `mandatory:"false" json:"timePaidOn"`

	// example
	PaidBy *string `mandatory:"false" json:"paidBy"`

	// Amount that paid
	AmountPaid *float32 `mandatory:"false" json:"amountPaid"`

	// Name on the echeck card
	NameOnCard *string `mandatory:"false" json:"nameOnCard"`

	// Account number of the card owner
	AccountNumber *string `mandatory:"false" json:"accountNumber"`

	// Routing number of the echeck card
	RoutingNumber *string `mandatory:"false" json:"routingNumber"`

	// Echeck card type
	CardType EcheckPaymentDetailCardTypeEnum `mandatory:"false" json:"cardType,omitempty"`
}

// GetTimePaidOn returns TimePaidOn
func (m EcheckPaymentDetail) GetTimePaidOn() *common.SDKTime {
	return m.TimePaidOn
}

// GetPaidBy returns PaidBy
func (m EcheckPaymentDetail) GetPaidBy() *string {
	return m.PaidBy
}

// GetAmountPaid returns AmountPaid
func (m EcheckPaymentDetail) GetAmountPaid() *float32 {
	return m.AmountPaid
}

func (m EcheckPaymentDetail) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EcheckPaymentDetail) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingEcheckPaymentDetailCardTypeEnum(string(m.CardType)); !ok && m.CardType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CardType: %s. Supported values are: %s.", m.CardType, strings.Join(GetEcheckPaymentDetailCardTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m EcheckPaymentDetail) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeEcheckPaymentDetail EcheckPaymentDetail
	s := struct {
		DiscriminatorParam string `json:"paymentMethod"`
		MarshalTypeEcheckPaymentDetail
	}{
		"ECHECK",
		(MarshalTypeEcheckPaymentDetail)(m),
	}

	return json.Marshal(&s)
}

// EcheckPaymentDetailCardTypeEnum Enum with underlying type: string
type EcheckPaymentDetailCardTypeEnum string

// Set of constants representing the allowable values for EcheckPaymentDetailCardTypeEnum
const (
	EcheckPaymentDetailCardTypeSaving            EcheckPaymentDetailCardTypeEnum = "SAVING"
	EcheckPaymentDetailCardTypeChecking          EcheckPaymentDetailCardTypeEnum = "CHECKING"
	EcheckPaymentDetailCardTypeCorporateChecking EcheckPaymentDetailCardTypeEnum = "CORPORATE_CHECKING"
)

var mappingEcheckPaymentDetailCardTypeEnum = map[string]EcheckPaymentDetailCardTypeEnum{
	"SAVING":             EcheckPaymentDetailCardTypeSaving,
	"CHECKING":           EcheckPaymentDetailCardTypeChecking,
	"CORPORATE_CHECKING": EcheckPaymentDetailCardTypeCorporateChecking,
}

var mappingEcheckPaymentDetailCardTypeEnumLowerCase = map[string]EcheckPaymentDetailCardTypeEnum{
	"saving":             EcheckPaymentDetailCardTypeSaving,
	"checking":           EcheckPaymentDetailCardTypeChecking,
	"corporate_checking": EcheckPaymentDetailCardTypeCorporateChecking,
}

// GetEcheckPaymentDetailCardTypeEnumValues Enumerates the set of values for EcheckPaymentDetailCardTypeEnum
func GetEcheckPaymentDetailCardTypeEnumValues() []EcheckPaymentDetailCardTypeEnum {
	values := make([]EcheckPaymentDetailCardTypeEnum, 0)
	for _, v := range mappingEcheckPaymentDetailCardTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetEcheckPaymentDetailCardTypeEnumStringValues Enumerates the set of values in String for EcheckPaymentDetailCardTypeEnum
func GetEcheckPaymentDetailCardTypeEnumStringValues() []string {
	return []string{
		"SAVING",
		"CHECKING",
		"CORPORATE_CHECKING",
	}
}

// GetMappingEcheckPaymentDetailCardTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEcheckPaymentDetailCardTypeEnum(val string) (EcheckPaymentDetailCardTypeEnum, bool) {
	enum, ok := mappingEcheckPaymentDetailCardTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
