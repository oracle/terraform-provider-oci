// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"strings"
)

// PipelineDiagnosticStateEnum Enum with underlying type: string
type PipelineDiagnosticStateEnum string

// Set of constants representing the allowable values for PipelineDiagnosticStateEnum
const (
	PipelineDiagnosticStateInProgress PipelineDiagnosticStateEnum = "IN_PROGRESS"
	PipelineDiagnosticStateSucceeded  PipelineDiagnosticStateEnum = "SUCCEEDED"
	PipelineDiagnosticStateFailed     PipelineDiagnosticStateEnum = "FAILED"
)

var mappingPipelineDiagnosticStateEnum = map[string]PipelineDiagnosticStateEnum{
	"IN_PROGRESS": PipelineDiagnosticStateInProgress,
	"SUCCEEDED":   PipelineDiagnosticStateSucceeded,
	"FAILED":      PipelineDiagnosticStateFailed,
}

var mappingPipelineDiagnosticStateEnumLowerCase = map[string]PipelineDiagnosticStateEnum{
	"in_progress": PipelineDiagnosticStateInProgress,
	"succeeded":   PipelineDiagnosticStateSucceeded,
	"failed":      PipelineDiagnosticStateFailed,
}

// GetPipelineDiagnosticStateEnumValues Enumerates the set of values for PipelineDiagnosticStateEnum
func GetPipelineDiagnosticStateEnumValues() []PipelineDiagnosticStateEnum {
	values := make([]PipelineDiagnosticStateEnum, 0)
	for _, v := range mappingPipelineDiagnosticStateEnum {
		values = append(values, v)
	}
	return values
}

// GetPipelineDiagnosticStateEnumStringValues Enumerates the set of values in String for PipelineDiagnosticStateEnum
func GetPipelineDiagnosticStateEnumStringValues() []string {
	return []string{
		"IN_PROGRESS",
		"SUCCEEDED",
		"FAILED",
	}
}

// GetMappingPipelineDiagnosticStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPipelineDiagnosticStateEnum(val string) (PipelineDiagnosticStateEnum, bool) {
	enum, ok := mappingPipelineDiagnosticStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
