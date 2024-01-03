// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ScheduleAuditReportDetails Details for the audit report schedule.
type ScheduleAuditReportDetails struct {

	// The time span of records in report to be scheduled.
	// <period-value><period>
	// Allowed period strings - "H","D","M","Y"
	// Each of the above fields potentially introduce constraints. A workRequest is created only
	// when period-value satisfies all the constraints. Constraints introduced:
	// 1. period = H (The allowed range for period-value is [1, 23])
	// 2. period = D (The allowed range for period-value is [1, 30])
	// 3. period = M (The allowed range for period-value is [1, 11])
	// 4. period = Y (The minimum period-value is 1)
	RecordTimeSpan *string `mandatory:"true" json:"recordTimeSpan"`

	// Specifies the limit on the number of rows in the report.
	RowLimit *int `mandatory:"false" json:"rowLimit"`
}

func (m ScheduleAuditReportDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ScheduleAuditReportDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ScheduleAuditReportDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeScheduleAuditReportDetails ScheduleAuditReportDetails
	s := struct {
		DiscriminatorParam string `json:"reportType"`
		MarshalTypeScheduleAuditReportDetails
	}{
		"AUDIT",
		(MarshalTypeScheduleAuditReportDetails)(m),
	}

	return json.Marshal(&s)
}
