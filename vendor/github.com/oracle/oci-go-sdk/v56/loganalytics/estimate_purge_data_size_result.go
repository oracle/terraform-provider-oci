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

// EstimatePurgeDataSizeResult purge data size in bytes
type EstimatePurgeDataSizeResult struct {

	// This is the size of data to be purged in bytes
	PurgeDataSizeInBytes *int64 `mandatory:"true" json:"purgeDataSizeInBytes"`
}

func (m EstimatePurgeDataSizeResult) String() string {
	return common.PointerString(m)
}
