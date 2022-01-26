// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v56/core"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func CoreIpsecConnectionTunnelRoutesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreIpsecConnectionTunnelRoutes,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"advertiser": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipsec_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tunnel_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tunnel_routes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"advertiser": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"age": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"as_path": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
						},
						"is_best_path": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"prefix": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readCoreIpsecConnectionTunnelRoutes(d *schema.ResourceData, m interface{}) error {
	sync := &CoreIpsecConnectionTunnelRoutesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

type CoreIpsecConnectionTunnelRoutesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.ListIPSecConnectionTunnelRoutesResponse
}

func (s *CoreIpsecConnectionTunnelRoutesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreIpsecConnectionTunnelRoutesDataSourceCrud) Get() error {
	request := oci_core.ListIPSecConnectionTunnelRoutesRequest{}

	if advertiser, ok := s.D.GetOkExists("advertiser"); ok {
		request.Advertiser = oci_core.TunnelRouteSummaryAdvertiserEnum(advertiser.(string))
	}

	if ipsecId, ok := s.D.GetOkExists("ipsec_id"); ok {
		tmp := ipsecId.(string)
		request.IpscId = &tmp
	}

	if tunnelId, ok := s.D.GetOkExists("tunnel_id"); ok {
		tmp := tunnelId.(string)
		request.TunnelId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListIPSecConnectionTunnelRoutes(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListIPSecConnectionTunnelRoutes(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreIpsecConnectionTunnelRoutesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreIpsecConnectionTunnelRoutesDataSource-", CoreIpsecConnectionTunnelRoutesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		ipsecConnectionTunnelRoute := map[string]interface{}{}

		ipsecConnectionTunnelRoute["advertiser"] = r.Advertiser

		if r.Age != nil {
			ipsecConnectionTunnelRoute["age"] = strconv.FormatInt(*r.Age, 10)
		}

		ipsecConnectionTunnelRoute["as_path"] = r.AsPath

		if r.IsBestPath != nil {
			ipsecConnectionTunnelRoute["is_best_path"] = *r.IsBestPath
		}

		if r.Prefix != nil {
			ipsecConnectionTunnelRoute["prefix"] = *r.Prefix
		}

		resources = append(resources, ipsecConnectionTunnelRoute)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreIpsecConnectionTunnelRoutesDataSource().Schema["tunnel_routes"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("tunnel_routes", resources); err != nil {
		return err
	}

	return nil
}
