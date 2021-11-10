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

// RepositoryCommit Commit object with commit information
type RepositoryCommit struct {

	// Commit hash pointed to by Ref name
	CommitId *string `mandatory:"true" json:"commitId"`

	// The commit message.
	CommitMessage *string `mandatory:"true" json:"commitMessage"`

	// The name of the author of the repository.
	AuthorName *string `mandatory:"false" json:"authorName"`

	// The email of the author of the repository.
	AuthorEmail *string `mandatory:"false" json:"authorEmail"`

	// The name of who create the commit.
	CommitterName *string `mandatory:"false" json:"committerName"`

	// The email of who create the commit.
	CommitterEmail *string `mandatory:"false" json:"committerEmail"`

	// An array of parent commit ids of created commit.
	ParentCommitIds []string `mandatory:"false" json:"parentCommitIds"`

	// The time at which commit was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Tree information for the specified commit
	TreeId *string `mandatory:"false" json:"treeId"`
}

func (m RepositoryCommit) String() string {
	return common.PointerString(m)
}
