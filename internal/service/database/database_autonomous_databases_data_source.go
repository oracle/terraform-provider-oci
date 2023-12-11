// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"
)

func DatabaseAutonomousDatabasesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseAutonomousDatabases,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"autonomous_container_database_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"db_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"db_workload": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"infrastructure_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_data_guard_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"is_free_tier": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"is_refreshable_clone": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"is_resource_pool_leader": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"resource_pool_leader_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"autonomous_databases": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(DatabaseAutonomousDatabaseResource()),
			},
		},
	}
}

func readDatabaseAutonomousDatabases(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousDatabasesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseAutonomousDatabasesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListAutonomousDatabasesResponse
}

func (s *DatabaseAutonomousDatabasesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseAutonomousDatabasesDataSourceCrud) Get() error {
	request := oci_database.ListAutonomousDatabasesRequest{}

	if autonomousContainerDatabaseId, ok := s.D.GetOkExists("autonomous_container_database_id"); ok {
		tmp := autonomousContainerDatabaseId.(string)
		request.AutonomousContainerDatabaseId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if dbVersion, ok := s.D.GetOkExists("db_version"); ok {
		tmp := dbVersion.(string)
		request.DbVersion = &tmp
	}

	if dbWorkload, ok := s.D.GetOkExists("db_workload"); ok {
		request.DbWorkload = oci_database.AutonomousDatabaseSummaryDbWorkloadEnum(dbWorkload.(string))
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if infrastructureType, ok := s.D.GetOkExists("infrastructure_type"); ok {
		request.InfrastructureType = oci_database.AutonomousDatabaseSummaryInfrastructureTypeEnum(infrastructureType.(string))
	}

	if isDataGuardEnabled, ok := s.D.GetOkExists("is_data_guard_enabled"); ok {
		tmp := isDataGuardEnabled.(bool)
		request.IsDataGuardEnabled = &tmp
	}

	if isFreeTier, ok := s.D.GetOkExists("is_free_tier"); ok {
		tmp := isFreeTier.(bool)
		request.IsFreeTier = &tmp
	}

	if isRefreshableClone, ok := s.D.GetOkExists("is_refreshable_clone"); ok {
		tmp := isRefreshableClone.(bool)
		request.IsRefreshableClone = &tmp
	}

	if isResourcePoolLeader, ok := s.D.GetOkExists("is_resource_pool_leader"); ok {
		tmp := isResourcePoolLeader.(bool)
		request.IsResourcePoolLeader = &tmp
	}

	if resourcePoolLeaderId, ok := s.D.GetOkExists("resource_pool_leader_id"); ok {
		tmp := resourcePoolLeaderId.(string)
		request.ResourcePoolLeaderId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_database.AutonomousDatabaseSummaryLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListAutonomousDatabases(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAutonomousDatabases(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseAutonomousDatabasesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseAutonomousDatabasesDataSource-", DatabaseAutonomousDatabasesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		autonomousDatabase := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.ActualUsedDataStorageSizeInTBs != nil {
			autonomousDatabase["actual_used_data_storage_size_in_tbs"] = *r.ActualUsedDataStorageSizeInTBs
		}

		if r.AllocatedStorageSizeInTBs != nil {
			autonomousDatabase["allocated_storage_size_in_tbs"] = *r.AllocatedStorageSizeInTBs
		}

		if r.ApexDetails != nil {
			autonomousDatabase["apex_details"] = []interface{}{AutonomousDatabaseApexToMap(r.ApexDetails)}
		} else {
			autonomousDatabase["apex_details"] = nil
		}

		if r.ArePrimaryWhitelistedIpsUsed != nil {
			autonomousDatabase["are_primary_whitelisted_ips_used"] = *r.ArePrimaryWhitelistedIpsUsed
		}
		if r.AutonomousContainerDatabaseId != nil {
			autonomousDatabase["autonomous_container_database_id"] = *r.AutonomousContainerDatabaseId
		}

		autonomousDatabase["autonomous_maintenance_schedule_type"] = r.AutonomousMaintenanceScheduleType

		autonomousDatabase["available_upgrade_versions"] = r.AvailableUpgradeVersions

		if r.BackupConfig != nil {
			autonomousDatabase["backup_config"] = []interface{}{AutonomousDatabaseBackupConfigToMap(r.BackupConfig)}
		} else {
			autonomousDatabase["backup_config"] = nil
		}

		if r.BackupRetentionPeriodInDays != nil {
			autonomousDatabase["backup_retention_period_in_days"] = *r.BackupRetentionPeriodInDays
		}

		if r.CharacterSet != nil {
			autonomousDatabase["character_set"] = *r.CharacterSet
		}

		if r.ComputeCount != nil {
			autonomousDatabase["compute_count"] = *r.ComputeCount
		}

		autonomousDatabase["compute_model"] = r.ComputeModel

		if r.ConnectionStrings != nil {
			autonomousDatabase["connection_strings"] = []interface{}{AutonomousDatabaseConnectionStringsToMap(r.ConnectionStrings)}
		} else {
			autonomousDatabase["connection_strings"] = nil
		}

		if r.ConnectionUrls != nil {
			autonomousDatabase["connection_urls"] = []interface{}{AutonomousDatabaseConnectionUrlsToMap(r.ConnectionUrls)}
		} else {
			autonomousDatabase["connection_urls"] = nil
		}

		if r.CpuCoreCount != nil {
			autonomousDatabase["cpu_core_count"] = *r.CpuCoreCount
		}

		customerContacts := []interface{}{}
		for _, item := range r.CustomerContacts {
			customerContacts = append(customerContacts, CustomerContactToMap(item))
		}
		autonomousDatabase["customer_contacts"] = customerContacts

		autonomousDatabase["data_safe_status"] = r.DataSafeStatus

		if r.DataStorageSizeInGBs != nil {
			autonomousDatabase["data_storage_size_in_gb"] = *r.DataStorageSizeInGBs
		}

		if r.DataStorageSizeInTBs != nil {
			autonomousDatabase["data_storage_size_in_tbs"] = *r.DataStorageSizeInTBs
		}

		autonomousDatabase["database_edition"] = r.DatabaseEdition

		autonomousDatabase["database_management_status"] = r.DatabaseManagementStatus

		autonomousDatabase["dataguard_region_type"] = r.DataguardRegionType

		if r.DbName != nil {
			autonomousDatabase["db_name"] = *r.DbName
		}

		dbToolsDetails := []interface{}{}
		for _, item := range r.DbToolsDetails {
			dbToolsDetails = append(dbToolsDetails, DatabaseToolToMap(item))
		}
		autonomousDatabase["db_tools_details"] = dbToolsDetails

		if r.DbVersion != nil {
			autonomousDatabase["db_version"] = *r.DbVersion
		}

		autonomousDatabase["db_workload"] = r.DbWorkload

		if r.DefinedTags != nil {
			autonomousDatabase["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		autonomousDatabase["disaster_recovery_region_type"] = r.DisasterRecoveryRegionType

		if r.DisplayName != nil {
			autonomousDatabase["display_name"] = *r.DisplayName
		}

		if r.FailedDataRecoveryInSeconds != nil {
			autonomousDatabase["failed_data_recovery_in_seconds"] = *r.FailedDataRecoveryInSeconds
		}

		autonomousDatabase["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			autonomousDatabase["id"] = *r.Id
		}

		if r.InMemoryAreaInGBs != nil {
			autonomousDatabase["in_memory_area_in_gbs"] = *r.InMemoryAreaInGBs
		}

		if r.InMemoryPercentage != nil {
			autonomousDatabase["in_memory_percentage"] = *r.InMemoryPercentage
		}

		autonomousDatabase["infrastructure_type"] = r.InfrastructureType

		if r.IsAccessControlEnabled != nil {
			autonomousDatabase["is_access_control_enabled"] = *r.IsAccessControlEnabled
		}

		if r.IsAutoScalingEnabled != nil {
			autonomousDatabase["is_auto_scaling_enabled"] = *r.IsAutoScalingEnabled
		}

		if r.IsAutoScalingForStorageEnabled != nil {
			autonomousDatabase["is_auto_scaling_for_storage_enabled"] = *r.IsAutoScalingForStorageEnabled
		}

		if r.IsDataGuardEnabled != nil {
			autonomousDatabase["is_data_guard_enabled"] = *r.IsDataGuardEnabled
		}

		if r.IsDedicated != nil {
			autonomousDatabase["is_dedicated"] = *r.IsDedicated
		}

		if r.IsDevTier != nil {
			autonomousDatabase["is_dev_tier"] = *r.IsDevTier
		}

		if r.IsFreeTier != nil {
			autonomousDatabase["is_free_tier"] = *r.IsFreeTier
		}

		if r.IsLocalDataGuardEnabled != nil {
			autonomousDatabase["is_local_data_guard_enabled"] = *r.IsLocalDataGuardEnabled
		}

		if r.IsMtlsConnectionRequired != nil {
			autonomousDatabase["is_mtls_connection_required"] = *r.IsMtlsConnectionRequired
		}

		if r.IsPreview != nil {
			autonomousDatabase["is_preview"] = *r.IsPreview
		}

		if r.IsReconnectCloneEnabled != nil {
			autonomousDatabase["is_reconnect_clone_enabled"] = *r.IsReconnectCloneEnabled
		}

		if r.IsRefreshableClone != nil {
			autonomousDatabase["is_refreshable_clone"] = *r.IsRefreshableClone
		}

		if r.IsRemoteDataGuardEnabled != nil {
			autonomousDatabase["is_remote_data_guard_enabled"] = *r.IsRemoteDataGuardEnabled
		}

		keyHistoryEntry := []interface{}{}
		for _, item := range r.KeyHistoryEntry {
			keyHistoryEntry = append(keyHistoryEntry, AutonomousDatabaseKeyHistoryEntryToMap(item))
		}
		autonomousDatabase["key_history_entry"] = keyHistoryEntry

		if r.KeyStoreId != nil {
			autonomousDatabase["key_store_id"] = *r.KeyStoreId
		}

		if r.KeyStoreWalletName != nil {
			autonomousDatabase["key_store_wallet_name"] = *r.KeyStoreWalletName
		}

		if r.KmsKeyId != nil {
			autonomousDatabase["kms_key_id"] = *r.KmsKeyId
		}

		if r.KmsKeyLifecycleDetails != nil {
			autonomousDatabase["kms_key_lifecycle_details"] = *r.KmsKeyLifecycleDetails
		}

		if r.KmsKeyVersionId != nil {
			autonomousDatabase["kms_key_version_id"] = *r.KmsKeyVersionId
		}

		autonomousDatabase["license_model"] = r.LicenseModel

		if r.LifecycleDetails != nil {
			autonomousDatabase["lifecycle_details"] = *r.LifecycleDetails
		}

		if r.LocalAdgAutoFailoverMaxDataLossLimit != nil {
			autonomousDatabase["local_adg_auto_failover_max_data_loss_limit"] = *r.LocalAdgAutoFailoverMaxDataLossLimit
		}

		autonomousDatabase["local_disaster_recovery_type"] = r.LocalDisasterRecoveryType

		if r.LocalStandbyDb != nil {
			autonomousDatabase["local_standby_db"] = []interface{}{AutonomousDatabaseStandbySummaryToMap(r.LocalStandbyDb)}
		} else {
			autonomousDatabase["local_standby_db"] = nil
		}

		if r.LongTermBackupSchedule != nil {
			autonomousDatabase["long_term_backup_schedule"] = []interface{}{LongTermBackUpScheduleDetailsToMap(r.LongTermBackupSchedule)}
		} else {
			autonomousDatabase["long_term_backup_schedule"] = nil
		}

		if r.MemoryPerOracleComputeUnitInGBs != nil {
			autonomousDatabase["memory_per_oracle_compute_unit_in_gbs"] = *r.MemoryPerOracleComputeUnitInGBs
		}

		if r.NcharacterSet != nil {
			autonomousDatabase["ncharacter_set"] = *r.NcharacterSet
		}

		autonomousDatabase["net_services_architecture"] = r.NetServicesArchitecture

		if r.NextLongTermBackupTimeStamp != nil {
			autonomousDatabase["next_long_term_backup_time_stamp"] = r.NextLongTermBackupTimeStamp.String()
		}

		autonomousDatabase["nsg_ids"] = r.NsgIds

		if r.OcpuCount != nil {
			autonomousDatabase["ocpu_count"] = *r.OcpuCount
		}

		autonomousDatabase["open_mode"] = r.OpenMode

		autonomousDatabase["operations_insights_status"] = r.OperationsInsightsStatus

		autonomousDatabase["peer_db_ids"] = r.PeerDbIds

		autonomousDatabase["permission_level"] = r.PermissionLevel

		if r.PrivateEndpoint != nil {
			autonomousDatabase["private_endpoint"] = *r.PrivateEndpoint
		}

		if r.PrivateEndpointIp != nil {
			autonomousDatabase["private_endpoint_ip"] = *r.PrivateEndpointIp
		}

		if r.PrivateEndpointLabel != nil {
			autonomousDatabase["private_endpoint_label"] = *r.PrivateEndpointLabel
		}

		autonomousDatabase["provisionable_cpus"] = r.ProvisionableCpus

		autonomousDatabase["refreshable_mode"] = r.RefreshableMode

		autonomousDatabase["refreshable_status"] = r.RefreshableStatus

		if r.RemoteDisasterRecoveryConfiguration != nil {
			autonomousDatabase["remote_disaster_recovery_configuration"] = []interface{}{DisasterRecoveryConfigurationToMap(r.RemoteDisasterRecoveryConfiguration)}
		} else {
			autonomousDatabase["remote_disaster_recovery_configuration"] = nil
		}

		if r.ResourcePoolLeaderId != nil {
			autonomousDatabase["resource_pool_leader_id"] = *r.ResourcePoolLeaderId
		}

		if r.ResourcePoolSummary != nil {
			autonomousDatabase["resource_pool_summary"] = []interface{}{ResourcePoolSummaryToMap(r.ResourcePoolSummary)}
		} else {
			autonomousDatabase["resource_pool_summary"] = nil
		}

		autonomousDatabase["role"] = r.Role

		scheduledOperations := []interface{}{}
		for _, item := range r.ScheduledOperations {
			scheduledOperations = append(scheduledOperations, ScheduledOperationDetailsToMap(item))
		}
		autonomousDatabase["scheduled_operations"] = scheduledOperations

		if r.ServiceConsoleUrl != nil {
			autonomousDatabase["service_console_url"] = *r.ServiceConsoleUrl
		}

		if r.SourceId != nil {
			autonomousDatabase["source_id"] = *r.SourceId
		}

		if r.StandbyDb != nil {
			autonomousDatabase["standby_db"] = []interface{}{AutonomousDatabaseStandbySummaryToMap(r.StandbyDb)}
		} else {
			autonomousDatabase["standby_db"] = nil
		}

		autonomousDatabase["standby_whitelisted_ips"] = r.StandbyWhitelistedIps

		autonomousDatabase["state"] = r.LifecycleState

		if r.SubnetId != nil {
			autonomousDatabase["subnet_id"] = *r.SubnetId
		}

		autonomousDatabase["supported_regions_to_clone_to"] = r.SupportedRegionsToCloneTo

		if r.SystemTags != nil {
			autonomousDatabase["system_tags"] = tfresource.SystemTagsToMap(r.SystemTags)
		}

		if r.TimeCreated != nil {
			autonomousDatabase["time_created"] = r.TimeCreated.String()
		}

		if r.TimeDataGuardRoleChanged != nil {
			autonomousDatabase["time_data_guard_role_changed"] = r.TimeDataGuardRoleChanged.String()
		}

		if r.TimeDeletionOfFreeAutonomousDatabase != nil {
			autonomousDatabase["time_deletion_of_free_autonomous_database"] = r.TimeDeletionOfFreeAutonomousDatabase.String()
		}

		if r.TimeDisasterRecoveryRoleChanged != nil {
			autonomousDatabase["time_disaster_recovery_role_changed"] = r.TimeDisasterRecoveryRoleChanged.String()
		}

		if r.TimeLocalDataGuardEnabled != nil {
			autonomousDatabase["time_local_data_guard_enabled"] = r.TimeLocalDataGuardEnabled.String()
		}

		if r.TimeMaintenanceBegin != nil {
			autonomousDatabase["time_maintenance_begin"] = r.TimeMaintenanceBegin.String()
		}

		if r.TimeMaintenanceEnd != nil {
			autonomousDatabase["time_maintenance_end"] = r.TimeMaintenanceEnd.String()
		}

		if r.TimeOfJoiningResourcePool != nil {
			autonomousDatabase["time_of_joining_resource_pool"] = r.TimeOfJoiningResourcePool.String()
		}

		if r.TimeOfLastFailover != nil {
			autonomousDatabase["time_of_last_failover"] = r.TimeOfLastFailover.String()
		}

		if r.TimeOfLastRefresh != nil {
			autonomousDatabase["time_of_last_refresh"] = r.TimeOfLastRefresh.String()
		}

		if r.TimeOfLastRefreshPoint != nil {
			autonomousDatabase["time_of_last_refresh_point"] = r.TimeOfLastRefreshPoint.String()
		}

		if r.TimeOfLastSwitchover != nil {
			autonomousDatabase["time_of_last_switchover"] = r.TimeOfLastSwitchover.String()
		}

		if r.TimeOfNextRefresh != nil {
			autonomousDatabase["time_of_next_refresh"] = r.TimeOfNextRefresh.String()
		}

		if r.TimeReclamationOfFreeAutonomousDatabase != nil {
			autonomousDatabase["time_reclamation_of_free_autonomous_database"] = r.TimeReclamationOfFreeAutonomousDatabase.String()
		}

		if r.TimeUntilReconnectCloneEnabled != nil {
			autonomousDatabase["time_until_reconnect_clone_enabled"] = r.TimeUntilReconnectCloneEnabled.String()
		}

		if r.TotalBackupStorageSizeInGBs != nil {
			autonomousDatabase["total_backup_storage_size_in_gbs"] = *r.TotalBackupStorageSizeInGBs
		}

		if r.UsedDataStorageSizeInGBs != nil {
			autonomousDatabase["used_data_storage_size_in_gbs"] = *r.UsedDataStorageSizeInGBs
		}

		if r.UsedDataStorageSizeInTBs != nil {
			autonomousDatabase["used_data_storage_size_in_tbs"] = *r.UsedDataStorageSizeInTBs
		}

		if r.VaultId != nil {
			autonomousDatabase["vault_id"] = *r.VaultId
		}

		autonomousDatabase["whitelisted_ips"] = r.WhitelistedIps

		resources = append(resources, autonomousDatabase)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabaseAutonomousDatabasesDataSource().Schema["autonomous_databases"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("autonomous_databases", resources); err != nil {
		return err
	}

	return nil
}
