// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SqlTuningAdvisorTaskSummaryReportIndexFindingSummary A summary for all the index findings in a SQL Tuning Advisor task. Includes the index's hash value, table name, schema, index name, reference count and index columns
type SqlTuningAdvisorTaskSummaryReportIndexFindingSummary struct {

	// Numerical representation of the index.
	IndexHashValue *int64 `mandatory:"true" json:"indexHashValue"`

	// Name of the index.
	IndexName *string `mandatory:"true" json:"indexName"`

	// Table's name related to the index.
	TableName *string `mandatory:"true" json:"tableName"`

	// Schema related to the index.
	Schema *string `mandatory:"true" json:"schema"`

	// The number of times the index is referenced within the SQL Tuning advisor task findings.
	ReferenceCount *int `mandatory:"true" json:"referenceCount"`

	// Columns of the index.
	IndexColumns []string `mandatory:"true" json:"indexColumns"`
}

func (m SqlTuningAdvisorTaskSummaryReportIndexFindingSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SqlTuningAdvisorTaskSummaryReportIndexFindingSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
