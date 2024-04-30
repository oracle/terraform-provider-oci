// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.cloud.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OptimizerStatisticsOperationTask The details of the Optimizer Statistics Collection task.
type OptimizerStatisticsOperationTask struct {

	// The name of the target object for which statistics are gathered.
	Target *string `mandatory:"true" json:"target"`

	// The type of target object.
	TargetType OptimizerStatisticsOperationTaskTargetTypeEnum `mandatory:"true" json:"targetType"`

	// The start time of the Optimizer Statistics Collection task.
	TimeStart *common.SDKTime `mandatory:"true" json:"timeStart"`

	// The end time of the Optimizer Statistics Collection task.
	TimeEnd *common.SDKTime `mandatory:"true" json:"timeEnd"`

	// The status of the Optimizer Statistics Collection task.
	Status OptimizerStatisticsOperationTaskStatusEnum `mandatory:"true" json:"status"`
}

func (m OptimizerStatisticsOperationTask) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OptimizerStatisticsOperationTask) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOptimizerStatisticsOperationTaskTargetTypeEnum(string(m.TargetType)); !ok && m.TargetType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TargetType: %s. Supported values are: %s.", m.TargetType, strings.Join(GetOptimizerStatisticsOperationTaskTargetTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOptimizerStatisticsOperationTaskStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetOptimizerStatisticsOperationTaskStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OptimizerStatisticsOperationTaskTargetTypeEnum Enum with underlying type: string
type OptimizerStatisticsOperationTaskTargetTypeEnum string

// Set of constants representing the allowable values for OptimizerStatisticsOperationTaskTargetTypeEnum
const (
	OptimizerStatisticsOperationTaskTargetTypeTable             OptimizerStatisticsOperationTaskTargetTypeEnum = "TABLE"
	OptimizerStatisticsOperationTaskTargetTypeGlobalTable       OptimizerStatisticsOperationTaskTargetTypeEnum = "GLOBAL_TABLE"
	OptimizerStatisticsOperationTaskTargetTypeCoordinatorTable  OptimizerStatisticsOperationTaskTargetTypeEnum = "COORDINATOR_TABLE"
	OptimizerStatisticsOperationTaskTargetTypeTablePartition    OptimizerStatisticsOperationTaskTargetTypeEnum = "TABLE_PARTITION"
	OptimizerStatisticsOperationTaskTargetTypeTableSubpartition OptimizerStatisticsOperationTaskTargetTypeEnum = "TABLE_SUBPARTITION"
	OptimizerStatisticsOperationTaskTargetTypeIndex             OptimizerStatisticsOperationTaskTargetTypeEnum = "INDEX"
	OptimizerStatisticsOperationTaskTargetTypeIndexPartition    OptimizerStatisticsOperationTaskTargetTypeEnum = "INDEX_PARTITION"
	OptimizerStatisticsOperationTaskTargetTypeIndexSubpartition OptimizerStatisticsOperationTaskTargetTypeEnum = "INDEX_SUBPARTITION"
)

var mappingOptimizerStatisticsOperationTaskTargetTypeEnum = map[string]OptimizerStatisticsOperationTaskTargetTypeEnum{
	"TABLE":              OptimizerStatisticsOperationTaskTargetTypeTable,
	"GLOBAL_TABLE":       OptimizerStatisticsOperationTaskTargetTypeGlobalTable,
	"COORDINATOR_TABLE":  OptimizerStatisticsOperationTaskTargetTypeCoordinatorTable,
	"TABLE_PARTITION":    OptimizerStatisticsOperationTaskTargetTypeTablePartition,
	"TABLE_SUBPARTITION": OptimizerStatisticsOperationTaskTargetTypeTableSubpartition,
	"INDEX":              OptimizerStatisticsOperationTaskTargetTypeIndex,
	"INDEX_PARTITION":    OptimizerStatisticsOperationTaskTargetTypeIndexPartition,
	"INDEX_SUBPARTITION": OptimizerStatisticsOperationTaskTargetTypeIndexSubpartition,
}

var mappingOptimizerStatisticsOperationTaskTargetTypeEnumLowerCase = map[string]OptimizerStatisticsOperationTaskTargetTypeEnum{
	"table":              OptimizerStatisticsOperationTaskTargetTypeTable,
	"global_table":       OptimizerStatisticsOperationTaskTargetTypeGlobalTable,
	"coordinator_table":  OptimizerStatisticsOperationTaskTargetTypeCoordinatorTable,
	"table_partition":    OptimizerStatisticsOperationTaskTargetTypeTablePartition,
	"table_subpartition": OptimizerStatisticsOperationTaskTargetTypeTableSubpartition,
	"index":              OptimizerStatisticsOperationTaskTargetTypeIndex,
	"index_partition":    OptimizerStatisticsOperationTaskTargetTypeIndexPartition,
	"index_subpartition": OptimizerStatisticsOperationTaskTargetTypeIndexSubpartition,
}

// GetOptimizerStatisticsOperationTaskTargetTypeEnumValues Enumerates the set of values for OptimizerStatisticsOperationTaskTargetTypeEnum
func GetOptimizerStatisticsOperationTaskTargetTypeEnumValues() []OptimizerStatisticsOperationTaskTargetTypeEnum {
	values := make([]OptimizerStatisticsOperationTaskTargetTypeEnum, 0)
	for _, v := range mappingOptimizerStatisticsOperationTaskTargetTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOptimizerStatisticsOperationTaskTargetTypeEnumStringValues Enumerates the set of values in String for OptimizerStatisticsOperationTaskTargetTypeEnum
func GetOptimizerStatisticsOperationTaskTargetTypeEnumStringValues() []string {
	return []string{
		"TABLE",
		"GLOBAL_TABLE",
		"COORDINATOR_TABLE",
		"TABLE_PARTITION",
		"TABLE_SUBPARTITION",
		"INDEX",
		"INDEX_PARTITION",
		"INDEX_SUBPARTITION",
	}
}

// GetMappingOptimizerStatisticsOperationTaskTargetTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOptimizerStatisticsOperationTaskTargetTypeEnum(val string) (OptimizerStatisticsOperationTaskTargetTypeEnum, bool) {
	enum, ok := mappingOptimizerStatisticsOperationTaskTargetTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// OptimizerStatisticsOperationTaskStatusEnum Enum with underlying type: string
type OptimizerStatisticsOperationTaskStatusEnum string

// Set of constants representing the allowable values for OptimizerStatisticsOperationTaskStatusEnum
const (
	OptimizerStatisticsOperationTaskStatusPending    OptimizerStatisticsOperationTaskStatusEnum = "PENDING"
	OptimizerStatisticsOperationTaskStatusInProgress OptimizerStatisticsOperationTaskStatusEnum = "IN_PROGRESS"
	OptimizerStatisticsOperationTaskStatusSkipped    OptimizerStatisticsOperationTaskStatusEnum = "SKIPPED"
	OptimizerStatisticsOperationTaskStatusTimedOut   OptimizerStatisticsOperationTaskStatusEnum = "TIMED_OUT"
	OptimizerStatisticsOperationTaskStatusCompleted  OptimizerStatisticsOperationTaskStatusEnum = "COMPLETED"
	OptimizerStatisticsOperationTaskStatusFailed     OptimizerStatisticsOperationTaskStatusEnum = "FAILED"
)

var mappingOptimizerStatisticsOperationTaskStatusEnum = map[string]OptimizerStatisticsOperationTaskStatusEnum{
	"PENDING":     OptimizerStatisticsOperationTaskStatusPending,
	"IN_PROGRESS": OptimizerStatisticsOperationTaskStatusInProgress,
	"SKIPPED":     OptimizerStatisticsOperationTaskStatusSkipped,
	"TIMED_OUT":   OptimizerStatisticsOperationTaskStatusTimedOut,
	"COMPLETED":   OptimizerStatisticsOperationTaskStatusCompleted,
	"FAILED":      OptimizerStatisticsOperationTaskStatusFailed,
}

var mappingOptimizerStatisticsOperationTaskStatusEnumLowerCase = map[string]OptimizerStatisticsOperationTaskStatusEnum{
	"pending":     OptimizerStatisticsOperationTaskStatusPending,
	"in_progress": OptimizerStatisticsOperationTaskStatusInProgress,
	"skipped":     OptimizerStatisticsOperationTaskStatusSkipped,
	"timed_out":   OptimizerStatisticsOperationTaskStatusTimedOut,
	"completed":   OptimizerStatisticsOperationTaskStatusCompleted,
	"failed":      OptimizerStatisticsOperationTaskStatusFailed,
}

// GetOptimizerStatisticsOperationTaskStatusEnumValues Enumerates the set of values for OptimizerStatisticsOperationTaskStatusEnum
func GetOptimizerStatisticsOperationTaskStatusEnumValues() []OptimizerStatisticsOperationTaskStatusEnum {
	values := make([]OptimizerStatisticsOperationTaskStatusEnum, 0)
	for _, v := range mappingOptimizerStatisticsOperationTaskStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetOptimizerStatisticsOperationTaskStatusEnumStringValues Enumerates the set of values in String for OptimizerStatisticsOperationTaskStatusEnum
func GetOptimizerStatisticsOperationTaskStatusEnumStringValues() []string {
	return []string{
		"PENDING",
		"IN_PROGRESS",
		"SKIPPED",
		"TIMED_OUT",
		"COMPLETED",
		"FAILED",
	}
}

// GetMappingOptimizerStatisticsOperationTaskStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOptimizerStatisticsOperationTaskStatusEnum(val string) (OptimizerStatisticsOperationTaskStatusEnum, bool) {
	enum, ok := mappingOptimizerStatisticsOperationTaskStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
