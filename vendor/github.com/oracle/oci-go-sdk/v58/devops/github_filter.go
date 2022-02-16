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

// GithubFilter The filter for GitHub events.
type GithubFilter struct {
	Include *GithubFilterAttributes `mandatory:"false" json:"include"`

	// The events, for example, PUSH, PULL_REQUEST_MERGE.
	Events []GithubFilterEventsEnum `mandatory:"false" json:"events,omitempty"`
}

func (m GithubFilter) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GithubFilter) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range m.Events {
		if _, ok := GetMappingGithubFilterEventsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Events: %s. Supported values are: %s.", val, strings.Join(GetGithubFilterEventsEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingGithubFilterEventsEnum = map[string]GithubFilterEventsEnum{
	"PUSH":                  GithubFilterEventsPush,
	"PULL_REQUEST_CREATED":  GithubFilterEventsPullRequestCreated,
	"PULL_REQUEST_UPDATED":  GithubFilterEventsPullRequestUpdated,
	"PULL_REQUEST_REOPENED": GithubFilterEventsPullRequestReopened,
	"PULL_REQUEST_MERGED":   GithubFilterEventsPullRequestMerged,
}

// GetGithubFilterEventsEnumValues Enumerates the set of values for GithubFilterEventsEnum
func GetGithubFilterEventsEnumValues() []GithubFilterEventsEnum {
	values := make([]GithubFilterEventsEnum, 0)
	for _, v := range mappingGithubFilterEventsEnum {
		values = append(values, v)
	}
	return values
}

// GetGithubFilterEventsEnumStringValues Enumerates the set of values in String for GithubFilterEventsEnum
func GetGithubFilterEventsEnumStringValues() []string {
	return []string{
		"PUSH",
		"PULL_REQUEST_CREATED",
		"PULL_REQUEST_UPDATED",
		"PULL_REQUEST_REOPENED",
		"PULL_REQUEST_MERGED",
	}
}

// GetMappingGithubFilterEventsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGithubFilterEventsEnum(val string) (GithubFilterEventsEnum, bool) {
	mappingGithubFilterEventsEnumIgnoreCase := make(map[string]GithubFilterEventsEnum)
	for k, v := range mappingGithubFilterEventsEnum {
		mappingGithubFilterEventsEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingGithubFilterEventsEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
