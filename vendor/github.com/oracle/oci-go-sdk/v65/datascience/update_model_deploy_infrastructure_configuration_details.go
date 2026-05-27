// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// UpdateModelDeployInfrastructureConfigurationDetails The Infrastructure configuration details for updating a model deployment.
type UpdateModelDeployInfrastructureConfigurationDetails interface {
}

type updatemodeldeployinfrastructureconfigurationdetails struct {
	JsonData           []byte
	InfrastructureType string `json:"infrastructureType"`
}

// UnmarshalJSON unmarshals json
func (m *updatemodeldeployinfrastructureconfigurationdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdatemodeldeployinfrastructureconfigurationdetails updatemodeldeployinfrastructureconfigurationdetails
	s := struct {
		Model Unmarshalerupdatemodeldeployinfrastructureconfigurationdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.InfrastructureType = s.Model.InfrastructureType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updatemodeldeployinfrastructureconfigurationdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.InfrastructureType {
	case "MANAGED_COMPUTE_CLUSTER":
		mm := UpdateManagedComputeClusterModelDeployInfrastructureConfigDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for UpdateModelDeployInfrastructureConfigurationDetails: %s.", m.InfrastructureType)
		return *m, nil
	}
}

func (m updatemodeldeployinfrastructureconfigurationdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updatemodeldeployinfrastructureconfigurationdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
