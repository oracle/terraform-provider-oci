// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.oracle.com/iaas/Content/devops/using/home.htm).
//

package devops

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PullRequestCommentSummary summary of a pullRequest comment
type PullRequestCommentSummary struct {

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// OCID of the pull request that this comment belongs to
	PullRequestId *string `mandatory:"true" json:"pullRequestId"`

	// Content of the Comment.
	Data *string `mandatory:"true" json:"data"`

	// Status of the Comment
	Status PullRequestCommentStatusEnum `mandatory:"true" json:"status"`

	// Creation timestamp. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	CreatedBy *PrincipalDetails `mandatory:"true" json:"createdBy"`

	// ID of parent Comment
	ParentId *string `mandatory:"false" json:"parentId"`

	// File path in the commit
	FilePath *string `mandatory:"false" json:"filePath"`

	// Commit SHA
	CommitId *string `mandatory:"false" json:"commitId"`

	// File path in the target commit
	FileType PullRequestCommentFileTypeEnum `mandatory:"false" json:"fileType,omitempty"`

	// Line number in the file
	LineNumber *int `mandatory:"false" json:"lineNumber"`

	Likes *PullRequestCommentLikeCollection `mandatory:"false" json:"likes"`

	// Latest update timestamp. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	UpdatedBy *PrincipalDetails `mandatory:"false" json:"updatedBy"`

	// Shows the status of an inline comments context
	ContextStatus PullRequestCommentContextStatusEnum `mandatory:"false" json:"contextStatus,omitempty"`

	// 4 line snippet to be displayed as context for inline comments
	CommentContext []DiffLineDetails `mandatory:"false" json:"commentContext"`
}

func (m PullRequestCommentSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PullRequestCommentSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPullRequestCommentStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetPullRequestCommentStatusEnumStringValues(), ",")))
	}

	if _, ok := GetMappingPullRequestCommentFileTypeEnum(string(m.FileType)); !ok && m.FileType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FileType: %s. Supported values are: %s.", m.FileType, strings.Join(GetPullRequestCommentFileTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPullRequestCommentContextStatusEnum(string(m.ContextStatus)); !ok && m.ContextStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ContextStatus: %s. Supported values are: %s.", m.ContextStatus, strings.Join(GetPullRequestCommentContextStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
