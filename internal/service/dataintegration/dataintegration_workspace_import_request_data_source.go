// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
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

func DataintegrationWorkspaceImportRequestDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["import_request_key"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["workspace_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DataintegrationWorkspaceImportRequestResource(), fieldMap, readSingularDataintegrationWorkspaceImportRequest)
}

func readSingularDataintegrationWorkspaceImportRequest(d *schema.ResourceData, m interface{}) error {
	sync := &DataintegrationWorkspaceImportRequestDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataIntegrationClient()

	return tfresource.ReadResource(sync)
}

type DataintegrationWorkspaceImportRequestDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dataintegration.DataIntegrationClient
	Res    *oci_dataintegration.GetImportRequestResponse
}

func (s *DataintegrationWorkspaceImportRequestDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataintegrationWorkspaceImportRequestDataSourceCrud) Get() error {
	request := oci_dataintegration.GetImportRequestRequest{}

	if importRequestKey, ok := s.D.GetOkExists("import_request_key"); ok {
		tmp := importRequestKey.(string)
		request.ImportRequestKey = &tmp
	}

	if workspaceId, ok := s.D.GetOkExists("workspace_id"); ok {
		tmp := workspaceId.(string)
		request.WorkspaceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "dataintegration")

	response, err := s.Client.GetImportRequest(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataintegrationWorkspaceImportRequestDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataintegrationWorkspaceImportRequestDataSource-", DataintegrationWorkspaceImportRequestDataSource(), s.D))

	if s.Res.BucketName != nil {
		s.D.Set("bucket", *s.Res.BucketName)
	}

	if s.Res.CreatedBy != nil {
		s.D.Set("created_by", *s.Res.CreatedBy)
	}

	s.D.Set("error_messages", s.Res.ErrorMessages)

	if s.Res.FileName != nil {
		s.D.Set("file_name", *s.Res.FileName)
	}

	if s.Res.ImportConflictResolution != nil {
		s.D.Set("import_conflict_resolution", []interface{}{ImportConflictResolutionToMap(s.Res.ImportConflictResolution)})
	} else {
		s.D.Set("import_conflict_resolution", nil)
	}

	importedObjects := []interface{}{}
	for _, item := range s.Res.ImportedObjects {
		importedObjects = append(importedObjects, ImportObjectMetadataSummaryToMap(item))
	}
	s.D.Set("imported_objects", importedObjects)

	if s.Res.Key != nil {
		s.D.Set("key", *s.Res.Key)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.ObjectKeyForImport != nil {
		s.D.Set("object_key_for_import", *s.Res.ObjectKeyForImport)
	}

	if s.Res.ObjectStorageRegion != nil {
		s.D.Set("object_storage_region", *s.Res.ObjectStorageRegion)
	}

	if s.Res.ObjectStorageTenancyId != nil {
		s.D.Set("object_storage_tenancy_id", *s.Res.ObjectStorageTenancyId)
	}

	s.D.Set("status", s.Res.Status)

	if s.Res.TimeEndedInMillis != nil {
		s.D.Set("time_ended_in_millis", strconv.FormatInt(*s.Res.TimeEndedInMillis, 10))
	}

	if s.Res.TimeStartedInMillis != nil {
		s.D.Set("time_started_in_millis", strconv.FormatInt(*s.Res.TimeStartedInMillis, 10))
	}

	if s.Res.TotalImportedObjectCount != nil {
		s.D.Set("total_imported_object_count", *s.Res.TotalImportedObjectCount)
	}

	return nil
}
