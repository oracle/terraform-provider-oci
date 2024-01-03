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

// ExecuteMergePullRequestDetails The information about the merge.
type ExecuteMergePullRequestDetails struct {

	// The commit message to be shown for this pull request in the destination branch after merge is done.
	CommitMessage *string `mandatory:"true" json:"commitMessage"`

	// What needs to happen after the merge is done successfully.
	PostMergeAction ExecuteMergePullRequestDetailsPostMergeActionEnum `mandatory:"false" json:"postMergeAction,omitempty"`

	// the strategy of merging.
	MergeStrategy MergeStrategyEnum `mandatory:"true" json:"mergeStrategy"`
}

func (m ExecuteMergePullRequestDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExecuteMergePullRequestDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExecuteMergePullRequestDetailsPostMergeActionEnum(string(m.PostMergeAction)); !ok && m.PostMergeAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PostMergeAction: %s. Supported values are: %s.", m.PostMergeAction, strings.Join(GetExecuteMergePullRequestDetailsPostMergeActionEnumStringValues(), ",")))
	}

	if _, ok := GetMappingMergeStrategyEnum(string(m.MergeStrategy)); !ok && m.MergeStrategy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MergeStrategy: %s. Supported values are: %s.", m.MergeStrategy, strings.Join(GetMergeStrategyEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ExecuteMergePullRequestDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeExecuteMergePullRequestDetails ExecuteMergePullRequestDetails
	s := struct {
		DiscriminatorParam string `json:"actionType"`
		MarshalTypeExecuteMergePullRequestDetails
	}{
		"EXECUTE",
		(MarshalTypeExecuteMergePullRequestDetails)(m),
	}

	return json.Marshal(&s)
}

// ExecuteMergePullRequestDetailsPostMergeActionEnum Enum with underlying type: string
type ExecuteMergePullRequestDetailsPostMergeActionEnum string

// Set of constants representing the allowable values for ExecuteMergePullRequestDetailsPostMergeActionEnum
const (
	ExecuteMergePullRequestDetailsPostMergeActionDeleteSourceBranch ExecuteMergePullRequestDetailsPostMergeActionEnum = "DELETE_SOURCE_BRANCH"
	ExecuteMergePullRequestDetailsPostMergeActionKeepSourceBranch   ExecuteMergePullRequestDetailsPostMergeActionEnum = "KEEP_SOURCE_BRANCH"
)

var mappingExecuteMergePullRequestDetailsPostMergeActionEnum = map[string]ExecuteMergePullRequestDetailsPostMergeActionEnum{
	"DELETE_SOURCE_BRANCH": ExecuteMergePullRequestDetailsPostMergeActionDeleteSourceBranch,
	"KEEP_SOURCE_BRANCH":   ExecuteMergePullRequestDetailsPostMergeActionKeepSourceBranch,
}

var mappingExecuteMergePullRequestDetailsPostMergeActionEnumLowerCase = map[string]ExecuteMergePullRequestDetailsPostMergeActionEnum{
	"delete_source_branch": ExecuteMergePullRequestDetailsPostMergeActionDeleteSourceBranch,
	"keep_source_branch":   ExecuteMergePullRequestDetailsPostMergeActionKeepSourceBranch,
}

// GetExecuteMergePullRequestDetailsPostMergeActionEnumValues Enumerates the set of values for ExecuteMergePullRequestDetailsPostMergeActionEnum
func GetExecuteMergePullRequestDetailsPostMergeActionEnumValues() []ExecuteMergePullRequestDetailsPostMergeActionEnum {
	values := make([]ExecuteMergePullRequestDetailsPostMergeActionEnum, 0)
	for _, v := range mappingExecuteMergePullRequestDetailsPostMergeActionEnum {
		values = append(values, v)
	}
	return values
}

// GetExecuteMergePullRequestDetailsPostMergeActionEnumStringValues Enumerates the set of values in String for ExecuteMergePullRequestDetailsPostMergeActionEnum
func GetExecuteMergePullRequestDetailsPostMergeActionEnumStringValues() []string {
	return []string{
		"DELETE_SOURCE_BRANCH",
		"KEEP_SOURCE_BRANCH",
	}
}

// GetMappingExecuteMergePullRequestDetailsPostMergeActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExecuteMergePullRequestDetailsPostMergeActionEnum(val string) (ExecuteMergePullRequestDetailsPostMergeActionEnum, bool) {
	enum, ok := mappingExecuteMergePullRequestDetailsPostMergeActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
