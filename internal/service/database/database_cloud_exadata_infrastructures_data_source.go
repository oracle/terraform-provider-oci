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

func DatabaseCloudExadataInfrastructuresDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseCloudExadataInfrastructures,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"cluster_placement_group_id": {
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
			"cloud_exadata_infrastructures": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(DatabaseCloudExadataInfrastructureResource()),
			},
		},
	}
}

func readDatabaseCloudExadataInfrastructures(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseCloudExadataInfrastructuresDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseCloudExadataInfrastructuresDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListCloudExadataInfrastructuresResponse
}

func (s *DatabaseCloudExadataInfrastructuresDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseCloudExadataInfrastructuresDataSourceCrud) Get() error {
	request := oci_database.ListCloudExadataInfrastructuresRequest{}

	if clusterPlacementGroupId, ok := s.D.GetOkExists("cluster_placement_group_id"); ok {
		tmp := clusterPlacementGroupId.(string)
		request.ClusterPlacementGroupId = &tmp
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
		request.LifecycleState = oci_database.CloudExadataInfrastructureSummaryLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListCloudExadataInfrastructures(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListCloudExadataInfrastructures(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseCloudExadataInfrastructuresDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseCloudExadataInfrastructuresDataSource-", DatabaseCloudExadataInfrastructuresDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		cloudExadataInfrastructure := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.ActivatedStorageCount != nil {
			cloudExadataInfrastructure["activated_storage_count"] = *r.ActivatedStorageCount
		}

		if r.AdditionalStorageCount != nil {
			cloudExadataInfrastructure["additional_storage_count"] = *r.AdditionalStorageCount
		}

		if r.AvailabilityDomain != nil {
			cloudExadataInfrastructure["availability_domain"] = *r.AvailabilityDomain
		}

		if r.AvailableStorageSizeInGBs != nil {
			cloudExadataInfrastructure["available_storage_size_in_gbs"] = *r.AvailableStorageSizeInGBs
		}

		if r.ClusterPlacementGroupId != nil {
			cloudExadataInfrastructure["cluster_placement_group_id"] = *r.ClusterPlacementGroupId
		}

		if r.ComputeCount != nil {
			cloudExadataInfrastructure["compute_count"] = *r.ComputeCount
		}

		if r.CpuCount != nil {
			cloudExadataInfrastructure["cpu_count"] = *r.CpuCount
		}

		customerContacts := []interface{}{}
		for _, item := range r.CustomerContacts {
			customerContacts = append(customerContacts, CustomerContactToMap(item))
		}
		cloudExadataInfrastructure["customer_contacts"] = customerContacts

		if r.DataStorageSizeInTBs != nil {
			cloudExadataInfrastructure["data_storage_size_in_tbs"] = *r.DataStorageSizeInTBs
		}

		if r.DbNodeStorageSizeInGBs != nil {
			cloudExadataInfrastructure["db_node_storage_size_in_gbs"] = *r.DbNodeStorageSizeInGBs
		}

		if r.DbServerVersion != nil {
			cloudExadataInfrastructure["db_server_version"] = *r.DbServerVersion
		}

		definedFileSystemConfigurations := []interface{}{}
		for _, item := range r.DefinedFileSystemConfigurations {
			definedFileSystemConfigurations = append(definedFileSystemConfigurations, DefinedFileSystemConfigurationToMap(item))
		}
		cloudExadataInfrastructure["defined_file_system_configurations"] = definedFileSystemConfigurations

		if r.DefinedTags != nil {
			cloudExadataInfrastructure["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			cloudExadataInfrastructure["display_name"] = *r.DisplayName
		}

		cloudExadataInfrastructure["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			cloudExadataInfrastructure["id"] = *r.Id
		}

		if r.LastMaintenanceRunId != nil {
			cloudExadataInfrastructure["last_maintenance_run_id"] = *r.LastMaintenanceRunId
		}

		if r.LifecycleDetails != nil {
			cloudExadataInfrastructure["lifecycle_details"] = *r.LifecycleDetails
		}

		if r.MaintenanceWindow != nil {
			cloudExadataInfrastructure["maintenance_window"] = []interface{}{MaintenanceWindowToMap(r.MaintenanceWindow)}
		} else {
			cloudExadataInfrastructure["maintenance_window"] = nil
		}

		if r.MaxCpuCount != nil {
			cloudExadataInfrastructure["max_cpu_count"] = *r.MaxCpuCount
		}

		if r.MaxDataStorageInTBs != nil {
			cloudExadataInfrastructure["max_data_storage_in_tbs"] = *r.MaxDataStorageInTBs
		}

		if r.MaxDbNodeStorageInGBs != nil {
			cloudExadataInfrastructure["max_db_node_storage_in_gbs"] = *r.MaxDbNodeStorageInGBs
		}

		if r.MaxMemoryInGBs != nil {
			cloudExadataInfrastructure["max_memory_in_gbs"] = *r.MaxMemoryInGBs
		}

		if r.MemorySizeInGBs != nil {
			cloudExadataInfrastructure["memory_size_in_gbs"] = *r.MemorySizeInGBs
		}

		if r.MonthlyDbServerVersion != nil {
			cloudExadataInfrastructure["monthly_db_server_version"] = *r.MonthlyDbServerVersion
		}

		if r.MonthlyStorageServerVersion != nil {
			cloudExadataInfrastructure["monthly_storage_server_version"] = *r.MonthlyStorageServerVersion
		}

		if r.NextMaintenanceRunId != nil {
			cloudExadataInfrastructure["next_maintenance_run_id"] = *r.NextMaintenanceRunId
		}

		if r.Shape != nil {
			cloudExadataInfrastructure["shape"] = *r.Shape
		}

		cloudExadataInfrastructure["state"] = r.LifecycleState

		if r.StorageCount != nil {
			cloudExadataInfrastructure["storage_count"] = *r.StorageCount
		}

		if r.StorageServerVersion != nil {
			cloudExadataInfrastructure["storage_server_version"] = *r.StorageServerVersion
		}

		if r.SystemTags != nil {
			cloudExadataInfrastructure["system_tags"] = tfresource.SystemTagsToMap(r.SystemTags)
		}

		if r.TimeCreated != nil {
			cloudExadataInfrastructure["time_created"] = r.TimeCreated.String()
		}

		if r.TotalStorageSizeInGBs != nil {
			cloudExadataInfrastructure["total_storage_size_in_gbs"] = *r.TotalStorageSizeInGBs
		}

		resources = append(resources, cloudExadataInfrastructure)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabaseCloudExadataInfrastructuresDataSource().Schema["cloud_exadata_infrastructures"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("cloud_exadata_infrastructures", resources); err != nil {
		return err
	}

	return nil
}
