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

// CreateBuildPipelineStageDetails The information about a new stage.
type CreateBuildPipelineStageDetails interface {

	// The OCID of the build pipeline.
	GetBuildPipelineId() *string

	GetBuildPipelineStagePredecessorCollection() *BuildPipelineStagePredecessorCollection

	// Stage display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
	GetDisplayName() *string

	// Optional description about the stage.
	GetDescription() *string

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}
}

type createbuildpipelinestagedetails struct {
	JsonData                                []byte
	BuildPipelineId                         *string                                  `mandatory:"true" json:"buildPipelineId"`
	BuildPipelineStagePredecessorCollection *BuildPipelineStagePredecessorCollection `mandatory:"true" json:"buildPipelineStagePredecessorCollection"`
	DisplayName                             *string                                  `mandatory:"false" json:"displayName"`
	Description                             *string                                  `mandatory:"false" json:"description"`
	FreeformTags                            map[string]string                        `mandatory:"false" json:"freeformTags"`
	DefinedTags                             map[string]map[string]interface{}        `mandatory:"false" json:"definedTags"`
	BuildPipelineStageType                  string                                   `json:"buildPipelineStageType"`
}

// UnmarshalJSON unmarshals json
func (m *createbuildpipelinestagedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatebuildpipelinestagedetails createbuildpipelinestagedetails
	s := struct {
		Model Unmarshalercreatebuildpipelinestagedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.BuildPipelineId = s.Model.BuildPipelineId
	m.BuildPipelineStagePredecessorCollection = s.Model.BuildPipelineStagePredecessorCollection
	m.DisplayName = s.Model.DisplayName
	m.Description = s.Model.Description
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.BuildPipelineStageType = s.Model.BuildPipelineStageType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createbuildpipelinestagedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.BuildPipelineStageType {
	case "DELIVER_ARTIFACT":
		mm := CreateDeliverArtifactStageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TRIGGER_DEPLOYMENT_PIPELINE":
		mm := CreateTriggerDeploymentStageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "WAIT":
		mm := CreateWaitStageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "BUILD":
		mm := CreateBuildStageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetBuildPipelineId returns BuildPipelineId
func (m createbuildpipelinestagedetails) GetBuildPipelineId() *string {
	return m.BuildPipelineId
}

//GetBuildPipelineStagePredecessorCollection returns BuildPipelineStagePredecessorCollection
func (m createbuildpipelinestagedetails) GetBuildPipelineStagePredecessorCollection() *BuildPipelineStagePredecessorCollection {
	return m.BuildPipelineStagePredecessorCollection
}

//GetDisplayName returns DisplayName
func (m createbuildpipelinestagedetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetDescription returns Description
func (m createbuildpipelinestagedetails) GetDescription() *string {
	return m.Description
}

//GetFreeformTags returns FreeformTags
func (m createbuildpipelinestagedetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m createbuildpipelinestagedetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m createbuildpipelinestagedetails) String() string {
	return common.PointerString(m)
}
