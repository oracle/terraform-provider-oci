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

// JobEnvironmentConfigurationDetails Environment configuration to capture job runtime dependencies.
type JobEnvironmentConfigurationDetails interface {
}

type jobenvironmentconfigurationdetails struct {
	JsonData           []byte
	JobEnvironmentType string `json:"jobEnvironmentType"`
}

// UnmarshalJSON unmarshals json
func (m *jobenvironmentconfigurationdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerjobenvironmentconfigurationdetails jobenvironmentconfigurationdetails
	s := struct {
		Model Unmarshalerjobenvironmentconfigurationdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.JobEnvironmentType = s.Model.JobEnvironmentType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *jobenvironmentconfigurationdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.JobEnvironmentType {
	case "OCIR_CONTAINER":
		mm := OcirContainerJobEnvironmentConfigurationDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for JobEnvironmentConfigurationDetails: %s.", m.JobEnvironmentType)
		return *m, nil
	}
}

func (m jobenvironmentconfigurationdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m jobenvironmentconfigurationdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// JobEnvironmentConfigurationDetailsJobEnvironmentTypeEnum Enum with underlying type: string
type JobEnvironmentConfigurationDetailsJobEnvironmentTypeEnum string

// Set of constants representing the allowable values for JobEnvironmentConfigurationDetailsJobEnvironmentTypeEnum
const (
	JobEnvironmentConfigurationDetailsJobEnvironmentTypeOcirContainer JobEnvironmentConfigurationDetailsJobEnvironmentTypeEnum = "OCIR_CONTAINER"
)

var mappingJobEnvironmentConfigurationDetailsJobEnvironmentTypeEnum = map[string]JobEnvironmentConfigurationDetailsJobEnvironmentTypeEnum{
	"OCIR_CONTAINER": JobEnvironmentConfigurationDetailsJobEnvironmentTypeOcirContainer,
}

var mappingJobEnvironmentConfigurationDetailsJobEnvironmentTypeEnumLowerCase = map[string]JobEnvironmentConfigurationDetailsJobEnvironmentTypeEnum{
	"ocir_container": JobEnvironmentConfigurationDetailsJobEnvironmentTypeOcirContainer,
}

// GetJobEnvironmentConfigurationDetailsJobEnvironmentTypeEnumValues Enumerates the set of values for JobEnvironmentConfigurationDetailsJobEnvironmentTypeEnum
func GetJobEnvironmentConfigurationDetailsJobEnvironmentTypeEnumValues() []JobEnvironmentConfigurationDetailsJobEnvironmentTypeEnum {
	values := make([]JobEnvironmentConfigurationDetailsJobEnvironmentTypeEnum, 0)
	for _, v := range mappingJobEnvironmentConfigurationDetailsJobEnvironmentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetJobEnvironmentConfigurationDetailsJobEnvironmentTypeEnumStringValues Enumerates the set of values in String for JobEnvironmentConfigurationDetailsJobEnvironmentTypeEnum
func GetJobEnvironmentConfigurationDetailsJobEnvironmentTypeEnumStringValues() []string {
	return []string{
		"OCIR_CONTAINER",
	}
}

// GetMappingJobEnvironmentConfigurationDetailsJobEnvironmentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJobEnvironmentConfigurationDetailsJobEnvironmentTypeEnum(val string) (JobEnvironmentConfigurationDetailsJobEnvironmentTypeEnum, bool) {
	enum, ok := mappingJobEnvironmentConfigurationDetailsJobEnvironmentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
