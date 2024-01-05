// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Analytics API
//
// Analytics API.
//

package analytics

import (
	"strings"
)

// LicenseTypeEnum Enum with underlying type: string
type LicenseTypeEnum string

// Set of constants representing the allowable values for LicenseTypeEnum
const (
	LicenseTypeLicenseIncluded     LicenseTypeEnum = "LICENSE_INCLUDED"
	LicenseTypeBringYourOwnLicense LicenseTypeEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingLicenseTypeEnum = map[string]LicenseTypeEnum{
	"LICENSE_INCLUDED":       LicenseTypeLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": LicenseTypeBringYourOwnLicense,
}

var mappingLicenseTypeEnumLowerCase = map[string]LicenseTypeEnum{
	"license_included":       LicenseTypeLicenseIncluded,
	"bring_your_own_license": LicenseTypeBringYourOwnLicense,
}

// GetLicenseTypeEnumValues Enumerates the set of values for LicenseTypeEnum
func GetLicenseTypeEnumValues() []LicenseTypeEnum {
	values := make([]LicenseTypeEnum, 0)
	for _, v := range mappingLicenseTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetLicenseTypeEnumStringValues Enumerates the set of values in String for LicenseTypeEnum
func GetLicenseTypeEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}

// GetMappingLicenseTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLicenseTypeEnum(val string) (LicenseTypeEnum, bool) {
	enum, ok := mappingLicenseTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
