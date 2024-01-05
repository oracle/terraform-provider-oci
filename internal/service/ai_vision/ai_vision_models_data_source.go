// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ai_vision

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_ai_vision "github.com/oracle/oci-go-sdk/v65/aivision"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func AiVisionModelsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readAiVisionModels,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"project_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"model_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(AiVisionModelResource()),
						},
					},
				},
			},
		},
	}
}

func readAiVisionModels(d *schema.ResourceData, m interface{}) error {
	sync := &AiVisionModelsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AiServiceVisionClient()

	return tfresource.ReadResource(sync)
}

type AiVisionModelsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_ai_vision.AIServiceVisionClient
	Res    *oci_ai_vision.ListModelsResponse
}

func (s *AiVisionModelsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AiVisionModelsDataSourceCrud) Get() error {
	request := oci_ai_vision.ListModelsRequest{}

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

	if projectId, ok := s.D.GetOkExists("project_id"); ok {
		tmp := projectId.(string)
		request.ProjectId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_ai_vision.ModelLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "ai_vision")

	response, err := s.Client.ListModels(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListModels(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *AiVisionModelsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("AiVisionModelsDataSource-", AiVisionModelsDataSource(), s.D))
	resources := []map[string]interface{}{}
	model := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ModelSummaryToMap(item))
	}
	model["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, AiVisionModelsDataSource().Schema["model_collection"].Elem.(*schema.Resource).Schema)
		model["items"] = items
	}

	resources = append(resources, model)
	if err := s.D.Set("model_collection", resources); err != nil {
		return err
	}

	return nil
}
