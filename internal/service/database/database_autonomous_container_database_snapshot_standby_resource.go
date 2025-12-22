// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_database "github.com/oracle/oci-go-sdk/v65/database"
	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseAutonomousContainerDatabaseSnapshotStandbyResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("2h"),
			Update: tfresource.GetTimeoutDuration("2h"),
			Delete: tfresource.GetTimeoutDuration("2h"),
		},
		Create: createDatabaseAutonomousContainerDatabaseSnapshotStandby,
		Read:   readDatabaseAutonomousContainerDatabaseSnapshotStandby,
		Delete: deleteDatabaseAutonomousContainerDatabaseSnapshotStandby,
		Schema: map[string]*schema.Schema{
			// Required
			"autonomous_container_database_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"role": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"connection_strings_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createDatabaseAutonomousContainerDatabaseSnapshotStandby(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousContainerDatabaseSnapshotStandbyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.CreateResource(d, sync)
}

func readDatabaseAutonomousContainerDatabaseSnapshotStandby(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteDatabaseAutonomousContainerDatabaseSnapshotStandby(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DatabaseAutonomousContainerDatabaseSnapshotStandbyResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.AutonomousContainerDatabase
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *DatabaseAutonomousContainerDatabaseSnapshotStandbyResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseAutonomousContainerDatabaseSnapshotStandbyResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.AutonomousContainerDatabaseLifecycleStateProvisioning),
		string(oci_database.AutonomousContainerDatabaseLifecycleStateRestoring),
	}
}

func (s *DatabaseAutonomousContainerDatabaseSnapshotStandbyResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.AutonomousContainerDatabaseLifecycleStateAvailable),
	}
}

func (s *DatabaseAutonomousContainerDatabaseSnapshotStandbyResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.AutonomousContainerDatabaseLifecycleStateTerminating),
		string(oci_database.AutonomousContainerDatabaseLifecycleStateUnavailable),
	}
}

func (s *DatabaseAutonomousContainerDatabaseSnapshotStandbyResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.AutonomousContainerDatabaseLifecycleStateTerminated),
	}
}

func (s *DatabaseAutonomousContainerDatabaseSnapshotStandbyResourceCrud) Create() error {
	request := oci_database.ConvertStandbyAutonomousContainerDatabaseRequest{}

	if autonomousContainerDatabaseId, ok := s.D.GetOkExists("autonomous_container_database_id"); ok {
		tmp := autonomousContainerDatabaseId.(string)
		request.AutonomousContainerDatabaseId = &tmp
	}

	if connectionStringsType, ok := s.D.GetOkExists("connection_strings_type"); ok {
		request.ConnectionStringsType = oci_database.ConvertStandbyAutonomousContainerDatabaseDetailsConnectionStringsTypeEnum(connectionStringsType.(string))
	}

	if role, ok := s.D.GetOkExists("role"); ok {
		request.Role = oci_database.ConvertStandbyAutonomousContainerDatabaseDetailsRoleEnum(role.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.ConvertStandbyAutonomousContainerDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	s.Res = &response.AutonomousContainerDatabase

	if workId != nil {
		var identifier *string
		var err error
		identifier, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "autonomouscontainerdatabase", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if identifier != nil {
			s.D.SetId(*identifier)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *DatabaseAutonomousContainerDatabaseSnapshotStandbyResourceCrud) SetData() error {
	associatedBackupConfigurationDetails := []interface{}{}
	for _, item := range s.Res.AssociatedBackupConfigurationDetails {
		associatedBackupConfigurationDetails = append(associatedBackupConfigurationDetails, BackupDestinationConfigurationSummaryToMap(item))
	}
	s.D.Set("associated_backup_configuration_details", associatedBackupConfigurationDetails)

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
		s.D.Set("backup_config", []interface{}{AdbdAutonomousContainerDatabaseBackupConfigToMap(s.Res.BackupConfig)})
	} else {
		s.D.Set("backup_config", nil)
	}

	backupDestinationPropertiesList := []interface{}{}
	for _, item := range s.Res.BackupDestinationPropertiesList {
		backupDestinationPropertiesList = append(backupDestinationPropertiesList, BackupDestinationPropertiesToMap(item))
	}
	s.D.Set("backup_destination_properties_list", backupDestinationPropertiesList)

	if s.Res.CloudAutonomousVmClusterId != nil {
		s.D.Set("cloud_autonomous_vm_cluster_id", *s.Res.CloudAutonomousVmClusterId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("compute_model", s.Res.ComputeModel)

	customerContacts := []interface{}{}
	for _, item := range s.Res.CustomerContacts {
		customerContacts = append(customerContacts, CustomerContactToMap(item))
	}
	s.D.Set("customer_contacts", customerContacts)

	if s.Res.Dataguard != nil {
		s.D.Set("dataguard", []interface{}{AdbdAutonomousContainerDatabaseDataguardToMap(*s.Res.Dataguard)})
	} else {
		s.D.Set("dataguard", nil)
	}

	dataguardGroupMembers := []interface{}{}
	for _, item := range s.Res.DataguardGroupMembers {
		dataguardGroupMembers = append(dataguardGroupMembers, AdbdAutonomousContainerDatabaseDataguardToMap(item))
	}
	s.D.Set("dataguard_group_members", dataguardGroupMembers)

	if s.Res.DbName != nil {
		s.D.Set("db_name", *s.Res.DbName)
	}

	if s.Res.DbSplitThreshold != nil {
		s.D.Set("db_split_threshold", *s.Res.DbSplitThreshold)
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

	s.D.Set("distribution_affinity", s.Res.DistributionAffinity)

	if s.Res.DstFileVersion != nil {
		s.D.Set("dst_file_version", *s.Res.DstFileVersion)
	}

	if s.Res.EncryptionKeyLocationDetails != nil {
		encryptionKeyLocationDetailsArray := []interface{}{}
		if encryptionKeyLocationDetailsMap := AdbdEncryptionKeyLocationDetailsToMap(&s.Res.EncryptionKeyLocationDetails); encryptionKeyLocationDetailsMap != nil {
			encryptionKeyLocationDetailsArray = append(encryptionKeyLocationDetailsArray, encryptionKeyLocationDetailsMap)
		}
		s.D.Set("encryption_key_location_details", encryptionKeyLocationDetailsArray)
	} else {
		s.D.Set("encryption_key_location_details", nil)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("infrastructure_type", s.Res.InfrastructureType)

	if s.Res.IsDataGuardEnabled != nil {
		s.D.Set("is_data_guard_enabled", *s.Res.IsDataGuardEnabled)
	}

	if s.Res.IsDstFileUpdateEnabled != nil {
		s.D.Set("is_dst_file_update_enabled", *s.Res.IsDstFileUpdateEnabled)
	}

	if s.Res.IsMultipleStandby != nil {
		s.D.Set("is_multiple_standby", *s.Res.IsMultipleStandby)
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

	s.D.Set("list_one_off_patches", s.Res.ListOneOffPatches)

	if s.Res.MaintenanceWindow != nil {
		s.D.Set("maintenance_window", []interface{}{MaintenanceWindowToMap(s.Res.MaintenanceWindow)})
	} else {
		s.D.Set("maintenance_window", nil)
	}

	if s.Res.MemoryPerComputeUnitInGBs != nil {
		s.D.Set("memory_per_compute_unit_in_gbs", *s.Res.MemoryPerComputeUnitInGBs)
	}

	if s.Res.MemoryPerOracleComputeUnitInGBs != nil {
		s.D.Set("memory_per_oracle_compute_unit_in_gbs", *s.Res.MemoryPerOracleComputeUnitInGBs)
	}

	s.D.Set("net_services_architecture", s.Res.NetServicesArchitecture)

	if s.Res.NextMaintenanceRunId != nil {
		s.D.Set("next_maintenance_run_id", *s.Res.NextMaintenanceRunId)
	}

	if s.Res.OkvEndPointGroupName != nil {
		s.D.Set("okv_end_point_group_name", *s.Res.OkvEndPointGroupName)
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

	if s.Res.RecoveryApplianceDetails != nil {
		s.D.Set("recovery_appliance_details", []interface{}{RecoveryApplianceDetailsToMap(s.Res.RecoveryApplianceDetails)})
	} else {
		s.D.Set("recovery_appliance_details", nil)
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

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

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

	if s.Res.VmFailoverReservation != nil {
		s.D.Set("vm_failover_reservation", *s.Res.VmFailoverReservation)
	}

	return nil
}

func AdbdAutonomousContainerDatabaseBackupConfigToMap(obj *oci_database.AutonomousContainerDatabaseBackupConfig) map[string]interface{} {
	result := map[string]interface{}{}

	backupDestinationDetails := []interface{}{}
	for _, item := range obj.BackupDestinationDetails {
		backupDestinationDetails = append(backupDestinationDetails, BackupDestinationDetailsToMap(item))
	}
	result["backup_destination_details"] = backupDestinationDetails

	if obj.RecoveryWindowInDays != nil {
		result["recovery_window_in_days"] = int(*obj.RecoveryWindowInDays)
	}

	return result
}

func AdbdAutonomousContainerDatabaseDataguardToMap(obj oci_database.AutonomousContainerDatabaseDataguard) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ApplyLag != nil {
		result["apply_lag"] = string(*obj.ApplyLag)
	}

	if obj.ApplyRate != nil {
		result["apply_rate"] = string(*obj.ApplyRate)
	}

	if obj.AutomaticFailoverTarget != nil {
		result["automatic_failover_target"] = string(*obj.AutomaticFailoverTarget)
	}

	if obj.AutonomousContainerDatabaseId != nil {
		result["autonomous_container_database_id"] = string(*obj.AutonomousContainerDatabaseId)
	}

	if obj.AvailabilityDomain != nil {
		result["availability_domain"] = string(*obj.AvailabilityDomain)
	}

	if obj.FastStartFailOverLagLimitInSeconds != nil {
		result["fast_start_fail_over_lag_limit_in_seconds"] = int(*obj.FastStartFailOverLagLimitInSeconds)
	}

	if obj.IsAutomaticFailoverEnabled != nil {
		result["is_automatic_failover_enabled"] = bool(*obj.IsAutomaticFailoverEnabled)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	result["protection_mode"] = string(obj.ProtectionMode)

	if obj.RedoTransportMode != nil {
		result["redo_transport_mode"] = string(*obj.RedoTransportMode)
	}

	result["role"] = string(obj.Role)

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeLagRefreshedOn != nil {
		result["time_lag_refreshed_on"] = obj.TimeLagRefreshedOn.String()
	}

	if obj.TimeLastRoleChanged != nil {
		result["time_last_role_changed"] = obj.TimeLastRoleChanged.String()
	}

	if obj.TimeLastSynced != nil {
		result["time_last_synced"] = obj.TimeLastSynced.String()
	}

	if obj.TransportLag != nil {
		result["transport_lag"] = string(*obj.TransportLag)
	}

	return result
}
