// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ai_language

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_ai_language "github.com/oracle/oci-go-sdk/v65/ailanguage"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func AiLanguageEndpointDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(AiLanguageEndpointResource(), fieldMap, readSingularAiLanguageEndpoint)
}

func readSingularAiLanguageEndpoint(d *schema.ResourceData, m interface{}) error {
	sync := &AiLanguageEndpointDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AiServiceLanguageClient()

	return tfresource.ReadResource(sync)
}

type AiLanguageEndpointDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_ai_language.AIServiceLanguageClient
	Res    *oci_ai_language.GetEndpointResponse
}

func (s *AiLanguageEndpointDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AiLanguageEndpointDataSourceCrud) Get() error {
	request := oci_ai_language.GetEndpointRequest{}

	// if endpointId, ok := s.D.GetOkExists("endpoint_id"); ok {
	// 	tmp := endpointId.(string)
	// 	request.EndpointId = &tmp
	// }

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.EndpointId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "ai_language")

	response, err := s.Client.GetEndpoint(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *AiLanguageEndpointDataSourceCrud) SetData() error {
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

	if s.Res.FreeformTags != nil {
		s.D.Set("freeform_tags", s.Res.FreeformTags)
	}

	// s.D.Set("freeform_tags", s.Res.FreeformTags)
	// s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.InferenceUnits != nil {
		s.D.Set("inference_units", *s.Res.InferenceUnits)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.ModelId != nil {
		s.D.Set("model_id", *s.Res.ModelId)
	}

	if s.Res.ProjectId != nil {
		s.D.Set("project_id", *s.Res.ProjectId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", s.Res.SystemTags)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
