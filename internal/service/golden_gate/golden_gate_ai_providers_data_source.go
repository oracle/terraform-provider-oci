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

func GoldenGateAiProvidersDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readGoldenGateAiProvidersWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ai_provider_collection": {
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
									"auth_type": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"default_base_url": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"models": {
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

func readGoldenGateAiProvidersWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &GoldenGateAiProvidersDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GoldenGateClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type GoldenGateAiProvidersDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_golden_gate.GoldenGateClient
	Res    *oci_golden_gate.ListAiProvidersResponse
}

func (s *GoldenGateAiProvidersDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *GoldenGateAiProvidersDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_golden_gate.ListAiProvidersRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "golden_gate")

	response, err := s.Client.ListAiProviders(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *GoldenGateAiProvidersDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("GoldenGateAiProvidersDataSource-", GoldenGateAiProvidersDataSource(), s.D))
	resources := []map[string]interface{}{}
	aiProvider := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AiProviderSummaryToMap(item))
	}
	aiProvider["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, GoldenGateAiProvidersDataSource().Schema["ai_provider_collection"].Elem.(*schema.Resource).Schema)
		aiProvider["items"] = items
	}

	resources = append(resources, aiProvider)
	if err := s.D.Set("ai_provider_collection", resources); err != nil {
		return err
	}

	return nil
}

func AiProviderSummaryToMap(obj oci_golden_gate.AiProviderSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["auth_type"] = obj.AuthType

	if obj.DefaultBaseUrl != nil {
		result["default_base_url"] = string(*obj.DefaultBaseUrl)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	models := []interface{}{}
	for _, item := range obj.Models {
		models = append(models, AiModelSummaryToMap(item))
	}
	result["models"] = models

	result["provider_type"] = string(obj.ProviderType)

	return result
}
