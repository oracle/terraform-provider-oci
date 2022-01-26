// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// UnprocessedDataBucket Configuration details of the bucket that stores unprocessed payloads.
type UnprocessedDataBucket struct {

	// Object Storage namespace.
	Namespace *string `mandatory:"true" json:"namespace"`

	// Name of the Object Storage bucket.
	BucketName *string `mandatory:"true" json:"bucketName"`

	// Flag that specifies if this configuration is enabled or not.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// The time when this record is created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The latest time when this record is updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m UnprocessedDataBucket) String() string {
	return common.PointerString(m)
}
