// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// JobOutBind The details of the job out-bind variable.
type JobOutBind struct {

	// The position of the out-bind variable.
	Position *int `mandatory:"true" json:"position"`

	// The datatype of the out-bind variable.
	DataType JobOutBindDataTypeEnum `mandatory:"true" json:"dataType"`
}

func (m JobOutBind) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m JobOutBind) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingJobOutBindDataTypeEnum(string(m.DataType)); !ok && m.DataType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DataType: %s. Supported values are: %s.", m.DataType, strings.Join(GetJobOutBindDataTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// JobOutBindDataTypeEnum Enum with underlying type: string
type JobOutBindDataTypeEnum string

// Set of constants representing the allowable values for JobOutBindDataTypeEnum
const (
	JobOutBindDataTypeNumber JobOutBindDataTypeEnum = "NUMBER"
	JobOutBindDataTypeString JobOutBindDataTypeEnum = "STRING"
	JobOutBindDataTypeClob   JobOutBindDataTypeEnum = "CLOB"
)

var mappingJobOutBindDataTypeEnum = map[string]JobOutBindDataTypeEnum{
	"NUMBER": JobOutBindDataTypeNumber,
	"STRING": JobOutBindDataTypeString,
	"CLOB":   JobOutBindDataTypeClob,
}

var mappingJobOutBindDataTypeEnumLowerCase = map[string]JobOutBindDataTypeEnum{
	"number": JobOutBindDataTypeNumber,
	"string": JobOutBindDataTypeString,
	"clob":   JobOutBindDataTypeClob,
}

// GetJobOutBindDataTypeEnumValues Enumerates the set of values for JobOutBindDataTypeEnum
func GetJobOutBindDataTypeEnumValues() []JobOutBindDataTypeEnum {
	values := make([]JobOutBindDataTypeEnum, 0)
	for _, v := range mappingJobOutBindDataTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetJobOutBindDataTypeEnumStringValues Enumerates the set of values in String for JobOutBindDataTypeEnum
func GetJobOutBindDataTypeEnumStringValues() []string {
	return []string{
		"NUMBER",
		"STRING",
		"CLOB",
	}
}

// GetMappingJobOutBindDataTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJobOutBindDataTypeEnum(val string) (JobOutBindDataTypeEnum, bool) {
	enum, ok := mappingJobOutBindDataTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
