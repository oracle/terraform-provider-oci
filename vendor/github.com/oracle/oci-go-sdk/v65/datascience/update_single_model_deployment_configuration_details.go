// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateSingleModelDeploymentConfigurationDetails The single model type deployment for update.
type UpdateSingleModelDeploymentConfigurationDetails struct {
	ModelConfigurationDetails *UpdateModelConfigurationDetails `mandatory:"false" json:"modelConfigurationDetails"`

	EnvironmentConfigurationDetails UpdateModelDeploymentEnvironmentConfigurationDetails `mandatory:"false" json:"environmentConfigurationDetails"`
}

func (m UpdateSingleModelDeploymentConfigurationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateSingleModelDeploymentConfigurationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateSingleModelDeploymentConfigurationDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateSingleModelDeploymentConfigurationDetails UpdateSingleModelDeploymentConfigurationDetails
	s := struct {
		DiscriminatorParam string `json:"deploymentType"`
		MarshalTypeUpdateSingleModelDeploymentConfigurationDetails
	}{
		"SINGLE_MODEL",
		(MarshalTypeUpdateSingleModelDeploymentConfigurationDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *UpdateSingleModelDeploymentConfigurationDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ModelConfigurationDetails       *UpdateModelConfigurationDetails                     `json:"modelConfigurationDetails"`
		EnvironmentConfigurationDetails updatemodeldeploymentenvironmentconfigurationdetails `json:"environmentConfigurationDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.ModelConfigurationDetails = model.ModelConfigurationDetails

	nn, e = model.EnvironmentConfigurationDetails.UnmarshalPolymorphicJSON(model.EnvironmentConfigurationDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.EnvironmentConfigurationDetails = nn.(UpdateModelDeploymentEnvironmentConfigurationDetails)
	} else {
		m.EnvironmentConfigurationDetails = nil
	}

	return
}
