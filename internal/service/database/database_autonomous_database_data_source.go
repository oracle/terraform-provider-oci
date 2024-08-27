// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"
)

func DatabaseAutonomousDatabaseDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["autonomous_database_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DatabaseAutonomousDatabaseResource(), fieldMap, readSingularDatabaseAutonomousDatabase)
}

func readSingularDatabaseAutonomousDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousDatabaseDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseAutonomousDatabaseDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.GetAutonomousDatabaseResponse
}

func (s *DatabaseAutonomousDatabaseDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseAutonomousDatabaseDataSourceCrud) Get() error {
	request := oci_database.GetAutonomousDatabaseRequest{}

	if autonomousDatabaseId, ok := s.D.GetOkExists("autonomous_database_id"); ok {
		tmp := autonomousDatabaseId.(string)
		request.AutonomousDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.GetAutonomousDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseAutonomousDatabaseDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.ActualUsedDataStorageSizeInTBs != nil {
		s.D.Set("actual_used_data_storage_size_in_tbs", *s.Res.ActualUsedDataStorageSizeInTBs)
	}

	if s.Res.AllocatedStorageSizeInTBs != nil {
		s.D.Set("allocated_storage_size_in_tbs", *s.Res.AllocatedStorageSizeInTBs)
	}

	if s.Res.ArePrimaryWhitelistedIpsUsed != nil {
		s.D.Set("are_primary_whitelisted_ips_used", *s.Res.ArePrimaryWhitelistedIpsUsed)
	}

	if s.Res.ApexDetails != nil {
		s.D.Set("apex_details", []interface{}{AutonomousDatabaseApexToMap(s.Res.ApexDetails)})
	} else {
		s.D.Set("apex_details", nil)
	}

	if s.Res.AutoRefreshFrequencyInSeconds != nil {
		s.D.Set("auto_refresh_frequency_in_seconds", *s.Res.AutoRefreshFrequencyInSeconds)
	}

	if s.Res.AutoRefreshPointLagInSeconds != nil {
		s.D.Set("auto_refresh_point_lag_in_seconds", *s.Res.AutoRefreshPointLagInSeconds)
	}

	if s.Res.AutonomousContainerDatabaseId != nil {
		s.D.Set("autonomous_container_database_id", *s.Res.AutonomousContainerDatabaseId)
	}

	s.D.Set("autonomous_maintenance_schedule_type", s.Res.AutonomousMaintenanceScheduleType)

	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	s.D.Set("available_upgrade_versions", s.Res.AvailableUpgradeVersions)

	if s.Res.BackupConfig != nil {
		s.D.Set("backup_config", []interface{}{AutonomousDatabaseBackupConfigToMap(s.Res.BackupConfig)})
	} else {
		s.D.Set("backup_config", nil)
	}

	if s.Res.BackupRetentionPeriodInDays != nil {
		s.D.Set("backup_retention_period_in_days", *s.Res.BackupRetentionPeriodInDays)
	}

	if s.Res.ByolComputeCountLimit != nil {
		s.D.Set("byol_compute_count_limit", *s.Res.ByolComputeCountLimit)
	}

	if s.Res.CharacterSet != nil {
		s.D.Set("character_set", *s.Res.CharacterSet)
	}

	if s.Res.ClusterPlacementGroupId != nil {
		s.D.Set("cluster_placement_group_id", *s.Res.ClusterPlacementGroupId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ComputeCount != nil {
		s.D.Set("compute_count", *s.Res.ComputeCount)
	}

	s.D.Set("compute_model", s.Res.ComputeModel)

	if s.Res.ConnectionStrings != nil {
		s.D.Set("connection_strings", []interface{}{AutonomousDatabaseConnectionStringsToMap(s.Res.ConnectionStrings)})
	} else {
		s.D.Set("connection_strings", nil)
	}

	if s.Res.ConnectionUrls != nil {
		s.D.Set("connection_urls", []interface{}{AutonomousDatabaseConnectionUrlsToMap(s.Res.ConnectionUrls)})
	} else {
		s.D.Set("connection_urls", nil)
	}

	if s.Res.CpuCoreCount != nil {
		s.D.Set("cpu_core_count", *s.Res.CpuCoreCount)
	}

	customerContacts := []interface{}{}
	for _, item := range s.Res.CustomerContacts {
		customerContacts = append(customerContacts, CustomerContactToMap(item))
	}
	s.D.Set("customer_contacts", customerContacts)

	s.D.Set("data_safe_status", s.Res.DataSafeStatus)

	if s.Res.DataStorageSizeInGBs != nil {
		s.D.Set("data_storage_size_in_gb", *s.Res.DataStorageSizeInGBs)
	}

	if s.Res.DataStorageSizeInTBs != nil {
		s.D.Set("data_storage_size_in_tbs", *s.Res.DataStorageSizeInTBs)
	}

	s.D.Set("database_edition", s.Res.DatabaseEdition)

	s.D.Set("database_management_status", s.Res.DatabaseManagementStatus)

	s.D.Set("dataguard_region_type", s.Res.DataguardRegionType)

	if s.Res.DbName != nil {
		s.D.Set("db_name", *s.Res.DbName)
	}

	dbToolsDetails := []interface{}{}
	for _, item := range s.Res.DbToolsDetails {
		dbToolsDetails = append(dbToolsDetails, DatabaseToolToMap(item))
	}
	s.D.Set("db_tools_details", dbToolsDetails)

	if s.Res.DbVersion != nil {
		s.D.Set("db_version", *s.Res.DbVersion)
	}

	s.D.Set("db_workload", s.Res.DbWorkload)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("disaster_recovery_region_type", s.Res.DisasterRecoveryRegionType)

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.FailedDataRecoveryInSeconds != nil {
		s.D.Set("failed_data_recovery_in_seconds", *s.Res.FailedDataRecoveryInSeconds)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.InMemoryAreaInGBs != nil {
		s.D.Set("in_memory_area_in_gbs", *s.Res.InMemoryAreaInGBs)
	}

	if s.Res.InMemoryPercentage != nil {
		s.D.Set("in_memory_percentage", *s.Res.InMemoryPercentage)
	}

	if s.Res.InMemoryAreaInGBs != nil {
		s.D.Set("in_memory_area_in_gbs", *s.Res.InMemoryAreaInGBs)
	}

	if s.Res.InMemoryPercentage != nil {
		s.D.Set("in_memory_percentage", *s.Res.InMemoryPercentage)
	}

	s.D.Set("infrastructure_type", s.Res.InfrastructureType)

	if s.Res.IsAccessControlEnabled != nil {
		s.D.Set("is_access_control_enabled", *s.Res.IsAccessControlEnabled)
	}

	if s.Res.IsAutoScalingEnabled != nil {
		s.D.Set("is_auto_scaling_enabled", *s.Res.IsAutoScalingEnabled)
	}

	if s.Res.IsAutoScalingForStorageEnabled != nil {
		s.D.Set("is_auto_scaling_for_storage_enabled", *s.Res.IsAutoScalingForStorageEnabled)
	}

	if s.Res.IsDataGuardEnabled != nil {
		s.D.Set("is_data_guard_enabled", *s.Res.IsDataGuardEnabled)
	}

	if s.Res.IsDedicated != nil {
		s.D.Set("is_dedicated", *s.Res.IsDedicated)
	}

	if s.Res.IsDevTier != nil {
		s.D.Set("is_dev_tier", *s.Res.IsDevTier)
	}

	if s.Res.IsFreeTier != nil {
		s.D.Set("is_free_tier", *s.Res.IsFreeTier)
	}

	if s.Res.IsLocalDataGuardEnabled != nil {
		s.D.Set("is_local_data_guard_enabled", *s.Res.IsLocalDataGuardEnabled)
	}

	if s.Res.IsMtlsConnectionRequired != nil {
		s.D.Set("is_mtls_connection_required", *s.Res.IsMtlsConnectionRequired)
	}

	if s.Res.IsPreview != nil {
		s.D.Set("is_preview", *s.Res.IsPreview)
	}

	if s.Res.IsReconnectCloneEnabled != nil {
		s.D.Set("is_reconnect_clone_enabled", *s.Res.IsReconnectCloneEnabled)
	}

	if s.Res.IsRefreshableClone != nil {
		s.D.Set("is_refreshable_clone", *s.Res.IsRefreshableClone)
	}

	if s.Res.IsRemoteDataGuardEnabled != nil {
		s.D.Set("is_remote_data_guard_enabled", *s.Res.IsRemoteDataGuardEnabled)
	}

	keyHistoryEntry := []interface{}{}
	for _, item := range s.Res.KeyHistoryEntry {
		keyHistoryEntry = append(keyHistoryEntry, AutonomousDatabaseKeyHistoryEntryToMap(item))
	}
	s.D.Set("key_history_entry", keyHistoryEntry)

	if s.Res.KeyStoreId != nil {
		s.D.Set("key_store_id", *s.Res.KeyStoreId)
	}

	if s.Res.KeyStoreWalletName != nil {
		s.D.Set("key_store_wallet_name", *s.Res.KeyStoreWalletName)
	}

	if s.Res.KmsKeyId != nil {
		s.D.Set("kms_key_id", *s.Res.KmsKeyId)
	}

	if s.Res.KmsKeyLifecycleDetails != nil {
		s.D.Set("kms_key_lifecycle_details", *s.Res.KmsKeyLifecycleDetails)
	}

	if s.Res.KmsKeyVersionId != nil {
		s.D.Set("kms_key_version_id", *s.Res.KmsKeyVersionId)
	}

	s.D.Set("license_model", s.Res.LicenseModel)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.LocalAdgAutoFailoverMaxDataLossLimit != nil {
		s.D.Set("local_adg_auto_failover_max_data_loss_limit", *s.Res.LocalAdgAutoFailoverMaxDataLossLimit)
	}

	s.D.Set("local_disaster_recovery_type", s.Res.LocalDisasterRecoveryType)

	if s.Res.LocalStandbyDb != nil {
		s.D.Set("local_standby_db", []interface{}{AutonomousDatabaseStandbySummaryToMap(s.Res.LocalStandbyDb)})
	} else {
		s.D.Set("local_standby_db", nil)
	}

	if s.Res.LongTermBackupSchedule != nil {
		s.D.Set("long_term_backup_schedule", []interface{}{LongTermBackUpScheduleDetailsToMap(s.Res.LongTermBackupSchedule)})
	} else {
		s.D.Set("long_term_backup_schedule", nil)
	}

	if s.Res.MemoryPerOracleComputeUnitInGBs != nil {
		s.D.Set("memory_per_oracle_compute_unit_in_gbs", *s.Res.MemoryPerOracleComputeUnitInGBs)
	}

	if s.Res.NcharacterSet != nil {
		s.D.Set("ncharacter_set", *s.Res.NcharacterSet)
	}

	s.D.Set("net_services_architecture", s.Res.NetServicesArchitecture)

	if s.Res.NextLongTermBackupTimeStamp != nil {
		s.D.Set("next_long_term_backup_time_stamp", s.Res.NextLongTermBackupTimeStamp.String())
	}

	s.D.Set("nsg_ids", s.Res.NsgIds)

	if s.Res.OcpuCount != nil {
		s.D.Set("ocpu_count", *s.Res.OcpuCount)
	}

	s.D.Set("open_mode", s.Res.OpenMode)

	s.D.Set("operations_insights_status", s.Res.OperationsInsightsStatus)

	s.D.Set("peer_db_ids", s.Res.PeerDbIds)

	s.D.Set("permission_level", s.Res.PermissionLevel)

	if s.Res.PrivateEndpoint != nil {
		s.D.Set("private_endpoint", *s.Res.PrivateEndpoint)
	}

	if s.Res.PrivateEndpointIp != nil {
		s.D.Set("private_endpoint_ip", *s.Res.PrivateEndpointIp)
	}

	if s.Res.PrivateEndpointLabel != nil {
		s.D.Set("private_endpoint_label", *s.Res.PrivateEndpointLabel)
	}

	s.D.Set("provisionable_cpus", s.Res.ProvisionableCpus)

	if s.Res.PublicConnectionUrls != nil {
		s.D.Set("public_connection_urls", []interface{}{AutonomousDatabaseConnectionUrlsToMap(s.Res.PublicConnectionUrls)})
	} else {
		s.D.Set("public_connection_urls", nil)
	}

	if s.Res.PublicEndpoint != nil {
		s.D.Set("public_endpoint", *s.Res.PublicEndpoint)
	}

	s.D.Set("refreshable_mode", s.Res.RefreshableMode)

	s.D.Set("refreshable_status", s.Res.RefreshableStatus)

	if s.Res.RemoteDisasterRecoveryConfiguration != nil {
		s.D.Set("remote_disaster_recovery_configuration", []interface{}{DisasterRecoveryConfigurationToMap(s.Res.RemoteDisasterRecoveryConfiguration)})
	} else {
		s.D.Set("remote_disaster_recovery_configuration", nil)
	}

	if s.Res.ResourcePoolLeaderId != nil {
		s.D.Set("resource_pool_leader_id", *s.Res.ResourcePoolLeaderId)
	}

	if s.Res.ResourcePoolSummary != nil {
		s.D.Set("resource_pool_summary", []interface{}{ResourcePoolSummaryToMap(s.Res.ResourcePoolSummary)})
	} else {
		s.D.Set("resource_pool_summary", nil)
	}

	s.D.Set("role", s.Res.Role)

	scheduledOperations := []interface{}{}
	for _, item := range s.Res.ScheduledOperations {
		scheduledOperations = append(scheduledOperations, ScheduledOperationDetailsToMap(item))
	}
	s.D.Set("scheduled_operations", scheduledOperations)

	if s.Res.ServiceConsoleUrl != nil {
		s.D.Set("service_console_url", *s.Res.ServiceConsoleUrl)
	}

	if s.Res.SourceId != nil {
		s.D.Set("source_id", *s.Res.SourceId)
	}

	if s.Res.StandbyDb != nil {
		s.D.Set("standby_db", []interface{}{AutonomousDatabaseStandbySummaryToMap(s.Res.StandbyDb)})
	} else {
		s.D.Set("standby_db", nil)
	}

	s.D.Set("standby_whitelisted_ips", s.Res.StandbyWhitelistedIps)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", *s.Res.SubnetId)
	}

	if s.Res.SubscriptionId != nil {
		s.D.Set("subscription_id", *s.Res.SubscriptionId)
	}

	s.D.Set("supported_regions_to_clone_to", s.Res.SupportedRegionsToCloneTo)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeDataGuardRoleChanged != nil {
		s.D.Set("time_data_guard_role_changed", s.Res.TimeDataGuardRoleChanged.String())
	}

	if s.Res.TimeDeletionOfFreeAutonomousDatabase != nil {
		s.D.Set("time_deletion_of_free_autonomous_database", s.Res.TimeDeletionOfFreeAutonomousDatabase.String())
	}

	if s.Res.TimeDisasterRecoveryRoleChanged != nil {
		s.D.Set("time_disaster_recovery_role_changed", s.Res.TimeDisasterRecoveryRoleChanged.String())
	}

	if s.Res.TimeLocalDataGuardEnabled != nil {
		s.D.Set("time_local_data_guard_enabled", s.Res.TimeLocalDataGuardEnabled.String())
	}

	if s.Res.TimeMaintenanceBegin != nil {
		s.D.Set("time_maintenance_begin", s.Res.TimeMaintenanceBegin.String())
	}

	if s.Res.TimeMaintenanceEnd != nil {
		s.D.Set("time_maintenance_end", s.Res.TimeMaintenanceEnd.String())
	}

	if s.Res.TimeOfAutoRefreshStart != nil {
		s.D.Set("time_of_auto_refresh_start", s.Res.TimeOfAutoRefreshStart.Format(time.RFC3339Nano))
	}

	if s.Res.TimeOfJoiningResourcePool != nil {
		s.D.Set("time_of_joining_resource_pool", s.Res.TimeOfJoiningResourcePool.String())
	}

	if s.Res.TimeOfLastFailover != nil {
		s.D.Set("time_of_last_failover", s.Res.TimeOfLastFailover.String())
	}

	if s.Res.TimeOfLastRefresh != nil {
		s.D.Set("time_of_last_refresh", s.Res.TimeOfLastRefresh.String())
	}

	if s.Res.TimeOfLastRefreshPoint != nil {
		s.D.Set("time_of_last_refresh_point", s.Res.TimeOfLastRefreshPoint.String())
	}

	if s.Res.TimeOfLastSwitchover != nil {
		s.D.Set("time_of_last_switchover", s.Res.TimeOfLastSwitchover.String())
	}

	if s.Res.TimeOfNextRefresh != nil {
		s.D.Set("time_of_next_refresh", s.Res.TimeOfNextRefresh.String())
	}

	if s.Res.TimeReclamationOfFreeAutonomousDatabase != nil {
		s.D.Set("time_reclamation_of_free_autonomous_database", s.Res.TimeReclamationOfFreeAutonomousDatabase.String())
	}

	if s.Res.TimeUntilReconnectCloneEnabled != nil {
		s.D.Set("time_until_reconnect_clone_enabled", s.Res.TimeUntilReconnectCloneEnabled.String())
	}

	if s.Res.TotalBackupStorageSizeInGBs != nil {
		s.D.Set("total_backup_storage_size_in_gbs", *s.Res.TotalBackupStorageSizeInGBs)
	}

	if s.Res.UsedDataStorageSizeInGBs != nil {
		s.D.Set("used_data_storage_size_in_gbs", *s.Res.UsedDataStorageSizeInGBs)
	}

	if s.Res.UsedDataStorageSizeInTBs != nil {
		s.D.Set("used_data_storage_size_in_tbs", *s.Res.UsedDataStorageSizeInTBs)
	}

	if s.Res.VaultId != nil {
		s.D.Set("vault_id", *s.Res.VaultId)
	}

	s.D.Set("whitelisted_ips", s.Res.WhitelistedIps)

	return nil
}
