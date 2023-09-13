// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"
)

func DatabaseCloudAutonomousVmClustersDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseCloudAutonomousVmClusters,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"availability_domain": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cloud_exadata_infrastructure_id": {
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
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cloud_autonomous_vm_clusters": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(DatabaseCloudAutonomousVmClusterResource()),
			},
		},
	}
}

func readDatabaseCloudAutonomousVmClusters(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseCloudAutonomousVmClustersDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseCloudAutonomousVmClustersDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListCloudAutonomousVmClustersResponse
}

func (s *DatabaseCloudAutonomousVmClustersDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseCloudAutonomousVmClustersDataSourceCrud) Get() error {
	request := oci_database.ListCloudAutonomousVmClustersRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if cloudExadataInfrastructureId, ok := s.D.GetOkExists("cloud_exadata_infrastructure_id"); ok {
		tmp := cloudExadataInfrastructureId.(string)
		request.CloudExadataInfrastructureId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_database.CloudAutonomousVmClusterSummaryLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListCloudAutonomousVmClusters(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListCloudAutonomousVmClusters(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseCloudAutonomousVmClustersDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseCloudAutonomousVmClustersDataSource-", DatabaseCloudAutonomousVmClustersDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		cloudAutonomousVmCluster := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.AutonomousDataStoragePercentage != nil {
			cloudAutonomousVmCluster["autonomous_data_storage_percentage"] = *r.AutonomousDataStoragePercentage
		}

		if r.AutonomousDataStorageSizeInTBs != nil {
			cloudAutonomousVmCluster["autonomous_data_storage_size_in_tbs"] = *r.AutonomousDataStorageSizeInTBs
		}

		if r.AvailabilityDomain != nil {
			cloudAutonomousVmCluster["availability_domain"] = *r.AvailabilityDomain
		}

		if r.AvailableAutonomousDataStorageSizeInTBs != nil {
			cloudAutonomousVmCluster["available_autonomous_data_storage_size_in_tbs"] = *r.AvailableAutonomousDataStorageSizeInTBs
		}

		if r.AvailableContainerDatabases != nil {
			cloudAutonomousVmCluster["available_container_databases"] = *r.AvailableContainerDatabases
		}

		if r.AvailableCpus != nil {
			cloudAutonomousVmCluster["available_cpus"] = *r.AvailableCpus
		}

		if r.CloudExadataInfrastructureId != nil {
			cloudAutonomousVmCluster["cloud_exadata_infrastructure_id"] = *r.CloudExadataInfrastructureId
		}

		if r.ClusterTimeZone != nil {
			cloudAutonomousVmCluster["cluster_time_zone"] = *r.ClusterTimeZone
		}

		cloudAutonomousVmCluster["compute_model"] = r.ComputeModel

		if r.CpuCoreCount != nil {
			cloudAutonomousVmCluster["cpu_core_count"] = *r.CpuCoreCount
		}

		if r.CpuCoreCountPerNode != nil {
			cloudAutonomousVmCluster["cpu_core_count_per_node"] = *r.CpuCoreCountPerNode
		}

		if r.CpuPercentage != nil {
			cloudAutonomousVmCluster["cpu_percentage"] = *r.CpuPercentage
		}

		if r.DataStorageSizeInGBs != nil {
			cloudAutonomousVmCluster["data_storage_size_in_gb"] = *r.DataStorageSizeInGBs
		}

		if r.DataStorageSizeInTBs != nil {
			cloudAutonomousVmCluster["data_storage_size_in_tbs"] = *r.DataStorageSizeInTBs
		}

		if r.DbNodeStorageSizeInGBs != nil {
			cloudAutonomousVmCluster["db_node_storage_size_in_gbs"] = *r.DbNodeStorageSizeInGBs
		}

		cloudAutonomousVmCluster["db_servers"] = r.DbServers

		if r.DefinedTags != nil {
			cloudAutonomousVmCluster["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.Description != nil {
			cloudAutonomousVmCluster["description"] = *r.Description
		}

		if r.DisplayName != nil {
			cloudAutonomousVmCluster["display_name"] = *r.DisplayName
		}

		if r.Domain != nil {
			cloudAutonomousVmCluster["domain"] = *r.Domain
		}

		if r.ExadataStorageInTBsLowestScaledValue != nil {
			cloudAutonomousVmCluster["exadata_storage_in_tbs_lowest_scaled_value"] = *r.ExadataStorageInTBsLowestScaledValue
		}

		cloudAutonomousVmCluster["freeform_tags"] = r.FreeformTags

		if r.Hostname != nil {
			cloudAutonomousVmCluster["hostname"] = *r.Hostname
		}

		if r.Id != nil {
			cloudAutonomousVmCluster["id"] = *r.Id
		}

		if r.IsMtlsEnabledVmCluster != nil {
			cloudAutonomousVmCluster["is_mtls_enabled_vm_cluster"] = *r.IsMtlsEnabledVmCluster
		}

		if r.LastMaintenanceRunId != nil {
			cloudAutonomousVmCluster["last_maintenance_run_id"] = *r.LastMaintenanceRunId
		}

		if r.LastUpdateHistoryEntryId != nil {
			cloudAutonomousVmCluster["last_update_history_entry_id"] = *r.LastUpdateHistoryEntryId
		}

		cloudAutonomousVmCluster["license_model"] = r.LicenseModel

		if r.LifecycleDetails != nil {
			cloudAutonomousVmCluster["lifecycle_details"] = *r.LifecycleDetails
		}

		if r.MaintenanceWindow != nil {
			cloudAutonomousVmCluster["maintenance_window"] = []interface{}{MaintenanceWindowToMap(r.MaintenanceWindow)}
		} else {
			cloudAutonomousVmCluster["maintenance_window"] = nil
		}

		if r.MaxAcdsLowestScaledValue != nil {
			cloudAutonomousVmCluster["max_acds_lowest_scaled_value"] = *r.MaxAcdsLowestScaledValue
		}

		if r.MemoryPerOracleComputeUnitInGBs != nil {
			cloudAutonomousVmCluster["memory_per_oracle_compute_unit_in_gbs"] = *r.MemoryPerOracleComputeUnitInGBs
		}

		if r.MemorySizeInGBs != nil {
			cloudAutonomousVmCluster["memory_size_in_gbs"] = *r.MemorySizeInGBs
		}

		if r.NextMaintenanceRunId != nil {
			cloudAutonomousVmCluster["next_maintenance_run_id"] = *r.NextMaintenanceRunId
		}

		if r.NodeCount != nil {
			cloudAutonomousVmCluster["node_count"] = *r.NodeCount
		}

		if r.NonProvisionableAutonomousContainerDatabases != nil {
			cloudAutonomousVmCluster["non_provisionable_autonomous_container_databases"] = *r.NonProvisionableAutonomousContainerDatabases
		}

		cloudAutonomousVmCluster["nsg_ids"] = r.NsgIds

		if r.OcpuCount != nil {
			cloudAutonomousVmCluster["ocpu_count"] = *r.OcpuCount
		}

		if r.OcpusLowestScaledValue != nil {
			cloudAutonomousVmCluster["ocpus_lowest_scaled_value"] = *r.OcpusLowestScaledValue
		}

		if r.ProvisionableAutonomousContainerDatabases != nil {
			cloudAutonomousVmCluster["provisionable_autonomous_container_databases"] = *r.ProvisionableAutonomousContainerDatabases
		}

		if r.ProvisionedAutonomousContainerDatabases != nil {
			cloudAutonomousVmCluster["provisioned_autonomous_container_databases"] = *r.ProvisionedAutonomousContainerDatabases
		}

		if r.ProvisionedCpus != nil {
			cloudAutonomousVmCluster["provisioned_cpus"] = *r.ProvisionedCpus
		}

		if r.ReclaimableCpus != nil {
			cloudAutonomousVmCluster["reclaimable_cpus"] = *r.ReclaimableCpus
		}

		if r.ReservedCpus != nil {
			cloudAutonomousVmCluster["reserved_cpus"] = *r.ReservedCpus
		}

		if r.ScanListenerPortNonTls != nil {
			cloudAutonomousVmCluster["scan_listener_port_non_tls"] = *r.ScanListenerPortNonTls
		}

		if r.ScanListenerPortTls != nil {
			cloudAutonomousVmCluster["scan_listener_port_tls"] = *r.ScanListenerPortTls
		}

		if r.Shape != nil {
			cloudAutonomousVmCluster["shape"] = *r.Shape
		}

		cloudAutonomousVmCluster["state"] = r.LifecycleState

		if r.SubnetId != nil {
			cloudAutonomousVmCluster["subnet_id"] = *r.SubnetId
		}

		if r.TimeCreated != nil {
			cloudAutonomousVmCluster["time_created"] = r.TimeCreated.String()
		}

		if r.TimeUpdated != nil {
			cloudAutonomousVmCluster["time_updated"] = r.TimeUpdated.String()
		}

		if r.TotalAutonomousDataStorageInTBs != nil {
			cloudAutonomousVmCluster["total_autonomous_data_storage_in_tbs"] = *r.TotalAutonomousDataStorageInTBs
		}

		if r.TotalContainerDatabases != nil {
			cloudAutonomousVmCluster["total_container_databases"] = *r.TotalContainerDatabases
		}

		if r.TotalCpus != nil {
			cloudAutonomousVmCluster["total_cpus"] = *r.TotalCpus
		}

		resources = append(resources, cloudAutonomousVmCluster)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabaseCloudAutonomousVmClustersDataSource().Schema["cloud_autonomous_vm_clusters"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("cloud_autonomous_vm_clusters", resources); err != nil {
		return err
	}

	return nil
}
