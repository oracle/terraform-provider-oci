// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package network_firewall

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_network_firewall "github.com/oracle/oci-go-sdk/v65/networkfirewall"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func NetworkFirewallNetworkFirewallPolicyUrlListsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readNetworkFirewallNetworkFirewallPolicyUrlLists,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_firewall_policy_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"url_list_summary_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     NetworkFirewallNetworkFirewallPolicyUrlListResource(),
						},
					},
				},
			},
		},
	}
}

func readNetworkFirewallNetworkFirewallPolicyUrlLists(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyUrlListsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()

	return tfresource.ReadResource(sync)
}

type NetworkFirewallNetworkFirewallPolicyUrlListsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_network_firewall.NetworkFirewallClient
	Res    *oci_network_firewall.ListUrlListsResponse
}

func (s *NetworkFirewallNetworkFirewallPolicyUrlListsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *NetworkFirewallNetworkFirewallPolicyUrlListsDataSourceCrud) Get() error {
	request := oci_network_firewall.ListUrlListsRequest{}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if networkFirewallPolicyId, ok := s.D.GetOkExists("network_firewall_policy_id"); ok {
		tmp := networkFirewallPolicyId.(string)
		request.NetworkFirewallPolicyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "network_firewall")

	response, err := s.Client.ListUrlLists(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListUrlLists(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicyUrlListsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("NetworkFirewallNetworkFirewallPolicyUrlListsDataSource-", NetworkFirewallNetworkFirewallPolicyUrlListsDataSource(), s.D))
	resources := []map[string]interface{}{}
	networkFirewallPolicyUrlList := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, UrlListSummaryToMap(item))
	}
	networkFirewallPolicyUrlList["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, NetworkFirewallNetworkFirewallPolicyUrlListsDataSource().Schema["url_list_summary_collection"].Elem.(*schema.Resource).Schema)
		networkFirewallPolicyUrlList["items"] = items
	}

	resources = append(resources, networkFirewallPolicyUrlList)
	if err := s.D.Set("url_list_summary_collection", resources); err != nil {
		return err
	}

	return nil
}
