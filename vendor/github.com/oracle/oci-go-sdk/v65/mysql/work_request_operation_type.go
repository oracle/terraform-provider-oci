// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"strings"
)

// WorkRequestOperationTypeEnum Enum with underlying type: string
type WorkRequestOperationTypeEnum string

// Set of constants representing the allowable values for WorkRequestOperationTypeEnum
const (
	WorkRequestOperationTypeCreateDbsystem                        WorkRequestOperationTypeEnum = "CREATE_DBSYSTEM"
	WorkRequestOperationTypeUpdateDbsystem                        WorkRequestOperationTypeEnum = "UPDATE_DBSYSTEM"
	WorkRequestOperationTypeDeleteDbsystem                        WorkRequestOperationTypeEnum = "DELETE_DBSYSTEM"
	WorkRequestOperationTypeStartDbsystem                         WorkRequestOperationTypeEnum = "START_DBSYSTEM"
	WorkRequestOperationTypeStopDbsystem                          WorkRequestOperationTypeEnum = "STOP_DBSYSTEM"
	WorkRequestOperationTypeRestartDbsystem                       WorkRequestOperationTypeEnum = "RESTART_DBSYSTEM"
	WorkRequestOperationTypeAddHeatwaveCluster                    WorkRequestOperationTypeEnum = "ADD_HEATWAVE_CLUSTER"
	WorkRequestOperationTypeUpdateHeatwaveCluster                 WorkRequestOperationTypeEnum = "UPDATE_HEATWAVE_CLUSTER"
	WorkRequestOperationTypeDeleteHeatwaveCluster                 WorkRequestOperationTypeEnum = "DELETE_HEATWAVE_CLUSTER"
	WorkRequestOperationTypeStartHeatwaveCluster                  WorkRequestOperationTypeEnum = "START_HEATWAVE_CLUSTER"
	WorkRequestOperationTypeStopHeatwaveCluster                   WorkRequestOperationTypeEnum = "STOP_HEATWAVE_CLUSTER"
	WorkRequestOperationTypeRestartHeatwaveCluster                WorkRequestOperationTypeEnum = "RESTART_HEATWAVE_CLUSTER"
	WorkRequestOperationTypeGenerateHeatwaveClusterMemoryEstimate WorkRequestOperationTypeEnum = "GENERATE_HEATWAVE_CLUSTER_MEMORY_ESTIMATE"
	WorkRequestOperationTypeCreateReplica                         WorkRequestOperationTypeEnum = "CREATE_REPLICA"
	WorkRequestOperationTypeUpdateReplica                         WorkRequestOperationTypeEnum = "UPDATE_REPLICA"
	WorkRequestOperationTypeDeleteReplica                         WorkRequestOperationTypeEnum = "DELETE_REPLICA"
	WorkRequestOperationTypeCreateChannel                         WorkRequestOperationTypeEnum = "CREATE_CHANNEL"
	WorkRequestOperationTypeUpdateChannel                         WorkRequestOperationTypeEnum = "UPDATE_CHANNEL"
	WorkRequestOperationTypeResumeChannel                         WorkRequestOperationTypeEnum = "RESUME_CHANNEL"
	WorkRequestOperationTypeResetChannel                          WorkRequestOperationTypeEnum = "RESET_CHANNEL"
	WorkRequestOperationTypeDeleteChannel                         WorkRequestOperationTypeEnum = "DELETE_CHANNEL"
	WorkRequestOperationTypeCreateBackup                          WorkRequestOperationTypeEnum = "CREATE_BACKUP"
	WorkRequestOperationTypeDeleteBackup                          WorkRequestOperationTypeEnum = "DELETE_BACKUP"
)

var mappingWorkRequestOperationTypeEnum = map[string]WorkRequestOperationTypeEnum{
	"CREATE_DBSYSTEM":                           WorkRequestOperationTypeCreateDbsystem,
	"UPDATE_DBSYSTEM":                           WorkRequestOperationTypeUpdateDbsystem,
	"DELETE_DBSYSTEM":                           WorkRequestOperationTypeDeleteDbsystem,
	"START_DBSYSTEM":                            WorkRequestOperationTypeStartDbsystem,
	"STOP_DBSYSTEM":                             WorkRequestOperationTypeStopDbsystem,
	"RESTART_DBSYSTEM":                          WorkRequestOperationTypeRestartDbsystem,
	"ADD_HEATWAVE_CLUSTER":                      WorkRequestOperationTypeAddHeatwaveCluster,
	"UPDATE_HEATWAVE_CLUSTER":                   WorkRequestOperationTypeUpdateHeatwaveCluster,
	"DELETE_HEATWAVE_CLUSTER":                   WorkRequestOperationTypeDeleteHeatwaveCluster,
	"START_HEATWAVE_CLUSTER":                    WorkRequestOperationTypeStartHeatwaveCluster,
	"STOP_HEATWAVE_CLUSTER":                     WorkRequestOperationTypeStopHeatwaveCluster,
	"RESTART_HEATWAVE_CLUSTER":                  WorkRequestOperationTypeRestartHeatwaveCluster,
	"GENERATE_HEATWAVE_CLUSTER_MEMORY_ESTIMATE": WorkRequestOperationTypeGenerateHeatwaveClusterMemoryEstimate,
	"CREATE_REPLICA":                            WorkRequestOperationTypeCreateReplica,
	"UPDATE_REPLICA":                            WorkRequestOperationTypeUpdateReplica,
	"DELETE_REPLICA":                            WorkRequestOperationTypeDeleteReplica,
	"CREATE_CHANNEL":                            WorkRequestOperationTypeCreateChannel,
	"UPDATE_CHANNEL":                            WorkRequestOperationTypeUpdateChannel,
	"RESUME_CHANNEL":                            WorkRequestOperationTypeResumeChannel,
	"RESET_CHANNEL":                             WorkRequestOperationTypeResetChannel,
	"DELETE_CHANNEL":                            WorkRequestOperationTypeDeleteChannel,
	"CREATE_BACKUP":                             WorkRequestOperationTypeCreateBackup,
	"DELETE_BACKUP":                             WorkRequestOperationTypeDeleteBackup,
}

var mappingWorkRequestOperationTypeEnumLowerCase = map[string]WorkRequestOperationTypeEnum{
	"create_dbsystem":                           WorkRequestOperationTypeCreateDbsystem,
	"update_dbsystem":                           WorkRequestOperationTypeUpdateDbsystem,
	"delete_dbsystem":                           WorkRequestOperationTypeDeleteDbsystem,
	"start_dbsystem":                            WorkRequestOperationTypeStartDbsystem,
	"stop_dbsystem":                             WorkRequestOperationTypeStopDbsystem,
	"restart_dbsystem":                          WorkRequestOperationTypeRestartDbsystem,
	"add_heatwave_cluster":                      WorkRequestOperationTypeAddHeatwaveCluster,
	"update_heatwave_cluster":                   WorkRequestOperationTypeUpdateHeatwaveCluster,
	"delete_heatwave_cluster":                   WorkRequestOperationTypeDeleteHeatwaveCluster,
	"start_heatwave_cluster":                    WorkRequestOperationTypeStartHeatwaveCluster,
	"stop_heatwave_cluster":                     WorkRequestOperationTypeStopHeatwaveCluster,
	"restart_heatwave_cluster":                  WorkRequestOperationTypeRestartHeatwaveCluster,
	"generate_heatwave_cluster_memory_estimate": WorkRequestOperationTypeGenerateHeatwaveClusterMemoryEstimate,
	"create_replica":                            WorkRequestOperationTypeCreateReplica,
	"update_replica":                            WorkRequestOperationTypeUpdateReplica,
	"delete_replica":                            WorkRequestOperationTypeDeleteReplica,
	"create_channel":                            WorkRequestOperationTypeCreateChannel,
	"update_channel":                            WorkRequestOperationTypeUpdateChannel,
	"resume_channel":                            WorkRequestOperationTypeResumeChannel,
	"reset_channel":                             WorkRequestOperationTypeResetChannel,
	"delete_channel":                            WorkRequestOperationTypeDeleteChannel,
	"create_backup":                             WorkRequestOperationTypeCreateBackup,
	"delete_backup":                             WorkRequestOperationTypeDeleteBackup,
}

// GetWorkRequestOperationTypeEnumValues Enumerates the set of values for WorkRequestOperationTypeEnum
func GetWorkRequestOperationTypeEnumValues() []WorkRequestOperationTypeEnum {
	values := make([]WorkRequestOperationTypeEnum, 0)
	for _, v := range mappingWorkRequestOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetWorkRequestOperationTypeEnumStringValues Enumerates the set of values in String for WorkRequestOperationTypeEnum
func GetWorkRequestOperationTypeEnumStringValues() []string {
	return []string{
		"CREATE_DBSYSTEM",
		"UPDATE_DBSYSTEM",
		"DELETE_DBSYSTEM",
		"START_DBSYSTEM",
		"STOP_DBSYSTEM",
		"RESTART_DBSYSTEM",
		"ADD_HEATWAVE_CLUSTER",
		"UPDATE_HEATWAVE_CLUSTER",
		"DELETE_HEATWAVE_CLUSTER",
		"START_HEATWAVE_CLUSTER",
		"STOP_HEATWAVE_CLUSTER",
		"RESTART_HEATWAVE_CLUSTER",
		"GENERATE_HEATWAVE_CLUSTER_MEMORY_ESTIMATE",
		"CREATE_REPLICA",
		"UPDATE_REPLICA",
		"DELETE_REPLICA",
		"CREATE_CHANNEL",
		"UPDATE_CHANNEL",
		"RESUME_CHANNEL",
		"RESET_CHANNEL",
		"DELETE_CHANNEL",
		"CREATE_BACKUP",
		"DELETE_BACKUP",
	}
}

// GetMappingWorkRequestOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkRequestOperationTypeEnum(val string) (WorkRequestOperationTypeEnum, bool) {
	enum, ok := mappingWorkRequestOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
