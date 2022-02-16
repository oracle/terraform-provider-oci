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

func CoreTunnelSecurityAssociationsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreTunnelSecurityAssociations,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"ipsec_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tunnel_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tunnel_security_associations": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"cpe_subnet": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"oracle_subnet": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"tunnel_sa_error_info": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"tunnel_sa_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readCoreTunnelSecurityAssociations(d *schema.ResourceData, m interface{}) error {
	sync := &CoreTunnelSecurityAssociationsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

type CoreTunnelSecurityAssociationsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.ListIPSecConnectionTunnelSecurityAssociationsResponse
}

func (s *CoreTunnelSecurityAssociationsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreTunnelSecurityAssociationsDataSourceCrud) Get() error {
	request := oci_core.ListIPSecConnectionTunnelSecurityAssociationsRequest{}

	if ipsecId, ok := s.D.GetOkExists("ipsec_id"); ok {
		tmp := ipsecId.(string)
		request.IpscId = &tmp
	}

	if tunnelId, ok := s.D.GetOkExists("tunnel_id"); ok {
		tmp := tunnelId.(string)
		request.TunnelId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListIPSecConnectionTunnelSecurityAssociations(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListIPSecConnectionTunnelSecurityAssociations(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreTunnelSecurityAssociationsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreTunnelSecurityAssociationsDataSource-", CoreTunnelSecurityAssociationsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		tunnelSecurityAssociation := map[string]interface{}{}

		if r.CpeSubnet != nil {
			tunnelSecurityAssociation["cpe_subnet"] = *r.CpeSubnet
		}

		if r.OracleSubnet != nil {
			tunnelSecurityAssociation["oracle_subnet"] = *r.OracleSubnet
		}

		if r.Time != nil {
			tunnelSecurityAssociation["time"] = *r.Time
		}

		if r.TunnelSaErrorInfo != nil {
			tunnelSecurityAssociation["tunnel_sa_error_info"] = *r.TunnelSaErrorInfo
		}

		tunnelSecurityAssociation["tunnel_sa_status"] = r.TunnelSaStatus

		resources = append(resources, tunnelSecurityAssociation)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreTunnelSecurityAssociationsDataSource().Schema["tunnel_security_associations"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("tunnel_security_associations", resources); err != nil {
		return err
	}

	return nil
}
