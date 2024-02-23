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

func DatabaseCloudExadataInfrastructureDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["cloud_exadata_infrastructure_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DatabaseCloudExadataInfrastructureResource(), fieldMap, readSingularDatabaseCloudExadataInfrastructure)
}

func readSingularDatabaseCloudExadataInfrastructure(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseCloudExadataInfrastructureDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseCloudExadataInfrastructureDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.GetCloudExadataInfrastructureResponse
}

func (s *DatabaseCloudExadataInfrastructureDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseCloudExadataInfrastructureDataSourceCrud) Get() error {
	request := oci_database.GetCloudExadataInfrastructureRequest{}

	if cloudExadataInfrastructureId, ok := s.D.GetOkExists("cloud_exadata_infrastructure_id"); ok {
		tmp := cloudExadataInfrastructureId.(string)
		request.CloudExadataInfrastructureId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.GetCloudExadataInfrastructure(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseCloudExadataInfrastructureDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.ActivatedStorageCount != nil {
		s.D.Set("activated_storage_count", *s.Res.ActivatedStorageCount)
	}

	if s.Res.AdditionalStorageCount != nil {
		s.D.Set("additional_storage_count", *s.Res.AdditionalStorageCount)
	}

	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.AvailableStorageSizeInGBs != nil {
		s.D.Set("available_storage_size_in_gbs", *s.Res.AvailableStorageSizeInGBs)
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

	if s.Res.CpuCount != nil {
		s.D.Set("cpu_count", *s.Res.CpuCount)
	}

	customerContacts := []interface{}{}
	for _, item := range s.Res.CustomerContacts {
		customerContacts = append(customerContacts, CustomerContactToMap(item))
	}
	s.D.Set("customer_contacts", customerContacts)

	if s.Res.DataStorageSizeInTBs != nil {
		s.D.Set("data_storage_size_in_tbs", *s.Res.DataStorageSizeInTBs)
	}

	if s.Res.DbNodeStorageSizeInGBs != nil {
		s.D.Set("db_node_storage_size_in_gbs", *s.Res.DbNodeStorageSizeInGBs)
	}

	if s.Res.DbServerVersion != nil {
		s.D.Set("db_server_version", *s.Res.DbServerVersion)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

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

	if s.Res.MaxCpuCount != nil {
		s.D.Set("max_cpu_count", *s.Res.MaxCpuCount)
	}

	if s.Res.MaxDataStorageInTBs != nil {
		s.D.Set("max_data_storage_in_tbs", *s.Res.MaxDataStorageInTBs)
	}

	if s.Res.MaxDbNodeStorageInGBs != nil {
		s.D.Set("max_db_node_storage_in_gbs", *s.Res.MaxDbNodeStorageInGBs)
	}

	if s.Res.MaxMemoryInGBs != nil {
		s.D.Set("max_memory_in_gbs", *s.Res.MaxMemoryInGBs)
	}

	if s.Res.MemorySizeInGBs != nil {
		s.D.Set("memory_size_in_gbs", *s.Res.MemorySizeInGBs)
	}

	if s.Res.MonthlyDbServerVersion != nil {
		s.D.Set("monthly_db_server_version", *s.Res.MonthlyDbServerVersion)
	}

	if s.Res.MonthlyStorageServerVersion != nil {
		s.D.Set("monthly_storage_server_version", *s.Res.MonthlyStorageServerVersion)
	}

	if s.Res.NextMaintenanceRunId != nil {
		s.D.Set("next_maintenance_run_id", *s.Res.NextMaintenanceRunId)
	}

	if s.Res.Shape != nil {
		s.D.Set("shape", *s.Res.Shape)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.StorageCount != nil {
		s.D.Set("storage_count", *s.Res.StorageCount)
	}

	if s.Res.StorageServerVersion != nil {
		s.D.Set("storage_server_version", *s.Res.StorageServerVersion)
	}

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TotalStorageSizeInGBs != nil {
		s.D.Set("total_storage_size_in_gbs", *s.Res.TotalStorageSizeInGBs)
	}

	return nil
}
