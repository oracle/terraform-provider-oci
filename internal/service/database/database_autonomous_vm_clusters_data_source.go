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

func DatabaseAutonomousVmClustersDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseAutonomousVmClusters,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"exadata_infrastructure_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"autonomous_vm_clusters": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(DatabaseAutonomousVmClusterResource()),
			},
		},
	}
}

func readDatabaseAutonomousVmClusters(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousVmClustersDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseAutonomousVmClustersDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListAutonomousVmClustersResponse
}

func (s *DatabaseAutonomousVmClustersDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseAutonomousVmClustersDataSourceCrud) Get() error {
	request := oci_database.ListAutonomousVmClustersRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if exadataInfrastructureId, ok := s.D.GetOkExists("exadata_infrastructure_id"); ok {
		tmp := exadataInfrastructureId.(string)
		request.ExadataInfrastructureId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_database.AutonomousVmClusterSummaryLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListAutonomousVmClusters(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAutonomousVmClusters(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseAutonomousVmClustersDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseAutonomousVmClustersDataSource-", DatabaseAutonomousVmClustersDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		autonomousVmCluster := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.AutonomousDataStoragePercentage != nil {
			autonomousVmCluster["autonomous_data_storage_percentage"] = *r.AutonomousDataStoragePercentage
		}

		if r.AutonomousDataStorageSizeInTBs != nil {
			autonomousVmCluster["autonomous_data_storage_size_in_tbs"] = *r.AutonomousDataStorageSizeInTBs
		}

		if r.AvailableAutonomousDataStorageSizeInTBs != nil {
			autonomousVmCluster["available_autonomous_data_storage_size_in_tbs"] = *r.AvailableAutonomousDataStorageSizeInTBs
		}

		if r.AvailableContainerDatabases != nil {
			autonomousVmCluster["available_container_databases"] = *r.AvailableContainerDatabases
		}

		if r.AvailableCpus != nil {
			autonomousVmCluster["available_cpus"] = *r.AvailableCpus
		}

		if r.AvailableDataStorageSizeInTBs != nil {
			autonomousVmCluster["available_data_storage_size_in_tbs"] = *r.AvailableDataStorageSizeInTBs
		}

		autonomousVmCluster["compute_model"] = r.ComputeModel

		if r.CpuCoreCountPerNode != nil {
			autonomousVmCluster["cpu_core_count_per_node"] = *r.CpuCoreCountPerNode
		}

		if r.CpuPercentage != nil {
			autonomousVmCluster["cpu_percentage"] = *r.CpuPercentage
		}

		if r.CpusEnabled != nil {
			autonomousVmCluster["cpus_enabled"] = *r.CpusEnabled
		}

		if r.CpusLowestScaledValue != nil {
			autonomousVmCluster["cpus_lowest_scaled_value"] = *r.CpusLowestScaledValue
		}

		if r.DataStorageSizeInGBs != nil {
			autonomousVmCluster["data_storage_size_in_gb"] = *r.DataStorageSizeInGBs
		}

		if r.DataStorageSizeInTBs != nil {
			autonomousVmCluster["data_storage_size_in_tbs"] = *r.DataStorageSizeInTBs
		}

		if r.DbNodeStorageSizeInGBs != nil {
			autonomousVmCluster["db_node_storage_size_in_gbs"] = *r.DbNodeStorageSizeInGBs
		}

		autonomousVmCluster["db_servers"] = r.DbServers

		if r.DefinedTags != nil {
			autonomousVmCluster["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			autonomousVmCluster["display_name"] = *r.DisplayName
		}

		if r.ExadataInfrastructureId != nil {
			autonomousVmCluster["exadata_infrastructure_id"] = *r.ExadataInfrastructureId
		}

		if r.ExadataStorageInTBsLowestScaledValue != nil {
			autonomousVmCluster["exadata_storage_in_tbs_lowest_scaled_value"] = *r.ExadataStorageInTBsLowestScaledValue
		}

		autonomousVmCluster["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			autonomousVmCluster["id"] = *r.Id
		}

		if r.IsLocalBackupEnabled != nil {
			autonomousVmCluster["is_local_backup_enabled"] = *r.IsLocalBackupEnabled
		}

		if r.IsMtlsEnabled != nil {
			autonomousVmCluster["is_mtls_enabled"] = *r.IsMtlsEnabled
		}

		if r.LastMaintenanceRunId != nil {
			autonomousVmCluster["last_maintenance_run_id"] = *r.LastMaintenanceRunId
		}

		autonomousVmCluster["license_model"] = r.LicenseModel

		if r.LifecycleDetails != nil {
			autonomousVmCluster["lifecycle_details"] = *r.LifecycleDetails
		}

		if r.MaintenanceWindow != nil {
			autonomousVmCluster["maintenance_window"] = []interface{}{MaintenanceWindowToMap(r.MaintenanceWindow)}
		} else {
			autonomousVmCluster["maintenance_window"] = nil
		}

		if r.MaxAcdsLowestScaledValue != nil {
			autonomousVmCluster["max_acds_lowest_scaled_value"] = *r.MaxAcdsLowestScaledValue
		}

		if r.MemoryPerOracleComputeUnitInGBs != nil {
			autonomousVmCluster["memory_per_oracle_compute_unit_in_gbs"] = *r.MemoryPerOracleComputeUnitInGBs
		}

		if r.MemorySizeInGBs != nil {
			autonomousVmCluster["memory_size_in_gbs"] = *r.MemorySizeInGBs
		}

		if r.NextMaintenanceRunId != nil {
			autonomousVmCluster["next_maintenance_run_id"] = *r.NextMaintenanceRunId
		}

		if r.NodeCount != nil {
			autonomousVmCluster["node_count"] = *r.NodeCount
		}

		if r.NonProvisionableAutonomousContainerDatabases != nil {
			autonomousVmCluster["non_provisionable_autonomous_container_databases"] = *r.NonProvisionableAutonomousContainerDatabases
		}

		if r.OcpusEnabled != nil {
			autonomousVmCluster["ocpus_enabled"] = *r.OcpusEnabled
		}

		if r.ProvisionableAutonomousContainerDatabases != nil {
			autonomousVmCluster["provisionable_autonomous_container_databases"] = *r.ProvisionableAutonomousContainerDatabases
		}

		if r.ProvisionedAutonomousContainerDatabases != nil {
			autonomousVmCluster["provisioned_autonomous_container_databases"] = *r.ProvisionedAutonomousContainerDatabases
		}

		if r.ProvisionedCpus != nil {
			autonomousVmCluster["provisioned_cpus"] = *r.ProvisionedCpus
		}

		if r.ReclaimableCpus != nil {
			autonomousVmCluster["reclaimable_cpus"] = *r.ReclaimableCpus
		}

		if r.ReservedCpus != nil {
			autonomousVmCluster["reserved_cpus"] = *r.ReservedCpus
		}

		if r.ScanListenerPortNonTls != nil {
			autonomousVmCluster["scan_listener_port_non_tls"] = *r.ScanListenerPortNonTls
		}

		if r.ScanListenerPortTls != nil {
			autonomousVmCluster["scan_listener_port_tls"] = *r.ScanListenerPortTls
		}

		autonomousVmCluster["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			autonomousVmCluster["time_created"] = r.TimeCreated.String()
		}

		if r.TimeDatabaseSslCertificateExpires != nil {
			autonomousVmCluster["time_database_ssl_certificate_expires"] = r.TimeDatabaseSslCertificateExpires.String()
		}

		if r.TimeOrdsCertificateExpires != nil {
			autonomousVmCluster["time_ords_certificate_expires"] = r.TimeOrdsCertificateExpires.String()
		}

		if r.TimeZone != nil {
			autonomousVmCluster["time_zone"] = *r.TimeZone
		}

		if r.TotalAutonomousDataStorageInTBs != nil {
			autonomousVmCluster["total_autonomous_data_storage_in_tbs"] = *r.TotalAutonomousDataStorageInTBs
		}

		if r.TotalContainerDatabases != nil {
			autonomousVmCluster["total_container_databases"] = *r.TotalContainerDatabases
		}

		if r.VmClusterNetworkId != nil {
			autonomousVmCluster["vm_cluster_network_id"] = *r.VmClusterNetworkId
		}

		resources = append(resources, autonomousVmCluster)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabaseAutonomousVmClustersDataSource().Schema["autonomous_vm_clusters"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("autonomous_vm_clusters", resources); err != nil {
		return err
	}

	return nil
}
