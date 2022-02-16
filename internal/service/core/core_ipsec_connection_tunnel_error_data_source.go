// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v58/core"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func CoreIpsecConnectionTunnelErrorDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularCoreIpsecConnectionTunnelError,
		Schema: map[string]*schema.Schema{
			"ipsec_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tunnel_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"error_code": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"error_description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"oci_resources_link": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"solution": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"timestamp": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularCoreIpsecConnectionTunnelError(d *schema.ResourceData, m interface{}) error {
	sync := &CoreIpsecConnectionTunnelErrorDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

type CoreIpsecConnectionTunnelErrorDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.GetIPSecConnectionTunnelErrorResponse
}

func (s *CoreIpsecConnectionTunnelErrorDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreIpsecConnectionTunnelErrorDataSourceCrud) Get() error {
	request := oci_core.GetIPSecConnectionTunnelErrorRequest{}

	if ipsecId, ok := s.D.GetOkExists("ipsec_id"); ok {
		tmp := ipsecId.(string)
		request.IpscId = &tmp
	}

	if tunnelId, ok := s.D.GetOkExists("tunnel_id"); ok {
		tmp := tunnelId.(string)
		request.TunnelId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.GetIPSecConnectionTunnelError(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CoreIpsecConnectionTunnelErrorDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.ErrorCode != nil {
		s.D.Set("error_code", *s.Res.ErrorCode)
	}

	if s.Res.ErrorDescription != nil {
		s.D.Set("error_description", *s.Res.ErrorDescription)
	}

	if s.Res.OciResourcesLink != nil {
		s.D.Set("oci_resources_link", *s.Res.OciResourcesLink)
	}

	if s.Res.Solution != nil {
		s.D.Set("solution", *s.Res.Solution)
	}

	if s.Res.Timestamp != nil {
		s.D.Set("timestamp", s.Res.Timestamp.String())
	}

	return nil
}
