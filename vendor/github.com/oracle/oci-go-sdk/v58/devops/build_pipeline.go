// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// BuildPipeline A set of stages forming a directed acyclic graph that defines the build process.
type BuildPipeline struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment where the build pipeline is created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the DevOps project.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// Optional description about the build pipeline.
	Description *string `mandatory:"false" json:"description"`

	// Build pipeline display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The time the build pipeline was created. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the build pipeline was updated. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the build pipeline.
	LifecycleState BuildPipelineLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	BuildPipelineParameters *BuildPipelineParameterCollection `mandatory:"false" json:"buildPipelineParameters"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m BuildPipeline) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BuildPipeline) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingBuildPipelineLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetBuildPipelineLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BuildPipelineLifecycleStateEnum Enum with underlying type: string
type BuildPipelineLifecycleStateEnum string

// Set of constants representing the allowable values for BuildPipelineLifecycleStateEnum
const (
	BuildPipelineLifecycleStateCreating BuildPipelineLifecycleStateEnum = "CREATING"
	BuildPipelineLifecycleStateUpdating BuildPipelineLifecycleStateEnum = "UPDATING"
	BuildPipelineLifecycleStateActive   BuildPipelineLifecycleStateEnum = "ACTIVE"
	BuildPipelineLifecycleStateInactive BuildPipelineLifecycleStateEnum = "INACTIVE"
	BuildPipelineLifecycleStateDeleting BuildPipelineLifecycleStateEnum = "DELETING"
	BuildPipelineLifecycleStateDeleted  BuildPipelineLifecycleStateEnum = "DELETED"
	BuildPipelineLifecycleStateFailed   BuildPipelineLifecycleStateEnum = "FAILED"
)

var mappingBuildPipelineLifecycleStateEnum = map[string]BuildPipelineLifecycleStateEnum{
	"CREATING": BuildPipelineLifecycleStateCreating,
	"UPDATING": BuildPipelineLifecycleStateUpdating,
	"ACTIVE":   BuildPipelineLifecycleStateActive,
	"INACTIVE": BuildPipelineLifecycleStateInactive,
	"DELETING": BuildPipelineLifecycleStateDeleting,
	"DELETED":  BuildPipelineLifecycleStateDeleted,
	"FAILED":   BuildPipelineLifecycleStateFailed,
}

// GetBuildPipelineLifecycleStateEnumValues Enumerates the set of values for BuildPipelineLifecycleStateEnum
func GetBuildPipelineLifecycleStateEnumValues() []BuildPipelineLifecycleStateEnum {
	values := make([]BuildPipelineLifecycleStateEnum, 0)
	for _, v := range mappingBuildPipelineLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetBuildPipelineLifecycleStateEnumStringValues Enumerates the set of values in String for BuildPipelineLifecycleStateEnum
func GetBuildPipelineLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingBuildPipelineLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBuildPipelineLifecycleStateEnum(val string) (BuildPipelineLifecycleStateEnum, bool) {
	mappingBuildPipelineLifecycleStateEnumIgnoreCase := make(map[string]BuildPipelineLifecycleStateEnum)
	for k, v := range mappingBuildPipelineLifecycleStateEnum {
		mappingBuildPipelineLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingBuildPipelineLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
