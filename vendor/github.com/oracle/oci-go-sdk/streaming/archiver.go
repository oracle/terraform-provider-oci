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
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The state of the stream archiver.
	LifecycleState ArchiverLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The name of the bucket.
	BucketName *string `mandatory:"false" json:"bucketName"`

	// The flag to create a new bucket or use existing one.
	UseExistingBucket *bool `mandatory:"false" json:"useExistingBucket"`

	// The start message.
	StartPosition ArchiverStartPositionEnum `mandatory:"false" json:"startPosition,omitempty"`

	// The batch rollover size in megabytes.
	BatchRolloverSizeInMBs *int `mandatory:"false" json:"batchRolloverSizeInMBs"`

	// The rollover time in seconds.
	BatchRolloverTimeInSeconds *int `mandatory:"false" json:"batchRolloverTimeInSeconds"`

	// If an operation failed this property contained the last error occurred.
	Error *ArchiverError `mandatory:"false" json:"error"`
}

func (m Archiver) String() string {
	return common.PointerString(m)
}
