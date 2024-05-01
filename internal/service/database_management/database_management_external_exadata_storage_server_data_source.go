// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseManagementExternalExadataStorageServerDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["external_exadata_storage_server_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DatabaseManagementExternalExadataStorageServerResource(), fieldMap, readSingularDatabaseManagementExternalExadataStorageServer)
}

func readSingularDatabaseManagementExternalExadataStorageServer(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalExadataStorageServerDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementExternalExadataStorageServerDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.GetExternalExadataStorageServerResponse
}

func (s *DatabaseManagementExternalExadataStorageServerDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementExternalExadataStorageServerDataSourceCrud) Get() error {
	request := oci_database_management.GetExternalExadataStorageServerRequest{}

	if externalExadataStorageServerId, ok := s.D.GetOkExists("external_exadata_storage_server_id"); ok {
		tmp := externalExadataStorageServerId.(string)
		request.ExternalExadataStorageServerId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.GetExternalExadataStorageServer(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseManagementExternalExadataStorageServerDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("additional_details", s.Res.AdditionalDetails)

	if s.Res.Connector != nil {
		s.D.Set("connector", []interface{}{ExternalExadataStorageConnectorSummaryToMap1(s.Res.Connector)})
	} else {
		s.D.Set("connector", nil)
	}

	if s.Res.CpuCount != nil {
		s.D.Set("cpu_count", *s.Res.CpuCount)
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

	if s.Res.InternalId != nil {
		s.D.Set("internal_id", *s.Res.InternalId)
	}

	if s.Res.IpAddress != nil {
		s.D.Set("ip_address", *s.Res.IpAddress)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.MakeModel != nil {
		s.D.Set("make_model", *s.Res.MakeModel)
	}

	if s.Res.MaxFlashDiskIOPS != nil {
		s.D.Set("max_flash_disk_iops", *s.Res.MaxFlashDiskIOPS)
	}

	if s.Res.MaxFlashDiskThroughput != nil {
		s.D.Set("max_flash_disk_throughput", *s.Res.MaxFlashDiskThroughput)
	}

	if s.Res.MaxHardDiskIOPS != nil {
		s.D.Set("max_hard_disk_iops", *s.Res.MaxHardDiskIOPS)
	}

	if s.Res.MaxHardDiskThroughput != nil {
		s.D.Set("max_hard_disk_throughput", *s.Res.MaxHardDiskThroughput)
	}

	if s.Res.MemoryGB != nil {
		s.D.Set("memory_gb", *s.Res.MemoryGB)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.Status != nil {
		s.D.Set("status", *s.Res.Status)
	}

	if s.Res.StorageGridId != nil {
		s.D.Set("storage_grid_id", *s.Res.StorageGridId)
	}

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.Version != nil {
		s.D.Set("version", *s.Res.Version)
	}

	return nil
}
