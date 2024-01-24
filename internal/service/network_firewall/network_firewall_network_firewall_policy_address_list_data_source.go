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

func NetworkFirewallNetworkFirewallPolicyAddressListDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["network_firewall_policy_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["name"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(NetworkFirewallNetworkFirewallPolicyAddressListResource(), fieldMap, readSingularNetworkFirewallNetworkFirewallPolicyAddressList)
}

func readSingularNetworkFirewallNetworkFirewallPolicyAddressList(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyAddressListDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()

	return tfresource.ReadResource(sync)
}

type NetworkFirewallNetworkFirewallPolicyAddressListDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_network_firewall.NetworkFirewallClient
	Res    *oci_network_firewall.GetAddressListResponse
}

func (s *NetworkFirewallNetworkFirewallPolicyAddressListDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *NetworkFirewallNetworkFirewallPolicyAddressListDataSourceCrud) Get() error {
	request := oci_network_firewall.GetAddressListRequest{}

	if addressListName, ok := s.D.GetOkExists("name"); ok {
		tmp := addressListName.(string)
		request.AddressListName = &tmp
	}

	if networkFirewallPolicyId, ok := s.D.GetOkExists("network_firewall_policy_id"); ok {
		tmp := networkFirewallPolicyId.(string)
		request.NetworkFirewallPolicyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "network_firewall")

	response, err := s.Client.GetAddressList(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicyAddressListDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("NetworkFirewallNetworkFirewallPolicyAddressListDataSource-", NetworkFirewallNetworkFirewallPolicyAddressListDataSource(), s.D))

	s.D.Set("addresses", s.Res.Addresses)

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.ParentResourceId != nil {
		s.D.Set("parent_resource_id", *s.Res.ParentResourceId)
	}

	if s.Res.TotalAddresses != nil {
		s.D.Set("total_addresses", *s.Res.TotalAddresses)
	}

	s.D.Set("type", s.Res.Type)

	return nil
}
