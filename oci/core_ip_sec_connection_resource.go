// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"bytes"
	"context"
	"fmt"

	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"

	oci_core "github.com/oracle/oci-go-sdk/core"
)

func CoreIpSecConnectionResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createCoreIpSecConnection,
		Read:     readCoreIpSecConnection,
		Update:   updateCoreIpSecConnection,
		Delete:   deleteCoreIpSecConnection,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cpe_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"drg_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"static_routes": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Optional
			"cpe_local_identifier": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cpe_local_identifier_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: definedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"tunnel_configuration": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 2,
				MinItems: 1,
				Set:      tunnelConfigurationHashCodeForSets,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"bgp_session_config": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"customer_bgp_asn": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"customer_interface_ip": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"oracle_interface_ip": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},

									// Computed
								},
							},
						},
						"display_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"routing": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"shared_secret": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},

			// Computed
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createCoreIpSecConnection(d *schema.ResourceData, m interface{}) error {
	sync := &CoreIpSecConnectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return CreateResource(d, sync)
}

func readCoreIpSecConnection(d *schema.ResourceData, m interface{}) error {
	sync := &CoreIpSecConnectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return ReadResource(sync)
}

func updateCoreIpSecConnection(d *schema.ResourceData, m interface{}) error {
	sync := &CoreIpSecConnectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return UpdateResource(d, sync)
}

func deleteCoreIpSecConnection(d *schema.ResourceData, m interface{}) error {
	sync := &CoreIpSecConnectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type CoreIpSecConnectionResourceCrud struct {
	BaseCrud
	Client                 *oci_core.VirtualNetworkClient
	Res                    *oci_core.IpSecConnection
	DisableNotFoundRetries bool
}

func (s *CoreIpSecConnectionResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CoreIpSecConnectionResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.IpSecConnectionLifecycleStateProvisioning),
	}
}

func (s *CoreIpSecConnectionResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.IpSecConnectionLifecycleStateAvailable),
	}
}

func (s *CoreIpSecConnectionResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.IpSecConnectionLifecycleStateTerminating),
	}
}

func (s *CoreIpSecConnectionResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.IpSecConnectionLifecycleStateTerminated),
	}
}

func (s *CoreIpSecConnectionResourceCrud) Create() error {
	request := oci_core.CreateIPSecConnectionRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if cpeId, ok := s.D.GetOkExists("cpe_id"); ok {
		tmp := cpeId.(string)
		request.CpeId = &tmp
	}

	if cpeLocalIdentifier, ok := s.D.GetOkExists("cpe_local_identifier"); ok {
		tmp := cpeLocalIdentifier.(string)
		request.CpeLocalIdentifier = &tmp
	}

	if cpeLocalIdentifierType, ok := s.D.GetOkExists("cpe_local_identifier_type"); ok {
		request.CpeLocalIdentifierType = oci_core.CreateIpSecConnectionDetailsCpeLocalIdentifierTypeEnum(cpeLocalIdentifierType.(string))
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if drgId, ok := s.D.GetOkExists("drg_id"); ok {
		tmp := drgId.(string)
		request.DrgId = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.StaticRoutes = []string{}
	if staticRoutes, ok := s.D.GetOkExists("static_routes"); ok {
		interfaces := staticRoutes.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		request.StaticRoutes = tmp
	}

	if tunnelConfiguration, ok := s.D.GetOkExists("tunnel_configuration"); ok {
		set := tunnelConfiguration.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_core.CreateIpSecConnectionTunnelDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := tunnelConfigurationHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "tunnel_configuration", stateDataIndex)
			converted, err := s.mapToCreateIPSecConnectionTunnelDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		request.TunnelConfiguration = tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateIPSecConnection(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.IpSecConnection
	return nil
}

func (s *CoreIpSecConnectionResourceCrud) Get() error {
	request := oci_core.GetIPSecConnectionRequest{}

	tmp := s.D.Id()
	request.IpscId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetIPSecConnection(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.IpSecConnection
	return nil
}

func (s *CoreIpSecConnectionResourceCrud) Update() error {
	request := oci_core.UpdateIPSecConnectionRequest{}

	if cpeLocalIdentifier, ok := s.D.GetOkExists("cpe_local_identifier"); ok {
		tmp := cpeLocalIdentifier.(string)
		request.CpeLocalIdentifier = &tmp
	}

	if cpeLocalIdentifierType, ok := s.D.GetOkExists("cpe_local_identifier_type"); ok {
		request.CpeLocalIdentifierType = oci_core.UpdateIpSecConnectionDetailsCpeLocalIdentifierTypeEnum(cpeLocalIdentifierType.(string))
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.IpscId = &tmp

	request.StaticRoutes = []string{}
	if staticRoutes, ok := s.D.GetOkExists("static_routes"); ok {
		interfaces := staticRoutes.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		request.StaticRoutes = tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateIPSecConnection(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.IpSecConnection
	return nil
}

func (s *CoreIpSecConnectionResourceCrud) Delete() error {
	request := oci_core.DeleteIPSecConnectionRequest{}

	tmp := s.D.Id()
	request.IpscId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteIPSecConnection(context.Background(), request)
	return err
}

func (s *CoreIpSecConnectionResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CpeId != nil {
		s.D.Set("cpe_id", *s.Res.CpeId)
	}

	if s.Res.CpeLocalIdentifier != nil {
		s.D.Set("cpe_local_identifier", *s.Res.CpeLocalIdentifier)
	}

	s.D.Set("cpe_local_identifier_type", s.Res.CpeLocalIdentifierType)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.DrgId != nil {
		s.D.Set("drg_id", *s.Res.DrgId)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("static_routes", s.Res.StaticRoutes)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func (s *CoreIpSecConnectionResourceCrud) mapToCreateIPSecConnectionTunnelDetails(fieldKeyFormat string) (oci_core.CreateIpSecConnectionTunnelDetails, error) {
	result := oci_core.CreateIpSecConnectionTunnelDetails{}

	if bgpSessionConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "bgp_session_config")); ok {
		if tmpList := bgpSessionConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "bgp_session_config"), 0)
			tmp, err := s.mapToCreateIPSecTunnelBgpSessionDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert bgp_session_config, encountered error: %v", err)
			}
			result.BgpSessionConfig = &tmp
		}
	}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if routing, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "routing")); ok {
		result.Routing = oci_core.CreateIpSecConnectionTunnelDetailsRoutingEnum(routing.(string))
	}

	if sharedSecret, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "shared_secret")); ok {
		tmp := sharedSecret.(string)
		result.SharedSecret = &tmp
	}

	return result, nil
}

func CreateIPSecConnectionTunnelDetailsToMap(obj oci_core.CreateIpSecConnectionTunnelDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BgpSessionConfig != nil {
		result["bgp_session_config"] = []interface{}{CreateIPSecTunnelBgpSessionDetailsToMap(obj.BgpSessionConfig)}
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["routing"] = string(obj.Routing)

	if obj.SharedSecret != nil {
		result["shared_secret"] = string(*obj.SharedSecret)
	}

	return result
}

func (s *CoreIpSecConnectionResourceCrud) mapToCreateIPSecTunnelBgpSessionDetails(fieldKeyFormat string) (oci_core.CreateIpSecTunnelBgpSessionDetails, error) {
	result := oci_core.CreateIpSecTunnelBgpSessionDetails{}

	if customerBgpAsn, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "customer_bgp_asn")); ok {
		tmp := customerBgpAsn.(string)
		result.CustomerBgpAsn = &tmp
	}

	if customerInterfaceIp, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "customer_interface_ip")); ok {
		tmp := customerInterfaceIp.(string)
		result.CustomerInterfaceIp = &tmp
	}

	if oracleInterfaceIp, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "oracle_interface_ip")); ok {
		tmp := oracleInterfaceIp.(string)
		result.OracleInterfaceIp = &tmp
	}

	return result, nil
}

func CreateIPSecTunnelBgpSessionDetailsToMap(obj *oci_core.CreateIpSecTunnelBgpSessionDetails) map[string]interface{} {
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

func tunnelConfigurationHashCodeForSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if bgpSessionConfig, ok := m["bgp_session_config"]; ok {
		if tmpList := bgpSessionConfig.([]interface{}); len(tmpList) > 0 {
			buf.WriteString("bgp_session_config-")
			bgpSessionConfigRaw := tmpList[0].(map[string]interface{})
			if customerBgpAsn, ok := bgpSessionConfigRaw["customer_bgp_asn"]; ok && customerBgpAsn != "" {
				buf.WriteString(fmt.Sprintf("%v-", customerBgpAsn))
			}
			if customerInterfaceIp, ok := bgpSessionConfigRaw["customer_interface_ip"]; ok && customerInterfaceIp != "" {
				buf.WriteString(fmt.Sprintf("%v-", customerInterfaceIp))
			}
			if oracleInterfaceIp, ok := bgpSessionConfigRaw["oracle_interface_ip"]; ok && oracleInterfaceIp != "" {
				buf.WriteString(fmt.Sprintf("%v-", oracleInterfaceIp))
			}
		}
	}
	if displayName, ok := m["display_name"]; ok && displayName != "" {
		buf.WriteString(fmt.Sprintf("%v-", displayName))
	}
	if routing, ok := m["routing"]; ok && routing != "" {
		buf.WriteString(fmt.Sprintf("%v-", routing))
	}
	if sharedSecret, ok := m["shared_secret"]; ok && sharedSecret != "" {
		buf.WriteString(fmt.Sprintf("%v-", sharedSecret))
	}
	return hashcode.String(buf.String())
}
