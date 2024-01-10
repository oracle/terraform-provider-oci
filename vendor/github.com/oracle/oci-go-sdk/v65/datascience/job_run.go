// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// JobRun A job run.
type JobRun struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the job run.
	Id *string `mandatory:"true" json:"id"`

	// The date and time the job run was accepted in the timestamp format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeAccepted *common.SDKTime `mandatory:"true" json:"timeAccepted"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user who created the job run.
	CreatedBy *string `mandatory:"true" json:"createdBy"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project to associate the job with.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where you want to create the job.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the job run.
	JobId *string `mandatory:"true" json:"jobId"`

	JobConfigurationOverrideDetails JobConfigurationDetails `mandatory:"true" json:"jobConfigurationOverrideDetails"`

	JobInfrastructureConfigurationDetails JobInfrastructureConfigurationDetails `mandatory:"true" json:"jobInfrastructureConfigurationDetails"`

	// The state of the job run.
	LifecycleState JobRunLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the job run request was started in the timestamp format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The date and time the job run request was finished in the timestamp format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`

	// A user-friendly display name for the resource.
	DisplayName *string `mandatory:"false" json:"displayName"`

	JobLogConfigurationOverrideDetails *JobLogConfigurationDetails `mandatory:"false" json:"jobLogConfigurationOverrideDetails"`

	// Collection of JobStorageMountConfigurationDetails.
	JobStorageMountConfigurationDetailsList []StorageMountConfigurationDetails `mandatory:"false" json:"jobStorageMountConfigurationDetailsList"`

	LogDetails *JobRunLogDetails `mandatory:"false" json:"logDetails"`

	// Details of the state of the job run.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m JobRun) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m JobRun) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingJobRunLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetJobRunLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *JobRun) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		TimeStarted                             *common.SDKTime                       `json:"timeStarted"`
		TimeFinished                            *common.SDKTime                       `json:"timeFinished"`
		DisplayName                             *string                               `json:"displayName"`
		JobLogConfigurationOverrideDetails      *JobLogConfigurationDetails           `json:"jobLogConfigurationOverrideDetails"`
		JobStorageMountConfigurationDetailsList []storagemountconfigurationdetails    `json:"jobStorageMountConfigurationDetailsList"`
		LogDetails                              *JobRunLogDetails                     `json:"logDetails"`
		LifecycleDetails                        *string                               `json:"lifecycleDetails"`
		FreeformTags                            map[string]string                     `json:"freeformTags"`
		DefinedTags                             map[string]map[string]interface{}     `json:"definedTags"`
		Id                                      *string                               `json:"id"`
		TimeAccepted                            *common.SDKTime                       `json:"timeAccepted"`
		CreatedBy                               *string                               `json:"createdBy"`
		ProjectId                               *string                               `json:"projectId"`
		CompartmentId                           *string                               `json:"compartmentId"`
		JobId                                   *string                               `json:"jobId"`
		JobConfigurationOverrideDetails         jobconfigurationdetails               `json:"jobConfigurationOverrideDetails"`
		JobInfrastructureConfigurationDetails   jobinfrastructureconfigurationdetails `json:"jobInfrastructureConfigurationDetails"`
		LifecycleState                          JobRunLifecycleStateEnum              `json:"lifecycleState"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.TimeStarted = model.TimeStarted

	m.TimeFinished = model.TimeFinished

	m.DisplayName = model.DisplayName

	m.JobLogConfigurationOverrideDetails = model.JobLogConfigurationOverrideDetails

	m.JobStorageMountConfigurationDetailsList = make([]StorageMountConfigurationDetails, len(model.JobStorageMountConfigurationDetailsList))
	for i, n := range model.JobStorageMountConfigurationDetailsList {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.JobStorageMountConfigurationDetailsList[i] = nn.(StorageMountConfigurationDetails)
		} else {
			m.JobStorageMountConfigurationDetailsList[i] = nil
		}
	}
	m.LogDetails = model.LogDetails

	m.LifecycleDetails = model.LifecycleDetails

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.Id = model.Id

	m.TimeAccepted = model.TimeAccepted

	m.CreatedBy = model.CreatedBy

	m.ProjectId = model.ProjectId

	m.CompartmentId = model.CompartmentId

	m.JobId = model.JobId

	nn, e = model.JobConfigurationOverrideDetails.UnmarshalPolymorphicJSON(model.JobConfigurationOverrideDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.JobConfigurationOverrideDetails = nn.(JobConfigurationDetails)
	} else {
		m.JobConfigurationOverrideDetails = nil
	}

	nn, e = model.JobInfrastructureConfigurationDetails.UnmarshalPolymorphicJSON(model.JobInfrastructureConfigurationDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.JobInfrastructureConfigurationDetails = nn.(JobInfrastructureConfigurationDetails)
	} else {
		m.JobInfrastructureConfigurationDetails = nil
	}

	m.LifecycleState = model.LifecycleState

	return
}
