// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.oracle.com/iaas/Content/devops/using/home.htm).
//

package devops

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DeliverArtifactStage Specifies the Deliver Artifacts stage.
type DeliverArtifactStage struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the DevOps project.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// The OCID of the build pipeline.
	BuildPipelineId *string `mandatory:"true" json:"buildPipelineId"`

	// The OCID of the compartment where the pipeline is created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	DeliverArtifactCollection *DeliverArtifactCollection `mandatory:"true" json:"deliverArtifactCollection"`

	// Stage display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Optional description about the build stage.
	Description *string `mandatory:"false" json:"description"`

	// The time the stage was created. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the stage was updated. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	BuildPipelineStagePredecessorCollection *BuildPipelineStagePredecessorCollection `mandatory:"false" json:"buildPipelineStagePredecessorCollection"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces. See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The current state of the stage.
	LifecycleState BuildPipelineStageLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

// GetId returns Id
func (m DeliverArtifactStage) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m DeliverArtifactStage) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m DeliverArtifactStage) GetDescription() *string {
	return m.Description
}

// GetProjectId returns ProjectId
func (m DeliverArtifactStage) GetProjectId() *string {
	return m.ProjectId
}

// GetBuildPipelineId returns BuildPipelineId
func (m DeliverArtifactStage) GetBuildPipelineId() *string {
	return m.BuildPipelineId
}

// GetCompartmentId returns CompartmentId
func (m DeliverArtifactStage) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetTimeCreated returns TimeCreated
func (m DeliverArtifactStage) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m DeliverArtifactStage) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleState returns LifecycleState
func (m DeliverArtifactStage) GetLifecycleState() BuildPipelineStageLifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m DeliverArtifactStage) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetBuildPipelineStagePredecessorCollection returns BuildPipelineStagePredecessorCollection
func (m DeliverArtifactStage) GetBuildPipelineStagePredecessorCollection() *BuildPipelineStagePredecessorCollection {
	return m.BuildPipelineStagePredecessorCollection
}

// GetFreeformTags returns FreeformTags
func (m DeliverArtifactStage) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m DeliverArtifactStage) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m DeliverArtifactStage) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m DeliverArtifactStage) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DeliverArtifactStage) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingBuildPipelineStageLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetBuildPipelineStageLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DeliverArtifactStage) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDeliverArtifactStage DeliverArtifactStage
	s := struct {
		DiscriminatorParam string `json:"buildPipelineStageType"`
		MarshalTypeDeliverArtifactStage
	}{
		"DELIVER_ARTIFACT",
		(MarshalTypeDeliverArtifactStage)(m),
	}

	return json.Marshal(&s)
}
