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

// JobDefinitionSummary A list of job definition resources. Job definitions define the harvest scope and includes the list of objects
// to be harvested along with a schedule. The list of objects is usually specified through a combination of object
// type, regular expressions, or specific names of objects and a sample size for the data harvested.
type JobDefinitionSummary struct {

	// Unique key of the job definition resource that is immutable.
	Key *string `mandatory:"true" json:"key"`

	// A user-friendly display name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Detailed description of the job definition.
	Description *string `mandatory:"false" json:"description"`

	// The data catalog's OCID.
	CatalogId *string `mandatory:"false" json:"catalogId"`

	// URI to the job definition instance in the API.
	Uri *string `mandatory:"false" json:"uri"`

	// Type of the job definition.
	JobType JobTypeEnum `mandatory:"false" json:"jobType,omitempty"`

	// Lifecycle state of the job definition.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Specify if sample data to be extracted as part of this harvest.
	IsSampleDataExtracted *bool `mandatory:"false" json:"isSampleDataExtracted"`

	// The date and time the job definition was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2019-03-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The key of the connection resource used in harvest, sampling, profiling jobs.
	ConnectionKey *string `mandatory:"false" json:"connectionKey"`

	// Time that the latest job execution started. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeLatestExecutionStarted *common.SDKTime `mandatory:"false" json:"timeLatestExecutionStarted"`

	// Time that the latest job execution ended or null if it hasn't yet completed.
	// An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeLatestExecutionEnded *common.SDKTime `mandatory:"false" json:"timeLatestExecutionEnded"`

	// Status of the latest job execution, such as running, paused, or completed.
	JobExecutionState JobExecutionStateEnum `mandatory:"false" json:"jobExecutionState,omitempty"`

	// Type of job schedule for the latest job executed.
	ScheduleType JobScheduleTypeEnum `mandatory:"false" json:"scheduleType,omitempty"`

	// Unique key of the data asset to which this job applies, if the job involves a data asset.
	DataAssetKey *string `mandatory:"false" json:"dataAssetKey"`

	// Unique key of the glossary to which this job applies, if the job involves a glossary.
	GlossaryKey *string `mandatory:"false" json:"glossaryKey"`
}

func (m JobDefinitionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m JobDefinitionSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingJobTypeEnum(string(m.JobType)); !ok && m.JobType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for JobType: %s. Supported values are: %s.", m.JobType, strings.Join(GetJobTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingJobExecutionStateEnum(string(m.JobExecutionState)); !ok && m.JobExecutionState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for JobExecutionState: %s. Supported values are: %s.", m.JobExecutionState, strings.Join(GetJobExecutionStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingJobScheduleTypeEnum(string(m.ScheduleType)); !ok && m.ScheduleType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ScheduleType: %s. Supported values are: %s.", m.ScheduleType, strings.Join(GetJobScheduleTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
