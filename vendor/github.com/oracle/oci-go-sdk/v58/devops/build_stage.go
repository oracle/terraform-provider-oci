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

// BuildStage Specifies the build stage.
type BuildStage struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the DevOps project.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// The OCID of the build pipeline.
	BuildPipelineId *string `mandatory:"true" json:"buildPipelineId"`

	// The OCID of the compartment where the pipeline is created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	BuildSourceCollection *BuildSourceCollection `mandatory:"true" json:"buildSourceCollection"`

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

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The path to the build specification file for this environment. The default location of the file if not specified is build_spec.yaml.
	BuildSpecFile *string `mandatory:"false" json:"buildSpecFile"`

	// Timeout for the build stage execution. Specify value in seconds.
	StageExecutionTimeoutInSeconds *int `mandatory:"false" json:"stageExecutionTimeoutInSeconds"`

	// Name of the build source where the build_spec.yml file is located. If not specified, then the first entry in the build source collection is chosen as primary build source.
	PrimaryBuildSource *string `mandatory:"false" json:"primaryBuildSource"`

	// Image name for the build environment.
	Image BuildStageImageEnum `mandatory:"true" json:"image"`

	// The current state of the stage.
	LifecycleState BuildPipelineStageLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

//GetId returns Id
func (m BuildStage) GetId() *string {
	return m.Id
}

//GetDisplayName returns DisplayName
func (m BuildStage) GetDisplayName() *string {
	return m.DisplayName
}

//GetDescription returns Description
func (m BuildStage) GetDescription() *string {
	return m.Description
}

//GetProjectId returns ProjectId
func (m BuildStage) GetProjectId() *string {
	return m.ProjectId
}

//GetBuildPipelineId returns BuildPipelineId
func (m BuildStage) GetBuildPipelineId() *string {
	return m.BuildPipelineId
}

//GetCompartmentId returns CompartmentId
func (m BuildStage) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetTimeCreated returns TimeCreated
func (m BuildStage) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetTimeUpdated returns TimeUpdated
func (m BuildStage) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

//GetLifecycleState returns LifecycleState
func (m BuildStage) GetLifecycleState() BuildPipelineStageLifecycleStateEnum {
	return m.LifecycleState
}

//GetLifecycleDetails returns LifecycleDetails
func (m BuildStage) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

//GetBuildPipelineStagePredecessorCollection returns BuildPipelineStagePredecessorCollection
func (m BuildStage) GetBuildPipelineStagePredecessorCollection() *BuildPipelineStagePredecessorCollection {
	return m.BuildPipelineStagePredecessorCollection
}

//GetFreeformTags returns FreeformTags
func (m BuildStage) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m BuildStage) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetSystemTags returns SystemTags
func (m BuildStage) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m BuildStage) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BuildStage) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBuildStageImageEnum(string(m.Image)); !ok && m.Image != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Image: %s. Supported values are: %s.", m.Image, strings.Join(GetBuildStageImageEnumStringValues(), ",")))
	}

	if _, ok := GetMappingBuildPipelineStageLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetBuildPipelineStageLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m BuildStage) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeBuildStage BuildStage
	s := struct {
		DiscriminatorParam string `json:"buildPipelineStageType"`
		MarshalTypeBuildStage
	}{
		"BUILD",
		(MarshalTypeBuildStage)(m),
	}

	return json.Marshal(&s)
}

// BuildStageImageEnum Enum with underlying type: string
type BuildStageImageEnum string

// Set of constants representing the allowable values for BuildStageImageEnum
const (
	BuildStageImageOl7X8664Standard10 BuildStageImageEnum = "OL7_X86_64_STANDARD_10"
)

var mappingBuildStageImageEnum = map[string]BuildStageImageEnum{
	"OL7_X86_64_STANDARD_10": BuildStageImageOl7X8664Standard10,
}

// GetBuildStageImageEnumValues Enumerates the set of values for BuildStageImageEnum
func GetBuildStageImageEnumValues() []BuildStageImageEnum {
	values := make([]BuildStageImageEnum, 0)
	for _, v := range mappingBuildStageImageEnum {
		values = append(values, v)
	}
	return values
}

// GetBuildStageImageEnumStringValues Enumerates the set of values in String for BuildStageImageEnum
func GetBuildStageImageEnumStringValues() []string {
	return []string{
		"OL7_X86_64_STANDARD_10",
	}
}

// GetMappingBuildStageImageEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBuildStageImageEnum(val string) (BuildStageImageEnum, bool) {
	mappingBuildStageImageEnumIgnoreCase := make(map[string]BuildStageImageEnum)
	for k, v := range mappingBuildStageImageEnum {
		mappingBuildStageImageEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingBuildStageImageEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
