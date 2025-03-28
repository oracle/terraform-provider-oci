// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// JavaFamilySortByEnum Enum with underlying type: string
type JavaFamilySortByEnum string

// Set of constants representing the allowable values for JavaFamilySortByEnum
const (
	JavaFamilySortByFamilyVersion        JavaFamilySortByEnum = "familyVersion"
	JavaFamilySortByEndOfSupportLifeDate JavaFamilySortByEnum = "endOfSupportLifeDate"
	JavaFamilySortBySupportType          JavaFamilySortByEnum = "supportType"
)

var mappingJavaFamilySortByEnum = map[string]JavaFamilySortByEnum{
	"familyVersion":        JavaFamilySortByFamilyVersion,
	"endOfSupportLifeDate": JavaFamilySortByEndOfSupportLifeDate,
	"supportType":          JavaFamilySortBySupportType,
}

var mappingJavaFamilySortByEnumLowerCase = map[string]JavaFamilySortByEnum{
	"familyversion":        JavaFamilySortByFamilyVersion,
	"endofsupportlifedate": JavaFamilySortByEndOfSupportLifeDate,
	"supporttype":          JavaFamilySortBySupportType,
}

// GetJavaFamilySortByEnumValues Enumerates the set of values for JavaFamilySortByEnum
func GetJavaFamilySortByEnumValues() []JavaFamilySortByEnum {
	values := make([]JavaFamilySortByEnum, 0)
	for _, v := range mappingJavaFamilySortByEnum {
		values = append(values, v)
	}
	return values
}

// GetJavaFamilySortByEnumStringValues Enumerates the set of values in String for JavaFamilySortByEnum
func GetJavaFamilySortByEnumStringValues() []string {
	return []string{
		"familyVersion",
		"endOfSupportLifeDate",
		"supportType",
	}
}

// GetMappingJavaFamilySortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJavaFamilySortByEnum(val string) (JavaFamilySortByEnum, bool) {
	enum, ok := mappingJavaFamilySortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
