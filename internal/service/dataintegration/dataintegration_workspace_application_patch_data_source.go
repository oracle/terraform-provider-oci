// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dataintegration

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_dataintegration "github.com/oracle/oci-go-sdk/v65/dataintegration"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataintegrationWorkspaceApplicationPatchDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["application_key"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["patch_key"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["workspace_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DataintegrationWorkspaceApplicationPatchResource(), fieldMap, readSingularDataintegrationWorkspaceApplicationPatch)
}

func readSingularDataintegrationWorkspaceApplicationPatch(d *schema.ResourceData, m interface{}) error {
	sync := &DataintegrationWorkspaceApplicationPatchDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataIntegrationClient()

	return tfresource.ReadResource(sync)
}

type DataintegrationWorkspaceApplicationPatchDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dataintegration.DataIntegrationClient
	Res    *oci_dataintegration.GetPatchResponse
}

func (s *DataintegrationWorkspaceApplicationPatchDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataintegrationWorkspaceApplicationPatchDataSourceCrud) Get() error {
	request := oci_dataintegration.GetPatchRequest{}

	if applicationKey, ok := s.D.GetOkExists("application_key"); ok {
		tmp := applicationKey.(string)
		request.ApplicationKey = &tmp
	}

	if patchKey, ok := s.D.GetOkExists("patch_key"); ok {
		tmp := patchKey.(string)
		request.PatchKey = &tmp
	}

	if workspaceId, ok := s.D.GetOkExists("workspace_id"); ok {
		tmp := workspaceId.(string)
		request.WorkspaceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "dataintegration")

	response, err := s.Client.GetPatch(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataintegrationWorkspaceApplicationPatchDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataintegrationWorkspaceApplicationPatchDataSource-", DataintegrationWorkspaceApplicationPatchDataSource(), s.D))

	if s.Res.ApplicationVersion != nil {
		s.D.Set("application_version", *s.Res.ApplicationVersion)
	}

	dependentObjectMetadata := []interface{}{}
	for _, item := range s.Res.DependentObjectMetadata {
		dependentObjectMetadata = append(dependentObjectMetadata, PatchObjectMetadataToMap(item))
	}
	s.D.Set("dependent_object_metadata", dependentObjectMetadata)

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("error_messages", s.Res.ErrorMessages)
	s.D.Set("error_messages", s.Res.ErrorMessages)

	if s.Res.Identifier != nil {
		s.D.Set("identifier", *s.Res.Identifier)
	}

	if s.Res.Key != nil {
		s.D.Set("key", *s.Res.Key)
	}

	s.D.Set("key_map", s.Res.KeyMap)
	s.D.Set("key_map", s.Res.KeyMap)

	if s.Res.Metadata != nil {
		s.D.Set("metadata", []interface{}{ObjectMetadataToMap(s.Res.Metadata)})
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
		s.D.Set("parent_ref", []interface{}{ParentReferenceToMap(s.Res.ParentRef)})
	} else {
		s.D.Set("parent_ref", nil)
	}

	patchObjectMetadata := []interface{}{}
	for _, item := range s.Res.PatchObjectMetadata {
		patchObjectMetadata = append(patchObjectMetadata, PatchObjectMetadataToMap(item))
	}
	s.D.Set("patch_object_metadata", patchObjectMetadata)

	s.D.Set("patch_status", s.Res.PatchStatus)

	s.D.Set("patch_type", s.Res.PatchType)

	if s.Res.TimePatched != nil {
		s.D.Set("time_patched", s.Res.TimePatched.String())
	}

	return nil
}
