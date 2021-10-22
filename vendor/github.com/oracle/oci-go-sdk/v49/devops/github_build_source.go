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
	"github.com/oracle/oci-go-sdk/v49/common"
)

// GithubBuildSource Github Build Source for Build Stage
type GithubBuildSource struct {

	// Name of the Build source. This must be unique within a BuildSourceCollection. The name can be used by customers to locate the working directory pertinent to this repository.
	Name *string `mandatory:"true" json:"name"`

	// Url for repository
	RepositoryUrl *string `mandatory:"true" json:"repositoryUrl"`

	// branch name
	Branch *string `mandatory:"true" json:"branch"`

	// Connection identifier pertinent to GITHUB source provider
	ConnectionId *string `mandatory:"true" json:"connectionId"`
}

//GetName returns Name
func (m GithubBuildSource) GetName() *string {
	return m.Name
}

//GetRepositoryUrl returns RepositoryUrl
func (m GithubBuildSource) GetRepositoryUrl() *string {
	return m.RepositoryUrl
}

//GetBranch returns Branch
func (m GithubBuildSource) GetBranch() *string {
	return m.Branch
}

func (m GithubBuildSource) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m GithubBuildSource) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeGithubBuildSource GithubBuildSource
	s := struct {
		DiscriminatorParam string `json:"connectionType"`
		MarshalTypeGithubBuildSource
	}{
		"GITHUB",
		(MarshalTypeGithubBuildSource)(m),
	}

	return json.Marshal(&s)
}
