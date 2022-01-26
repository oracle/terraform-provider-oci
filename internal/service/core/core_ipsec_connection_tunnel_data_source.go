// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v56/core"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func CoreIpSecConnectionTunnelDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularCoreIpSecConnectionTunnel,
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
			"bgp_session_info": {
				Type:     schema.TypeList,
				Computed: true,
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
			"encryption_domain_config": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"cpe_traffic_selector": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"oracle_traffic_selector": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
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
	}
}

func readSingularCoreIpSecConnectionTunnel(d *schema.ResourceData, m interface{}) error {
	sync := &CoreIpSecConnectionTunnelDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()
	return tfresource.ReadResource(sync)
}

type CoreIpSecConnectionTunnelDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.GetIPSecConnectionTunnelResponse
}

func (s *CoreIpSecConnectionTunnelDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreIpSecConnectionTunnelDataSourceCrud) Get() error {
	request := oci_core.GetIPSecConnectionTunnelRequest{}

	if ipsecId, ok := s.D.GetOkExists("ipsec_id"); ok {
		tmp := ipsecId.(string)
		request.IpscId = &tmp
	}

	if tunnelId, ok := s.D.GetOkExists("tunnel_id"); ok {
		tmp := tunnelId.(string)
		request.TunnelId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.GetIPSecConnectionTunnel(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CoreIpSecConnectionTunnelDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.BgpSessionInfo != nil {
		s.D.Set("bgp_session_info", []interface{}{BgpSessionInfoToMap(s.Res.BgpSessionInfo)})
	} else {
		s.D.Set("bgp_session_info", nil)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CpeIp != nil {
		s.D.Set("cpe_ip", *s.Res.CpeIp)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.EncryptionDomainConfig != nil {
		s.D.Set("encryption_domain_config", []interface{}{EncryptionDomainConfigToMap(s.Res.EncryptionDomainConfig)})
	} else {
		s.D.Set("encryption_domain_config", nil)
	}

	s.D.Set("ike_version", s.Res.IkeVersion)

	s.D.Set("routing", s.Res.Routing)

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("status", s.Res.Status)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeStatusUpdated != nil {
		s.D.Set("time_status_updated", s.Res.TimeStatusUpdated.String())
	}

	if s.Res.VpnIp != nil {
		s.D.Set("vpn_ip", *s.Res.VpnIp)
	}

	return nil
}

func BgpSessionInfoToMap(obj *oci_core.BgpSessionInfo) map[string]interface{} {
	result := map[string]interface{}{}

	result["bgp_ipv6state"] = string(obj.BgpIpv6State)

	result["bgp_state"] = string(obj.BgpState)

	if obj.CustomerBgpAsn != nil {
		result["customer_bgp_asn"] = string(*obj.CustomerBgpAsn)
	}

	if obj.CustomerInterfaceIp != nil {
		result["customer_interface_ip"] = string(*obj.CustomerInterfaceIp)
	}

	if obj.CustomerInterfaceIpv6 != nil {
		result["customer_interface_ipv6"] = string(*obj.CustomerInterfaceIpv6)
	}

	if obj.OracleBgpAsn != nil {
		result["oracle_bgp_asn"] = string(*obj.OracleBgpAsn)
	}

	if obj.OracleInterfaceIp != nil {
		result["oracle_interface_ip"] = string(*obj.OracleInterfaceIp)
	}

	if obj.OracleInterfaceIpv6 != nil {
		result["oracle_interface_ipv6"] = string(*obj.OracleInterfaceIpv6)
	}

	return result
}

func UpdateIPSecTunnelBgpSessionDetailsToMap(obj *oci_core.UpdateIpSecTunnelBgpSessionDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CustomerBgpAsn != nil {
		result["customer_bgp_asn"] = string(*obj.CustomerBgpAsn)
	}

	if obj.CustomerInterfaceIp != nil {
		result["customer_interface_ip"] = string(*obj.CustomerInterfaceIp)
	}

	if obj.OracleInterfaceIp != nil {
		result["oracle_interface_ip"] = string(*obj.OracleInterfaceIp)
	}

	return result
}
