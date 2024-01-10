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

// AddressVerificationCodeEnum Enum with underlying type: string
type AddressVerificationCodeEnum string

// Set of constants representing the allowable values for AddressVerificationCodeEnum
const (
	AddressVerificationCodeVerified          AddressVerificationCodeEnum = "VERIFIED"
	AddressVerificationCodePartiallyVerified AddressVerificationCodeEnum = "PARTIALLY_VERIFIED"
	AddressVerificationCodeAmbiguous         AddressVerificationCodeEnum = "AMBIGUOUS"
	AddressVerificationCodeReverted          AddressVerificationCodeEnum = "REVERTED"
	AddressVerificationCodeUnverified        AddressVerificationCodeEnum = "UNVERIFIED"
)

var mappingAddressVerificationCodeEnum = map[string]AddressVerificationCodeEnum{
	"VERIFIED":           AddressVerificationCodeVerified,
	"PARTIALLY_VERIFIED": AddressVerificationCodePartiallyVerified,
	"AMBIGUOUS":          AddressVerificationCodeAmbiguous,
	"REVERTED":           AddressVerificationCodeReverted,
	"UNVERIFIED":         AddressVerificationCodeUnverified,
}

var mappingAddressVerificationCodeEnumLowerCase = map[string]AddressVerificationCodeEnum{
	"verified":           AddressVerificationCodeVerified,
	"partially_verified": AddressVerificationCodePartiallyVerified,
	"ambiguous":          AddressVerificationCodeAmbiguous,
	"reverted":           AddressVerificationCodeReverted,
	"unverified":         AddressVerificationCodeUnverified,
}

// GetAddressVerificationCodeEnumValues Enumerates the set of values for AddressVerificationCodeEnum
func GetAddressVerificationCodeEnumValues() []AddressVerificationCodeEnum {
	values := make([]AddressVerificationCodeEnum, 0)
	for _, v := range mappingAddressVerificationCodeEnum {
		values = append(values, v)
	}
	return values
}

// GetAddressVerificationCodeEnumStringValues Enumerates the set of values in String for AddressVerificationCodeEnum
func GetAddressVerificationCodeEnumStringValues() []string {
	return []string{
		"VERIFIED",
		"PARTIALLY_VERIFIED",
		"AMBIGUOUS",
		"REVERTED",
		"UNVERIFIED",
	}
}

// GetMappingAddressVerificationCodeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAddressVerificationCodeEnum(val string) (AddressVerificationCodeEnum, bool) {
	enum, ok := mappingAddressVerificationCodeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
