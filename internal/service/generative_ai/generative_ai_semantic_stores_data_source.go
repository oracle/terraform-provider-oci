// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package generative_ai

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_generative_ai "github.com/oracle/oci-go-sdk/v65/generativeai"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func GenerativeAiSemanticStoresDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readGenerativeAiSemanticStoresWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"data_source_querying_connection_id": {
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
			"state": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"semantic_store_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(GenerativeAiSemanticStoreResource()),
						},
					},
				},
			},
		},
	}
}

func readGenerativeAiSemanticStoresWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &GenerativeAiSemanticStoresDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type GenerativeAiSemanticStoresDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_generative_ai.GenerativeAiClient
	Res    *oci_generative_ai.ListSemanticStoresResponse
}

func (s *GenerativeAiSemanticStoresDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *GenerativeAiSemanticStoresDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_generative_ai.ListSemanticStoresRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if dataSourceQueryingConnectionId, ok := s.D.GetOkExists("data_source_querying_connection_id"); ok {
		tmp := dataSourceQueryingConnectionId.(string)
		request.DataSourceQueryingConnectionId = &tmp
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
		lifecycleStateEnum := oci_generative_ai.ListSemanticStoresLifecycleStateEnum(state.(string))
		request.LifecycleState = []oci_generative_ai.ListSemanticStoresLifecycleStateEnum{lifecycleStateEnum}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "generative_ai")

	response, err := s.Client.ListSemanticStores(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSemanticStores(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *GenerativeAiSemanticStoresDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("GenerativeAiSemanticStoresDataSource-", GenerativeAiSemanticStoresDataSource(), s.D))
	resources := []map[string]interface{}{}
	semanticStore := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, SemanticStoreSummaryToMap(item))
	}
	semanticStore["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, GenerativeAiSemanticStoresDataSource().Schema["semantic_store_collection"].Elem.(*schema.Resource).Schema)
		semanticStore["items"] = items
	}

	resources = append(resources, semanticStore)
	if err := s.D.Set("semantic_store_collection", resources); err != nil {
		return err
	}

	return nil
}
