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

// BuildRun Each time you attempt to run a build pipeline you create one build run.
// A build can be running currently, or it can be a record of the run that happened in the past.
// The set of build runs constitutes a build pipeline's history.
type BuildRun struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	BuildRunSource BuildRunSource `mandatory:"true" json:"buildRunSource"`

	// Build run display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The OCID of the compartment where the build is running.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The OCID of the DevOps project.
	ProjectId *string `mandatory:"false" json:"projectId"`

	// The OCID of the build pipeline.
	BuildPipelineId *string `mandatory:"false" json:"buildPipelineId"`

	BuildRunArguments *BuildRunArgumentCollection `mandatory:"false" json:"buildRunArguments"`

	// The time the build run was created. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the build run was updated. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the build run.
	LifecycleState BuildRunLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	BuildRunProgress *BuildRunProgress `mandatory:"false" json:"buildRunProgress"`

	CommitInfo *CommitInfo `mandatory:"false" json:"commitInfo"`

	BuildOutputs *BuildOutputs `mandatory:"false" json:"buildOutputs"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m BuildRun) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BuildRun) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingBuildRunLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetBuildRunLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *BuildRun) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName       *string                           `json:"displayName"`
		CompartmentId     *string                           `json:"compartmentId"`
		ProjectId         *string                           `json:"projectId"`
		BuildPipelineId   *string                           `json:"buildPipelineId"`
		BuildRunArguments *BuildRunArgumentCollection       `json:"buildRunArguments"`
		TimeCreated       *common.SDKTime                   `json:"timeCreated"`
		TimeUpdated       *common.SDKTime                   `json:"timeUpdated"`
		LifecycleState    BuildRunLifecycleStateEnum        `json:"lifecycleState"`
		LifecycleDetails  *string                           `json:"lifecycleDetails"`
		BuildRunProgress  *BuildRunProgress                 `json:"buildRunProgress"`
		CommitInfo        *CommitInfo                       `json:"commitInfo"`
		BuildOutputs      *BuildOutputs                     `json:"buildOutputs"`
		FreeformTags      map[string]string                 `json:"freeformTags"`
		DefinedTags       map[string]map[string]interface{} `json:"definedTags"`
		SystemTags        map[string]map[string]interface{} `json:"systemTags"`
		Id                *string                           `json:"id"`
		BuildRunSource    buildrunsource                    `json:"buildRunSource"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.CompartmentId = model.CompartmentId

	m.ProjectId = model.ProjectId

	m.BuildPipelineId = model.BuildPipelineId

	m.BuildRunArguments = model.BuildRunArguments

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleState = model.LifecycleState

	m.LifecycleDetails = model.LifecycleDetails

	m.BuildRunProgress = model.BuildRunProgress

	m.CommitInfo = model.CommitInfo

	m.BuildOutputs = model.BuildOutputs

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

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

// BuildRunLifecycleStateEnum Enum with underlying type: string
type BuildRunLifecycleStateEnum string

// Set of constants representing the allowable values for BuildRunLifecycleStateEnum
const (
	BuildRunLifecycleStateAccepted   BuildRunLifecycleStateEnum = "ACCEPTED"
	BuildRunLifecycleStateInProgress BuildRunLifecycleStateEnum = "IN_PROGRESS"
	BuildRunLifecycleStateFailed     BuildRunLifecycleStateEnum = "FAILED"
	BuildRunLifecycleStateSucceeded  BuildRunLifecycleStateEnum = "SUCCEEDED"
	BuildRunLifecycleStateCanceling  BuildRunLifecycleStateEnum = "CANCELING"
	BuildRunLifecycleStateCanceled   BuildRunLifecycleStateEnum = "CANCELED"
	BuildRunLifecycleStateDeleting   BuildRunLifecycleStateEnum = "DELETING"
)

var mappingBuildRunLifecycleStateEnum = map[string]BuildRunLifecycleStateEnum{
	"ACCEPTED":    BuildRunLifecycleStateAccepted,
	"IN_PROGRESS": BuildRunLifecycleStateInProgress,
	"FAILED":      BuildRunLifecycleStateFailed,
	"SUCCEEDED":   BuildRunLifecycleStateSucceeded,
	"CANCELING":   BuildRunLifecycleStateCanceling,
	"CANCELED":    BuildRunLifecycleStateCanceled,
	"DELETING":    BuildRunLifecycleStateDeleting,
}

var mappingBuildRunLifecycleStateEnumLowerCase = map[string]BuildRunLifecycleStateEnum{
	"accepted":    BuildRunLifecycleStateAccepted,
	"in_progress": BuildRunLifecycleStateInProgress,
	"failed":      BuildRunLifecycleStateFailed,
	"succeeded":   BuildRunLifecycleStateSucceeded,
	"canceling":   BuildRunLifecycleStateCanceling,
	"canceled":    BuildRunLifecycleStateCanceled,
	"deleting":    BuildRunLifecycleStateDeleting,
}

// GetBuildRunLifecycleStateEnumValues Enumerates the set of values for BuildRunLifecycleStateEnum
func GetBuildRunLifecycleStateEnumValues() []BuildRunLifecycleStateEnum {
	values := make([]BuildRunLifecycleStateEnum, 0)
	for _, v := range mappingBuildRunLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetBuildRunLifecycleStateEnumStringValues Enumerates the set of values in String for BuildRunLifecycleStateEnum
func GetBuildRunLifecycleStateEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"FAILED",
		"SUCCEEDED",
		"CANCELING",
		"CANCELED",
		"DELETING",
	}
}

// GetMappingBuildRunLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBuildRunLifecycleStateEnum(val string) (BuildRunLifecycleStateEnum, bool) {
	enum, ok := mappingBuildRunLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
