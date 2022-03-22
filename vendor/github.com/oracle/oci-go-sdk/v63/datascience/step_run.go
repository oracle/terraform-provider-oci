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
	"github.com/oracle/oci-go-sdk/v63/common"
	"strings"
)

// StepRun Detail of each StepRun
type StepRun interface {

	// The date and time the pipeline step run was started in the timestamp format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	GetTimeStarted() *common.SDKTime

	// The name of the step being run
	GetStepName() *string

	// The date and time the pipeline step run finshed executing in the timestamp format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	GetTimeFinished() *common.SDKTime

	// The state of the step run.
	GetLifecycleState() StepRunLifecycleStateEnum

	// Details of the state of the step run.
	GetLifecycleDetails() *string
}

type steprun struct {
	JsonData         []byte
	TimeStarted      *common.SDKTime           `mandatory:"true" json:"timeStarted"`
	StepName         *string                   `mandatory:"true" json:"stepName"`
	TimeFinished     *common.SDKTime           `mandatory:"false" json:"timeFinished"`
	LifecycleState   StepRunLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
	LifecycleDetails *string                   `mandatory:"false" json:"lifecycleDetails"`
	StepType         string                    `json:"stepType"`
}

// UnmarshalJSON unmarshals json
func (m *steprun) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalersteprun steprun
	s := struct {
		Model Unmarshalersteprun
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
func (m *steprun) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.StepType {
	case "ML_JOB":
		mm := MlJobStepRun{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetTimeStarted returns TimeStarted
func (m steprun) GetTimeStarted() *common.SDKTime {
	return m.TimeStarted
}

//GetStepName returns StepName
func (m steprun) GetStepName() *string {
	return m.StepName
}

//GetTimeFinished returns TimeFinished
func (m steprun) GetTimeFinished() *common.SDKTime {
	return m.TimeFinished
}

//GetLifecycleState returns LifecycleState
func (m steprun) GetLifecycleState() StepRunLifecycleStateEnum {
	return m.LifecycleState
}

//GetLifecycleDetails returns LifecycleDetails
func (m steprun) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

func (m steprun) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m steprun) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingStepRunLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetStepRunLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// StepRunLifecycleStateEnum Enum with underlying type: string
type StepRunLifecycleStateEnum string

// Set of constants representing the allowable values for StepRunLifecycleStateEnum
const (
	StepRunLifecycleStateWaiting    StepRunLifecycleStateEnum = "WAITING"
	StepRunLifecycleStateAccepted   StepRunLifecycleStateEnum = "ACCEPTED"
	StepRunLifecycleStateInProgress StepRunLifecycleStateEnum = "IN_PROGRESS"
	StepRunLifecycleStateFailed     StepRunLifecycleStateEnum = "FAILED"
	StepRunLifecycleStateSucceeded  StepRunLifecycleStateEnum = "SUCCEEDED"
	StepRunLifecycleStateCanceling  StepRunLifecycleStateEnum = "CANCELING"
	StepRunLifecycleStateCanceled   StepRunLifecycleStateEnum = "CANCELED"
	StepRunLifecycleStateDeleted    StepRunLifecycleStateEnum = "DELETED"
	StepRunLifecycleStateSkipped    StepRunLifecycleStateEnum = "SKIPPED"
)

var mappingStepRunLifecycleStateEnum = map[string]StepRunLifecycleStateEnum{
	"WAITING":     StepRunLifecycleStateWaiting,
	"ACCEPTED":    StepRunLifecycleStateAccepted,
	"IN_PROGRESS": StepRunLifecycleStateInProgress,
	"FAILED":      StepRunLifecycleStateFailed,
	"SUCCEEDED":   StepRunLifecycleStateSucceeded,
	"CANCELING":   StepRunLifecycleStateCanceling,
	"CANCELED":    StepRunLifecycleStateCanceled,
	"DELETED":     StepRunLifecycleStateDeleted,
	"SKIPPED":     StepRunLifecycleStateSkipped,
}

var mappingStepRunLifecycleStateEnumLowerCase = map[string]StepRunLifecycleStateEnum{
	"waiting":     StepRunLifecycleStateWaiting,
	"accepted":    StepRunLifecycleStateAccepted,
	"in_progress": StepRunLifecycleStateInProgress,
	"failed":      StepRunLifecycleStateFailed,
	"succeeded":   StepRunLifecycleStateSucceeded,
	"canceling":   StepRunLifecycleStateCanceling,
	"canceled":    StepRunLifecycleStateCanceled,
	"deleted":     StepRunLifecycleStateDeleted,
	"skipped":     StepRunLifecycleStateSkipped,
}

// GetStepRunLifecycleStateEnumValues Enumerates the set of values for StepRunLifecycleStateEnum
func GetStepRunLifecycleStateEnumValues() []StepRunLifecycleStateEnum {
	values := make([]StepRunLifecycleStateEnum, 0)
	for _, v := range mappingStepRunLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetStepRunLifecycleStateEnumStringValues Enumerates the set of values in String for StepRunLifecycleStateEnum
func GetStepRunLifecycleStateEnumStringValues() []string {
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

// GetMappingStepRunLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingStepRunLifecycleStateEnum(val string) (StepRunLifecycleStateEnum, bool) {
	enum, ok := mappingStepRunLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// StepRunStepTypeEnum Enum with underlying type: string
type StepRunStepTypeEnum string

// Set of constants representing the allowable values for StepRunStepTypeEnum
const (
	StepRunStepTypeMlJob StepRunStepTypeEnum = "ML_JOB"
)

var mappingStepRunStepTypeEnum = map[string]StepRunStepTypeEnum{
	"ML_JOB": StepRunStepTypeMlJob,
}

var mappingStepRunStepTypeEnumLowerCase = map[string]StepRunStepTypeEnum{
	"ml_job": StepRunStepTypeMlJob,
}

// GetStepRunStepTypeEnumValues Enumerates the set of values for StepRunStepTypeEnum
func GetStepRunStepTypeEnumValues() []StepRunStepTypeEnum {
	values := make([]StepRunStepTypeEnum, 0)
	for _, v := range mappingStepRunStepTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetStepRunStepTypeEnumStringValues Enumerates the set of values in String for StepRunStepTypeEnum
func GetStepRunStepTypeEnumStringValues() []string {
	return []string{
		"ML_JOB",
	}
}

// GetMappingStepRunStepTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingStepRunStepTypeEnum(val string) (StepRunStepTypeEnum, bool) {
	enum, ok := mappingStepRunStepTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
