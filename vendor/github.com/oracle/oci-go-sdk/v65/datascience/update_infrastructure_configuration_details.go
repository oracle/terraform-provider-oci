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

// UpdateInfrastructureConfigurationDetails The infrastructure configuration details for update.
type UpdateInfrastructureConfigurationDetails interface {
}

type updateinfrastructureconfigurationdetails struct {
	JsonData           []byte
	InfrastructureType string `json:"infrastructureType"`
}

// UnmarshalJSON unmarshals json
func (m *updateinfrastructureconfigurationdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdateinfrastructureconfigurationdetails updateinfrastructureconfigurationdetails
	s := struct {
		Model Unmarshalerupdateinfrastructureconfigurationdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.InfrastructureType = s.Model.InfrastructureType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updateinfrastructureconfigurationdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.InfrastructureType {
	case "INSTANCE_POOL":
		mm := UpdateInstancePoolInfrastructureConfigurationDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for UpdateInfrastructureConfigurationDetails: %s.", m.InfrastructureType)
		return *m, nil
	}
}

func (m updateinfrastructureconfigurationdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updateinfrastructureconfigurationdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
