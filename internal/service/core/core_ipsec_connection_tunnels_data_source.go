// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CoreIpSecConnectionTunnelsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreIpSecConnectionTunnels,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"ipsec_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ip_sec_connection_tunnels": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(CoreIpSecConnectionTunnelManagementResource()),
			},
		},
	}
}

func readCoreIpSecConnectionTunnels(d *schema.ResourceData, m interface{}) error {
	sync := &CoreIpSecConnectionTunnelsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

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

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreIpSecConnectionTunnelsDataSource-", CoreIpSecConnectionTunnelsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		ipSecConnectionTunnel := map[string]interface{}{}

		ipSecConnectionTunnel["associated_virtual_circuits"] = r.AssociatedVirtualCircuits
		ipSecConnectionTunnel["associated_virtual_circuits"] = r.AssociatedVirtualCircuits

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

		ipSecConnectionTunnel["dpd_mode"] = r.DpdMode

		if r.DpdTimeoutInSec != nil {
			ipSecConnectionTunnel["dpd_timeout_in_sec"] = *r.DpdTimeoutInSec
		}

		if r.EncryptionDomainConfig != nil {
			ipSecConnectionTunnel["encryption_domain_config"] = []interface{}{EncryptionDomainConfigToMap(r.EncryptionDomainConfig)}
		} else {
			ipSecConnectionTunnel["encryption_domain_config"] = nil
		}

		if r.Id != nil {
			ipSecConnectionTunnel["id"] = *r.Id
		}

		ipSecConnectionTunnel["ike_version"] = r.IkeVersion

		ipSecConnectionTunnel["nat_translation_enabled"] = r.NatTranslationEnabled

		ipSecConnectionTunnel["oracle_can_initiate"] = r.OracleCanInitiate

		if r.PhaseOneDetails != nil {
			ipSecConnectionTunnel["phase_one_details"] = []interface{}{TunnelPhaseOneDetailsToMap(r.PhaseOneDetails)}
		} else {
			ipSecConnectionTunnel["phase_one_details"] = nil
		}

		if r.PhaseTwoDetails != nil {
			ipSecConnectionTunnel["phase_two_details"] = []interface{}{TunnelPhaseTwoDetailsToMap(r.PhaseTwoDetails)}
		} else {
			ipSecConnectionTunnel["phase_two_details"] = nil
		}

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
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreIpSecConnectionTunnelsDataSource().Schema["ip_sec_connection_tunnels"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("ip_sec_connection_tunnels", resources); err != nil {
		return err
	}

	return nil
}

func EncryptionDomainConfigToMap(obj *oci_core.EncryptionDomainConfig) map[string]interface{} {
	result := map[string]interface{}{}

	result["cpe_traffic_selector"] = obj.CpeTrafficSelector

	result["oracle_traffic_selector"] = obj.OracleTrafficSelector

	return result
}

func TunnelPhaseOneDetailsToMap(obj *oci_core.TunnelPhaseOneDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CustomAuthenticationAlgorithm != nil {
		result["custom_authentication_algorithm"] = string(*obj.CustomAuthenticationAlgorithm)
	}

	if obj.CustomDhGroup != nil {
		result["custom_dh_group"] = string(*obj.CustomDhGroup)
	}

	if obj.CustomEncryptionAlgorithm != nil {
		result["custom_encryption_algorithm"] = string(*obj.CustomEncryptionAlgorithm)
	}

	if obj.IsCustomPhaseOneConfig != nil {
		result["is_custom_phase_one_config"] = bool(*obj.IsCustomPhaseOneConfig)
	}

	if obj.IsIkeEstablished != nil {
		result["is_ike_established"] = bool(*obj.IsIkeEstablished)
	}

	if obj.Lifetime != nil {
		result["lifetime"] = int(*obj.Lifetime)
	}

	if obj.NegotiatedAuthenticationAlgorithm != nil {
		result["negotiated_authentication_algorithm"] = string(*obj.NegotiatedAuthenticationAlgorithm)
	}

	if obj.NegotiatedDhGroup != nil {
		result["negotiated_dh_group"] = string(*obj.NegotiatedDhGroup)
	}

	if obj.NegotiatedEncryptionAlgorithm != nil {
		result["negotiated_encryption_algorithm"] = string(*obj.NegotiatedEncryptionAlgorithm)
	}

	if obj.RemainingLifetime != nil {
		result["remaining_lifetime"] = strconv.FormatInt(*obj.RemainingLifetime, 10)
		result["remaining_lifetime_int"] = int(*obj.RemainingLifetime)
	}

	if obj.RemainingLifetimeLastRetrieved != nil {
		result["remaining_lifetime_last_retrieved"] = obj.RemainingLifetimeLastRetrieved.String()
	}

	return result
}

func TunnelPhaseTwoDetailsToMap(obj *oci_core.TunnelPhaseTwoDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CustomAuthenticationAlgorithm != nil {
		result["custom_authentication_algorithm"] = string(*obj.CustomAuthenticationAlgorithm)
	}

	if obj.CustomEncryptionAlgorithm != nil {
		result["custom_encryption_algorithm"] = string(*obj.CustomEncryptionAlgorithm)
	}

	if obj.DhGroup != nil {
		result["dh_group"] = string(*obj.DhGroup)
	}

	if obj.IsCustomPhaseTwoConfig != nil {
		result["is_custom_phase_two_config"] = bool(*obj.IsCustomPhaseTwoConfig)
	}

	if obj.IsEspEstablished != nil {
		result["is_esp_established"] = bool(*obj.IsEspEstablished)
	}

	if obj.IsPfsEnabled != nil {
		result["is_pfs_enabled"] = bool(*obj.IsPfsEnabled)
	}

	if obj.Lifetime != nil {
		result["lifetime"] = int(*obj.Lifetime)
	}

	if obj.NegotiatedAuthenticationAlgorithm != nil {
		result["negotiated_authentication_algorithm"] = string(*obj.NegotiatedAuthenticationAlgorithm)
	}

	if obj.NegotiatedDhGroup != nil {
		result["negotiated_dh_group"] = string(*obj.NegotiatedDhGroup)
	}

	if obj.NegotiatedEncryptionAlgorithm != nil {
		result["negotiated_encryption_algorithm"] = string(*obj.NegotiatedEncryptionAlgorithm)
	}

	if obj.RemainingLifetime != nil {
		result["remaining_lifetime"] = strconv.FormatInt(*obj.RemainingLifetime, 10)
		result["remaining_lifetime_int"] = int(*obj.RemainingLifetime)
	}

	if obj.RemainingLifetimeLastRetrieved != nil {
		result["remaining_lifetime_last_retrieved"] = obj.RemainingLifetimeLastRetrieved.String()
	}

	return result
}
