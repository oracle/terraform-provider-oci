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

// BuildRunnerShapeConfig The information about build runner.
type BuildRunnerShapeConfig interface {
}

type buildrunnershapeconfig struct {
	JsonData        []byte
	BuildRunnerType string `json:"buildRunnerType"`
}

// UnmarshalJSON unmarshals json
func (m *buildrunnershapeconfig) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerbuildrunnershapeconfig buildrunnershapeconfig
	s := struct {
		Model Unmarshalerbuildrunnershapeconfig
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.BuildRunnerType = s.Model.BuildRunnerType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *buildrunnershapeconfig) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.BuildRunnerType {
	case "DEFAULT":
		mm := DefaultBuildRunnerShapeConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CUSTOM":
		mm := CustomBuildRunnerShapeConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for BuildRunnerShapeConfig: %s.", m.BuildRunnerType)
		return *m, nil
	}
}

func (m buildrunnershapeconfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m buildrunnershapeconfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BuildRunnerShapeConfigBuildRunnerTypeEnum Enum with underlying type: string
type BuildRunnerShapeConfigBuildRunnerTypeEnum string

// Set of constants representing the allowable values for BuildRunnerShapeConfigBuildRunnerTypeEnum
const (
	BuildRunnerShapeConfigBuildRunnerTypeCustom  BuildRunnerShapeConfigBuildRunnerTypeEnum = "CUSTOM"
	BuildRunnerShapeConfigBuildRunnerTypeDefault BuildRunnerShapeConfigBuildRunnerTypeEnum = "DEFAULT"
)

var mappingBuildRunnerShapeConfigBuildRunnerTypeEnum = map[string]BuildRunnerShapeConfigBuildRunnerTypeEnum{
	"CUSTOM":  BuildRunnerShapeConfigBuildRunnerTypeCustom,
	"DEFAULT": BuildRunnerShapeConfigBuildRunnerTypeDefault,
}

var mappingBuildRunnerShapeConfigBuildRunnerTypeEnumLowerCase = map[string]BuildRunnerShapeConfigBuildRunnerTypeEnum{
	"custom":  BuildRunnerShapeConfigBuildRunnerTypeCustom,
	"default": BuildRunnerShapeConfigBuildRunnerTypeDefault,
}

// GetBuildRunnerShapeConfigBuildRunnerTypeEnumValues Enumerates the set of values for BuildRunnerShapeConfigBuildRunnerTypeEnum
func GetBuildRunnerShapeConfigBuildRunnerTypeEnumValues() []BuildRunnerShapeConfigBuildRunnerTypeEnum {
	values := make([]BuildRunnerShapeConfigBuildRunnerTypeEnum, 0)
	for _, v := range mappingBuildRunnerShapeConfigBuildRunnerTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetBuildRunnerShapeConfigBuildRunnerTypeEnumStringValues Enumerates the set of values in String for BuildRunnerShapeConfigBuildRunnerTypeEnum
func GetBuildRunnerShapeConfigBuildRunnerTypeEnumStringValues() []string {
	return []string{
		"CUSTOM",
		"DEFAULT",
	}
}

// GetMappingBuildRunnerShapeConfigBuildRunnerTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBuildRunnerShapeConfigBuildRunnerTypeEnum(val string) (BuildRunnerShapeConfigBuildRunnerTypeEnum, bool) {
	enum, ok := mappingBuildRunnerShapeConfigBuildRunnerTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
