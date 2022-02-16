// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"strings"
)

// AuditTrailStatusEnum Enum with underlying type: string
type AuditTrailStatusEnum string

// Set of constants representing the allowable values for AuditTrailStatusEnum
const (
	AuditTrailStatusStarting   AuditTrailStatusEnum = "STARTING"
	AuditTrailStatusCollecting AuditTrailStatusEnum = "COLLECTING"
	AuditTrailStatusRecovering AuditTrailStatusEnum = "RECOVERING"
	AuditTrailStatusIdle       AuditTrailStatusEnum = "IDLE"
	AuditTrailStatusStopping   AuditTrailStatusEnum = "STOPPING"
	AuditTrailStatusStopped    AuditTrailStatusEnum = "STOPPED"
	AuditTrailStatusResuming   AuditTrailStatusEnum = "RESUMING"
	AuditTrailStatusRetrying   AuditTrailStatusEnum = "RETRYING"
)

var mappingAuditTrailStatusEnum = map[string]AuditTrailStatusEnum{
	"STARTING":   AuditTrailStatusStarting,
	"COLLECTING": AuditTrailStatusCollecting,
	"RECOVERING": AuditTrailStatusRecovering,
	"IDLE":       AuditTrailStatusIdle,
	"STOPPING":   AuditTrailStatusStopping,
	"STOPPED":    AuditTrailStatusStopped,
	"RESUMING":   AuditTrailStatusResuming,
	"RETRYING":   AuditTrailStatusRetrying,
}

// GetAuditTrailStatusEnumValues Enumerates the set of values for AuditTrailStatusEnum
func GetAuditTrailStatusEnumValues() []AuditTrailStatusEnum {
	values := make([]AuditTrailStatusEnum, 0)
	for _, v := range mappingAuditTrailStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetAuditTrailStatusEnumStringValues Enumerates the set of values in String for AuditTrailStatusEnum
func GetAuditTrailStatusEnumStringValues() []string {
	return []string{
		"STARTING",
		"COLLECTING",
		"RECOVERING",
		"IDLE",
		"STOPPING",
		"STOPPED",
		"RESUMING",
		"RETRYING",
	}
}

// GetMappingAuditTrailStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAuditTrailStatusEnum(val string) (AuditTrailStatusEnum, bool) {
	mappingAuditTrailStatusEnumIgnoreCase := make(map[string]AuditTrailStatusEnum)
	for k, v := range mappingAuditTrailStatusEnum {
		mappingAuditTrailStatusEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingAuditTrailStatusEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
