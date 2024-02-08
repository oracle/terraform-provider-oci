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

// PipelineStepDetails A step in the pipeline.
type PipelineStepDetails interface {

	// The name of the step. It must be unique within the pipeline. This is used to create the pipeline DAG.
	GetStepName() *string

	// A short description of the step.
	GetDescription() *string

	// The list of step names this current step depends on for execution.
	GetDependsOn() []string

	GetStepConfigurationDetails() *PipelineStepConfigurationDetails
}

type pipelinestepdetails struct {
	JsonData                 []byte
	Description              *string                           `mandatory:"false" json:"description"`
	DependsOn                []string                          `mandatory:"false" json:"dependsOn"`
	StepConfigurationDetails *PipelineStepConfigurationDetails `mandatory:"false" json:"stepConfigurationDetails"`
	StepName                 *string                           `mandatory:"true" json:"stepName"`
	StepType                 string                            `json:"stepType"`
}

// UnmarshalJSON unmarshals json
func (m *pipelinestepdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerpipelinestepdetails pipelinestepdetails
	s := struct {
		Model Unmarshalerpipelinestepdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.StepName = s.Model.StepName
	m.Description = s.Model.Description
	m.DependsOn = s.Model.DependsOn
	m.StepConfigurationDetails = s.Model.StepConfigurationDetails
	m.StepType = s.Model.StepType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *pipelinestepdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.StepType {
	case "ML_JOB":
		mm := PipelineMlJobStepDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CUSTOM_SCRIPT":
		mm := PipelineCustomScriptStepDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for PipelineStepDetails: %s.", m.StepType)
		return *m, nil
	}
}

// GetDescription returns Description
func (m pipelinestepdetails) GetDescription() *string {
	return m.Description
}

// GetDependsOn returns DependsOn
func (m pipelinestepdetails) GetDependsOn() []string {
	return m.DependsOn
}

// GetStepConfigurationDetails returns StepConfigurationDetails
func (m pipelinestepdetails) GetStepConfigurationDetails() *PipelineStepConfigurationDetails {
	return m.StepConfigurationDetails
}

// GetStepName returns StepName
func (m pipelinestepdetails) GetStepName() *string {
	return m.StepName
}

func (m pipelinestepdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m pipelinestepdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PipelineStepDetailsStepTypeEnum Enum with underlying type: string
type PipelineStepDetailsStepTypeEnum string

// Set of constants representing the allowable values for PipelineStepDetailsStepTypeEnum
const (
	PipelineStepDetailsStepTypeMlJob        PipelineStepDetailsStepTypeEnum = "ML_JOB"
	PipelineStepDetailsStepTypeCustomScript PipelineStepDetailsStepTypeEnum = "CUSTOM_SCRIPT"
)

var mappingPipelineStepDetailsStepTypeEnum = map[string]PipelineStepDetailsStepTypeEnum{
	"ML_JOB":        PipelineStepDetailsStepTypeMlJob,
	"CUSTOM_SCRIPT": PipelineStepDetailsStepTypeCustomScript,
}

var mappingPipelineStepDetailsStepTypeEnumLowerCase = map[string]PipelineStepDetailsStepTypeEnum{
	"ml_job":        PipelineStepDetailsStepTypeMlJob,
	"custom_script": PipelineStepDetailsStepTypeCustomScript,
}

// GetPipelineStepDetailsStepTypeEnumValues Enumerates the set of values for PipelineStepDetailsStepTypeEnum
func GetPipelineStepDetailsStepTypeEnumValues() []PipelineStepDetailsStepTypeEnum {
	values := make([]PipelineStepDetailsStepTypeEnum, 0)
	for _, v := range mappingPipelineStepDetailsStepTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPipelineStepDetailsStepTypeEnumStringValues Enumerates the set of values in String for PipelineStepDetailsStepTypeEnum
func GetPipelineStepDetailsStepTypeEnumStringValues() []string {
	return []string{
		"ML_JOB",
		"CUSTOM_SCRIPT",
	}
}

// GetMappingPipelineStepDetailsStepTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPipelineStepDetailsStepTypeEnum(val string) (PipelineStepDetailsStepTypeEnum, bool) {
	enum, ok := mappingPipelineStepDetailsStepTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
