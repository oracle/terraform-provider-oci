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

func UsageProxyUsagelimitsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readUsageProxyUsagelimits,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"limit_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"subscription_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"usage_limit_collection": {
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
									"action": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"alert_level": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"created_by": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"entitlement_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"limit_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"max_hard_limit": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"limit": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"modified_by": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"resource_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"service_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"sku_part_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_created": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_modified": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"value_type": {
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

func readUsageProxyUsagelimits(d *schema.ResourceData, m interface{}) error {
	sync := &UsageProxyUsagelimitsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).UsagelimitsClient()

	return tfresource.ReadResource(sync)
}

type UsageProxyUsagelimitsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_usage_proxy.UsagelimitsClient
	Res    *oci_usage_proxy.ListUsageLimitsResponse
}

func (s *UsageProxyUsagelimitsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *UsageProxyUsagelimitsDataSourceCrud) Get() error {
	request := oci_usage_proxy.ListUsageLimitsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if limitType, ok := s.D.GetOkExists("limit_type"); ok {
		tmp := limitType.(string)
		request.LimitType = &tmp
	}

	if resourceType, ok := s.D.GetOkExists("resource_type"); ok {
		tmp := resourceType.(string)
		request.ResourceType = &tmp
	}

	if serviceType, ok := s.D.GetOkExists("service_type"); ok {
		tmp := serviceType.(string)
		request.ServiceType = &tmp
	}

	if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
		tmp := subscriptionId.(string)
		request.SubscriptionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "usage_proxy")

	response, err := s.Client.ListUsageLimits(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListUsageLimits(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *UsageProxyUsagelimitsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("UsageProxyUsagelimitsDataSource-", UsageProxyUsagelimitsDataSource(), s.D))
	resources := []map[string]interface{}{}
	usagelimit := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, UsageLimitSummaryToMap(item))
	}
	usagelimit["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, UsageProxyUsagelimitsDataSource().Schema["usage_limit_collection"].Elem.(*schema.Resource).Schema)
		usagelimit["items"] = items
	}

	resources = append(resources, usagelimit)
	if err := s.D.Set("usage_limit_collection", resources); err != nil {
		return err
	}

	return nil
}

func UsageLimitSummaryToMap(obj oci_usage_proxy.UsageLimitSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["action"] = string(obj.Action)

	if obj.AlertLevel != nil {
		result["alert_level"] = float32(*obj.AlertLevel)
	}

	if obj.CreatedBy != nil {
		result["created_by"] = string(*obj.CreatedBy)
	}

	if obj.EntitlementId != nil {
		result["entitlement_id"] = string(*obj.EntitlementId)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Limit != nil {
		result["limit"] = string(*obj.Limit)
	}

	result["limit_type"] = string(obj.LimitType)

	if obj.MaxHardLimit != nil {
		result["max_hard_limit"] = string(*obj.MaxHardLimit)
	}

	if obj.ModifiedBy != nil {
		result["modified_by"] = string(*obj.ModifiedBy)
	}

	if obj.ResourceName != nil {
		result["resource_name"] = string(*obj.ResourceName)
	}

	if obj.ServiceName != nil {
		result["service_name"] = string(*obj.ServiceName)
	}

	if obj.SkuPartId != nil {
		result["sku_part_id"] = string(*obj.SkuPartId)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeModified != nil {
		result["time_modified"] = obj.TimeModified.String()
	}

	result["value_type"] = string(obj.ValueType)

	return result
}
