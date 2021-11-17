// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps APIs to create a DevOps project to group the pipelines,  add reference to target deployment environments, add artifacts to deploy,  and create deployment pipelines needed to deploy your software.
//

package devops

import (
	"github.com/oracle/oci-go-sdk/v52/common"
)

// RepositorySummary Summary of the Repository.
type RepositorySummary struct {

	// The OCID of the repository. This value is unique and immutable.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the repository's Compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the Project containing the repository.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// Unique name of a repository. This value is mutable.
	Name *string `mandatory:"false" json:"name"`

	// Tenancy unique namespace. This value is fetched from Project Service
	Namespace *string `mandatory:"false" json:"namespace"`

	// Unique name of a project in a namespace. This value is fetched from Project Service
	ProjectName *string `mandatory:"false" json:"projectName"`

	// The description of this repository. Avoid entering confidential information
	Description *string `mandatory:"false" json:"description"`

	// The default branch of the repository
	DefaultBranch *string `mandatory:"false" json:"defaultBranch"`

	// Type of repository
	RepositoryType RepositoryRepositoryTypeEnum `mandatory:"false" json:"repositoryType,omitempty"`

	// ssh url user utilized to git clone, pull and push
	SshUrl *string `mandatory:"false" json:"sshUrl"`

	// http url user utilized to git clone, pull and push
	HttpUrl *string `mandatory:"false" json:"httpUrl"`

	MirrorRepositoryConfig *MirrorRepositoryConfig `mandatory:"false" json:"mirrorRepositoryConfig"`

	// The time the the Repository was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the Repository was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the Repository.
	LifecycleState RepositoryLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m RepositorySummary) String() string {
	return common.PointerString(m)
}
