// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Fleets API
//
// The APIs for the Fleet Management (https://docs.oracle.com/en-us/iaas/jms/doc/fleet-management.html) feature of Java Management Service to monitor and manage the usage of Java in your enterprise. Use these APIs to manage fleets, configure managed instances to report to fleets, and gain insights into the Java workloads running on these instances by carrying out basic and advanced features.
//

package jms

import (
	"strings"
)

// JavaReleaseSortByEnum Enum with underlying type: string
type JavaReleaseSortByEnum string

// Set of constants representing the allowable values for JavaReleaseSortByEnum
const (
	JavaReleaseSortByReleaseDate    JavaReleaseSortByEnum = "releaseDate"
	JavaReleaseSortByReleaseVersion JavaReleaseSortByEnum = "releaseVersion"
	JavaReleaseSortByFamilyVersion  JavaReleaseSortByEnum = "familyVersion"
	JavaReleaseSortByLicenseType    JavaReleaseSortByEnum = "licenseType"
)

var mappingJavaReleaseSortByEnum = map[string]JavaReleaseSortByEnum{
	"releaseDate":    JavaReleaseSortByReleaseDate,
	"releaseVersion": JavaReleaseSortByReleaseVersion,
	"familyVersion":  JavaReleaseSortByFamilyVersion,
	"licenseType":    JavaReleaseSortByLicenseType,
}

var mappingJavaReleaseSortByEnumLowerCase = map[string]JavaReleaseSortByEnum{
	"releasedate":    JavaReleaseSortByReleaseDate,
	"releaseversion": JavaReleaseSortByReleaseVersion,
	"familyversion":  JavaReleaseSortByFamilyVersion,
	"licensetype":    JavaReleaseSortByLicenseType,
}

// GetJavaReleaseSortByEnumValues Enumerates the set of values for JavaReleaseSortByEnum
func GetJavaReleaseSortByEnumValues() []JavaReleaseSortByEnum {
	values := make([]JavaReleaseSortByEnum, 0)
	for _, v := range mappingJavaReleaseSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetJavaReleaseSortByEnumStringValues Enumerates the set of values in String for JavaReleaseSortByEnum
func GetJavaReleaseSortByEnumStringValues() []string {
	return []string{
		"releaseDate",
		"releaseVersion",
		"familyVersion",
		"licenseType",
	}
}

// GetMappingJavaReleaseSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJavaReleaseSortByEnum(val string) (JavaReleaseSortByEnum, bool) {
	enum, ok := mappingJavaReleaseSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
