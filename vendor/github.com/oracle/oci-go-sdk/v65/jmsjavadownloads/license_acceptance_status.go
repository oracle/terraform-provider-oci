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

// LicenseAcceptanceStatusEnum Enum with underlying type: string
type LicenseAcceptanceStatusEnum string

// Set of constants representing the allowable values for LicenseAcceptanceStatusEnum
const (
	LicenseAcceptanceStatusAccepted LicenseAcceptanceStatusEnum = "ACCEPTED"
	LicenseAcceptanceStatusRevoked  LicenseAcceptanceStatusEnum = "REVOKED"
)

var mappingLicenseAcceptanceStatusEnum = map[string]LicenseAcceptanceStatusEnum{
	"ACCEPTED": LicenseAcceptanceStatusAccepted,
	"REVOKED":  LicenseAcceptanceStatusRevoked,
}

var mappingLicenseAcceptanceStatusEnumLowerCase = map[string]LicenseAcceptanceStatusEnum{
	"accepted": LicenseAcceptanceStatusAccepted,
	"revoked":  LicenseAcceptanceStatusRevoked,
}

// GetLicenseAcceptanceStatusEnumValues Enumerates the set of values for LicenseAcceptanceStatusEnum
func GetLicenseAcceptanceStatusEnumValues() []LicenseAcceptanceStatusEnum {
	values := make([]LicenseAcceptanceStatusEnum, 0)
	for _, v := range mappingLicenseAcceptanceStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetLicenseAcceptanceStatusEnumStringValues Enumerates the set of values in String for LicenseAcceptanceStatusEnum
func GetLicenseAcceptanceStatusEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"REVOKED",
	}
}

// GetMappingLicenseAcceptanceStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLicenseAcceptanceStatusEnum(val string) (LicenseAcceptanceStatusEnum, bool) {
	enum, ok := mappingLicenseAcceptanceStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
