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

// LicenseAcceptanceSortByEnum Enum with underlying type: string
type LicenseAcceptanceSortByEnum string

// Set of constants representing the allowable values for LicenseAcceptanceSortByEnum
const (
	LicenseAcceptanceSortByTimeAccepted            LicenseAcceptanceSortByEnum = "timeAccepted"
	LicenseAcceptanceSortByTimeLastUpdated         LicenseAcceptanceSortByEnum = "timeLastUpdated"
	LicenseAcceptanceSortByLicenseAcceptanceStatus LicenseAcceptanceSortByEnum = "licenseAcceptanceStatus"
)

var mappingLicenseAcceptanceSortByEnum = map[string]LicenseAcceptanceSortByEnum{
	"timeAccepted":            LicenseAcceptanceSortByTimeAccepted,
	"timeLastUpdated":         LicenseAcceptanceSortByTimeLastUpdated,
	"licenseAcceptanceStatus": LicenseAcceptanceSortByLicenseAcceptanceStatus,
}

var mappingLicenseAcceptanceSortByEnumLowerCase = map[string]LicenseAcceptanceSortByEnum{
	"timeaccepted":            LicenseAcceptanceSortByTimeAccepted,
	"timelastupdated":         LicenseAcceptanceSortByTimeLastUpdated,
	"licenseacceptancestatus": LicenseAcceptanceSortByLicenseAcceptanceStatus,
}

// GetLicenseAcceptanceSortByEnumValues Enumerates the set of values for LicenseAcceptanceSortByEnum
func GetLicenseAcceptanceSortByEnumValues() []LicenseAcceptanceSortByEnum {
	values := make([]LicenseAcceptanceSortByEnum, 0)
	for _, v := range mappingLicenseAcceptanceSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetLicenseAcceptanceSortByEnumStringValues Enumerates the set of values in String for LicenseAcceptanceSortByEnum
func GetLicenseAcceptanceSortByEnumStringValues() []string {
	return []string{
		"timeAccepted",
		"timeLastUpdated",
		"licenseAcceptanceStatus",
	}
}

// GetMappingLicenseAcceptanceSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLicenseAcceptanceSortByEnum(val string) (LicenseAcceptanceSortByEnum, bool) {
	enum, ok := mappingLicenseAcceptanceSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
