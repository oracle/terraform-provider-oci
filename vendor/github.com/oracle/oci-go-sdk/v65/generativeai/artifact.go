// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generative AI Service Management API
//
// OCI Generative AI is a fully managed service that provides a set of state-of-the-art, customizable large language models (LLMs) that cover a wide range of use cases for text generation, summarization, and text embeddings.
// Use the Generative AI service management API to create and manage DedicatedAiCluster, Endpoint, Model, and WorkRequest in the Generative AI service. For example, create a custom model by fine-tuning an out-of-the-box model using your own data, on a fine-tuning dedicated AI cluster. Then, create a hosting dedicated AI cluster with an endpoint to host your custom model.
// To access your custom model endpoints, or to try the out-of-the-box models to generate text, summarize, and create text embeddings see the Generative AI Inference API (https://docs.oracle.com/iaas/api/#/en/generative-ai-inference/latest/).
// To learn more about the service, see the Generative AI documentation (https://docs.oracle.com/iaas/Content/generative-ai/home.htm).
//

package generativeai

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Artifact Container/artifact configuration for the deployment.
type Artifact interface {

	// if put artifact to a table, the id is needed
	GetId() *string

	// The date and time the artifact was created.
	GetTimeCreated() *common.SDKTime

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the application.
	GetHostedDeploymentId() *string

	// The current status of the artifact.
	GetStatus() ArtifactStatusEnum
}

type artifact struct {
	JsonData           []byte
	Id                 *string            `mandatory:"false" json:"id"`
	TimeCreated        *common.SDKTime    `mandatory:"false" json:"timeCreated"`
	HostedDeploymentId *string            `mandatory:"false" json:"hostedDeploymentId"`
	Status             ArtifactStatusEnum `mandatory:"false" json:"status,omitempty"`
	ArtifactType       string             `json:"artifactType"`
}

// UnmarshalJSON unmarshals json
func (m *artifact) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerartifact artifact
	s := struct {
		Model Unmarshalerartifact
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.TimeCreated = s.Model.TimeCreated
	m.HostedDeploymentId = s.Model.HostedDeploymentId
	m.Status = s.Model.Status
	m.ArtifactType = s.Model.ArtifactType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *artifact) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ArtifactType {
	case "SIMPLE_DOCKER_ARTIFACT":
		mm := SingleDockerArtifact{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for Artifact: %s.", m.ArtifactType)
		return *m, nil
	}
}

// GetId returns Id
func (m artifact) GetId() *string {
	return m.Id
}

// GetTimeCreated returns TimeCreated
func (m artifact) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetHostedDeploymentId returns HostedDeploymentId
func (m artifact) GetHostedDeploymentId() *string {
	return m.HostedDeploymentId
}

// GetStatus returns Status
func (m artifact) GetStatus() ArtifactStatusEnum {
	return m.Status
}

func (m artifact) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m artifact) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingArtifactStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetArtifactStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ArtifactStatusEnum Enum with underlying type: string
type ArtifactStatusEnum string

// Set of constants representing the allowable values for ArtifactStatusEnum
const (
	ArtifactStatusActive   ArtifactStatusEnum = "ACTIVE"
	ArtifactStatusInactive ArtifactStatusEnum = "INACTIVE"
	ArtifactStatusUpdating ArtifactStatusEnum = "UPDATING"
)

var mappingArtifactStatusEnum = map[string]ArtifactStatusEnum{
	"ACTIVE":   ArtifactStatusActive,
	"INACTIVE": ArtifactStatusInactive,
	"UPDATING": ArtifactStatusUpdating,
}

var mappingArtifactStatusEnumLowerCase = map[string]ArtifactStatusEnum{
	"active":   ArtifactStatusActive,
	"inactive": ArtifactStatusInactive,
	"updating": ArtifactStatusUpdating,
}

// GetArtifactStatusEnumValues Enumerates the set of values for ArtifactStatusEnum
func GetArtifactStatusEnumValues() []ArtifactStatusEnum {
	values := make([]ArtifactStatusEnum, 0)
	for _, v := range mappingArtifactStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetArtifactStatusEnumStringValues Enumerates the set of values in String for ArtifactStatusEnum
func GetArtifactStatusEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
	}
}

// GetMappingArtifactStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingArtifactStatusEnum(val string) (ArtifactStatusEnum, bool) {
	enum, ok := mappingArtifactStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ArtifactArtifactTypeEnum Enum with underlying type: string
type ArtifactArtifactTypeEnum string

// Set of constants representing the allowable values for ArtifactArtifactTypeEnum
const (
	ArtifactArtifactTypeSimpleDockerArtifact ArtifactArtifactTypeEnum = "SIMPLE_DOCKER_ARTIFACT"
)

var mappingArtifactArtifactTypeEnum = map[string]ArtifactArtifactTypeEnum{
	"SIMPLE_DOCKER_ARTIFACT": ArtifactArtifactTypeSimpleDockerArtifact,
}

var mappingArtifactArtifactTypeEnumLowerCase = map[string]ArtifactArtifactTypeEnum{
	"simple_docker_artifact": ArtifactArtifactTypeSimpleDockerArtifact,
}

// GetArtifactArtifactTypeEnumValues Enumerates the set of values for ArtifactArtifactTypeEnum
func GetArtifactArtifactTypeEnumValues() []ArtifactArtifactTypeEnum {
	values := make([]ArtifactArtifactTypeEnum, 0)
	for _, v := range mappingArtifactArtifactTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetArtifactArtifactTypeEnumStringValues Enumerates the set of values in String for ArtifactArtifactTypeEnum
func GetArtifactArtifactTypeEnumStringValues() []string {
	return []string{
		"SIMPLE_DOCKER_ARTIFACT",
	}
}

// GetMappingArtifactArtifactTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingArtifactArtifactTypeEnum(val string) (ArtifactArtifactTypeEnum, bool) {
	enum, ok := mappingArtifactArtifactTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
