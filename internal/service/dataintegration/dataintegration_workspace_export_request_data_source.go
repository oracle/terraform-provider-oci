// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dataintegration

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_dataintegration "github.com/oracle/oci-go-sdk/v65/dataintegration"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataintegrationWorkspaceExportRequestDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["export_request_key"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["workspace_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DataintegrationWorkspaceExportRequestResource(), fieldMap, readSingularDataintegrationWorkspaceExportRequest)
}

func readSingularDataintegrationWorkspaceExportRequest(d *schema.ResourceData, m interface{}) error {
	sync := &DataintegrationWorkspaceExportRequestDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataIntegrationClient()

	return tfresource.ReadResource(sync)
}

type DataintegrationWorkspaceExportRequestDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dataintegration.DataIntegrationClient
	Res    *oci_dataintegration.GetExportRequestResponse
}

func (s *DataintegrationWorkspaceExportRequestDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataintegrationWorkspaceExportRequestDataSourceCrud) Get() error {
	request := oci_dataintegration.GetExportRequestRequest{}

	if exportRequestKey, ok := s.D.GetOkExists("export_request_key"); ok {
		tmp := exportRequestKey.(string)
		request.ExportRequestKey = &tmp
	}

	if workspaceId, ok := s.D.GetOkExists("workspace_id"); ok {
		tmp := workspaceId.(string)
		request.WorkspaceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "dataintegration")

	response, err := s.Client.GetExportRequest(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataintegrationWorkspaceExportRequestDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataintegrationWorkspaceExportRequestDataSource-", DataintegrationWorkspaceExportRequestDataSource(), s.D))

	if s.Res.AreReferencesIncluded != nil {
		s.D.Set("are_references_included", *s.Res.AreReferencesIncluded)
	}

	if s.Res.BucketName != nil {
		s.D.Set("bucket", *s.Res.BucketName)
	}

	if s.Res.CreatedBy != nil {
		s.D.Set("created_by", *s.Res.CreatedBy)
	}

	s.D.Set("error_messages", s.Res.ErrorMessages)

	exportedItems := []interface{}{}
	for _, item := range s.Res.ExportedItems {
		exportedItems = append(exportedItems, ExportObjectMetadataSummaryToMap(item))
	}
	s.D.Set("exported_items", exportedItems)

	if s.Res.FileName != nil {
		s.D.Set("file_name", *s.Res.FileName)
	}

	s.D.Set("filters", s.Res.Filters)

	if s.Res.IsObjectOverwriteEnabled != nil {
		s.D.Set("is_object_overwrite_enabled", *s.Res.IsObjectOverwriteEnabled)
	}

	if s.Res.Key != nil {
		s.D.Set("key", *s.Res.Key)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("object_keys", s.Res.ObjectKeys)

	if s.Res.ObjectStorageRegion != nil {
		s.D.Set("object_storage_region", *s.Res.ObjectStorageRegion)
	}

	if s.Res.ObjectStorageTenancyId != nil {
		s.D.Set("object_storage_tenancy_id", *s.Res.ObjectStorageTenancyId)
	}

	if s.Res.ReferencedItems != nil {
		s.D.Set("referenced_items", *s.Res.ReferencedItems)
	}

	s.D.Set("status", s.Res.Status)

	if s.Res.TimeEndedInMillis != nil {
		s.D.Set("time_ended_in_millis", strconv.FormatInt(*s.Res.TimeEndedInMillis, 10))
	}

	if s.Res.TimeStartedInMillis != nil {
		s.D.Set("time_started_in_millis", strconv.FormatInt(*s.Res.TimeStartedInMillis, 10))
	}

	if s.Res.TotalExportedObjectCount != nil {
		s.D.Set("total_exported_object_count", *s.Res.TotalExportedObjectCount)
	}

	return nil
}
