// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"fmt"

	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v58/core"
)

func CoreIpSecConnectionTunnelManagementResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
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
					string(oci_core.UpdateIpSecConnectionTunnelDetailsRoutingPolicy),
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
			"encryption_domain_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"cpe_traffic_selector": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"oracle_traffic_selector": {
							Type:     schema.TypeList,
							Computed: true,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dpd_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dpd_timeout_in_sec": {
				Type:     schema.TypeInt,
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
			"nat_translation_enabled": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"oracle_can_initiate": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"phase_one_details": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"custom_authentication_algorithm": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"custom_dh_group": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"custom_encryption_algorithm": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_custom_phase_one_config": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_ike_established": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"lifetime": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"negotiated_authentication_algorithm": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"negotiated_dh_group": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"negotiated_encryption_algorithm": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"remaining_lifetime": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"remaining_lifetime_last_retrieved": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"phase_two_details": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"custom_authentication_algorithm": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"custom_encryption_algorithm": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"dh_group": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_custom_phase_two_config": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_esp_established": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_pfs_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"lifetime": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"negotiated_authentication_algorithm": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"negotiated_dh_group": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"negotiated_encryption_algorithm": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"remaining_lifetime": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"remaining_lifetime_last_retrieved": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"shared_secret": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: utils.ValidateNotEmptyString(),
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
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()
	return tfresource.CreateResource(d, sync)
}

func readCoreIpSecConnectionTunnelManagement(d *schema.ResourceData, m interface{}) error {
	sync := &CoreIpSecConnectionTunnelManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()
	return tfresource.ReadResource(sync)
}

func updateCoreIpSecConnectionTunnelManagement(d *schema.ResourceData, m interface{}) error {
	sync := &CoreIpSecConnectionTunnelManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()
	return tfresource.UpdateResource(d, sync)
}

func deleteCoreIpSecConnectionTunnelManagement(d *schema.ResourceData, m interface{}) error {
	sync := &CoreIpSecConnectionTunnelManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type CoreIpSecConnectionTunnelManagementResourceCrud struct {
	tfresource.BaseCrud
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

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
			if conditionErr := tfresource.WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate)); conditionErr != nil {
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

	if natTranslation, ok := s.D.GetOkExists("nat_translation_enabled"); ok {
		request.NatTranslationEnabled = oci_core.UpdateIpSecConnectionTunnelDetailsNatTranslationEnabledEnum(natTranslation.(string))
	}

	if oracleInitiation, ok := s.D.GetOkExists("oracle_can_initiate"); ok {
		request.OracleInitiation = oci_core.UpdateIpSecConnectionTunnelDetailsOracleInitiationEnum(oracleInitiation.(string))
	}

	dpdConfig := &oci_core.DpdConfig{}
	if dpdMode, ok := s.D.GetOkExists("dpd_mode"); ok {
		dpdConfig.DpdMode = oci_core.DpdConfigDpdModeEnum(dpdMode.(string))
	}

	if dpdTimeout, ok := s.D.GetOkExists("dpd_timeout_in_sec"); ok {
		tmp := dpdTimeout.(int)
		dpdConfig.DpdTimeoutInSec = &tmp
	}
	request.DpdConfig = dpdConfig

	if _, ok := s.D.GetOkExists("phase_one_details"); ok {
		phaseOneDetails := &oci_core.PhaseOneConfigDetails{}
		fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "phase_one_details", 0)
		if customAuthenticationAlgorithm, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "custom_authentication_algorithm")); ok {
			phaseOneDetails.AuthenticationAlgorithm = oci_core.PhaseOneConfigDetailsAuthenticationAlgorithmEnum(customAuthenticationAlgorithm.(string))
		}

		if customDhGroup, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "custom_dh_group")); ok {
			phaseOneDetails.DiffieHelmanGroup = oci_core.PhaseOneConfigDetailsDiffieHelmanGroupEnum(customDhGroup.(string))
		}

		if customEncryptionAlgorithm, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "custom_encryption_algorithm")); ok {
			phaseOneDetails.EncryptionAlgorithm = oci_core.PhaseOneConfigDetailsEncryptionAlgorithmEnum(customEncryptionAlgorithm.(string))
		}

		if isCustomPhaseOneConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_custom_phase_one_config")); ok {
			tmp := isCustomPhaseOneConfig.(bool)
			phaseOneDetails.IsCustomPhaseOneConfig = &tmp
		}

		if lifetime, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "lifetime")); ok {
			tmp := lifetime.(int)
			phaseOneDetails.LifetimeInSeconds = &tmp
		}

		request.PhaseOneConfig = phaseOneDetails
	}

	if _, ok := s.D.GetOkExists("phase_two_details"); ok {
		phaseTwoDetails := &oci_core.PhaseTwoConfigDetails{}
		fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "phase_two_details", 0)
		if customAuthenticationAlgorithm, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "custom_authentication_algorithm")); ok {
			phaseTwoDetails.AuthenticationAlgorithm = oci_core.PhaseTwoConfigDetailsAuthenticationAlgorithmEnum(customAuthenticationAlgorithm.(string))
		}

		if customDhGroup, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "dh_group")); ok {
			phaseTwoDetails.PfsDhGroup = oci_core.PhaseTwoConfigDetailsPfsDhGroupEnum(customDhGroup.(string))
		}

		if customEncryptionAlgorithm, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "custom_encryption_algorithm")); ok {
			phaseTwoDetails.EncryptionAlgorithm = oci_core.PhaseTwoConfigDetailsEncryptionAlgorithmEnum(customEncryptionAlgorithm.(string))
		}

		if isCustomPhaseTwoConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_custom_phase_two_config")); ok {
			tmp := isCustomPhaseTwoConfig.(bool)
			phaseTwoDetails.IsCustomPhaseTwoConfig = &tmp
		}

		if lifetime, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "lifetime")); ok {
			tmp := lifetime.(int)
			phaseTwoDetails.LifetimeInSeconds = &tmp
		}

		if isPfsEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_pfs_enabled")); ok {
			tmp := isPfsEnabled.(bool)
			phaseTwoDetails.IsPfsEnabled = &tmp
		}

		request.PhaseTwoConfig = phaseTwoDetails
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

	if request.Routing == oci_core.UpdateIpSecConnectionTunnelDetailsRoutingPolicy {
		if _, ok := s.D.GetOkExists("encryption_domain_config"); ok {
			EncryptionDomainDetails := &oci_core.UpdateIpSecTunnelEncryptionDomainDetails{}
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "encryption_domain_config", 0)
			if oracleTrafficSelector, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "oracle_traffic_selector")); ok {
				interfaces := oracleTrafficSelector.([]interface{})
				tmp := make([]string, len(interfaces))
				for i := range interfaces {
					if interfaces[i] != nil {
						tmp[i] = interfaces[i].(string)
					}
				}
				if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "oracle_traffic_selector")) {
					EncryptionDomainDetails.OracleTrafficSelector = tmp
				}
			}

			if cpeTrafficSelector, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cpe_traffic_selector")); ok {
				interfaces := cpeTrafficSelector.([]interface{})
				tmp := make([]string, len(interfaces))
				for i := range interfaces {
					if interfaces[i] != nil {
						tmp[i] = interfaces[i].(string)
					}
				}
				if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "cpe_traffic_selector")) {
					EncryptionDomainDetails.CpeTrafficSelector = tmp
				}
			}

			request.EncryptionDomainConfig = EncryptionDomainDetails
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")
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

	if s.Res.EncryptionDomainConfig != nil {
		s.D.Set("encryption_domain_config", []interface{}{EncryptionDomainConfigToMap(s.Res.EncryptionDomainConfig)})
	} else {
		if _, ok := s.D.GetOkExists("encryption_domain_config"); !ok {
			s.D.Set("encryption_domain_config", nil)
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

	s.D.Set("dpd_mode", s.Res.DpdMode)

	if s.Res.DpdTimeoutInSec != nil {
		s.D.Set("dpd_timeout_in_sec", *s.Res.DpdTimeoutInSec)
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
