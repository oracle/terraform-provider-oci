// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps APIs to create a DevOps project to group the pipelines,  add reference to target deployment environments, add artifacts to deploy,  and create deployment pipelines needed to deploy your software.
//

package devops

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v52/common"
)

// BuildRunProgress The run progress details of a BuildRun.
type BuildRunProgress struct {

	// The time the the BuildRun is started. An RFC3339 formatted datetime string
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The time the BuildRun is finished. An RFC3339 formatted datetime string
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`

	// Map of stage OCIDs to BuildPipelineStageRunProgress model.
	BuildPipelineStageRunProgress map[string]BuildPipelineStageRunProgress `mandatory:"false" json:"buildPipelineStageRunProgress"`
}

func (m BuildRunProgress) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *BuildRunProgress) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		TimeStarted                   *common.SDKTime                          `json:"timeStarted"`
		TimeFinished                  *common.SDKTime                          `json:"timeFinished"`
		BuildPipelineStageRunProgress map[string]buildpipelinestagerunprogress `json:"buildPipelineStageRunProgress"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.TimeStarted = model.TimeStarted

	m.TimeFinished = model.TimeFinished

	m.BuildPipelineStageRunProgress = make(map[string]BuildPipelineStageRunProgress)
	for k, v := range model.BuildPipelineStageRunProgress {
		nn, e = v.UnmarshalPolymorphicJSON(v.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.BuildPipelineStageRunProgress[k] = nn.(BuildPipelineStageRunProgress)
		} else {
			m.BuildPipelineStageRunProgress[k] = nil
		}
	}

	return
}
