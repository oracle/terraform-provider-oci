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

func CoreIpSecConnectionTunnelDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["tunnel_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["ipsec_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(CoreIpSecConnectionTunnelManagementResource(), fieldMap, readSingularCoreIpSecConnectionTunnel)
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

	s.D.Set("dpd_mode", s.Res.DpdMode)

	if s.Res.DpdTimeoutInSec != nil {
		s.D.Set("dpd_timeout_in_sec", s.Res.DpdTimeoutInSec)
	}

	s.D.Set("nat_translation_enabled", s.Res.NatTranslationEnabled)

	s.D.Set("oracle_can_initiate", s.Res.OracleCanInitiate)

	if s.Res.PhaseOneDetails != nil {
		s.D.Set("phase_one_details", []interface{}{TunnelPhaseOneDetailsToMap(s.Res.PhaseOneDetails)})
	} else {
		if _, ok := s.D.GetOkExists("phase_one_details"); !ok {
			s.D.Set("phase_one_details", nil)
		}
	}

	if s.Res.PhaseTwoDetails != nil {
		s.D.Set("phase_two_details", []interface{}{TunnelPhaseTwoDetailsToMap(s.Res.PhaseTwoDetails)})
	} else {
		if _, ok := s.D.GetOkExists("phase_two_details"); !ok {
			s.D.Set("phase_two_details", nil)
		}
	}

	return nil
}

func BgpSessionInfoToMap(obj *oci_core.BgpSessionInfo) map[string]interface{} {
	result := map[string]interface{}{}

	result["bgp_ipv6state"] = string(obj.BgpIpv6State)
	result["bgp_ipv6_state"] = string(obj.BgpIpv6State)

	result["bgp_state"] = string(obj.BgpState)

	if obj.CustomerBgpAsn != nil { // nil when static routing but still pass bgp ips
		result["oracle_bgp_asn"] = string(*obj.OracleBgpAsn)
	}

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
