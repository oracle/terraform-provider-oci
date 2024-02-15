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

// AdhocQueryResultStateEnum Enum with underlying type: string
type AdhocQueryResultStateEnum string

// Set of constants representing the allowable values for AdhocQueryResultStateEnum
const (
	AdhocQueryResultStateAccepted           AdhocQueryResultStateEnum = "ACCEPTED"
	AdhocQueryResultStateCreated            AdhocQueryResultStateEnum = "CREATED"
	AdhocQueryResultStateInProgress         AdhocQueryResultStateEnum = "IN_PROGRESS"
	AdhocQueryResultStateCompleted          AdhocQueryResultStateEnum = "COMPLETED"
	AdhocQueryResultStatePartiallyCompleted AdhocQueryResultStateEnum = "PARTIALLY_COMPLETED"
	AdhocQueryResultStateFailed             AdhocQueryResultStateEnum = "FAILED"
	AdhocQueryResultStateExpired            AdhocQueryResultStateEnum = "EXPIRED"
)

var mappingAdhocQueryResultStateEnum = map[string]AdhocQueryResultStateEnum{
	"ACCEPTED":            AdhocQueryResultStateAccepted,
	"CREATED":             AdhocQueryResultStateCreated,
	"IN_PROGRESS":         AdhocQueryResultStateInProgress,
	"COMPLETED":           AdhocQueryResultStateCompleted,
	"PARTIALLY_COMPLETED": AdhocQueryResultStatePartiallyCompleted,
	"FAILED":              AdhocQueryResultStateFailed,
	"EXPIRED":             AdhocQueryResultStateExpired,
}

var mappingAdhocQueryResultStateEnumLowerCase = map[string]AdhocQueryResultStateEnum{
	"accepted":            AdhocQueryResultStateAccepted,
	"created":             AdhocQueryResultStateCreated,
	"in_progress":         AdhocQueryResultStateInProgress,
	"completed":           AdhocQueryResultStateCompleted,
	"partially_completed": AdhocQueryResultStatePartiallyCompleted,
	"failed":              AdhocQueryResultStateFailed,
	"expired":             AdhocQueryResultStateExpired,
}

// GetAdhocQueryResultStateEnumValues Enumerates the set of values for AdhocQueryResultStateEnum
func GetAdhocQueryResultStateEnumValues() []AdhocQueryResultStateEnum {
	values := make([]AdhocQueryResultStateEnum, 0)
	for _, v := range mappingAdhocQueryResultStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAdhocQueryResultStateEnumStringValues Enumerates the set of values in String for AdhocQueryResultStateEnum
func GetAdhocQueryResultStateEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"CREATED",
		"IN_PROGRESS",
		"COMPLETED",
		"PARTIALLY_COMPLETED",
		"FAILED",
		"EXPIRED",
	}
}

// GetMappingAdhocQueryResultStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAdhocQueryResultStateEnum(val string) (AdhocQueryResultStateEnum, bool) {
	enum, ok := mappingAdhocQueryResultStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
