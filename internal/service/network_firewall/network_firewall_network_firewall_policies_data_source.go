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

func NetworkFirewallNetworkFirewallPoliciesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readNetworkFirewallNetworkFirewallPolicies,
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
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_firewall_policy_summary_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(NetworkFirewallNetworkFirewallPolicyResource()),
						},
					},
				},
			},
		},
	}
}

func readNetworkFirewallNetworkFirewallPolicies(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPoliciesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()

	return tfresource.ReadResource(sync)
}

type NetworkFirewallNetworkFirewallPoliciesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_network_firewall.NetworkFirewallClient
	Res    *oci_network_firewall.ListNetworkFirewallPoliciesResponse
}

func (s *NetworkFirewallNetworkFirewallPoliciesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *NetworkFirewallNetworkFirewallPoliciesDataSourceCrud) Get() error {
	request := oci_network_firewall.ListNetworkFirewallPoliciesRequest{}

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

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_network_firewall.ListNetworkFirewallPoliciesLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "network_firewall")

	response, err := s.Client.ListNetworkFirewallPolicies(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListNetworkFirewallPolicies(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *NetworkFirewallNetworkFirewallPoliciesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("NetworkFirewallNetworkFirewallPoliciesDataSource-", NetworkFirewallNetworkFirewallPoliciesDataSource(), s.D))
	resources := []map[string]interface{}{}
	networkFirewallPolicy := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, NetworkFirewallPolicySummaryToMap(item))
	}
	networkFirewallPolicy["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, NetworkFirewallNetworkFirewallPoliciesDataSource().Schema["network_firewall_policy_summary_collection"].Elem.(*schema.Resource).Schema)
		networkFirewallPolicy["items"] = items
	}

	resources = append(resources, networkFirewallPolicy)
	if err := s.D.Set("network_firewall_policy_summary_collection", resources); err != nil {
		return err
	}

	return nil
}
