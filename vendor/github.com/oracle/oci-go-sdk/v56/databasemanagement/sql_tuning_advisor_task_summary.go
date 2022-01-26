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

// SqlTuningAdvisorTaskSummary The summary for a SQL Tuning Advisor task.
type SqlTuningAdvisorTaskSummary struct {

	// Unique identifier of the task. It is not the OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	SqlTuningAdvisorTaskId *int64 `mandatory:"true" json:"sqlTuningAdvisorTaskId"`

	// The instance id of the task. It is not the OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	InstanceId *int `mandatory:"false" json:"instanceId"`

	// The name of the task.
	Name *string `mandatory:"false" json:"name"`

	// The description of the task.
	Description *string `mandatory:"false" json:"description"`

	// The owner of the task.
	Owner *string `mandatory:"false" json:"owner"`

	// Creation date of the task.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The status of the task.
	TaskStatus SqlTuningTaskStatusTypesEnum `mandatory:"false" json:"taskStatus,omitempty"`

	// Days to expire the task. If the value equals -1 then the task has no expiration time (UNLIMITED).
	DaysToExpire *int `mandatory:"false" json:"daysToExpire"`

	// Start timestamp of task execution.
	TimeExecutionStarted *common.SDKTime `mandatory:"false" json:"timeExecutionStarted"`

	// End timestamp of task execution.
	TimeExecutionEnded *common.SDKTime `mandatory:"false" json:"timeExecutionEnded"`

	// The total number of SQL statements related to the SQL tuning advisor task.
	TotalSqlStatements *int `mandatory:"false" json:"totalSqlStatements"`

	// Number of recommendations produced.
	RecommendationCount *int `mandatory:"false" json:"recommendationCount"`
}

func (m SqlTuningAdvisorTaskSummary) String() string {
	return common.PointerString(m)
}
