// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dataintegration

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_dataintegration "github.com/oracle/oci-go-sdk/v65/dataintegration"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataintegrationWorkspaceFolderDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["folder_key"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}

	fieldMap["workspace_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DataintegrationWorkspaceFolderResource(), fieldMap, readSingularDataintegrationWorkspaceFolder)
}

func readSingularDataintegrationWorkspaceFolder(d *schema.ResourceData, m interface{}) error {
	sync := &DataintegrationWorkspaceFolderDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataIntegrationClient()

	return tfresource.ReadResource(sync)
}

type DataintegrationWorkspaceFolderDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dataintegration.DataIntegrationClient
	Res    *oci_dataintegration.GetFolderResponse
}

func (s *DataintegrationWorkspaceFolderDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataintegrationWorkspaceFolderDataSourceCrud) Get() error {
	request := oci_dataintegration.GetFolderRequest{}

	if folderKey, ok := s.D.GetOkExists("folder_key"); ok {
		tmp := folderKey.(string)
		request.FolderKey = &tmp
	}

	if workspaceId, ok := s.D.GetOkExists("workspace_id"); ok {
		tmp := workspaceId.(string)
		request.WorkspaceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "dataintegration")

	response, err := s.Client.GetFolder(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataintegrationWorkspaceFolderDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataintegrationWorkspaceFolderDataSource-", DataintegrationWorkspaceFolderDataSource(), s.D))

	if s.Res.CategoryName != nil {
		s.D.Set("category_name", *s.Res.CategoryName)
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.Identifier != nil {
		s.D.Set("identifier", *s.Res.Identifier)
	}

	if s.Res.Key != nil {
		s.D.Set("key", *s.Res.Key)
	}

	s.D.Set("key_map", s.Res.KeyMap)

	if s.Res.Metadata != nil {
		s.D.Set("metadata", []interface{}{DataintegrationFolderObjectMetadataToMap(s.Res.Metadata)})
	} else {
		s.D.Set("metadata", nil)
	}

	if s.Res.ModelType != nil {
		s.D.Set("model_type", *s.Res.ModelType)
	}

	if s.Res.ModelVersion != nil {
		s.D.Set("model_version", *s.Res.ModelVersion)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.ObjectStatus != nil {
		s.D.Set("object_status", *s.Res.ObjectStatus)
	}

	if s.Res.ObjectVersion != nil {
		s.D.Set("object_version", *s.Res.ObjectVersion)
	}

	if s.Res.ParentRef != nil {
		s.D.Set("parent_ref", []interface{}{DataintegrationFolderParentReferenceToMap(s.Res.ParentRef)})
	} else {
		s.D.Set("parent_ref", nil)
	}

	return nil
}
