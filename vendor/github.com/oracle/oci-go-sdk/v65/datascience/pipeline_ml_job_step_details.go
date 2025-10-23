// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// PipelineMlJobStepDetails The type of step where the job is pre-created by the user.
type PipelineMlJobStepDetails struct {

	// The name of the step. It must be unique within the pipeline. This is used to create the pipeline DAG.
	StepName *string `mandatory:"true" json:"stepName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the job to be used as a step.
	JobId *string `mandatory:"true" json:"jobId"`

	// A short description of the step.
	Description *string `mandatory:"false" json:"description"`

	// The list of step names this current step depends on for execution.
	DependsOn []string `mandatory:"false" json:"dependsOn"`

	StepConfigurationDetails *PipelineStepConfigurationDetails `mandatory:"false" json:"stepConfigurationDetails"`

	// Name used when creating the steprun.
	StepRunName *string `mandatory:"false" json:"stepRunName"`

	StepParameters PipelineStepParameterDetails `mandatory:"false" json:"stepParameters"`
}

// GetStepName returns StepName
func (m PipelineMlJobStepDetails) GetStepName() *string {
	return m.StepName
}

// GetDescription returns Description
func (m PipelineMlJobStepDetails) GetDescription() *string {
	return m.Description
}

// GetDependsOn returns DependsOn
func (m PipelineMlJobStepDetails) GetDependsOn() []string {
	return m.DependsOn
}

// GetStepConfigurationDetails returns StepConfigurationDetails
func (m PipelineMlJobStepDetails) GetStepConfigurationDetails() *PipelineStepConfigurationDetails {
	return m.StepConfigurationDetails
}

func (m PipelineMlJobStepDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PipelineMlJobStepDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PipelineMlJobStepDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePipelineMlJobStepDetails PipelineMlJobStepDetails
	s := struct {
		DiscriminatorParam string `json:"stepType"`
		MarshalTypePipelineMlJobStepDetails
	}{
		"ML_JOB",
		(MarshalTypePipelineMlJobStepDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *PipelineMlJobStepDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description              *string                           `json:"description"`
		DependsOn                []string                          `json:"dependsOn"`
		StepConfigurationDetails *PipelineStepConfigurationDetails `json:"stepConfigurationDetails"`
		StepRunName              *string                           `json:"stepRunName"`
		StepParameters           pipelinestepparameterdetails      `json:"stepParameters"`
		StepName                 *string                           `json:"stepName"`
		JobId                    *string                           `json:"jobId"`
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

	m.StepRunName = model.StepRunName

	nn, e = model.StepParameters.UnmarshalPolymorphicJSON(model.StepParameters.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.StepParameters = nn.(PipelineStepParameterDetails)
	} else {
		m.StepParameters = nil
	}

	m.StepName = model.StepName

	m.JobId = model.JobId

	return
}
