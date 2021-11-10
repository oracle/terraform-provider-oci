// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps APIs to create a DevOps project to group the pipelines,  add reference to target deployment environments, add artifacts to deploy,  and create deployment pipelines needed to deploy your software.
//

package devops

import (
	"github.com/oracle/oci-go-sdk/v51/common"
)

// CreateRepositoryDetails The information about new Repository.
type CreateRepositoryDetails struct {

	// Unique name of a repository.
	Name *string `mandatory:"true" json:"name"`

	// The OCID of the Project containing the repository.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// The default branch of the repository.
	DefaultBranch *string `mandatory:"false" json:"defaultBranch"`

	// Type of repository
	RepositoryType RepositoryRepositoryTypeEnum `mandatory:"false" json:"repositoryType,omitempty"`

	MirrorRepositoryConfig *MirrorRepositoryConfig `mandatory:"false" json:"mirrorRepositoryConfig"`

	// The description of this repository. Avoid entering confidential information
	Description *string `mandatory:"false" json:"description"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateRepositoryDetails) String() string {
	return common.PointerString(m)
}
