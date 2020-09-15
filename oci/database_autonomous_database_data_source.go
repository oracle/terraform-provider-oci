// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v25/database"
)

func init() {
	RegisterDatasource("oci_database_autonomous_database", DatabaseAutonomousDatabaseDataSource())
}

func DatabaseAutonomousDatabaseDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["autonomous_database_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(DatabaseAutonomousDatabaseResource(), fieldMap, readSingularDatabaseAutonomousDatabase)
}

func readSingularDatabaseAutonomousDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousDatabaseDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient()

	return ReadResource(sync)
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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "database")

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

	if s.Res.AutonomousContainerDatabaseId != nil {
		s.D.Set("autonomous_container_database_id", *s.Res.AutonomousContainerDatabaseId)
	}

	s.D.Set("available_upgrade_versions", s.Res.AvailableUpgradeVersions)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

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

	s.D.Set("data_safe_status", s.Res.DataSafeStatus)

	if s.Res.DataStorageSizeInTBs != nil {
		s.D.Set("data_storage_size_in_tbs", *s.Res.DataStorageSizeInTBs)
	}

	if s.Res.DbName != nil {
		s.D.Set("db_name", *s.Res.DbName)
	}

	if s.Res.DbVersion != nil {
		s.D.Set("db_version", *s.Res.DbVersion)
	}

	s.D.Set("db_workload", s.Res.DbWorkload)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.FailedDataRecoveryInSeconds != nil {
		s.D.Set("failed_data_recovery_in_seconds", *s.Res.FailedDataRecoveryInSeconds)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("infrastructure_type", s.Res.InfrastructureType)

	if s.Res.IsAutoScalingEnabled != nil {
		s.D.Set("is_auto_scaling_enabled", *s.Res.IsAutoScalingEnabled)
	}

	if s.Res.IsDataGuardEnabled != nil {
		s.D.Set("is_data_guard_enabled", *s.Res.IsDataGuardEnabled)
	}

	if s.Res.IsDedicated != nil {
		s.D.Set("is_dedicated", *s.Res.IsDedicated)
	}

	if s.Res.IsFreeTier != nil {
		s.D.Set("is_free_tier", *s.Res.IsFreeTier)
	}

	if s.Res.IsPreview != nil {
		s.D.Set("is_preview", *s.Res.IsPreview)
	}

	if s.Res.IsRefreshableClone != nil {
		s.D.Set("is_refreshable_clone", *s.Res.IsRefreshableClone)
	}

	s.D.Set("license_model", s.Res.LicenseModel)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("nsg_ids", s.Res.NsgIds)

	s.D.Set("open_mode", s.Res.OpenMode)

	if s.Res.PrivateEndpoint != nil {
		s.D.Set("private_endpoint", *s.Res.PrivateEndpoint)
	}

	if s.Res.PrivateEndpointIp != nil {
		s.D.Set("private_endpoint_ip", *s.Res.PrivateEndpointIp)
	}

	if s.Res.PrivateEndpointLabel != nil {
		s.D.Set("private_endpoint_label", *s.Res.PrivateEndpointLabel)
	}

	s.D.Set("refreshable_mode", s.Res.RefreshableMode)

	s.D.Set("refreshable_status", s.Res.RefreshableStatus)

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

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", *s.Res.SubnetId)
	}

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", systemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeDeletionOfFreeAutonomousDatabase != nil {
		s.D.Set("time_deletion_of_free_autonomous_database", s.Res.TimeDeletionOfFreeAutonomousDatabase.String())
	}

	if s.Res.TimeMaintenanceBegin != nil {
		s.D.Set("time_maintenance_begin", s.Res.TimeMaintenanceBegin.String())
	}

	if s.Res.TimeMaintenanceEnd != nil {
		s.D.Set("time_maintenance_end", s.Res.TimeMaintenanceEnd.String())
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

	if s.Res.UsedDataStorageSizeInTBs != nil {
		s.D.Set("used_data_storage_size_in_tbs", *s.Res.UsedDataStorageSizeInTBs)
	}

	s.D.Set("whitelisted_ips", s.Res.WhitelistedIps)

	return nil
}
