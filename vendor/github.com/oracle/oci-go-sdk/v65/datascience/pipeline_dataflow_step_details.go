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

// PipelineDataflowStepDetails The type of step where the dataflow application is pre-created by the user.
type PipelineDataflowStepDetails struct {

	// The name of the step. It must be unique within the pipeline. This is used to create the pipeline DAG.
	StepName *string `mandatory:"true" json:"stepName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the dataflow application to be used as a step.
	ApplicationId *string `mandatory:"true" json:"applicationId"`

	// A short description of the step.
	Description *string `mandatory:"false" json:"description"`

	// The list of step names this current step depends on for execution.
	DependsOn []string `mandatory:"false" json:"dependsOn"`

	StepConfigurationDetails *PipelineStepConfigurationDetails `mandatory:"false" json:"stepConfigurationDetails"`

	StepDataflowConfigurationDetails *PipelineDataflowConfigurationDetails `mandatory:"false" json:"stepDataflowConfigurationDetails"`
}

// GetStepName returns StepName
func (m PipelineDataflowStepDetails) GetStepName() *string {
	return m.StepName
}

// GetDescription returns Description
func (m PipelineDataflowStepDetails) GetDescription() *string {
	return m.Description
}

// GetDependsOn returns DependsOn
func (m PipelineDataflowStepDetails) GetDependsOn() []string {
	return m.DependsOn
}

// GetStepConfigurationDetails returns StepConfigurationDetails
func (m PipelineDataflowStepDetails) GetStepConfigurationDetails() *PipelineStepConfigurationDetails {
	return m.StepConfigurationDetails
}

func (m PipelineDataflowStepDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PipelineDataflowStepDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PipelineDataflowStepDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePipelineDataflowStepDetails PipelineDataflowStepDetails
	s := struct {
		DiscriminatorParam string `json:"stepType"`
		MarshalTypePipelineDataflowStepDetails
	}{
		"DATAFLOW",
		(MarshalTypePipelineDataflowStepDetails)(m),
	}

	return json.Marshal(&s)
}
