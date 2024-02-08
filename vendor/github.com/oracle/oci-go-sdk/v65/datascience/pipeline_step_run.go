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

// PipelineStepRun Detail of each StepRun.
type PipelineStepRun interface {

	// The date and time the pipeline step run was started in the timestamp format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	GetTimeStarted() *common.SDKTime

	// The name of the step.
	GetStepName() *string

	// The date and time the pipeline step run finshed executing in the timestamp format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	GetTimeFinished() *common.SDKTime

	// The state of the step run.
	GetLifecycleState() PipelineStepRunLifecycleStateEnum

	// Details of the state of the step run.
	GetLifecycleDetails() *string
}

type pipelinesteprun struct {
	JsonData         []byte
	TimeFinished     *common.SDKTime                   `mandatory:"false" json:"timeFinished"`
	LifecycleState   PipelineStepRunLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
	LifecycleDetails *string                           `mandatory:"false" json:"lifecycleDetails"`
	TimeStarted      *common.SDKTime                   `mandatory:"true" json:"timeStarted"`
	StepName         *string                           `mandatory:"true" json:"stepName"`
	StepType         string                            `json:"stepType"`
}

// UnmarshalJSON unmarshals json
func (m *pipelinesteprun) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerpipelinesteprun pipelinesteprun
	s := struct {
		Model Unmarshalerpipelinesteprun
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.TimeStarted = s.Model.TimeStarted
	m.StepName = s.Model.StepName
	m.TimeFinished = s.Model.TimeFinished
	m.LifecycleState = s.Model.LifecycleState
	m.LifecycleDetails = s.Model.LifecycleDetails
	m.StepType = s.Model.StepType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *pipelinesteprun) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.StepType {
	case "CUSTOM_SCRIPT":
		mm := PipelineCustomScriptStepRun{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ML_JOB":
		mm := PipelineMlJobStepRun{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for PipelineStepRun: %s.", m.StepType)
		return *m, nil
	}
}

// GetTimeFinished returns TimeFinished
func (m pipelinesteprun) GetTimeFinished() *common.SDKTime {
	return m.TimeFinished
}

// GetLifecycleState returns LifecycleState
func (m pipelinesteprun) GetLifecycleState() PipelineStepRunLifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m pipelinesteprun) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetTimeStarted returns TimeStarted
func (m pipelinesteprun) GetTimeStarted() *common.SDKTime {
	return m.TimeStarted
}

// GetStepName returns StepName
func (m pipelinesteprun) GetStepName() *string {
	return m.StepName
}

func (m pipelinesteprun) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m pipelinesteprun) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingPipelineStepRunLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetPipelineStepRunLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PipelineStepRunLifecycleStateEnum Enum with underlying type: string
type PipelineStepRunLifecycleStateEnum string

// Set of constants representing the allowable values for PipelineStepRunLifecycleStateEnum
const (
	PipelineStepRunLifecycleStateWaiting    PipelineStepRunLifecycleStateEnum = "WAITING"
	PipelineStepRunLifecycleStateAccepted   PipelineStepRunLifecycleStateEnum = "ACCEPTED"
	PipelineStepRunLifecycleStateInProgress PipelineStepRunLifecycleStateEnum = "IN_PROGRESS"
	PipelineStepRunLifecycleStateFailed     PipelineStepRunLifecycleStateEnum = "FAILED"
	PipelineStepRunLifecycleStateSucceeded  PipelineStepRunLifecycleStateEnum = "SUCCEEDED"
	PipelineStepRunLifecycleStateCanceling  PipelineStepRunLifecycleStateEnum = "CANCELING"
	PipelineStepRunLifecycleStateCanceled   PipelineStepRunLifecycleStateEnum = "CANCELED"
	PipelineStepRunLifecycleStateDeleted    PipelineStepRunLifecycleStateEnum = "DELETED"
	PipelineStepRunLifecycleStateSkipped    PipelineStepRunLifecycleStateEnum = "SKIPPED"
)

var mappingPipelineStepRunLifecycleStateEnum = map[string]PipelineStepRunLifecycleStateEnum{
	"WAITING":     PipelineStepRunLifecycleStateWaiting,
	"ACCEPTED":    PipelineStepRunLifecycleStateAccepted,
	"IN_PROGRESS": PipelineStepRunLifecycleStateInProgress,
	"FAILED":      PipelineStepRunLifecycleStateFailed,
	"SUCCEEDED":   PipelineStepRunLifecycleStateSucceeded,
	"CANCELING":   PipelineStepRunLifecycleStateCanceling,
	"CANCELED":    PipelineStepRunLifecycleStateCanceled,
	"DELETED":     PipelineStepRunLifecycleStateDeleted,
	"SKIPPED":     PipelineStepRunLifecycleStateSkipped,
}

var mappingPipelineStepRunLifecycleStateEnumLowerCase = map[string]PipelineStepRunLifecycleStateEnum{
	"waiting":     PipelineStepRunLifecycleStateWaiting,
	"accepted":    PipelineStepRunLifecycleStateAccepted,
	"in_progress": PipelineStepRunLifecycleStateInProgress,
	"failed":      PipelineStepRunLifecycleStateFailed,
	"succeeded":   PipelineStepRunLifecycleStateSucceeded,
	"canceling":   PipelineStepRunLifecycleStateCanceling,
	"canceled":    PipelineStepRunLifecycleStateCanceled,
	"deleted":     PipelineStepRunLifecycleStateDeleted,
	"skipped":     PipelineStepRunLifecycleStateSkipped,
}

// GetPipelineStepRunLifecycleStateEnumValues Enumerates the set of values for PipelineStepRunLifecycleStateEnum
func GetPipelineStepRunLifecycleStateEnumValues() []PipelineStepRunLifecycleStateEnum {
	values := make([]PipelineStepRunLifecycleStateEnum, 0)
	for _, v := range mappingPipelineStepRunLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetPipelineStepRunLifecycleStateEnumStringValues Enumerates the set of values in String for PipelineStepRunLifecycleStateEnum
func GetPipelineStepRunLifecycleStateEnumStringValues() []string {
	return []string{
		"WAITING",
		"ACCEPTED",
		"IN_PROGRESS",
		"FAILED",
		"SUCCEEDED",
		"CANCELING",
		"CANCELED",
		"DELETED",
		"SKIPPED",
	}
}

// GetMappingPipelineStepRunLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPipelineStepRunLifecycleStateEnum(val string) (PipelineStepRunLifecycleStateEnum, bool) {
	enum, ok := mappingPipelineStepRunLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// PipelineStepRunStepTypeEnum Enum with underlying type: string
type PipelineStepRunStepTypeEnum string

// Set of constants representing the allowable values for PipelineStepRunStepTypeEnum
const (
	PipelineStepRunStepTypeMlJob        PipelineStepRunStepTypeEnum = "ML_JOB"
	PipelineStepRunStepTypeCustomScript PipelineStepRunStepTypeEnum = "CUSTOM_SCRIPT"
)

var mappingPipelineStepRunStepTypeEnum = map[string]PipelineStepRunStepTypeEnum{
	"ML_JOB":        PipelineStepRunStepTypeMlJob,
	"CUSTOM_SCRIPT": PipelineStepRunStepTypeCustomScript,
}

var mappingPipelineStepRunStepTypeEnumLowerCase = map[string]PipelineStepRunStepTypeEnum{
	"ml_job":        PipelineStepRunStepTypeMlJob,
	"custom_script": PipelineStepRunStepTypeCustomScript,
}

// GetPipelineStepRunStepTypeEnumValues Enumerates the set of values for PipelineStepRunStepTypeEnum
func GetPipelineStepRunStepTypeEnumValues() []PipelineStepRunStepTypeEnum {
	values := make([]PipelineStepRunStepTypeEnum, 0)
	for _, v := range mappingPipelineStepRunStepTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPipelineStepRunStepTypeEnumStringValues Enumerates the set of values in String for PipelineStepRunStepTypeEnum
func GetPipelineStepRunStepTypeEnumStringValues() []string {
	return []string{
		"ML_JOB",
		"CUSTOM_SCRIPT",
	}
}

// GetMappingPipelineStepRunStepTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPipelineStepRunStepTypeEnum(val string) (PipelineStepRunStepTypeEnum, bool) {
	enum, ok := mappingPipelineStepRunStepTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
