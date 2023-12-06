// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package usage_proxy

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_usage_proxy "github.com/oracle/oci-go-sdk/v65/usage"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func UsageProxyResourcesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readUsageProxyResources,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"entitlement_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"resources_collection": {
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
									"child_resources": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"daily_unit_display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"hourly_unit_display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"instance_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_purchased": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"raw_unit_display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"servicename": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"skus": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"cloud_credit_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"sku_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"sku_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"usage_data_type": {
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

func readUsageProxyResources(d *schema.ResourceData, m interface{}) error {
	sync := &UsageProxyResourcesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ResourcesClient()

	return tfresource.ReadResource(sync)
}

type UsageProxyResourcesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_usage_proxy.ResourcesClient
	Res    *oci_usage_proxy.ListResourcesResponse
}

func (s *UsageProxyResourcesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *UsageProxyResourcesDataSourceCrud) Get() error {
	request := oci_usage_proxy.ListResourcesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if entitlementId, ok := s.D.GetOkExists("entitlement_id"); ok {
		tmp := entitlementId.(string)
		request.EntitlementId = &tmp
	}

	if serviceName, ok := s.D.GetOkExists("service_name"); ok {
		tmp := serviceName.(string)
		request.ServiceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "usage_proxy")

	response, err := s.Client.ListResources(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListResources(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *UsageProxyResourcesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("UsageProxyResourcesDataSource-", UsageProxyResourcesDataSource(), s.D))
	resources := []map[string]interface{}{}
	resource := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ResourceSummaryToMap(item))
	}
	resource["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, UsageProxyResourcesDataSource().Schema["resources_collection"].Elem.(*schema.Resource).Schema)
		resource["items"] = items
	}

	resources = append(resources, resource)
	if err := s.D.Set("resources_collection", resources); err != nil {
		return err
	}

	return nil
}

func ResourceSummaryToMap(obj oci_usage_proxy.ResourceSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["child_resources"] = obj.ChildResources

	if obj.DailyUnitDisplayName != nil {
		result["daily_unit_display_name"] = string(*obj.DailyUnitDisplayName)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.HourlyUnitDisplayName != nil {
		result["hourly_unit_display_name"] = string(*obj.HourlyUnitDisplayName)
	}

	if obj.InstanceType != nil {
		result["instance_type"] = string(*obj.InstanceType)
	}

	if obj.IsPurchased != nil {
		result["is_purchased"] = bool(*obj.IsPurchased)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.RawUnitDisplayName != nil {
		result["raw_unit_display_name"] = string(*obj.RawUnitDisplayName)
	}

	if obj.Servicename != nil {
		result["servicename"] = string(*obj.Servicename)
	}

	skus := []interface{}{}
	for _, item := range obj.Skus {
		skus = append(skus, SkuProductsToMap(item))
	}
	result["skus"] = skus

	result["usage_data_type"] = string(obj.UsageDataType)

	return result
}

func SkuProductsToMap(obj oci_usage_proxy.SkuProducts) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CloudCreditType != nil {
		result["cloud_credit_type"] = string(*obj.CloudCreditType)
	}

	if obj.SkuId != nil {
		result["sku_id"] = string(*obj.SkuId)
	}

	if obj.SkuType != nil {
		result["sku_type"] = string(*obj.SkuType)
	}

	return result
}
