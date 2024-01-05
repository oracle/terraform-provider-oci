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

// CreatePipelineDetails The information about new Pipeline.
type CreatePipelineDetails struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project to associate the pipeline with.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where you want to create the pipeline.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Array of step details for each step.
	StepDetails []PipelineStepDetails `mandatory:"true" json:"stepDetails"`

	// A user-friendly display name for the resource.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A short description of the pipeline.
	Description *string `mandatory:"false" json:"description"`

	ConfigurationDetails PipelineConfigurationDetails `mandatory:"false" json:"configurationDetails"`

	LogConfigurationDetails *PipelineLogConfigurationDetails `mandatory:"false" json:"logConfigurationDetails"`

	InfrastructureConfigurationDetails *PipelineInfrastructureConfigurationDetails `mandatory:"false" json:"infrastructureConfigurationDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreatePipelineDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreatePipelineDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreatePipelineDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName                        *string                                     `json:"displayName"`
		Description                        *string                                     `json:"description"`
		ConfigurationDetails               pipelineconfigurationdetails                `json:"configurationDetails"`
		LogConfigurationDetails            *PipelineLogConfigurationDetails            `json:"logConfigurationDetails"`
		InfrastructureConfigurationDetails *PipelineInfrastructureConfigurationDetails `json:"infrastructureConfigurationDetails"`
		FreeformTags                       map[string]string                           `json:"freeformTags"`
		DefinedTags                        map[string]map[string]interface{}           `json:"definedTags"`
		ProjectId                          *string                                     `json:"projectId"`
		CompartmentId                      *string                                     `json:"compartmentId"`
		StepDetails                        []pipelinestepdetails                       `json:"stepDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	nn, e = model.ConfigurationDetails.UnmarshalPolymorphicJSON(model.ConfigurationDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ConfigurationDetails = nn.(PipelineConfigurationDetails)
	} else {
		m.ConfigurationDetails = nil
	}

	m.LogConfigurationDetails = model.LogConfigurationDetails

	m.InfrastructureConfigurationDetails = model.InfrastructureConfigurationDetails

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.ProjectId = model.ProjectId

	m.CompartmentId = model.CompartmentId

	m.StepDetails = make([]PipelineStepDetails, len(model.StepDetails))
	for i, n := range model.StepDetails {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.StepDetails[i] = nn.(PipelineStepDetails)
		} else {
			m.StepDetails[i] = nil
		}
	}
	return
}
