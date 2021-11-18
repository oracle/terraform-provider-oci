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

// GithubFilter The filter for Github events
type GithubFilter struct {
	Include *GithubFilterAttributes `mandatory:"false" json:"include"`

	// The events, example PUSH, PULL_REQUEST_MERGE etc.
	Events []GithubFilterEventsEnum `mandatory:"false" json:"events,omitempty"`
}

func (m GithubFilter) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m GithubFilter) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeGithubFilter GithubFilter
	s := struct {
		DiscriminatorParam string `json:"triggerSource"`
		MarshalTypeGithubFilter
	}{
		"GITHUB",
		(MarshalTypeGithubFilter)(m),
	}

	return json.Marshal(&s)
}

// GithubFilterEventsEnum Enum with underlying type: string
type GithubFilterEventsEnum string

// Set of constants representing the allowable values for GithubFilterEventsEnum
const (
	GithubFilterEventsPush                GithubFilterEventsEnum = "PUSH"
	GithubFilterEventsPullRequestCreated  GithubFilterEventsEnum = "PULL_REQUEST_CREATED"
	GithubFilterEventsPullRequestUpdated  GithubFilterEventsEnum = "PULL_REQUEST_UPDATED"
	GithubFilterEventsPullRequestReopened GithubFilterEventsEnum = "PULL_REQUEST_REOPENED"
	GithubFilterEventsPullRequestMerged   GithubFilterEventsEnum = "PULL_REQUEST_MERGED"
)

var mappingGithubFilterEvents = map[string]GithubFilterEventsEnum{
	"PUSH":                  GithubFilterEventsPush,
	"PULL_REQUEST_CREATED":  GithubFilterEventsPullRequestCreated,
	"PULL_REQUEST_UPDATED":  GithubFilterEventsPullRequestUpdated,
	"PULL_REQUEST_REOPENED": GithubFilterEventsPullRequestReopened,
	"PULL_REQUEST_MERGED":   GithubFilterEventsPullRequestMerged,
}

// GetGithubFilterEventsEnumValues Enumerates the set of values for GithubFilterEventsEnum
func GetGithubFilterEventsEnumValues() []GithubFilterEventsEnum {
	values := make([]GithubFilterEventsEnum, 0)
	for _, v := range mappingGithubFilterEvents {
		values = append(values, v)
	}
	return values
}
