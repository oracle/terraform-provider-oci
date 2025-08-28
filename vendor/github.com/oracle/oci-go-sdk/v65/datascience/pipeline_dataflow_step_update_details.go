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

// PipelineDataflowStepUpdateDetails The type of step where the dataflow application is pre-created by the user.
type PipelineDataflowStepUpdateDetails struct {

	// The name of the step.
	StepName *string `mandatory:"true" json:"stepName"`

	// A short description of the step.
	Description *string `mandatory:"false" json:"description"`

	StepConfigurationDetails *PipelineStepConfigurationDetails `mandatory:"false" json:"stepConfigurationDetails"`

	StepDataflowConfigurationDetails *PipelineDataflowConfigurationDetails `mandatory:"false" json:"stepDataflowConfigurationDetails"`
}

// GetStepName returns StepName
func (m PipelineDataflowStepUpdateDetails) GetStepName() *string {
	return m.StepName
}

// GetDescription returns Description
func (m PipelineDataflowStepUpdateDetails) GetDescription() *string {
	return m.Description
}

// GetStepConfigurationDetails returns StepConfigurationDetails
func (m PipelineDataflowStepUpdateDetails) GetStepConfigurationDetails() *PipelineStepConfigurationDetails {
	return m.StepConfigurationDetails
}

func (m PipelineDataflowStepUpdateDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PipelineDataflowStepUpdateDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PipelineDataflowStepUpdateDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePipelineDataflowStepUpdateDetails PipelineDataflowStepUpdateDetails
	s := struct {
		DiscriminatorParam string `json:"stepType"`
		MarshalTypePipelineDataflowStepUpdateDetails
	}{
		"DATAFLOW",
		(MarshalTypePipelineDataflowStepUpdateDetails)(m),
	}

	return json.Marshal(&s)
}
