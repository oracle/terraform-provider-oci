// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package network_firewall

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_network_firewall "github.com/oracle/oci-go-sdk/v65/networkfirewall"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func NetworkFirewallNetworkFirewallPolicyApplicationDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["network_firewall_policy_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["name"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(NetworkFirewallNetworkFirewallPolicyApplicationResource(), fieldMap, readSingularApplication)
}

func readSingularApplication(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyApplicationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()

	return tfresource.ReadResource(sync)
}

type NetworkFirewallNetworkFirewallPolicyApplicationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_network_firewall.NetworkFirewallClient
	Res    *oci_network_firewall.GetApplicationResponse
}

func (s *NetworkFirewallNetworkFirewallPolicyApplicationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *NetworkFirewallNetworkFirewallPolicyApplicationDataSourceCrud) Get() error {
	request := oci_network_firewall.GetApplicationRequest{}

	if applicationName, ok := s.D.GetOkExists("name"); ok {
		tmp := applicationName.(string)
		request.ApplicationName = &tmp
	}

	if networkFirewallPolicyId, ok := s.D.GetOkExists("network_firewall_policy_id"); ok {
		tmp := networkFirewallPolicyId.(string)
		request.NetworkFirewallPolicyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "network_firewall")

	response, err := s.Client.GetApplication(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicyApplicationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("NetworkFirewallNetworkFirewallPolicyApplicationDataSource-", NetworkFirewallNetworkFirewallPolicyApplicationDataSource(), s.D))
	switch v := (s.Res.Application).(type) {
	case oci_network_firewall.IcmpApplication:
		s.D.Set("type", "ICMP")

		if v.IcmpCode != nil {
			s.D.Set("icmp_code", *v.IcmpCode)
		}

		if v.IcmpType != nil {
			s.D.Set("icmp_type", *v.IcmpType)
		}

		if v.Name != nil {
			s.D.Set("name", *v.Name)
		}

		if v.ParentResourceId != nil {
			s.D.Set("parent_resource_id", *v.ParentResourceId)
		}
	case oci_network_firewall.Icmp6Application:
		s.D.Set("type", "ICMP_V6")

		if v.IcmpCode != nil {
			s.D.Set("icmp_code", *v.IcmpCode)
		}

		if v.IcmpType != nil {
			s.D.Set("icmp_type", *v.IcmpType)
		}

		if v.Name != nil {
			s.D.Set("name", *v.Name)
		}

		if v.ParentResourceId != nil {
			s.D.Set("parent_resource_id", *v.ParentResourceId)
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", s.Res.Application)
		return nil
	}

	return nil
}
