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
	"github.com/oracle/oci-go-sdk/v56/common"
)

// WaitStageRunProgress Specifies Wait stage specific run details.
type WaitStageRunProgress struct {

	// Build Run display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
	StageDisplayName *string `mandatory:"false" json:"stageDisplayName"`

	// The stage OCID.
	BuildPipelineStageId *string `mandatory:"false" json:"buildPipelineStageId"`

	// The time the stage started executing. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The time the stage finished executing. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`

	BuildPipelineStagePredecessors *BuildPipelineStagePredecessorCollection `mandatory:"false" json:"buildPipelineStagePredecessors"`

	// The current status of the stage.
	Status BuildPipelineStageRunProgressStatusEnum `mandatory:"false" json:"status,omitempty"`
}

//GetStageDisplayName returns StageDisplayName
func (m WaitStageRunProgress) GetStageDisplayName() *string {
	return m.StageDisplayName
}

//GetBuildPipelineStageId returns BuildPipelineStageId
func (m WaitStageRunProgress) GetBuildPipelineStageId() *string {
	return m.BuildPipelineStageId
}

//GetTimeStarted returns TimeStarted
func (m WaitStageRunProgress) GetTimeStarted() *common.SDKTime {
	return m.TimeStarted
}

//GetTimeFinished returns TimeFinished
func (m WaitStageRunProgress) GetTimeFinished() *common.SDKTime {
	return m.TimeFinished
}

//GetStatus returns Status
func (m WaitStageRunProgress) GetStatus() BuildPipelineStageRunProgressStatusEnum {
	return m.Status
}

//GetBuildPipelineStagePredecessors returns BuildPipelineStagePredecessors
func (m WaitStageRunProgress) GetBuildPipelineStagePredecessors() *BuildPipelineStagePredecessorCollection {
	return m.BuildPipelineStagePredecessors
}

func (m WaitStageRunProgress) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m WaitStageRunProgress) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeWaitStageRunProgress WaitStageRunProgress
	s := struct {
		DiscriminatorParam string `json:"buildPipelineStageType"`
		MarshalTypeWaitStageRunProgress
	}{
		"WAIT",
		(MarshalTypeWaitStageRunProgress)(m),
	}

	return json.Marshal(&s)
}
