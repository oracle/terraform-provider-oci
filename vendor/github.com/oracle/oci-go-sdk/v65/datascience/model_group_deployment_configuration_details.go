// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// ModelGroupDeploymentConfigurationDetails The model group type deployment.
type ModelGroupDeploymentConfigurationDetails struct {
	ModelGroupConfigurationDetails *ModelGroupConfigurationDetails `mandatory:"true" json:"modelGroupConfigurationDetails"`

	InfrastructureConfigurationDetails InfrastructureConfigurationDetails `mandatory:"true" json:"infrastructureConfigurationDetails"`

	EnvironmentConfigurationDetails ModelDeploymentEnvironmentConfigurationDetails `mandatory:"false" json:"environmentConfigurationDetails"`
}

func (m ModelGroupDeploymentConfigurationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ModelGroupDeploymentConfigurationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ModelGroupDeploymentConfigurationDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeModelGroupDeploymentConfigurationDetails ModelGroupDeploymentConfigurationDetails
	s := struct {
		DiscriminatorParam string `json:"deploymentType"`
		MarshalTypeModelGroupDeploymentConfigurationDetails
	}{
		"MODEL_GROUP",
		(MarshalTypeModelGroupDeploymentConfigurationDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *ModelGroupDeploymentConfigurationDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		EnvironmentConfigurationDetails    modeldeploymentenvironmentconfigurationdetails `json:"environmentConfigurationDetails"`
		ModelGroupConfigurationDetails     *ModelGroupConfigurationDetails                `json:"modelGroupConfigurationDetails"`
		InfrastructureConfigurationDetails infrastructureconfigurationdetails             `json:"infrastructureConfigurationDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.EnvironmentConfigurationDetails.UnmarshalPolymorphicJSON(model.EnvironmentConfigurationDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.EnvironmentConfigurationDetails = nn.(ModelDeploymentEnvironmentConfigurationDetails)
	} else {
		m.EnvironmentConfigurationDetails = nil
	}

	m.ModelGroupConfigurationDetails = model.ModelGroupConfigurationDetails

	nn, e = model.InfrastructureConfigurationDetails.UnmarshalPolymorphicJSON(model.InfrastructureConfigurationDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.InfrastructureConfigurationDetails = nn.(InfrastructureConfigurationDetails)
	} else {
		m.InfrastructureConfigurationDetails = nil
	}

	return
}
