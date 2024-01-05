// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package usage_proxy

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_usage_proxy "github.com/oracle/oci-go-sdk/v65/usage"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func UsageProxyResourceQuotasDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readUsageProxyResourceQuotas,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"service_entitlement": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"resource_quotum_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"is_allowed": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"affected_resource": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"balance": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"is_allowed": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_dependency": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_overage": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"purchased_limit": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"service": {
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

func readUsageProxyResourceQuotas(d *schema.ResourceData, m interface{}) error {
	sync := &UsageProxyResourceQuotasDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ResourcesClient()

	return tfresource.ReadResource(sync)
}

type UsageProxyResourceQuotasDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_usage_proxy.ResourcesClient
	Res    *oci_usage_proxy.ListResourceQuotaResponse
}

func (s *UsageProxyResourceQuotasDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *UsageProxyResourceQuotasDataSourceCrud) Get() error {
	request := oci_usage_proxy.ListResourceQuotaRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if serviceEntitlement, ok := s.D.GetOkExists("service_entitlement"); ok {
		tmp := serviceEntitlement.(string)
		request.ServiceEntitlement = &tmp
	}

	if serviceName, ok := s.D.GetOkExists("service_name"); ok {
		tmp := serviceName.(string)
		request.ServiceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "usage_proxy")

	response, err := s.Client.ListResourceQuota(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListResourceQuota(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *UsageProxyResourceQuotasDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("UsageProxyResourceQuotasDataSource-", UsageProxyResourceQuotasDataSource(), s.D))
	resources := []map[string]interface{}{}
	resourceQuota := map[string]interface{}{}

	if s.Res.IsAllowed != nil {
		resourceQuota["is_allowed"] = *s.Res.IsAllowed
	}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ResourceQuotumSummaryToMap(item))
	}
	resourceQuota["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, UsageProxyResourceQuotasDataSource().Schema["resource_quotum_collection"].Elem.(*schema.Resource).Schema)
		resourceQuota["items"] = items
	}

	resources = append(resources, resourceQuota)
	if err := s.D.Set("resource_quotum_collection", resources); err != nil {
		return err
	}

	return nil
}

func ResourceQuotumSummaryToMap(obj oci_usage_proxy.ResourceQuotumSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AffectedResource != nil {
		result["affected_resource"] = string(*obj.AffectedResource)
	}

	if obj.Balance != nil {
		result["balance"] = float64(*obj.Balance)
	}

	if obj.IsAllowed != nil {
		result["is_allowed"] = bool(*obj.IsAllowed)
	}

	if obj.IsDependency != nil {
		result["is_dependency"] = bool(*obj.IsDependency)
	}

	if obj.IsOverage != nil {
		result["is_overage"] = bool(*obj.IsOverage)
	}

	if obj.Limit != nil {
		result["limit"] = float64(*obj.Limit)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.PurchasedLimit != nil {
		result["purchased_limit"] = float64(*obj.PurchasedLimit)
	}

	if obj.Service != nil {
		result["service"] = string(*obj.Service)
	}

	return result
}
