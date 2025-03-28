// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Full Stack Disaster Recovery API
//
// Use the Full Stack Disaster Recovery (DR) API to manage disaster recovery for business applications.
// Full Stack DR is an OCI disaster recovery orchestration and management service that provides comprehensive disaster
// recovery capabilities for all layers of an application stack, including infrastructure, middleware, database,
// and application.
//

package disasterrecovery

import (
	"strings"
)

// DrPlanStepTypeEnum Enum with underlying type: string
type DrPlanStepTypeEnum string

// Set of constants representing the allowable values for DrPlanStepTypeEnum
const (
	DrPlanStepTypeComputeInstanceStopPrecheck                                           DrPlanStepTypeEnum = "COMPUTE_INSTANCE_STOP_PRECHECK"
	DrPlanStepTypeComputeInstanceLaunchPrecheck                                         DrPlanStepTypeEnum = "COMPUTE_INSTANCE_LAUNCH_PRECHECK"
	DrPlanStepTypeComputeInstanceTerminatePrecheck                                      DrPlanStepTypeEnum = "COMPUTE_INSTANCE_TERMINATE_PRECHECK"
	DrPlanStepTypeComputeInstanceRemovePrecheck                                         DrPlanStepTypeEnum = "COMPUTE_INSTANCE_REMOVE_PRECHECK"
	DrPlanStepTypeVolumeGroupRestoreSwitchoverPrecheck                                  DrPlanStepTypeEnum = "VOLUME_GROUP_RESTORE_SWITCHOVER_PRECHECK"
	DrPlanStepTypeVolumeGroupRestoreFailoverPrecheck                                    DrPlanStepTypeEnum = "VOLUME_GROUP_RESTORE_FAILOVER_PRECHECK"
	DrPlanStepTypeDatabaseSwitchoverPrecheck                                            DrPlanStepTypeEnum = "DATABASE_SWITCHOVER_PRECHECK"
	DrPlanStepTypeDatabaseFailoverPrecheck                                              DrPlanStepTypeEnum = "DATABASE_FAILOVER_PRECHECK"
	DrPlanStepTypeAutonomousDatabaseSwitchoverPrecheck                                  DrPlanStepTypeEnum = "AUTONOMOUS_DATABASE_SWITCHOVER_PRECHECK"
	DrPlanStepTypeAutonomousDatabaseFailoverPrecheck                                    DrPlanStepTypeEnum = "AUTONOMOUS_DATABASE_FAILOVER_PRECHECK"
	DrPlanStepTypeAutonomousContainerDatabaseSwitchoverPrecheck                         DrPlanStepTypeEnum = "AUTONOMOUS_CONTAINER_DATABASE_SWITCHOVER_PRECHECK"
	DrPlanStepTypeAutonomousContainerDatabaseFailoverPrecheck                           DrPlanStepTypeEnum = "AUTONOMOUS_CONTAINER_DATABASE_FAILOVER_PRECHECK"
	DrPlanStepTypeAutonomousContainerDatabaseStartDrillConvertToSnapshotStandbyPrecheck DrPlanStepTypeEnum = "AUTONOMOUS_CONTAINER_DATABASE_START_DRILL_CONVERT_TO_SNAPSHOT_STANDBY_PRECHECK"
	DrPlanStepTypeAutonomousContainerDatabaseStopDrillConvertToPhysicalStandbyPrecheck  DrPlanStepTypeEnum = "AUTONOMOUS_CONTAINER_DATABASE_STOP_DRILL_CONVERT_TO_PHYSICAL_STANDBY_PRECHECK"
	DrPlanStepTypeAutonomousContainerDatabaseSwitchover                                 DrPlanStepTypeEnum = "AUTONOMOUS_CONTAINER_DATABASE_SWITCHOVER"
	DrPlanStepTypeAutonomousContainerDatabaseFailover                                   DrPlanStepTypeEnum = "AUTONOMOUS_CONTAINER_DATABASE_FAILOVER"
	DrPlanStepTypeAutonomousContainerDatabaseStartDrillConvertToSnapshotStandby         DrPlanStepTypeEnum = "AUTONOMOUS_CONTAINER_DATABASE_START_DRILL_CONVERT_TO_SNAPSHOT_STANDBY"
	DrPlanStepTypeAutonomousContainerDatabaseStopDrillConvertToPhysicalStandby          DrPlanStepTypeEnum = "AUTONOMOUS_CONTAINER_DATABASE_STOP_DRILL_CONVERT_TO_PHYSICAL_STANDBY"
	DrPlanStepTypeAutonomousDatabaseStartDrillCreateCloneStandbyPrecheck                DrPlanStepTypeEnum = "AUTONOMOUS_DATABASE_START_DRILL_CREATE_CLONE_STANDBY_PRECHECK"
	DrPlanStepTypeAutonomousDatabaseStopDrillDeleteCloneStandbyPrecheck                 DrPlanStepTypeEnum = "AUTONOMOUS_DATABASE_STOP_DRILL_DELETE_CLONE_STANDBY_PRECHECK"
	DrPlanStepTypeAutonomousDatabaseStartDrillConvertToSnapshotStandbyPrecheck          DrPlanStepTypeEnum = "AUTONOMOUS_DATABASE_START_DRILL_CONVERT_TO_SNAPSHOT_STANDBY_PRECHECK"
	DrPlanStepTypeAutonomousDatabaseStopDrillConvertToPhysicalStandbyPrecheck           DrPlanStepTypeEnum = "AUTONOMOUS_DATABASE_STOP_DRILL_CONVERT_TO_PHYSICAL_STANDBY_PRECHECK"
	DrPlanStepTypeAutonomousDatabaseStartDrillCreateCloneStandby                        DrPlanStepTypeEnum = "AUTONOMOUS_DATABASE_START_DRILL_CREATE_CLONE_STANDBY"
	DrPlanStepTypeAutonomousDatabaseStopDrillDeleteCloneStandby                         DrPlanStepTypeEnum = "AUTONOMOUS_DATABASE_STOP_DRILL_DELETE_CLONE_STANDBY"
	DrPlanStepTypeAutonomousDatabaseStartDrillConvertToSnapshotStandby                  DrPlanStepTypeEnum = "AUTONOMOUS_DATABASE_START_DRILL_CONVERT_TO_SNAPSHOT_STANDBY"
	DrPlanStepTypeAutonomousDatabaseStopDrillConvertToPhysicalStandby                   DrPlanStepTypeEnum = "AUTONOMOUS_DATABASE_STOP_DRILL_CONVERT_TO_PHYSICAL_STANDBY"
	DrPlanStepTypeUserDefinedPrecheck                                                   DrPlanStepTypeEnum = "USER_DEFINED_PRECHECK"
	DrPlanStepTypeComputeInstanceLaunch                                                 DrPlanStepTypeEnum = "COMPUTE_INSTANCE_LAUNCH"
	DrPlanStepTypeComputeInstanceStop                                                   DrPlanStepTypeEnum = "COMPUTE_INSTANCE_STOP"
	DrPlanStepTypeComputeInstanceTerminate                                              DrPlanStepTypeEnum = "COMPUTE_INSTANCE_TERMINATE"
	DrPlanStepTypeComputeInstanceRemove                                                 DrPlanStepTypeEnum = "COMPUTE_INSTANCE_REMOVE"
	DrPlanStepTypeDatabaseSwitchover                                                    DrPlanStepTypeEnum = "DATABASE_SWITCHOVER"
	DrPlanStepTypeDatabaseFailover                                                      DrPlanStepTypeEnum = "DATABASE_FAILOVER"
	DrPlanStepTypeAutonomousDatabaseSwitchover                                          DrPlanStepTypeEnum = "AUTONOMOUS_DATABASE_SWITCHOVER"
	DrPlanStepTypeAutonomousDatabaseFailover                                            DrPlanStepTypeEnum = "AUTONOMOUS_DATABASE_FAILOVER"
	DrPlanStepTypeVolumeGroupRestoreSwitchover                                          DrPlanStepTypeEnum = "VOLUME_GROUP_RESTORE_SWITCHOVER"
	DrPlanStepTypeVolumeGroupRestoreFailover                                            DrPlanStepTypeEnum = "VOLUME_GROUP_RESTORE_FAILOVER"
	DrPlanStepTypeVolumeGroupReverse                                                    DrPlanStepTypeEnum = "VOLUME_GROUP_REVERSE"
	DrPlanStepTypeVolumeGroupDelete                                                     DrPlanStepTypeEnum = "VOLUME_GROUP_DELETE"
	DrPlanStepTypeVolumeGroupRemove                                                     DrPlanStepTypeEnum = "VOLUME_GROUP_REMOVE"
	DrPlanStepTypeVolumeGroupTerminate                                                  DrPlanStepTypeEnum = "VOLUME_GROUP_TERMINATE"
	DrPlanStepTypeUserDefined                                                           DrPlanStepTypeEnum = "USER_DEFINED"
	DrPlanStepTypeVolumeGroupRestoreStartDrillPrecheck                                  DrPlanStepTypeEnum = "VOLUME_GROUP_RESTORE_START_DRILL_PRECHECK"
	DrPlanStepTypeVolumeGroupRemovePrecheck                                             DrPlanStepTypeEnum = "VOLUME_GROUP_REMOVE_PRECHECK"
	DrPlanStepTypeVolumeGroupTerminatePrecheck                                          DrPlanStepTypeEnum = "VOLUME_GROUP_TERMINATE_PRECHECK"
	DrPlanStepTypeVolumeGroupRestoreStartDrill                                          DrPlanStepTypeEnum = "VOLUME_GROUP_RESTORE_START_DRILL"
	DrPlanStepTypeAutonomousDatabaseCreateClonePrecheck                                 DrPlanStepTypeEnum = "AUTONOMOUS_DATABASE_CREATE_CLONE_PRECHECK"
	DrPlanStepTypeAutonomousDatabaseDeleteClonePrecheck                                 DrPlanStepTypeEnum = "AUTONOMOUS_DATABASE_DELETE_CLONE_PRECHECK"
	DrPlanStepTypeLoadBalancerUpdatePrimaryBackendSetPrecheck                           DrPlanStepTypeEnum = "LOAD_BALANCER_UPDATE_PRIMARY_BACKEND_SET_PRECHECK"
	DrPlanStepTypeLoadBalancerUpdateStandbyBackendSetPrecheck                           DrPlanStepTypeEnum = "LOAD_BALANCER_UPDATE_STANDBY_BACKEND_SET_PRECHECK"
	DrPlanStepTypeFileSystemSwitchoverPrecheck                                          DrPlanStepTypeEnum = "FILE_SYSTEM_SWITCHOVER_PRECHECK"
	DrPlanStepTypeFileSystemFailoverPrecheck                                            DrPlanStepTypeEnum = "FILE_SYSTEM_FAILOVER_PRECHECK"
	DrPlanStepTypeFileSystemStartDrillPrecheck                                          DrPlanStepTypeEnum = "FILE_SYSTEM_START_DRILL_PRECHECK"
	DrPlanStepTypeFileSystemStopDrillPrecheck                                           DrPlanStepTypeEnum = "FILE_SYSTEM_STOP_DRILL_PRECHECK"
	DrPlanStepTypeFileSystemRemovePrecheck                                              DrPlanStepTypeEnum = "FILE_SYSTEM_REMOVE_PRECHECK"
	DrPlanStepTypeFileSystemTerminatePrecheck                                           DrPlanStepTypeEnum = "FILE_SYSTEM_TERMINATE_PRECHECK"
	DrPlanStepTypeFileSystemMountPrecheck                                               DrPlanStepTypeEnum = "FILE_SYSTEM_MOUNT_PRECHECK"
	DrPlanStepTypeFileSystemUnmountPrecheck                                             DrPlanStepTypeEnum = "FILE_SYSTEM_UNMOUNT_PRECHECK"
	DrPlanStepTypeComputeInstanceStartPrecheck                                          DrPlanStepTypeEnum = "COMPUTE_INSTANCE_START_PRECHECK"
	DrPlanStepTypeComputeInstanceAttachBlockVolumesPrecheck                             DrPlanStepTypeEnum = "COMPUTE_INSTANCE_ATTACH_BLOCK_VOLUMES_PRECHECK"
	DrPlanStepTypeComputeInstanceDetachBlockVolumesPrecheck                             DrPlanStepTypeEnum = "COMPUTE_INSTANCE_DETACH_BLOCK_VOLUMES_PRECHECK"
	DrPlanStepTypeComputeInstanceMountBlockVolumesPrecheck                              DrPlanStepTypeEnum = "COMPUTE_INSTANCE_MOUNT_BLOCK_VOLUMES_PRECHECK"
	DrPlanStepTypeComputeInstanceUnmountBlockVolumesPrecheck                            DrPlanStepTypeEnum = "COMPUTE_INSTANCE_UNMOUNT_BLOCK_VOLUMES_PRECHECK"
	DrPlanStepTypeComputeCapacityReservationStartDrillPrecheck                          DrPlanStepTypeEnum = "COMPUTE_CAPACITY_RESERVATION_START_DRILL_PRECHECK"
	DrPlanStepTypeComputeCapacityAvailabilityStartDrillPrecheck                         DrPlanStepTypeEnum = "COMPUTE_CAPACITY_AVAILABILITY_START_DRILL_PRECHECK"
	DrPlanStepTypeAutonomousDatabaseCreateClone                                         DrPlanStepTypeEnum = "AUTONOMOUS_DATABASE_CREATE_CLONE"
	DrPlanStepTypeAutonomousDatabaseDeleteClone                                         DrPlanStepTypeEnum = "AUTONOMOUS_DATABASE_DELETE_CLONE"
	DrPlanStepTypeLoadBalancerUpdatePrimaryBackendSet                                   DrPlanStepTypeEnum = "LOAD_BALANCER_UPDATE_PRIMARY_BACKEND_SET"
	DrPlanStepTypeLoadBalancerUpdateStandbyBackendSet                                   DrPlanStepTypeEnum = "LOAD_BALANCER_UPDATE_STANDBY_BACKEND_SET"
	DrPlanStepTypeFileSystemSwitchover                                                  DrPlanStepTypeEnum = "FILE_SYSTEM_SWITCHOVER"
	DrPlanStepTypeFileSystemFailover                                                    DrPlanStepTypeEnum = "FILE_SYSTEM_FAILOVER"
	DrPlanStepTypeFileSystemRemove                                                      DrPlanStepTypeEnum = "FILE_SYSTEM_REMOVE"
	DrPlanStepTypeFileSystemReverse                                                     DrPlanStepTypeEnum = "FILE_SYSTEM_REVERSE"
	DrPlanStepTypeFileSystemTerminate                                                   DrPlanStepTypeEnum = "FILE_SYSTEM_TERMINATE"
	DrPlanStepTypeFileSystemStartDrill                                                  DrPlanStepTypeEnum = "FILE_SYSTEM_START_DRILL"
	DrPlanStepTypeFileSystemStopDrill                                                   DrPlanStepTypeEnum = "FILE_SYSTEM_STOP_DRILL"
	DrPlanStepTypeComputeInstanceStart                                                  DrPlanStepTypeEnum = "COMPUTE_INSTANCE_START"
	DrPlanStepTypeComputeInstanceAttachBlockVolumes                                     DrPlanStepTypeEnum = "COMPUTE_INSTANCE_ATTACH_BLOCK_VOLUMES"
	DrPlanStepTypeComputeInstanceDetachBlockVolumes                                     DrPlanStepTypeEnum = "COMPUTE_INSTANCE_DETACH_BLOCK_VOLUMES"
	DrPlanStepTypeFileSystemMount                                                       DrPlanStepTypeEnum = "FILE_SYSTEM_MOUNT"
	DrPlanStepTypeFileSystemUnmount                                                     DrPlanStepTypeEnum = "FILE_SYSTEM_UNMOUNT"
	DrPlanStepTypeComputeCapacityReservationSwitchoverPrecheck                          DrPlanStepTypeEnum = "COMPUTE_CAPACITY_RESERVATION_SWITCHOVER_PRECHECK"
	DrPlanStepTypeComputeCapacityReservationFailoverPrecheck                            DrPlanStepTypeEnum = "COMPUTE_CAPACITY_RESERVATION_FAILOVER_PRECHECK"
	DrPlanStepTypeComputeCapacityAvailabilitySwitchoverPrecheck                         DrPlanStepTypeEnum = "COMPUTE_CAPACITY_AVAILABILITY_SWITCHOVER_PRECHECK"
	DrPlanStepTypeComputeCapacityAvailabilityFailoverPrecheck                           DrPlanStepTypeEnum = "COMPUTE_CAPACITY_AVAILABILITY_FAILOVER_PRECHECK"
	DrPlanStepTypeOkeClusterSwitchoverCleanupPrimaryPrecheck                            DrPlanStepTypeEnum = "OKE_CLUSTER_SWITCHOVER_CLEANUP_PRIMARY_PRECHECK"
	DrPlanStepTypeOkeClusterSwitchoverRestoreStandbyPrecheck                            DrPlanStepTypeEnum = "OKE_CLUSTER_SWITCHOVER_RESTORE_STANDBY_PRECHECK"
	DrPlanStepTypeOkeClusterFailoverRestoreStandbyPrecheck                              DrPlanStepTypeEnum = "OKE_CLUSTER_FAILOVER_RESTORE_STANDBY_PRECHECK"
	DrPlanStepTypeOkeClusterStopDrillCleanupStandbyPrecheck                             DrPlanStepTypeEnum = "OKE_CLUSTER_STOP_DRILL_CLEANUP_STANDBY_PRECHECK"
	DrPlanStepTypeOkeClusterStartDrillRestoreStandbyPrecheck                            DrPlanStepTypeEnum = "OKE_CLUSTER_START_DRILL_RESTORE_STANDBY_PRECHECK"
	DrPlanStepTypeOkeClusterSwitchoverCleanupPrimary                                    DrPlanStepTypeEnum = "OKE_CLUSTER_SWITCHOVER_CLEANUP_PRIMARY"
	DrPlanStepTypeOkeClusterSwitchoverScaleUpStandby                                    DrPlanStepTypeEnum = "OKE_CLUSTER_SWITCHOVER_SCALE_UP_STANDBY"
	DrPlanStepTypeOkeClusterSwitchoverRestoreStandby                                    DrPlanStepTypeEnum = "OKE_CLUSTER_SWITCHOVER_RESTORE_STANDBY"
	DrPlanStepTypeOkeClusterSwitchoverReverseBackupStandby                              DrPlanStepTypeEnum = "OKE_CLUSTER_SWITCHOVER_REVERSE_BACKUP_STANDBY"
	DrPlanStepTypeOkeClusterSwitchoverScaleDownPrimary                                  DrPlanStepTypeEnum = "OKE_CLUSTER_SWITCHOVER_SCALE_DOWN_PRIMARY"
	DrPlanStepTypeOkeClusterFailoverRestoreStandby                                      DrPlanStepTypeEnum = "OKE_CLUSTER_FAILOVER_RESTORE_STANDBY"
	DrPlanStepTypeOkeClusterFailoverScaleUpStandby                                      DrPlanStepTypeEnum = "OKE_CLUSTER_FAILOVER_SCALE_UP_STANDBY"
	DrPlanStepTypeOkeClusterStartDrillRestoreStandby                                    DrPlanStepTypeEnum = "OKE_CLUSTER_START_DRILL_RESTORE_STANDBY"
	DrPlanStepTypeOkeClusterStartDrillScaleUpStandby                                    DrPlanStepTypeEnum = "OKE_CLUSTER_START_DRILL_SCALE_UP_STANDBY"
	DrPlanStepTypeOkeClusterStopDrillCleanupStandby                                     DrPlanStepTypeEnum = "OKE_CLUSTER_STOP_DRILL_CLEANUP_STANDBY"
	DrPlanStepTypeOkeClusterStopDrillScaleDownStandby                                   DrPlanStepTypeEnum = "OKE_CLUSTER_STOP_DRILL_SCALE_DOWN_STANDBY"
	DrPlanStepTypeObjectStorageBucketSwitchoverDeleteReplicationPrimaryPrecheck         DrPlanStepTypeEnum = "OBJECT_STORAGE_BUCKET_SWITCHOVER_DELETE_REPLICATION_PRIMARY_PRECHECK"
	DrPlanStepTypeObjectStorageBucketSwitchoverSetupReverseReplicationStandbyPrecheck   DrPlanStepTypeEnum = "OBJECT_STORAGE_BUCKET_SWITCHOVER_SETUP_REVERSE_REPLICATION_STANDBY_PRECHECK"
	DrPlanStepTypeObjectStorageBucketFailoverDeleteReplicationStandbyPrecheck           DrPlanStepTypeEnum = "OBJECT_STORAGE_BUCKET_FAILOVER_DELETE_REPLICATION_STANDBY_PRECHECK"
	DrPlanStepTypeObjectStorageBucketFailoverSetupReverseReplicationStandbyPrecheck     DrPlanStepTypeEnum = "OBJECT_STORAGE_BUCKET_FAILOVER_SETUP_REVERSE_REPLICATION_STANDBY_PRECHECK"
	DrPlanStepTypeObjectStorageBucketSwitchoverDeleteReplicationPrimary                 DrPlanStepTypeEnum = "OBJECT_STORAGE_BUCKET_SWITCHOVER_DELETE_REPLICATION_PRIMARY"
	DrPlanStepTypeObjectStorageBucketSwitchoverSetupReverseReplicationStandby           DrPlanStepTypeEnum = "OBJECT_STORAGE_BUCKET_SWITCHOVER_SETUP_REVERSE_REPLICATION_STANDBY"
	DrPlanStepTypeObjectStorageBucketFailoverDeleteReplicationStandby                   DrPlanStepTypeEnum = "OBJECT_STORAGE_BUCKET_FAILOVER_DELETE_REPLICATION_STANDBY"
	DrPlanStepTypeObjectStorageBucketFailoverSetupReverseReplicationStandby             DrPlanStepTypeEnum = "OBJECT_STORAGE_BUCKET_FAILOVER_SETUP_REVERSE_REPLICATION_STANDBY"
)

var mappingDrPlanStepTypeEnum = map[string]DrPlanStepTypeEnum{
	"COMPUTE_INSTANCE_STOP_PRECHECK":                                                 DrPlanStepTypeComputeInstanceStopPrecheck,
	"COMPUTE_INSTANCE_LAUNCH_PRECHECK":                                               DrPlanStepTypeComputeInstanceLaunchPrecheck,
	"COMPUTE_INSTANCE_TERMINATE_PRECHECK":                                            DrPlanStepTypeComputeInstanceTerminatePrecheck,
	"COMPUTE_INSTANCE_REMOVE_PRECHECK":                                               DrPlanStepTypeComputeInstanceRemovePrecheck,
	"VOLUME_GROUP_RESTORE_SWITCHOVER_PRECHECK":                                       DrPlanStepTypeVolumeGroupRestoreSwitchoverPrecheck,
	"VOLUME_GROUP_RESTORE_FAILOVER_PRECHECK":                                         DrPlanStepTypeVolumeGroupRestoreFailoverPrecheck,
	"DATABASE_SWITCHOVER_PRECHECK":                                                   DrPlanStepTypeDatabaseSwitchoverPrecheck,
	"DATABASE_FAILOVER_PRECHECK":                                                     DrPlanStepTypeDatabaseFailoverPrecheck,
	"AUTONOMOUS_DATABASE_SWITCHOVER_PRECHECK":                                        DrPlanStepTypeAutonomousDatabaseSwitchoverPrecheck,
	"AUTONOMOUS_DATABASE_FAILOVER_PRECHECK":                                          DrPlanStepTypeAutonomousDatabaseFailoverPrecheck,
	"AUTONOMOUS_CONTAINER_DATABASE_SWITCHOVER_PRECHECK":                              DrPlanStepTypeAutonomousContainerDatabaseSwitchoverPrecheck,
	"AUTONOMOUS_CONTAINER_DATABASE_FAILOVER_PRECHECK":                                DrPlanStepTypeAutonomousContainerDatabaseFailoverPrecheck,
	"AUTONOMOUS_CONTAINER_DATABASE_START_DRILL_CONVERT_TO_SNAPSHOT_STANDBY_PRECHECK": DrPlanStepTypeAutonomousContainerDatabaseStartDrillConvertToSnapshotStandbyPrecheck,
	"AUTONOMOUS_CONTAINER_DATABASE_STOP_DRILL_CONVERT_TO_PHYSICAL_STANDBY_PRECHECK":  DrPlanStepTypeAutonomousContainerDatabaseStopDrillConvertToPhysicalStandbyPrecheck,
	"AUTONOMOUS_CONTAINER_DATABASE_SWITCHOVER":                                       DrPlanStepTypeAutonomousContainerDatabaseSwitchover,
	"AUTONOMOUS_CONTAINER_DATABASE_FAILOVER":                                         DrPlanStepTypeAutonomousContainerDatabaseFailover,
	"AUTONOMOUS_CONTAINER_DATABASE_START_DRILL_CONVERT_TO_SNAPSHOT_STANDBY":          DrPlanStepTypeAutonomousContainerDatabaseStartDrillConvertToSnapshotStandby,
	"AUTONOMOUS_CONTAINER_DATABASE_STOP_DRILL_CONVERT_TO_PHYSICAL_STANDBY":           DrPlanStepTypeAutonomousContainerDatabaseStopDrillConvertToPhysicalStandby,
	"AUTONOMOUS_DATABASE_START_DRILL_CREATE_CLONE_STANDBY_PRECHECK":                  DrPlanStepTypeAutonomousDatabaseStartDrillCreateCloneStandbyPrecheck,
	"AUTONOMOUS_DATABASE_STOP_DRILL_DELETE_CLONE_STANDBY_PRECHECK":                   DrPlanStepTypeAutonomousDatabaseStopDrillDeleteCloneStandbyPrecheck,
	"AUTONOMOUS_DATABASE_START_DRILL_CONVERT_TO_SNAPSHOT_STANDBY_PRECHECK":           DrPlanStepTypeAutonomousDatabaseStartDrillConvertToSnapshotStandbyPrecheck,
	"AUTONOMOUS_DATABASE_STOP_DRILL_CONVERT_TO_PHYSICAL_STANDBY_PRECHECK":            DrPlanStepTypeAutonomousDatabaseStopDrillConvertToPhysicalStandbyPrecheck,
	"AUTONOMOUS_DATABASE_START_DRILL_CREATE_CLONE_STANDBY":                           DrPlanStepTypeAutonomousDatabaseStartDrillCreateCloneStandby,
	"AUTONOMOUS_DATABASE_STOP_DRILL_DELETE_CLONE_STANDBY":                            DrPlanStepTypeAutonomousDatabaseStopDrillDeleteCloneStandby,
	"AUTONOMOUS_DATABASE_START_DRILL_CONVERT_TO_SNAPSHOT_STANDBY":                    DrPlanStepTypeAutonomousDatabaseStartDrillConvertToSnapshotStandby,
	"AUTONOMOUS_DATABASE_STOP_DRILL_CONVERT_TO_PHYSICAL_STANDBY":                     DrPlanStepTypeAutonomousDatabaseStopDrillConvertToPhysicalStandby,
	"USER_DEFINED_PRECHECK":                                                          DrPlanStepTypeUserDefinedPrecheck,
	"COMPUTE_INSTANCE_LAUNCH":                                                        DrPlanStepTypeComputeInstanceLaunch,
	"COMPUTE_INSTANCE_STOP":                                                          DrPlanStepTypeComputeInstanceStop,
	"COMPUTE_INSTANCE_TERMINATE":                                                     DrPlanStepTypeComputeInstanceTerminate,
	"COMPUTE_INSTANCE_REMOVE":                                                        DrPlanStepTypeComputeInstanceRemove,
	"DATABASE_SWITCHOVER":                                                            DrPlanStepTypeDatabaseSwitchover,
	"DATABASE_FAILOVER":                                                              DrPlanStepTypeDatabaseFailover,
	"AUTONOMOUS_DATABASE_SWITCHOVER":                                                 DrPlanStepTypeAutonomousDatabaseSwitchover,
	"AUTONOMOUS_DATABASE_FAILOVER":                                                   DrPlanStepTypeAutonomousDatabaseFailover,
	"VOLUME_GROUP_RESTORE_SWITCHOVER":                                                DrPlanStepTypeVolumeGroupRestoreSwitchover,
	"VOLUME_GROUP_RESTORE_FAILOVER":                                                  DrPlanStepTypeVolumeGroupRestoreFailover,
	"VOLUME_GROUP_REVERSE":                                                           DrPlanStepTypeVolumeGroupReverse,
	"VOLUME_GROUP_DELETE":                                                            DrPlanStepTypeVolumeGroupDelete,
	"VOLUME_GROUP_REMOVE":                                                            DrPlanStepTypeVolumeGroupRemove,
	"VOLUME_GROUP_TERMINATE":                                                         DrPlanStepTypeVolumeGroupTerminate,
	"USER_DEFINED":                                                                   DrPlanStepTypeUserDefined,
	"VOLUME_GROUP_RESTORE_START_DRILL_PRECHECK":                                      DrPlanStepTypeVolumeGroupRestoreStartDrillPrecheck,
	"VOLUME_GROUP_REMOVE_PRECHECK":                                                   DrPlanStepTypeVolumeGroupRemovePrecheck,
	"VOLUME_GROUP_TERMINATE_PRECHECK":                                                DrPlanStepTypeVolumeGroupTerminatePrecheck,
	"VOLUME_GROUP_RESTORE_START_DRILL":                                               DrPlanStepTypeVolumeGroupRestoreStartDrill,
	"AUTONOMOUS_DATABASE_CREATE_CLONE_PRECHECK":                                      DrPlanStepTypeAutonomousDatabaseCreateClonePrecheck,
	"AUTONOMOUS_DATABASE_DELETE_CLONE_PRECHECK":                                      DrPlanStepTypeAutonomousDatabaseDeleteClonePrecheck,
	"LOAD_BALANCER_UPDATE_PRIMARY_BACKEND_SET_PRECHECK":                              DrPlanStepTypeLoadBalancerUpdatePrimaryBackendSetPrecheck,
	"LOAD_BALANCER_UPDATE_STANDBY_BACKEND_SET_PRECHECK":                              DrPlanStepTypeLoadBalancerUpdateStandbyBackendSetPrecheck,
	"FILE_SYSTEM_SWITCHOVER_PRECHECK":                                                DrPlanStepTypeFileSystemSwitchoverPrecheck,
	"FILE_SYSTEM_FAILOVER_PRECHECK":                                                  DrPlanStepTypeFileSystemFailoverPrecheck,
	"FILE_SYSTEM_START_DRILL_PRECHECK":                                               DrPlanStepTypeFileSystemStartDrillPrecheck,
	"FILE_SYSTEM_STOP_DRILL_PRECHECK":                                                DrPlanStepTypeFileSystemStopDrillPrecheck,
	"FILE_SYSTEM_REMOVE_PRECHECK":                                                    DrPlanStepTypeFileSystemRemovePrecheck,
	"FILE_SYSTEM_TERMINATE_PRECHECK":                                                 DrPlanStepTypeFileSystemTerminatePrecheck,
	"FILE_SYSTEM_MOUNT_PRECHECK":                                                     DrPlanStepTypeFileSystemMountPrecheck,
	"FILE_SYSTEM_UNMOUNT_PRECHECK":                                                   DrPlanStepTypeFileSystemUnmountPrecheck,
	"COMPUTE_INSTANCE_START_PRECHECK":                                                DrPlanStepTypeComputeInstanceStartPrecheck,
	"COMPUTE_INSTANCE_ATTACH_BLOCK_VOLUMES_PRECHECK":                                 DrPlanStepTypeComputeInstanceAttachBlockVolumesPrecheck,
	"COMPUTE_INSTANCE_DETACH_BLOCK_VOLUMES_PRECHECK":                                 DrPlanStepTypeComputeInstanceDetachBlockVolumesPrecheck,
	"COMPUTE_INSTANCE_MOUNT_BLOCK_VOLUMES_PRECHECK":                                  DrPlanStepTypeComputeInstanceMountBlockVolumesPrecheck,
	"COMPUTE_INSTANCE_UNMOUNT_BLOCK_VOLUMES_PRECHECK":                                DrPlanStepTypeComputeInstanceUnmountBlockVolumesPrecheck,
	"COMPUTE_CAPACITY_RESERVATION_START_DRILL_PRECHECK":                              DrPlanStepTypeComputeCapacityReservationStartDrillPrecheck,
	"COMPUTE_CAPACITY_AVAILABILITY_START_DRILL_PRECHECK":                             DrPlanStepTypeComputeCapacityAvailabilityStartDrillPrecheck,
	"AUTONOMOUS_DATABASE_CREATE_CLONE":                                               DrPlanStepTypeAutonomousDatabaseCreateClone,
	"AUTONOMOUS_DATABASE_DELETE_CLONE":                                               DrPlanStepTypeAutonomousDatabaseDeleteClone,
	"LOAD_BALANCER_UPDATE_PRIMARY_BACKEND_SET":                                       DrPlanStepTypeLoadBalancerUpdatePrimaryBackendSet,
	"LOAD_BALANCER_UPDATE_STANDBY_BACKEND_SET":                                       DrPlanStepTypeLoadBalancerUpdateStandbyBackendSet,
	"FILE_SYSTEM_SWITCHOVER":                                                         DrPlanStepTypeFileSystemSwitchover,
	"FILE_SYSTEM_FAILOVER":                                                           DrPlanStepTypeFileSystemFailover,
	"FILE_SYSTEM_REMOVE":                                                             DrPlanStepTypeFileSystemRemove,
	"FILE_SYSTEM_REVERSE":                                                            DrPlanStepTypeFileSystemReverse,
	"FILE_SYSTEM_TERMINATE":                                                          DrPlanStepTypeFileSystemTerminate,
	"FILE_SYSTEM_START_DRILL":                                                        DrPlanStepTypeFileSystemStartDrill,
	"FILE_SYSTEM_STOP_DRILL":                                                         DrPlanStepTypeFileSystemStopDrill,
	"COMPUTE_INSTANCE_START":                                                         DrPlanStepTypeComputeInstanceStart,
	"COMPUTE_INSTANCE_ATTACH_BLOCK_VOLUMES":                                          DrPlanStepTypeComputeInstanceAttachBlockVolumes,
	"COMPUTE_INSTANCE_DETACH_BLOCK_VOLUMES":                                          DrPlanStepTypeComputeInstanceDetachBlockVolumes,
	"FILE_SYSTEM_MOUNT":                                                              DrPlanStepTypeFileSystemMount,
	"FILE_SYSTEM_UNMOUNT":                                                            DrPlanStepTypeFileSystemUnmount,
	"COMPUTE_CAPACITY_RESERVATION_SWITCHOVER_PRECHECK":                               DrPlanStepTypeComputeCapacityReservationSwitchoverPrecheck,
	"COMPUTE_CAPACITY_RESERVATION_FAILOVER_PRECHECK":                                 DrPlanStepTypeComputeCapacityReservationFailoverPrecheck,
	"COMPUTE_CAPACITY_AVAILABILITY_SWITCHOVER_PRECHECK":                              DrPlanStepTypeComputeCapacityAvailabilitySwitchoverPrecheck,
	"COMPUTE_CAPACITY_AVAILABILITY_FAILOVER_PRECHECK":                                DrPlanStepTypeComputeCapacityAvailabilityFailoverPrecheck,
	"OKE_CLUSTER_SWITCHOVER_CLEANUP_PRIMARY_PRECHECK":                                DrPlanStepTypeOkeClusterSwitchoverCleanupPrimaryPrecheck,
	"OKE_CLUSTER_SWITCHOVER_RESTORE_STANDBY_PRECHECK":                                DrPlanStepTypeOkeClusterSwitchoverRestoreStandbyPrecheck,
	"OKE_CLUSTER_FAILOVER_RESTORE_STANDBY_PRECHECK":                                  DrPlanStepTypeOkeClusterFailoverRestoreStandbyPrecheck,
	"OKE_CLUSTER_STOP_DRILL_CLEANUP_STANDBY_PRECHECK":                                DrPlanStepTypeOkeClusterStopDrillCleanupStandbyPrecheck,
	"OKE_CLUSTER_START_DRILL_RESTORE_STANDBY_PRECHECK":                               DrPlanStepTypeOkeClusterStartDrillRestoreStandbyPrecheck,
	"OKE_CLUSTER_SWITCHOVER_CLEANUP_PRIMARY":                                         DrPlanStepTypeOkeClusterSwitchoverCleanupPrimary,
	"OKE_CLUSTER_SWITCHOVER_SCALE_UP_STANDBY":                                        DrPlanStepTypeOkeClusterSwitchoverScaleUpStandby,
	"OKE_CLUSTER_SWITCHOVER_RESTORE_STANDBY":                                         DrPlanStepTypeOkeClusterSwitchoverRestoreStandby,
	"OKE_CLUSTER_SWITCHOVER_REVERSE_BACKUP_STANDBY":                                  DrPlanStepTypeOkeClusterSwitchoverReverseBackupStandby,
	"OKE_CLUSTER_SWITCHOVER_SCALE_DOWN_PRIMARY":                                      DrPlanStepTypeOkeClusterSwitchoverScaleDownPrimary,
	"OKE_CLUSTER_FAILOVER_RESTORE_STANDBY":                                           DrPlanStepTypeOkeClusterFailoverRestoreStandby,
	"OKE_CLUSTER_FAILOVER_SCALE_UP_STANDBY":                                          DrPlanStepTypeOkeClusterFailoverScaleUpStandby,
	"OKE_CLUSTER_START_DRILL_RESTORE_STANDBY":                                        DrPlanStepTypeOkeClusterStartDrillRestoreStandby,
	"OKE_CLUSTER_START_DRILL_SCALE_UP_STANDBY":                                       DrPlanStepTypeOkeClusterStartDrillScaleUpStandby,
	"OKE_CLUSTER_STOP_DRILL_CLEANUP_STANDBY":                                         DrPlanStepTypeOkeClusterStopDrillCleanupStandby,
	"OKE_CLUSTER_STOP_DRILL_SCALE_DOWN_STANDBY":                                      DrPlanStepTypeOkeClusterStopDrillScaleDownStandby,
	"OBJECT_STORAGE_BUCKET_SWITCHOVER_DELETE_REPLICATION_PRIMARY_PRECHECK":           DrPlanStepTypeObjectStorageBucketSwitchoverDeleteReplicationPrimaryPrecheck,
	"OBJECT_STORAGE_BUCKET_SWITCHOVER_SETUP_REVERSE_REPLICATION_STANDBY_PRECHECK":    DrPlanStepTypeObjectStorageBucketSwitchoverSetupReverseReplicationStandbyPrecheck,
	"OBJECT_STORAGE_BUCKET_FAILOVER_DELETE_REPLICATION_STANDBY_PRECHECK":             DrPlanStepTypeObjectStorageBucketFailoverDeleteReplicationStandbyPrecheck,
	"OBJECT_STORAGE_BUCKET_FAILOVER_SETUP_REVERSE_REPLICATION_STANDBY_PRECHECK":      DrPlanStepTypeObjectStorageBucketFailoverSetupReverseReplicationStandbyPrecheck,
	"OBJECT_STORAGE_BUCKET_SWITCHOVER_DELETE_REPLICATION_PRIMARY":                    DrPlanStepTypeObjectStorageBucketSwitchoverDeleteReplicationPrimary,
	"OBJECT_STORAGE_BUCKET_SWITCHOVER_SETUP_REVERSE_REPLICATION_STANDBY":             DrPlanStepTypeObjectStorageBucketSwitchoverSetupReverseReplicationStandby,
	"OBJECT_STORAGE_BUCKET_FAILOVER_DELETE_REPLICATION_STANDBY":                      DrPlanStepTypeObjectStorageBucketFailoverDeleteReplicationStandby,
	"OBJECT_STORAGE_BUCKET_FAILOVER_SETUP_REVERSE_REPLICATION_STANDBY":               DrPlanStepTypeObjectStorageBucketFailoverSetupReverseReplicationStandby,
}

var mappingDrPlanStepTypeEnumLowerCase = map[string]DrPlanStepTypeEnum{
	"compute_instance_stop_precheck":                                                 DrPlanStepTypeComputeInstanceStopPrecheck,
	"compute_instance_launch_precheck":                                               DrPlanStepTypeComputeInstanceLaunchPrecheck,
	"compute_instance_terminate_precheck":                                            DrPlanStepTypeComputeInstanceTerminatePrecheck,
	"compute_instance_remove_precheck":                                               DrPlanStepTypeComputeInstanceRemovePrecheck,
	"volume_group_restore_switchover_precheck":                                       DrPlanStepTypeVolumeGroupRestoreSwitchoverPrecheck,
	"volume_group_restore_failover_precheck":                                         DrPlanStepTypeVolumeGroupRestoreFailoverPrecheck,
	"database_switchover_precheck":                                                   DrPlanStepTypeDatabaseSwitchoverPrecheck,
	"database_failover_precheck":                                                     DrPlanStepTypeDatabaseFailoverPrecheck,
	"autonomous_database_switchover_precheck":                                        DrPlanStepTypeAutonomousDatabaseSwitchoverPrecheck,
	"autonomous_database_failover_precheck":                                          DrPlanStepTypeAutonomousDatabaseFailoverPrecheck,
	"autonomous_container_database_switchover_precheck":                              DrPlanStepTypeAutonomousContainerDatabaseSwitchoverPrecheck,
	"autonomous_container_database_failover_precheck":                                DrPlanStepTypeAutonomousContainerDatabaseFailoverPrecheck,
	"autonomous_container_database_start_drill_convert_to_snapshot_standby_precheck": DrPlanStepTypeAutonomousContainerDatabaseStartDrillConvertToSnapshotStandbyPrecheck,
	"autonomous_container_database_stop_drill_convert_to_physical_standby_precheck":  DrPlanStepTypeAutonomousContainerDatabaseStopDrillConvertToPhysicalStandbyPrecheck,
	"autonomous_container_database_switchover":                                       DrPlanStepTypeAutonomousContainerDatabaseSwitchover,
	"autonomous_container_database_failover":                                         DrPlanStepTypeAutonomousContainerDatabaseFailover,
	"autonomous_container_database_start_drill_convert_to_snapshot_standby":          DrPlanStepTypeAutonomousContainerDatabaseStartDrillConvertToSnapshotStandby,
	"autonomous_container_database_stop_drill_convert_to_physical_standby":           DrPlanStepTypeAutonomousContainerDatabaseStopDrillConvertToPhysicalStandby,
	"autonomous_database_start_drill_create_clone_standby_precheck":                  DrPlanStepTypeAutonomousDatabaseStartDrillCreateCloneStandbyPrecheck,
	"autonomous_database_stop_drill_delete_clone_standby_precheck":                   DrPlanStepTypeAutonomousDatabaseStopDrillDeleteCloneStandbyPrecheck,
	"autonomous_database_start_drill_convert_to_snapshot_standby_precheck":           DrPlanStepTypeAutonomousDatabaseStartDrillConvertToSnapshotStandbyPrecheck,
	"autonomous_database_stop_drill_convert_to_physical_standby_precheck":            DrPlanStepTypeAutonomousDatabaseStopDrillConvertToPhysicalStandbyPrecheck,
	"autonomous_database_start_drill_create_clone_standby":                           DrPlanStepTypeAutonomousDatabaseStartDrillCreateCloneStandby,
	"autonomous_database_stop_drill_delete_clone_standby":                            DrPlanStepTypeAutonomousDatabaseStopDrillDeleteCloneStandby,
	"autonomous_database_start_drill_convert_to_snapshot_standby":                    DrPlanStepTypeAutonomousDatabaseStartDrillConvertToSnapshotStandby,
	"autonomous_database_stop_drill_convert_to_physical_standby":                     DrPlanStepTypeAutonomousDatabaseStopDrillConvertToPhysicalStandby,
	"user_defined_precheck":                                                          DrPlanStepTypeUserDefinedPrecheck,
	"compute_instance_launch":                                                        DrPlanStepTypeComputeInstanceLaunch,
	"compute_instance_stop":                                                          DrPlanStepTypeComputeInstanceStop,
	"compute_instance_terminate":                                                     DrPlanStepTypeComputeInstanceTerminate,
	"compute_instance_remove":                                                        DrPlanStepTypeComputeInstanceRemove,
	"database_switchover":                                                            DrPlanStepTypeDatabaseSwitchover,
	"database_failover":                                                              DrPlanStepTypeDatabaseFailover,
	"autonomous_database_switchover":                                                 DrPlanStepTypeAutonomousDatabaseSwitchover,
	"autonomous_database_failover":                                                   DrPlanStepTypeAutonomousDatabaseFailover,
	"volume_group_restore_switchover":                                                DrPlanStepTypeVolumeGroupRestoreSwitchover,
	"volume_group_restore_failover":                                                  DrPlanStepTypeVolumeGroupRestoreFailover,
	"volume_group_reverse":                                                           DrPlanStepTypeVolumeGroupReverse,
	"volume_group_delete":                                                            DrPlanStepTypeVolumeGroupDelete,
	"volume_group_remove":                                                            DrPlanStepTypeVolumeGroupRemove,
	"volume_group_terminate":                                                         DrPlanStepTypeVolumeGroupTerminate,
	"user_defined":                                                                   DrPlanStepTypeUserDefined,
	"volume_group_restore_start_drill_precheck":                                      DrPlanStepTypeVolumeGroupRestoreStartDrillPrecheck,
	"volume_group_remove_precheck":                                                   DrPlanStepTypeVolumeGroupRemovePrecheck,
	"volume_group_terminate_precheck":                                                DrPlanStepTypeVolumeGroupTerminatePrecheck,
	"volume_group_restore_start_drill":                                               DrPlanStepTypeVolumeGroupRestoreStartDrill,
	"autonomous_database_create_clone_precheck":                                      DrPlanStepTypeAutonomousDatabaseCreateClonePrecheck,
	"autonomous_database_delete_clone_precheck":                                      DrPlanStepTypeAutonomousDatabaseDeleteClonePrecheck,
	"load_balancer_update_primary_backend_set_precheck":                              DrPlanStepTypeLoadBalancerUpdatePrimaryBackendSetPrecheck,
	"load_balancer_update_standby_backend_set_precheck":                              DrPlanStepTypeLoadBalancerUpdateStandbyBackendSetPrecheck,
	"file_system_switchover_precheck":                                                DrPlanStepTypeFileSystemSwitchoverPrecheck,
	"file_system_failover_precheck":                                                  DrPlanStepTypeFileSystemFailoverPrecheck,
	"file_system_start_drill_precheck":                                               DrPlanStepTypeFileSystemStartDrillPrecheck,
	"file_system_stop_drill_precheck":                                                DrPlanStepTypeFileSystemStopDrillPrecheck,
	"file_system_remove_precheck":                                                    DrPlanStepTypeFileSystemRemovePrecheck,
	"file_system_terminate_precheck":                                                 DrPlanStepTypeFileSystemTerminatePrecheck,
	"file_system_mount_precheck":                                                     DrPlanStepTypeFileSystemMountPrecheck,
	"file_system_unmount_precheck":                                                   DrPlanStepTypeFileSystemUnmountPrecheck,
	"compute_instance_start_precheck":                                                DrPlanStepTypeComputeInstanceStartPrecheck,
	"compute_instance_attach_block_volumes_precheck":                                 DrPlanStepTypeComputeInstanceAttachBlockVolumesPrecheck,
	"compute_instance_detach_block_volumes_precheck":                                 DrPlanStepTypeComputeInstanceDetachBlockVolumesPrecheck,
	"compute_instance_mount_block_volumes_precheck":                                  DrPlanStepTypeComputeInstanceMountBlockVolumesPrecheck,
	"compute_instance_unmount_block_volumes_precheck":                                DrPlanStepTypeComputeInstanceUnmountBlockVolumesPrecheck,
	"compute_capacity_reservation_start_drill_precheck":                              DrPlanStepTypeComputeCapacityReservationStartDrillPrecheck,
	"compute_capacity_availability_start_drill_precheck":                             DrPlanStepTypeComputeCapacityAvailabilityStartDrillPrecheck,
	"autonomous_database_create_clone":                                               DrPlanStepTypeAutonomousDatabaseCreateClone,
	"autonomous_database_delete_clone":                                               DrPlanStepTypeAutonomousDatabaseDeleteClone,
	"load_balancer_update_primary_backend_set":                                       DrPlanStepTypeLoadBalancerUpdatePrimaryBackendSet,
	"load_balancer_update_standby_backend_set":                                       DrPlanStepTypeLoadBalancerUpdateStandbyBackendSet,
	"file_system_switchover":                                                         DrPlanStepTypeFileSystemSwitchover,
	"file_system_failover":                                                           DrPlanStepTypeFileSystemFailover,
	"file_system_remove":                                                             DrPlanStepTypeFileSystemRemove,
	"file_system_reverse":                                                            DrPlanStepTypeFileSystemReverse,
	"file_system_terminate":                                                          DrPlanStepTypeFileSystemTerminate,
	"file_system_start_drill":                                                        DrPlanStepTypeFileSystemStartDrill,
	"file_system_stop_drill":                                                         DrPlanStepTypeFileSystemStopDrill,
	"compute_instance_start":                                                         DrPlanStepTypeComputeInstanceStart,
	"compute_instance_attach_block_volumes":                                          DrPlanStepTypeComputeInstanceAttachBlockVolumes,
	"compute_instance_detach_block_volumes":                                          DrPlanStepTypeComputeInstanceDetachBlockVolumes,
	"file_system_mount":                                                              DrPlanStepTypeFileSystemMount,
	"file_system_unmount":                                                            DrPlanStepTypeFileSystemUnmount,
	"compute_capacity_reservation_switchover_precheck":                               DrPlanStepTypeComputeCapacityReservationSwitchoverPrecheck,
	"compute_capacity_reservation_failover_precheck":                                 DrPlanStepTypeComputeCapacityReservationFailoverPrecheck,
	"compute_capacity_availability_switchover_precheck":                              DrPlanStepTypeComputeCapacityAvailabilitySwitchoverPrecheck,
	"compute_capacity_availability_failover_precheck":                                DrPlanStepTypeComputeCapacityAvailabilityFailoverPrecheck,
	"oke_cluster_switchover_cleanup_primary_precheck":                                DrPlanStepTypeOkeClusterSwitchoverCleanupPrimaryPrecheck,
	"oke_cluster_switchover_restore_standby_precheck":                                DrPlanStepTypeOkeClusterSwitchoverRestoreStandbyPrecheck,
	"oke_cluster_failover_restore_standby_precheck":                                  DrPlanStepTypeOkeClusterFailoverRestoreStandbyPrecheck,
	"oke_cluster_stop_drill_cleanup_standby_precheck":                                DrPlanStepTypeOkeClusterStopDrillCleanupStandbyPrecheck,
	"oke_cluster_start_drill_restore_standby_precheck":                               DrPlanStepTypeOkeClusterStartDrillRestoreStandbyPrecheck,
	"oke_cluster_switchover_cleanup_primary":                                         DrPlanStepTypeOkeClusterSwitchoverCleanupPrimary,
	"oke_cluster_switchover_scale_up_standby":                                        DrPlanStepTypeOkeClusterSwitchoverScaleUpStandby,
	"oke_cluster_switchover_restore_standby":                                         DrPlanStepTypeOkeClusterSwitchoverRestoreStandby,
	"oke_cluster_switchover_reverse_backup_standby":                                  DrPlanStepTypeOkeClusterSwitchoverReverseBackupStandby,
	"oke_cluster_switchover_scale_down_primary":                                      DrPlanStepTypeOkeClusterSwitchoverScaleDownPrimary,
	"oke_cluster_failover_restore_standby":                                           DrPlanStepTypeOkeClusterFailoverRestoreStandby,
	"oke_cluster_failover_scale_up_standby":                                          DrPlanStepTypeOkeClusterFailoverScaleUpStandby,
	"oke_cluster_start_drill_restore_standby":                                        DrPlanStepTypeOkeClusterStartDrillRestoreStandby,
	"oke_cluster_start_drill_scale_up_standby":                                       DrPlanStepTypeOkeClusterStartDrillScaleUpStandby,
	"oke_cluster_stop_drill_cleanup_standby":                                         DrPlanStepTypeOkeClusterStopDrillCleanupStandby,
	"oke_cluster_stop_drill_scale_down_standby":                                      DrPlanStepTypeOkeClusterStopDrillScaleDownStandby,
	"object_storage_bucket_switchover_delete_replication_primary_precheck":           DrPlanStepTypeObjectStorageBucketSwitchoverDeleteReplicationPrimaryPrecheck,
	"object_storage_bucket_switchover_setup_reverse_replication_standby_precheck":    DrPlanStepTypeObjectStorageBucketSwitchoverSetupReverseReplicationStandbyPrecheck,
	"object_storage_bucket_failover_delete_replication_standby_precheck":             DrPlanStepTypeObjectStorageBucketFailoverDeleteReplicationStandbyPrecheck,
	"object_storage_bucket_failover_setup_reverse_replication_standby_precheck":      DrPlanStepTypeObjectStorageBucketFailoverSetupReverseReplicationStandbyPrecheck,
	"object_storage_bucket_switchover_delete_replication_primary":                    DrPlanStepTypeObjectStorageBucketSwitchoverDeleteReplicationPrimary,
	"object_storage_bucket_switchover_setup_reverse_replication_standby":             DrPlanStepTypeObjectStorageBucketSwitchoverSetupReverseReplicationStandby,
	"object_storage_bucket_failover_delete_replication_standby":                      DrPlanStepTypeObjectStorageBucketFailoverDeleteReplicationStandby,
	"object_storage_bucket_failover_setup_reverse_replication_standby":               DrPlanStepTypeObjectStorageBucketFailoverSetupReverseReplicationStandby,
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
		"AUTONOMOUS_CONTAINER_DATABASE_SWITCHOVER_PRECHECK",
		"AUTONOMOUS_CONTAINER_DATABASE_FAILOVER_PRECHECK",
		"AUTONOMOUS_CONTAINER_DATABASE_START_DRILL_CONVERT_TO_SNAPSHOT_STANDBY_PRECHECK",
		"AUTONOMOUS_CONTAINER_DATABASE_STOP_DRILL_CONVERT_TO_PHYSICAL_STANDBY_PRECHECK",
		"AUTONOMOUS_CONTAINER_DATABASE_SWITCHOVER",
		"AUTONOMOUS_CONTAINER_DATABASE_FAILOVER",
		"AUTONOMOUS_CONTAINER_DATABASE_START_DRILL_CONVERT_TO_SNAPSHOT_STANDBY",
		"AUTONOMOUS_CONTAINER_DATABASE_STOP_DRILL_CONVERT_TO_PHYSICAL_STANDBY",
		"AUTONOMOUS_DATABASE_START_DRILL_CREATE_CLONE_STANDBY_PRECHECK",
		"AUTONOMOUS_DATABASE_STOP_DRILL_DELETE_CLONE_STANDBY_PRECHECK",
		"AUTONOMOUS_DATABASE_START_DRILL_CONVERT_TO_SNAPSHOT_STANDBY_PRECHECK",
		"AUTONOMOUS_DATABASE_STOP_DRILL_CONVERT_TO_PHYSICAL_STANDBY_PRECHECK",
		"AUTONOMOUS_DATABASE_START_DRILL_CREATE_CLONE_STANDBY",
		"AUTONOMOUS_DATABASE_STOP_DRILL_DELETE_CLONE_STANDBY",
		"AUTONOMOUS_DATABASE_START_DRILL_CONVERT_TO_SNAPSHOT_STANDBY",
		"AUTONOMOUS_DATABASE_STOP_DRILL_CONVERT_TO_PHYSICAL_STANDBY",
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
		"VOLUME_GROUP_RESTORE_START_DRILL_PRECHECK",
		"VOLUME_GROUP_REMOVE_PRECHECK",
		"VOLUME_GROUP_TERMINATE_PRECHECK",
		"VOLUME_GROUP_RESTORE_START_DRILL",
		"AUTONOMOUS_DATABASE_CREATE_CLONE_PRECHECK",
		"AUTONOMOUS_DATABASE_DELETE_CLONE_PRECHECK",
		"LOAD_BALANCER_UPDATE_PRIMARY_BACKEND_SET_PRECHECK",
		"LOAD_BALANCER_UPDATE_STANDBY_BACKEND_SET_PRECHECK",
		"FILE_SYSTEM_SWITCHOVER_PRECHECK",
		"FILE_SYSTEM_FAILOVER_PRECHECK",
		"FILE_SYSTEM_START_DRILL_PRECHECK",
		"FILE_SYSTEM_STOP_DRILL_PRECHECK",
		"FILE_SYSTEM_REMOVE_PRECHECK",
		"FILE_SYSTEM_TERMINATE_PRECHECK",
		"FILE_SYSTEM_MOUNT_PRECHECK",
		"FILE_SYSTEM_UNMOUNT_PRECHECK",
		"COMPUTE_INSTANCE_START_PRECHECK",
		"COMPUTE_INSTANCE_ATTACH_BLOCK_VOLUMES_PRECHECK",
		"COMPUTE_INSTANCE_DETACH_BLOCK_VOLUMES_PRECHECK",
		"COMPUTE_INSTANCE_MOUNT_BLOCK_VOLUMES_PRECHECK",
		"COMPUTE_INSTANCE_UNMOUNT_BLOCK_VOLUMES_PRECHECK",
		"COMPUTE_CAPACITY_RESERVATION_START_DRILL_PRECHECK",
		"COMPUTE_CAPACITY_AVAILABILITY_START_DRILL_PRECHECK",
		"AUTONOMOUS_DATABASE_CREATE_CLONE",
		"AUTONOMOUS_DATABASE_DELETE_CLONE",
		"LOAD_BALANCER_UPDATE_PRIMARY_BACKEND_SET",
		"LOAD_BALANCER_UPDATE_STANDBY_BACKEND_SET",
		"FILE_SYSTEM_SWITCHOVER",
		"FILE_SYSTEM_FAILOVER",
		"FILE_SYSTEM_REMOVE",
		"FILE_SYSTEM_REVERSE",
		"FILE_SYSTEM_TERMINATE",
		"FILE_SYSTEM_START_DRILL",
		"FILE_SYSTEM_STOP_DRILL",
		"COMPUTE_INSTANCE_START",
		"COMPUTE_INSTANCE_ATTACH_BLOCK_VOLUMES",
		"COMPUTE_INSTANCE_DETACH_BLOCK_VOLUMES",
		"FILE_SYSTEM_MOUNT",
		"FILE_SYSTEM_UNMOUNT",
		"COMPUTE_CAPACITY_RESERVATION_SWITCHOVER_PRECHECK",
		"COMPUTE_CAPACITY_RESERVATION_FAILOVER_PRECHECK",
		"COMPUTE_CAPACITY_AVAILABILITY_SWITCHOVER_PRECHECK",
		"COMPUTE_CAPACITY_AVAILABILITY_FAILOVER_PRECHECK",
		"OKE_CLUSTER_SWITCHOVER_CLEANUP_PRIMARY_PRECHECK",
		"OKE_CLUSTER_SWITCHOVER_RESTORE_STANDBY_PRECHECK",
		"OKE_CLUSTER_FAILOVER_RESTORE_STANDBY_PRECHECK",
		"OKE_CLUSTER_STOP_DRILL_CLEANUP_STANDBY_PRECHECK",
		"OKE_CLUSTER_START_DRILL_RESTORE_STANDBY_PRECHECK",
		"OKE_CLUSTER_SWITCHOVER_CLEANUP_PRIMARY",
		"OKE_CLUSTER_SWITCHOVER_SCALE_UP_STANDBY",
		"OKE_CLUSTER_SWITCHOVER_RESTORE_STANDBY",
		"OKE_CLUSTER_SWITCHOVER_REVERSE_BACKUP_STANDBY",
		"OKE_CLUSTER_SWITCHOVER_SCALE_DOWN_PRIMARY",
		"OKE_CLUSTER_FAILOVER_RESTORE_STANDBY",
		"OKE_CLUSTER_FAILOVER_SCALE_UP_STANDBY",
		"OKE_CLUSTER_START_DRILL_RESTORE_STANDBY",
		"OKE_CLUSTER_START_DRILL_SCALE_UP_STANDBY",
		"OKE_CLUSTER_STOP_DRILL_CLEANUP_STANDBY",
		"OKE_CLUSTER_STOP_DRILL_SCALE_DOWN_STANDBY",
		"OBJECT_STORAGE_BUCKET_SWITCHOVER_DELETE_REPLICATION_PRIMARY_PRECHECK",
		"OBJECT_STORAGE_BUCKET_SWITCHOVER_SETUP_REVERSE_REPLICATION_STANDBY_PRECHECK",
		"OBJECT_STORAGE_BUCKET_FAILOVER_DELETE_REPLICATION_STANDBY_PRECHECK",
		"OBJECT_STORAGE_BUCKET_FAILOVER_SETUP_REVERSE_REPLICATION_STANDBY_PRECHECK",
		"OBJECT_STORAGE_BUCKET_SWITCHOVER_DELETE_REPLICATION_PRIMARY",
		"OBJECT_STORAGE_BUCKET_SWITCHOVER_SETUP_REVERSE_REPLICATION_STANDBY",
		"OBJECT_STORAGE_BUCKET_FAILOVER_DELETE_REPLICATION_STANDBY",
		"OBJECT_STORAGE_BUCKET_FAILOVER_SETUP_REVERSE_REPLICATION_STANDBY",
	}
}

// GetMappingDrPlanStepTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDrPlanStepTypeEnum(val string) (DrPlanStepTypeEnum, bool) {
	enum, ok := mappingDrPlanStepTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
