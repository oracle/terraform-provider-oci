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

// GenericDeployArtifactSource Specifies the Artifact Registry source details.
type GenericDeployArtifactSource struct {

	// The OCID of a repository
	RepositoryId *string `mandatory:"true" json:"repositoryId"`

	// Specifies the artifact path in the repository.
	DeployArtifactPath *string `mandatory:"true" json:"deployArtifactPath"`

	// Users can set this as a placeholder value that refers to a pipeline parameter, for example, ${appVersion}.
	DeployArtifactVersion *string `mandatory:"true" json:"deployArtifactVersion"`
}

func (m GenericDeployArtifactSource) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m GenericDeployArtifactSource) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeGenericDeployArtifactSource GenericDeployArtifactSource
	s := struct {
		DiscriminatorParam string `json:"deployArtifactSourceType"`
		MarshalTypeGenericDeployArtifactSource
	}{
		"GENERIC_ARTIFACT",
		(MarshalTypeGenericDeployArtifactSource)(m),
	}

	return json.Marshal(&s)
}
