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

// CreateArchiverDetails Represents the parameters of the stream archiver.
type CreateArchiverDetails struct {

	// The name of the bucket.
	BucketName *string `mandatory:"true" json:"bucketName"`

	// The flag to create a new bucket or use existing one.
	UseExistingBucket *bool `mandatory:"true" json:"useExistingBucket"`

	// The start message.
	StartPosition ArchiverStartPositionEnum `mandatory:"true" json:"startPosition"`

	// The batch rollover size in megabytes.
	BatchRolloverSizeInMBs *int `mandatory:"true" json:"batchRolloverSizeInMBs"`

	// The rollover time in seconds.
	BatchRolloverTimeInSeconds *int `mandatory:"true" json:"batchRolloverTimeInSeconds"`
}

func (m CreateArchiverDetails) String() string {
	return common.PointerString(m)
}

// CreateArchiverDetailsStartPositionEnum is an alias to type: ArchiverStartPositionEnum
// Consider using ArchiverStartPositionEnum instead
// Deprecated
type CreateArchiverDetailsStartPositionEnum = ArchiverStartPositionEnum

// Set of constants representing the allowable values for ArchiverStartPositionEnum
// Deprecated
const (
	CreateArchiverDetailsStartPositionLatest      ArchiverStartPositionEnum = "LATEST"
	CreateArchiverDetailsStartPositionTrimHorizon ArchiverStartPositionEnum = "TRIM_HORIZON"
)

// GetCreateArchiverDetailsStartPositionEnumValues Enumerates the set of values for ArchiverStartPositionEnum
// Consider using GetArchiverStartPositionEnumValue
// Deprecated
var GetCreateArchiverDetailsStartPositionEnumValues = GetArchiverStartPositionEnumValues
