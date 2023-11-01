// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for the operating system environments in your private data centers through a single management console. For more information, see Overview of OS Management Hub (https://docs.cloud.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"strings"
)

// ClassificationTypesEnum Enum with underlying type: string
type ClassificationTypesEnum string

// Set of constants representing the allowable values for ClassificationTypesEnum
const (
	ClassificationTypesSecurity    ClassificationTypesEnum = "SECURITY"
	ClassificationTypesBugfix      ClassificationTypesEnum = "BUGFIX"
	ClassificationTypesEnhancement ClassificationTypesEnum = "ENHANCEMENT"
	ClassificationTypesOther       ClassificationTypesEnum = "OTHER"
)

var mappingClassificationTypesEnum = map[string]ClassificationTypesEnum{
	"SECURITY":    ClassificationTypesSecurity,
	"BUGFIX":      ClassificationTypesBugfix,
	"ENHANCEMENT": ClassificationTypesEnhancement,
	"OTHER":       ClassificationTypesOther,
}

var mappingClassificationTypesEnumLowerCase = map[string]ClassificationTypesEnum{
	"security":    ClassificationTypesSecurity,
	"bugfix":      ClassificationTypesBugfix,
	"enhancement": ClassificationTypesEnhancement,
	"other":       ClassificationTypesOther,
}

// GetClassificationTypesEnumValues Enumerates the set of values for ClassificationTypesEnum
func GetClassificationTypesEnumValues() []ClassificationTypesEnum {
	values := make([]ClassificationTypesEnum, 0)
	for _, v := range mappingClassificationTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetClassificationTypesEnumStringValues Enumerates the set of values in String for ClassificationTypesEnum
func GetClassificationTypesEnumStringValues() []string {
	return []string{
		"SECURITY",
		"BUGFIX",
		"ENHANCEMENT",
		"OTHER",
	}
}

// GetMappingClassificationTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingClassificationTypesEnum(val string) (ClassificationTypesEnum, bool) {
	enum, ok := mappingClassificationTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
