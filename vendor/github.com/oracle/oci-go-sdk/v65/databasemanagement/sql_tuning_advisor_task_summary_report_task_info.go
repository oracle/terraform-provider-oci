// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SqlTuningAdvisorTaskSummaryReportTaskInfo The general information regarding the SQL Tuning Advisor task.
type SqlTuningAdvisorTaskSummaryReportTaskInfo struct {

	// The ID of the SQL Tuning Advisor task. This is not the OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	Id *int64 `mandatory:"true" json:"id"`

	// The name of the SQL Tuning Advisor task.
	Name *string `mandatory:"true" json:"name"`

	// The owner of the SQL Tuning Advisor task.
	Owner *string `mandatory:"true" json:"owner"`

	// The start time of the task execution.
	TimeStarted *common.SDKTime `mandatory:"true" json:"timeStarted"`

	// The end time of the task execution.
	TimeEnded *common.SDKTime `mandatory:"true" json:"timeEnded"`

	// The description of the SQL Tuning Advisor task. This is not defined for Auto SQL Tuning tasks.
	Description *string `mandatory:"false" json:"description"`

	// The status of the SQL Tuning Advisor task. This is not defined for Auto SQL Tuning tasks.
	Status SqlTuningTaskStatusTypesEnum `mandatory:"false" json:"status,omitempty"`

	// The total running time in seconds. This is not defined for Auto SQL Tuning tasks.
	RunningTime *int `mandatory:"false" json:"runningTime"`
}

func (m SqlTuningAdvisorTaskSummaryReportTaskInfo) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SqlTuningAdvisorTaskSummaryReportTaskInfo) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSqlTuningTaskStatusTypesEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetSqlTuningTaskStatusTypesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
