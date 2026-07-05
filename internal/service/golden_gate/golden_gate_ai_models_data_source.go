// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package golden_gate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_golden_gate "github.com/oracle/oci-go-sdk/v65/goldengate"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func GoldenGateAiModelsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readGoldenGateAiModelsWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"provider_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"region": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tenancy_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ai_model_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"key": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"provider_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func readGoldenGateAiModelsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &GoldenGateAiModelsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GoldenGateClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type GoldenGateAiModelsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_golden_gate.GoldenGateClient
	Res    *oci_golden_gate.ListAiModelsResponse
}

func (s *GoldenGateAiModelsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *GoldenGateAiModelsDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_golden_gate.ListAiModelsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if providerType, ok := s.D.GetOkExists("provider_type"); ok {
		request.ProviderType = oci_golden_gate.AiProviderSummaryProviderTypeEnum(providerType.(string))
	}

	if region, ok := s.D.GetOkExists("region"); ok {
		tmp := region.(string)
		request.Region = &tmp
	}

	if tenancyId, ok := s.D.GetOkExists("tenancy_id"); ok {
		tmp := tenancyId.(string)
		request.TenancyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "golden_gate")

	response, err := s.Client.ListAiModels(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *GoldenGateAiModelsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("GoldenGateAiModelsDataSource-", GoldenGateAiModelsDataSource(), s.D))
	resources := []map[string]interface{}{}
	aiModel := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AiModelSummaryToMap(item))
	}
	aiModel["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, GoldenGateAiModelsDataSource().Schema["ai_model_collection"].Elem.(*schema.Resource).Schema)
		aiModel["items"] = items
	}

	resources = append(resources, aiModel)
	if err := s.D.Set("ai_model_collection", resources); err != nil {
		return err
	}

	return nil
}

func AiModelSummaryToMap(obj oci_golden_gate.AiModelSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	result["provider_type"] = string(obj.ProviderType)

	return result
}
