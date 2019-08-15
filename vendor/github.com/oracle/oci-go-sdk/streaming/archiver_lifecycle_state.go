// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Streaming Service API
//
// The API for the Streaming Service.
//

package streaming

// ArchiverLifecycleStateEnum Enum with underlying type: string
type ArchiverLifecycleStateEnum string

// Set of constants representing the allowable values for ArchiverLifecycleStateEnum
const (
	ArchiverLifecycleStateCreating ArchiverLifecycleStateEnum = "CREATING"
	ArchiverLifecycleStateStopped  ArchiverLifecycleStateEnum = "STOPPED"
	ArchiverLifecycleStateStarting ArchiverLifecycleStateEnum = "STARTING"
	ArchiverLifecycleStateRunning  ArchiverLifecycleStateEnum = "RUNNING"
	ArchiverLifecycleStateStopping ArchiverLifecycleStateEnum = "STOPPING"
	ArchiverLifecycleStateUpdating ArchiverLifecycleStateEnum = "UPDATING"
)

var mappingArchiverLifecycleState = map[string]ArchiverLifecycleStateEnum{
	"CREATING": ArchiverLifecycleStateCreating,
	"STOPPED":  ArchiverLifecycleStateStopped,
	"STARTING": ArchiverLifecycleStateStarting,
	"RUNNING":  ArchiverLifecycleStateRunning,
	"STOPPING": ArchiverLifecycleStateStopping,
	"UPDATING": ArchiverLifecycleStateUpdating,
}

// GetArchiverLifecycleStateEnumValues Enumerates the set of values for ArchiverLifecycleStateEnum
func GetArchiverLifecycleStateEnumValues() []ArchiverLifecycleStateEnum {
	values := make([]ArchiverLifecycleStateEnum, 0)
	for _, v := range mappingArchiverLifecycleState {
		values = append(values, v)
	}
	return values
}
