// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// LogLocationBucketDetails Details to access log file in the specified Object Storage bucket, if any.
type LogLocationBucketDetails struct {

	// Name of the bucket containing the log file.
	BucketName *string `mandatory:"true" json:"bucketName"`

	// Object Storage namespace.
	Namespace *string `mandatory:"true" json:"namespace"`

	// Log object name.
	ObjectName *string `mandatory:"true" json:"objectName"`
}

func (m LogLocationBucketDetails) String() string {
	return common.PointerString(m)
}
