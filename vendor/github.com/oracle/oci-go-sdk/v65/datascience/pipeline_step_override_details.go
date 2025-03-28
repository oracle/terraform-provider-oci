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

// PipelineStepOverrideDetails Override details of the step. Only Step Configuration is allowed to be overridden.
type PipelineStepOverrideDetails struct {

	// The name of the step.
	StepName *string `mandatory:"true" json:"stepName"`

	StepConfigurationDetails *PipelineStepConfigurationDetails `mandatory:"true" json:"stepConfigurationDetails"`

	StepContainerConfigurationDetails PipelineContainerConfigurationDetails `mandatory:"false" json:"stepContainerConfigurationDetails"`

	StepDataflowConfigurationDetails *PipelineDataflowConfigurationDetails `mandatory:"false" json:"stepDataflowConfigurationDetails"`
}

func (m PipelineStepOverrideDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PipelineStepOverrideDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *PipelineStepOverrideDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		StepContainerConfigurationDetails pipelinecontainerconfigurationdetails `json:"stepContainerConfigurationDetails"`
		StepDataflowConfigurationDetails  *PipelineDataflowConfigurationDetails `json:"stepDataflowConfigurationDetails"`
		StepName                          *string                               `json:"stepName"`
		StepConfigurationDetails          *PipelineStepConfigurationDetails     `json:"stepConfigurationDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.StepContainerConfigurationDetails.UnmarshalPolymorphicJSON(model.StepContainerConfigurationDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.StepContainerConfigurationDetails = nn.(PipelineContainerConfigurationDetails)
	} else {
		m.StepContainerConfigurationDetails = nil
	}

	m.StepDataflowConfigurationDetails = model.StepDataflowConfigurationDetails

	m.StepName = model.StepName

	m.StepConfigurationDetails = model.StepConfigurationDetails

	return
}
