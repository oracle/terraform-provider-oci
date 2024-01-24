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

func NetworkFirewallNetworkFirewallPolicyApplicationGroupDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["name"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["network_firewall_policy_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(NetworkFirewallNetworkFirewallPolicyApplicationGroupResource(), fieldMap, readSingularNetworkFirewallNetworkFirewallPolicyApplicationGroup)
}

func readSingularNetworkFirewallNetworkFirewallPolicyApplicationGroup(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyApplicationGroupDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()

	return tfresource.ReadResource(sync)
}

type NetworkFirewallNetworkFirewallPolicyApplicationGroupDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_network_firewall.NetworkFirewallClient
	Res    *oci_network_firewall.GetApplicationGroupResponse
}

func (s *NetworkFirewallNetworkFirewallPolicyApplicationGroupDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *NetworkFirewallNetworkFirewallPolicyApplicationGroupDataSourceCrud) Get() error {
	request := oci_network_firewall.GetApplicationGroupRequest{}

	if applicationGroupName, ok := s.D.GetOkExists("name"); ok {
		tmp := applicationGroupName.(string)
		request.ApplicationGroupName = &tmp
	}

	if networkFirewallPolicyId, ok := s.D.GetOkExists("network_firewall_policy_id"); ok {
		tmp := networkFirewallPolicyId.(string)
		request.NetworkFirewallPolicyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "network_firewall")

	response, err := s.Client.GetApplicationGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicyApplicationGroupDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("NetworkFirewallNetworkFirewallPolicyApplicationGroupDataSource-", NetworkFirewallNetworkFirewallPolicyApplicationGroupDataSource(), s.D))

	s.D.Set("apps", s.Res.Apps)

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.ParentResourceId != nil {
		s.D.Set("parent_resource_id", *s.Res.ParentResourceId)
	}

	if s.Res.TotalApps != nil {
		s.D.Set("total_apps", *s.Res.TotalApps)
	}

	return nil
}
