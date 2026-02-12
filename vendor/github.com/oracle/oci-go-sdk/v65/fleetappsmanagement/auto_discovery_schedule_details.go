// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AutoDiscoveryScheduleDetails The schedule details for auto-discovery, if enabled.
// The auto-discovery time for every occurrence will adhere to the schedule start time.
// Example: To run the auto-discovery every day at 9:00AM UTC, the `recurrences`
// will be `FREQ=DAILY;INTERVAL=1` and `timeScheduleStart` could be `2016-08-25T09:00:00Z`.
type AutoDiscoveryScheduleDetails struct {

	// The date and time the auto-discovery starts, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	TimeScheduleStart *common.SDKTime `mandatory:"true" json:"timeScheduleStart"`

	// Recurrence rule for the frequency for auto-discovery, in the format defined by RFC 5545 section 3.3.10 (https://tools.ietf.org/html/rfc5545#section-3.3.10).
	// Support is only for recurrence pattern as specified in the examples below
	// Example:
	// - `FREQ=DAILY;INTERVAL=1` (every day)
	// - `FREQ=WEEKLY;BYDAY=MO,WE,FR` (every Monday, Wednesday, and Friday)
	//
	//
	//
	//
	Recurrences *string `mandatory:"true" json:"recurrences"`
}

func (m AutoDiscoveryScheduleDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutoDiscoveryScheduleDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
