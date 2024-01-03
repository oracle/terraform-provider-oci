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

// ApprovalRule A rule which must be satisfied for matching pull requests before the pull request can be merged
type ApprovalRule struct {

	// Name which is used to uniquely identify an approval rule.
	Name *string `mandatory:"true" json:"name"`

	// Minimum number of approvals which must be provided by the reviewers specified in the list before the rule can be satisfied
	MinApprovalsCount *int `mandatory:"true" json:"minApprovalsCount"`

	// Branch name where pull requests targeting the branch must satisfy the approval rule. This value being null means the rule applies to all pull requests
	DestinationBranch *string `mandatory:"false" json:"destinationBranch"`

	// List of users who must provide approvals up to the minApprovalsCount specified in the rule. An empty list means the approvals can come from any user.
	Reviewers []PrincipalDetails `mandatory:"false" json:"reviewers"`
}

func (m ApprovalRule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ApprovalRule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
