// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// ndcs-control-plane API
//
// The control plane API for NoSQL Database Cloud Service HTTPS
// provides endpoints to perform NDCS operations, including creation
// and deletion of tables and indexes; population and access of data
// in tables; and access of table usage metrics.
//

package nosql

import (
	"github.com/oracle/oci-go-sdk/common"
)

// TableUsageSummary TableUsageSummary represents a single usage record, or slice, that includes
// information about read and write throughput consumed during that period
// as well as the current information regarding storage capacity. In
// addition the count of throttling exceptions for the period is reported.
type TableUsageSummary struct {

	// The length of the sampling period.
	SecondsInPeriod *int `mandatory:"false" json:"secondsInPeriod"`

	// Read throughput during the sampling period.
	ReadUnits *int `mandatory:"false" json:"readUnits"`

	// Write throughput during the sampling period.
	WriteUnits *int `mandatory:"false" json:"writeUnits"`

	// The size of the table, in GB.
	StorageInGBs *int `mandatory:"false" json:"storageInGBs"`

	// The number of times reads were throttled due to exceeding
	// the read throughput limit.
	ReadThrottleCount *int `mandatory:"false" json:"readThrottleCount"`

	// The number of times writes were throttled due to exceeding
	// the write throughput limit.
	WriteThrottleCount *int `mandatory:"false" json:"writeThrottleCount"`

	// The number of times writes were throttled because the table
	// exceeded its size limit.
	StorageThrottleCount *int `mandatory:"false" json:"storageThrottleCount"`
}

func (m TableUsageSummary) String() string {
	return common.PointerString(m)
}
