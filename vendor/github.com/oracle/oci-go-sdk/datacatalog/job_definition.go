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

// JobDefinition Representation of a job definition resource. Job definitions define the harvest scope and includes the list
// of objects to be harvested along with a schedule. The list of objects is usually specified through a combination
// of object type, regular expressions, or specific names of objects and a sample size for the data harvested.
type JobDefinition struct {

	// Unique key of the job definition resource that is immutable.
	Key *string `mandatory:"true" json:"key"`

	// A user-friendly display name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The data catalog's OCID.
	CatalogId *string `mandatory:"false" json:"catalogId"`

	// Type of the job definition.
	JobType JobTypeEnum `mandatory:"false" json:"jobType,omitempty"`

	// Specifies if the job definition is incremental or full.
	IsIncremental *bool `mandatory:"false" json:"isIncremental"`

	// The key of the data asset for which the job is defined.
	DataAssetKey *string `mandatory:"false" json:"dataAssetKey"`

	// Detailed description of the job definition.
	Description *string `mandatory:"false" json:"description"`

	// The key of the default connection resource to be used for harvest, sampling, profiling jobs.
	// This may be overridden in each job instance.
	ConnectionKey *string `mandatory:"false" json:"connectionKey"`

	// Version of the job definition object. Used internally but can be visible to users.
	InternalVersion *string `mandatory:"false" json:"internalVersion"`

	// Lifecycle state of the job definition.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The date and time the job definition was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2019-03-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The last time that any change was made to the data asset. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// OCID of the user who created this job definition.
	CreatedById *string `mandatory:"false" json:"createdById"`

	// OCID of the user who updated this job definition.
	UpdatedById *string `mandatory:"false" json:"updatedById"`

	// URI to the job definition instance in the API.
	Uri *string `mandatory:"false" json:"uri"`

	// Specify if sample data to be extracted as part of this harvest.
	IsSampleDataExtracted *bool `mandatory:"false" json:"isSampleDataExtracted"`

	// Specify the sample data size in MB, specified as number of rows, for this metadata harvest.
	SampleDataSizeInMBs *int `mandatory:"false" json:"sampleDataSizeInMBs"`

	// Time that the latest job execution started. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeLatestExecutionStarted *common.SDKTime `mandatory:"false" json:"timeLatestExecutionStarted"`

	// Time that the latest job execution ended or null if it hasn't yet completed.
	// An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeLatestExecutionEnded *common.SDKTime `mandatory:"false" json:"timeLatestExecutionEnded"`

	// Status of the latest job execution, such as running, paused, or completed.
	JobExecutionState JobExecutionStateEnum `mandatory:"false" json:"jobExecutionState,omitempty"`

	// Type of job schedule for the latest job executed.
	ScheduleType JobScheduleTypeEnum `mandatory:"false" json:"scheduleType,omitempty"`

	// A map of maps that contains the properties which are specific to the job type. Each job type
	// definition may define it's set of required and optional properties. The map keys are category names and the
	// values are maps of property name to property value. Every property is contained inside of a category. Most
	// job definitions have required properties within the "default" category.
	// Example: `{"properties": { "default": { "host": "host1", "port": "1521", "database": "orcl"}}}`
	Properties map[string]map[string]string `mandatory:"false" json:"properties"`
}

func (m JobDefinition) String() string {
	return common.PointerString(m)
}
