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

// FailedMetricRecord The record of a single metric object that failed input validation and the reason for the failure.
type FailedMetricRecord struct {

	// An error message indicating the reason that the indicated metric object failed input validation.
	Message *string `mandatory:"true" json:"message"`

	// Identifier of a metric object that failed input validation.
	MetricData *MetricDataDetails `mandatory:"true" json:"metricData"`
}

func (m FailedMetricRecord) String() string {
	return common.PointerString(m)
}
