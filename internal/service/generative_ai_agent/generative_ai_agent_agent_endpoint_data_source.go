// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package generative_ai_agent

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_generative_ai_agent "github.com/oracle/oci-go-sdk/v65/generativeaiagent"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func GenerativeAiAgentAgentEndpointDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["agent_endpoint_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(GenerativeAiAgentAgentEndpointResource(), fieldMap, readSingularGenerativeAiAgentAgentEndpoint)
}

func readSingularGenerativeAiAgentAgentEndpoint(d *schema.ResourceData, m interface{}) error {
	sync := &GenerativeAiAgentAgentEndpointDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiAgentClient()

	return tfresource.ReadResource(sync)
}

type GenerativeAiAgentAgentEndpointDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_generative_ai_agent.GenerativeAiAgentClient
	Res    *oci_generative_ai_agent.GetAgentEndpointResponse
}

func (s *GenerativeAiAgentAgentEndpointDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *GenerativeAiAgentAgentEndpointDataSourceCrud) Get() error {
	request := oci_generative_ai_agent.GetAgentEndpointRequest{}

	if agentEndpointId, ok := s.D.GetOkExists("agent_endpoint_id"); ok {
		tmp := agentEndpointId.(string)
		request.AgentEndpointId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "generative_ai_agent")

	response, err := s.Client.GetAgentEndpoint(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *GenerativeAiAgentAgentEndpointDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AgentId != nil {
		s.D.Set("agent_id", *s.Res.AgentId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ContentModerationConfig != nil {
		s.D.Set("content_moderation_config", []interface{}{ContentModerationConfigToMap(s.Res.ContentModerationConfig)})
	} else {
		s.D.Set("content_moderation_config", nil)
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

	if s.Res.SessionConfig != nil {
		s.D.Set("session_config", []interface{}{SessionConfigToMap(s.Res.SessionConfig)})
	} else {
		s.D.Set("session_config", nil)
	}

	if s.Res.ShouldEnableCitation != nil {
		s.D.Set("should_enable_citation", *s.Res.ShouldEnableCitation)
	}

	if s.Res.ShouldEnableSession != nil {
		s.D.Set("should_enable_session", *s.Res.ShouldEnableSession)
	}

	if s.Res.ShouldEnableTrace != nil {
		s.D.Set("should_enable_trace", *s.Res.ShouldEnableTrace)
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
