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

func GenerativeAiAgentToolDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["tool_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(GenerativeAiAgentToolResource(), fieldMap, readSingularGenerativeAiAgentTool)
}

func readSingularGenerativeAiAgentTool(d *schema.ResourceData, m interface{}) error {
	sync := &GenerativeAiAgentToolDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiAgentClient()

	return tfresource.ReadResource(sync)
}

type GenerativeAiAgentToolDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_generative_ai_agent.GenerativeAiAgentClient
	Res    *oci_generative_ai_agent.GetToolResponse
}

func (s *GenerativeAiAgentToolDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *GenerativeAiAgentToolDataSourceCrud) Get() error {
	request := oci_generative_ai_agent.GetToolRequest{}

	if toolId, ok := s.D.GetOkExists("tool_id"); ok {
		tmp := toolId.(string)
		request.ToolId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "generative_ai_agent")

	response, err := s.Client.GetTool(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *GenerativeAiAgentToolDataSourceCrud) SetData() error {
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

	s.D.Set("metadata", s.Res.Metadata)

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

	if s.Res.ToolConfig != nil {
		toolConfigArray := []interface{}{}
		if toolConfigMap := ToolConfigToMap(&s.Res.ToolConfig); toolConfigMap != nil {
			toolConfigArray = append(toolConfigArray, toolConfigMap)
		}
		s.D.Set("tool_config", toolConfigArray)
	} else {
		s.D.Set("tool_config", nil)
	}

	return nil
}
