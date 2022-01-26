// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// BuildRunSource The source from which the build run is triggered.
type BuildRunSource interface {
}

type buildrunsource struct {
	JsonData   []byte
	SourceType string `json:"sourceType"`
}

// UnmarshalJSON unmarshals json
func (m *buildrunsource) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerbuildrunsource buildrunsource
	s := struct {
		Model Unmarshalerbuildrunsource
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.SourceType = s.Model.SourceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *buildrunsource) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.SourceType {
	case "GITHUB":
		mm := GithubBuildRunSource{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DEVOPS_CODE_REPOSITORY":
		mm := DevopsCodeRepositoryBuildRunSource{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MANUAL":
		mm := ManualBuildRunSource{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GITLAB":
		mm := GitlabBuildRunSource{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m buildrunsource) String() string {
	return common.PointerString(m)
}

// BuildRunSourceSourceTypeEnum Enum with underlying type: string
type BuildRunSourceSourceTypeEnum string

// Set of constants representing the allowable values for BuildRunSourceSourceTypeEnum
const (
	BuildRunSourceSourceTypeManual               BuildRunSourceSourceTypeEnum = "MANUAL"
	BuildRunSourceSourceTypeGithub               BuildRunSourceSourceTypeEnum = "GITHUB"
	BuildRunSourceSourceTypeGitlab               BuildRunSourceSourceTypeEnum = "GITLAB"
	BuildRunSourceSourceTypeDevopsCodeRepository BuildRunSourceSourceTypeEnum = "DEVOPS_CODE_REPOSITORY"
)

var mappingBuildRunSourceSourceType = map[string]BuildRunSourceSourceTypeEnum{
	"MANUAL":                 BuildRunSourceSourceTypeManual,
	"GITHUB":                 BuildRunSourceSourceTypeGithub,
	"GITLAB":                 BuildRunSourceSourceTypeGitlab,
	"DEVOPS_CODE_REPOSITORY": BuildRunSourceSourceTypeDevopsCodeRepository,
}

// GetBuildRunSourceSourceTypeEnumValues Enumerates the set of values for BuildRunSourceSourceTypeEnum
func GetBuildRunSourceSourceTypeEnumValues() []BuildRunSourceSourceTypeEnum {
	values := make([]BuildRunSourceSourceTypeEnum, 0)
	for _, v := range mappingBuildRunSourceSourceType {
		values = append(values, v)
	}
	return values
}
