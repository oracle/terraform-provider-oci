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

func NetworkFirewallNetworkFirewallPolicyTunnelInspectionRuleDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["network_firewall_policy_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["tunnel_inspection_rule_name"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(NetworkFirewallNetworkFirewallPolicyTunnelInspectionRuleResource(), fieldMap, readSingularNetworkFirewallNetworkFirewallPolicyTunnelInspectionRule)
}

func readSingularNetworkFirewallNetworkFirewallPolicyTunnelInspectionRule(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyTunnelInspectionRuleDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()

	return tfresource.ReadResource(sync)
}

type NetworkFirewallNetworkFirewallPolicyTunnelInspectionRuleDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_network_firewall.NetworkFirewallClient
	Res    *oci_network_firewall.GetTunnelInspectionRuleResponse
}

func (s *NetworkFirewallNetworkFirewallPolicyTunnelInspectionRuleDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *NetworkFirewallNetworkFirewallPolicyTunnelInspectionRuleDataSourceCrud) Get() error {
	request := oci_network_firewall.GetTunnelInspectionRuleRequest{}

	if networkFirewallPolicyId, ok := s.D.GetOkExists("network_firewall_policy_id"); ok {
		tmp := networkFirewallPolicyId.(string)
		request.NetworkFirewallPolicyId = &tmp
	}

	if tunnelInspectionRuleName, ok := s.D.GetOkExists("tunnel_inspection_rule_name"); ok {
		tmp := tunnelInspectionRuleName.(string)
		request.TunnelInspectionRuleName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "network_firewall")

	response, err := s.Client.GetTunnelInspectionRule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicyTunnelInspectionRuleDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("NetworkFirewallNetworkFirewallPolicyTunnelInspectionRuleDataSource-", NetworkFirewallNetworkFirewallPolicyTunnelInspectionRuleDataSource(), s.D))
	switch v := (s.Res.TunnelInspectionRule).(type) {
	case oci_network_firewall.VxlanInspectionRule:
		s.D.Set("protocol", "VXLAN")

		if v.Condition != nil {
			s.D.Set("condition", []interface{}{VxlanInspectionRuleMatchCriteriaToMap(v.Condition)})
		} else {
			s.D.Set("condition", nil)
		}

		if v.Profile != nil {
			s.D.Set("profile", []interface{}{VxlanInspectionRuleProfileToMap(v.Profile)})
		} else {
			s.D.Set("profile", nil)
		}

		s.D.Set("action", v.Action)

		if v.Name != nil {
			s.D.Set("name", *v.Name)
			s.D.Set("tunnel_inspection_rule_name", *v.Name)
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
		log.Printf("[WARN] Received 'protocol' of unknown type %v", s.Res.TunnelInspectionRule)
		return nil
	}

	return nil
}
