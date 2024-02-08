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
	case "ME_STANDALONE":
		mm := ManagedEgressStandaloneJobInfrastructureConfigurationDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "STANDALONE":
		mm := StandaloneJobInfrastructureConfigurationDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for JobInfrastructureConfigurationDetails: %s.", m.JobInfrastructureType)
		return *m, nil
	}
}

func (m jobinfrastructureconfigurationdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m jobinfrastructureconfigurationdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// JobInfrastructureConfigurationDetailsJobInfrastructureTypeEnum Enum with underlying type: string
type JobInfrastructureConfigurationDetailsJobInfrastructureTypeEnum string

// Set of constants representing the allowable values for JobInfrastructureConfigurationDetailsJobInfrastructureTypeEnum
const (
	JobInfrastructureConfigurationDetailsJobInfrastructureTypeStandalone   JobInfrastructureConfigurationDetailsJobInfrastructureTypeEnum = "STANDALONE"
	JobInfrastructureConfigurationDetailsJobInfrastructureTypeMeStandalone JobInfrastructureConfigurationDetailsJobInfrastructureTypeEnum = "ME_STANDALONE"
)

var mappingJobInfrastructureConfigurationDetailsJobInfrastructureTypeEnum = map[string]JobInfrastructureConfigurationDetailsJobInfrastructureTypeEnum{
	"STANDALONE":    JobInfrastructureConfigurationDetailsJobInfrastructureTypeStandalone,
	"ME_STANDALONE": JobInfrastructureConfigurationDetailsJobInfrastructureTypeMeStandalone,
}

var mappingJobInfrastructureConfigurationDetailsJobInfrastructureTypeEnumLowerCase = map[string]JobInfrastructureConfigurationDetailsJobInfrastructureTypeEnum{
	"standalone":    JobInfrastructureConfigurationDetailsJobInfrastructureTypeStandalone,
	"me_standalone": JobInfrastructureConfigurationDetailsJobInfrastructureTypeMeStandalone,
}

// GetJobInfrastructureConfigurationDetailsJobInfrastructureTypeEnumValues Enumerates the set of values for JobInfrastructureConfigurationDetailsJobInfrastructureTypeEnum
func GetJobInfrastructureConfigurationDetailsJobInfrastructureTypeEnumValues() []JobInfrastructureConfigurationDetailsJobInfrastructureTypeEnum {
	values := make([]JobInfrastructureConfigurationDetailsJobInfrastructureTypeEnum, 0)
	for _, v := range mappingJobInfrastructureConfigurationDetailsJobInfrastructureTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetJobInfrastructureConfigurationDetailsJobInfrastructureTypeEnumStringValues Enumerates the set of values in String for JobInfrastructureConfigurationDetailsJobInfrastructureTypeEnum
func GetJobInfrastructureConfigurationDetailsJobInfrastructureTypeEnumStringValues() []string {
	return []string{
		"STANDALONE",
		"ME_STANDALONE",
	}
}

// GetMappingJobInfrastructureConfigurationDetailsJobInfrastructureTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJobInfrastructureConfigurationDetailsJobInfrastructureTypeEnum(val string) (JobInfrastructureConfigurationDetailsJobInfrastructureTypeEnum, bool) {
	enum, ok := mappingJobInfrastructureConfigurationDetailsJobInfrastructureTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
