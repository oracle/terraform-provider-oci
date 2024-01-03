// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management API
//
// API for the OS Management service. Use these API operations for working
// with Managed instances and Managed instance groups.
//

package osmanagement

import (
	"strings"
)

// IsEligibleForInstallationEnum Enum with underlying type: string
type IsEligibleForInstallationEnum string

// Set of constants representing the allowable values for IsEligibleForInstallationEnum
const (
	IsEligibleForInstallationInstallable    IsEligibleForInstallationEnum = "INSTALLABLE"
	IsEligibleForInstallationNotInstallable IsEligibleForInstallationEnum = "NOT_INSTALLABLE"
	IsEligibleForInstallationUnknown        IsEligibleForInstallationEnum = "UNKNOWN"
)

var mappingIsEligibleForInstallationEnum = map[string]IsEligibleForInstallationEnum{
	"INSTALLABLE":     IsEligibleForInstallationInstallable,
	"NOT_INSTALLABLE": IsEligibleForInstallationNotInstallable,
	"UNKNOWN":         IsEligibleForInstallationUnknown,
}

var mappingIsEligibleForInstallationEnumLowerCase = map[string]IsEligibleForInstallationEnum{
	"installable":     IsEligibleForInstallationInstallable,
	"not_installable": IsEligibleForInstallationNotInstallable,
	"unknown":         IsEligibleForInstallationUnknown,
}

// GetIsEligibleForInstallationEnumValues Enumerates the set of values for IsEligibleForInstallationEnum
func GetIsEligibleForInstallationEnumValues() []IsEligibleForInstallationEnum {
	values := make([]IsEligibleForInstallationEnum, 0)
	for _, v := range mappingIsEligibleForInstallationEnum {
		values = append(values, v)
	}
	return values
}

// GetIsEligibleForInstallationEnumStringValues Enumerates the set of values in String for IsEligibleForInstallationEnum
func GetIsEligibleForInstallationEnumStringValues() []string {
	return []string{
		"INSTALLABLE",
		"NOT_INSTALLABLE",
		"UNKNOWN",
	}
}

// GetMappingIsEligibleForInstallationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIsEligibleForInstallationEnum(val string) (IsEligibleForInstallationEnum, bool) {
	enum, ok := mappingIsEligibleForInstallationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
