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

// DevopsCodeRepositoryBuildRunSource Specifies details of build run through Devops Code Repository.
type DevopsCodeRepositoryBuildRunSource struct {

	// The Trigger that invoked this build run
	TriggerId *string `mandatory:"true" json:"triggerId"`

	TriggerInfo *TriggerInfo `mandatory:"true" json:"triggerInfo"`

	// The Devops Code Repository RepoId that invoked this build run
	RepositoryId *string `mandatory:"true" json:"repositoryId"`
}

func (m DevopsCodeRepositoryBuildRunSource) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m DevopsCodeRepositoryBuildRunSource) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDevopsCodeRepositoryBuildRunSource DevopsCodeRepositoryBuildRunSource
	s := struct {
		DiscriminatorParam string `json:"sourceType"`
		MarshalTypeDevopsCodeRepositoryBuildRunSource
	}{
		"DEVOPS_CODE_REPOSITORY",
		(MarshalTypeDevopsCodeRepositoryBuildRunSource)(m),
	}

	return json.Marshal(&s)
}
