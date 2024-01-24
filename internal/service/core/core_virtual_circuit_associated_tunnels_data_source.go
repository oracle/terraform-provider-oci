// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CoreVirtualCircuitAssociatedTunnelsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreVirtualCircuitAssociatedTunnels,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"virtual_circuit_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"virtual_circuit_associated_tunnel_details": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"ipsec_connection_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"tunnel_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"tunnel_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readCoreVirtualCircuitAssociatedTunnels(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVirtualCircuitAssociatedTunnelsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

type CoreVirtualCircuitAssociatedTunnelsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.ListVirtualCircuitAssociatedTunnelsResponse
}

func (s *CoreVirtualCircuitAssociatedTunnelsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreVirtualCircuitAssociatedTunnelsDataSourceCrud) Get() error {
	request := oci_core.ListVirtualCircuitAssociatedTunnelsRequest{}

	if virtualCircuitId, ok := s.D.GetOkExists("virtual_circuit_id"); ok {
		tmp := virtualCircuitId.(string)
		request.VirtualCircuitId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListVirtualCircuitAssociatedTunnels(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListVirtualCircuitAssociatedTunnels(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreVirtualCircuitAssociatedTunnelsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreVirtualCircuitAssociatedTunnelsDataSource-", CoreVirtualCircuitAssociatedTunnelsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		virtualCircuitAssociatedTunnel := map[string]interface{}{}

		if r.IpsecConnectionId != nil {
			virtualCircuitAssociatedTunnel["ipsec_connection_id"] = *r.IpsecConnectionId
		}

		if r.TunnelId != nil {
			virtualCircuitAssociatedTunnel["tunnel_id"] = *r.TunnelId
		}

		virtualCircuitAssociatedTunnel["tunnel_type"] = r.TunnelType

		resources = append(resources, virtualCircuitAssociatedTunnel)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreVirtualCircuitAssociatedTunnelsDataSource().Schema["virtual_circuit_associated_tunnel_details"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("virtual_circuit_associated_tunnel_details", resources); err != nil {
		return err
	}

	return nil
}
