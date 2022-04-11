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
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CustomScriptStepDetails The type of step where user provides the step artifact to be executed on an execution engine managed by the pipelines service.
type CustomScriptStepDetails struct {

	// The name of the step. It must be unique within the pipeline. This is used to create the pipeline DAG.
	StepName *string `mandatory:"true" json:"stepName"`

	// A short description of the step.
	Description *string `mandatory:"false" json:"description"`

	// The list of step names this current step depends on for execution.
	DependsOn []string `mandatory:"false" json:"dependsOn"`

	StepConfigurationDetails PipelineConfigurationDetails `mandatory:"false" json:"stepConfigurationDetails"`

	StepInfrastructureConfigurationDetails *PipelineInfrastructureConfigurationDetails `mandatory:"false" json:"stepInfrastructureConfigurationDetails"`

	// A flag to indicate whether the artifact has been uploaded for this step or not.
	IsArtifactUploaded *bool `mandatory:"false" json:"isArtifactUploaded"`
}

//GetStepName returns StepName
func (m CustomScriptStepDetails) GetStepName() *string {
	return m.StepName
}

//GetDescription returns Description
func (m CustomScriptStepDetails) GetDescription() *string {
	return m.Description
}

//GetDependsOn returns DependsOn
func (m CustomScriptStepDetails) GetDependsOn() []string {
	return m.DependsOn
}

//GetStepConfigurationDetails returns StepConfigurationDetails
func (m CustomScriptStepDetails) GetStepConfigurationDetails() PipelineConfigurationDetails {
	return m.StepConfigurationDetails
}

func (m CustomScriptStepDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CustomScriptStepDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CustomScriptStepDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCustomScriptStepDetails CustomScriptStepDetails
	s := struct {
		DiscriminatorParam string `json:"stepType"`
		MarshalTypeCustomScriptStepDetails
	}{
		"CUSTOM_SCRIPT",
		(MarshalTypeCustomScriptStepDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CustomScriptStepDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description                            *string                                     `json:"description"`
		DependsOn                              []string                                    `json:"dependsOn"`
		StepConfigurationDetails               pipelineconfigurationdetails                `json:"stepConfigurationDetails"`
		StepInfrastructureConfigurationDetails *PipelineInfrastructureConfigurationDetails `json:"stepInfrastructureConfigurationDetails"`
		IsArtifactUploaded                     *bool                                       `json:"isArtifactUploaded"`
		StepName                               *string                                     `json:"stepName"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.DependsOn = make([]string, len(model.DependsOn))
	for i, n := range model.DependsOn {
		m.DependsOn[i] = n
	}

	nn, e = model.StepConfigurationDetails.UnmarshalPolymorphicJSON(model.StepConfigurationDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.StepConfigurationDetails = nn.(PipelineConfigurationDetails)
	} else {
		m.StepConfigurationDetails = nil
	}

	m.StepInfrastructureConfigurationDetails = model.StepInfrastructureConfigurationDetails

	m.IsArtifactUploaded = model.IsArtifactUploaded

	m.StepName = model.StepName

	return
}
