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

// PipelineCustomScriptStepUpdateDetails The type of step where user provides the step artifact to be executed on an execution engine managed by the pipelines service.
type PipelineCustomScriptStepUpdateDetails struct {

	// The name of the step.
	StepName *string `mandatory:"true" json:"stepName"`

	// A short description of the step.
	Description *string `mandatory:"false" json:"description"`

	StepConfigurationDetails *PipelineStepConfigurationDetails `mandatory:"false" json:"stepConfigurationDetails"`

	StepInfrastructureConfigurationDetails *PipelineInfrastructureConfigurationDetails `mandatory:"false" json:"stepInfrastructureConfigurationDetails"`

	// The storage mount details to mount to the instance running the pipeline step.
	StepStorageMountConfigurationDetailsList []StorageMountConfigurationDetails `mandatory:"false" json:"stepStorageMountConfigurationDetailsList"`
}

// GetStepName returns StepName
func (m PipelineCustomScriptStepUpdateDetails) GetStepName() *string {
	return m.StepName
}

// GetDescription returns Description
func (m PipelineCustomScriptStepUpdateDetails) GetDescription() *string {
	return m.Description
}

// GetStepConfigurationDetails returns StepConfigurationDetails
func (m PipelineCustomScriptStepUpdateDetails) GetStepConfigurationDetails() *PipelineStepConfigurationDetails {
	return m.StepConfigurationDetails
}

func (m PipelineCustomScriptStepUpdateDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PipelineCustomScriptStepUpdateDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PipelineCustomScriptStepUpdateDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePipelineCustomScriptStepUpdateDetails PipelineCustomScriptStepUpdateDetails
	s := struct {
		DiscriminatorParam string `json:"stepType"`
		MarshalTypePipelineCustomScriptStepUpdateDetails
	}{
		"CUSTOM_SCRIPT",
		(MarshalTypePipelineCustomScriptStepUpdateDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *PipelineCustomScriptStepUpdateDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description                              *string                                     `json:"description"`
		StepConfigurationDetails                 *PipelineStepConfigurationDetails           `json:"stepConfigurationDetails"`
		StepInfrastructureConfigurationDetails   *PipelineInfrastructureConfigurationDetails `json:"stepInfrastructureConfigurationDetails"`
		StepStorageMountConfigurationDetailsList []storagemountconfigurationdetails          `json:"stepStorageMountConfigurationDetailsList"`
		StepName                                 *string                                     `json:"stepName"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.StepConfigurationDetails = model.StepConfigurationDetails

	m.StepInfrastructureConfigurationDetails = model.StepInfrastructureConfigurationDetails

	m.StepStorageMountConfigurationDetailsList = make([]StorageMountConfigurationDetails, len(model.StepStorageMountConfigurationDetailsList))
	for i, n := range model.StepStorageMountConfigurationDetailsList {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.StepStorageMountConfigurationDetailsList[i] = nn.(StorageMountConfigurationDetails)
		} else {
			m.StepStorageMountConfigurationDetailsList[i] = nil
		}
	}
	m.StepName = model.StepName

	return
}
