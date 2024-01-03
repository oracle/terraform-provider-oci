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

// ThirdPartyValidationTypeEnum Enum with underlying type: string
type ThirdPartyValidationTypeEnum string

// Set of constants representing the allowable values for ThirdPartyValidationTypeEnum
const (
	ThirdPartyValidationTypeOptional ThirdPartyValidationTypeEnum = "OPTIONAL"
	ThirdPartyValidationTypeRequired ThirdPartyValidationTypeEnum = "REQUIRED"
	ThirdPartyValidationTypeNever    ThirdPartyValidationTypeEnum = "NEVER"
)

var mappingThirdPartyValidationTypeEnum = map[string]ThirdPartyValidationTypeEnum{
	"OPTIONAL": ThirdPartyValidationTypeOptional,
	"REQUIRED": ThirdPartyValidationTypeRequired,
	"NEVER":    ThirdPartyValidationTypeNever,
}

var mappingThirdPartyValidationTypeEnumLowerCase = map[string]ThirdPartyValidationTypeEnum{
	"optional": ThirdPartyValidationTypeOptional,
	"required": ThirdPartyValidationTypeRequired,
	"never":    ThirdPartyValidationTypeNever,
}

// GetThirdPartyValidationTypeEnumValues Enumerates the set of values for ThirdPartyValidationTypeEnum
func GetThirdPartyValidationTypeEnumValues() []ThirdPartyValidationTypeEnum {
	values := make([]ThirdPartyValidationTypeEnum, 0)
	for _, v := range mappingThirdPartyValidationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetThirdPartyValidationTypeEnumStringValues Enumerates the set of values in String for ThirdPartyValidationTypeEnum
func GetThirdPartyValidationTypeEnumStringValues() []string {
	return []string{
		"OPTIONAL",
		"REQUIRED",
		"NEVER",
	}
}

// GetMappingThirdPartyValidationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingThirdPartyValidationTypeEnum(val string) (ThirdPartyValidationTypeEnum, bool) {
	enum, ok := mappingThirdPartyValidationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
