// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Digital Assistant Service Instance API
//
// API to create and maintain Oracle Digital Assistant service instances.
//

package oda

import (
	"strings"
)

// TranslationServiceEnum Enum with underlying type: string
type TranslationServiceEnum string

// Set of constants representing the allowable values for TranslationServiceEnum
const (
	TranslationServiceGoogle    TranslationServiceEnum = "GOOGLE"
	TranslationServiceMicrosoft TranslationServiceEnum = "MICROSOFT"
)

var mappingTranslationServiceEnum = map[string]TranslationServiceEnum{
	"GOOGLE":    TranslationServiceGoogle,
	"MICROSOFT": TranslationServiceMicrosoft,
}

var mappingTranslationServiceEnumLowerCase = map[string]TranslationServiceEnum{
	"google":    TranslationServiceGoogle,
	"microsoft": TranslationServiceMicrosoft,
}

// GetTranslationServiceEnumValues Enumerates the set of values for TranslationServiceEnum
func GetTranslationServiceEnumValues() []TranslationServiceEnum {
	values := make([]TranslationServiceEnum, 0)
	for _, v := range mappingTranslationServiceEnum {
		values = append(values, v)
	}
	return values
}

// GetTranslationServiceEnumStringValues Enumerates the set of values in String for TranslationServiceEnum
func GetTranslationServiceEnumStringValues() []string {
	return []string{
		"GOOGLE",
		"MICROSOFT",
	}
}

// GetMappingTranslationServiceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTranslationServiceEnum(val string) (TranslationServiceEnum, bool) {
	enum, ok := mappingTranslationServiceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
