// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// SqlTuningAdvisorTaskSummaryReportStatementCounts The number of statements in the SQL Tuning Advisor summary report.
type SqlTuningAdvisorTaskSummaryReportStatementCounts struct {

	// The number of distinct SQL statements.
	DistinctSql *int `mandatory:"true" json:"distinctSql"`

	// The total number of SQL statements.
	TotalSql *int `mandatory:"true" json:"totalSql"`

	// The number of distinct SQL statements with findings.
	FindingCount *int `mandatory:"true" json:"findingCount"`

	// The number of distinct SQL statements with errors.
	ErrorCount *int `mandatory:"true" json:"errorCount"`
}

func (m SqlTuningAdvisorTaskSummaryReportStatementCounts) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SqlTuningAdvisorTaskSummaryReportStatementCounts) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
