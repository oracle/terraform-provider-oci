// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Streaming Service API
//
// The API for the Streaming Service.
//

package streaming

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Archiver Represents the current state of the stream archiver.
type Archiver struct {

	// Time when the resource was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The state of the stream archiver.
	LifecycleState ArchiverLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The name of the bucket.
	BucketName *string `mandatory:"true" json:"bucketName"`

	// The flag to create a new bucket or use existing one.
	UseExistingBucket *bool `mandatory:"false" json:"useExistingBucket"`

	// The start message.
	StartPosition ArchiverStartPositionEnum `mandatory:"false" json:"startPosition,omitempty"`

	// The batch rollover size in megabytes.
	BatchRolloverSizeInMBs *int `mandatory:"false" json:"batchRolloverSizeInMBs"`

	// The rollover time in seconds.
	BatchRolloverTimeInSeconds *int `mandatory:"false" json:"batchRolloverTimeInSeconds"`

	Error *ArchiverError `mandatory:"false" json:"error"`
}

func (m Archiver) String() string {
	return common.PointerString(m)
}

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
