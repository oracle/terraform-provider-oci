// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// JobExecutionResultDetails The job execution result details.
type JobExecutionResultDetails interface {
}

type jobexecutionresultdetails struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *jobexecutionresultdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerjobexecutionresultdetails jobexecutionresultdetails
	s := struct {
		Model Unmarshalerjobexecutionresultdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *jobexecutionresultdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "OBJECT_STORAGE":
		mm := ObjectStorageJobExecutionResultDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for JobExecutionResultDetails: %s.", m.Type)
		return *m, nil
	}
}

func (m jobexecutionresultdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m jobexecutionresultdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// JobExecutionResultDetailsTypeEnum Enum with underlying type: string
type JobExecutionResultDetailsTypeEnum string

// Set of constants representing the allowable values for JobExecutionResultDetailsTypeEnum
const (
	JobExecutionResultDetailsTypeObjectStorage JobExecutionResultDetailsTypeEnum = "OBJECT_STORAGE"
)

var mappingJobExecutionResultDetailsTypeEnum = map[string]JobExecutionResultDetailsTypeEnum{
	"OBJECT_STORAGE": JobExecutionResultDetailsTypeObjectStorage,
}

var mappingJobExecutionResultDetailsTypeEnumLowerCase = map[string]JobExecutionResultDetailsTypeEnum{
	"object_storage": JobExecutionResultDetailsTypeObjectStorage,
}

// GetJobExecutionResultDetailsTypeEnumValues Enumerates the set of values for JobExecutionResultDetailsTypeEnum
func GetJobExecutionResultDetailsTypeEnumValues() []JobExecutionResultDetailsTypeEnum {
	values := make([]JobExecutionResultDetailsTypeEnum, 0)
	for _, v := range mappingJobExecutionResultDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetJobExecutionResultDetailsTypeEnumStringValues Enumerates the set of values in String for JobExecutionResultDetailsTypeEnum
func GetJobExecutionResultDetailsTypeEnumStringValues() []string {
	return []string{
		"OBJECT_STORAGE",
	}
}

// GetMappingJobExecutionResultDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJobExecutionResultDetailsTypeEnum(val string) (JobExecutionResultDetailsTypeEnum, bool) {
	enum, ok := mappingJobExecutionResultDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
