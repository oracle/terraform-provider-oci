// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard and Security Zones API
//
// Use the Cloud Guard and Security Zones API to automate processes that you would otherwise perform through the Cloud Guard Console or the Security Zones Console. For more information on these services, see the Cloud Guard (https://docs.cloud.oracle.com/iaas/cloud-guard/home.htm) and Security Zones (https://docs.cloud.oracle.com/iaas/security-zone/home.htm) documentation.
// **Note:** For Cloud Guard, you can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
//

package cloudguard

import (
	"strings"
)

// AdhocQueryStatusEnum Enum with underlying type: string
type AdhocQueryStatusEnum string

// Set of constants representing the allowable values for AdhocQueryStatusEnum
const (
	AdhocQueryStatusCreating           AdhocQueryStatusEnum = "CREATING"
	AdhocQueryStatusCreated            AdhocQueryStatusEnum = "CREATED"
	AdhocQueryStatusInProgress         AdhocQueryStatusEnum = "IN_PROGRESS"
	AdhocQueryStatusPartiallyCompleted AdhocQueryStatusEnum = "PARTIALLY_COMPLETED"
	AdhocQueryStatusExpired            AdhocQueryStatusEnum = "EXPIRED"
	AdhocQueryStatusCompleted          AdhocQueryStatusEnum = "COMPLETED"
	AdhocQueryStatusFailed             AdhocQueryStatusEnum = "FAILED"
)

var mappingAdhocQueryStatusEnum = map[string]AdhocQueryStatusEnum{
	"CREATING":            AdhocQueryStatusCreating,
	"CREATED":             AdhocQueryStatusCreated,
	"IN_PROGRESS":         AdhocQueryStatusInProgress,
	"PARTIALLY_COMPLETED": AdhocQueryStatusPartiallyCompleted,
	"EXPIRED":             AdhocQueryStatusExpired,
	"COMPLETED":           AdhocQueryStatusCompleted,
	"FAILED":              AdhocQueryStatusFailed,
}

var mappingAdhocQueryStatusEnumLowerCase = map[string]AdhocQueryStatusEnum{
	"creating":            AdhocQueryStatusCreating,
	"created":             AdhocQueryStatusCreated,
	"in_progress":         AdhocQueryStatusInProgress,
	"partially_completed": AdhocQueryStatusPartiallyCompleted,
	"expired":             AdhocQueryStatusExpired,
	"completed":           AdhocQueryStatusCompleted,
	"failed":              AdhocQueryStatusFailed,
}

// GetAdhocQueryStatusEnumValues Enumerates the set of values for AdhocQueryStatusEnum
func GetAdhocQueryStatusEnumValues() []AdhocQueryStatusEnum {
	values := make([]AdhocQueryStatusEnum, 0)
	for _, v := range mappingAdhocQueryStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetAdhocQueryStatusEnumStringValues Enumerates the set of values in String for AdhocQueryStatusEnum
func GetAdhocQueryStatusEnumStringValues() []string {
	return []string{
		"CREATING",
		"CREATED",
		"IN_PROGRESS",
		"PARTIALLY_COMPLETED",
		"EXPIRED",
		"COMPLETED",
		"FAILED",
	}
}

// GetMappingAdhocQueryStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAdhocQueryStatusEnum(val string) (AdhocQueryStatusEnum, bool) {
	enum, ok := mappingAdhocQueryStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
