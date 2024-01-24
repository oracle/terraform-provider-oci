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

func NetworkFirewallNetworkFirewallPolicyDecryptionRuleDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["name"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["network_firewall_policy_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(NetworkFirewallNetworkFirewallPolicyDecryptionRuleResource(), fieldMap, readSingularNetworkFirewallNetworkFirewallPolicyDecryptionRule)
}

func readSingularNetworkFirewallNetworkFirewallPolicyDecryptionRule(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyDecryptionRuleDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()

	return tfresource.ReadResource(sync)
}

type NetworkFirewallNetworkFirewallPolicyDecryptionRuleDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_network_firewall.NetworkFirewallClient
	Res    *oci_network_firewall.GetDecryptionRuleResponse
}

func (s *NetworkFirewallNetworkFirewallPolicyDecryptionRuleDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *NetworkFirewallNetworkFirewallPolicyDecryptionRuleDataSourceCrud) Get() error {
	request := oci_network_firewall.GetDecryptionRuleRequest{}

	if decryptionRuleName, ok := s.D.GetOkExists("name"); ok {
		tmp := decryptionRuleName.(string)
		request.DecryptionRuleName = &tmp
	}

	if networkFirewallPolicyId, ok := s.D.GetOkExists("network_firewall_policy_id"); ok {
		tmp := networkFirewallPolicyId.(string)
		request.NetworkFirewallPolicyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "network_firewall")

	response, err := s.Client.GetDecryptionRule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicyDecryptionRuleDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("NetworkFirewallNetworkFirewallPolicyDecryptionRuleDataSource-", NetworkFirewallNetworkFirewallPolicyDecryptionRuleDataSource(), s.D))

	s.D.Set("action", s.Res.Action)

	if s.Res.Condition != nil {
		s.D.Set("condition", []interface{}{DecryptionRuleMatchCriteriaToMap(s.Res.Condition)})
	} else {
		s.D.Set("condition", nil)
	}

	if s.Res.DecryptionProfile != nil {
		s.D.Set("decryption_profile", *s.Res.DecryptionProfile)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.ParentResourceId != nil {
		s.D.Set("parent_resource_id", *s.Res.ParentResourceId)
	}

	if s.Res.Position != nil {
		s.D.Set("position", []interface{}{RulePositionToMapDecRule(s.Res.Position)})
	} else {
		s.D.Set("position", nil)
	}

	if s.Res.Secret != nil {
		s.D.Set("secret", *s.Res.Secret)
	}

	return nil
}
