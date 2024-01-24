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

func DatabaseAutonomousContainerDatabaseDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["autonomous_container_database_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DatabaseAutonomousContainerDatabaseResource(), fieldMap, readSingularDatabaseAutonomousContainerDatabase)
}

func readSingularDatabaseAutonomousContainerDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousContainerDatabaseDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseAutonomousContainerDatabaseDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.GetAutonomousContainerDatabaseResponse
}

func (s *DatabaseAutonomousContainerDatabaseDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseAutonomousContainerDatabaseDataSourceCrud) Get() error {
	request := oci_database.GetAutonomousContainerDatabaseRequest{}

	if autonomousContainerDatabaseId, ok := s.D.GetOkExists("autonomous_container_database_id"); ok {
		tmp := autonomousContainerDatabaseId.(string)
		request.AutonomousContainerDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.GetAutonomousContainerDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseAutonomousContainerDatabaseDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AutonomousExadataInfrastructureId != nil {
		s.D.Set("autonomous_exadata_infrastructure_id", *s.Res.AutonomousExadataInfrastructureId)
	}

	if s.Res.AutonomousVmClusterId != nil {
		s.D.Set("autonomous_vm_cluster_id", *s.Res.AutonomousVmClusterId)
	}

	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.AvailableCpus != nil {
		s.D.Set("available_cpus", *s.Res.AvailableCpus)
	}

	if s.Res.BackupConfig != nil {
		s.D.Set("backup_config", []interface{}{AutonomousContainerDatabaseBackupConfigToMap(s.Res.BackupConfig, nil, true)})
	} else {
		s.D.Set("backup_config", nil)
	}

	if s.Res.CloudAutonomousVmClusterId != nil {
		s.D.Set("cloud_autonomous_vm_cluster_id", *s.Res.CloudAutonomousVmClusterId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("compute_model", s.Res.ComputeModel)

	if s.Res.DbName != nil {
		s.D.Set("db_name", *s.Res.DbName)
	}

	if s.Res.DbUniqueName != nil {
		s.D.Set("db_unique_name", *s.Res.DbUniqueName)
	}

	if s.Res.DbVersion != nil {
		s.D.Set("db_version", *s.Res.DbVersion)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.DstFileVersion != nil {
		s.D.Set("dst_file_version", *s.Res.DstFileVersion)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("infrastructure_type", s.Res.InfrastructureType)

	if s.Res.IsDstFileUpdateEnabled != nil {
		s.D.Set("is_dst_file_update_enabled", *s.Res.IsDstFileUpdateEnabled)
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

	if s.Res.KmsKeyVersionId != nil {
		s.D.Set("kms_key_version_id", *s.Res.KmsKeyVersionId)
	}

	if s.Res.LargestProvisionableAutonomousDatabaseInCpus != nil {
		s.D.Set("largest_provisionable_autonomous_database_in_cpus", *s.Res.LargestProvisionableAutonomousDatabaseInCpus)
	}

	if s.Res.LastMaintenanceRunId != nil {
		s.D.Set("last_maintenance_run_id", *s.Res.LastMaintenanceRunId)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.MaintenanceWindow != nil {
		s.D.Set("maintenance_window", []interface{}{MaintenanceWindowToMap(s.Res.MaintenanceWindow)})
	} else {
		s.D.Set("maintenance_window", nil)
	}

	if s.Res.MemoryPerOracleComputeUnitInGBs != nil {
		s.D.Set("memory_per_oracle_compute_unit_in_gbs", *s.Res.MemoryPerOracleComputeUnitInGBs)
	}

	if s.Res.NextMaintenanceRunId != nil {
		s.D.Set("next_maintenance_run_id", *s.Res.NextMaintenanceRunId)
	}

	if s.Res.PatchId != nil {
		s.D.Set("patch_id", *s.Res.PatchId)
	}

	s.D.Set("patch_model", s.Res.PatchModel)

	s.D.Set("provisionable_cpus", s.Res.ProvisionableCpus)

	if s.Res.ProvisionedCpus != nil {
		s.D.Set("provisioned_cpus", *s.Res.ProvisionedCpus)
	}

	if s.Res.ReclaimableCpus != nil {
		s.D.Set("reclaimable_cpus", *s.Res.ReclaimableCpus)
	}

	if s.Res.ReservedCpus != nil {
		s.D.Set("reserved_cpus", *s.Res.ReservedCpus)
	}

	s.D.Set("role", s.Res.Role)

	s.D.Set("service_level_agreement_type", s.Res.ServiceLevelAgreementType)

	if s.Res.StandbyMaintenanceBufferInDays != nil {
		s.D.Set("standby_maintenance_buffer_in_days", *s.Res.StandbyMaintenanceBufferInDays)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeOfLastBackup != nil {
		s.D.Set("time_of_last_backup", s.Res.TimeOfLastBackup.String())
	}

	if s.Res.TimeSnapshotStandbyRevert != nil {
		s.D.Set("time_snapshot_standby_revert", s.Res.TimeSnapshotStandbyRevert.String())
	}

	if s.Res.TotalCpus != nil {
		s.D.Set("total_cpus", *s.Res.TotalCpus)
	}

	if s.Res.VaultId != nil {
		s.D.Set("vault_id", *s.Res.VaultId)
	}

	s.D.Set("version_preference", s.Res.VersionPreference)

	return nil
}
