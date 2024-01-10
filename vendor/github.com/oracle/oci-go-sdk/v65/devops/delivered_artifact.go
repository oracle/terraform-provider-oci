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

// DeliveredArtifact Details of the artifacts delivered through the Deliver Artifacts stage.
type DeliveredArtifact interface {

	// The OCID of the deployment artifact definition.
	GetDeployArtifactId() *string

	// Name of the output artifact defined in the build specification file.
	GetOutputArtifactName() *string
}

type deliveredartifact struct {
	JsonData           []byte
	DeployArtifactId   *string `mandatory:"true" json:"deployArtifactId"`
	OutputArtifactName *string `mandatory:"true" json:"outputArtifactName"`
	ArtifactType       string  `json:"artifactType"`
}

// UnmarshalJSON unmarshals json
func (m *deliveredartifact) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdeliveredartifact deliveredartifact
	s := struct {
		Model Unmarshalerdeliveredartifact
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DeployArtifactId = s.Model.DeployArtifactId
	m.OutputArtifactName = s.Model.OutputArtifactName
	m.ArtifactType = s.Model.ArtifactType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *deliveredartifact) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ArtifactType {
	case "OCIR":
		mm := ContainerRegistryDeliveredArtifact{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GENERIC_ARTIFACT":
		mm := GenericDeliveredArtifact{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for DeliveredArtifact: %s.", m.ArtifactType)
		return *m, nil
	}
}

// GetDeployArtifactId returns DeployArtifactId
func (m deliveredartifact) GetDeployArtifactId() *string {
	return m.DeployArtifactId
}

// GetOutputArtifactName returns OutputArtifactName
func (m deliveredartifact) GetOutputArtifactName() *string {
	return m.OutputArtifactName
}

func (m deliveredartifact) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m deliveredartifact) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DeliveredArtifactArtifactTypeEnum Enum with underlying type: string
type DeliveredArtifactArtifactTypeEnum string

// Set of constants representing the allowable values for DeliveredArtifactArtifactTypeEnum
const (
	DeliveredArtifactArtifactTypeGenericArtifact DeliveredArtifactArtifactTypeEnum = "GENERIC_ARTIFACT"
	DeliveredArtifactArtifactTypeOcir            DeliveredArtifactArtifactTypeEnum = "OCIR"
)

var mappingDeliveredArtifactArtifactTypeEnum = map[string]DeliveredArtifactArtifactTypeEnum{
	"GENERIC_ARTIFACT": DeliveredArtifactArtifactTypeGenericArtifact,
	"OCIR":             DeliveredArtifactArtifactTypeOcir,
}

var mappingDeliveredArtifactArtifactTypeEnumLowerCase = map[string]DeliveredArtifactArtifactTypeEnum{
	"generic_artifact": DeliveredArtifactArtifactTypeGenericArtifact,
	"ocir":             DeliveredArtifactArtifactTypeOcir,
}

// GetDeliveredArtifactArtifactTypeEnumValues Enumerates the set of values for DeliveredArtifactArtifactTypeEnum
func GetDeliveredArtifactArtifactTypeEnumValues() []DeliveredArtifactArtifactTypeEnum {
	values := make([]DeliveredArtifactArtifactTypeEnum, 0)
	for _, v := range mappingDeliveredArtifactArtifactTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDeliveredArtifactArtifactTypeEnumStringValues Enumerates the set of values in String for DeliveredArtifactArtifactTypeEnum
func GetDeliveredArtifactArtifactTypeEnumStringValues() []string {
	return []string{
		"GENERIC_ARTIFACT",
		"OCIR",
	}
}

// GetMappingDeliveredArtifactArtifactTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDeliveredArtifactArtifactTypeEnum(val string) (DeliveredArtifactArtifactTypeEnum, bool) {
	enum, ok := mappingDeliveredArtifactArtifactTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
