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

// CreateJobExecutionDetails Properties for creating a new job execution.
type CreateJobExecutionDetails struct {

	// Sub-type of this job execution.
	SubType *string `mandatory:"false" json:"subType"`

	// Type of the job execution.
	JobType JobTypeEnum `mandatory:"false" json:"jobType,omitempty"`

	// The unique key of the parent execution or null if this job execution has no parent.
	ParentKey *string `mandatory:"false" json:"parentKey"`

	// Time that job execution started. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// Time that the job execution ended or null if it hasn't yet completed.
	// An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeEnded *common.SDKTime `mandatory:"false" json:"timeEnded"`

	// Status of the job execution, such as running, paused, or completed.
	LifecycleState JobExecutionStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Error code returned from the job execution or null if job is still running or didn't return an error.
	ErrorCode *string `mandatory:"false" json:"errorCode"`

	// Error message returned from the job execution or null if job is still running or didn't return an error.
	ErrorMessage *string `mandatory:"false" json:"errorMessage"`

	// The unique key of the triggering external scheduler resource or null if this job execution is not externally triggered.
	ScheduleInstanceKey *string `mandatory:"false" json:"scheduleInstanceKey"`

	// Process identifier related to the job execution if the job is an external job.
	ProcessKey *string `mandatory:"false" json:"processKey"`

	// If the job is an external process, then a URL of the job for accessing this resource and its status.
	ExternalUrl *string `mandatory:"false" json:"externalUrl"`

	// An identifier used for log message correlation.
	EventKey *string `mandatory:"false" json:"eventKey"`

	// The key of the associated data entity resource.
	DataEntityKey *string `mandatory:"false" json:"dataEntityKey"`

	// A map of maps that contains the execution context properties which are specific to a job execution. Each job
	// execution may define it's set of required and optional properties. The map keys are category names and the
	// values are maps of property name to property value. Every property is contained inside of a category. Most
	// job executions have required properties within the "default" category.
	// Example: `{"properties": { "default": { "host": "host1", "port": "1521", "database": "orcl"}}}`
	Properties map[string]map[string]string `mandatory:"false" json:"properties"`
}

func (m CreateJobExecutionDetails) String() string {
	return common.PointerString(m)
}
