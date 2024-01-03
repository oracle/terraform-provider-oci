// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// VerifyConfiguration The Verify stage configuration specifies a build service to run a pipeline for the recommended code changes.
// The build pipeline will be initiated to ensure that there is no breaking change after the dependency versions
// have been updated in source to avoid vulnerabilities.
type VerifyConfiguration interface {
}

type verifyconfiguration struct {
	JsonData         []byte
	BuildServiceType string `json:"buildServiceType"`
}

// UnmarshalJSON unmarshals json
func (m *verifyconfiguration) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerverifyconfiguration verifyconfiguration
	s := struct {
		Model Unmarshalerverifyconfiguration
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.BuildServiceType = s.Model.BuildServiceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *verifyconfiguration) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.BuildServiceType {
	case "JENKINS_PIPELINE":
		mm := JenkinsPipelineConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "NONE":
		mm := NoneVerifyConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_DEVOPS_BUILD":
		mm := OciDevOpsBuildConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GITHUB_ACTIONS":
		mm := GitHubActionsConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GITLAB_PIPELINE":
		mm := GitLabPipelineConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for VerifyConfiguration: %s.", m.BuildServiceType)
		return *m, nil
	}
}

func (m verifyconfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m verifyconfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// VerifyConfigurationBuildServiceTypeEnum Enum with underlying type: string
type VerifyConfigurationBuildServiceTypeEnum string

// Set of constants representing the allowable values for VerifyConfigurationBuildServiceTypeEnum
const (
	VerifyConfigurationBuildServiceTypeOciDevopsBuild  VerifyConfigurationBuildServiceTypeEnum = "OCI_DEVOPS_BUILD"
	VerifyConfigurationBuildServiceTypeGitlabPipeline  VerifyConfigurationBuildServiceTypeEnum = "GITLAB_PIPELINE"
	VerifyConfigurationBuildServiceTypeGithubActions   VerifyConfigurationBuildServiceTypeEnum = "GITHUB_ACTIONS"
	VerifyConfigurationBuildServiceTypeJenkinsPipeline VerifyConfigurationBuildServiceTypeEnum = "JENKINS_PIPELINE"
	VerifyConfigurationBuildServiceTypeNone            VerifyConfigurationBuildServiceTypeEnum = "NONE"
)

var mappingVerifyConfigurationBuildServiceTypeEnum = map[string]VerifyConfigurationBuildServiceTypeEnum{
	"OCI_DEVOPS_BUILD": VerifyConfigurationBuildServiceTypeOciDevopsBuild,
	"GITLAB_PIPELINE":  VerifyConfigurationBuildServiceTypeGitlabPipeline,
	"GITHUB_ACTIONS":   VerifyConfigurationBuildServiceTypeGithubActions,
	"JENKINS_PIPELINE": VerifyConfigurationBuildServiceTypeJenkinsPipeline,
	"NONE":             VerifyConfigurationBuildServiceTypeNone,
}

var mappingVerifyConfigurationBuildServiceTypeEnumLowerCase = map[string]VerifyConfigurationBuildServiceTypeEnum{
	"oci_devops_build": VerifyConfigurationBuildServiceTypeOciDevopsBuild,
	"gitlab_pipeline":  VerifyConfigurationBuildServiceTypeGitlabPipeline,
	"github_actions":   VerifyConfigurationBuildServiceTypeGithubActions,
	"jenkins_pipeline": VerifyConfigurationBuildServiceTypeJenkinsPipeline,
	"none":             VerifyConfigurationBuildServiceTypeNone,
}

// GetVerifyConfigurationBuildServiceTypeEnumValues Enumerates the set of values for VerifyConfigurationBuildServiceTypeEnum
func GetVerifyConfigurationBuildServiceTypeEnumValues() []VerifyConfigurationBuildServiceTypeEnum {
	values := make([]VerifyConfigurationBuildServiceTypeEnum, 0)
	for _, v := range mappingVerifyConfigurationBuildServiceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetVerifyConfigurationBuildServiceTypeEnumStringValues Enumerates the set of values in String for VerifyConfigurationBuildServiceTypeEnum
func GetVerifyConfigurationBuildServiceTypeEnumStringValues() []string {
	return []string{
		"OCI_DEVOPS_BUILD",
		"GITLAB_PIPELINE",
		"GITHUB_ACTIONS",
		"JENKINS_PIPELINE",
		"NONE",
	}
}

// GetMappingVerifyConfigurationBuildServiceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVerifyConfigurationBuildServiceTypeEnum(val string) (VerifyConfigurationBuildServiceTypeEnum, bool) {
	enum, ok := mappingVerifyConfigurationBuildServiceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
