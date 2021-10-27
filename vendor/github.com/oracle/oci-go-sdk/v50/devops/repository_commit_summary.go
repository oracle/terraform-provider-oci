// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps APIs to create a DevOps project to group the pipelines,  add reference to target deployment environments, add artifacts to deploy,  and create deployment pipelines needed to deploy your software.
//

package devops

import (
	"github.com/oracle/oci-go-sdk/v50/common"
)

// RepositoryCommitSummary Commit summary with commit information
type RepositoryCommitSummary struct {

	// Commit hash pointed to by Ref name
	CommitId *string `mandatory:"true" json:"commitId"`

	// The commit message.
	CommitMessage *string `mandatory:"true" json:"commitMessage"`

	// The name of the author of the repository.
	AuthorName *string `mandatory:"true" json:"authorName"`

	// The email of the author of the repository.
	AuthorEmail *string `mandatory:"true" json:"authorEmail"`

	// The name of who create the commit.
	CommitterName *string `mandatory:"true" json:"committerName"`

	// The email of who create the commit.
	CommitterEmail *string `mandatory:"true" json:"committerEmail"`

	// An array of parent commit ids of created commit.
	ParentCommitIds []string `mandatory:"true" json:"parentCommitIds"`

	// The time to create the commit.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Tree information for the specified commit
	TreeId *string `mandatory:"true" json:"treeId"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m RepositoryCommitSummary) String() string {
	return common.PointerString(m)
}
