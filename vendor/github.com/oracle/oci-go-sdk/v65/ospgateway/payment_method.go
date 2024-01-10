// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// PaymentMethodEnum Enum with underlying type: string
type PaymentMethodEnum string

// Set of constants representing the allowable values for PaymentMethodEnum
const (
	PaymentMethodCreditCard PaymentMethodEnum = "CREDIT_CARD"
	PaymentMethodPaypal     PaymentMethodEnum = "PAYPAL"
)

var mappingPaymentMethodEnum = map[string]PaymentMethodEnum{
	"CREDIT_CARD": PaymentMethodCreditCard,
	"PAYPAL":      PaymentMethodPaypal,
}

var mappingPaymentMethodEnumLowerCase = map[string]PaymentMethodEnum{
	"credit_card": PaymentMethodCreditCard,
	"paypal":      PaymentMethodPaypal,
}

// GetPaymentMethodEnumValues Enumerates the set of values for PaymentMethodEnum
func GetPaymentMethodEnumValues() []PaymentMethodEnum {
	values := make([]PaymentMethodEnum, 0)
	for _, v := range mappingPaymentMethodEnum {
		values = append(values, v)
	}
	return values
}

// GetPaymentMethodEnumStringValues Enumerates the set of values in String for PaymentMethodEnum
func GetPaymentMethodEnumStringValues() []string {
	return []string{
		"CREDIT_CARD",
		"PAYPAL",
	}
}

// GetMappingPaymentMethodEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPaymentMethodEnum(val string) (PaymentMethodEnum, bool) {
	enum, ok := mappingPaymentMethodEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
