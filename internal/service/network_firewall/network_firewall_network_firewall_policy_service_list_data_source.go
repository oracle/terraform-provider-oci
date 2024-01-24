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

func NetworkFirewallNetworkFirewallPolicyServiceListDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["network_firewall_policy_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["name"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(NetworkFirewallNetworkFirewallPolicyServiceListResource(), fieldMap, readSingularNetworkFirewallNetworkFirewallPolicyServiceList)
}

func readSingularNetworkFirewallNetworkFirewallPolicyServiceList(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyServiceListDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()

	return tfresource.ReadResource(sync)
}

type NetworkFirewallNetworkFirewallPolicyServiceListDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_network_firewall.NetworkFirewallClient
	Res    *oci_network_firewall.GetServiceListResponse
}

func (s *NetworkFirewallNetworkFirewallPolicyServiceListDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *NetworkFirewallNetworkFirewallPolicyServiceListDataSourceCrud) Get() error {
	request := oci_network_firewall.GetServiceListRequest{}

	if networkFirewallPolicyId, ok := s.D.GetOkExists("network_firewall_policy_id"); ok {
		tmp := networkFirewallPolicyId.(string)
		request.NetworkFirewallPolicyId = &tmp
	}

	if serviceListName, ok := s.D.GetOkExists("name"); ok {
		tmp := serviceListName.(string)
		request.ServiceListName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "network_firewall")

	response, err := s.Client.GetServiceList(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicyServiceListDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("NetworkFirewallNetworkFirewallPolicyServiceListDataSource-", NetworkFirewallNetworkFirewallPolicyServiceListDataSource(), s.D))

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.ParentResourceId != nil {
		s.D.Set("parent_resource_id", *s.Res.ParentResourceId)
	}

	s.D.Set("services", s.Res.Services)

	if s.Res.TotalServices != nil {
		s.D.Set("total_services", *s.Res.TotalServices)
	}

	return nil
}
