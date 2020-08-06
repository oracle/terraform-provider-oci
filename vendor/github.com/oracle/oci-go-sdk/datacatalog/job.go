// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
//

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Job Details of a job. Jobs are scheduled instances of a job definition.
type Job struct {

	// Unique key of the job resource.
	Key *string `mandatory:"true" json:"key"`

	// A user-friendly display name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Detailed description of the job.
	Description *string `mandatory:"false" json:"description"`

	// The data catalog's OCID.
	CatalogId *string `mandatory:"false" json:"catalogId"`

	// Lifecycle state for job.
	LifecycleState JobLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The date and time the job was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2019-03-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Time that this job was last updated. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Type of the job.
	JobType JobTypeEnum `mandatory:"false" json:"jobType,omitempty"`

	// Schedule specified in the cron expression format that has seven fields for second, minute, hour, day-of-month, month, day-of-week, year.
	// It can also include special characters like * for all and ? for any. There are also pre-defined schedules that can be specified using
	// special strings. For example, @hourly will run the job every hour.
	ScheduleCronExpression *string `mandatory:"false" json:"scheduleCronExpression"`

	// Date that the schedule should be operational. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeScheduleBegin *common.SDKTime `mandatory:"false" json:"timeScheduleBegin"`

	// Date that the schedule should end from being operational. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeScheduleEnd *common.SDKTime `mandatory:"false" json:"timeScheduleEnd"`

	// Type of job schedule that is inferred from the scheduling properties.
	ScheduleType JobScheduleTypeEnum `mandatory:"false" json:"scheduleType,omitempty"`

	// The key of the connection used by the job. This connection will override the default connection specified in
	// the associated job definition. All executions will use this connection.
	ConnectionKey *string `mandatory:"false" json:"connectionKey"`

	// The unique key of the job definition resource that defined the scope of this job.
	JobDefinitionKey *string `mandatory:"false" json:"jobDefinitionKey"`

	// Internal version of the job resource.
	InternalVersion *string `mandatory:"false" json:"internalVersion"`

	// The total number of executions for this job schedule.
	ExecutionCount *int `mandatory:"false" json:"executionCount"`

	// The date and time of the most recent execution for this Job, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2019-03-25T21:10:29.600Z`
	TimeOfLatestExecution *common.SDKTime `mandatory:"false" json:"timeOfLatestExecution"`

	// OCID of the user who created this job.
	CreatedById *string `mandatory:"false" json:"createdById"`

	// OCID of the user who updated this job.
	UpdatedById *string `mandatory:"false" json:"updatedById"`

	// The display name of the job definition resource that defined the scope of this job.
	JobDefinitionName *string `mandatory:"false" json:"jobDefinitionName"`

	// Error code returned from the latest job execution for this job. Useful when the latest Job execution is in FAILED state.
	ErrorCode *string `mandatory:"false" json:"errorCode"`

	// Error message returned from the latest job execution for this job. Useful when the latest Job Execution is in a FAILED state.
	ErrorMessage *string `mandatory:"false" json:"errorMessage"`

	// URI to the job instance in the API.
	Uri *string `mandatory:"false" json:"uri"`
}

func (m Job) String() string {
	return common.PointerString(m)
}
