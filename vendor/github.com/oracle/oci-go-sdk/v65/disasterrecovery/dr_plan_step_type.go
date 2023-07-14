// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Full Stack Disaster Recovery API
//
// Use the Full Stack Disaster Recovery (FSDR) API to manage disaster recovery for business applications.
// FSDR is an OCI disaster recovery orchestration and management service that provides comprehensive disaster recovery
// capabilities for all layers of an application stack, including infrastructure, middleware, database, and application.
//

package disasterrecovery

import (
	"strings"
)

// DrPlanStepTypeEnum Enum with underlying type: string
type DrPlanStepTypeEnum string

// Set of constants representing the allowable values for DrPlanStepTypeEnum
const (
	DrPlanStepTypeComputeInstanceStopPrecheck                   DrPlanStepTypeEnum = "COMPUTE_INSTANCE_STOP_PRECHECK"
	DrPlanStepTypeComputeInstanceLaunchPrecheck                 DrPlanStepTypeEnum = "COMPUTE_INSTANCE_LAUNCH_PRECHECK"
	DrPlanStepTypeComputeInstanceTerminatePrecheck              DrPlanStepTypeEnum = "COMPUTE_INSTANCE_TERMINATE_PRECHECK"
	DrPlanStepTypeComputeInstanceRemovePrecheck                 DrPlanStepTypeEnum = "COMPUTE_INSTANCE_REMOVE_PRECHECK"
	DrPlanStepTypeVolumeGroupRestoreSwitchoverPrecheck          DrPlanStepTypeEnum = "VOLUME_GROUP_RESTORE_SWITCHOVER_PRECHECK"
	DrPlanStepTypeVolumeGroupRestoreFailoverPrecheck            DrPlanStepTypeEnum = "VOLUME_GROUP_RESTORE_FAILOVER_PRECHECK"
	DrPlanStepTypeDatabaseSwitchoverPrecheck                    DrPlanStepTypeEnum = "DATABASE_SWITCHOVER_PRECHECK"
	DrPlanStepTypeDatabaseFailoverPrecheck                      DrPlanStepTypeEnum = "DATABASE_FAILOVER_PRECHECK"
	DrPlanStepTypeAutonomousDatabaseSwitchoverPrecheck          DrPlanStepTypeEnum = "AUTONOMOUS_DATABASE_SWITCHOVER_PRECHECK"
	DrPlanStepTypeAutonomousDatabaseFailoverPrecheck            DrPlanStepTypeEnum = "AUTONOMOUS_DATABASE_FAILOVER_PRECHECK"
	DrPlanStepTypeUserDefinedPrecheck                           DrPlanStepTypeEnum = "USER_DEFINED_PRECHECK"
	DrPlanStepTypeComputeInstanceLaunch                         DrPlanStepTypeEnum = "COMPUTE_INSTANCE_LAUNCH"
	DrPlanStepTypeComputeInstanceStop                           DrPlanStepTypeEnum = "COMPUTE_INSTANCE_STOP"
	DrPlanStepTypeComputeInstanceTerminate                      DrPlanStepTypeEnum = "COMPUTE_INSTANCE_TERMINATE"
	DrPlanStepTypeComputeInstanceRemove                         DrPlanStepTypeEnum = "COMPUTE_INSTANCE_REMOVE"
	DrPlanStepTypeDatabaseSwitchover                            DrPlanStepTypeEnum = "DATABASE_SWITCHOVER"
	DrPlanStepTypeDatabaseFailover                              DrPlanStepTypeEnum = "DATABASE_FAILOVER"
	DrPlanStepTypeAutonomousDatabaseSwitchover                  DrPlanStepTypeEnum = "AUTONOMOUS_DATABASE_SWITCHOVER"
	DrPlanStepTypeAutonomousDatabaseFailover                    DrPlanStepTypeEnum = "AUTONOMOUS_DATABASE_FAILOVER"
	DrPlanStepTypeVolumeGroupRestoreSwitchover                  DrPlanStepTypeEnum = "VOLUME_GROUP_RESTORE_SWITCHOVER"
	DrPlanStepTypeVolumeGroupRestoreFailover                    DrPlanStepTypeEnum = "VOLUME_GROUP_RESTORE_FAILOVER"
	DrPlanStepTypeVolumeGroupReverse                            DrPlanStepTypeEnum = "VOLUME_GROUP_REVERSE"
	DrPlanStepTypeVolumeGroupDelete                             DrPlanStepTypeEnum = "VOLUME_GROUP_DELETE"
	DrPlanStepTypeVolumeGroupRemove                             DrPlanStepTypeEnum = "VOLUME_GROUP_REMOVE"
	DrPlanStepTypeVolumeGroupTerminate                          DrPlanStepTypeEnum = "VOLUME_GROUP_TERMINATE"
	DrPlanStepTypeUserDefined                                   DrPlanStepTypeEnum = "USER_DEFINED"
	DrPlanStepTypeComputeCapacityReservationSwitchoverPrecheck  DrPlanStepTypeEnum = "COMPUTE_CAPACITY_RESERVATION_SWITCHOVER_PRECHECK"
	DrPlanStepTypeComputeCapacityReservationFailoverPrecheck    DrPlanStepTypeEnum = "COMPUTE_CAPACITY_RESERVATION_FAILOVER_PRECHECK"
	DrPlanStepTypeComputeCapacityAvailabilitySwitchoverPrecheck DrPlanStepTypeEnum = "COMPUTE_CAPACITY_AVAILABILITY_SWITCHOVER_PRECHECK"
	DrPlanStepTypeComputeCapacityAvailabilityFailoverPrecheck   DrPlanStepTypeEnum = "COMPUTE_CAPACITY_AVAILABILITY_FAILOVER_PRECHECK"
)

var mappingDrPlanStepTypeEnum = map[string]DrPlanStepTypeEnum{
	"COMPUTE_INSTANCE_STOP_PRECHECK":                    DrPlanStepTypeComputeInstanceStopPrecheck,
	"COMPUTE_INSTANCE_LAUNCH_PRECHECK":                  DrPlanStepTypeComputeInstanceLaunchPrecheck,
	"COMPUTE_INSTANCE_TERMINATE_PRECHECK":               DrPlanStepTypeComputeInstanceTerminatePrecheck,
	"COMPUTE_INSTANCE_REMOVE_PRECHECK":                  DrPlanStepTypeComputeInstanceRemovePrecheck,
	"VOLUME_GROUP_RESTORE_SWITCHOVER_PRECHECK":          DrPlanStepTypeVolumeGroupRestoreSwitchoverPrecheck,
	"VOLUME_GROUP_RESTORE_FAILOVER_PRECHECK":            DrPlanStepTypeVolumeGroupRestoreFailoverPrecheck,
	"DATABASE_SWITCHOVER_PRECHECK":                      DrPlanStepTypeDatabaseSwitchoverPrecheck,
	"DATABASE_FAILOVER_PRECHECK":                        DrPlanStepTypeDatabaseFailoverPrecheck,
	"AUTONOMOUS_DATABASE_SWITCHOVER_PRECHECK":           DrPlanStepTypeAutonomousDatabaseSwitchoverPrecheck,
	"AUTONOMOUS_DATABASE_FAILOVER_PRECHECK":             DrPlanStepTypeAutonomousDatabaseFailoverPrecheck,
	"USER_DEFINED_PRECHECK":                             DrPlanStepTypeUserDefinedPrecheck,
	"COMPUTE_INSTANCE_LAUNCH":                           DrPlanStepTypeComputeInstanceLaunch,
	"COMPUTE_INSTANCE_STOP":                             DrPlanStepTypeComputeInstanceStop,
	"COMPUTE_INSTANCE_TERMINATE":                        DrPlanStepTypeComputeInstanceTerminate,
	"COMPUTE_INSTANCE_REMOVE":                           DrPlanStepTypeComputeInstanceRemove,
	"DATABASE_SWITCHOVER":                               DrPlanStepTypeDatabaseSwitchover,
	"DATABASE_FAILOVER":                                 DrPlanStepTypeDatabaseFailover,
	"AUTONOMOUS_DATABASE_SWITCHOVER":                    DrPlanStepTypeAutonomousDatabaseSwitchover,
	"AUTONOMOUS_DATABASE_FAILOVER":                      DrPlanStepTypeAutonomousDatabaseFailover,
	"VOLUME_GROUP_RESTORE_SWITCHOVER":                   DrPlanStepTypeVolumeGroupRestoreSwitchover,
	"VOLUME_GROUP_RESTORE_FAILOVER":                     DrPlanStepTypeVolumeGroupRestoreFailover,
	"VOLUME_GROUP_REVERSE":                              DrPlanStepTypeVolumeGroupReverse,
	"VOLUME_GROUP_DELETE":                               DrPlanStepTypeVolumeGroupDelete,
	"VOLUME_GROUP_REMOVE":                               DrPlanStepTypeVolumeGroupRemove,
	"VOLUME_GROUP_TERMINATE":                            DrPlanStepTypeVolumeGroupTerminate,
	"USER_DEFINED":                                      DrPlanStepTypeUserDefined,
	"COMPUTE_CAPACITY_RESERVATION_SWITCHOVER_PRECHECK":  DrPlanStepTypeComputeCapacityReservationSwitchoverPrecheck,
	"COMPUTE_CAPACITY_RESERVATION_FAILOVER_PRECHECK":    DrPlanStepTypeComputeCapacityReservationFailoverPrecheck,
	"COMPUTE_CAPACITY_AVAILABILITY_SWITCHOVER_PRECHECK": DrPlanStepTypeComputeCapacityAvailabilitySwitchoverPrecheck,
	"COMPUTE_CAPACITY_AVAILABILITY_FAILOVER_PRECHECK":   DrPlanStepTypeComputeCapacityAvailabilityFailoverPrecheck,
}

var mappingDrPlanStepTypeEnumLowerCase = map[string]DrPlanStepTypeEnum{
	"compute_instance_stop_precheck":                    DrPlanStepTypeComputeInstanceStopPrecheck,
	"compute_instance_launch_precheck":                  DrPlanStepTypeComputeInstanceLaunchPrecheck,
	"compute_instance_terminate_precheck":               DrPlanStepTypeComputeInstanceTerminatePrecheck,
	"compute_instance_remove_precheck":                  DrPlanStepTypeComputeInstanceRemovePrecheck,
	"volume_group_restore_switchover_precheck":          DrPlanStepTypeVolumeGroupRestoreSwitchoverPrecheck,
	"volume_group_restore_failover_precheck":            DrPlanStepTypeVolumeGroupRestoreFailoverPrecheck,
	"database_switchover_precheck":                      DrPlanStepTypeDatabaseSwitchoverPrecheck,
	"database_failover_precheck":                        DrPlanStepTypeDatabaseFailoverPrecheck,
	"autonomous_database_switchover_precheck":           DrPlanStepTypeAutonomousDatabaseSwitchoverPrecheck,
	"autonomous_database_failover_precheck":             DrPlanStepTypeAutonomousDatabaseFailoverPrecheck,
	"user_defined_precheck":                             DrPlanStepTypeUserDefinedPrecheck,
	"compute_instance_launch":                           DrPlanStepTypeComputeInstanceLaunch,
	"compute_instance_stop":                             DrPlanStepTypeComputeInstanceStop,
	"compute_instance_terminate":                        DrPlanStepTypeComputeInstanceTerminate,
	"compute_instance_remove":                           DrPlanStepTypeComputeInstanceRemove,
	"database_switchover":                               DrPlanStepTypeDatabaseSwitchover,
	"database_failover":                                 DrPlanStepTypeDatabaseFailover,
	"autonomous_database_switchover":                    DrPlanStepTypeAutonomousDatabaseSwitchover,
	"autonomous_database_failover":                      DrPlanStepTypeAutonomousDatabaseFailover,
	"volume_group_restore_switchover":                   DrPlanStepTypeVolumeGroupRestoreSwitchover,
	"volume_group_restore_failover":                     DrPlanStepTypeVolumeGroupRestoreFailover,
	"volume_group_reverse":                              DrPlanStepTypeVolumeGroupReverse,
	"volume_group_delete":                               DrPlanStepTypeVolumeGroupDelete,
	"volume_group_remove":                               DrPlanStepTypeVolumeGroupRemove,
	"volume_group_terminate":                            DrPlanStepTypeVolumeGroupTerminate,
	"user_defined":                                      DrPlanStepTypeUserDefined,
	"compute_capacity_reservation_switchover_precheck":  DrPlanStepTypeComputeCapacityReservationSwitchoverPrecheck,
	"compute_capacity_reservation_failover_precheck":    DrPlanStepTypeComputeCapacityReservationFailoverPrecheck,
	"compute_capacity_availability_switchover_precheck": DrPlanStepTypeComputeCapacityAvailabilitySwitchoverPrecheck,
	"compute_capacity_availability_failover_precheck":   DrPlanStepTypeComputeCapacityAvailabilityFailoverPrecheck,
}

// GetDrPlanStepTypeEnumValues Enumerates the set of values for DrPlanStepTypeEnum
func GetDrPlanStepTypeEnumValues() []DrPlanStepTypeEnum {
	values := make([]DrPlanStepTypeEnum, 0)
	for _, v := range mappingDrPlanStepTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDrPlanStepTypeEnumStringValues Enumerates the set of values in String for DrPlanStepTypeEnum
func GetDrPlanStepTypeEnumStringValues() []string {
	return []string{
		"COMPUTE_INSTANCE_STOP_PRECHECK",
		"COMPUTE_INSTANCE_LAUNCH_PRECHECK",
		"COMPUTE_INSTANCE_TERMINATE_PRECHECK",
		"COMPUTE_INSTANCE_REMOVE_PRECHECK",
		"VOLUME_GROUP_RESTORE_SWITCHOVER_PRECHECK",
		"VOLUME_GROUP_RESTORE_FAILOVER_PRECHECK",
		"DATABASE_SWITCHOVER_PRECHECK",
		"DATABASE_FAILOVER_PRECHECK",
		"AUTONOMOUS_DATABASE_SWITCHOVER_PRECHECK",
		"AUTONOMOUS_DATABASE_FAILOVER_PRECHECK",
		"USER_DEFINED_PRECHECK",
		"COMPUTE_INSTANCE_LAUNCH",
		"COMPUTE_INSTANCE_STOP",
		"COMPUTE_INSTANCE_TERMINATE",
		"COMPUTE_INSTANCE_REMOVE",
		"DATABASE_SWITCHOVER",
		"DATABASE_FAILOVER",
		"AUTONOMOUS_DATABASE_SWITCHOVER",
		"AUTONOMOUS_DATABASE_FAILOVER",
		"VOLUME_GROUP_RESTORE_SWITCHOVER",
		"VOLUME_GROUP_RESTORE_FAILOVER",
		"VOLUME_GROUP_REVERSE",
		"VOLUME_GROUP_DELETE",
		"VOLUME_GROUP_REMOVE",
		"VOLUME_GROUP_TERMINATE",
		"USER_DEFINED",
		"COMPUTE_CAPACITY_RESERVATION_SWITCHOVER_PRECHECK",
		"COMPUTE_CAPACITY_RESERVATION_FAILOVER_PRECHECK",
		"COMPUTE_CAPACITY_AVAILABILITY_SWITCHOVER_PRECHECK",
		"COMPUTE_CAPACITY_AVAILABILITY_FAILOVER_PRECHECK",
	}
}

// GetMappingDrPlanStepTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDrPlanStepTypeEnum(val string) (DrPlanStepTypeEnum, bool) {
	enum, ok := mappingDrPlanStepTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
