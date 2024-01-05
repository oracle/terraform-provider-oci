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

func DataintegrationWorkspaceApplicationDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["application_key"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["workspace_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DataintegrationWorkspaceApplicationResource(), fieldMap, readSingularDataintegrationWorkspaceApplication)
}

func readSingularDataintegrationWorkspaceApplication(d *schema.ResourceData, m interface{}) error {
	sync := &DataintegrationWorkspaceApplicationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataIntegrationClient()

	return tfresource.ReadResource(sync)
}

type DataintegrationWorkspaceApplicationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dataintegration.DataIntegrationClient
	Res    *oci_dataintegration.GetApplicationResponse
}

func (s *DataintegrationWorkspaceApplicationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataintegrationWorkspaceApplicationDataSourceCrud) Get() error {
	request := oci_dataintegration.GetApplicationRequest{}

	if applicationKey, ok := s.D.GetOkExists("application_key"); ok {
		tmp := applicationKey.(string)
		request.ApplicationKey = &tmp
	}

	if workspaceId, ok := s.D.GetOkExists("workspace_id"); ok {
		tmp := workspaceId.(string)
		request.WorkspaceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "dataintegration")

	response, err := s.Client.GetApplication(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataintegrationWorkspaceApplicationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.ApplicationVersion != nil {
		s.D.Set("application_version", *s.Res.ApplicationVersion)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	dependentObjectMetadata := []interface{}{}
	for _, item := range s.Res.DependentObjectMetadata {
		dependentObjectMetadata = append(dependentObjectMetadata, PatchObjectMetadataToMap(item))
	}
	s.D.Set("dependent_object_metadata", dependentObjectMetadata)

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Identifier != nil {
		s.D.Set("identifier", *s.Res.Identifier)
	}

	if s.Res.Key != nil {
		s.D.Set("key", *s.Res.Key)
	}

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

	publishedObjectMetadataList := []interface{}{}
	for _, value := range s.Res.PublishedObjectMetadata {
		publishedObjectMetadataList = append(publishedObjectMetadataList, value)
	}
	s.D.Set("published_object_metadata", publishedObjectMetadataList)

	if s.Res.SourceApplicationInfo != nil {
		s.D.Set("source_application_info", []interface{}{SourceApplicationInfoToMap(s.Res.SourceApplicationInfo)})
	} else {
		s.D.Set("source_application_info", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimePatched != nil {
		s.D.Set("time_patched", s.Res.TimePatched.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
