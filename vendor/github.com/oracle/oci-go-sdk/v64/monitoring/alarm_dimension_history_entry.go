// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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
	"fmt"
	"github.com/oracle/oci-go-sdk/v64/common"
	"strings"
)

// AlarmDimensionHistoryEntry A timestamped alarm history entry including a description (`summary`).
// If the entry corresponds to a state transition, such as OK to Firing, then the entry also includes a transition timestamp.
// The entry applies to one or more of the dimension key-value pairs specified in the request (from `RetrieveDimensionHistoryDetails`).
type AlarmDimensionHistoryEntry struct {

	// Description of the alarm history entry.
	// Example 1 - alarm status is FIRING:
	//   `The alarm stream state is FIRING`
	// Example 2 - the alarm transitioned to a different status:
	//   `State transitioned from OK to Firing`
	Summary *string `mandatory:"true" json:"summary"`

	// Timestamp for this alarm history entry. Format defined by RFC3339.
	// Example: `2022-02-01T01:02:29.600Z`
	Timestamp *common.SDKTime `mandatory:"true" json:"timestamp"`

	// Timestamp for the transition of the alarm state. For example, the time when the alarm transitioned from OK to Firing.
	// Available for state transition entries only. Note: A three-minute lag for this value accounts for any late-arriving metrics.
	// Example: `2022-02-01T0:59:00.789Z`
	TimestampTriggered *common.SDKTime `mandatory:"false" json:"timestampTriggered"`
}

func (m AlarmDimensionHistoryEntry) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AlarmDimensionHistoryEntry) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
