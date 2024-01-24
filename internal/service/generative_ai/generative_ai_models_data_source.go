// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package generative_ai

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_generative_ai "github.com/oracle/oci-go-sdk/v65/generativeai"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func GenerativeAiModelsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readGenerativeAiModels,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"capability": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
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
			"vendor": {
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
							Elem:     tfresource.GetDataSourceItemSchema(GenerativeAiModelResource()),
						},
					},
				},
			},
		},
	}
}

func readGenerativeAiModels(d *schema.ResourceData, m interface{}) error {
	sync := &GenerativeAiModelsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiClient()

	return tfresource.ReadResource(sync)
}

type GenerativeAiModelsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_generative_ai.GenerativeAiClient
	Res    *oci_generative_ai.ListModelsResponse
}

func (s *GenerativeAiModelsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *GenerativeAiModelsDataSourceCrud) Get() error {
	request := oci_generative_ai.ListModelsRequest{}

	if capability, ok := s.D.GetOkExists("capability"); ok {
		interfaces := capability.([]interface{})
		tmp := make([]oci_generative_ai.ModelCapabilityEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_generative_ai.ModelCapabilityEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("capability") {
			request.Capability = tmp
		}
	}

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
		request.LifecycleState = oci_generative_ai.ModelLifecycleStateEnum(state.(string))
	}

	if vendor, ok := s.D.GetOkExists("vendor"); ok {
		tmp := vendor.(string)
		request.Vendor = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "generative_ai")

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

func (s *GenerativeAiModelsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("GenerativeAiModelsDataSource-", GenerativeAiModelsDataSource(), s.D))
	resources := []map[string]interface{}{}
	model := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ModelSummaryToMap(item))
	}
	model["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, GenerativeAiModelsDataSource().Schema["model_collection"].Elem.(*schema.Resource).Schema)
		model["items"] = items
	}

	resources = append(resources, model)
	if err := s.D.Set("model_collection", resources); err != nil {
		return err
	}

	return nil
}
