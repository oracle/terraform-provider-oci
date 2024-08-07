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

// DeployArtifactSource Specifies source of an artifact.
type DeployArtifactSource interface {
}

type deployartifactsource struct {
	JsonData                 []byte
	DeployArtifactSourceType string `json:"deployArtifactSourceType"`
}

// UnmarshalJSON unmarshals json
func (m *deployartifactsource) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdeployartifactsource deployartifactsource
	s := struct {
		Model Unmarshalerdeployartifactsource
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DeployArtifactSourceType = s.Model.DeployArtifactSourceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *deployartifactsource) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.DeployArtifactSourceType {
	case "GENERIC_ARTIFACT":
		mm := GenericDeployArtifactSource{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HELM_CHART":
		mm := HelmRepositoryDeployArtifactSource{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCIR":
		mm := OcirDeployArtifactSource{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HELM_COMMAND_SPEC":
		mm := HelmCommandSpecArtifactSource{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "INLINE":
		mm := InlineDeployArtifactSource{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for DeployArtifactSource: %s.", m.DeployArtifactSourceType)
		return *m, nil
	}
}

func (m deployartifactsource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m deployartifactsource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DeployArtifactSourceDeployArtifactSourceTypeEnum Enum with underlying type: string
type DeployArtifactSourceDeployArtifactSourceTypeEnum string

// Set of constants representing the allowable values for DeployArtifactSourceDeployArtifactSourceTypeEnum
const (
	DeployArtifactSourceDeployArtifactSourceTypeInline          DeployArtifactSourceDeployArtifactSourceTypeEnum = "INLINE"
	DeployArtifactSourceDeployArtifactSourceTypeOcir            DeployArtifactSourceDeployArtifactSourceTypeEnum = "OCIR"
	DeployArtifactSourceDeployArtifactSourceTypeGenericArtifact DeployArtifactSourceDeployArtifactSourceTypeEnum = "GENERIC_ARTIFACT"
	DeployArtifactSourceDeployArtifactSourceTypeHelmChart       DeployArtifactSourceDeployArtifactSourceTypeEnum = "HELM_CHART"
	DeployArtifactSourceDeployArtifactSourceTypeHelmCommandSpec DeployArtifactSourceDeployArtifactSourceTypeEnum = "HELM_COMMAND_SPEC"
)

var mappingDeployArtifactSourceDeployArtifactSourceTypeEnum = map[string]DeployArtifactSourceDeployArtifactSourceTypeEnum{
	"INLINE":            DeployArtifactSourceDeployArtifactSourceTypeInline,
	"OCIR":              DeployArtifactSourceDeployArtifactSourceTypeOcir,
	"GENERIC_ARTIFACT":  DeployArtifactSourceDeployArtifactSourceTypeGenericArtifact,
	"HELM_CHART":        DeployArtifactSourceDeployArtifactSourceTypeHelmChart,
	"HELM_COMMAND_SPEC": DeployArtifactSourceDeployArtifactSourceTypeHelmCommandSpec,
}

var mappingDeployArtifactSourceDeployArtifactSourceTypeEnumLowerCase = map[string]DeployArtifactSourceDeployArtifactSourceTypeEnum{
	"inline":            DeployArtifactSourceDeployArtifactSourceTypeInline,
	"ocir":              DeployArtifactSourceDeployArtifactSourceTypeOcir,
	"generic_artifact":  DeployArtifactSourceDeployArtifactSourceTypeGenericArtifact,
	"helm_chart":        DeployArtifactSourceDeployArtifactSourceTypeHelmChart,
	"helm_command_spec": DeployArtifactSourceDeployArtifactSourceTypeHelmCommandSpec,
}

// GetDeployArtifactSourceDeployArtifactSourceTypeEnumValues Enumerates the set of values for DeployArtifactSourceDeployArtifactSourceTypeEnum
func GetDeployArtifactSourceDeployArtifactSourceTypeEnumValues() []DeployArtifactSourceDeployArtifactSourceTypeEnum {
	values := make([]DeployArtifactSourceDeployArtifactSourceTypeEnum, 0)
	for _, v := range mappingDeployArtifactSourceDeployArtifactSourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDeployArtifactSourceDeployArtifactSourceTypeEnumStringValues Enumerates the set of values in String for DeployArtifactSourceDeployArtifactSourceTypeEnum
func GetDeployArtifactSourceDeployArtifactSourceTypeEnumStringValues() []string {
	return []string{
		"INLINE",
		"OCIR",
		"GENERIC_ARTIFACT",
		"HELM_CHART",
		"HELM_COMMAND_SPEC",
	}
}

// GetMappingDeployArtifactSourceDeployArtifactSourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDeployArtifactSourceDeployArtifactSourceTypeEnum(val string) (DeployArtifactSourceDeployArtifactSourceTypeEnum, bool) {
	enum, ok := mappingDeployArtifactSourceDeployArtifactSourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
