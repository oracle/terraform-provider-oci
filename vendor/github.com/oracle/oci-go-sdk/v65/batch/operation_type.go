// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Batch API
//
// Use the Batch Control Plane API to encapsulate and manage all aspects of computationally intensive jobs.
//

package batch

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateBatchContext       OperationTypeEnum = "CREATE_BATCH_CONTEXT"
	OperationTypeUpdateBatchContext       OperationTypeEnum = "UPDATE_BATCH_CONTEXT"
	OperationTypeDeleteBatchContext       OperationTypeEnum = "DELETE_BATCH_CONTEXT"
	OperationTypeStartBatchContext        OperationTypeEnum = "START_BATCH_CONTEXT"
	OperationTypeStopBatchContext         OperationTypeEnum = "STOP_BATCH_CONTEXT"
	OperationTypeMoveBatchContext         OperationTypeEnum = "MOVE_BATCH_CONTEXT"
	OperationTypeCreateBatchJob           OperationTypeEnum = "CREATE_BATCH_JOB"
	OperationTypeUpdateBatchJob           OperationTypeEnum = "UPDATE_BATCH_JOB"
	OperationTypeDeleteBatchJob           OperationTypeEnum = "DELETE_BATCH_JOB"
	OperationTypeMoveBatchJob             OperationTypeEnum = "MOVE_BATCH_JOB"
	OperationTypePauseBatchJob            OperationTypeEnum = "PAUSE_BATCH_JOB"
	OperationTypeUnpauseBatchJob          OperationTypeEnum = "UNPAUSE_BATCH_JOB"
	OperationTypeMoveBatchJobPool         OperationTypeEnum = "MOVE_BATCH_JOB_POOL"
	OperationTypeUpdateBatchJobPool       OperationTypeEnum = "UPDATE_BATCH_JOB_POOL"
	OperationTypeStartBatchJobPool        OperationTypeEnum = "START_BATCH_JOB_POOL"
	OperationTypeStopBatchJobPool         OperationTypeEnum = "STOP_BATCH_JOB_POOL"
	OperationTypeMoveBatchTaskEnvironment OperationTypeEnum = "MOVE_BATCH_TASK_ENVIRONMENT"
	OperationTypeMoveBatchTaskProfile     OperationTypeEnum = "MOVE_BATCH_TASK_PROFILE"
	OperationTypeInternal                 OperationTypeEnum = "INTERNAL"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_BATCH_CONTEXT":        OperationTypeCreateBatchContext,
	"UPDATE_BATCH_CONTEXT":        OperationTypeUpdateBatchContext,
	"DELETE_BATCH_CONTEXT":        OperationTypeDeleteBatchContext,
	"START_BATCH_CONTEXT":         OperationTypeStartBatchContext,
	"STOP_BATCH_CONTEXT":          OperationTypeStopBatchContext,
	"MOVE_BATCH_CONTEXT":          OperationTypeMoveBatchContext,
	"CREATE_BATCH_JOB":            OperationTypeCreateBatchJob,
	"UPDATE_BATCH_JOB":            OperationTypeUpdateBatchJob,
	"DELETE_BATCH_JOB":            OperationTypeDeleteBatchJob,
	"MOVE_BATCH_JOB":              OperationTypeMoveBatchJob,
	"PAUSE_BATCH_JOB":             OperationTypePauseBatchJob,
	"UNPAUSE_BATCH_JOB":           OperationTypeUnpauseBatchJob,
	"MOVE_BATCH_JOB_POOL":         OperationTypeMoveBatchJobPool,
	"UPDATE_BATCH_JOB_POOL":       OperationTypeUpdateBatchJobPool,
	"START_BATCH_JOB_POOL":        OperationTypeStartBatchJobPool,
	"STOP_BATCH_JOB_POOL":         OperationTypeStopBatchJobPool,
	"MOVE_BATCH_TASK_ENVIRONMENT": OperationTypeMoveBatchTaskEnvironment,
	"MOVE_BATCH_TASK_PROFILE":     OperationTypeMoveBatchTaskProfile,
	"INTERNAL":                    OperationTypeInternal,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_batch_context":        OperationTypeCreateBatchContext,
	"update_batch_context":        OperationTypeUpdateBatchContext,
	"delete_batch_context":        OperationTypeDeleteBatchContext,
	"start_batch_context":         OperationTypeStartBatchContext,
	"stop_batch_context":          OperationTypeStopBatchContext,
	"move_batch_context":          OperationTypeMoveBatchContext,
	"create_batch_job":            OperationTypeCreateBatchJob,
	"update_batch_job":            OperationTypeUpdateBatchJob,
	"delete_batch_job":            OperationTypeDeleteBatchJob,
	"move_batch_job":              OperationTypeMoveBatchJob,
	"pause_batch_job":             OperationTypePauseBatchJob,
	"unpause_batch_job":           OperationTypeUnpauseBatchJob,
	"move_batch_job_pool":         OperationTypeMoveBatchJobPool,
	"update_batch_job_pool":       OperationTypeUpdateBatchJobPool,
	"start_batch_job_pool":        OperationTypeStartBatchJobPool,
	"stop_batch_job_pool":         OperationTypeStopBatchJobPool,
	"move_batch_task_environment": OperationTypeMoveBatchTaskEnvironment,
	"move_batch_task_profile":     OperationTypeMoveBatchTaskProfile,
	"internal":                    OperationTypeInternal,
}

// GetOperationTypeEnumValues Enumerates the set of values for OperationTypeEnum
func GetOperationTypeEnumValues() []OperationTypeEnum {
	values := make([]OperationTypeEnum, 0)
	for _, v := range mappingOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOperationTypeEnumStringValues Enumerates the set of values in String for OperationTypeEnum
func GetOperationTypeEnumStringValues() []string {
	return []string{
		"CREATE_BATCH_CONTEXT",
		"UPDATE_BATCH_CONTEXT",
		"DELETE_BATCH_CONTEXT",
		"START_BATCH_CONTEXT",
		"STOP_BATCH_CONTEXT",
		"MOVE_BATCH_CONTEXT",
		"CREATE_BATCH_JOB",
		"UPDATE_BATCH_JOB",
		"DELETE_BATCH_JOB",
		"MOVE_BATCH_JOB",
		"PAUSE_BATCH_JOB",
		"UNPAUSE_BATCH_JOB",
		"MOVE_BATCH_JOB_POOL",
		"UPDATE_BATCH_JOB_POOL",
		"START_BATCH_JOB_POOL",
		"STOP_BATCH_JOB_POOL",
		"MOVE_BATCH_TASK_ENVIRONMENT",
		"MOVE_BATCH_TASK_PROFILE",
		"INTERNAL",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
