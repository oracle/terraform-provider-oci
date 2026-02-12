// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// GitRepoContentDetails Contains data about file and branch of the script.
type GitRepoContentDetails struct {

	// OCID of the associated source provider platform configuration.
	SourceProvider *string `mandatory:"true" json:"sourceProvider"`

	// Branch name in the repository.
	Branch *string `mandatory:"false" json:"branch"`

	// Path to the file from the root of the repository.
	FilePath *string `mandatory:"false" json:"filePath"`

	// Content can be single file or whole repo. File type can only have Token based source provider and Repo type can only have key based source provider.
	ContentSelection GitRepoContentDetailsContentSelectionEnum `mandatory:"true" json:"contentSelection"`
}

func (m GitRepoContentDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GitRepoContentDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGitRepoContentDetailsContentSelectionEnum(string(m.ContentSelection)); !ok && m.ContentSelection != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ContentSelection: %s. Supported values are: %s.", m.ContentSelection, strings.Join(GetGitRepoContentDetailsContentSelectionEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m GitRepoContentDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeGitRepoContentDetails GitRepoContentDetails
	s := struct {
		DiscriminatorParam string `json:"sourceType"`
		MarshalTypeGitRepoContentDetails
	}{
		"GIT_REPO",
		(MarshalTypeGitRepoContentDetails)(m),
	}

	return json.Marshal(&s)
}

// GitRepoContentDetailsContentSelectionEnum Enum with underlying type: string
type GitRepoContentDetailsContentSelectionEnum string

// Set of constants representing the allowable values for GitRepoContentDetailsContentSelectionEnum
const (
	GitRepoContentDetailsContentSelectionFile       GitRepoContentDetailsContentSelectionEnum = "FILE"
	GitRepoContentDetailsContentSelectionRepository GitRepoContentDetailsContentSelectionEnum = "REPOSITORY"
)

var mappingGitRepoContentDetailsContentSelectionEnum = map[string]GitRepoContentDetailsContentSelectionEnum{
	"FILE":       GitRepoContentDetailsContentSelectionFile,
	"REPOSITORY": GitRepoContentDetailsContentSelectionRepository,
}

var mappingGitRepoContentDetailsContentSelectionEnumLowerCase = map[string]GitRepoContentDetailsContentSelectionEnum{
	"file":       GitRepoContentDetailsContentSelectionFile,
	"repository": GitRepoContentDetailsContentSelectionRepository,
}

// GetGitRepoContentDetailsContentSelectionEnumValues Enumerates the set of values for GitRepoContentDetailsContentSelectionEnum
func GetGitRepoContentDetailsContentSelectionEnumValues() []GitRepoContentDetailsContentSelectionEnum {
	values := make([]GitRepoContentDetailsContentSelectionEnum, 0)
	for _, v := range mappingGitRepoContentDetailsContentSelectionEnum {
		values = append(values, v)
	}
	return values
}

// GetGitRepoContentDetailsContentSelectionEnumStringValues Enumerates the set of values in String for GitRepoContentDetailsContentSelectionEnum
func GetGitRepoContentDetailsContentSelectionEnumStringValues() []string {
	return []string{
		"FILE",
		"REPOSITORY",
	}
}

// GetMappingGitRepoContentDetailsContentSelectionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGitRepoContentDetailsContentSelectionEnum(val string) (GitRepoContentDetailsContentSelectionEnum, bool) {
	enum, ok := mappingGitRepoContentDetailsContentSelectionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
