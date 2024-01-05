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

// JobInBind The details of the job in-bind variable.
type JobInBind struct {

	// The position of the in-bind variable.
	Position *int `mandatory:"true" json:"position"`

	// The datatype of the in-bind variable.
	DataType JobInBindDataTypeEnum `mandatory:"true" json:"dataType"`

	// The values for the in-bind variable.
	Values []string `mandatory:"true" json:"values"`

	// The Oracle schema object name for the predefined type of array.
	ArrayTypeName *string `mandatory:"false" json:"arrayTypeName"`
}

func (m JobInBind) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m JobInBind) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingJobInBindDataTypeEnum(string(m.DataType)); !ok && m.DataType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DataType: %s. Supported values are: %s.", m.DataType, strings.Join(GetJobInBindDataTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// JobInBindDataTypeEnum Enum with underlying type: string
type JobInBindDataTypeEnum string

// Set of constants representing the allowable values for JobInBindDataTypeEnum
const (
	JobInBindDataTypeNumber JobInBindDataTypeEnum = "NUMBER"
	JobInBindDataTypeString JobInBindDataTypeEnum = "STRING"
	JobInBindDataTypeClob   JobInBindDataTypeEnum = "CLOB"
)

var mappingJobInBindDataTypeEnum = map[string]JobInBindDataTypeEnum{
	"NUMBER": JobInBindDataTypeNumber,
	"STRING": JobInBindDataTypeString,
	"CLOB":   JobInBindDataTypeClob,
}

var mappingJobInBindDataTypeEnumLowerCase = map[string]JobInBindDataTypeEnum{
	"number": JobInBindDataTypeNumber,
	"string": JobInBindDataTypeString,
	"clob":   JobInBindDataTypeClob,
}

// GetJobInBindDataTypeEnumValues Enumerates the set of values for JobInBindDataTypeEnum
func GetJobInBindDataTypeEnumValues() []JobInBindDataTypeEnum {
	values := make([]JobInBindDataTypeEnum, 0)
	for _, v := range mappingJobInBindDataTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetJobInBindDataTypeEnumStringValues Enumerates the set of values in String for JobInBindDataTypeEnum
func GetJobInBindDataTypeEnumStringValues() []string {
	return []string{
		"NUMBER",
		"STRING",
		"CLOB",
	}
}

// GetMappingJobInBindDataTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJobInBindDataTypeEnum(val string) (JobInBindDataTypeEnum, bool) {
	enum, ok := mappingJobInBindDataTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
