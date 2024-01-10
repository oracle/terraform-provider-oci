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

// SqlTuningAdvisorTaskFindingSummary A summary of the findings of the objects in a tuning task that match a given filter.
// This includes the kind of findings that were reported, whether the benefits were analyzed, and the number of benefits obtained.
type SqlTuningAdvisorTaskFindingSummary struct {

	// The unique identifier of the SQL Tuning Advisor task. This is not the OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	SqlTuningAdvisorTaskId *int64 `mandatory:"true" json:"sqlTuningAdvisorTaskId"`

	// The key of the object to which these recommendations apply.
	// This is not the OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	SqlTuningAdvisorTaskObjectId *int64 `mandatory:"true" json:"sqlTuningAdvisorTaskObjectId"`

	// The execution id of the analyzed SQL object. This is not the OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	SqlTuningAdvisorTaskObjectExecutionId *int64 `mandatory:"true" json:"sqlTuningAdvisorTaskObjectExecutionId"`

	// The text of the SQL statement.
	SqlText *string `mandatory:"true" json:"sqlText"`

	// The parsing schema of the object.
	ParsingSchema *string `mandatory:"true" json:"parsingSchema"`

	// The unique key of this SQL statement.
	SqlKey *string `mandatory:"true" json:"sqlKey"`

	// The time benefit (in seconds) for the highest-rated finding for this object.
	DbTimeBenefit *float32 `mandatory:"false" json:"dbTimeBenefit"`

	// The per-execution percentage benefit.
	PerExecutionPercentage *int `mandatory:"false" json:"perExecutionPercentage"`

	// Indicates whether a statistics recommendation was reported for this SQL statement.
	IsStatsFindingPresent *bool `mandatory:"false" json:"isStatsFindingPresent"`

	// Indicates whether a SQL Profile recommendation was reported for this SQL statement.
	IsSqlProfileFindingPresent *bool `mandatory:"false" json:"isSqlProfileFindingPresent"`

	// Indicates whether a SQL Profile recommendation has been implemented for this SQL statement.
	IsSqlProfileFindingImplemented *bool `mandatory:"false" json:"isSqlProfileFindingImplemented"`

	// Indicates whether an index recommendation was reported for this SQL statement.
	IsIndexFindingPresent *bool `mandatory:"false" json:"isIndexFindingPresent"`

	// Indicates whether a restructure SQL recommendation was reported for this SQL statement.
	IsRestructureSqlFindingPresent *bool `mandatory:"false" json:"isRestructureSqlFindingPresent"`

	// Indicates whether an alternative execution plan was reported for this SQL statement.
	IsAlternativePlanFindingPresent *bool `mandatory:"false" json:"isAlternativePlanFindingPresent"`

	// Indicates whether a miscellaneous finding was reported for this SQL statement.
	IsMiscellaneousFindingPresent *bool `mandatory:"false" json:"isMiscellaneousFindingPresent"`

	// Indicates whether there is an error in this SQL statement.
	IsErrorFindingPresent *bool `mandatory:"false" json:"isErrorFindingPresent"`

	// Indicates whether the task timed out.
	IsTimeoutFindingPresent *bool `mandatory:"false" json:"isTimeoutFindingPresent"`
}

func (m SqlTuningAdvisorTaskFindingSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SqlTuningAdvisorTaskFindingSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
