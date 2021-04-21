// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

// WorkRequestOperationTypeEnum Enum with underlying type: string
type WorkRequestOperationTypeEnum string

// Set of constants representing the allowable values for WorkRequestOperationTypeEnum
const (
	WorkRequestOperationTypeCreateDbsystem                         WorkRequestOperationTypeEnum = "CREATE_DBSYSTEM"
	WorkRequestOperationTypeUpdateDbsystem                         WorkRequestOperationTypeEnum = "UPDATE_DBSYSTEM"
	WorkRequestOperationTypeDeleteDbsystem                         WorkRequestOperationTypeEnum = "DELETE_DBSYSTEM"
	WorkRequestOperationTypeStartDbsystem                          WorkRequestOperationTypeEnum = "START_DBSYSTEM"
	WorkRequestOperationTypeStopDbsystem                           WorkRequestOperationTypeEnum = "STOP_DBSYSTEM"
	WorkRequestOperationTypeRestartDbsystem                        WorkRequestOperationTypeEnum = "RESTART_DBSYSTEM"
	WorkRequestOperationTypeAddAnalyticsCluster                    WorkRequestOperationTypeEnum = "ADD_ANALYTICS_CLUSTER"
	WorkRequestOperationTypeUpdateAnalyticsCluster                 WorkRequestOperationTypeEnum = "UPDATE_ANALYTICS_CLUSTER"
	WorkRequestOperationTypeDeleteAnalyticsCluster                 WorkRequestOperationTypeEnum = "DELETE_ANALYTICS_CLUSTER"
	WorkRequestOperationTypeStartAnalyticsCluster                  WorkRequestOperationTypeEnum = "START_ANALYTICS_CLUSTER"
	WorkRequestOperationTypeStopAnalyticsCluster                   WorkRequestOperationTypeEnum = "STOP_ANALYTICS_CLUSTER"
	WorkRequestOperationTypeRestartAnalyticsCluster                WorkRequestOperationTypeEnum = "RESTART_ANALYTICS_CLUSTER"
	WorkRequestOperationTypeGenerateAnalyticsClusterMemoryEstimate WorkRequestOperationTypeEnum = "GENERATE_ANALYTICS_CLUSTER_MEMORY_ESTIMATE"
	WorkRequestOperationTypeAddHeatwaveCluster                     WorkRequestOperationTypeEnum = "ADD_HEATWAVE_CLUSTER"
	WorkRequestOperationTypeUpdateHeatwaveCluster                  WorkRequestOperationTypeEnum = "UPDATE_HEATWAVE_CLUSTER"
	WorkRequestOperationTypeDeleteHeatwaveCluster                  WorkRequestOperationTypeEnum = "DELETE_HEATWAVE_CLUSTER"
	WorkRequestOperationTypeStartHeatwaveCluster                   WorkRequestOperationTypeEnum = "START_HEATWAVE_CLUSTER"
	WorkRequestOperationTypeStopHeatwaveCluster                    WorkRequestOperationTypeEnum = "STOP_HEATWAVE_CLUSTER"
	WorkRequestOperationTypeRestartHeatwaveCluster                 WorkRequestOperationTypeEnum = "RESTART_HEATWAVE_CLUSTER"
	WorkRequestOperationTypeGenerateHeatwaveClusterMemoryEstimate  WorkRequestOperationTypeEnum = "GENERATE_HEATWAVE_CLUSTER_MEMORY_ESTIMATE"
)

var mappingWorkRequestOperationType = map[string]WorkRequestOperationTypeEnum{
	"CREATE_DBSYSTEM":                            WorkRequestOperationTypeCreateDbsystem,
	"UPDATE_DBSYSTEM":                            WorkRequestOperationTypeUpdateDbsystem,
	"DELETE_DBSYSTEM":                            WorkRequestOperationTypeDeleteDbsystem,
	"START_DBSYSTEM":                             WorkRequestOperationTypeStartDbsystem,
	"STOP_DBSYSTEM":                              WorkRequestOperationTypeStopDbsystem,
	"RESTART_DBSYSTEM":                           WorkRequestOperationTypeRestartDbsystem,
	"ADD_ANALYTICS_CLUSTER":                      WorkRequestOperationTypeAddAnalyticsCluster,
	"UPDATE_ANALYTICS_CLUSTER":                   WorkRequestOperationTypeUpdateAnalyticsCluster,
	"DELETE_ANALYTICS_CLUSTER":                   WorkRequestOperationTypeDeleteAnalyticsCluster,
	"START_ANALYTICS_CLUSTER":                    WorkRequestOperationTypeStartAnalyticsCluster,
	"STOP_ANALYTICS_CLUSTER":                     WorkRequestOperationTypeStopAnalyticsCluster,
	"RESTART_ANALYTICS_CLUSTER":                  WorkRequestOperationTypeRestartAnalyticsCluster,
	"GENERATE_ANALYTICS_CLUSTER_MEMORY_ESTIMATE": WorkRequestOperationTypeGenerateAnalyticsClusterMemoryEstimate,
	"ADD_HEATWAVE_CLUSTER":                       WorkRequestOperationTypeAddHeatwaveCluster,
	"UPDATE_HEATWAVE_CLUSTER":                    WorkRequestOperationTypeUpdateHeatwaveCluster,
	"DELETE_HEATWAVE_CLUSTER":                    WorkRequestOperationTypeDeleteHeatwaveCluster,
	"START_HEATWAVE_CLUSTER":                     WorkRequestOperationTypeStartHeatwaveCluster,
	"STOP_HEATWAVE_CLUSTER":                      WorkRequestOperationTypeStopHeatwaveCluster,
	"RESTART_HEATWAVE_CLUSTER":                   WorkRequestOperationTypeRestartHeatwaveCluster,
	"GENERATE_HEATWAVE_CLUSTER_MEMORY_ESTIMATE":  WorkRequestOperationTypeGenerateHeatwaveClusterMemoryEstimate,
}

// GetWorkRequestOperationTypeEnumValues Enumerates the set of values for WorkRequestOperationTypeEnum
func GetWorkRequestOperationTypeEnumValues() []WorkRequestOperationTypeEnum {
	values := make([]WorkRequestOperationTypeEnum, 0)
	for _, v := range mappingWorkRequestOperationType {
		values = append(values, v)
	}
	return values
}
