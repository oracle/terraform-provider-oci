// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// LicenseTypeEnum Enum with underlying type: string
type LicenseTypeEnum string

// Set of constants representing the allowable values for LicenseTypeEnum
const (
	LicenseTypeOtn        LicenseTypeEnum = "OTN"
	LicenseTypeNftc       LicenseTypeEnum = "NFTC"
	LicenseTypeRestricted LicenseTypeEnum = "RESTRICTED"
)

var mappingLicenseTypeEnum = map[string]LicenseTypeEnum{
	"OTN":        LicenseTypeOtn,
	"NFTC":       LicenseTypeNftc,
	"RESTRICTED": LicenseTypeRestricted,
}

var mappingLicenseTypeEnumLowerCase = map[string]LicenseTypeEnum{
	"otn":        LicenseTypeOtn,
	"nftc":       LicenseTypeNftc,
	"restricted": LicenseTypeRestricted,
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
		"OTN",
		"NFTC",
		"RESTRICTED",
	}
}

// GetMappingLicenseTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLicenseTypeEnum(val string) (LicenseTypeEnum, bool) {
	enum, ok := mappingLicenseTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
