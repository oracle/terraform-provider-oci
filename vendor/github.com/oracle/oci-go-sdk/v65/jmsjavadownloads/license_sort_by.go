// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Download API
//
// The APIs for the download engine of the Java Management Service.
//

package jmsjavadownloads

import (
	"strings"
)

// LicenseSortByEnum Enum with underlying type: string
type LicenseSortByEnum string

// Set of constants representing the allowable values for LicenseSortByEnum
const (
	LicenseSortByLicenseType LicenseSortByEnum = "licenseType"
	LicenseSortByDisplayName LicenseSortByEnum = "displayName"
)

var mappingLicenseSortByEnum = map[string]LicenseSortByEnum{
	"licenseType": LicenseSortByLicenseType,
	"displayName": LicenseSortByDisplayName,
}

var mappingLicenseSortByEnumLowerCase = map[string]LicenseSortByEnum{
	"licensetype": LicenseSortByLicenseType,
	"displayname": LicenseSortByDisplayName,
}

// GetLicenseSortByEnumValues Enumerates the set of values for LicenseSortByEnum
func GetLicenseSortByEnumValues() []LicenseSortByEnum {
	values := make([]LicenseSortByEnum, 0)
	for _, v := range mappingLicenseSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetLicenseSortByEnumStringValues Enumerates the set of values in String for LicenseSortByEnum
func GetLicenseSortByEnumStringValues() []string {
	return []string{
		"licenseType",
		"displayName",
	}
}

// GetMappingLicenseSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLicenseSortByEnum(val string) (LicenseSortByEnum, bool) {
	enum, ok := mappingLicenseSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
