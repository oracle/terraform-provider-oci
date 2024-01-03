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

// CodeSearchResultSummary Object wrapping the matching code snippets results in a file and the metadata about the file and repo.
type CodeSearchResultSummary struct {

	// The OCID of the repository the file belongs to.
	RepositoryId *string `mandatory:"true" json:"repositoryId"`

	// Name of the project.
	ProjectName *string `mandatory:"true" json:"projectName"`

	// Name of the repository.
	RepositoryName *string `mandatory:"true" json:"repositoryName"`

	// Relative path of the file with respect to repository folder.
	FilePath *string `mandatory:"true" json:"filePath"`

	// Name of the ref that contains the current version of the file.
	RefName *string `mandatory:"true" json:"refName"`

	// The OCID of the containing repository's compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Name of the file.
	FileName *string `mandatory:"false" json:"fileName"`

	// Extension of the file.
	FileExtension *string `mandatory:"false" json:"fileExtension"`

	// Commit that contains the current version of the file.
	CommitId *string `mandatory:"false" json:"commitId"`

	// List of matchng code snippets in the file.
	CodeSnippets []CodeSnippet `mandatory:"false" json:"codeSnippets"`
}

func (m CodeSearchResultSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CodeSearchResultSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
