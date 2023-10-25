// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Monitoring API
//
// Use the Monitoring API to manage metric queries and alarms for assessing the health, capacity, and performance of your cloud resources.
// Endpoints vary by operation. For PostMetric, use the `telemetry-ingestion` endpoints; for all other operations, use the `telemetry` endpoints.
// For more information, see
// the Monitoring documentation (https://docs.cloud.oracle.com/iaas/Content/Monitoring/home.htm).
//

package monitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PostMetricDataResponseDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
