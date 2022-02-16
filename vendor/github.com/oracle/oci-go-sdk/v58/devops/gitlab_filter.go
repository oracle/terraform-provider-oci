// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// GitlabFilter The filter for GitLab events.
type GitlabFilter struct {
	Include *GitlabFilterAttributes `mandatory:"false" json:"include"`

	// The events, for example, PUSH, PULL_REQUEST_MERGE.
	Events []GitlabFilterEventsEnum `mandatory:"false" json:"events,omitempty"`
}

func (m GitlabFilter) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GitlabFilter) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range m.Events {
		if _, ok := GetMappingGitlabFilterEventsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Events: %s. Supported values are: %s.", val, strings.Join(GetGitlabFilterEventsEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m GitlabFilter) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeGitlabFilter GitlabFilter
	s := struct {
		DiscriminatorParam string `json:"triggerSource"`
		MarshalTypeGitlabFilter
	}{
		"GITLAB",
		(MarshalTypeGitlabFilter)(m),
	}

	return json.Marshal(&s)
}

// GitlabFilterEventsEnum Enum with underlying type: string
type GitlabFilterEventsEnum string

// Set of constants representing the allowable values for GitlabFilterEventsEnum
const (
	GitlabFilterEventsPush                GitlabFilterEventsEnum = "PUSH"
	GitlabFilterEventsPullRequestCreated  GitlabFilterEventsEnum = "PULL_REQUEST_CREATED"
	GitlabFilterEventsPullRequestUpdated  GitlabFilterEventsEnum = "PULL_REQUEST_UPDATED"
	GitlabFilterEventsPullRequestReopened GitlabFilterEventsEnum = "PULL_REQUEST_REOPENED"
	GitlabFilterEventsPullRequestMerged   GitlabFilterEventsEnum = "PULL_REQUEST_MERGED"
)

var mappingGitlabFilterEventsEnum = map[string]GitlabFilterEventsEnum{
	"PUSH":                  GitlabFilterEventsPush,
	"PULL_REQUEST_CREATED":  GitlabFilterEventsPullRequestCreated,
	"PULL_REQUEST_UPDATED":  GitlabFilterEventsPullRequestUpdated,
	"PULL_REQUEST_REOPENED": GitlabFilterEventsPullRequestReopened,
	"PULL_REQUEST_MERGED":   GitlabFilterEventsPullRequestMerged,
}

// GetGitlabFilterEventsEnumValues Enumerates the set of values for GitlabFilterEventsEnum
func GetGitlabFilterEventsEnumValues() []GitlabFilterEventsEnum {
	values := make([]GitlabFilterEventsEnum, 0)
	for _, v := range mappingGitlabFilterEventsEnum {
		values = append(values, v)
	}
	return values
}

// GetGitlabFilterEventsEnumStringValues Enumerates the set of values in String for GitlabFilterEventsEnum
func GetGitlabFilterEventsEnumStringValues() []string {
	return []string{
		"PUSH",
		"PULL_REQUEST_CREATED",
		"PULL_REQUEST_UPDATED",
		"PULL_REQUEST_REOPENED",
		"PULL_REQUEST_MERGED",
	}
}

// GetMappingGitlabFilterEventsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGitlabFilterEventsEnum(val string) (GitlabFilterEventsEnum, bool) {
	mappingGitlabFilterEventsEnumIgnoreCase := make(map[string]GitlabFilterEventsEnum)
	for k, v := range mappingGitlabFilterEventsEnum {
		mappingGitlabFilterEventsEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingGitlabFilterEventsEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
