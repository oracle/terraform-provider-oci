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

func AiLanguageModelsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readAiLanguageModels,
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
							Elem:     tfresource.GetDataSourceItemSchema(AiLanguageModelResource()),
						},
					},
				},
			},
		},
	}
}

func readAiLanguageModels(d *schema.ResourceData, m interface{}) error {
	sync := &AiLanguageModelsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AiServiceLanguageClient()

	return tfresource.ReadResource(sync)
}

type AiLanguageModelsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_ai_language.AIServiceLanguageClient
	Res    *oci_ai_language.ListModelsResponse
}

func (s *AiLanguageModelsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AiLanguageModelsDataSourceCrud) Get() error {
	request := oci_ai_language.ListModelsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	// if modelId, ok := s.D.GetOkExists("id"); ok {
	// 	tmp := modelId.(string)
	// 	request.ModelId = &tmp
	// }

	if projectId, ok := s.D.GetOkExists("project_id"); ok {
		tmp := projectId.(string)
		request.ProjectId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_ai_language.ModelLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "ai_language")

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

func (s *AiLanguageModelsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("AiLanguageModelsDataSource-", AiLanguageModelsDataSource(), s.D))
	resources := []map[string]interface{}{}
	model := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ModelSummaryToMap(item))
	}
	model["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, AiLanguageModelsDataSource().Schema["model_collection"].Elem.(*schema.Resource).Schema)
		model["items"] = items
	}

	resources = append(resources, model)
	if err := s.D.Set("model_collection", resources); err != nil {
		return err
	}

	return nil
}
