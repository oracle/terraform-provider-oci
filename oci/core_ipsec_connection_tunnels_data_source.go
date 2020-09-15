// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v25/core"
)

func init() {
	RegisterDatasource("oci_core_ipsec_connection_tunnels", CoreIpSecConnectionTunnelsDataSource())
}

func CoreIpSecConnectionTunnelsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreIpSecConnectionTunnels,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"ipsec_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ip_sec_connection_tunnels": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"bgp_session_info": {
							Type:     schema.TypeList,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"bgp_state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"customer_bgp_asn": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"customer_interface_ip": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"oracle_bgp_asn": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"oracle_interface_ip": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"compartment_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"cpe_ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ike_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"routing": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_status_updated": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"vpn_ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readCoreIpSecConnectionTunnels(d *schema.ResourceData, m interface{}) error {
	sync := &CoreIpSecConnectionTunnelsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient()

	return ReadResource(sync)
}

type CoreIpSecConnectionTunnelsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.ListIPSecConnectionTunnelsResponse
}

func (s *CoreIpSecConnectionTunnelsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreIpSecConnectionTunnelsDataSourceCrud) Get() error {
	request := oci_core.ListIPSecConnectionTunnelsRequest{}

	if ipsecId, ok := s.D.GetOkExists("ipsec_id"); ok {
		tmp := ipsecId.(string)
		request.IpscId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := s.Client.ListIPSecConnectionTunnels(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListIPSecConnectionTunnels(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreIpSecConnectionTunnelsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		ipSecConnectionTunnel := map[string]interface{}{}

		if r.BgpSessionInfo != nil {
			ipSecConnectionTunnel["bgp_session_info"] = []interface{}{BgpSessionInfoToMap(r.BgpSessionInfo)}
		} else {
			ipSecConnectionTunnel["bgp_session_info"] = nil
		}

		if r.CompartmentId != nil {
			ipSecConnectionTunnel["compartment_id"] = *r.CompartmentId
		}

		if r.CpeIp != nil {
			ipSecConnectionTunnel["cpe_ip"] = *r.CpeIp
		}

		if r.DisplayName != nil {
			ipSecConnectionTunnel["display_name"] = *r.DisplayName
		}

		if r.Id != nil {
			ipSecConnectionTunnel["id"] = *r.Id
		}

		ipSecConnectionTunnel["ike_version"] = r.IkeVersion

		ipSecConnectionTunnel["routing"] = r.Routing

		ipSecConnectionTunnel["state"] = r.LifecycleState

		ipSecConnectionTunnel["status"] = r.Status

		if r.TimeCreated != nil {
			ipSecConnectionTunnel["time_created"] = r.TimeCreated.String()
		}

		if r.TimeStatusUpdated != nil {
			ipSecConnectionTunnel["time_status_updated"] = r.TimeStatusUpdated.String()
		}

		if r.VpnIp != nil {
			ipSecConnectionTunnel["vpn_ip"] = *r.VpnIp
		}

		resources = append(resources, ipSecConnectionTunnel)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, CoreIpSecConnectionTunnelsDataSource().Schema["ip_sec_connection_tunnels"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("ip_sec_connection_tunnels", resources); err != nil {
		return err
	}

	return nil
}
