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

// SqlTuningAdvisorTaskSummaryReportTaskInfo SQL Tuning advisor task general info.
type SqlTuningAdvisorTaskSummaryReportTaskInfo struct {

	// The SQL Tuning Advisor task id. It is not the OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	Id *int64 `mandatory:"true" json:"id"`

	// The SQL Tuning Advisor task name.
	Name *string `mandatory:"true" json:"name"`

	// The SQL Tuning Advisor task user owner.
	Owner *string `mandatory:"true" json:"owner"`

	// Start timestamp of task execution.
	TimeStarted *common.SDKTime `mandatory:"true" json:"timeStarted"`

	// End timestamp of task execution.
	TimeEnded *common.SDKTime `mandatory:"true" json:"timeEnded"`

	// The SQL Tuning Advisor task description. Not defined on Auto SQL Tuning tasks.
	Description *string `mandatory:"false" json:"description"`

	// The SQL Tuning Advisor task status. Not defined on Auto SQL Tuning tasks.
	Status SqlTuningTaskStatusTypesEnum `mandatory:"false" json:"status,omitempty"`

	// The total running time in seconds. Not defined on Auto SQL Tuning tasks.
	RunningTime *int `mandatory:"false" json:"runningTime"`
}

func (m SqlTuningAdvisorTaskSummaryReportTaskInfo) String() string {
	return common.PointerString(m)
}
