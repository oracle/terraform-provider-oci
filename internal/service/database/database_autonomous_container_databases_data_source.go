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

func DatabaseAutonomousContainerDatabasesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseAutonomousContainerDatabases,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"autonomous_exadata_infrastructure_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"autonomous_vm_cluster_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"availability_domain": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cloud_autonomous_vm_cluster_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"infrastructure_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_level_agreement_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"autonomous_container_databases": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(DatabaseAutonomousContainerDatabaseResource()),
			},
		},
	}
}

func readDatabaseAutonomousContainerDatabases(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousContainerDatabasesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseAutonomousContainerDatabasesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListAutonomousContainerDatabasesResponse
}

func (s *DatabaseAutonomousContainerDatabasesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseAutonomousContainerDatabasesDataSourceCrud) Get() error {
	request := oci_database.ListAutonomousContainerDatabasesRequest{}

	if autonomousExadataInfrastructureId, ok := s.D.GetOkExists("autonomous_exadata_infrastructure_id"); ok {
		tmp := autonomousExadataInfrastructureId.(string)
		request.AutonomousExadataInfrastructureId = &tmp
	}

	if autonomousVmClusterId, ok := s.D.GetOkExists("autonomous_vm_cluster_id"); ok {
		tmp := autonomousVmClusterId.(string)
		request.AutonomousVmClusterId = &tmp
	}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if cloudAutonomousVmClusterId, ok := s.D.GetOkExists("cloud_autonomous_vm_cluster_id"); ok {
		tmp := cloudAutonomousVmClusterId.(string)
		request.CloudAutonomousVmClusterId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if infrastructureType, ok := s.D.GetOkExists("infrastructure_type"); ok {
		request.InfrastructureType = oci_database.AutonomousContainerDatabaseSummaryInfrastructureTypeEnum(infrastructureType.(string))
	}

	if serviceLevelAgreementType, ok := s.D.GetOkExists("service_level_agreement_type"); ok {
		tmp := serviceLevelAgreementType.(string)
		request.ServiceLevelAgreementType = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_database.AutonomousContainerDatabaseSummaryLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListAutonomousContainerDatabases(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAutonomousContainerDatabases(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseAutonomousContainerDatabasesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseAutonomousContainerDatabasesDataSource-", DatabaseAutonomousContainerDatabasesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		autonomousContainerDatabase := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.AutonomousExadataInfrastructureId != nil {
			autonomousContainerDatabase["autonomous_exadata_infrastructure_id"] = *r.AutonomousExadataInfrastructureId
		}

		if r.AutonomousVmClusterId != nil {
			autonomousContainerDatabase["autonomous_vm_cluster_id"] = *r.AutonomousVmClusterId
		}

		if r.AvailabilityDomain != nil {
			autonomousContainerDatabase["availability_domain"] = *r.AvailabilityDomain
		}

		if r.AvailableCpus != nil {
			autonomousContainerDatabase["available_cpus"] = *r.AvailableCpus
		}

		if r.BackupConfig != nil {
			autonomousContainerDatabase["backup_config"] = []interface{}{AutonomousContainerDatabaseBackupConfigToMap(r.BackupConfig, nil, true)}
		} else {
			autonomousContainerDatabase["backup_config"] = nil
		}

		if r.CloudAutonomousVmClusterId != nil {
			autonomousContainerDatabase["cloud_autonomous_vm_cluster_id"] = *r.CloudAutonomousVmClusterId
		}

		autonomousContainerDatabase["compute_model"] = r.ComputeModel

		if r.DbName != nil {
			autonomousContainerDatabase["db_name"] = *r.DbName
		}

		if r.DbSplitThreshold != nil {
			autonomousContainerDatabase["db_split_threshold"] = *r.DbSplitThreshold
		}

		if r.DbUniqueName != nil {
			autonomousContainerDatabase["db_unique_name"] = *r.DbUniqueName
		}

		if r.DbVersion != nil {
			autonomousContainerDatabase["db_version"] = *r.DbVersion
		}

		if r.DefinedTags != nil {
			autonomousContainerDatabase["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			autonomousContainerDatabase["display_name"] = *r.DisplayName
		}

		autonomousContainerDatabase["distribution_affinity"] = r.DistributionAffinity

		if r.DstFileVersion != nil {
			autonomousContainerDatabase["dst_file_version"] = *r.DstFileVersion
		}

		autonomousContainerDatabase["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			autonomousContainerDatabase["id"] = *r.Id
		}

		autonomousContainerDatabase["infrastructure_type"] = r.InfrastructureType

		if r.IsDstFileUpdateEnabled != nil {
			autonomousContainerDatabase["is_dst_file_update_enabled"] = *r.IsDstFileUpdateEnabled
		}

		keyHistoryEntry := []interface{}{}
		for _, item := range r.KeyHistoryEntry {
			keyHistoryEntry = append(keyHistoryEntry, AutonomousDatabaseKeyHistoryEntryToMap(item))
		}
		autonomousContainerDatabase["key_history_entry"] = keyHistoryEntry

		if r.KeyStoreId != nil {
			autonomousContainerDatabase["key_store_id"] = *r.KeyStoreId
		}

		if r.KeyStoreWalletName != nil {
			autonomousContainerDatabase["key_store_wallet_name"] = *r.KeyStoreWalletName
		}

		if r.KmsKeyId != nil {
			autonomousContainerDatabase["kms_key_id"] = *r.KmsKeyId
		}

		if r.KmsKeyVersionId != nil {
			autonomousContainerDatabase["kms_key_version_id"] = *r.KmsKeyVersionId
		}

		if r.LargestProvisionableAutonomousDatabaseInCpus != nil {
			autonomousContainerDatabase["largest_provisionable_autonomous_database_in_cpus"] = *r.LargestProvisionableAutonomousDatabaseInCpus
		}

		if r.LastMaintenanceRunId != nil {
			autonomousContainerDatabase["last_maintenance_run_id"] = *r.LastMaintenanceRunId
		}

		if r.LifecycleDetails != nil {
			autonomousContainerDatabase["lifecycle_details"] = *r.LifecycleDetails
		}

		autonomousContainerDatabase["list_one_off_patches"] = r.ListOneOffPatches

		if r.MaintenanceWindow != nil {
			autonomousContainerDatabase["maintenance_window"] = []interface{}{MaintenanceWindowToMap(r.MaintenanceWindow)}
		} else {
			autonomousContainerDatabase["maintenance_window"] = nil
		}

		if r.MemoryPerOracleComputeUnitInGBs != nil {
			autonomousContainerDatabase["memory_per_oracle_compute_unit_in_gbs"] = *r.MemoryPerOracleComputeUnitInGBs
		}

		autonomousContainerDatabase["net_services_architecture"] = r.NetServicesArchitecture

		if r.NextMaintenanceRunId != nil {
			autonomousContainerDatabase["next_maintenance_run_id"] = *r.NextMaintenanceRunId
		}

		if r.PatchId != nil {
			autonomousContainerDatabase["patch_id"] = *r.PatchId
		}

		autonomousContainerDatabase["patch_model"] = r.PatchModel

		autonomousContainerDatabase["provisionable_cpus"] = r.ProvisionableCpus

		if r.ProvisionedCpus != nil {
			autonomousContainerDatabase["provisioned_cpus"] = *r.ProvisionedCpus
		}

		if r.ReclaimableCpus != nil {
			autonomousContainerDatabase["reclaimable_cpus"] = *r.ReclaimableCpus
		}

		if r.ReservedCpus != nil {
			autonomousContainerDatabase["reserved_cpus"] = *r.ReservedCpus
		}

		autonomousContainerDatabase["role"] = r.Role

		autonomousContainerDatabase["service_level_agreement_type"] = r.ServiceLevelAgreementType

		if r.StandbyMaintenanceBufferInDays != nil {
			autonomousContainerDatabase["standby_maintenance_buffer_in_days"] = *r.StandbyMaintenanceBufferInDays
		}

		autonomousContainerDatabase["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			autonomousContainerDatabase["time_created"] = r.TimeCreated.String()
		}

		if r.TimeOfLastBackup != nil {
			autonomousContainerDatabase["time_of_last_backup"] = r.TimeOfLastBackup.String()
		}

		if r.TimeSnapshotStandbyRevert != nil {
			autonomousContainerDatabase["time_snapshot_standby_revert"] = r.TimeSnapshotStandbyRevert.String()
		}

		if r.TotalCpus != nil {
			autonomousContainerDatabase["total_cpus"] = *r.TotalCpus
		}

		if r.VaultId != nil {
			autonomousContainerDatabase["vault_id"] = *r.VaultId
		}

		autonomousContainerDatabase["version_preference"] = r.VersionPreference

		if r.VmFailoverReservation != nil {
			autonomousContainerDatabase["vm_failover_reservation"] = *r.VmFailoverReservation
		}

		resources = append(resources, autonomousContainerDatabase)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabaseAutonomousContainerDatabasesDataSource().Schema["autonomous_container_databases"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("autonomous_container_databases", resources); err != nil {
		return err
	}

	return nil
}
