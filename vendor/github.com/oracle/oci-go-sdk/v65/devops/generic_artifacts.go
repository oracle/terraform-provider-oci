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

// GenericArtifacts Details of artifact generated via pipeline run
type GenericArtifacts struct {

	// Name of stage step at which this output is generated.
	StepName *string `mandatory:"true" json:"stepName"`

	// Name of artifact.
	Name *string `mandatory:"true" json:"name"`

	LocationDetails GenericArtifactLocationDetails `mandatory:"true" json:"locationDetails"`
}

// GetStepName returns StepName
func (m GenericArtifacts) GetStepName() *string {
	return m.StepName
}

func (m GenericArtifacts) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GenericArtifacts) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m GenericArtifacts) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeGenericArtifacts GenericArtifacts
	s := struct {
		DiscriminatorParam string `json:"outputType"`
		MarshalTypeGenericArtifacts
	}{
		"ARTIFACT",
		(MarshalTypeGenericArtifacts)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *GenericArtifacts) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		StepName        *string                        `json:"stepName"`
		Name            *string                        `json:"name"`
		LocationDetails genericartifactlocationdetails `json:"locationDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.StepName = model.StepName

	m.Name = model.Name

	nn, e = model.LocationDetails.UnmarshalPolymorphicJSON(model.LocationDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.LocationDetails = nn.(GenericArtifactLocationDetails)
	} else {
		m.LocationDetails = nil
	}

	return
}
