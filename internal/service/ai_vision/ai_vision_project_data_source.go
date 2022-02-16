// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ai_vision

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_ai_vision "github.com/oracle/oci-go-sdk/v58/aivision"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func AiVisionProjectDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["project_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(AiVisionProjectResource(), fieldMap, readSingularAiVisionProject)
}

func readSingularAiVisionProject(d *schema.ResourceData, m interface{}) error {
	sync := &AiVisionProjectDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AiServiceVisionClient()

	return tfresource.ReadResource(sync)
}

type AiVisionProjectDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_ai_vision.AIServiceVisionClient
	Res    *oci_ai_vision.GetProjectResponse
}

func (s *AiVisionProjectDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AiVisionProjectDataSourceCrud) Get() error {
	request := oci_ai_vision.GetProjectRequest{}

	if projectId, ok := s.D.GetOkExists("project_id"); ok {
		tmp := projectId.(string)
		request.ProjectId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "ai_vision")

	response, err := s.Client.GetProject(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *AiVisionProjectDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
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

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
