// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BuildRunProgress The run progress details of a build run.
type BuildRunProgress struct {

	// The time the build run started. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The time the build run finished. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`

	// Map of stage OCIDs to build pipeline stage run progress model.
	BuildPipelineStageRunProgress map[string]BuildPipelineStageRunProgress `mandatory:"false" json:"buildPipelineStageRunProgress"`
}

func (m BuildRunProgress) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BuildRunProgress) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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
