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

// BuildStageRunProgress Specifies the run details for Build stage.
type BuildStageRunProgress struct {
	BuildSourceCollection *BuildSourceCollection `mandatory:"true" json:"buildSourceCollection"`

	// Build Run display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
	StageDisplayName *string `mandatory:"false" json:"stageDisplayName"`

	// The stage OCID.
	BuildPipelineStageId *string `mandatory:"false" json:"buildPipelineStageId"`

	// The time the stage started executing. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The time the stage finished executing. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`

	BuildPipelineStagePredecessors *BuildPipelineStagePredecessorCollection `mandatory:"false" json:"buildPipelineStagePredecessors"`

	// Name of Build Runner shape where this Build Stage is running.
	ActualBuildRunnerShape *string `mandatory:"false" json:"actualBuildRunnerShape"`

	ActualBuildRunnerShapeConfig *ActualBuildRunnerShapeConfig `mandatory:"false" json:"actualBuildRunnerShapeConfig"`

	// The path to the build specification file for this Environment. The default location if not specified is build_spec.yaml
	BuildSpecFile *string `mandatory:"false" json:"buildSpecFile"`

	// Timeout for the Build Stage Execution. Value in seconds.
	StageExecutionTimeoutInSeconds *int `mandatory:"false" json:"stageExecutionTimeoutInSeconds"`

	// Name of the BuildSource in which the build_spec.yml file need to be located. If not specified, the 1st entry in the BuildSource collection will be chosen as Primary.
	PrimaryBuildSource *string `mandatory:"false" json:"primaryBuildSource"`

	// The details about all the steps in a Build stage
	Steps []BuildStageRunStep `mandatory:"false" json:"steps"`

	ExportedVariables *ExportedVariableCollection `mandatory:"false" json:"exportedVariables"`

	// Image name for the Build Environment
	Image BuildStageRunProgressImageEnum `mandatory:"true" json:"image"`

	// The current status of the stage.
	Status BuildPipelineStageRunProgressStatusEnum `mandatory:"false" json:"status,omitempty"`
}

//GetStageDisplayName returns StageDisplayName
func (m BuildStageRunProgress) GetStageDisplayName() *string {
	return m.StageDisplayName
}

//GetBuildPipelineStageId returns BuildPipelineStageId
func (m BuildStageRunProgress) GetBuildPipelineStageId() *string {
	return m.BuildPipelineStageId
}

//GetTimeStarted returns TimeStarted
func (m BuildStageRunProgress) GetTimeStarted() *common.SDKTime {
	return m.TimeStarted
}

//GetTimeFinished returns TimeFinished
func (m BuildStageRunProgress) GetTimeFinished() *common.SDKTime {
	return m.TimeFinished
}

//GetStatus returns Status
func (m BuildStageRunProgress) GetStatus() BuildPipelineStageRunProgressStatusEnum {
	return m.Status
}

//GetBuildPipelineStagePredecessors returns BuildPipelineStagePredecessors
func (m BuildStageRunProgress) GetBuildPipelineStagePredecessors() *BuildPipelineStagePredecessorCollection {
	return m.BuildPipelineStagePredecessors
}

func (m BuildStageRunProgress) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BuildStageRunProgress) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBuildStageRunProgressImageEnum(string(m.Image)); !ok && m.Image != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Image: %s. Supported values are: %s.", m.Image, strings.Join(GetBuildStageRunProgressImageEnumStringValues(), ",")))
	}

	if _, ok := GetMappingBuildPipelineStageRunProgressStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetBuildPipelineStageRunProgressStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m BuildStageRunProgress) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeBuildStageRunProgress BuildStageRunProgress
	s := struct {
		DiscriminatorParam string `json:"buildPipelineStageType"`
		MarshalTypeBuildStageRunProgress
	}{
		"BUILD",
		(MarshalTypeBuildStageRunProgress)(m),
	}

	return json.Marshal(&s)
}

// BuildStageRunProgressImageEnum Enum with underlying type: string
type BuildStageRunProgressImageEnum string

// Set of constants representing the allowable values for BuildStageRunProgressImageEnum
const (
	BuildStageRunProgressImageOl7X8664Standard10 BuildStageRunProgressImageEnum = "OL7_X86_64_STANDARD_10"
)

var mappingBuildStageRunProgressImageEnum = map[string]BuildStageRunProgressImageEnum{
	"OL7_X86_64_STANDARD_10": BuildStageRunProgressImageOl7X8664Standard10,
}

// GetBuildStageRunProgressImageEnumValues Enumerates the set of values for BuildStageRunProgressImageEnum
func GetBuildStageRunProgressImageEnumValues() []BuildStageRunProgressImageEnum {
	values := make([]BuildStageRunProgressImageEnum, 0)
	for _, v := range mappingBuildStageRunProgressImageEnum {
		values = append(values, v)
	}
	return values
}

// GetBuildStageRunProgressImageEnumStringValues Enumerates the set of values in String for BuildStageRunProgressImageEnum
func GetBuildStageRunProgressImageEnumStringValues() []string {
	return []string{
		"OL7_X86_64_STANDARD_10",
	}
}

// GetMappingBuildStageRunProgressImageEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBuildStageRunProgressImageEnum(val string) (BuildStageRunProgressImageEnum, bool) {
	mappingBuildStageRunProgressImageEnumIgnoreCase := make(map[string]BuildStageRunProgressImageEnum)
	for k, v := range mappingBuildStageRunProgressImageEnum {
		mappingBuildStageRunProgressImageEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingBuildStageRunProgressImageEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
