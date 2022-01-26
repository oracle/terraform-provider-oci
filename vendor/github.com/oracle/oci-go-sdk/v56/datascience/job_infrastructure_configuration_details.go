// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// JobInfrastructureConfigurationDetails The job infrastructure configuration details (shape, block storage, etc.)
type JobInfrastructureConfigurationDetails interface {
}

type jobinfrastructureconfigurationdetails struct {
	JsonData              []byte
	JobInfrastructureType string `json:"jobInfrastructureType"`
}

// UnmarshalJSON unmarshals json
func (m *jobinfrastructureconfigurationdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerjobinfrastructureconfigurationdetails jobinfrastructureconfigurationdetails
	s := struct {
		Model Unmarshalerjobinfrastructureconfigurationdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.JobInfrastructureType = s.Model.JobInfrastructureType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *jobinfrastructureconfigurationdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.JobInfrastructureType {
	case "STANDALONE":
		mm := StandaloneJobInfrastructureConfigurationDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m jobinfrastructureconfigurationdetails) String() string {
	return common.PointerString(m)
}

// JobInfrastructureConfigurationDetailsJobInfrastructureTypeEnum Enum with underlying type: string
type JobInfrastructureConfigurationDetailsJobInfrastructureTypeEnum string

// Set of constants representing the allowable values for JobInfrastructureConfigurationDetailsJobInfrastructureTypeEnum
const (
	JobInfrastructureConfigurationDetailsJobInfrastructureTypeStandalone JobInfrastructureConfigurationDetailsJobInfrastructureTypeEnum = "STANDALONE"
)

var mappingJobInfrastructureConfigurationDetailsJobInfrastructureType = map[string]JobInfrastructureConfigurationDetailsJobInfrastructureTypeEnum{
	"STANDALONE": JobInfrastructureConfigurationDetailsJobInfrastructureTypeStandalone,
}

// GetJobInfrastructureConfigurationDetailsJobInfrastructureTypeEnumValues Enumerates the set of values for JobInfrastructureConfigurationDetailsJobInfrastructureTypeEnum
func GetJobInfrastructureConfigurationDetailsJobInfrastructureTypeEnumValues() []JobInfrastructureConfigurationDetailsJobInfrastructureTypeEnum {
	values := make([]JobInfrastructureConfigurationDetailsJobInfrastructureTypeEnum, 0)
	for _, v := range mappingJobInfrastructureConfigurationDetailsJobInfrastructureType {
		values = append(values, v)
	}
	return values
}
