// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps APIs to create a DevOps project to group the pipelines,  add reference to target deployment environments, add artifacts to deploy,  and create deployment pipelines needed to deploy your software.
//

package devops

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v50/common"
)

// ContainerRegistryDeliveredArtifact Details of OCIR Artifact delivered via DeliverArtifactStage.
type ContainerRegistryDeliveredArtifact struct {

	// The OCID of the deploy artifact definition
	DeployArtifactId *string `mandatory:"true" json:"deployArtifactId"`

	// Name of the output artifact defined in the build spec
	OutputArtifactName *string `mandatory:"true" json:"outputArtifactName"`

	// The Hash of the OCIR artifact pushed by the DeliverArtifactStage
	DeliveredArtifactHash *string `mandatory:"true" json:"deliveredArtifactHash"`

	// The imageUri of the OCIR artifact pushed by the DeliverArtifactStage
	ImageUri *string `mandatory:"false" json:"imageUri"`
}

//GetDeployArtifactId returns DeployArtifactId
func (m ContainerRegistryDeliveredArtifact) GetDeployArtifactId() *string {
	return m.DeployArtifactId
}

//GetOutputArtifactName returns OutputArtifactName
func (m ContainerRegistryDeliveredArtifact) GetOutputArtifactName() *string {
	return m.OutputArtifactName
}

func (m ContainerRegistryDeliveredArtifact) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ContainerRegistryDeliveredArtifact) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeContainerRegistryDeliveredArtifact ContainerRegistryDeliveredArtifact
	s := struct {
		DiscriminatorParam string `json:"artifactType"`
		MarshalTypeContainerRegistryDeliveredArtifact
	}{
		"OCIR",
		(MarshalTypeContainerRegistryDeliveredArtifact)(m),
	}

	return json.Marshal(&s)
}
