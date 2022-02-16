// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v58/database"
)

func DatabaseCloudAutonomousVmClusterDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["cloud_autonomous_vm_cluster_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DatabaseCloudAutonomousVmClusterResource(), fieldMap, readSingularDatabaseCloudAutonomousVmCluster)
}

func readSingularDatabaseCloudAutonomousVmCluster(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseCloudAutonomousVmClusterDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseCloudAutonomousVmClusterDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.GetCloudAutonomousVmClusterResponse
}

func (s *DatabaseCloudAutonomousVmClusterDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseCloudAutonomousVmClusterDataSourceCrud) Get() error {
	request := oci_database.GetCloudAutonomousVmClusterRequest{}

	if cloudAutonomousVmClusterId, ok := s.D.GetOkExists("cloud_autonomous_vm_cluster_id"); ok {
		tmp := cloudAutonomousVmClusterId.(string)
		request.CloudAutonomousVmClusterId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.GetCloudAutonomousVmCluster(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseCloudAutonomousVmClusterDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.CloudExadataInfrastructureId != nil {
		s.D.Set("cloud_exadata_infrastructure_id", *s.Res.CloudExadataInfrastructureId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CpuCoreCount != nil {
		s.D.Set("cpu_core_count", *s.Res.CpuCoreCount)
	}

	if s.Res.DataStorageSizeInGBs != nil {
		s.D.Set("data_storage_size_in_gb", *s.Res.DataStorageSizeInGBs)
	}

	if s.Res.DataStorageSizeInTBs != nil {
		s.D.Set("data_storage_size_in_tbs", *s.Res.DataStorageSizeInTBs)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.Domain != nil {
		s.D.Set("domain", *s.Res.Domain)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Hostname != nil {
		s.D.Set("hostname", *s.Res.Hostname)
	}

	if s.Res.LastMaintenanceRunId != nil {
		s.D.Set("last_maintenance_run_id", *s.Res.LastMaintenanceRunId)
	}

	if s.Res.LastUpdateHistoryEntryId != nil {
		s.D.Set("last_update_history_entry_id", *s.Res.LastUpdateHistoryEntryId)
	}

	s.D.Set("license_model", s.Res.LicenseModel)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.MemorySizeInGBs != nil {
		s.D.Set("memory_size_in_gbs", *s.Res.MemorySizeInGBs)
	}

	if s.Res.NextMaintenanceRunId != nil {
		s.D.Set("next_maintenance_run_id", *s.Res.NextMaintenanceRunId)
	}

	if s.Res.NodeCount != nil {
		s.D.Set("node_count", *s.Res.NodeCount)
	}

	s.D.Set("nsg_ids", s.Res.NsgIds)

	if s.Res.OcpuCount != nil {
		s.D.Set("ocpu_count", *s.Res.OcpuCount)
	}

	if s.Res.Shape != nil {
		s.D.Set("shape", *s.Res.Shape)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", *s.Res.SubnetId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
