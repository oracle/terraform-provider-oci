// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"strings"
)

// LicenseModelEnum Enum with underlying type: string
type LicenseModelEnum string

// Set of constants representing the allowable values for LicenseModelEnum
const (
	LicenseModelLicenseIncluded     LicenseModelEnum = "LICENSE_INCLUDED"
	LicenseModelBringYourOwnLicense LicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingLicenseModelEnum = map[string]LicenseModelEnum{
	"LICENSE_INCLUDED":       LicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": LicenseModelBringYourOwnLicense,
}

var mappingLicenseModelEnumLowerCase = map[string]LicenseModelEnum{
	"license_included":       LicenseModelLicenseIncluded,
	"bring_your_own_license": LicenseModelBringYourOwnLicense,
}

// GetLicenseModelEnumValues Enumerates the set of values for LicenseModelEnum
func GetLicenseModelEnumValues() []LicenseModelEnum {
	values := make([]LicenseModelEnum, 0)
	for _, v := range mappingLicenseModelEnum {
		values = append(values, v)
	}
	return values
}

// GetLicenseModelEnumStringValues Enumerates the set of values in String for LicenseModelEnum
func GetLicenseModelEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}

// GetMappingLicenseModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLicenseModelEnum(val string) (LicenseModelEnum, bool) {
	enum, ok := mappingLicenseModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
