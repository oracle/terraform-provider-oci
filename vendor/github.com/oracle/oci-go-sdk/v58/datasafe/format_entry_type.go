// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"strings"
)

// FormatEntryTypeEnum Enum with underlying type: string
type FormatEntryTypeEnum string

// Set of constants representing the allowable values for FormatEntryTypeEnum
const (
	FormatEntryTypeDeleteRows                  FormatEntryTypeEnum = "DELETE_ROWS"
	FormatEntryTypeDeterministicSubstitution   FormatEntryTypeEnum = "DETERMINISTIC_SUBSTITUTION"
	FormatEntryTypeDeterministicEncryption     FormatEntryTypeEnum = "DETERMINISTIC_ENCRYPTION"
	FormatEntryTypeDeterministicEncryptionDate FormatEntryTypeEnum = "DETERMINISTIC_ENCRYPTION_DATE"
	FormatEntryTypeFixedNumber                 FormatEntryTypeEnum = "FIXED_NUMBER"
	FormatEntryTypeFixedString                 FormatEntryTypeEnum = "FIXED_STRING"
	FormatEntryTypeLibraryMaskingFormat        FormatEntryTypeEnum = "LIBRARY_MASKING_FORMAT"
	FormatEntryTypeNullValue                   FormatEntryTypeEnum = "NULL_VALUE"
	FormatEntryTypePostProcessingFunction      FormatEntryTypeEnum = "POST_PROCESSING_FUNCTION"
	FormatEntryTypePreserveOriginalData        FormatEntryTypeEnum = "PRESERVE_ORIGINAL_DATA"
	FormatEntryTypeRandomDate                  FormatEntryTypeEnum = "RANDOM_DATE"
	FormatEntryTypeRandomDecimalNumber         FormatEntryTypeEnum = "RANDOM_DECIMAL_NUMBER"
	FormatEntryTypeRandomDigits                FormatEntryTypeEnum = "RANDOM_DIGITS"
	FormatEntryTypeRandomList                  FormatEntryTypeEnum = "RANDOM_LIST"
	FormatEntryTypeRandomNumber                FormatEntryTypeEnum = "RANDOM_NUMBER"
	FormatEntryTypeRandomString                FormatEntryTypeEnum = "RANDOM_STRING"
	FormatEntryTypeRandomSubstitution          FormatEntryTypeEnum = "RANDOM_SUBSTITUTION"
	FormatEntryTypeRegularExpression           FormatEntryTypeEnum = "REGULAR_EXPRESSION"
	FormatEntryTypeShuffle                     FormatEntryTypeEnum = "SHUFFLE"
	FormatEntryTypeSqlExpression               FormatEntryTypeEnum = "SQL_EXPRESSION"
	FormatEntryTypeSubstring                   FormatEntryTypeEnum = "SUBSTRING"
	FormatEntryTypeTruncateTable               FormatEntryTypeEnum = "TRUNCATE_TABLE"
	FormatEntryTypeUserDefinedFunction         FormatEntryTypeEnum = "USER_DEFINED_FUNCTION"
)

var mappingFormatEntryTypeEnum = map[string]FormatEntryTypeEnum{
	"DELETE_ROWS":                   FormatEntryTypeDeleteRows,
	"DETERMINISTIC_SUBSTITUTION":    FormatEntryTypeDeterministicSubstitution,
	"DETERMINISTIC_ENCRYPTION":      FormatEntryTypeDeterministicEncryption,
	"DETERMINISTIC_ENCRYPTION_DATE": FormatEntryTypeDeterministicEncryptionDate,
	"FIXED_NUMBER":                  FormatEntryTypeFixedNumber,
	"FIXED_STRING":                  FormatEntryTypeFixedString,
	"LIBRARY_MASKING_FORMAT":        FormatEntryTypeLibraryMaskingFormat,
	"NULL_VALUE":                    FormatEntryTypeNullValue,
	"POST_PROCESSING_FUNCTION":      FormatEntryTypePostProcessingFunction,
	"PRESERVE_ORIGINAL_DATA":        FormatEntryTypePreserveOriginalData,
	"RANDOM_DATE":                   FormatEntryTypeRandomDate,
	"RANDOM_DECIMAL_NUMBER":         FormatEntryTypeRandomDecimalNumber,
	"RANDOM_DIGITS":                 FormatEntryTypeRandomDigits,
	"RANDOM_LIST":                   FormatEntryTypeRandomList,
	"RANDOM_NUMBER":                 FormatEntryTypeRandomNumber,
	"RANDOM_STRING":                 FormatEntryTypeRandomString,
	"RANDOM_SUBSTITUTION":           FormatEntryTypeRandomSubstitution,
	"REGULAR_EXPRESSION":            FormatEntryTypeRegularExpression,
	"SHUFFLE":                       FormatEntryTypeShuffle,
	"SQL_EXPRESSION":                FormatEntryTypeSqlExpression,
	"SUBSTRING":                     FormatEntryTypeSubstring,
	"TRUNCATE_TABLE":                FormatEntryTypeTruncateTable,
	"USER_DEFINED_FUNCTION":         FormatEntryTypeUserDefinedFunction,
}

// GetFormatEntryTypeEnumValues Enumerates the set of values for FormatEntryTypeEnum
func GetFormatEntryTypeEnumValues() []FormatEntryTypeEnum {
	values := make([]FormatEntryTypeEnum, 0)
	for _, v := range mappingFormatEntryTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetFormatEntryTypeEnumStringValues Enumerates the set of values in String for FormatEntryTypeEnum
func GetFormatEntryTypeEnumStringValues() []string {
	return []string{
		"DELETE_ROWS",
		"DETERMINISTIC_SUBSTITUTION",
		"DETERMINISTIC_ENCRYPTION",
		"DETERMINISTIC_ENCRYPTION_DATE",
		"FIXED_NUMBER",
		"FIXED_STRING",
		"LIBRARY_MASKING_FORMAT",
		"NULL_VALUE",
		"POST_PROCESSING_FUNCTION",
		"PRESERVE_ORIGINAL_DATA",
		"RANDOM_DATE",
		"RANDOM_DECIMAL_NUMBER",
		"RANDOM_DIGITS",
		"RANDOM_LIST",
		"RANDOM_NUMBER",
		"RANDOM_STRING",
		"RANDOM_SUBSTITUTION",
		"REGULAR_EXPRESSION",
		"SHUFFLE",
		"SQL_EXPRESSION",
		"SUBSTRING",
		"TRUNCATE_TABLE",
		"USER_DEFINED_FUNCTION",
	}
}

// GetMappingFormatEntryTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFormatEntryTypeEnum(val string) (FormatEntryTypeEnum, bool) {
	mappingFormatEntryTypeEnumIgnoreCase := make(map[string]FormatEntryTypeEnum)
	for k, v := range mappingFormatEntryTypeEnum {
		mappingFormatEntryTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingFormatEntryTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
