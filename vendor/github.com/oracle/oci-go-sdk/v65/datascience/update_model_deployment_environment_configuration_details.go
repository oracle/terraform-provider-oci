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

// UpdateModelDeploymentEnvironmentConfigurationDetails The configuration to carry the environment details thats used in Model Deployment update
type UpdateModelDeploymentEnvironmentConfigurationDetails interface {
}

type updatemodeldeploymentenvironmentconfigurationdetails struct {
	JsonData                     []byte
	EnvironmentConfigurationType string `json:"environmentConfigurationType"`
}

// UnmarshalJSON unmarshals json
func (m *updatemodeldeploymentenvironmentconfigurationdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdatemodeldeploymentenvironmentconfigurationdetails updatemodeldeploymentenvironmentconfigurationdetails
	s := struct {
		Model Unmarshalerupdatemodeldeploymentenvironmentconfigurationdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.EnvironmentConfigurationType = s.Model.EnvironmentConfigurationType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updatemodeldeploymentenvironmentconfigurationdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.EnvironmentConfigurationType {
	case "DEFAULT":
		mm := UpdateDefaultModelDeploymentEnvironmentConfigurationDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCIR_CONTAINER":
		mm := UpdateOcirModelDeploymentEnvironmentConfigurationDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for UpdateModelDeploymentEnvironmentConfigurationDetails: %s.", m.EnvironmentConfigurationType)
		return *m, nil
	}
}

func (m updatemodeldeploymentenvironmentconfigurationdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updatemodeldeploymentenvironmentconfigurationdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
