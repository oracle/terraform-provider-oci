// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud VMware Solution API
//
// Use the Oracle Cloud VMware API to create SDDCs and manage ESXi hosts and software.
// For more information, see Oracle Cloud VMware Solution (https://docs.cloud.oracle.com/iaas/Content/VMware/Concepts/ocvsoverview.htm).
//

package ocvp

import (
	"strings"
)

// SddcTypesEnum Enum with underlying type: string
type SddcTypesEnum string

// Set of constants representing the allowable values for SddcTypesEnum
const (
	SddcTypesProduction    SddcTypesEnum = "PRODUCTION"
	SddcTypesNonProduction SddcTypesEnum = "NON_PRODUCTION"
)

var mappingSddcTypesEnum = map[string]SddcTypesEnum{
	"PRODUCTION":     SddcTypesProduction,
	"NON_PRODUCTION": SddcTypesNonProduction,
}

var mappingSddcTypesEnumLowerCase = map[string]SddcTypesEnum{
	"production":     SddcTypesProduction,
	"non_production": SddcTypesNonProduction,
}

// GetSddcTypesEnumValues Enumerates the set of values for SddcTypesEnum
func GetSddcTypesEnumValues() []SddcTypesEnum {
	values := make([]SddcTypesEnum, 0)
	for _, v := range mappingSddcTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetSddcTypesEnumStringValues Enumerates the set of values in String for SddcTypesEnum
func GetSddcTypesEnumStringValues() []string {
	return []string{
		"PRODUCTION",
		"NON_PRODUCTION",
	}
}

// GetMappingSddcTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSddcTypesEnum(val string) (SddcTypesEnum, bool) {
	enum, ok := mappingSddcTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
