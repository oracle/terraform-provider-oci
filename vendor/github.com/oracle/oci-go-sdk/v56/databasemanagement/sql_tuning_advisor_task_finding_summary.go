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
	"github.com/oracle/oci-go-sdk/v56/common"
)

// SqlTuningAdvisorTaskFindingSummary A summary for all the findings of objects in a tuning task that match a given certain filter.
// Includes what kind of findings were found, whether benefits were analyzed, and how many benefits can be obtained.
type SqlTuningAdvisorTaskFindingSummary struct {

	// Unique identifier of the task. It is not the OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	SqlTuningAdvisorTaskId *int64 `mandatory:"true" json:"sqlTuningAdvisorTaskId"`

	// Key of the object to which these recommendations apply.
	// It is not the OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	SqlTuningAdvisorTaskObjectId *int64 `mandatory:"true" json:"sqlTuningAdvisorTaskObjectId"`

	// Execution id of the analyzed SQL object. It is not the OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	SqlTuningAdvisorTaskObjectExecutionId *int64 `mandatory:"true" json:"sqlTuningAdvisorTaskObjectExecutionId"`

	// Text of the SQL statement.
	SqlText *string `mandatory:"true" json:"sqlText"`

	// Parsing schema of the object.
	ParsingSchema *string `mandatory:"true" json:"parsingSchema"`

	// Unique key of this SQL statement
	SqlKey *string `mandatory:"true" json:"sqlKey"`

	// Time benefit in seconds for the highest-rated finding for this object.
	DbTimeBenefit *float32 `mandatory:"false" json:"dbTimeBenefit"`

	// The per-execution percentage benefit.
	PerExecutionPercentage *int `mandatory:"false" json:"perExecutionPercentage"`

	// Whether a statistics recommendation was found for this SQL statement.
	IsStatsFindingPresent *bool `mandatory:"false" json:"isStatsFindingPresent"`

	// Whether a SQL Profile recommendation was found for this SQL statement.
	IsSqlProfileFindingPresent *bool `mandatory:"false" json:"isSqlProfileFindingPresent"`

	// Whether a SQL Profile recommendation has been implemented for this SQL statement.
	IsSqlProfileFindingImplemented *bool `mandatory:"false" json:"isSqlProfileFindingImplemented"`

	// Whether an index recommendation was found for this SQL statement.
	IsIndexFindingPresent *bool `mandatory:"false" json:"isIndexFindingPresent"`

	// Whether a restructure SQL recommendation was found for this SQL statement.
	IsRestructureSqlFindingPresent *bool `mandatory:"false" json:"isRestructureSqlFindingPresent"`

	// Whether an alternative execution plan was found for this SQL statement.
	IsAlternativePlanFindingPresent *bool `mandatory:"false" json:"isAlternativePlanFindingPresent"`

	// Whether a miscellaneous finding was found for this SQL statement.
	IsMiscellaneousFindingPresent *bool `mandatory:"false" json:"isMiscellaneousFindingPresent"`

	// Whether there is an error in this SQL statement.
	IsErrorFindingPresent *bool `mandatory:"false" json:"isErrorFindingPresent"`

	// Whether the task timed out.
	IsTimeoutFindingPresent *bool `mandatory:"false" json:"isTimeoutFindingPresent"`
}

func (m SqlTuningAdvisorTaskFindingSummary) String() string {
	return common.PointerString(m)
}
