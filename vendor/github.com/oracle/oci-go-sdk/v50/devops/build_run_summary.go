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
	"github.com/oracle/oci-go-sdk/v50/common"
)

// BuildRunSummary Summary of the BuildRun.
type BuildRunSummary struct {

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Project Identifier
	ProjectId *string `mandatory:"true" json:"projectId"`

	// Pipeline Identifier
	BuildPipelineId *string `mandatory:"true" json:"buildPipelineId"`

	BuildRunSource BuildRunSource `mandatory:"true" json:"buildRunSource"`

	// BuildRun identifier which can be renamed and is not necessarily unique
	DisplayName *string `mandatory:"false" json:"displayName"`

	BuildRunArguments *BuildRunArgumentCollection `mandatory:"false" json:"buildRunArguments"`

	BuildRunProgressSummary *BuildRunProgressSummary `mandatory:"false" json:"buildRunProgressSummary"`

	// The time the the BuildRun was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the BuildRun was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the BuildRun.
	LifecycleState BuildRunLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	CommitInfo *CommitInfo `mandatory:"false" json:"commitInfo"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m BuildRunSummary) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *BuildRunSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName             *string                           `json:"displayName"`
		BuildRunArguments       *BuildRunArgumentCollection       `json:"buildRunArguments"`
		BuildRunProgressSummary *BuildRunProgressSummary          `json:"buildRunProgressSummary"`
		TimeCreated             *common.SDKTime                   `json:"timeCreated"`
		TimeUpdated             *common.SDKTime                   `json:"timeUpdated"`
		LifecycleState          BuildRunLifecycleStateEnum        `json:"lifecycleState"`
		LifecycleDetails        *string                           `json:"lifecycleDetails"`
		CommitInfo              *CommitInfo                       `json:"commitInfo"`
		FreeformTags            map[string]string                 `json:"freeformTags"`
		DefinedTags             map[string]map[string]interface{} `json:"definedTags"`
		SystemTags              map[string]map[string]interface{} `json:"systemTags"`
		Id                      *string                           `json:"id"`
		CompartmentId           *string                           `json:"compartmentId"`
		ProjectId               *string                           `json:"projectId"`
		BuildPipelineId         *string                           `json:"buildPipelineId"`
		BuildRunSource          buildrunsource                    `json:"buildRunSource"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.BuildRunArguments = model.BuildRunArguments

	m.BuildRunProgressSummary = model.BuildRunProgressSummary

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleState = model.LifecycleState

	m.LifecycleDetails = model.LifecycleDetails

	m.CommitInfo = model.CommitInfo

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.ProjectId = model.ProjectId

	m.BuildPipelineId = model.BuildPipelineId

	nn, e = model.BuildRunSource.UnmarshalPolymorphicJSON(model.BuildRunSource.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.BuildRunSource = nn.(BuildRunSource)
	} else {
		m.BuildRunSource = nil
	}

	return
}
