// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package network_firewall

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_network_firewall "github.com/oracle/oci-go-sdk/v65/networkfirewall"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func NetworkFirewallNetworkFirewallPolicyDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["network_firewall_policy_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(NetworkFirewallNetworkFirewallPolicyResource(), fieldMap, readSingularNetworkFirewallNetworkFirewallPolicy)
}

func readSingularNetworkFirewallNetworkFirewallPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()

	return tfresource.ReadResource(sync)
}

type NetworkFirewallNetworkFirewallPolicyDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_network_firewall.NetworkFirewallClient
	Res    *oci_network_firewall.GetNetworkFirewallPolicyResponse
}

func (s *NetworkFirewallNetworkFirewallPolicyDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *NetworkFirewallNetworkFirewallPolicyDataSourceCrud) Get() error {
	request := oci_network_firewall.GetNetworkFirewallPolicyRequest{}

	if networkFirewallPolicyId, ok := s.D.GetOkExists("network_firewall_policy_id"); ok {
		tmp := networkFirewallPolicyId.(string)
		request.NetworkFirewallPolicyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "network_firewall")

	response, err := s.Client.GetNetworkFirewallPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicyDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	//s.D.Set("application_lists", s.Res.ApplicationLists)
	if s.Res.ApplicationLists != nil {
		s.D.Set("application_lists", propertyApplicationListsToMap(s.Res.ApplicationLists))
	} else {
		s.D.Set("application_lists", nil)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	//s.D.Set("decryption_profiles", s.Res.DecryptionProfiles)

	if s.Res.DecryptionProfiles != nil {
		s.D.Set("decryption_profiles", DecryptionProfileMapToMap(s.Res.DecryptionProfiles))
	} else {
		s.D.Set("decryption_profiles", nil)
	}

	decryptionRules := []interface{}{}
	for _, item := range s.Res.DecryptionRules {
		decryptionRules = append(decryptionRules, DecryptionRuleToMap(item))
	}
	s.D.Set("decryption_rules", decryptionRules)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	//s.D.Set("ip_address_lists", s.Res.IpAddressLists)
	if s.Res.IpAddressLists != nil {
		s.D.Set("ip_address_lists", ipAddressListsToMap(s.Res.IpAddressLists))
	} else {
		s.D.Set("ip_address_lists", nil)
	}

	if s.Res.IsFirewallAttached != nil {
		s.D.Set("is_firewall_attached", *s.Res.IsFirewallAttached)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	//s.D.Set("mapped_secrets", s.Res.MappedSecrets)
	if s.Res.MappedSecrets != nil {
		s.D.Set("mapped_secrets", MappedSecretsToMap(s.Res.MappedSecrets))
	} else {
		s.D.Set("mapped_secrets", nil)
	}

	securityRules := []interface{}{}
	for _, item := range s.Res.SecurityRules {
		securityRules = append(securityRules, SecurityRuleToMap(item))
	}
	s.D.Set("security_rules", securityRules)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	//s.D.Set("url_lists", s.Res.UrlLists)
	if s.Res.UrlLists != nil {
		s.D.Set("url_lists", propertyUrlListsToMap(s.Res.UrlLists))
	} else {
		s.D.Set("url_lists", nil)
	}
	return nil
}
