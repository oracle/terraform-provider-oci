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

// JobRunUseCaseConfigurationDetails The use-case configuration details
type JobRunUseCaseConfigurationDetails interface {
}

type jobrunusecaseconfigurationdetails struct {
	JsonData    []byte
	UseCaseType string `json:"useCaseType"`
}

// UnmarshalJSON unmarshals json
func (m *jobrunusecaseconfigurationdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerjobrunusecaseconfigurationdetails jobrunusecaseconfigurationdetails
	s := struct {
		Model Unmarshalerjobrunusecaseconfigurationdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.UseCaseType = s.Model.UseCaseType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *jobrunusecaseconfigurationdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.UseCaseType {
	case "GENERIC":
		mm := GenericJobRunUseCaseConfigurationDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for JobRunUseCaseConfigurationDetails: %s.", m.UseCaseType)
		return *m, nil
	}
}

func (m jobrunusecaseconfigurationdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m jobrunusecaseconfigurationdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// JobRunUseCaseConfigurationDetailsUseCaseTypeEnum Enum with underlying type: string
type JobRunUseCaseConfigurationDetailsUseCaseTypeEnum string

// Set of constants representing the allowable values for JobRunUseCaseConfigurationDetailsUseCaseTypeEnum
const (
	JobRunUseCaseConfigurationDetailsUseCaseTypeGeneric JobRunUseCaseConfigurationDetailsUseCaseTypeEnum = "GENERIC"
)

var mappingJobRunUseCaseConfigurationDetailsUseCaseTypeEnum = map[string]JobRunUseCaseConfigurationDetailsUseCaseTypeEnum{
	"GENERIC": JobRunUseCaseConfigurationDetailsUseCaseTypeGeneric,
}

var mappingJobRunUseCaseConfigurationDetailsUseCaseTypeEnumLowerCase = map[string]JobRunUseCaseConfigurationDetailsUseCaseTypeEnum{
	"generic": JobRunUseCaseConfigurationDetailsUseCaseTypeGeneric,
}

// GetJobRunUseCaseConfigurationDetailsUseCaseTypeEnumValues Enumerates the set of values for JobRunUseCaseConfigurationDetailsUseCaseTypeEnum
func GetJobRunUseCaseConfigurationDetailsUseCaseTypeEnumValues() []JobRunUseCaseConfigurationDetailsUseCaseTypeEnum {
	values := make([]JobRunUseCaseConfigurationDetailsUseCaseTypeEnum, 0)
	for _, v := range mappingJobRunUseCaseConfigurationDetailsUseCaseTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetJobRunUseCaseConfigurationDetailsUseCaseTypeEnumStringValues Enumerates the set of values in String for JobRunUseCaseConfigurationDetailsUseCaseTypeEnum
func GetJobRunUseCaseConfigurationDetailsUseCaseTypeEnumStringValues() []string {
	return []string{
		"GENERIC",
	}
}

// GetMappingJobRunUseCaseConfigurationDetailsUseCaseTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJobRunUseCaseConfigurationDetailsUseCaseTypeEnum(val string) (JobRunUseCaseConfigurationDetailsUseCaseTypeEnum, bool) {
	enum, ok := mappingJobRunUseCaseConfigurationDetailsUseCaseTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
