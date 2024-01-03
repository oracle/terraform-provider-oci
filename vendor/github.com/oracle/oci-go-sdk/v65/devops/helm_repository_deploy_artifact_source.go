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

// HelmRepositoryDeployArtifactSource Specifies Helm chart source details.
type HelmRepositoryDeployArtifactSource struct {

	// The URL of an OCIR repository.
	ChartUrl *string `mandatory:"true" json:"chartUrl"`

	// Users can set this as a placeholder value that refers to a pipeline parameter, for example, ${appVersion}.
	DeployArtifactVersion *string `mandatory:"true" json:"deployArtifactVersion"`

	HelmVerificationKeySource VerificationKeySource `mandatory:"false" json:"helmVerificationKeySource"`
}

func (m HelmRepositoryDeployArtifactSource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HelmRepositoryDeployArtifactSource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m HelmRepositoryDeployArtifactSource) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeHelmRepositoryDeployArtifactSource HelmRepositoryDeployArtifactSource
	s := struct {
		DiscriminatorParam string `json:"deployArtifactSourceType"`
		MarshalTypeHelmRepositoryDeployArtifactSource
	}{
		"HELM_CHART",
		(MarshalTypeHelmRepositoryDeployArtifactSource)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *HelmRepositoryDeployArtifactSource) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		HelmVerificationKeySource verificationkeysource `json:"helmVerificationKeySource"`
		ChartUrl                  *string               `json:"chartUrl"`
		DeployArtifactVersion     *string               `json:"deployArtifactVersion"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.HelmVerificationKeySource.UnmarshalPolymorphicJSON(model.HelmVerificationKeySource.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.HelmVerificationKeySource = nn.(VerificationKeySource)
	} else {
		m.HelmVerificationKeySource = nil
	}

	m.ChartUrl = model.ChartUrl

	m.DeployArtifactVersion = model.DeployArtifactVersion

	return
}
