// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

// EmBridgeLatestImportProcessingStatusEnum Enum with underlying type: string
type EmBridgeLatestImportProcessingStatusEnum string

// Set of constants representing the allowable values for EmBridgeLatestImportProcessingStatusEnum
const (
	EmBridgeLatestImportProcessingStatusNotStarted     EmBridgeLatestImportProcessingStatusEnum = "NOT_STARTED"
	EmBridgeLatestImportProcessingStatusSuccess        EmBridgeLatestImportProcessingStatusEnum = "SUCCESS"
	EmBridgeLatestImportProcessingStatusInProgress     EmBridgeLatestImportProcessingStatusEnum = "IN_PROGRESS"
	EmBridgeLatestImportProcessingStatusFailed         EmBridgeLatestImportProcessingStatusEnum = "FAILED"
	EmBridgeLatestImportProcessingStatusPartialSuccess EmBridgeLatestImportProcessingStatusEnum = "PARTIAL_SUCCESS"
)

var mappingEmBridgeLatestImportProcessingStatus = map[string]EmBridgeLatestImportProcessingStatusEnum{
	"NOT_STARTED":     EmBridgeLatestImportProcessingStatusNotStarted,
	"SUCCESS":         EmBridgeLatestImportProcessingStatusSuccess,
	"IN_PROGRESS":     EmBridgeLatestImportProcessingStatusInProgress,
	"FAILED":          EmBridgeLatestImportProcessingStatusFailed,
	"PARTIAL_SUCCESS": EmBridgeLatestImportProcessingStatusPartialSuccess,
}

// GetEmBridgeLatestImportProcessingStatusEnumValues Enumerates the set of values for EmBridgeLatestImportProcessingStatusEnum
func GetEmBridgeLatestImportProcessingStatusEnumValues() []EmBridgeLatestImportProcessingStatusEnum {
	values := make([]EmBridgeLatestImportProcessingStatusEnum, 0)
	for _, v := range mappingEmBridgeLatestImportProcessingStatus {
		values = append(values, v)
	}
	return values
}
