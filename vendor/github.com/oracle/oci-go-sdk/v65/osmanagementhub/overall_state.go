// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for the operating system environments in your private data centers through a single management console. For more information, see Overview of OS Management Hub (https://docs.cloud.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"strings"
)

// OverallStateEnum Enum with underlying type: string
type OverallStateEnum string

// Set of constants representing the allowable values for OverallStateEnum
const (
	OverallStateNormal            OverallStateEnum = "NORMAL"
	OverallStateRegistrationerror OverallStateEnum = "REGISTRATIONERROR"
	OverallStateSyncing           OverallStateEnum = "SYNCING"
	OverallStateSyncfailed        OverallStateEnum = "SYNCFAILED"
	OverallStateWarning           OverallStateEnum = "WARNING"
	OverallStateError             OverallStateEnum = "ERROR"
	OverallStateUnavailable       OverallStateEnum = "UNAVAILABLE"
)

var mappingOverallStateEnum = map[string]OverallStateEnum{
	"NORMAL":            OverallStateNormal,
	"REGISTRATIONERROR": OverallStateRegistrationerror,
	"SYNCING":           OverallStateSyncing,
	"SYNCFAILED":        OverallStateSyncfailed,
	"WARNING":           OverallStateWarning,
	"ERROR":             OverallStateError,
	"UNAVAILABLE":       OverallStateUnavailable,
}

var mappingOverallStateEnumLowerCase = map[string]OverallStateEnum{
	"normal":            OverallStateNormal,
	"registrationerror": OverallStateRegistrationerror,
	"syncing":           OverallStateSyncing,
	"syncfailed":        OverallStateSyncfailed,
	"warning":           OverallStateWarning,
	"error":             OverallStateError,
	"unavailable":       OverallStateUnavailable,
}

// GetOverallStateEnumValues Enumerates the set of values for OverallStateEnum
func GetOverallStateEnumValues() []OverallStateEnum {
	values := make([]OverallStateEnum, 0)
	for _, v := range mappingOverallStateEnum {
		values = append(values, v)
	}
	return values
}

// GetOverallStateEnumStringValues Enumerates the set of values in String for OverallStateEnum
func GetOverallStateEnumStringValues() []string {
	return []string{
		"NORMAL",
		"REGISTRATIONERROR",
		"SYNCING",
		"SYNCFAILED",
		"WARNING",
		"ERROR",
		"UNAVAILABLE",
	}
}

// GetMappingOverallStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOverallStateEnum(val string) (OverallStateEnum, bool) {
	enum, ok := mappingOverallStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
