// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Flow API
//
// Use the Data Flow APIs to run any Apache Spark application at any scale without deploying or managing any infrastructure.
//

package dataflow

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// RunLogSummary A summary of a log associated with a particular run.
type RunLogSummary struct {

	// The name of the log.
	// Example: spark_driver_stderr_20190917T114000Z.log.gz
	Name *string `mandatory:"true" json:"name"`

	// The runId associated with the log.
	RunId *string `mandatory:"true" json:"runId"`

	// The source of the log such as driver and executor.
	Source RunLogSummarySourceEnum `mandatory:"true" json:"source"`

	// The type of log such as stdout and stderr.
	Type RunLogSummaryTypeEnum `mandatory:"true" json:"type"`

	// The size of the object in bytes.
	SizeInBytes *int64 `mandatory:"false" json:"sizeInBytes"`

	// The date and time the object was created, as described in RFC 2616 (https://tools.ietf.org/rfc/rfc2616), section 14.29.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m RunLogSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RunLogSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRunLogSummarySourceEnum(string(m.Source)); !ok && m.Source != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Source: %s. Supported values are: %s.", m.Source, strings.Join(GetRunLogSummarySourceEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRunLogSummaryTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetRunLogSummaryTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RunLogSummarySourceEnum Enum with underlying type: string
type RunLogSummarySourceEnum string

// Set of constants representing the allowable values for RunLogSummarySourceEnum
const (
	RunLogSummarySourceApplication RunLogSummarySourceEnum = "APPLICATION"
	RunLogSummarySourceDriver      RunLogSummarySourceEnum = "DRIVER"
	RunLogSummarySourceExecutor    RunLogSummarySourceEnum = "EXECUTOR"
)

var mappingRunLogSummarySourceEnum = map[string]RunLogSummarySourceEnum{
	"APPLICATION": RunLogSummarySourceApplication,
	"DRIVER":      RunLogSummarySourceDriver,
	"EXECUTOR":    RunLogSummarySourceExecutor,
}

// GetRunLogSummarySourceEnumValues Enumerates the set of values for RunLogSummarySourceEnum
func GetRunLogSummarySourceEnumValues() []RunLogSummarySourceEnum {
	values := make([]RunLogSummarySourceEnum, 0)
	for _, v := range mappingRunLogSummarySourceEnum {
		values = append(values, v)
	}
	return values
}

// GetRunLogSummarySourceEnumStringValues Enumerates the set of values in String for RunLogSummarySourceEnum
func GetRunLogSummarySourceEnumStringValues() []string {
	return []string{
		"APPLICATION",
		"DRIVER",
		"EXECUTOR",
	}
}

// GetMappingRunLogSummarySourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRunLogSummarySourceEnum(val string) (RunLogSummarySourceEnum, bool) {
	mappingRunLogSummarySourceEnumIgnoreCase := make(map[string]RunLogSummarySourceEnum)
	for k, v := range mappingRunLogSummarySourceEnum {
		mappingRunLogSummarySourceEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingRunLogSummarySourceEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// RunLogSummaryTypeEnum Enum with underlying type: string
type RunLogSummaryTypeEnum string

// Set of constants representing the allowable values for RunLogSummaryTypeEnum
const (
	RunLogSummaryTypeStderr RunLogSummaryTypeEnum = "STDERR"
	RunLogSummaryTypeStdout RunLogSummaryTypeEnum = "STDOUT"
)

var mappingRunLogSummaryTypeEnum = map[string]RunLogSummaryTypeEnum{
	"STDERR": RunLogSummaryTypeStderr,
	"STDOUT": RunLogSummaryTypeStdout,
}

// GetRunLogSummaryTypeEnumValues Enumerates the set of values for RunLogSummaryTypeEnum
func GetRunLogSummaryTypeEnumValues() []RunLogSummaryTypeEnum {
	values := make([]RunLogSummaryTypeEnum, 0)
	for _, v := range mappingRunLogSummaryTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRunLogSummaryTypeEnumStringValues Enumerates the set of values in String for RunLogSummaryTypeEnum
func GetRunLogSummaryTypeEnumStringValues() []string {
	return []string{
		"STDERR",
		"STDOUT",
	}
}

// GetMappingRunLogSummaryTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRunLogSummaryTypeEnum(val string) (RunLogSummaryTypeEnum, bool) {
	mappingRunLogSummaryTypeEnumIgnoreCase := make(map[string]RunLogSummaryTypeEnum)
	for k, v := range mappingRunLogSummaryTypeEnum {
		mappingRunLogSummaryTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingRunLogSummaryTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
