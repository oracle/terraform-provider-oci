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

// JreSecurityStatusEnum Enum with underlying type: string
type JreSecurityStatusEnum string

// Set of constants representing the allowable values for JreSecurityStatusEnum
const (
	JreSecurityStatusEarlyAccess     JreSecurityStatusEnum = "EARLY_ACCESS"
	JreSecurityStatusUnknown         JreSecurityStatusEnum = "UNKNOWN"
	JreSecurityStatusUpToDate        JreSecurityStatusEnum = "UP_TO_DATE"
	JreSecurityStatusUpdateRequired  JreSecurityStatusEnum = "UPDATE_REQUIRED"
	JreSecurityStatusUpgradeRequired JreSecurityStatusEnum = "UPGRADE_REQUIRED"
)

var mappingJreSecurityStatusEnum = map[string]JreSecurityStatusEnum{
	"EARLY_ACCESS":     JreSecurityStatusEarlyAccess,
	"UNKNOWN":          JreSecurityStatusUnknown,
	"UP_TO_DATE":       JreSecurityStatusUpToDate,
	"UPDATE_REQUIRED":  JreSecurityStatusUpdateRequired,
	"UPGRADE_REQUIRED": JreSecurityStatusUpgradeRequired,
}

var mappingJreSecurityStatusEnumLowerCase = map[string]JreSecurityStatusEnum{
	"early_access":     JreSecurityStatusEarlyAccess,
	"unknown":          JreSecurityStatusUnknown,
	"up_to_date":       JreSecurityStatusUpToDate,
	"update_required":  JreSecurityStatusUpdateRequired,
	"upgrade_required": JreSecurityStatusUpgradeRequired,
}

// GetJreSecurityStatusEnumValues Enumerates the set of values for JreSecurityStatusEnum
func GetJreSecurityStatusEnumValues() []JreSecurityStatusEnum {
	values := make([]JreSecurityStatusEnum, 0)
	for _, v := range mappingJreSecurityStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetJreSecurityStatusEnumStringValues Enumerates the set of values in String for JreSecurityStatusEnum
func GetJreSecurityStatusEnumStringValues() []string {
	return []string{
		"EARLY_ACCESS",
		"UNKNOWN",
		"UP_TO_DATE",
		"UPDATE_REQUIRED",
		"UPGRADE_REQUIRED",
	}
}

// GetMappingJreSecurityStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJreSecurityStatusEnum(val string) (JreSecurityStatusEnum, bool) {
	enum, ok := mappingJreSecurityStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
