// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package tenantmanagercontrolplane

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_tenantmanagercontrolplane "github.com/oracle/oci-go-sdk/v65/tenantmanagercontrolplane"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func TenantmanagercontrolplaneLinkFeaturesDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readTenantmanagercontrolplaneLinkFeaturesWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"link_features_collection": {
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
									"feature": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"partner_service_console_url": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"user_guide_url": {
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

func readTenantmanagercontrolplaneLinkFeaturesWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &TenantmanagercontrolplaneLinkFeaturesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LinkFeaturesClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type TenantmanagercontrolplaneLinkFeaturesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_tenantmanagercontrolplane.LinkFeaturesClient
	Res    *oci_tenantmanagercontrolplane.ListLinkFeaturesResponse
}

func (s *TenantmanagercontrolplaneLinkFeaturesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *TenantmanagercontrolplaneLinkFeaturesDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_tenantmanagercontrolplane.ListLinkFeaturesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "tenantmanagercontrolplane")

	response, err := s.Client.ListLinkFeatures(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *TenantmanagercontrolplaneLinkFeaturesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("TenantmanagercontrolplaneLinkFeaturesDataSource-", TenantmanagercontrolplaneLinkFeaturesDataSource(), s.D))
	resources := []map[string]interface{}{}
	linkFeature := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, LinkFeatureSummaryToMap(item))
	}
	linkFeature["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, TenantmanagercontrolplaneLinkFeaturesDataSource().Schema["link_features_collection"].Elem.(*schema.Resource).Schema)
		linkFeature["items"] = items
	}

	resources = append(resources, linkFeature)
	if err := s.D.Set("link_features_collection", resources); err != nil {
		return err
	}

	return nil
}

func LinkFeatureSummaryToMap(obj oci_tenantmanagercontrolplane.LinkFeatureSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Feature != nil {
		result["feature"] = string(*obj.Feature)
	}

	if obj.PartnerServiceConsoleUrl != nil {
		result["partner_service_console_url"] = string(*obj.PartnerServiceConsoleUrl)
	}

	if obj.UserGuideUrl != nil {
		result["user_guide_url"] = string(*obj.UserGuideUrl)
	}

	return result
}
