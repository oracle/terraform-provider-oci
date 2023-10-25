// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// NoSQL Database API
//
// The control plane API for NoSQL Database Cloud Service HTTPS
// provides endpoints to perform NDCS operations, including creation
// and deletion of tables and indexes; population and access of data
// in tables; and access of table usage metrics.
//

package nosql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
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

	// The percentage of allowed per-shard usage for the table shard with the highest usage.
	MaxShardSizeUsageInPercent *int `mandatory:"false" json:"maxShardSizeUsageInPercent"`

	// The time stamp of this usage record.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`
}

func (m TableUsageSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TableUsageSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
