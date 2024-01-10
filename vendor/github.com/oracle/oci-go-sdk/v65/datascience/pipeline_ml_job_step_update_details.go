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

// PipelineMlJobStepUpdateDetails The type of step where the job is pre-created by the user.
type PipelineMlJobStepUpdateDetails struct {

	// The name of the step.
	StepName *string `mandatory:"true" json:"stepName"`

	// A short description of the step.
	Description *string `mandatory:"false" json:"description"`

	StepConfigurationDetails *PipelineStepConfigurationDetails `mandatory:"false" json:"stepConfigurationDetails"`
}

// GetStepName returns StepName
func (m PipelineMlJobStepUpdateDetails) GetStepName() *string {
	return m.StepName
}

// GetDescription returns Description
func (m PipelineMlJobStepUpdateDetails) GetDescription() *string {
	return m.Description
}

// GetStepConfigurationDetails returns StepConfigurationDetails
func (m PipelineMlJobStepUpdateDetails) GetStepConfigurationDetails() *PipelineStepConfigurationDetails {
	return m.StepConfigurationDetails
}

func (m PipelineMlJobStepUpdateDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PipelineMlJobStepUpdateDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PipelineMlJobStepUpdateDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePipelineMlJobStepUpdateDetails PipelineMlJobStepUpdateDetails
	s := struct {
		DiscriminatorParam string `json:"stepType"`
		MarshalTypePipelineMlJobStepUpdateDetails
	}{
		"ML_JOB",
		(MarshalTypePipelineMlJobStepUpdateDetails)(m),
	}

	return json.Marshal(&s)
}
