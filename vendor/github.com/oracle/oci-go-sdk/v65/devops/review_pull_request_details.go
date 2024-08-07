// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ReviewPullRequestDetails Details to submit pull request review
type ReviewPullRequestDetails struct {

	// The review action taken
	Action ReviewPullRequestDetailsActionEnum `mandatory:"true" json:"action"`
}

func (m ReviewPullRequestDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ReviewPullRequestDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingReviewPullRequestDetailsActionEnum(string(m.Action)); !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetReviewPullRequestDetailsActionEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ReviewPullRequestDetailsActionEnum Enum with underlying type: string
type ReviewPullRequestDetailsActionEnum string

// Set of constants representing the allowable values for ReviewPullRequestDetailsActionEnum
const (
	ReviewPullRequestDetailsActionApprove   ReviewPullRequestDetailsActionEnum = "APPROVE"
	ReviewPullRequestDetailsActionUnapprove ReviewPullRequestDetailsActionEnum = "UNAPPROVE"
)

var mappingReviewPullRequestDetailsActionEnum = map[string]ReviewPullRequestDetailsActionEnum{
	"APPROVE":   ReviewPullRequestDetailsActionApprove,
	"UNAPPROVE": ReviewPullRequestDetailsActionUnapprove,
}

var mappingReviewPullRequestDetailsActionEnumLowerCase = map[string]ReviewPullRequestDetailsActionEnum{
	"approve":   ReviewPullRequestDetailsActionApprove,
	"unapprove": ReviewPullRequestDetailsActionUnapprove,
}

// GetReviewPullRequestDetailsActionEnumValues Enumerates the set of values for ReviewPullRequestDetailsActionEnum
func GetReviewPullRequestDetailsActionEnumValues() []ReviewPullRequestDetailsActionEnum {
	values := make([]ReviewPullRequestDetailsActionEnum, 0)
	for _, v := range mappingReviewPullRequestDetailsActionEnum {
		values = append(values, v)
	}
	return values
}

// GetReviewPullRequestDetailsActionEnumStringValues Enumerates the set of values in String for ReviewPullRequestDetailsActionEnum
func GetReviewPullRequestDetailsActionEnumStringValues() []string {
	return []string{
		"APPROVE",
		"UNAPPROVE",
	}
}

// GetMappingReviewPullRequestDetailsActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReviewPullRequestDetailsActionEnum(val string) (ReviewPullRequestDetailsActionEnum, bool) {
	enum, ok := mappingReviewPullRequestDetailsActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
