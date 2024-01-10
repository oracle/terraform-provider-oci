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

// ResponderExecutionStatusEnum Enum with underlying type: string
type ResponderExecutionStatusEnum string

// Set of constants representing the allowable values for ResponderExecutionStatusEnum
const (
	ResponderExecutionStatusStarted              ResponderExecutionStatusEnum = "STARTED"
	ResponderExecutionStatusAwaitingConfirmation ResponderExecutionStatusEnum = "AWAITING_CONFIRMATION"
	ResponderExecutionStatusSucceeded            ResponderExecutionStatusEnum = "SUCCEEDED"
	ResponderExecutionStatusFailed               ResponderExecutionStatusEnum = "FAILED"
	ResponderExecutionStatusSkipped              ResponderExecutionStatusEnum = "SKIPPED"
)

var mappingResponderExecutionStatusEnum = map[string]ResponderExecutionStatusEnum{
	"STARTED":               ResponderExecutionStatusStarted,
	"AWAITING_CONFIRMATION": ResponderExecutionStatusAwaitingConfirmation,
	"SUCCEEDED":             ResponderExecutionStatusSucceeded,
	"FAILED":                ResponderExecutionStatusFailed,
	"SKIPPED":               ResponderExecutionStatusSkipped,
}

var mappingResponderExecutionStatusEnumLowerCase = map[string]ResponderExecutionStatusEnum{
	"started":               ResponderExecutionStatusStarted,
	"awaiting_confirmation": ResponderExecutionStatusAwaitingConfirmation,
	"succeeded":             ResponderExecutionStatusSucceeded,
	"failed":                ResponderExecutionStatusFailed,
	"skipped":               ResponderExecutionStatusSkipped,
}

// GetResponderExecutionStatusEnumValues Enumerates the set of values for ResponderExecutionStatusEnum
func GetResponderExecutionStatusEnumValues() []ResponderExecutionStatusEnum {
	values := make([]ResponderExecutionStatusEnum, 0)
	for _, v := range mappingResponderExecutionStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetResponderExecutionStatusEnumStringValues Enumerates the set of values in String for ResponderExecutionStatusEnum
func GetResponderExecutionStatusEnumStringValues() []string {
	return []string{
		"STARTED",
		"AWAITING_CONFIRMATION",
		"SUCCEEDED",
		"FAILED",
		"SKIPPED",
	}
}

// GetMappingResponderExecutionStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResponderExecutionStatusEnum(val string) (ResponderExecutionStatusEnum, bool) {
	enum, ok := mappingResponderExecutionStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
