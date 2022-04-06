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

// StepDetails A step in the pipeline
type StepDetails interface {

	// The name of the step. It must be unique within the pipeline. This is used to create the pipeline DAG.
	GetStepName() *string

	// The list of step names this current step depends on for execution
	GetDependsOn() []string

	// A short description of the step
	GetDescription() *string
}

type stepdetails struct {
	JsonData    []byte
	StepName    *string  `mandatory:"true" json:"stepName"`
	DependsOn   []string `mandatory:"true" json:"dependsOn"`
	Description *string  `mandatory:"false" json:"description"`
	StepType    string   `json:"stepType"`
}

// UnmarshalJSON unmarshals json
func (m *stepdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerstepdetails stepdetails
	s := struct {
		Model Unmarshalerstepdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.StepName = s.Model.StepName
	m.DependsOn = s.Model.DependsOn
	m.Description = s.Model.Description
	m.StepType = s.Model.StepType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *stepdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.StepType {
	case "ML_JOB":
		mm := MlJobStepDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetStepName returns StepName
func (m stepdetails) GetStepName() *string {
	return m.StepName
}

//GetDependsOn returns DependsOn
func (m stepdetails) GetDependsOn() []string {
	return m.DependsOn
}

//GetDescription returns Description
func (m stepdetails) GetDescription() *string {
	return m.Description
}

func (m stepdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m stepdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// StepDetailsStepTypeEnum Enum with underlying type: string
type StepDetailsStepTypeEnum string

// Set of constants representing the allowable values for StepDetailsStepTypeEnum
const (
	StepDetailsStepTypeMlJob StepDetailsStepTypeEnum = "ML_JOB"
)

var mappingStepDetailsStepTypeEnum = map[string]StepDetailsStepTypeEnum{
	"ML_JOB": StepDetailsStepTypeMlJob,
}

var mappingStepDetailsStepTypeEnumLowerCase = map[string]StepDetailsStepTypeEnum{
	"ml_job": StepDetailsStepTypeMlJob,
}

// GetStepDetailsStepTypeEnumValues Enumerates the set of values for StepDetailsStepTypeEnum
func GetStepDetailsStepTypeEnumValues() []StepDetailsStepTypeEnum {
	values := make([]StepDetailsStepTypeEnum, 0)
	for _, v := range mappingStepDetailsStepTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetStepDetailsStepTypeEnumStringValues Enumerates the set of values in String for StepDetailsStepTypeEnum
func GetStepDetailsStepTypeEnumStringValues() []string {
	return []string{
		"ML_JOB",
	}
}

// GetMappingStepDetailsStepTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingStepDetailsStepTypeEnum(val string) (StepDetailsStepTypeEnum, bool) {
	enum, ok := mappingStepDetailsStepTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
