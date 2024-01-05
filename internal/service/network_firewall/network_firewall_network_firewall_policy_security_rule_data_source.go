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

func NetworkFirewallNetworkFirewallPolicySecurityRuleDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["network_firewall_policy_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["name"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(NetworkFirewallNetworkFirewallPolicySecurityRuleResource(), fieldMap, readSingularNetworkFirewallNetworkFirewallPolicySecurityRule)
}

func readSingularNetworkFirewallNetworkFirewallPolicySecurityRule(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicySecurityRuleDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()

	return tfresource.ReadResource(sync)
}

type NetworkFirewallNetworkFirewallPolicySecurityRuleDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_network_firewall.NetworkFirewallClient
	Res    *oci_network_firewall.GetSecurityRuleResponse
}

func (s *NetworkFirewallNetworkFirewallPolicySecurityRuleDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *NetworkFirewallNetworkFirewallPolicySecurityRuleDataSourceCrud) Get() error {
	request := oci_network_firewall.GetSecurityRuleRequest{}

	if networkFirewallPolicyId, ok := s.D.GetOkExists("network_firewall_policy_id"); ok {
		tmp := networkFirewallPolicyId.(string)
		request.NetworkFirewallPolicyId = &tmp
	}

	if securityRuleName, ok := s.D.GetOkExists("name"); ok {
		tmp := securityRuleName.(string)
		request.SecurityRuleName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "network_firewall")

	response, err := s.Client.GetSecurityRule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicySecurityRuleDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("NetworkFirewallNetworkFirewallPolicySecurityRuleDataSource-", NetworkFirewallNetworkFirewallPolicySecurityRuleDataSource(), s.D))

	s.D.Set("action", s.Res.Action)

	if s.Res.Condition != nil {
		s.D.Set("condition", []interface{}{SecurityRuleMatchCriteriaToMap(s.Res.Condition)})
	} else {
		s.D.Set("condition", nil)
	}

	s.D.Set("inspection", s.Res.Inspection)

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.ParentResourceId != nil {
		s.D.Set("parent_resource_id", *s.Res.ParentResourceId)
	}

	if s.Res.Position != nil {
		s.D.Set("position", []interface{}{RulePositionToMap(s.Res.Position)})
	} else {
		s.D.Set("position", nil)
	}

	return nil
}
