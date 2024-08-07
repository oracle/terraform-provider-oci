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

// PullRequestLifecycleActivitySummary activity describing a pull request state change
type PullRequestLifecycleActivitySummary struct {

	// activity identifier
	Id *string `mandatory:"true" json:"id"`

	Principal *PrincipalDetails `mandatory:"true" json:"principal"`

	// pullRequest OCID
	PullRequestId *string `mandatory:"true" json:"pullRequestId"`

	// The time the action was performed. An RFC3339 formatted datetime string
	TimeOccurred *common.SDKTime `mandatory:"true" json:"timeOccurred"`

	// The state of a pull request after an action.
	State PullRequestLifecycleActivitySummaryStateEnum `mandatory:"true" json:"state"`
}

// GetId returns Id
func (m PullRequestLifecycleActivitySummary) GetId() *string {
	return m.Id
}

// GetPrincipal returns Principal
func (m PullRequestLifecycleActivitySummary) GetPrincipal() *PrincipalDetails {
	return m.Principal
}

// GetPullRequestId returns PullRequestId
func (m PullRequestLifecycleActivitySummary) GetPullRequestId() *string {
	return m.PullRequestId
}

// GetTimeOccurred returns TimeOccurred
func (m PullRequestLifecycleActivitySummary) GetTimeOccurred() *common.SDKTime {
	return m.TimeOccurred
}

func (m PullRequestLifecycleActivitySummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PullRequestLifecycleActivitySummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPullRequestLifecycleActivitySummaryStateEnum(string(m.State)); !ok && m.State != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for State: %s. Supported values are: %s.", m.State, strings.Join(GetPullRequestLifecycleActivitySummaryStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PullRequestLifecycleActivitySummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePullRequestLifecycleActivitySummary PullRequestLifecycleActivitySummary
	s := struct {
		DiscriminatorParam string `json:"activityType"`
		MarshalTypePullRequestLifecycleActivitySummary
	}{
		"LIFECYCLE",
		(MarshalTypePullRequestLifecycleActivitySummary)(m),
	}

	return json.Marshal(&s)
}

// PullRequestLifecycleActivitySummaryStateEnum Enum with underlying type: string
type PullRequestLifecycleActivitySummaryStateEnum string

// Set of constants representing the allowable values for PullRequestLifecycleActivitySummaryStateEnum
const (
	PullRequestLifecycleActivitySummaryStateOpened   PullRequestLifecycleActivitySummaryStateEnum = "OPENED"
	PullRequestLifecycleActivitySummaryStateClosed   PullRequestLifecycleActivitySummaryStateEnum = "CLOSED"
	PullRequestLifecycleActivitySummaryStateMerged   PullRequestLifecycleActivitySummaryStateEnum = "MERGED"
	PullRequestLifecycleActivitySummaryStateReopened PullRequestLifecycleActivitySummaryStateEnum = "REOPENED"
)

var mappingPullRequestLifecycleActivitySummaryStateEnum = map[string]PullRequestLifecycleActivitySummaryStateEnum{
	"OPENED":   PullRequestLifecycleActivitySummaryStateOpened,
	"CLOSED":   PullRequestLifecycleActivitySummaryStateClosed,
	"MERGED":   PullRequestLifecycleActivitySummaryStateMerged,
	"REOPENED": PullRequestLifecycleActivitySummaryStateReopened,
}

var mappingPullRequestLifecycleActivitySummaryStateEnumLowerCase = map[string]PullRequestLifecycleActivitySummaryStateEnum{
	"opened":   PullRequestLifecycleActivitySummaryStateOpened,
	"closed":   PullRequestLifecycleActivitySummaryStateClosed,
	"merged":   PullRequestLifecycleActivitySummaryStateMerged,
	"reopened": PullRequestLifecycleActivitySummaryStateReopened,
}

// GetPullRequestLifecycleActivitySummaryStateEnumValues Enumerates the set of values for PullRequestLifecycleActivitySummaryStateEnum
func GetPullRequestLifecycleActivitySummaryStateEnumValues() []PullRequestLifecycleActivitySummaryStateEnum {
	values := make([]PullRequestLifecycleActivitySummaryStateEnum, 0)
	for _, v := range mappingPullRequestLifecycleActivitySummaryStateEnum {
		values = append(values, v)
	}
	return values
}

// GetPullRequestLifecycleActivitySummaryStateEnumStringValues Enumerates the set of values in String for PullRequestLifecycleActivitySummaryStateEnum
func GetPullRequestLifecycleActivitySummaryStateEnumStringValues() []string {
	return []string{
		"OPENED",
		"CLOSED",
		"MERGED",
		"REOPENED",
	}
}

// GetMappingPullRequestLifecycleActivitySummaryStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPullRequestLifecycleActivitySummaryStateEnum(val string) (PullRequestLifecycleActivitySummaryStateEnum, bool) {
	enum, ok := mappingPullRequestLifecycleActivitySummaryStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
