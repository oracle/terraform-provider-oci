// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Dependency Management API
//
// Use the Application Dependency Management API to create knowledge bases and vulnerability audits.  For more information, see ADM (https://docs.oracle.com/iaas/Content/application-dependency-management/home.htm).
//

package adm

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// GithubCodeSnippetSourceLocation Github Source location of a code snippet.
type GithubCodeSnippetSourceLocation struct {
	SourceCodeUrl *string `mandatory:"true" json:"sourceCodeUrl"`

	FilePath *string `mandatory:"true" json:"filePath"`
}

func (m GithubCodeSnippetSourceLocation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GithubCodeSnippetSourceLocation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m GithubCodeSnippetSourceLocation) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeGithubCodeSnippetSourceLocation GithubCodeSnippetSourceLocation
	s := struct {
		DiscriminatorParam string `json:"origin"`
		MarshalTypeGithubCodeSnippetSourceLocation
	}{
		"GITHUB",
		(MarshalTypeGithubCodeSnippetSourceLocation)(m),
	}

	return json.Marshal(&s)
}
