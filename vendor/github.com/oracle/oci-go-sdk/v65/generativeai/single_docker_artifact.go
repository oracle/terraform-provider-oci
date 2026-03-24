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

// SingleDockerArtifact Container/artifact configuration for the deployment.
type SingleDockerArtifact struct {

	// if put artifact to a table, the id is needed
	Id *string `mandatory:"false" json:"id"`

	// The date and time the artifact was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the application.
	HostedDeploymentId *string `mandatory:"false" json:"hostedDeploymentId"`

	// image url.
	ContainerUri *string `mandatory:"false" json:"containerUri"`

	// image tag.
	Tag *string `mandatory:"false" json:"tag"`

	// The current status of the artifact.
	Status ArtifactStatusEnum `mandatory:"false" json:"status,omitempty"`
}

// GetId returns Id
func (m SingleDockerArtifact) GetId() *string {
	return m.Id
}

// GetTimeCreated returns TimeCreated
func (m SingleDockerArtifact) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetHostedDeploymentId returns HostedDeploymentId
func (m SingleDockerArtifact) GetHostedDeploymentId() *string {
	return m.HostedDeploymentId
}

// GetStatus returns Status
func (m SingleDockerArtifact) GetStatus() ArtifactStatusEnum {
	return m.Status
}

func (m SingleDockerArtifact) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SingleDockerArtifact) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingArtifactStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetArtifactStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m SingleDockerArtifact) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeSingleDockerArtifact SingleDockerArtifact
	s := struct {
		DiscriminatorParam string `json:"artifactType"`
		MarshalTypeSingleDockerArtifact
	}{
		"SIMPLE_DOCKER_ARTIFACT",
		(MarshalTypeSingleDockerArtifact)(m),
	}

	return json.Marshal(&s)
}
