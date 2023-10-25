// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Dependency Management API
//
// Use the Application Dependency Management API to create knowledge bases and vulnerability audits.  For more information, see ADM (https://docs.cloud.oracle.com/Content/application-dependency-management/home.htm).
//

package adm

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExternalScmConfiguration An external SCM configuration extends a SCM Configuration with necessary data to reach and use the Source Code Management tool/platform used by a Remediation Recipe.
// An external SCM in ADM refers to GitHub, or GitLab.
type ExternalScmConfiguration struct {

	// The branch used by ADM to patch vulnerabilities.
	Branch *string `mandatory:"true" json:"branch"`

	// If true, the Pull Request (PR) will be merged after the verify stage completes successfully
	// If false, the PR with the proposed changes must be reviewed and manually merged.
	IsAutomergeEnabled *bool `mandatory:"true" json:"isAutomergeEnabled"`

	// The repository URL for the SCM.
	// For Non-Enterprise GitHub the expected format is https://github.com/[owner]/[repoName]
	// For Enterprise GitHub the expected format is http(s)://[hostname]/api/v3/repos/[owner]/[repoName]
	// For GitLab the expected format is https://gitlab.com/[groupName]/[repoName]
	RepositoryUrl *string `mandatory:"true" json:"repositoryUrl"`

	// The Oracle Cloud Identifier (OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)) of the Private Access Token (PAT) Secret.
	// The secret provides the credentials necessary to authenticate against the SCM.
	PatSecretId *string `mandatory:"true" json:"patSecretId"`

	// The location of the build file relative to the root of the repository. Only Maven build files (POM) are currently supported.
	// If this property is not specified, ADM will use the build file located at the root of the repository.
	BuildFileLocation *string `mandatory:"false" json:"buildFileLocation"`

	// The username for the SCM (to perform operations such as cloning or pushing via HTTP).
	Username *string `mandatory:"false" json:"username"`

	// The type of External Source Code Management.
	ExternalScmType ExternalScmConfigurationExternalScmTypeEnum `mandatory:"true" json:"externalScmType"`
}

// GetBranch returns Branch
func (m ExternalScmConfiguration) GetBranch() *string {
	return m.Branch
}

// GetBuildFileLocation returns BuildFileLocation
func (m ExternalScmConfiguration) GetBuildFileLocation() *string {
	return m.BuildFileLocation
}

// GetIsAutomergeEnabled returns IsAutomergeEnabled
func (m ExternalScmConfiguration) GetIsAutomergeEnabled() *bool {
	return m.IsAutomergeEnabled
}

func (m ExternalScmConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExternalScmConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExternalScmConfigurationExternalScmTypeEnum(string(m.ExternalScmType)); !ok && m.ExternalScmType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExternalScmType: %s. Supported values are: %s.", m.ExternalScmType, strings.Join(GetExternalScmConfigurationExternalScmTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ExternalScmConfiguration) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeExternalScmConfiguration ExternalScmConfiguration
	s := struct {
		DiscriminatorParam string `json:"scmType"`
		MarshalTypeExternalScmConfiguration
	}{
		"EXTERNAL_SCM",
		(MarshalTypeExternalScmConfiguration)(m),
	}

	return json.Marshal(&s)
}

// ExternalScmConfigurationExternalScmTypeEnum Enum with underlying type: string
type ExternalScmConfigurationExternalScmTypeEnum string

// Set of constants representing the allowable values for ExternalScmConfigurationExternalScmTypeEnum
const (
	ExternalScmConfigurationExternalScmTypeGithub ExternalScmConfigurationExternalScmTypeEnum = "GITHUB"
	ExternalScmConfigurationExternalScmTypeGitlab ExternalScmConfigurationExternalScmTypeEnum = "GITLAB"
)

var mappingExternalScmConfigurationExternalScmTypeEnum = map[string]ExternalScmConfigurationExternalScmTypeEnum{
	"GITHUB": ExternalScmConfigurationExternalScmTypeGithub,
	"GITLAB": ExternalScmConfigurationExternalScmTypeGitlab,
}

var mappingExternalScmConfigurationExternalScmTypeEnumLowerCase = map[string]ExternalScmConfigurationExternalScmTypeEnum{
	"github": ExternalScmConfigurationExternalScmTypeGithub,
	"gitlab": ExternalScmConfigurationExternalScmTypeGitlab,
}

// GetExternalScmConfigurationExternalScmTypeEnumValues Enumerates the set of values for ExternalScmConfigurationExternalScmTypeEnum
func GetExternalScmConfigurationExternalScmTypeEnumValues() []ExternalScmConfigurationExternalScmTypeEnum {
	values := make([]ExternalScmConfigurationExternalScmTypeEnum, 0)
	for _, v := range mappingExternalScmConfigurationExternalScmTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetExternalScmConfigurationExternalScmTypeEnumStringValues Enumerates the set of values in String for ExternalScmConfigurationExternalScmTypeEnum
func GetExternalScmConfigurationExternalScmTypeEnumStringValues() []string {
	return []string{
		"GITHUB",
		"GITLAB",
	}
}

// GetMappingExternalScmConfigurationExternalScmTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExternalScmConfigurationExternalScmTypeEnum(val string) (ExternalScmConfigurationExternalScmTypeEnum, bool) {
	enum, ok := mappingExternalScmConfigurationExternalScmTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
