// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
// For more information, see Data Catalog (https://docs.oracle.com/iaas/data-catalog/home.htm).
//

package datacatalog

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
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

	// Interval on which the job will be run. Value is specified as a cron-supported time specification "nickname".
	// The following subset of those is supported: @monthly, @weekly, @daily, @hourly.
	// For metastore sync, an additional option @default is supported, which will schedule jobs at a more granular frequency.
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

	// Unique key of the data asset to which this job applies, if the job involves a data asset.
	DataAssetKey *string `mandatory:"false" json:"dataAssetKey"`

	// Unique key of the glossary to which this job applies.
	GlossaryKey *string `mandatory:"false" json:"glossaryKey"`

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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Job) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingJobLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetJobLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingJobTypeEnum(string(m.JobType)); !ok && m.JobType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for JobType: %s. Supported values are: %s.", m.JobType, strings.Join(GetJobTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingJobScheduleTypeEnum(string(m.ScheduleType)); !ok && m.ScheduleType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ScheduleType: %s. Supported values are: %s.", m.ScheduleType, strings.Join(GetJobScheduleTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
