// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v28/core"
)

func init() {
	RegisterResource("oci_core_ipsec_connection_tunnel_management", CoreIpSecConnectionTunnelManagementResource())
}

func CoreIpSecConnectionTunnelManagementResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: DefaultTimeout,
		Create:   createCoreIpSecConnectionTunnelManagement,
		Read:     readCoreIpSecConnectionTunnelManagement,
		Update:   updateCoreIpSecConnectionTunnelManagement,
		Delete:   deleteCoreIpSecConnectionTunnelManagement,
		Schema: map[string]*schema.Schema{
			"ipsec_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tunnel_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"routing": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					string(oci_core.UpdateIpSecConnectionTunnelDetailsRoutingBgp),
					string(oci_core.UpdateIpSecConnectionTunnelDetailsRoutingStatic),
				}, true),
			},
			// Optional
			"bgp_session_info": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"customer_bgp_asn": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"customer_interface_ip": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"oracle_interface_ip": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
						"bgp_state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"oracle_bgp_asn": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ike_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ValidateFunc: validation.StringInSlice([]string{
					string(oci_core.UpdateIpSecConnectionTunnelDetailsIkeVersionV1),
					string(oci_core.UpdateIpSecConnectionTunnelDetailsIkeVersionV2),
				}, true),
			},
			"shared_secret": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateNotEmptyString(),
			},
			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cpe_ip": {
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

func createCoreIpSecConnectionTunnelManagement(d *schema.ResourceData, m interface{}) error {
	sync := &CoreIpSecConnectionTunnelManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient()
	return CreateResource(d, sync)
}

func readCoreIpSecConnectionTunnelManagement(d *schema.ResourceData, m interface{}) error {
	sync := &CoreIpSecConnectionTunnelManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient()
	return ReadResource(sync)
}

func updateCoreIpSecConnectionTunnelManagement(d *schema.ResourceData, m interface{}) error {
	sync := &CoreIpSecConnectionTunnelManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient()
	return UpdateResource(d, sync)
}

func deleteCoreIpSecConnectionTunnelManagement(d *schema.ResourceData, m interface{}) error {
	sync := &CoreIpSecConnectionTunnelManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient()
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type CoreIpSecConnectionTunnelManagementResourceCrud struct {
	BaseCrud
	Client                 *oci_core.VirtualNetworkClient
	Res                    *oci_core.IpSecConnectionTunnel
	ResSecret              *oci_core.IpSecConnectionTunnelSharedSecret
	DisableNotFoundRetries bool
}

func (s *CoreIpSecConnectionTunnelManagementResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CoreIpSecConnectionTunnelManagementResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.IpSecConnectionTunnelLifecycleStateProvisioning),
	}
}

func (s *CoreIpSecConnectionTunnelManagementResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.IpSecConnectionTunnelLifecycleStateAvailable),
	}
}

func (s *CoreIpSecConnectionTunnelManagementResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_core.IpSecConnectionTunnelLifecycleStateProvisioning),
	}
}

func (s *CoreIpSecConnectionTunnelManagementResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_core.IpSecConnectionTunnelLifecycleStateAvailable),
	}
}

func (s *CoreIpSecConnectionTunnelManagementResourceCrud) Create() error {
	return s.Update()
}

func (s *CoreIpSecConnectionTunnelManagementResourceCrud) Get() error {
	request := oci_core.GetIPSecConnectionTunnelRequest{}

	if ipsecId, ok := s.D.GetOkExists("ipsec_id"); ok {
		tmp := ipsecId.(string)
		request.IpscId = &tmp
	}

	if tunnelId, ok := s.D.GetOkExists("tunnel_id"); ok {
		tmp := tunnelId.(string)
		request.TunnelId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := s.Client.GetIPSecConnectionTunnel(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.IpSecConnectionTunnel

	secretRequest := oci_core.GetIPSecConnectionTunnelSharedSecretRequest{}

	secretRequest.IpscId = request.IpscId

	secretRequest.TunnelId = request.TunnelId

	secretResponse, err := s.Client.GetIPSecConnectionTunnelSharedSecret(context.Background(), secretRequest)
	if err != nil {
		return err
	}

	s.ResSecret = &secretResponse.IpSecConnectionTunnelSharedSecret

	return nil
}

func (s *CoreIpSecConnectionTunnelManagementResourceCrud) Update() error {
	if s.D.HasChange("shared_secret") {
		if sharedSecret, ok := s.D.GetOkExists("shared_secret"); ok {
			secretUpdateRequest := oci_core.UpdateIPSecConnectionTunnelSharedSecretRequest{}
			tmp := sharedSecret.(string)
			secretUpdateRequest.SharedSecret = &tmp

			if ipscId, ok := s.D.GetOkExists("ipsec_id"); ok {
				tmp := ipscId.(string)
				secretUpdateRequest.IpscId = &tmp
			}

			if tunnelId, ok := s.D.GetOkExists("tunnel_id"); ok {
				tmp := tunnelId.(string)
				secretUpdateRequest.TunnelId = &tmp
			}
			_, err := s.Client.UpdateIPSecConnectionTunnelSharedSecret(context.Background(), secretUpdateRequest)
			if err != nil {
				return err
			}

			retentionPolicyFunc := func() bool { return s.Res.LifecycleState == oci_core.IpSecConnectionTunnelLifecycleStateAvailable }
			if conditionErr := WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate)); conditionErr != nil {
				return conditionErr
			}
		}
	}

	request := oci_core.UpdateIPSecConnectionTunnelRequest{}

	if ipscId, ok := s.D.GetOkExists("ipsec_id"); ok {
		tmp := ipscId.(string)
		request.IpscId = &tmp
	}

	if tunnelId, ok := s.D.GetOkExists("tunnel_id"); ok {
		tmp := tunnelId.(string)
		request.TunnelId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if ikeVersion, ok := s.D.GetOkExists("ike_version"); ok {
		request.IkeVersion = oci_core.UpdateIpSecConnectionTunnelDetailsIkeVersionEnum(ikeVersion.(string))
	}

	if routing, ok := s.D.GetOkExists("routing"); ok {
		request.Routing = oci_core.UpdateIpSecConnectionTunnelDetailsRoutingEnum(routing.(string))
	}

	if request.Routing == oci_core.UpdateIpSecConnectionTunnelDetailsRoutingBgp {
		if _, ok := s.D.GetOkExists("bgp_session_info"); ok {
			BgpSessionDeatails := &oci_core.UpdateIpSecTunnelBgpSessionDetails{}
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "bgp_session_info", 0)
			if customerBgpAsn, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "customer_bgp_asn")); ok {
				tmp := customerBgpAsn.(string)
				BgpSessionDeatails.CustomerBgpAsn = &tmp
			}

			if customerInterfaceIp, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "customer_interface_ip")); ok {
				tmp := customerInterfaceIp.(string)
				BgpSessionDeatails.CustomerInterfaceIp = &tmp
			}

			if oracleInterfaceIp, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "oracle_interface_ip")); ok {
				tmp := oracleInterfaceIp.(string)
				BgpSessionDeatails.OracleInterfaceIp = &tmp
			}

			request.BgpSessionConfig = BgpSessionDeatails
		}
	}
	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")
	response, err := s.Client.UpdateIPSecConnectionTunnel(context.Background(), request)
	if err != nil {
		return err
	}
	s.Res = &response.IpSecConnectionTunnel
	return nil
}

func (s *CoreIpSecConnectionTunnelManagementResourceCrud) Delete() error {
	return nil
}

func (s *CoreIpSecConnectionTunnelManagementResourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.BgpSessionInfo != nil {
		s.D.Set("bgp_session_info", []interface{}{BgpSessionInfoToMap(s.Res.BgpSessionInfo)})
	} else {
		if _, ok := s.D.GetOkExists("bgp_session_info"); !ok {
			s.D.Set("bgp_session_info", nil)
		}
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

	if s.ResSecret != nil && s.ResSecret.SharedSecret != nil {
		s.D.Set("shared_secret", *(s.ResSecret.SharedSecret))
	}

	return nil
}
