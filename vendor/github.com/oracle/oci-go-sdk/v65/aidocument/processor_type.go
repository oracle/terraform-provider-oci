// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Document Understanding API
//
// Document AI helps customers perform various analysis on their documents. If a customer has lots of documents, they can process them in batch using asynchronous API endpoints.
//

package aidocument

import (
	"strings"
)

// ProcessorTypeEnum Enum with underlying type: string
type ProcessorTypeEnum string

// Set of constants representing the allowable values for ProcessorTypeEnum
const (
	ProcessorTypeGeneral ProcessorTypeEnum = "GENERAL"
)

var mappingProcessorTypeEnum = map[string]ProcessorTypeEnum{
	"GENERAL": ProcessorTypeGeneral,
}

var mappingProcessorTypeEnumLowerCase = map[string]ProcessorTypeEnum{
	"general": ProcessorTypeGeneral,
}

// GetProcessorTypeEnumValues Enumerates the set of values for ProcessorTypeEnum
func GetProcessorTypeEnumValues() []ProcessorTypeEnum {
	values := make([]ProcessorTypeEnum, 0)
	for _, v := range mappingProcessorTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetProcessorTypeEnumStringValues Enumerates the set of values in String for ProcessorTypeEnum
func GetProcessorTypeEnumStringValues() []string {
	return []string{
		"GENERAL",
	}
}

// GetMappingProcessorTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingProcessorTypeEnum(val string) (ProcessorTypeEnum, bool) {
	enum, ok := mappingProcessorTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
