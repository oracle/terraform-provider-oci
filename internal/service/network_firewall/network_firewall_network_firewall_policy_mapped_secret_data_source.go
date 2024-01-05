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

func NetworkFirewallNetworkFirewallPolicyMappedSecretDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["name"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["network_firewall_policy_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(NetworkFirewallNetworkFirewallPolicyMappedSecretResource(), fieldMap, readSingularNetworkFirewallNetworkFirewallPolicyMappedSecret)
}

func readSingularNetworkFirewallNetworkFirewallPolicyMappedSecret(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyMappedSecretDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()

	return tfresource.ReadResource(sync)
}

type NetworkFirewallNetworkFirewallPolicyMappedSecretDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_network_firewall.NetworkFirewallClient
	Res    *oci_network_firewall.GetMappedSecretResponse
}

func (s *NetworkFirewallNetworkFirewallPolicyMappedSecretDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *NetworkFirewallNetworkFirewallPolicyMappedSecretDataSourceCrud) Get() error {
	request := oci_network_firewall.GetMappedSecretRequest{}

	if mappedSecretName, ok := s.D.GetOkExists("name"); ok {
		tmp := mappedSecretName.(string)
		request.MappedSecretName = &tmp
	}

	if networkFirewallPolicyId, ok := s.D.GetOkExists("network_firewall_policy_id"); ok {
		tmp := networkFirewallPolicyId.(string)
		request.NetworkFirewallPolicyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "network_firewall")

	response, err := s.Client.GetMappedSecret(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicyMappedSecretDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("NetworkFirewallNetworkFirewallPolicyMappedSecretDataSource-", NetworkFirewallNetworkFirewallPolicyMappedSecretDataSource(), s.D))
	switch v := (s.Res.MappedSecret).(type) {
	case oci_network_firewall.VaultMappedSecret:
		s.D.Set("source", "OCI_VAULT")

		if v.VaultSecretId != nil {
			s.D.Set("vault_secret_id", *v.VaultSecretId)
		}

		if v.VersionNumber != nil {
			s.D.Set("version_number", *v.VersionNumber)
		}

		if v.Name != nil {
			s.D.Set("name", *v.Name)
		}

		if v.ParentResourceId != nil {
			s.D.Set("parent_resource_id", *v.ParentResourceId)
		}

		s.D.Set("type", v.Type)
	default:
		log.Printf("[WARN] Received 'source' of unknown type %v", s.Res.MappedSecret)
		return nil
	}

	return nil
}
