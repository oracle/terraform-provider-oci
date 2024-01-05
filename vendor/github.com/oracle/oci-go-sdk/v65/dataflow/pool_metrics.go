// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Flow API
//
// Use the Data Flow APIs to run any Apache Spark application at any scale without deploying or managing any infrastructure.
//

package dataflow

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PoolMetrics A collection of metrics related to a particular pool.
type PoolMetrics struct {

	// The last time this pool was started.
	TimeLastStarted *common.SDKTime `mandatory:"false" json:"timeLastStarted"`

	// The last time this pool was stopped.
	TimeLastStopped *common.SDKTime `mandatory:"false" json:"timeLastStopped"`

	// The last time a run used this pool.
	TimeLastUsed *common.SDKTime `mandatory:"false" json:"timeLastUsed"`

	// The last time the mertics were updated for this.
	TimeLastMetricsUpdated *common.SDKTime `mandatory:"false" json:"timeLastMetricsUpdated"`

	// The number of runs that are currently running that are using this pool.
	ActiveRunsCount *int64 `mandatory:"false" json:"activeRunsCount"`

	// A count of the nodes that are currently being used for each shape in this pool.
	ActivelyUsedNodeCount []NodeCount `mandatory:"false" json:"activelyUsedNodeCount"`
}

func (m PoolMetrics) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PoolMetrics) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
