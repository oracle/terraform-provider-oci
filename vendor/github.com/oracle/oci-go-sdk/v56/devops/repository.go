// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// Repository Repositories containing the source code to build and deploy.
type Repository struct {

	// The OCID of the repository. This value is unique and immutable.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the repository's compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the DevOps project containing the repository.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// Unique name of a repository. This value is mutable.
	Name *string `mandatory:"false" json:"name"`

	// Tenancy unique namespace.
	Namespace *string `mandatory:"false" json:"namespace"`

	// Unique project name in a namespace.
	ProjectName *string `mandatory:"false" json:"projectName"`

	// SSH URL that you use to git clone, pull and push.
	SshUrl *string `mandatory:"false" json:"sshUrl"`

	// HTTP URL that you use to git clone, pull and push.
	HttpUrl *string `mandatory:"false" json:"httpUrl"`

	// Details of the repository. Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// The default branch of the repository.
	DefaultBranch *string `mandatory:"false" json:"defaultBranch"`

	// Type of repository:
	// Mirrored - Repository created by mirroring an existing repository.
	// Hosted - Repository created and hosted using OCI DevOps code repository.
	RepositoryType RepositoryRepositoryTypeEnum `mandatory:"false" json:"repositoryType,omitempty"`

	MirrorRepositoryConfig *MirrorRepositoryConfig `mandatory:"false" json:"mirrorRepositoryConfig"`

	// The time the repository was created. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the repository was updated. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the repository.
	LifecycleState RepositoryLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecyleDetails *string `mandatory:"false" json:"lifecyleDetails"`

	// The count of the branches present in the repository.
	BranchCount *int `mandatory:"false" json:"branchCount"`

	// The count of the commits present in the repository.
	CommitCount *int `mandatory:"false" json:"commitCount"`

	// The size of the repository in bytes.
	SizeInBytes *int64 `mandatory:"false" json:"sizeInBytes"`

	// Trigger build events supported for this repository:
	// Push - Build is triggered when a push event occurs.
	// Commit updates - Build is triggered when new commits are mirrored into a repository.
	TriggerBuildEvents []RepositoryTriggerBuildEventsEnum `mandatory:"false" json:"triggerBuildEvents,omitempty"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m Repository) String() string {
	return common.PointerString(m)
}

// RepositoryRepositoryTypeEnum Enum with underlying type: string
type RepositoryRepositoryTypeEnum string

// Set of constants representing the allowable values for RepositoryRepositoryTypeEnum
const (
	RepositoryRepositoryTypeMirrored RepositoryRepositoryTypeEnum = "MIRRORED"
	RepositoryRepositoryTypeHosted   RepositoryRepositoryTypeEnum = "HOSTED"
)

var mappingRepositoryRepositoryType = map[string]RepositoryRepositoryTypeEnum{
	"MIRRORED": RepositoryRepositoryTypeMirrored,
	"HOSTED":   RepositoryRepositoryTypeHosted,
}

// GetRepositoryRepositoryTypeEnumValues Enumerates the set of values for RepositoryRepositoryTypeEnum
func GetRepositoryRepositoryTypeEnumValues() []RepositoryRepositoryTypeEnum {
	values := make([]RepositoryRepositoryTypeEnum, 0)
	for _, v := range mappingRepositoryRepositoryType {
		values = append(values, v)
	}
	return values
}

// RepositoryLifecycleStateEnum Enum with underlying type: string
type RepositoryLifecycleStateEnum string

// Set of constants representing the allowable values for RepositoryLifecycleStateEnum
const (
	RepositoryLifecycleStateActive   RepositoryLifecycleStateEnum = "ACTIVE"
	RepositoryLifecycleStateCreating RepositoryLifecycleStateEnum = "CREATING"
	RepositoryLifecycleStateDeleted  RepositoryLifecycleStateEnum = "DELETED"
)

var mappingRepositoryLifecycleState = map[string]RepositoryLifecycleStateEnum{
	"ACTIVE":   RepositoryLifecycleStateActive,
	"CREATING": RepositoryLifecycleStateCreating,
	"DELETED":  RepositoryLifecycleStateDeleted,
}

// GetRepositoryLifecycleStateEnumValues Enumerates the set of values for RepositoryLifecycleStateEnum
func GetRepositoryLifecycleStateEnumValues() []RepositoryLifecycleStateEnum {
	values := make([]RepositoryLifecycleStateEnum, 0)
	for _, v := range mappingRepositoryLifecycleState {
		values = append(values, v)
	}
	return values
}

// RepositoryTriggerBuildEventsEnum Enum with underlying type: string
type RepositoryTriggerBuildEventsEnum string

// Set of constants representing the allowable values for RepositoryTriggerBuildEventsEnum
const (
	RepositoryTriggerBuildEventsPush          RepositoryTriggerBuildEventsEnum = "PUSH"
	RepositoryTriggerBuildEventsCommitUpdates RepositoryTriggerBuildEventsEnum = "COMMIT_UPDATES"
)

var mappingRepositoryTriggerBuildEvents = map[string]RepositoryTriggerBuildEventsEnum{
	"PUSH":           RepositoryTriggerBuildEventsPush,
	"COMMIT_UPDATES": RepositoryTriggerBuildEventsCommitUpdates,
}

// GetRepositoryTriggerBuildEventsEnumValues Enumerates the set of values for RepositoryTriggerBuildEventsEnum
func GetRepositoryTriggerBuildEventsEnumValues() []RepositoryTriggerBuildEventsEnum {
	values := make([]RepositoryTriggerBuildEventsEnum, 0)
	for _, v := range mappingRepositoryTriggerBuildEvents {
		values = append(values, v)
	}
	return values
}
