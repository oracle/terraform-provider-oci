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

// CreatePipelineRunDetails The information about new PipelineRun.
type CreatePipelineRunDetails struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where you want to create the pipeline run.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the pipeline for which pipeline run is created.
	PipelineId *string `mandatory:"true" json:"pipelineId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project to associate the pipeline run with.
	ProjectId *string `mandatory:"false" json:"projectId"`

	// A user-friendly display name for the resource.
	DisplayName *string `mandatory:"false" json:"displayName"`

	ConfigurationOverrideDetails PipelineConfigurationDetails `mandatory:"false" json:"configurationOverrideDetails"`

	LogConfigurationOverrideDetails *PipelineLogConfigurationDetails `mandatory:"false" json:"logConfigurationOverrideDetails"`

	// Array of step override details. Only Step Configuration is allowed to be overridden.
	StepOverrideDetails []PipelineStepOverrideDetails `mandatory:"false" json:"stepOverrideDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m CreatePipelineRunDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreatePipelineRunDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreatePipelineRunDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ProjectId                       *string                           `json:"projectId"`
		DisplayName                     *string                           `json:"displayName"`
		ConfigurationOverrideDetails    pipelineconfigurationdetails      `json:"configurationOverrideDetails"`
		LogConfigurationOverrideDetails *PipelineLogConfigurationDetails  `json:"logConfigurationOverrideDetails"`
		StepOverrideDetails             []PipelineStepOverrideDetails     `json:"stepOverrideDetails"`
		FreeformTags                    map[string]string                 `json:"freeformTags"`
		DefinedTags                     map[string]map[string]interface{} `json:"definedTags"`
		SystemTags                      map[string]map[string]interface{} `json:"systemTags"`
		CompartmentId                   *string                           `json:"compartmentId"`
		PipelineId                      *string                           `json:"pipelineId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.ProjectId = model.ProjectId

	m.DisplayName = model.DisplayName

	nn, e = model.ConfigurationOverrideDetails.UnmarshalPolymorphicJSON(model.ConfigurationOverrideDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ConfigurationOverrideDetails = nn.(PipelineConfigurationDetails)
	} else {
		m.ConfigurationOverrideDetails = nil
	}

	m.LogConfigurationOverrideDetails = model.LogConfigurationOverrideDetails

	m.StepOverrideDetails = make([]PipelineStepOverrideDetails, len(model.StepOverrideDetails))
	copy(m.StepOverrideDetails, model.StepOverrideDetails)
	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.CompartmentId = model.CompartmentId

	m.PipelineId = model.PipelineId

	return
}
