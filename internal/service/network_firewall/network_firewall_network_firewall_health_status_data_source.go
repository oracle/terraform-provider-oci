// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package network_firewall

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_network_firewall "github.com/oracle/oci-go-sdk/v65/networkfirewall"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func NetworkFirewallNetworkFirewallHealthStatusDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readSingularNetworkFirewallNetworkFirewallHealthStatusWithContext,
		Schema: map[string]*schema.Schema{
			"network_firewall_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularNetworkFirewallNetworkFirewallHealthStatusWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &NetworkFirewallNetworkFirewallHealthStatusDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type NetworkFirewallNetworkFirewallHealthStatusDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_network_firewall.NetworkFirewallClient
	Res    *oci_network_firewall.GetNetworkFirewallHealthStatusResponse
}

func (s *NetworkFirewallNetworkFirewallHealthStatusDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *NetworkFirewallNetworkFirewallHealthStatusDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_network_firewall.GetNetworkFirewallHealthStatusRequest{}

	if networkFirewallId, ok := s.D.GetOkExists("network_firewall_id"); ok {
		tmp := networkFirewallId.(string)
		request.NetworkFirewallId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "network_firewall")

	response, err := s.Client.GetNetworkFirewallHealthStatus(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *NetworkFirewallNetworkFirewallHealthStatusDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("NetworkFirewallNetworkFirewallHealthStatusDataSource-", NetworkFirewallNetworkFirewallHealthStatusDataSource(), s.D))

	s.D.Set("status", s.Res.Status)

	return nil
}
