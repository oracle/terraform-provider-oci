// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"strings"
)

// LicenseTypeEnum Enum with underlying type: string
type LicenseTypeEnum string

// Set of constants representing the allowable values for LicenseTypeEnum
const (
	LicenseTypeStandardEdition   LicenseTypeEnum = "STANDARD_EDITION"
	LicenseTypeEnterpriseEdition LicenseTypeEnum = "ENTERPRISE_EDITION"
)

var mappingLicenseTypeEnum = map[string]LicenseTypeEnum{
	"STANDARD_EDITION":   LicenseTypeStandardEdition,
	"ENTERPRISE_EDITION": LicenseTypeEnterpriseEdition,
}

var mappingLicenseTypeEnumLowerCase = map[string]LicenseTypeEnum{
	"standard_edition":   LicenseTypeStandardEdition,
	"enterprise_edition": LicenseTypeEnterpriseEdition,
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
		"STANDARD_EDITION",
		"ENTERPRISE_EDITION",
	}
}

// GetMappingLicenseTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLicenseTypeEnum(val string) (LicenseTypeEnum, bool) {
	enum, ok := mappingLicenseTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
