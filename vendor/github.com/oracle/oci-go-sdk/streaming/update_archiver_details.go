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

// UpdateArchiverDetails The update stream archiver parameters.
type UpdateArchiverDetails struct {

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
}

func (m UpdateArchiverDetails) String() string {
	return common.PointerString(m)
}

// UpdateArchiverDetailsStartPositionEnum is an alias to type: ArchiverStartPositionEnum
// Consider using ArchiverStartPositionEnum instead
// Deprecated
type UpdateArchiverDetailsStartPositionEnum = ArchiverStartPositionEnum

// Set of constants representing the allowable values for ArchiverStartPositionEnum
// Deprecated
const (
	UpdateArchiverDetailsStartPositionLatest      ArchiverStartPositionEnum = "LATEST"
	UpdateArchiverDetailsStartPositionTrimHorizon ArchiverStartPositionEnum = "TRIM_HORIZON"
)

// GetUpdateArchiverDetailsStartPositionEnumValues Enumerates the set of values for ArchiverStartPositionEnum
// Consider using GetArchiverStartPositionEnumValue
// Deprecated
var GetUpdateArchiverDetailsStartPositionEnumValues = GetArchiverStartPositionEnumValues
