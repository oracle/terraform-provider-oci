// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// JobExecutionSummary A list of job executions. A job execution is a unit of work being executed on behalf of a job.
type JobExecutionSummary struct {

	// Unique key of the job execution resource.
	Key *string `mandatory:"true" json:"key"`

	// The unique key of the parent job.
	JobKey *string `mandatory:"false" json:"jobKey"`

	// Type of the job execution.
	JobType JobTypeEnum `mandatory:"false" json:"jobType,omitempty"`

	// The unique key of the parent execution or null if this job execution has no parent.
	ParentKey *string `mandatory:"false" json:"parentKey"`

	// The unique key of the triggering external scheduler resource or null if this job execution is not externally triggered.
	ScheduleInstanceKey *string `mandatory:"false" json:"scheduleInstanceKey"`

	// Status of the job execution, such as running, paused, or completed.
	LifecycleState JobExecutionStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The date and time the job execution was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2019-03-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Time that job execution started. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// Time that the job execution ended or null if it hasn't yet completed.
	// An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeEnded *common.SDKTime `mandatory:"false" json:"timeEnded"`

	// URI to the job execution instance in the API.
	Uri *string `mandatory:"false" json:"uri"`
}

func (m JobExecutionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m JobExecutionSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingJobTypeEnum(string(m.JobType)); !ok && m.JobType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for JobType: %s. Supported values are: %s.", m.JobType, strings.Join(GetJobTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingJobExecutionStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetJobExecutionStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
