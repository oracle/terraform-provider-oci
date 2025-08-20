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

func GenerativeAiGenerativeAiPrivateEndpointsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readGenerativeAiGenerativeAiPrivateEndpoints,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"generative_ai_private_endpoint_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(GenerativeAiGenerativeAiPrivateEndpointResource()),
						},
					},
				},
			},
		},
	}
}

func readGenerativeAiGenerativeAiPrivateEndpoints(d *schema.ResourceData, m interface{}) error {
	sync := &GenerativeAiGenerativeAiPrivateEndpointsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiClient()

	return tfresource.ReadResource(sync)
}

type GenerativeAiGenerativeAiPrivateEndpointsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_generative_ai.GenerativeAiClient
	Res    *oci_generative_ai.ListGenerativeAiPrivateEndpointsResponse
}

func (s *GenerativeAiGenerativeAiPrivateEndpointsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *GenerativeAiGenerativeAiPrivateEndpointsDataSourceCrud) Get() error {
	request := oci_generative_ai.ListGenerativeAiPrivateEndpointsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_generative_ai.GenerativeAiPrivateEndpointLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "generative_ai")

	response, err := s.Client.ListGenerativeAiPrivateEndpoints(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListGenerativeAiPrivateEndpoints(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *GenerativeAiGenerativeAiPrivateEndpointsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("GenerativeAiGenerativeAiPrivateEndpointsDataSource-", GenerativeAiGenerativeAiPrivateEndpointsDataSource(), s.D))
	resources := []map[string]interface{}{}
	generativeAiPrivateEndpoint := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, GenerativeAiPrivateEndpointSummaryToMap(item))
	}
	generativeAiPrivateEndpoint["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, GenerativeAiGenerativeAiPrivateEndpointsDataSource().Schema["generative_ai_private_endpoint_collection"].Elem.(*schema.Resource).Schema)
		generativeAiPrivateEndpoint["items"] = items
	}

	resources = append(resources, generativeAiPrivateEndpoint)
	if err := s.D.Set("generative_ai_private_endpoint_collection", resources); err != nil {
		return err
	}

	return nil
}
