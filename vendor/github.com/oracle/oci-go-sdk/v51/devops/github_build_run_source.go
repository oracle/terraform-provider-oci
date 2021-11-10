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
	"github.com/oracle/oci-go-sdk/v51/common"
)

// GithubBuildRunSource Specifies details of build run through Github.
type GithubBuildRunSource struct {

	// The Trigger that invoked this build run
	TriggerId *string `mandatory:"true" json:"triggerId"`

	TriggerInfo *TriggerInfo `mandatory:"true" json:"triggerInfo"`
}

func (m GithubBuildRunSource) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m GithubBuildRunSource) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeGithubBuildRunSource GithubBuildRunSource
	s := struct {
		DiscriminatorParam string `json:"sourceType"`
		MarshalTypeGithubBuildRunSource
	}{
		"GITHUB",
		(MarshalTypeGithubBuildRunSource)(m),
	}

	return json.Marshal(&s)
}
