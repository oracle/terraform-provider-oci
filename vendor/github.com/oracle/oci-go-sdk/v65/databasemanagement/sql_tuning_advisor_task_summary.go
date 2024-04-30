// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.cloud.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SqlTuningAdvisorTaskSummary The summary of a SQL Tuning Advisor task.
type SqlTuningAdvisorTaskSummary struct {

	// The unique identifier of the SQL Tuning Advisor task. This is not the OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	SqlTuningAdvisorTaskId *int64 `mandatory:"true" json:"sqlTuningAdvisorTaskId"`

	// The instance ID of the SQL Tuning Advisor task. This is not the OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	InstanceId *int `mandatory:"false" json:"instanceId"`

	// The name of the SQL Tuning Advisor task.
	Name *string `mandatory:"false" json:"name"`

	// The description of the SQL Tuning Advisor task.
	Description *string `mandatory:"false" json:"description"`

	// The owner of the SQL Tuning Advisor task.
	Owner *string `mandatory:"false" json:"owner"`

	// The Creation date of the SQL Tuning Advisor task.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The status of the SQL Tuning Advisor task.
	TaskStatus SqlTuningTaskStatusTypesEnum `mandatory:"false" json:"taskStatus,omitempty"`

	// The number of days left before the task expires. If the value equals -1, then the task has no expiration time (UNLIMITED).
	DaysToExpire *int `mandatory:"false" json:"daysToExpire"`

	// The start time of the task execution.
	TimeExecutionStarted *common.SDKTime `mandatory:"false" json:"timeExecutionStarted"`

	// The end time of the task execution.
	TimeExecutionEnded *common.SDKTime `mandatory:"false" json:"timeExecutionEnded"`

	// The total number of SQL statements related to the SQL Tuning Advisor task.
	TotalSqlStatements *int `mandatory:"false" json:"totalSqlStatements"`

	// The number of recommendations provided for the SQL Tuning Advisor task.
	RecommendationCount *int `mandatory:"false" json:"recommendationCount"`
}

func (m SqlTuningAdvisorTaskSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SqlTuningAdvisorTaskSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSqlTuningTaskStatusTypesEnum(string(m.TaskStatus)); !ok && m.TaskStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TaskStatus: %s. Supported values are: %s.", m.TaskStatus, strings.Join(GetSqlTuningTaskStatusTypesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
