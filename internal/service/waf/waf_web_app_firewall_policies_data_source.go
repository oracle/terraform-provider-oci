// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package waf

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_waf "github.com/oracle/oci-go-sdk/v65/waf"
)

func WafWebAppFirewallPoliciesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readWafWebAppFirewallPolicies,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
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
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"web_app_firewall_policy_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(WafWebAppFirewallPolicyResource()),
						},
					},
				},
			},
		},
	}
}

func readWafWebAppFirewallPolicies(d *schema.ResourceData, m interface{}) error {
	sync := &WafWebAppFirewallPoliciesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).WafClient()

	return tfresource.ReadResource(sync)
}

type WafWebAppFirewallPoliciesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_waf.WafClient
	Res    *oci_waf.ListWebAppFirewallPoliciesResponse
}

func (s *WafWebAppFirewallPoliciesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *WafWebAppFirewallPoliciesDataSourceCrud) Get() error {
	request := oci_waf.ListWebAppFirewallPoliciesRequest{}

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

	// allows filtering by multiple values
	if state, ok := s.D.GetOkExists("state"); ok {
		interfaces := state.([]interface{})
		tmp := make([]oci_waf.WebAppFirewallPolicyLifecycleStateEnum, len(interfaces))

		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_waf.WebAppFirewallPolicyLifecycleStateEnum(interfaces[i].(string))
			}
		}

		if len(tmp) != 0 || s.D.HasChange("state") {
			request.LifecycleState = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "waf")

	response, err := s.Client.ListWebAppFirewallPolicies(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListWebAppFirewallPolicies(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *WafWebAppFirewallPoliciesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("WafWebAppFirewallPoliciesDataSource-", WafWebAppFirewallPoliciesDataSource(), s.D))
	resources := []map[string]interface{}{}
	webAppFirewallPolicy := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, WebAppFirewallPolicySummaryToMap(item))
	}
	webAppFirewallPolicy["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, WafWebAppFirewallPoliciesDataSource().Schema["web_app_firewall_policy_collection"].Elem.(*schema.Resource).Schema)
		webAppFirewallPolicy["items"] = items
	}

	resources = append(resources, webAppFirewallPolicy)
	if err := s.D.Set("web_app_firewall_policy_collection", resources); err != nil {
		return err
	}

	return nil
}
