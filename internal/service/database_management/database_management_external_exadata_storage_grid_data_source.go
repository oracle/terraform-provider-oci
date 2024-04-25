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

func DatabaseManagementExternalExadataStorageGridDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["external_exadata_storage_grid_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DatabaseManagementExternalExadataStorageGridResource(), fieldMap, readSingularDatabaseManagementExternalExadataStorageGrid)
}

func readSingularDatabaseManagementExternalExadataStorageGrid(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalExadataStorageGridDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementExternalExadataStorageGridDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.GetExternalExadataStorageGridResponse
}

func (s *DatabaseManagementExternalExadataStorageGridDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementExternalExadataStorageGridDataSourceCrud) Get() error {
	request := oci_database_management.GetExternalExadataStorageGridRequest{}

	if externalExadataStorageGridId, ok := s.D.GetOkExists("external_exadata_storage_grid_id"); ok {
		tmp := externalExadataStorageGridId.(string)
		request.ExternalExadataStorageGridId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.GetExternalExadataStorageGrid(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseManagementExternalExadataStorageGridDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("additional_details", s.Res.AdditionalDetails)

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

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.ServerCount != nil {
		s.D.Set("server_count", *s.Res.ServerCount)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.Status != nil {
		s.D.Set("status", *s.Res.Status)
	}

	storageServers := []interface{}{}
	for _, item := range s.Res.StorageServers {
		storageServers = append(storageServers, ExternalExadataStorageServerSummaryToMap(item))
	}
	s.D.Set("storage_servers", storageServers)

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

func ExternalExadataStorageServerSummaryToMap(obj oci_database_management.ExternalExadataStorageServerSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["additional_details"] = obj.AdditionalDetails

	if obj.ConnectorId != nil {
		result["connector_id"] = string(*obj.ConnectorId)
	}

	if obj.CpuCount != nil {
		result["cpu_count"] = float32(*obj.CpuCount)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.InternalId != nil {
		result["internal_id"] = string(*obj.InternalId)
	}

	if obj.IpAddress != nil {
		result["ip_address"] = string(*obj.IpAddress)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.MakeModel != nil {
		result["make_model"] = string(*obj.MakeModel)
	}

	if obj.MaxFlashDiskIOPS != nil {
		result["max_flash_disk_iops"] = int(*obj.MaxFlashDiskIOPS)
	}

	if obj.MaxFlashDiskThroughput != nil {
		result["max_flash_disk_throughput"] = int(*obj.MaxFlashDiskThroughput)
	}

	if obj.MaxHardDiskIOPS != nil {
		result["max_hard_disk_iops"] = int(*obj.MaxHardDiskIOPS)
	}

	if obj.MaxHardDiskThroughput != nil {
		result["max_hard_disk_throughput"] = int(*obj.MaxHardDiskThroughput)
	}

	if obj.MemoryGB != nil {
		result["memory_gb"] = float64(*obj.MemoryGB)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.Status != nil {
		result["status"] = string(*obj.Status)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	if obj.Version != nil {
		result["version"] = string(*obj.Version)
	}

	return result
}
