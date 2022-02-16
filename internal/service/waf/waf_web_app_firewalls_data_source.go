// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package waf

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_waf "github.com/oracle/oci-go-sdk/v58/waf"
)

func WafWebAppFirewallsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readWafWebAppFirewalls,
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
			"web_app_firewall_policy_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"web_app_firewall_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(WafWebAppFirewallResource()),
						},
					},
				},
			},
		},
	}
}

func readWafWebAppFirewalls(d *schema.ResourceData, m interface{}) error {
	sync := &WafWebAppFirewallsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).WafClient()

	return tfresource.ReadResource(sync)
}

type WafWebAppFirewallsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_waf.WafClient
	Res    *oci_waf.ListWebAppFirewallsResponse
}

func (s *WafWebAppFirewallsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *WafWebAppFirewallsDataSourceCrud) Get() error {
	request := oci_waf.ListWebAppFirewallsRequest{}

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
		tmp := make([]oci_waf.WebAppFirewallLifecycleStateEnum, len(interfaces))

		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_waf.WebAppFirewallLifecycleStateEnum(interfaces[i].(string))
			}
		}

		if len(tmp) != 0 || s.D.HasChange("state") {
			request.LifecycleState = tmp
		}
	}

	if webAppFirewallPolicyId, ok := s.D.GetOkExists("web_app_firewall_policy_id"); ok {
		tmp := webAppFirewallPolicyId.(string)
		request.WebAppFirewallPolicyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "waf")

	response, err := s.Client.ListWebAppFirewalls(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListWebAppFirewalls(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *WafWebAppFirewallsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("WafWebAppFirewallsDataSource-", WafWebAppFirewallsDataSource(), s.D))
	resources := []map[string]interface{}{}
	webAppFirewall := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, WebAppFirewallSummaryToMap(item))
	}
	webAppFirewall["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, WafWebAppFirewallsDataSource().Schema["web_app_firewall_collection"].Elem.(*schema.Resource).Schema)
		webAppFirewall["items"] = items
	}

	resources = append(resources, webAppFirewall)
	if err := s.D.Set("web_app_firewall_collection", resources); err != nil {
		return err
	}

	return nil
}
