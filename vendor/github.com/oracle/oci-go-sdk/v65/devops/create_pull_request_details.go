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

// CreatePullRequestDetails The information about new Pull Request.
type CreatePullRequestDetails struct {

	// Pull Request title
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The source branch of the pull request.
	SourceBranch *string `mandatory:"true" json:"sourceBranch"`

	// The OCID of the repository.
	RepositoryId *string `mandatory:"true" json:"repositoryId"`

	// The destination branch of the pull request. If not provided, default branch will be used as the destination branch.
	DestinationBranch *string `mandatory:"false" json:"destinationBranch"`

	// The OCID of the forked repository that will act as the source of the changes to be included in the pull request against the parent repository.
	SourceRepositoryId *string `mandatory:"false" json:"sourceRepositoryId"`

	// Details of the pull request. Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// Reviewers for this pull request.
	Reviewers []CreateReviewerDetails `mandatory:"false" json:"reviewers"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreatePullRequestDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreatePullRequestDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
