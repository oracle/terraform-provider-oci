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

// Recurrence Alarm suppression recurring schedule. Only one recurrence condition is supported for suppression.
type Recurrence struct {

	// This specify frequency and time of recurring suppression. suppressionRecurrence is following standard RFC 5545 section.
	// Please refer this for details:https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.10
	// FREQ = frequency of recurring suppression. Only WEEKLY and DAILY supported.
	// BYDAY = Comma separated days for Weekly suppression.
	// BYHOUR, BYMINUTE, BYSECOND = These specify recurring suppression start time in UTC after `timeSuppressFrom` value. Defaults to "00:00:00" if not specified.
	// Other Rules are not supported.
	// Example:
	// To create recurring suppression on every Monday and Tuesday at 10:00:00 UTC
	// FREQ=WEEKLY;BYDAY=MO,TU;BYHOUR=10
	// To create recurring suppression every day at 21:30 UTC
	// FREQ=DAILY;BYHOUR=21;BYMINUTE=30
	SuppressionRecurrence *string `mandatory:"true" json:"suppressionRecurrence"`

	// Recurring suppression duration. example: PT10M, PT1H
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
