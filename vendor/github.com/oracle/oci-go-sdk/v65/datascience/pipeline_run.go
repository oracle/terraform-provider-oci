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

// PipelineRun Description of PipelineRun.
type PipelineRun struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the pipeline run.
	Id *string `mandatory:"true" json:"id"`

	// The date and time the pipeline run was accepted in the timestamp format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeAccepted *common.SDKTime `mandatory:"true" json:"timeAccepted"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user who created the pipeline run.
	CreatedBy *string `mandatory:"true" json:"createdBy"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project to associate the pipeline run with.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where you want to create the pipeline run.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the pipeline.
	PipelineId *string `mandatory:"true" json:"pipelineId"`

	// A user-friendly display name for the resource.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Array of StepRun object for each step.
	StepRuns []PipelineStepRun `mandatory:"true" json:"stepRuns"`

	// The current state of the pipeline run.
	LifecycleState PipelineRunLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the pipeline run request was started in the timestamp format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The date and time the pipeline run was updated in the timestamp format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The date and time the pipeline run request was finished in the timestamp format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`

	ConfigurationDetails PipelineConfigurationDetails `mandatory:"false" json:"configurationDetails"`

	ConfigurationOverrideDetails PipelineConfigurationDetails `mandatory:"false" json:"configurationOverrideDetails"`

	LogConfigurationOverrideDetails *PipelineLogConfigurationDetails `mandatory:"false" json:"logConfigurationOverrideDetails"`

	// Array of step override details. Only Step Configuration is allowed to be overridden.
	StepOverrideDetails []PipelineStepOverrideDetails `mandatory:"false" json:"stepOverrideDetails"`

	LogDetails *PipelineRunLogDetails `mandatory:"false" json:"logDetails"`

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

func (m PipelineRun) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PipelineRun) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPipelineRunLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetPipelineRunLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *PipelineRun) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		TimeStarted                     *common.SDKTime                   `json:"timeStarted"`
		TimeUpdated                     *common.SDKTime                   `json:"timeUpdated"`
		TimeFinished                    *common.SDKTime                   `json:"timeFinished"`
		ConfigurationDetails            pipelineconfigurationdetails      `json:"configurationDetails"`
		ConfigurationOverrideDetails    pipelineconfigurationdetails      `json:"configurationOverrideDetails"`
		LogConfigurationOverrideDetails *PipelineLogConfigurationDetails  `json:"logConfigurationOverrideDetails"`
		StepOverrideDetails             []PipelineStepOverrideDetails     `json:"stepOverrideDetails"`
		LogDetails                      *PipelineRunLogDetails            `json:"logDetails"`
		LifecycleDetails                *string                           `json:"lifecycleDetails"`
		FreeformTags                    map[string]string                 `json:"freeformTags"`
		DefinedTags                     map[string]map[string]interface{} `json:"definedTags"`
		SystemTags                      map[string]map[string]interface{} `json:"systemTags"`
		Id                              *string                           `json:"id"`
		TimeAccepted                    *common.SDKTime                   `json:"timeAccepted"`
		CreatedBy                       *string                           `json:"createdBy"`
		ProjectId                       *string                           `json:"projectId"`
		CompartmentId                   *string                           `json:"compartmentId"`
		PipelineId                      *string                           `json:"pipelineId"`
		DisplayName                     *string                           `json:"displayName"`
		StepRuns                        []pipelinesteprun                 `json:"stepRuns"`
		LifecycleState                  PipelineRunLifecycleStateEnum     `json:"lifecycleState"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.TimeStarted = model.TimeStarted

	m.TimeUpdated = model.TimeUpdated

	m.TimeFinished = model.TimeFinished

	nn, e = model.ConfigurationDetails.UnmarshalPolymorphicJSON(model.ConfigurationDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ConfigurationDetails = nn.(PipelineConfigurationDetails)
	} else {
		m.ConfigurationDetails = nil
	}

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
	m.LogDetails = model.LogDetails

	m.LifecycleDetails = model.LifecycleDetails

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.TimeAccepted = model.TimeAccepted

	m.CreatedBy = model.CreatedBy

	m.ProjectId = model.ProjectId

	m.CompartmentId = model.CompartmentId

	m.PipelineId = model.PipelineId

	m.DisplayName = model.DisplayName

	m.StepRuns = make([]PipelineStepRun, len(model.StepRuns))
	for i, n := range model.StepRuns {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.StepRuns[i] = nn.(PipelineStepRun)
		} else {
			m.StepRuns[i] = nil
		}
	}
	m.LifecycleState = model.LifecycleState

	return
}
