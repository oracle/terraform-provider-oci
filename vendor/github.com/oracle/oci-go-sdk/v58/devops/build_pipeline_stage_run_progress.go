// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// BuildPipelineStageRunProgress The details about the run progress of a stage in a build run.
type BuildPipelineStageRunProgress interface {

	// Build Run display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
	GetStageDisplayName() *string

	// The stage OCID.
	GetBuildPipelineStageId() *string

	// The time the stage started executing. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	GetTimeStarted() *common.SDKTime

	// The time the stage finished executing. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	GetTimeFinished() *common.SDKTime

	// The current status of the stage.
	GetStatus() BuildPipelineStageRunProgressStatusEnum

	GetBuildPipelineStagePredecessors() *BuildPipelineStagePredecessorCollection
}

type buildpipelinestagerunprogress struct {
	JsonData                       []byte
	StageDisplayName               *string                                  `mandatory:"false" json:"stageDisplayName"`
	BuildPipelineStageId           *string                                  `mandatory:"false" json:"buildPipelineStageId"`
	TimeStarted                    *common.SDKTime                          `mandatory:"false" json:"timeStarted"`
	TimeFinished                   *common.SDKTime                          `mandatory:"false" json:"timeFinished"`
	Status                         BuildPipelineStageRunProgressStatusEnum  `mandatory:"false" json:"status,omitempty"`
	BuildPipelineStagePredecessors *BuildPipelineStagePredecessorCollection `mandatory:"false" json:"buildPipelineStagePredecessors"`
	BuildPipelineStageType         string                                   `json:"buildPipelineStageType"`
}

// UnmarshalJSON unmarshals json
func (m *buildpipelinestagerunprogress) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerbuildpipelinestagerunprogress buildpipelinestagerunprogress
	s := struct {
		Model Unmarshalerbuildpipelinestagerunprogress
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.StageDisplayName = s.Model.StageDisplayName
	m.BuildPipelineStageId = s.Model.BuildPipelineStageId
	m.TimeStarted = s.Model.TimeStarted
	m.TimeFinished = s.Model.TimeFinished
	m.Status = s.Model.Status
	m.BuildPipelineStagePredecessors = s.Model.BuildPipelineStagePredecessors
	m.BuildPipelineStageType = s.Model.BuildPipelineStageType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *buildpipelinestagerunprogress) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.BuildPipelineStageType {
	case "DELIVER_ARTIFACT":
		mm := DeliverArtifactStageRunProgress{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "WAIT":
		mm := WaitStageRunProgress{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TRIGGER_DEPLOYMENT_PIPELINE":
		mm := TriggerDeploymentPipelineStageRunProgress{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "BUILD":
		mm := BuildStageRunProgress{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetStageDisplayName returns StageDisplayName
func (m buildpipelinestagerunprogress) GetStageDisplayName() *string {
	return m.StageDisplayName
}

//GetBuildPipelineStageId returns BuildPipelineStageId
func (m buildpipelinestagerunprogress) GetBuildPipelineStageId() *string {
	return m.BuildPipelineStageId
}

//GetTimeStarted returns TimeStarted
func (m buildpipelinestagerunprogress) GetTimeStarted() *common.SDKTime {
	return m.TimeStarted
}

//GetTimeFinished returns TimeFinished
func (m buildpipelinestagerunprogress) GetTimeFinished() *common.SDKTime {
	return m.TimeFinished
}

//GetStatus returns Status
func (m buildpipelinestagerunprogress) GetStatus() BuildPipelineStageRunProgressStatusEnum {
	return m.Status
}

//GetBuildPipelineStagePredecessors returns BuildPipelineStagePredecessors
func (m buildpipelinestagerunprogress) GetBuildPipelineStagePredecessors() *BuildPipelineStagePredecessorCollection {
	return m.BuildPipelineStagePredecessors
}

func (m buildpipelinestagerunprogress) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m buildpipelinestagerunprogress) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingBuildPipelineStageRunProgressStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetBuildPipelineStageRunProgressStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BuildPipelineStageRunProgressStatusEnum Enum with underlying type: string
type BuildPipelineStageRunProgressStatusEnum string

// Set of constants representing the allowable values for BuildPipelineStageRunProgressStatusEnum
const (
	BuildPipelineStageRunProgressStatusAccepted   BuildPipelineStageRunProgressStatusEnum = "ACCEPTED"
	BuildPipelineStageRunProgressStatusInProgress BuildPipelineStageRunProgressStatusEnum = "IN_PROGRESS"
	BuildPipelineStageRunProgressStatusFailed     BuildPipelineStageRunProgressStatusEnum = "FAILED"
	BuildPipelineStageRunProgressStatusSucceeded  BuildPipelineStageRunProgressStatusEnum = "SUCCEEDED"
	BuildPipelineStageRunProgressStatusCanceling  BuildPipelineStageRunProgressStatusEnum = "CANCELING"
	BuildPipelineStageRunProgressStatusCanceled   BuildPipelineStageRunProgressStatusEnum = "CANCELED"
)

var mappingBuildPipelineStageRunProgressStatusEnum = map[string]BuildPipelineStageRunProgressStatusEnum{
	"ACCEPTED":    BuildPipelineStageRunProgressStatusAccepted,
	"IN_PROGRESS": BuildPipelineStageRunProgressStatusInProgress,
	"FAILED":      BuildPipelineStageRunProgressStatusFailed,
	"SUCCEEDED":   BuildPipelineStageRunProgressStatusSucceeded,
	"CANCELING":   BuildPipelineStageRunProgressStatusCanceling,
	"CANCELED":    BuildPipelineStageRunProgressStatusCanceled,
}

// GetBuildPipelineStageRunProgressStatusEnumValues Enumerates the set of values for BuildPipelineStageRunProgressStatusEnum
func GetBuildPipelineStageRunProgressStatusEnumValues() []BuildPipelineStageRunProgressStatusEnum {
	values := make([]BuildPipelineStageRunProgressStatusEnum, 0)
	for _, v := range mappingBuildPipelineStageRunProgressStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetBuildPipelineStageRunProgressStatusEnumStringValues Enumerates the set of values in String for BuildPipelineStageRunProgressStatusEnum
func GetBuildPipelineStageRunProgressStatusEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"FAILED",
		"SUCCEEDED",
		"CANCELING",
		"CANCELED",
	}
}

// GetMappingBuildPipelineStageRunProgressStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBuildPipelineStageRunProgressStatusEnum(val string) (BuildPipelineStageRunProgressStatusEnum, bool) {
	mappingBuildPipelineStageRunProgressStatusEnumIgnoreCase := make(map[string]BuildPipelineStageRunProgressStatusEnum)
	for k, v := range mappingBuildPipelineStageRunProgressStatusEnum {
		mappingBuildPipelineStageRunProgressStatusEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingBuildPipelineStageRunProgressStatusEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
