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

func GenerativeAiAgentAgentEndpointsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readGenerativeAiAgentAgentEndpoints,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"agent_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"agent_endpoint_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(GenerativeAiAgentAgentEndpointResource()),
						},
					},
				},
			},
		},
	}
}

func readGenerativeAiAgentAgentEndpoints(d *schema.ResourceData, m interface{}) error {
	sync := &GenerativeAiAgentAgentEndpointsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiAgentClient()

	return tfresource.ReadResource(sync)
}

type GenerativeAiAgentAgentEndpointsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_generative_ai_agent.GenerativeAiAgentClient
	Res    *oci_generative_ai_agent.ListAgentEndpointsResponse
}

func (s *GenerativeAiAgentAgentEndpointsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *GenerativeAiAgentAgentEndpointsDataSourceCrud) Get() error {
	request := oci_generative_ai_agent.ListAgentEndpointsRequest{}

	if agentId, ok := s.D.GetOkExists("agent_id"); ok {
		tmp := agentId.(string)
		request.AgentId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_generative_ai_agent.AgentEndpointLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "generative_ai_agent")

	response, err := s.Client.ListAgentEndpoints(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAgentEndpoints(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *GenerativeAiAgentAgentEndpointsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("GenerativeAiAgentAgentEndpointsDataSource-", GenerativeAiAgentAgentEndpointsDataSource(), s.D))
	resources := []map[string]interface{}{}
	agentEndpoint := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AgentEndpointSummaryToMap(item))
	}
	agentEndpoint["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, GenerativeAiAgentAgentEndpointsDataSource().Schema["agent_endpoint_collection"].Elem.(*schema.Resource).Schema)
		agentEndpoint["items"] = items
	}

	resources = append(resources, agentEndpoint)
	if err := s.D.Set("agent_endpoint_collection", resources); err != nil {
		return err
	}

	return nil
}
