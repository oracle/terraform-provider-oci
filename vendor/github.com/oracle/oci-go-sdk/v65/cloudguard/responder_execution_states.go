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

// ResponderExecutionStatesEnum Enum with underlying type: string
type ResponderExecutionStatesEnum string

// Set of constants representing the allowable values for ResponderExecutionStatesEnum
const (
	ResponderExecutionStatesStarted              ResponderExecutionStatesEnum = "STARTED"
	ResponderExecutionStatesAwaitingConfirmation ResponderExecutionStatesEnum = "AWAITING_CONFIRMATION"
	ResponderExecutionStatesAwaitingInput        ResponderExecutionStatesEnum = "AWAITING_INPUT"
	ResponderExecutionStatesSucceeded            ResponderExecutionStatesEnum = "SUCCEEDED"
	ResponderExecutionStatesFailed               ResponderExecutionStatesEnum = "FAILED"
	ResponderExecutionStatesSkipped              ResponderExecutionStatesEnum = "SKIPPED"
	ResponderExecutionStatesAll                  ResponderExecutionStatesEnum = "ALL"
)

var mappingResponderExecutionStatesEnum = map[string]ResponderExecutionStatesEnum{
	"STARTED":               ResponderExecutionStatesStarted,
	"AWAITING_CONFIRMATION": ResponderExecutionStatesAwaitingConfirmation,
	"AWAITING_INPUT":        ResponderExecutionStatesAwaitingInput,
	"SUCCEEDED":             ResponderExecutionStatesSucceeded,
	"FAILED":                ResponderExecutionStatesFailed,
	"SKIPPED":               ResponderExecutionStatesSkipped,
	"ALL":                   ResponderExecutionStatesAll,
}

var mappingResponderExecutionStatesEnumLowerCase = map[string]ResponderExecutionStatesEnum{
	"started":               ResponderExecutionStatesStarted,
	"awaiting_confirmation": ResponderExecutionStatesAwaitingConfirmation,
	"awaiting_input":        ResponderExecutionStatesAwaitingInput,
	"succeeded":             ResponderExecutionStatesSucceeded,
	"failed":                ResponderExecutionStatesFailed,
	"skipped":               ResponderExecutionStatesSkipped,
	"all":                   ResponderExecutionStatesAll,
}

// GetResponderExecutionStatesEnumValues Enumerates the set of values for ResponderExecutionStatesEnum
func GetResponderExecutionStatesEnumValues() []ResponderExecutionStatesEnum {
	values := make([]ResponderExecutionStatesEnum, 0)
	for _, v := range mappingResponderExecutionStatesEnum {
		values = append(values, v)
	}
	return values
}

// GetResponderExecutionStatesEnumStringValues Enumerates the set of values in String for ResponderExecutionStatesEnum
func GetResponderExecutionStatesEnumStringValues() []string {
	return []string{
		"STARTED",
		"AWAITING_CONFIRMATION",
		"AWAITING_INPUT",
		"SUCCEEDED",
		"FAILED",
		"SKIPPED",
		"ALL",
	}
}

// GetMappingResponderExecutionStatesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResponderExecutionStatesEnum(val string) (ResponderExecutionStatesEnum, bool) {
	enum, ok := mappingResponderExecutionStatesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
