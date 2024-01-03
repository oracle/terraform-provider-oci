// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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
	AuditTrailStatusStarting         AuditTrailStatusEnum = "STARTING"
	AuditTrailStatusCollecting       AuditTrailStatusEnum = "COLLECTING"
	AuditTrailStatusRecovering       AuditTrailStatusEnum = "RECOVERING"
	AuditTrailStatusIdle             AuditTrailStatusEnum = "IDLE"
	AuditTrailStatusStopping         AuditTrailStatusEnum = "STOPPING"
	AuditTrailStatusStopped          AuditTrailStatusEnum = "STOPPED"
	AuditTrailStatusResuming         AuditTrailStatusEnum = "RESUMING"
	AuditTrailStatusRetrying         AuditTrailStatusEnum = "RETRYING"
	AuditTrailStatusNotStarted       AuditTrailStatusEnum = "NOT_STARTED"
	AuditTrailStatusStoppedNeedsAttn AuditTrailStatusEnum = "STOPPED_NEEDS_ATTN"
	AuditTrailStatusStoppedFailed    AuditTrailStatusEnum = "STOPPED_FAILED"
)

var mappingAuditTrailStatusEnum = map[string]AuditTrailStatusEnum{
	"STARTING":           AuditTrailStatusStarting,
	"COLLECTING":         AuditTrailStatusCollecting,
	"RECOVERING":         AuditTrailStatusRecovering,
	"IDLE":               AuditTrailStatusIdle,
	"STOPPING":           AuditTrailStatusStopping,
	"STOPPED":            AuditTrailStatusStopped,
	"RESUMING":           AuditTrailStatusResuming,
	"RETRYING":           AuditTrailStatusRetrying,
	"NOT_STARTED":        AuditTrailStatusNotStarted,
	"STOPPED_NEEDS_ATTN": AuditTrailStatusStoppedNeedsAttn,
	"STOPPED_FAILED":     AuditTrailStatusStoppedFailed,
}

var mappingAuditTrailStatusEnumLowerCase = map[string]AuditTrailStatusEnum{
	"starting":           AuditTrailStatusStarting,
	"collecting":         AuditTrailStatusCollecting,
	"recovering":         AuditTrailStatusRecovering,
	"idle":               AuditTrailStatusIdle,
	"stopping":           AuditTrailStatusStopping,
	"stopped":            AuditTrailStatusStopped,
	"resuming":           AuditTrailStatusResuming,
	"retrying":           AuditTrailStatusRetrying,
	"not_started":        AuditTrailStatusNotStarted,
	"stopped_needs_attn": AuditTrailStatusStoppedNeedsAttn,
	"stopped_failed":     AuditTrailStatusStoppedFailed,
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
		"NOT_STARTED",
		"STOPPED_NEEDS_ATTN",
		"STOPPED_FAILED",
	}
}

// GetMappingAuditTrailStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAuditTrailStatusEnum(val string) (AuditTrailStatusEnum, bool) {
	enum, ok := mappingAuditTrailStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
