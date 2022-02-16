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

func DatabaseAutonomousVmClusterDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["autonomous_vm_cluster_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DatabaseAutonomousVmClusterResource(), fieldMap, readSingularDatabaseAutonomousVmCluster)
}

func readSingularDatabaseAutonomousVmCluster(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousVmClusterDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseAutonomousVmClusterDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.GetAutonomousVmClusterResponse
}

func (s *DatabaseAutonomousVmClusterDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseAutonomousVmClusterDataSourceCrud) Get() error {
	request := oci_database.GetAutonomousVmClusterRequest{}

	if autonomousVmClusterId, ok := s.D.GetOkExists("autonomous_vm_cluster_id"); ok {
		tmp := autonomousVmClusterId.(string)
		request.AutonomousVmClusterId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.GetAutonomousVmCluster(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseAutonomousVmClusterDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AvailableCpus != nil {
		s.D.Set("available_cpus", *s.Res.AvailableCpus)
	}

	if s.Res.AvailableDataStorageSizeInTBs != nil {
		s.D.Set("available_data_storage_size_in_tbs", *s.Res.AvailableDataStorageSizeInTBs)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CpusEnabled != nil {
		s.D.Set("cpus_enabled", *s.Res.CpusEnabled)
	}

	if s.Res.DataStorageSizeInGBs != nil {
		s.D.Set("data_storage_size_in_gb", *s.Res.DataStorageSizeInGBs)
	}

	if s.Res.DataStorageSizeInTBs != nil {
		s.D.Set("data_storage_size_in_tbs", *s.Res.DataStorageSizeInTBs)
	}

	if s.Res.DbNodeStorageSizeInGBs != nil {
		s.D.Set("db_node_storage_size_in_gbs", *s.Res.DbNodeStorageSizeInGBs)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.ExadataInfrastructureId != nil {
		s.D.Set("exadata_infrastructure_id", *s.Res.ExadataInfrastructureId)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsLocalBackupEnabled != nil {
		s.D.Set("is_local_backup_enabled", *s.Res.IsLocalBackupEnabled)
	}

	s.D.Set("license_model", s.Res.LicenseModel)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.MemorySizeInGBs != nil {
		s.D.Set("memory_size_in_gbs", *s.Res.MemorySizeInGBs)
	}

	if s.Res.OcpusEnabled != nil {
		s.D.Set("ocpus_enabled", *s.Res.OcpusEnabled)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeZone != nil {
		s.D.Set("time_zone", *s.Res.TimeZone)
	}

	if s.Res.VmClusterNetworkId != nil {
		s.D.Set("vm_cluster_network_id", *s.Res.VmClusterNetworkId)
	}

	return nil
}
