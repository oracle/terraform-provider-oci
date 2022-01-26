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

// GenericDeliveredArtifact Details of the generic artifacts delivered through the Deliver Artifacts stage.
type GenericDeliveredArtifact struct {

	// The OCID of the deployment artifact definition.
	DeployArtifactId *string `mandatory:"true" json:"deployArtifactId"`

	// Name of the output artifact defined in the build specification file.
	OutputArtifactName *string `mandatory:"true" json:"outputArtifactName"`

	// The OCID of the artifact pushed by the Deliver Artifacts stage.
	DeliveredArtifactId *string `mandatory:"true" json:"deliveredArtifactId"`

	// The OCID of the artifact registry repository used by the DeliverArtifactStage
	ArtifactRepositoryId *string `mandatory:"false" json:"artifactRepositoryId"`

	// Path of the repository where artifact was pushed
	Path *string `mandatory:"false" json:"path"`

	// Version of the artifact pushed
	Version *string `mandatory:"false" json:"version"`
}

//GetDeployArtifactId returns DeployArtifactId
func (m GenericDeliveredArtifact) GetDeployArtifactId() *string {
	return m.DeployArtifactId
}

//GetOutputArtifactName returns OutputArtifactName
func (m GenericDeliveredArtifact) GetOutputArtifactName() *string {
	return m.OutputArtifactName
}

func (m GenericDeliveredArtifact) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m GenericDeliveredArtifact) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeGenericDeliveredArtifact GenericDeliveredArtifact
	s := struct {
		DiscriminatorParam string `json:"artifactType"`
		MarshalTypeGenericDeliveredArtifact
	}{
		"GENERIC_ARTIFACT",
		(MarshalTypeGenericDeliveredArtifact)(m),
	}

	return json.Marshal(&s)
}
