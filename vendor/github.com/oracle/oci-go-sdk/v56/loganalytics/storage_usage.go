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

// StorageUsage This is the storage usage information of a tenancy in Logan Analytics application
type StorageUsage struct {

	// This is the number of bytes of active data (non-archived)
	ActiveDataSizeInBytes *int64 `mandatory:"true" json:"activeDataSizeInBytes"`

	// This is the number of bytes of archived data in object storage
	ArchivedDataSizeInBytes *int64 `mandatory:"true" json:"archivedDataSizeInBytes"`

	// This is the number of bytes of recalled data from archived in object store
	RecalledArchivedDataSizeInBytes *int64 `mandatory:"true" json:"recalledArchivedDataSizeInBytes"`
}

func (m StorageUsage) String() string {
	return common.PointerString(m)
}
