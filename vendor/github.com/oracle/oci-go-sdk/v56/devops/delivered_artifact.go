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
		return *m, nil
	}
}

//GetDeployArtifactId returns DeployArtifactId
func (m deliveredartifact) GetDeployArtifactId() *string {
	return m.DeployArtifactId
}

//GetOutputArtifactName returns OutputArtifactName
func (m deliveredartifact) GetOutputArtifactName() *string {
	return m.OutputArtifactName
}

func (m deliveredartifact) String() string {
	return common.PointerString(m)
}

// DeliveredArtifactArtifactTypeEnum Enum with underlying type: string
type DeliveredArtifactArtifactTypeEnum string

// Set of constants representing the allowable values for DeliveredArtifactArtifactTypeEnum
const (
	DeliveredArtifactArtifactTypeGenericArtifact DeliveredArtifactArtifactTypeEnum = "GENERIC_ARTIFACT"
	DeliveredArtifactArtifactTypeOcir            DeliveredArtifactArtifactTypeEnum = "OCIR"
)

var mappingDeliveredArtifactArtifactType = map[string]DeliveredArtifactArtifactTypeEnum{
	"GENERIC_ARTIFACT": DeliveredArtifactArtifactTypeGenericArtifact,
	"OCIR":             DeliveredArtifactArtifactTypeOcir,
}

// GetDeliveredArtifactArtifactTypeEnumValues Enumerates the set of values for DeliveredArtifactArtifactTypeEnum
func GetDeliveredArtifactArtifactTypeEnumValues() []DeliveredArtifactArtifactTypeEnum {
	values := make([]DeliveredArtifactArtifactTypeEnum, 0)
	for _, v := range mappingDeliveredArtifactArtifactType {
		values = append(values, v)
	}
	return values
}
