// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ai_data_platform

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_ai_data_platform "github.com/oracle/oci-go-sdk/v65/aidataplatform"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func AiDataPlatformAiDataPlatformsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readAiDataPlatformAiDataPlatformsWithContext,
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
			"exclude_lifecycle_state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"include_legacy": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ai_data_platform_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(AiDataPlatformAiDataPlatformResource()),
						},
					},
				},
			},
		},
	}
}

func readAiDataPlatformAiDataPlatformsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &AiDataPlatformAiDataPlatformsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AiDataPlatformClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type AiDataPlatformAiDataPlatformsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_ai_data_platform.AiDataPlatformClient
	Res    *oci_ai_data_platform.ListAiDataPlatformsResponse
}

func (s *AiDataPlatformAiDataPlatformsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AiDataPlatformAiDataPlatformsDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_ai_data_platform.ListAiDataPlatformsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if excludeLifecycleState, ok := s.D.GetOkExists("exclude_lifecycle_state"); ok {
		request.ExcludeLifecycleState = oci_ai_data_platform.AiDataPlatformLifecycleStateEnum(excludeLifecycleState.(string))
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if includeLegacy, ok := s.D.GetOkExists("include_legacy"); ok {
		tmp := includeLegacy.(string)
		request.IncludeLegacy = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_ai_data_platform.AiDataPlatformLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "ai_data_platform")

	response, err := s.Client.ListAiDataPlatforms(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAiDataPlatforms(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *AiDataPlatformAiDataPlatformsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("AiDataPlatformAiDataPlatformsDataSource-", AiDataPlatformAiDataPlatformsDataSource(), s.D))
	resources := []map[string]interface{}{}
	aiDataPlatform := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AiDataPlatformSummaryToMap(item))
	}
	aiDataPlatform["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, AiDataPlatformAiDataPlatformsDataSource().Schema["ai_data_platform_collection"].Elem.(*schema.Resource).Schema)
		aiDataPlatform["items"] = items
	}

	resources = append(resources, aiDataPlatform)
	if err := s.D.Set("ai_data_platform_collection", resources); err != nil {
		return err
	}

	return nil
}
