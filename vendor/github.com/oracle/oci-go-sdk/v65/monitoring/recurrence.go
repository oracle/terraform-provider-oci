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
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Recurrence Alarm suppression recurring schedule. Only one recurrence condition is supported within the list of preconditions for a suppression (`suppressionConditions`).
type Recurrence struct {

	// Frequency and start time of the recurring suppression. The format follows
	// the iCalendar specification (RFC 5545, section 3.3.10) (https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.10).
	// Supported rule parts:
	// * `FREQ`: Frequency of the recurring suppression: `WEEKLY` or `DAILY` only.
	// * `BYDAY`: Comma separated days. Use with weekly suppressions only. Supported values: `MO`, `TU`, `WE`, `TH`, `FR`, `SA` ,`SU`.
	// * `BYHOUR`, `BYMINUTE`, `BYSECOND`: Start time in UTC, after `timeSuppressFrom` value. Default is 00:00:00 UTC after `timeSuppressFrom`.
	SuppressionRecurrence *string `mandatory:"true" json:"suppressionRecurrence"`

	// Duration of the recurring suppression. Specified as a string in ISO 8601 format. Minimum: `PT1M` (1 minute). Maximum: `PT24H` (24 hours).
	SuppressionDuration *string `mandatory:"true" json:"suppressionDuration"`
}

func (m Recurrence) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Recurrence) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m Recurrence) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeRecurrence Recurrence
	s := struct {
		DiscriminatorParam string `json:"conditionType"`
		MarshalTypeRecurrence
	}{
		"RECURRENCE",
		(MarshalTypeRecurrence)(m),
	}

	return json.Marshal(&s)
}
