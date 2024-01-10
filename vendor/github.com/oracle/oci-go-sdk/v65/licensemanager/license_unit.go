// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// License Manager API
//
// Use the License Manager API to manage product licenses and license records. For more information, see License Manager Overview (https://docs.cloud.oracle.com/iaas/Content/LicenseManager/Concepts/licensemanageroverview.htm).
//

package licensemanager

import (
	"strings"
)

// LicenseUnitEnum Enum with underlying type: string
type LicenseUnitEnum string

// Set of constants representing the allowable values for LicenseUnitEnum
const (
	LicenseUnitOcpu          LicenseUnitEnum = "OCPU"
	LicenseUnitNamedUserPlus LicenseUnitEnum = "NAMED_USER_PLUS"
	LicenseUnitProcessors    LicenseUnitEnum = "PROCESSORS"
)

var mappingLicenseUnitEnum = map[string]LicenseUnitEnum{
	"OCPU":            LicenseUnitOcpu,
	"NAMED_USER_PLUS": LicenseUnitNamedUserPlus,
	"PROCESSORS":      LicenseUnitProcessors,
}

var mappingLicenseUnitEnumLowerCase = map[string]LicenseUnitEnum{
	"ocpu":            LicenseUnitOcpu,
	"named_user_plus": LicenseUnitNamedUserPlus,
	"processors":      LicenseUnitProcessors,
}

// GetLicenseUnitEnumValues Enumerates the set of values for LicenseUnitEnum
func GetLicenseUnitEnumValues() []LicenseUnitEnum {
	values := make([]LicenseUnitEnum, 0)
	for _, v := range mappingLicenseUnitEnum {
		values = append(values, v)
	}
	return values
}

// GetLicenseUnitEnumStringValues Enumerates the set of values in String for LicenseUnitEnum
func GetLicenseUnitEnumStringValues() []string {
	return []string{
		"OCPU",
		"NAMED_USER_PLUS",
		"PROCESSORS",
	}
}

// GetMappingLicenseUnitEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLicenseUnitEnum(val string) (LicenseUnitEnum, bool) {
	enum, ok := mappingLicenseUnitEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
