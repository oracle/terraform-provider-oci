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

// Pipeline A Pipeline to orchestrate and execute machine learning workflows.
type Pipeline struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the pipeline.
	Id *string `mandatory:"true" json:"id"`

	// The date and time the resource was created in the timestamp format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: 2020-08-06T21:10:29.41Z
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user who created the pipeline.
	CreatedBy *string `mandatory:"true" json:"createdBy"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project to associate the pipeline with.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where you want to create the pipeline.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly display name for the resource.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Array of step details for each step.
	StepDetails []PipelineStepDetails `mandatory:"true" json:"stepDetails"`

	// The current state of the pipeline.
	LifecycleState PipelineLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the resource was updated in the timestamp format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: 2020-08-06T21:10:29.41Z
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A short description of the pipeline.
	Description *string `mandatory:"false" json:"description"`

	ConfigurationDetails PipelineConfigurationDetails `mandatory:"false" json:"configurationDetails"`

	LogConfigurationDetails *PipelineLogConfigurationDetails `mandatory:"false" json:"logConfigurationDetails"`

	InfrastructureConfigurationDetails *PipelineInfrastructureConfigurationDetails `mandatory:"false" json:"infrastructureConfigurationDetails"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in 'Failed' state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

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

func (m Pipeline) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Pipeline) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPipelineLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetPipelineLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *Pipeline) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		TimeUpdated                        *common.SDKTime                             `json:"timeUpdated"`
		Description                        *string                                     `json:"description"`
		ConfigurationDetails               pipelineconfigurationdetails                `json:"configurationDetails"`
		LogConfigurationDetails            *PipelineLogConfigurationDetails            `json:"logConfigurationDetails"`
		InfrastructureConfigurationDetails *PipelineInfrastructureConfigurationDetails `json:"infrastructureConfigurationDetails"`
		LifecycleDetails                   *string                                     `json:"lifecycleDetails"`
		FreeformTags                       map[string]string                           `json:"freeformTags"`
		DefinedTags                        map[string]map[string]interface{}           `json:"definedTags"`
		SystemTags                         map[string]map[string]interface{}           `json:"systemTags"`
		Id                                 *string                                     `json:"id"`
		TimeCreated                        *common.SDKTime                             `json:"timeCreated"`
		CreatedBy                          *string                                     `json:"createdBy"`
		ProjectId                          *string                                     `json:"projectId"`
		CompartmentId                      *string                                     `json:"compartmentId"`
		DisplayName                        *string                                     `json:"displayName"`
		StepDetails                        []pipelinestepdetails                       `json:"stepDetails"`
		LifecycleState                     PipelineLifecycleStateEnum                  `json:"lifecycleState"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.TimeUpdated = model.TimeUpdated

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

	m.LifecycleDetails = model.LifecycleDetails

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.TimeCreated = model.TimeCreated

	m.CreatedBy = model.CreatedBy

	m.ProjectId = model.ProjectId

	m.CompartmentId = model.CompartmentId

	m.DisplayName = model.DisplayName

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
	m.LifecycleState = model.LifecycleState

	return
}
