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

// BuildSource Build source required for the Build stage.
type BuildSource interface {

	// Name of the build source. This must be unique within a build source collection. The name can be used by customers to locate the working directory pertinent to this repository.
	GetName() *string

	// URL for the repository.
	GetRepositoryUrl() *string

	// Branch name.
	GetBranch() *string
}

type buildsource struct {
	JsonData       []byte
	Name           *string `mandatory:"true" json:"name"`
	RepositoryUrl  *string `mandatory:"true" json:"repositoryUrl"`
	Branch         *string `mandatory:"true" json:"branch"`
	ConnectionType string  `json:"connectionType"`
}

// UnmarshalJSON unmarshals json
func (m *buildsource) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerbuildsource buildsource
	s := struct {
		Model Unmarshalerbuildsource
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Name = s.Model.Name
	m.RepositoryUrl = s.Model.RepositoryUrl
	m.Branch = s.Model.Branch
	m.ConnectionType = s.Model.ConnectionType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *buildsource) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ConnectionType {
	case "VBS":
		mm := VbsBuildSource{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "BITBUCKET_SERVER":
		mm := BitbucketServerBuildSource{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GITHUB":
		mm := GithubBuildSource{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "BITBUCKET_CLOUD":
		mm := BitbucketCloudBuildSource{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GITLAB_SERVER":
		mm := GitlabServerBuildSource{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DEVOPS_CODE_REPOSITORY":
		mm := DevopsCodeRepositoryBuildSource{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GITLAB":
		mm := GitlabBuildSource{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for BuildSource: %s.", m.ConnectionType)
		return *m, nil
	}
}

// GetName returns Name
func (m buildsource) GetName() *string {
	return m.Name
}

// GetRepositoryUrl returns RepositoryUrl
func (m buildsource) GetRepositoryUrl() *string {
	return m.RepositoryUrl
}

// GetBranch returns Branch
func (m buildsource) GetBranch() *string {
	return m.Branch
}

func (m buildsource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m buildsource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BuildSourceConnectionTypeEnum Enum with underlying type: string
type BuildSourceConnectionTypeEnum string

// Set of constants representing the allowable values for BuildSourceConnectionTypeEnum
const (
	BuildSourceConnectionTypeGithub               BuildSourceConnectionTypeEnum = "GITHUB"
	BuildSourceConnectionTypeGitlab               BuildSourceConnectionTypeEnum = "GITLAB"
	BuildSourceConnectionTypeGitlabServer         BuildSourceConnectionTypeEnum = "GITLAB_SERVER"
	BuildSourceConnectionTypeBitbucketCloud       BuildSourceConnectionTypeEnum = "BITBUCKET_CLOUD"
	BuildSourceConnectionTypeBitbucketServer      BuildSourceConnectionTypeEnum = "BITBUCKET_SERVER"
	BuildSourceConnectionTypeDevopsCodeRepository BuildSourceConnectionTypeEnum = "DEVOPS_CODE_REPOSITORY"
	BuildSourceConnectionTypeVbs                  BuildSourceConnectionTypeEnum = "VBS"
)

var mappingBuildSourceConnectionTypeEnum = map[string]BuildSourceConnectionTypeEnum{
	"GITHUB":                 BuildSourceConnectionTypeGithub,
	"GITLAB":                 BuildSourceConnectionTypeGitlab,
	"GITLAB_SERVER":          BuildSourceConnectionTypeGitlabServer,
	"BITBUCKET_CLOUD":        BuildSourceConnectionTypeBitbucketCloud,
	"BITBUCKET_SERVER":       BuildSourceConnectionTypeBitbucketServer,
	"DEVOPS_CODE_REPOSITORY": BuildSourceConnectionTypeDevopsCodeRepository,
	"VBS":                    BuildSourceConnectionTypeVbs,
}

var mappingBuildSourceConnectionTypeEnumLowerCase = map[string]BuildSourceConnectionTypeEnum{
	"github":                 BuildSourceConnectionTypeGithub,
	"gitlab":                 BuildSourceConnectionTypeGitlab,
	"gitlab_server":          BuildSourceConnectionTypeGitlabServer,
	"bitbucket_cloud":        BuildSourceConnectionTypeBitbucketCloud,
	"bitbucket_server":       BuildSourceConnectionTypeBitbucketServer,
	"devops_code_repository": BuildSourceConnectionTypeDevopsCodeRepository,
	"vbs":                    BuildSourceConnectionTypeVbs,
}

// GetBuildSourceConnectionTypeEnumValues Enumerates the set of values for BuildSourceConnectionTypeEnum
func GetBuildSourceConnectionTypeEnumValues() []BuildSourceConnectionTypeEnum {
	values := make([]BuildSourceConnectionTypeEnum, 0)
	for _, v := range mappingBuildSourceConnectionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetBuildSourceConnectionTypeEnumStringValues Enumerates the set of values in String for BuildSourceConnectionTypeEnum
func GetBuildSourceConnectionTypeEnumStringValues() []string {
	return []string{
		"GITHUB",
		"GITLAB",
		"GITLAB_SERVER",
		"BITBUCKET_CLOUD",
		"BITBUCKET_SERVER",
		"DEVOPS_CODE_REPOSITORY",
		"VBS",
	}
}

// GetMappingBuildSourceConnectionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBuildSourceConnectionTypeEnum(val string) (BuildSourceConnectionTypeEnum, bool) {
	enum, ok := mappingBuildSourceConnectionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
