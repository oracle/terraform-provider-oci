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

// PipelineContainerStepDetails The type of step where user provides the container details for an execution engine managed by the pipelines service.
type PipelineContainerStepDetails struct {

	// The name of the step. It must be unique within the pipeline. This is used to create the pipeline DAG.
	StepName *string `mandatory:"true" json:"stepName"`

	StepContainerConfigurationDetails PipelineContainerConfigurationDetails `mandatory:"true" json:"stepContainerConfigurationDetails"`

	// A short description of the step.
	Description *string `mandatory:"false" json:"description"`

	// The list of step names this current step depends on for execution.
	DependsOn []string `mandatory:"false" json:"dependsOn"`

	StepConfigurationDetails *PipelineStepConfigurationDetails `mandatory:"false" json:"stepConfigurationDetails"`

	StepInfrastructureConfigurationDetails *PipelineInfrastructureConfigurationDetails `mandatory:"false" json:"stepInfrastructureConfigurationDetails"`

	// A flag to indicate whether the artifact has been uploaded for this step or not.
	IsArtifactUploaded *bool `mandatory:"false" json:"isArtifactUploaded"`
}

// GetStepName returns StepName
func (m PipelineContainerStepDetails) GetStepName() *string {
	return m.StepName
}

// GetDescription returns Description
func (m PipelineContainerStepDetails) GetDescription() *string {
	return m.Description
}

// GetDependsOn returns DependsOn
func (m PipelineContainerStepDetails) GetDependsOn() []string {
	return m.DependsOn
}

// GetStepConfigurationDetails returns StepConfigurationDetails
func (m PipelineContainerStepDetails) GetStepConfigurationDetails() *PipelineStepConfigurationDetails {
	return m.StepConfigurationDetails
}

func (m PipelineContainerStepDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PipelineContainerStepDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PipelineContainerStepDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePipelineContainerStepDetails PipelineContainerStepDetails
	s := struct {
		DiscriminatorParam string `json:"stepType"`
		MarshalTypePipelineContainerStepDetails
	}{
		"CONTAINER",
		(MarshalTypePipelineContainerStepDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *PipelineContainerStepDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description                            *string                                     `json:"description"`
		DependsOn                              []string                                    `json:"dependsOn"`
		StepConfigurationDetails               *PipelineStepConfigurationDetails           `json:"stepConfigurationDetails"`
		StepInfrastructureConfigurationDetails *PipelineInfrastructureConfigurationDetails `json:"stepInfrastructureConfigurationDetails"`
		IsArtifactUploaded                     *bool                                       `json:"isArtifactUploaded"`
		StepName                               *string                                     `json:"stepName"`
		StepContainerConfigurationDetails      pipelinecontainerconfigurationdetails       `json:"stepContainerConfigurationDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.DependsOn = make([]string, len(model.DependsOn))
	copy(m.DependsOn, model.DependsOn)
	m.StepConfigurationDetails = model.StepConfigurationDetails

	m.StepInfrastructureConfigurationDetails = model.StepInfrastructureConfigurationDetails

	m.IsArtifactUploaded = model.IsArtifactUploaded

	m.StepName = model.StepName

	nn, e = model.StepContainerConfigurationDetails.UnmarshalPolymorphicJSON(model.StepContainerConfigurationDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.StepContainerConfigurationDetails = nn.(PipelineContainerConfigurationDetails)
	} else {
		m.StepContainerConfigurationDetails = nil
	}

	return
}
