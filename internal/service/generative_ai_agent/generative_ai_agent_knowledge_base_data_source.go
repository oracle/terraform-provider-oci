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

func GenerativeAiAgentKnowledgeBaseDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["knowledge_base_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(GenerativeAiAgentKnowledgeBaseResource(), fieldMap, readSingularGenerativeAiAgentKnowledgeBase)
}

func readSingularGenerativeAiAgentKnowledgeBase(d *schema.ResourceData, m interface{}) error {
	sync := &GenerativeAiAgentKnowledgeBaseDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiAgentClient()

	return tfresource.ReadResource(sync)
}

type GenerativeAiAgentKnowledgeBaseDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_generative_ai_agent.GenerativeAiAgentClient
	Res    *oci_generative_ai_agent.GetKnowledgeBaseResponse
}

func (s *GenerativeAiAgentKnowledgeBaseDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *GenerativeAiAgentKnowledgeBaseDataSourceCrud) Get() error {
	request := oci_generative_ai_agent.GetKnowledgeBaseRequest{}

	if knowledgeBaseId, ok := s.D.GetOkExists("knowledge_base_id"); ok {
		tmp := knowledgeBaseId.(string)
		request.KnowledgeBaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "generative_ai_agent")

	response, err := s.Client.GetKnowledgeBase(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *GenerativeAiAgentKnowledgeBaseDataSourceCrud) SetData() error {
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

	if s.Res.IndexConfig != nil {
		indexConfigArray := []interface{}{}
		if indexConfigMap := IndexConfigToMap(&s.Res.IndexConfig); indexConfigMap != nil {
			indexConfigArray = append(indexConfigArray, indexConfigMap)
		}
		s.D.Set("index_config", indexConfigArray)
	} else {
		s.D.Set("index_config", nil)
	}

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
