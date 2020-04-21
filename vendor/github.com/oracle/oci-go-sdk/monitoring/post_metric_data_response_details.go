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

// PostMetricDataResponseDetails The response object returned from a PostMetricData operation.
type PostMetricDataResponseDetails struct {

	// The number of metric objects that failed input validation.
	FailedMetricsCount *int `mandatory:"true" json:"failedMetricsCount"`

	// A list of records identifying metric objects that failed input validation
	// and the reasons for the failures.
	FailedMetrics []FailedMetricRecord `mandatory:"false" json:"failedMetrics"`
}

func (m PostMetricDataResponseDetails) String() string {
	return common.PointerString(m)
}
