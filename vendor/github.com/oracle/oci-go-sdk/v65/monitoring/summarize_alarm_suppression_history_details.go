// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Monitoring API
//
// Use the Monitoring API to manage metric queries and alarms for assessing the health, capacity, and performance of your cloud resources.
// Endpoints vary by operation. For PostMetricData, use the `telemetry-ingestion` endpoints; for all other operations, use the `telemetry` endpoints.
// For more information, see
// the Monitoring documentation (https://docs.cloud.oracle.com/iaas/Content/Monitoring/home.htm).
//

package monitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SummarizeAlarmSuppressionHistoryDetails The configuration details for returning history of suppressions for the specified alarm.
type SummarizeAlarmSuppressionHistoryDetails struct {

	// A filter to suppress only alarm state entries that include the set of specified dimension key-value pairs.
	// If you specify {"availabilityDomain": "phx-ad-1"}
	// and the alarm state entry corresponds to the set {"availabilityDomain": "phx-ad-1" and "resourceId": "ocid1.instance.region1.phx.exampleuniqueID"},
	// then this alarm will be included for suppression.
	// Example: `{"resourceId": "ocid1.instance.region1.phx.exampleuniqueID"}`
	Dimensions map[string]string `mandatory:"false" json:"dimensions"`

	// A filter to return only entries with "timeSuppressFrom" time occurring on or after the specified time.
	// The value cannot be a future time.
	// Format defined by RFC3339.
	// Example: `2023-02-01T01:02:29.600Z`
	TimeSuppressFromGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" json:"timeSuppressFromGreaterThanOrEqualTo"`

	// A filter to return only entries with "timeSuppressFrom" time occurring before the specified time.
	// The value cannot be a future time.
	// Format defined by RFC3339.
	// Example: `2023-02-01T01:02:29.600Z`
	TimeSuppressFromLessThan *common.SDKTime `mandatory:"false" json:"timeSuppressFromLessThan"`
}

func (m SummarizeAlarmSuppressionHistoryDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SummarizeAlarmSuppressionHistoryDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
