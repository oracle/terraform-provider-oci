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

// PipelineStepUpdateDetails The details of the step to update.
type PipelineStepUpdateDetails interface {

	// The name of the step.
	GetStepName() *string

	// A short description of the step.
	GetDescription() *string

	GetStepConfigurationDetails() *PipelineStepConfigurationDetails
}

type pipelinestepupdatedetails struct {
	JsonData                 []byte
	Description              *string                           `mandatory:"false" json:"description"`
	StepConfigurationDetails *PipelineStepConfigurationDetails `mandatory:"false" json:"stepConfigurationDetails"`
	StepName                 *string                           `mandatory:"true" json:"stepName"`
	StepType                 string                            `json:"stepType"`
}

// UnmarshalJSON unmarshals json
func (m *pipelinestepupdatedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerpipelinestepupdatedetails pipelinestepupdatedetails
	s := struct {
		Model Unmarshalerpipelinestepupdatedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.StepName = s.Model.StepName
	m.Description = s.Model.Description
	m.StepConfigurationDetails = s.Model.StepConfigurationDetails
	m.StepType = s.Model.StepType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *pipelinestepupdatedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.StepType {
	case "ML_JOB":
		mm := PipelineMlJobStepUpdateDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CUSTOM_SCRIPT":
		mm := PipelineCustomScriptStepUpdateDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for PipelineStepUpdateDetails: %s.", m.StepType)
		return *m, nil
	}
}

// GetDescription returns Description
func (m pipelinestepupdatedetails) GetDescription() *string {
	return m.Description
}

// GetStepConfigurationDetails returns StepConfigurationDetails
func (m pipelinestepupdatedetails) GetStepConfigurationDetails() *PipelineStepConfigurationDetails {
	return m.StepConfigurationDetails
}

// GetStepName returns StepName
func (m pipelinestepupdatedetails) GetStepName() *string {
	return m.StepName
}

func (m pipelinestepupdatedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m pipelinestepupdatedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PipelineStepUpdateDetailsStepTypeEnum Enum with underlying type: string
type PipelineStepUpdateDetailsStepTypeEnum string

// Set of constants representing the allowable values for PipelineStepUpdateDetailsStepTypeEnum
const (
	PipelineStepUpdateDetailsStepTypeMlJob        PipelineStepUpdateDetailsStepTypeEnum = "ML_JOB"
	PipelineStepUpdateDetailsStepTypeCustomScript PipelineStepUpdateDetailsStepTypeEnum = "CUSTOM_SCRIPT"
)

var mappingPipelineStepUpdateDetailsStepTypeEnum = map[string]PipelineStepUpdateDetailsStepTypeEnum{
	"ML_JOB":        PipelineStepUpdateDetailsStepTypeMlJob,
	"CUSTOM_SCRIPT": PipelineStepUpdateDetailsStepTypeCustomScript,
}

var mappingPipelineStepUpdateDetailsStepTypeEnumLowerCase = map[string]PipelineStepUpdateDetailsStepTypeEnum{
	"ml_job":        PipelineStepUpdateDetailsStepTypeMlJob,
	"custom_script": PipelineStepUpdateDetailsStepTypeCustomScript,
}

// GetPipelineStepUpdateDetailsStepTypeEnumValues Enumerates the set of values for PipelineStepUpdateDetailsStepTypeEnum
func GetPipelineStepUpdateDetailsStepTypeEnumValues() []PipelineStepUpdateDetailsStepTypeEnum {
	values := make([]PipelineStepUpdateDetailsStepTypeEnum, 0)
	for _, v := range mappingPipelineStepUpdateDetailsStepTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPipelineStepUpdateDetailsStepTypeEnumStringValues Enumerates the set of values in String for PipelineStepUpdateDetailsStepTypeEnum
func GetPipelineStepUpdateDetailsStepTypeEnumStringValues() []string {
	return []string{
		"ML_JOB",
		"CUSTOM_SCRIPT",
	}
}

// GetMappingPipelineStepUpdateDetailsStepTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPipelineStepUpdateDetailsStepTypeEnum(val string) (PipelineStepUpdateDetailsStepTypeEnum, bool) {
	enum, ok := mappingPipelineStepUpdateDetailsStepTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
