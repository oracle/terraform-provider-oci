// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// GitlabServerFilter The filter for GitLab self-hosted events.
type GitlabServerFilter struct {
	Include *GitlabServerFilterAttributes `mandatory:"false" json:"include"`

	Exclude *GitlabServerFilterExclusionAttributes `mandatory:"false" json:"exclude"`

	// The events, for example, PUSH, PULL_REQUEST_MERGE.
	Events []GitlabServerFilterEventsEnum `mandatory:"false" json:"events,omitempty"`
}

func (m GitlabServerFilter) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GitlabServerFilter) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range m.Events {
		if _, ok := GetMappingGitlabServerFilterEventsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Events: %s. Supported values are: %s.", val, strings.Join(GetGitlabServerFilterEventsEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m GitlabServerFilter) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeGitlabServerFilter GitlabServerFilter
	s := struct {
		DiscriminatorParam string `json:"triggerSource"`
		MarshalTypeGitlabServerFilter
	}{
		"GITLAB_SERVER",
		(MarshalTypeGitlabServerFilter)(m),
	}

	return json.Marshal(&s)
}

// GitlabServerFilterEventsEnum Enum with underlying type: string
type GitlabServerFilterEventsEnum string

// Set of constants representing the allowable values for GitlabServerFilterEventsEnum
const (
	GitlabServerFilterEventsPush                GitlabServerFilterEventsEnum = "PUSH"
	GitlabServerFilterEventsPullRequestCreated  GitlabServerFilterEventsEnum = "PULL_REQUEST_CREATED"
	GitlabServerFilterEventsPullRequestUpdated  GitlabServerFilterEventsEnum = "PULL_REQUEST_UPDATED"
	GitlabServerFilterEventsPullRequestReopened GitlabServerFilterEventsEnum = "PULL_REQUEST_REOPENED"
	GitlabServerFilterEventsPullRequestMerged   GitlabServerFilterEventsEnum = "PULL_REQUEST_MERGED"
)

var mappingGitlabServerFilterEventsEnum = map[string]GitlabServerFilterEventsEnum{
	"PUSH":                  GitlabServerFilterEventsPush,
	"PULL_REQUEST_CREATED":  GitlabServerFilterEventsPullRequestCreated,
	"PULL_REQUEST_UPDATED":  GitlabServerFilterEventsPullRequestUpdated,
	"PULL_REQUEST_REOPENED": GitlabServerFilterEventsPullRequestReopened,
	"PULL_REQUEST_MERGED":   GitlabServerFilterEventsPullRequestMerged,
}

var mappingGitlabServerFilterEventsEnumLowerCase = map[string]GitlabServerFilterEventsEnum{
	"push":                  GitlabServerFilterEventsPush,
	"pull_request_created":  GitlabServerFilterEventsPullRequestCreated,
	"pull_request_updated":  GitlabServerFilterEventsPullRequestUpdated,
	"pull_request_reopened": GitlabServerFilterEventsPullRequestReopened,
	"pull_request_merged":   GitlabServerFilterEventsPullRequestMerged,
}

// GetGitlabServerFilterEventsEnumValues Enumerates the set of values for GitlabServerFilterEventsEnum
func GetGitlabServerFilterEventsEnumValues() []GitlabServerFilterEventsEnum {
	values := make([]GitlabServerFilterEventsEnum, 0)
	for _, v := range mappingGitlabServerFilterEventsEnum {
		values = append(values, v)
	}
	return values
}

// GetGitlabServerFilterEventsEnumStringValues Enumerates the set of values in String for GitlabServerFilterEventsEnum
func GetGitlabServerFilterEventsEnumStringValues() []string {
	return []string{
		"PUSH",
		"PULL_REQUEST_CREATED",
		"PULL_REQUEST_UPDATED",
		"PULL_REQUEST_REOPENED",
		"PULL_REQUEST_MERGED",
	}
}

// GetMappingGitlabServerFilterEventsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGitlabServerFilterEventsEnum(val string) (GitlabServerFilterEventsEnum, bool) {
	enum, ok := mappingGitlabServerFilterEventsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
