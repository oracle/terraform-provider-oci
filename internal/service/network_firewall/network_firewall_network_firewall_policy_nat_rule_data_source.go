// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package network_firewall

import (
	"context"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_network_firewall "github.com/oracle/oci-go-sdk/v65/networkfirewall"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func NetworkFirewallNetworkFirewallPolicyNatRuleDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["nat_rule_name"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["network_firewall_policy_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(NetworkFirewallNetworkFirewallPolicyNatRuleResource(), fieldMap, readSingularNetworkFirewallNetworkFirewallPolicyNatRule)
}

func readSingularNetworkFirewallNetworkFirewallPolicyNatRule(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyNatRuleDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()

	return tfresource.ReadResource(sync)
}

type NetworkFirewallNetworkFirewallPolicyNatRuleDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_network_firewall.NetworkFirewallClient
	Res    *oci_network_firewall.GetNatRuleResponse
}

func (s *NetworkFirewallNetworkFirewallPolicyNatRuleDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *NetworkFirewallNetworkFirewallPolicyNatRuleDataSourceCrud) Get() error {
	request := oci_network_firewall.GetNatRuleRequest{}

	if natRuleName, ok := s.D.GetOkExists("nat_rule_name"); ok {
		tmp := natRuleName.(string)
		request.NatRuleName = &tmp
	}

	if networkFirewallPolicyId, ok := s.D.GetOkExists("network_firewall_policy_id"); ok {
		tmp := networkFirewallPolicyId.(string)
		request.NetworkFirewallPolicyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "network_firewall")

	response, err := s.Client.GetNatRule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicyNatRuleDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("NetworkFirewallNetworkFirewallPolicyNatRuleDataSource-", NetworkFirewallNetworkFirewallPolicyNatRuleDataSource(), s.D))
	switch v := (s.Res.NatRule).(type) {
	case oci_network_firewall.NatV4NatRule:
		s.D.Set("type", "NATV4")

		s.D.Set("action", v.Action)

		if v.Condition != nil {
			s.D.Set("condition", []interface{}{NatRuleMatchCriteriaToMap(v.Condition)})
		} else {
			s.D.Set("condition", nil)
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.Name != nil {
			s.D.Set("name", *v.Name)
			s.D.Set("nat_rule_name", *v.Name)
		}

		if v.ParentResourceId != nil {
			s.D.Set("parent_resource_id", *v.ParentResourceId)
		}

		if v.Position != nil {
			s.D.Set("position", []interface{}{RulePositionToMap(v.Position)})
		} else {
			s.D.Set("position", nil)
		}

		if v.PriorityOrder != nil {
			s.D.Set("priority_order", strconv.FormatInt(*v.PriorityOrder, 10))
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", s.Res.NatRule)
		return nil
	}

	return nil
}
