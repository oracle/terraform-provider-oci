// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OSP Gateway API
//
// This site describes all the Rest endpoints of OSP Gateway.
//

package ospgateway

import (
	"strings"
)

// CreditCardTypeEnum Enum with underlying type: string
type CreditCardTypeEnum string

// Set of constants representing the allowable values for CreditCardTypeEnum
const (
	CreditCardTypeVisa       CreditCardTypeEnum = "VISA"
	CreditCardTypeAmex       CreditCardTypeEnum = "AMEX"
	CreditCardTypeMastercard CreditCardTypeEnum = "MASTERCARD"
	CreditCardTypeDiscover   CreditCardTypeEnum = "DISCOVER"
	CreditCardTypeJcb        CreditCardTypeEnum = "JCB"
	CreditCardTypeDiner      CreditCardTypeEnum = "DINER"
	CreditCardTypeElo        CreditCardTypeEnum = "ELO"
)

var mappingCreditCardTypeEnum = map[string]CreditCardTypeEnum{
	"VISA":       CreditCardTypeVisa,
	"AMEX":       CreditCardTypeAmex,
	"MASTERCARD": CreditCardTypeMastercard,
	"DISCOVER":   CreditCardTypeDiscover,
	"JCB":        CreditCardTypeJcb,
	"DINER":      CreditCardTypeDiner,
	"ELO":        CreditCardTypeElo,
}

// GetCreditCardTypeEnumValues Enumerates the set of values for CreditCardTypeEnum
func GetCreditCardTypeEnumValues() []CreditCardTypeEnum {
	values := make([]CreditCardTypeEnum, 0)
	for _, v := range mappingCreditCardTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreditCardTypeEnumStringValues Enumerates the set of values in String for CreditCardTypeEnum
func GetCreditCardTypeEnumStringValues() []string {
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

// GetMappingCreditCardTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreditCardTypeEnum(val string) (CreditCardTypeEnum, bool) {
	mappingCreditCardTypeEnumIgnoreCase := make(map[string]CreditCardTypeEnum)
	for k, v := range mappingCreditCardTypeEnum {
		mappingCreditCardTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingCreditCardTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
