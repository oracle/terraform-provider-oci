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

func TenantmanagercontrolplaneAssignedSubscriptionAvailableRegionsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readTenantmanagercontrolplaneAssignedSubscriptionAvailableRegionsWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"assigned_subscription_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"available_region_collection": {
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
									"region_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"system_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
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

func readTenantmanagercontrolplaneAssignedSubscriptionAvailableRegionsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &TenantmanagercontrolplaneAssignedSubscriptionAvailableRegionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OrganizationsSubscriptionClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type TenantmanagercontrolplaneAssignedSubscriptionAvailableRegionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_tenantmanagercontrolplane.SubscriptionClient
	Res    *oci_tenantmanagercontrolplane.ListAssignedSubscriptionAvailableRegionsResponse
}

func (s *TenantmanagercontrolplaneAssignedSubscriptionAvailableRegionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *TenantmanagercontrolplaneAssignedSubscriptionAvailableRegionsDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_tenantmanagercontrolplane.ListAssignedSubscriptionAvailableRegionsRequest{}

	if assignedSubscriptionId, ok := s.D.GetOkExists("assigned_subscription_id"); ok {
		tmp := assignedSubscriptionId.(string)
		request.AssignedSubscriptionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "tenantmanagercontrolplane")

	response, err := s.Client.ListAssignedSubscriptionAvailableRegions(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAssignedSubscriptionAvailableRegions(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *TenantmanagercontrolplaneAssignedSubscriptionAvailableRegionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("TenantmanagercontrolplaneAssignedSubscriptionAvailableRegionsDataSource-", TenantmanagercontrolplaneAssignedSubscriptionAvailableRegionsDataSource(), s.D))
	resources := []map[string]interface{}{}
	assignedSubscriptionAvailableRegion := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AssignedSubscriptionAvailableRegionSummaryToMap(item))
	}
	assignedSubscriptionAvailableRegion["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, TenantmanagercontrolplaneAssignedSubscriptionAvailableRegionsDataSource().Schema["available_region_collection"].Elem.(*schema.Resource).Schema)
		assignedSubscriptionAvailableRegion["items"] = items
	}

	resources = append(resources, assignedSubscriptionAvailableRegion)
	if err := s.D.Set("available_region_collection", resources); err != nil {
		return err
	}

	return nil
}

func AssignedSubscriptionAvailableRegionSummaryToMap(obj oci_tenantmanagercontrolplane.AvailableRegionSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.RegionName != nil {
		result["region_name"] = string(*obj.RegionName)
	}

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	return result
}
