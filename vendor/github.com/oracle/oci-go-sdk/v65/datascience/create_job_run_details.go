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

// CreateJobRunDetails Parameters needed to create a new job run.
type CreateJobRunDetails struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project to associate the job with.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where you want to create the job.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the job to create a run for.
	JobId *string `mandatory:"true" json:"jobId"`

	// A user-friendly display name for the resource.
	DisplayName *string `mandatory:"false" json:"displayName"`

	JobConfigurationOverrideDetails JobConfigurationDetails `mandatory:"false" json:"jobConfigurationOverrideDetails"`

	JobLogConfigurationOverrideDetails *JobLogConfigurationDetails `mandatory:"false" json:"jobLogConfigurationOverrideDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateJobRunDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateJobRunDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateJobRunDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName                        *string                           `json:"displayName"`
		JobConfigurationOverrideDetails    jobconfigurationdetails           `json:"jobConfigurationOverrideDetails"`
		JobLogConfigurationOverrideDetails *JobLogConfigurationDetails       `json:"jobLogConfigurationOverrideDetails"`
		FreeformTags                       map[string]string                 `json:"freeformTags"`
		DefinedTags                        map[string]map[string]interface{} `json:"definedTags"`
		ProjectId                          *string                           `json:"projectId"`
		CompartmentId                      *string                           `json:"compartmentId"`
		JobId                              *string                           `json:"jobId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	nn, e = model.JobConfigurationOverrideDetails.UnmarshalPolymorphicJSON(model.JobConfigurationOverrideDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.JobConfigurationOverrideDetails = nn.(JobConfigurationDetails)
	} else {
		m.JobConfigurationOverrideDetails = nil
	}

	m.JobLogConfigurationOverrideDetails = model.JobLogConfigurationOverrideDetails

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.ProjectId = model.ProjectId

	m.CompartmentId = model.CompartmentId

	m.JobId = model.JobId

	return
}
