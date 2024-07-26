// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package generative_ai

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_generative_ai "github.com/oracle/oci-go-sdk/v65/generativeai"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func GenerativeAiEndpointDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["endpoint_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(GenerativeAiEndpointResource(), fieldMap, readSingularGenerativeAiEndpoint)
}

func readSingularGenerativeAiEndpoint(d *schema.ResourceData, m interface{}) error {
	sync := &GenerativeAiEndpointDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiClient()

	return tfresource.ReadResource(sync)
}

type GenerativeAiEndpointDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_generative_ai.GenerativeAiClient
	Res    *oci_generative_ai.GetEndpointResponse
}

func (s *GenerativeAiEndpointDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *GenerativeAiEndpointDataSourceCrud) Get() error {
	request := oci_generative_ai.GetEndpointRequest{}

	if endpointId, ok := s.D.GetOkExists("endpoint_id"); ok {
		tmp := endpointId.(string)
		request.EndpointId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "generative_ai")

	response, err := s.Client.GetEndpoint(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *GenerativeAiEndpointDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ContentModerationConfig != nil {
		s.D.Set("content_moderation_config", []interface{}{ContentModerationConfigToMap(s.Res.ContentModerationConfig)})
	} else {
		s.D.Set("content_moderation_config", nil)
	}

	if s.Res.DedicatedAiClusterId != nil {
		s.D.Set("dedicated_ai_cluster_id", *s.Res.DedicatedAiClusterId)
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

	if s.Res.ModelId != nil {
		s.D.Set("model_id", *s.Res.ModelId)
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
