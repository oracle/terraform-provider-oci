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

func NetworkFirewallNetworkFirewallPolicyServiceDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["network_firewall_policy_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["name"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(NetworkFirewallNetworkFirewallPolicyServiceResource(), fieldMap, readSingularNetworkFirewallNetworkFirewallPolicyService)
}

func readSingularNetworkFirewallNetworkFirewallPolicyService(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyServiceDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()

	return tfresource.ReadResource(sync)
}

type NetworkFirewallNetworkFirewallPolicyServiceDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_network_firewall.NetworkFirewallClient
	Res    *oci_network_firewall.GetServiceResponse
}

func (s *NetworkFirewallNetworkFirewallPolicyServiceDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *NetworkFirewallNetworkFirewallPolicyServiceDataSourceCrud) Get() error {
	request := oci_network_firewall.GetServiceRequest{}

	if networkFirewallPolicyId, ok := s.D.GetOkExists("network_firewall_policy_id"); ok {
		tmp := networkFirewallPolicyId.(string)
		request.NetworkFirewallPolicyId = &tmp
	}

	if serviceName, ok := s.D.GetOkExists("name"); ok {
		tmp := serviceName.(string)
		request.ServiceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "network_firewall")

	response, err := s.Client.GetService(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicyServiceDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("NetworkFirewallNetworkFirewallPolicyServiceDataSource-", NetworkFirewallNetworkFirewallPolicyServiceDataSource(), s.D))
	switch v := (s.Res.Service).(type) {
	case oci_network_firewall.TcpService:
		s.D.Set("type", "TCP_SERVICE")

		portRanges := []interface{}{}
		for _, item := range v.PortRanges {
			portRanges = append(portRanges, PortRangeToMap(item))
		}
		s.D.Set("port_ranges", portRanges)

		if v.Name != nil {
			s.D.Set("name", *v.Name)
		}

		if v.ParentResourceId != nil {
			s.D.Set("parent_resource_id", *v.ParentResourceId)
		}
	case oci_network_firewall.UdpService:
		s.D.Set("type", "UDP_SERVICE")

		portRanges := []interface{}{}
		for _, item := range v.PortRanges {
			portRanges = append(portRanges, PortRangeToMap(item))
		}
		s.D.Set("port_ranges", portRanges)

		if v.Name != nil {
			s.D.Set("name", *v.Name)
		}

		if v.ParentResourceId != nil {
			s.D.Set("parent_resource_id", *v.ParentResourceId)
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", s.Res.Service)
		return nil
	}

	return nil
}
