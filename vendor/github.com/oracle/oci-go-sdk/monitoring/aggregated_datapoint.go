// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Monitoring API
//
// Use the Monitoring API to manage metric queries and alarms for assessing the health, capacity, and performance of your cloud resources.
// Endpoints vary by operation. For PostMetric, use the `telemetry-ingestion` endpoints; for all other operations, use the `telemetry` endpoints.
// For information about monitoring, see Monitoring Overview (https://docs.cloud.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm).
//

package monitoring

import (
	"github.com/oracle/oci-go-sdk/common"
)

// AggregatedDatapoint A timestamp-value pair returned for the specified request.
type AggregatedDatapoint struct {

	// The date and time associated with the value of this data point. Format defined by RFC3339.
	// Example: `2019-02-01T01:02:29.600Z`
	Timestamp *common.SDKTime `mandatory:"true" json:"timestamp"`

	// Numeric value of the metric.
	// Example: `10.4`
	Value *float64 `mandatory:"true" json:"value"`
}

func (m AggregatedDatapoint) String() string {
	return common.PointerString(m)
}
