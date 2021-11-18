// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps APIs to create a DevOps project to group the pipelines,  add reference to target deployment environments, add artifacts to deploy,  and create deployment pipelines needed to deploy your software.
//

package devops

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v52/common"
)

// RepositoryBranch Branch related information
type RepositoryBranch struct {

	// Unique Ref name inside a repository
	RefName *string `mandatory:"true" json:"refName"`

	// Unique full ref name inside a repository
	FullRefName *string `mandatory:"true" json:"fullRefName"`

	// The OCID of the repository containing the ref.
	RepositoryId *string `mandatory:"true" json:"repositoryId"`

	// Commit ID pointed to by the new branch.
	CommitId *string `mandatory:"true" json:"commitId"`
}

//GetRefName returns RefName
func (m RepositoryBranch) GetRefName() *string {
	return m.RefName
}

//GetFullRefName returns FullRefName
func (m RepositoryBranch) GetFullRefName() *string {
	return m.FullRefName
}

//GetRepositoryId returns RepositoryId
func (m RepositoryBranch) GetRepositoryId() *string {
	return m.RepositoryId
}

func (m RepositoryBranch) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m RepositoryBranch) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeRepositoryBranch RepositoryBranch
	s := struct {
		DiscriminatorParam string `json:"refType"`
		MarshalTypeRepositoryBranch
	}{
		"BRANCH",
		(MarshalTypeRepositoryBranch)(m),
	}

	return json.Marshal(&s)
}
