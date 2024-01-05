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

func NetworkFirewallNetworkFirewallPolicyMappedSecretsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readNetworkFirewallNetworkFirewallPolicyMappedSecrets,
		Schema: map[string]*schema.Schema{
			"network_firewall_policy_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"mapped_secret_summary_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     NetworkFirewallNetworkFirewallPolicyMappedSecretResource(),
						},
					},
				},
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func readNetworkFirewallNetworkFirewallPolicyMappedSecrets(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyMappedSecretsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()

	return tfresource.ReadResource(sync)
}

type NetworkFirewallNetworkFirewallPolicyMappedSecretsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_network_firewall.NetworkFirewallClient
	Res    *oci_network_firewall.ListMappedSecretsResponse
}

func (s *NetworkFirewallNetworkFirewallPolicyMappedSecretsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *NetworkFirewallNetworkFirewallPolicyMappedSecretsDataSourceCrud) Get() error {
	request := oci_network_firewall.ListMappedSecretsRequest{}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if networkFirewallPolicyId, ok := s.D.GetOkExists("network_firewall_policy_id"); ok {
		tmp := networkFirewallPolicyId.(string)
		request.NetworkFirewallPolicyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "network_firewall")

	response, err := s.Client.ListMappedSecrets(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListMappedSecrets(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicyMappedSecretsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("NetworkFirewallNetworkFirewallPolicyMappedSecretsDataSource-", NetworkFirewallNetworkFirewallPolicyMappedSecretsDataSource(), s.D))
	resources := []map[string]interface{}{}
	networkFirewallPolicyMappedSecret := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, MappedSecretSummaryToMap(item))
	}
	networkFirewallPolicyMappedSecret["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, NetworkFirewallNetworkFirewallPolicyMappedSecretsDataSource().Schema["mapped_secret_summary_collection"].Elem.(*schema.Resource).Schema)
		networkFirewallPolicyMappedSecret["items"] = items
	}

	resources = append(resources, networkFirewallPolicyMappedSecret)
	if err := s.D.Set("mapped_secret_summary_collection", resources); err != nil {
		return err
	}

	return nil
}
