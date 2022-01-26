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

// RepositoryCommit Commit object with commit information.
type RepositoryCommit struct {

	// Commit hash pointed to by reference name.
	CommitId *string `mandatory:"true" json:"commitId"`

	// The commit message.
	CommitMessage *string `mandatory:"true" json:"commitMessage"`

	// Name of the author of the repository.
	AuthorName *string `mandatory:"false" json:"authorName"`

	// Email of the author of the repository.
	AuthorEmail *string `mandatory:"false" json:"authorEmail"`

	// Name of who creates the commit.
	CommitterName *string `mandatory:"false" json:"committerName"`

	// Email of who creates the commit.
	CommitterEmail *string `mandatory:"false" json:"committerEmail"`

	// An array of parent commit IDs of created commit.
	ParentCommitIds []string `mandatory:"false" json:"parentCommitIds"`

	// The time at which commit was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Tree information for the specified commit.
	TreeId *string `mandatory:"false" json:"treeId"`
}

func (m RepositoryCommit) String() string {
	return common.PointerString(m)
}
