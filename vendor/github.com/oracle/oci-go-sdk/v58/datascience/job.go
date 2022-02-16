// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// Job A job for training models.
type Job struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the job.
	Id *string `mandatory:"true" json:"id"`

	// The date and time the resource was created in the timestamp format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: 2020-08-06T21:10:29.41Z
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user who created the job.
	CreatedBy *string `mandatory:"true" json:"createdBy"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project to associate the job with.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where you want to create the job.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	JobConfigurationDetails JobConfigurationDetails `mandatory:"true" json:"jobConfigurationDetails"`

	JobInfrastructureConfigurationDetails JobInfrastructureConfigurationDetails `mandatory:"true" json:"jobInfrastructureConfigurationDetails"`

	// The state of the job.
	LifecycleState JobLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// A user-friendly display name for the resource.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A short description of the job.
	Description *string `mandatory:"false" json:"description"`

	JobLogConfigurationDetails *JobLogConfigurationDetails `mandatory:"false" json:"jobLogConfigurationDetails"`

	// The state of the job.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
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

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *Job) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName                           *string                               `json:"displayName"`
		Description                           *string                               `json:"description"`
		JobLogConfigurationDetails            *JobLogConfigurationDetails           `json:"jobLogConfigurationDetails"`
		LifecycleDetails                      *string                               `json:"lifecycleDetails"`
		FreeformTags                          map[string]string                     `json:"freeformTags"`
		DefinedTags                           map[string]map[string]interface{}     `json:"definedTags"`
		Id                                    *string                               `json:"id"`
		TimeCreated                           *common.SDKTime                       `json:"timeCreated"`
		CreatedBy                             *string                               `json:"createdBy"`
		ProjectId                             *string                               `json:"projectId"`
		CompartmentId                         *string                               `json:"compartmentId"`
		JobConfigurationDetails               jobconfigurationdetails               `json:"jobConfigurationDetails"`
		JobInfrastructureConfigurationDetails jobinfrastructureconfigurationdetails `json:"jobInfrastructureConfigurationDetails"`
		LifecycleState                        JobLifecycleStateEnum                 `json:"lifecycleState"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	m.JobLogConfigurationDetails = model.JobLogConfigurationDetails

	m.LifecycleDetails = model.LifecycleDetails

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.Id = model.Id

	m.TimeCreated = model.TimeCreated

	m.CreatedBy = model.CreatedBy

	m.ProjectId = model.ProjectId

	m.CompartmentId = model.CompartmentId

	nn, e = model.JobConfigurationDetails.UnmarshalPolymorphicJSON(model.JobConfigurationDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.JobConfigurationDetails = nn.(JobConfigurationDetails)
	} else {
		m.JobConfigurationDetails = nil
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
