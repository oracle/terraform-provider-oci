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

// CreatePullRequestCommentDetails The information about new Comment.
type CreatePullRequestCommentDetails struct {

	// Content of the Comment.
	Data *string `mandatory:"true" json:"data"`

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
}

func (m CreatePullRequestCommentDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreatePullRequestCommentDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingPullRequestCommentFileTypeEnum(string(m.FileType)); !ok && m.FileType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FileType: %s. Supported values are: %s.", m.FileType, strings.Join(GetPullRequestCommentFileTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
