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
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// JobConfigurationDetails The job configuration details
type JobConfigurationDetails interface {
}

type jobconfigurationdetails struct {
	JsonData []byte
	JobType  string `json:"jobType"`
}

// UnmarshalJSON unmarshals json
func (m *jobconfigurationdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerjobconfigurationdetails jobconfigurationdetails
	s := struct {
		Model Unmarshalerjobconfigurationdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.JobType = s.Model.JobType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *jobconfigurationdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.JobType {
	case "DEFAULT":
		mm := DefaultJobConfigurationDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m jobconfigurationdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m jobconfigurationdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// JobConfigurationDetailsJobTypeEnum Enum with underlying type: string
type JobConfigurationDetailsJobTypeEnum string

// Set of constants representing the allowable values for JobConfigurationDetailsJobTypeEnum
const (
	JobConfigurationDetailsJobTypeDefault JobConfigurationDetailsJobTypeEnum = "DEFAULT"
)

var mappingJobConfigurationDetailsJobTypeEnum = map[string]JobConfigurationDetailsJobTypeEnum{
	"DEFAULT": JobConfigurationDetailsJobTypeDefault,
}

// GetJobConfigurationDetailsJobTypeEnumValues Enumerates the set of values for JobConfigurationDetailsJobTypeEnum
func GetJobConfigurationDetailsJobTypeEnumValues() []JobConfigurationDetailsJobTypeEnum {
	values := make([]JobConfigurationDetailsJobTypeEnum, 0)
	for _, v := range mappingJobConfigurationDetailsJobTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetJobConfigurationDetailsJobTypeEnumStringValues Enumerates the set of values in String for JobConfigurationDetailsJobTypeEnum
func GetJobConfigurationDetailsJobTypeEnumStringValues() []string {
	return []string{
		"DEFAULT",
	}
}

// GetMappingJobConfigurationDetailsJobTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJobConfigurationDetailsJobTypeEnum(val string) (JobConfigurationDetailsJobTypeEnum, bool) {
	mappingJobConfigurationDetailsJobTypeEnumIgnoreCase := make(map[string]JobConfigurationDetailsJobTypeEnum)
	for k, v := range mappingJobConfigurationDetailsJobTypeEnum {
		mappingJobConfigurationDetailsJobTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingJobConfigurationDetailsJobTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
