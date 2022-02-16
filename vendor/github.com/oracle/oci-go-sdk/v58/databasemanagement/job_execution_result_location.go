// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// JobExecutionResultLocation The location of the job execution result.
type JobExecutionResultLocation interface {
}

type jobexecutionresultlocation struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *jobexecutionresultlocation) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerjobexecutionresultlocation jobexecutionresultlocation
	s := struct {
		Model Unmarshalerjobexecutionresultlocation
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *jobexecutionresultlocation) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "OBJECT_STORAGE":
		mm := ObjectStorageJobExecutionResultLocation{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m jobexecutionresultlocation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m jobexecutionresultlocation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// JobExecutionResultLocationTypeEnum Enum with underlying type: string
type JobExecutionResultLocationTypeEnum string

// Set of constants representing the allowable values for JobExecutionResultLocationTypeEnum
const (
	JobExecutionResultLocationTypeObjectStorage JobExecutionResultLocationTypeEnum = "OBJECT_STORAGE"
)

var mappingJobExecutionResultLocationTypeEnum = map[string]JobExecutionResultLocationTypeEnum{
	"OBJECT_STORAGE": JobExecutionResultLocationTypeObjectStorage,
}

// GetJobExecutionResultLocationTypeEnumValues Enumerates the set of values for JobExecutionResultLocationTypeEnum
func GetJobExecutionResultLocationTypeEnumValues() []JobExecutionResultLocationTypeEnum {
	values := make([]JobExecutionResultLocationTypeEnum, 0)
	for _, v := range mappingJobExecutionResultLocationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetJobExecutionResultLocationTypeEnumStringValues Enumerates the set of values in String for JobExecutionResultLocationTypeEnum
func GetJobExecutionResultLocationTypeEnumStringValues() []string {
	return []string{
		"OBJECT_STORAGE",
	}
}

// GetMappingJobExecutionResultLocationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJobExecutionResultLocationTypeEnum(val string) (JobExecutionResultLocationTypeEnum, bool) {
	mappingJobExecutionResultLocationTypeEnumIgnoreCase := make(map[string]JobExecutionResultLocationTypeEnum)
	for k, v := range mappingJobExecutionResultLocationTypeEnum {
		mappingJobExecutionResultLocationTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingJobExecutionResultLocationTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
