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

// EstimateReleaseDataSizeResult This is the size and time range of data to be released
type EstimateReleaseDataSizeResult struct {

	// This is the end of the time range of data to be released.  timeDataStarted and timeDataEnded delineate
	// the time range of the recalled data to be released.  They may not be exact the same as the
	// parameters in the request input (EstimateReleaseDataSizeDetails).
	TimeDataEnded *common.SDKTime `mandatory:"true" json:"timeDataEnded"`

	// This is the start of the time range of data to be released
	TimeDataStarted *common.SDKTime `mandatory:"true" json:"timeDataStarted"`

	// This is the size in bytes
	SizeInBytes *int64 `mandatory:"true" json:"sizeInBytes"`
}

func (m EstimateReleaseDataSizeResult) String() string {
	return common.PointerString(m)
}
