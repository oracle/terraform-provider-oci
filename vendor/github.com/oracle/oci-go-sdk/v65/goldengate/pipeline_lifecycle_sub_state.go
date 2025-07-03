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

// PipelineLifecycleSubStateEnum Enum with underlying type: string
type PipelineLifecycleSubStateEnum string

// Set of constants representing the allowable values for PipelineLifecycleSubStateEnum
const (
	PipelineLifecycleSubStateStarting    PipelineLifecycleSubStateEnum = "STARTING"
	PipelineLifecycleSubStateStopping    PipelineLifecycleSubStateEnum = "STOPPING"
	PipelineLifecycleSubStateStopped     PipelineLifecycleSubStateEnum = "STOPPED"
	PipelineLifecycleSubStateMoving      PipelineLifecycleSubStateEnum = "MOVING"
	PipelineLifecycleSubStateRunning     PipelineLifecycleSubStateEnum = "RUNNING"
	PipelineLifecycleSubStatePausing     PipelineLifecycleSubStateEnum = "PAUSING"
	PipelineLifecycleSubStatePaused      PipelineLifecycleSubStateEnum = "PAUSED"
	PipelineLifecycleSubStateStartFailed PipelineLifecycleSubStateEnum = "START_FAILED"
	PipelineLifecycleSubStateStopFailed  PipelineLifecycleSubStateEnum = "STOP_FAILED"
	PipelineLifecycleSubStatePauseFailed PipelineLifecycleSubStateEnum = "PAUSE_FAILED"
)

var mappingPipelineLifecycleSubStateEnum = map[string]PipelineLifecycleSubStateEnum{
	"STARTING":     PipelineLifecycleSubStateStarting,
	"STOPPING":     PipelineLifecycleSubStateStopping,
	"STOPPED":      PipelineLifecycleSubStateStopped,
	"MOVING":       PipelineLifecycleSubStateMoving,
	"RUNNING":      PipelineLifecycleSubStateRunning,
	"PAUSING":      PipelineLifecycleSubStatePausing,
	"PAUSED":       PipelineLifecycleSubStatePaused,
	"START_FAILED": PipelineLifecycleSubStateStartFailed,
	"STOP_FAILED":  PipelineLifecycleSubStateStopFailed,
	"PAUSE_FAILED": PipelineLifecycleSubStatePauseFailed,
}

var mappingPipelineLifecycleSubStateEnumLowerCase = map[string]PipelineLifecycleSubStateEnum{
	"starting":     PipelineLifecycleSubStateStarting,
	"stopping":     PipelineLifecycleSubStateStopping,
	"stopped":      PipelineLifecycleSubStateStopped,
	"moving":       PipelineLifecycleSubStateMoving,
	"running":      PipelineLifecycleSubStateRunning,
	"pausing":      PipelineLifecycleSubStatePausing,
	"paused":       PipelineLifecycleSubStatePaused,
	"start_failed": PipelineLifecycleSubStateStartFailed,
	"stop_failed":  PipelineLifecycleSubStateStopFailed,
	"pause_failed": PipelineLifecycleSubStatePauseFailed,
}

// GetPipelineLifecycleSubStateEnumValues Enumerates the set of values for PipelineLifecycleSubStateEnum
func GetPipelineLifecycleSubStateEnumValues() []PipelineLifecycleSubStateEnum {
	values := make([]PipelineLifecycleSubStateEnum, 0)
	for _, v := range mappingPipelineLifecycleSubStateEnum {
		values = append(values, v)
	}
	return values
}

// GetPipelineLifecycleSubStateEnumStringValues Enumerates the set of values in String for PipelineLifecycleSubStateEnum
func GetPipelineLifecycleSubStateEnumStringValues() []string {
	return []string{
		"STARTING",
		"STOPPING",
		"STOPPED",
		"MOVING",
		"RUNNING",
		"PAUSING",
		"PAUSED",
		"START_FAILED",
		"STOP_FAILED",
		"PAUSE_FAILED",
	}
}

// GetMappingPipelineLifecycleSubStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPipelineLifecycleSubStateEnum(val string) (PipelineLifecycleSubStateEnum, bool) {
	enum, ok := mappingPipelineLifecycleSubStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
