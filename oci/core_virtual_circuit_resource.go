// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"bytes"
	"context"
	"fmt"

	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"

	"strings"

	oci_core "github.com/oracle/oci-go-sdk/core"
)

func VirtualCircuitResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createVirtualCircuit,
		Read:     readVirtualCircuit,
		Update:   updateVirtualCircuit,
		Delete:   deleteVirtualCircuit,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"bandwidth_shape_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cross_connect_mappings": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"bgp_md5auth_key": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cross_connect_or_cross_connect_group_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"customer_bgp_peering_ip": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"oracle_bgp_peering_ip": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vlan": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"customer_bgp_asn": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gateway_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"provider_service_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"public_prefixes": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Set:      publicPrefixesHashCodeForSets,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"cidr_block": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"region": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"bgp_management": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"bgp_session_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"oracle_bgp_asn": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"provider_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"reference_comment": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
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

func createVirtualCircuit(d *schema.ResourceData, m interface{}) error {
	sync := &VirtualCircuitResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return CreateResource(d, sync)
}

func readVirtualCircuit(d *schema.ResourceData, m interface{}) error {
	sync := &VirtualCircuitResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return ReadResource(sync)
}

func updateVirtualCircuit(d *schema.ResourceData, m interface{}) error {
	sync := &VirtualCircuitResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return UpdateResource(d, sync)
}

func deleteVirtualCircuit(d *schema.ResourceData, m interface{}) error {
	sync := &VirtualCircuitResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type VirtualCircuitResourceCrud struct {
	BaseCrud
	Client                 *oci_core.VirtualNetworkClient
	Res                    *oci_core.VirtualCircuit
	DisableNotFoundRetries bool
}

func (s *VirtualCircuitResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *VirtualCircuitResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.VirtualCircuitLifecycleStateVerifying),
		string(oci_core.VirtualCircuitLifecycleStateProvisioning),
	}
}

func (s *VirtualCircuitResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.VirtualCircuitLifecycleStatePendingProvider),
		string(oci_core.VirtualCircuitLifecycleStateProvisioned),
	}
}

func (s *VirtualCircuitResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.VirtualCircuitLifecycleStateTerminating),
	}
}

func (s *VirtualCircuitResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.VirtualCircuitLifecycleStateTerminated),
	}
}

func (s *VirtualCircuitResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_core.VirtualCircuitLifecycleStateProvisioning),
	}
}

func (s *VirtualCircuitResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_core.VirtualCircuitLifecycleStatePendingProvider),
		string(oci_core.VirtualCircuitLifecycleStateProvisioned),
	}
}

func (s *VirtualCircuitResourceCrud) Create() error {
	request := oci_core.CreateVirtualCircuitRequest{}

	if bandwidthShapeName, ok := s.D.GetOkExists("bandwidth_shape_name"); ok {
		tmp := bandwidthShapeName.(string)
		request.BandwidthShapeName = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if crossConnectMappings, ok := s.D.GetOkExists("cross_connect_mappings"); ok {
		interfaces := crossConnectMappings.([]interface{})
		tmp := make([]oci_core.CrossConnectMapping, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "cross_connect_mappings", stateDataIndex)
			converted, err := s.mapToCrossConnectMapping(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		request.CrossConnectMappings = tmp
	}

	if customerBgpAsn, ok := s.D.GetOkExists("customer_bgp_asn"); ok {
		tmp := customerBgpAsn.(int)
		request.CustomerBgpAsn = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if gatewayId, ok := s.D.GetOkExists("gateway_id"); ok {
		tmp := gatewayId.(string)
		request.GatewayId = &tmp
	}

	if providerServiceId, ok := s.D.GetOkExists("provider_service_id"); ok {
		tmp := providerServiceId.(string)
		request.ProviderServiceId = &tmp
	}

	if publicPrefixes, ok := s.D.GetOkExists("public_prefixes"); ok {
		set := publicPrefixes.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_core.CreateVirtualCircuitPublicPrefixDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := publicPrefixesHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "public_prefixes", stateDataIndex)
			converted, err := s.mapToCreateVirtualCircuitPublicPrefixDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		request.PublicPrefixes = tmp
	}

	// Virtual Circuit of type 'PRIVATE' does not support publicPrefixes in payload
	if len(request.PublicPrefixes) == 0 {
		request.PublicPrefixes = nil
	}

	if region, ok := s.D.GetOkExists("region"); ok {
		tmp := region.(string)
		request.Region = &tmp
	}

	if type_, ok := s.D.GetOkExists("type"); ok {
		request.Type = oci_core.CreateVirtualCircuitDetailsTypeEnum(type_.(string))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateVirtualCircuit(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.VirtualCircuit
	return nil
}

func (s *VirtualCircuitResourceCrud) Get() error {
	request := oci_core.GetVirtualCircuitRequest{}

	tmp := s.D.Id()
	request.VirtualCircuitId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetVirtualCircuit(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.VirtualCircuit

	// VirtualCircuitPublicPrefixes are returned from another API, make a List call on them if the VirtualCircuit is of type 'PUBLIC'
	if type_, ok := s.D.GetOkExists("type"); ok && strings.EqualFold(type_.(string), string(oci_core.VirtualCircuitTypePublic)) && s.Res.PublicPrefixes == nil {
		request2 := oci_core.ListVirtualCircuitPublicPrefixesRequest{}
		request2.VirtualCircuitId = request.VirtualCircuitId

		response2, err2 := s.Client.ListVirtualCircuitPublicPrefixes(context.Background(), request2)

		publicPrefixes := []string{}
		for _, item := range response2.Items {
			publicPrefixes = append(publicPrefixes, *item.CidrBlock)
		}

		s.Res.PublicPrefixes = publicPrefixes
		if err2 != nil {
			return err2
		}
	}
	return nil
}

func (s *VirtualCircuitResourceCrud) Update() error {
	request := oci_core.UpdateVirtualCircuitRequest{}

	if bandwidthShapeName, ok := s.D.GetOkExists("bandwidth_shape_name"); ok {
		tmp := bandwidthShapeName.(string)
		request.BandwidthShapeName = &tmp
	}

	request.CrossConnectMappings = []oci_core.CrossConnectMapping{}
	if crossConnectMappings, ok := s.D.GetOkExists("cross_connect_mappings"); ok {
		interfaces := crossConnectMappings.([]interface{})
		tmp := make([]oci_core.CrossConnectMapping, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "cross_connect_mappings", stateDataIndex)
			converted, err := s.mapToCrossConnectMapping(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		request.CrossConnectMappings = tmp
	}

	if customerBgpAsn, ok := s.D.GetOkExists("customer_bgp_asn"); ok {
		tmp := customerBgpAsn.(int)
		request.CustomerBgpAsn = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if gatewayId, ok := s.D.GetOkExists("gateway_id"); ok {
		tmp := gatewayId.(string)
		request.GatewayId = &tmp
	}

	if providerState, ok := s.D.GetOkExists("provider_state"); ok {
		request.ProviderState = oci_core.UpdateVirtualCircuitDetailsProviderStateEnum(providerState.(string))
	}

	if referenceComment, ok := s.D.GetOkExists("reference_comment"); ok {
		tmp := referenceComment.(string)
		request.ReferenceComment = &tmp
	}

	tmp := s.D.Id()
	request.VirtualCircuitId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateVirtualCircuit(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.VirtualCircuit

	return nil
}

func (s *VirtualCircuitResourceCrud) Delete() error {
	request := oci_core.DeleteVirtualCircuitRequest{}

	tmp := s.D.Id()
	request.VirtualCircuitId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteVirtualCircuit(context.Background(), request)
	return err
}

func (s *VirtualCircuitResourceCrud) SetData() error {
	if s.Res.BandwidthShapeName != nil {
		s.D.Set("bandwidth_shape_name", *s.Res.BandwidthShapeName)
	}

	s.D.Set("bgp_management", s.Res.BgpManagement)

	s.D.Set("bgp_session_state", s.Res.BgpSessionState)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	crossConnectMappings := []interface{}{}
	for _, item := range s.Res.CrossConnectMappings {
		crossConnectMappings = append(crossConnectMappings, CrossConnectMappingToMap(item))
	}
	s.D.Set("cross_connect_mappings", crossConnectMappings)

	if s.Res.CustomerBgpAsn != nil {
		s.D.Set("customer_bgp_asn", *s.Res.CustomerBgpAsn)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.GatewayId != nil {
		s.D.Set("gateway_id", *s.Res.GatewayId)
	}

	if s.Res.OracleBgpAsn != nil {
		s.D.Set("oracle_bgp_asn", *s.Res.OracleBgpAsn)
	}

	if s.Res.ProviderServiceId != nil {
		s.D.Set("provider_service_id", *s.Res.ProviderServiceId)
	}

	s.D.Set("provider_state", s.Res.ProviderState)

	if s.Res.PublicPrefixes != nil {
		publicPrefixes := []interface{}{}
		for _, item := range s.Res.PublicPrefixes {
			publicPrefixes = append(publicPrefixes, CreateVirtualCircuitPublicPrefixDetailsToMap(item))
		}
		s.D.Set("public_prefixes", schema.NewSet(publicPrefixesHashCodeForSets, publicPrefixes))
	}

	if s.Res.ReferenceComment != nil {
		s.D.Set("reference_comment", *s.Res.ReferenceComment)
	}

	if s.Res.Region != nil {
		s.D.Set("region", *s.Res.Region)
	}

	s.D.Set("service_type", s.Res.ServiceType)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	s.D.Set("type", s.Res.Type)

	return nil
}

func (s *VirtualCircuitResourceCrud) mapToCreateVirtualCircuitPublicPrefixDetails(fieldKeyFormat string) (oci_core.CreateVirtualCircuitPublicPrefixDetails, error) {
	result := oci_core.CreateVirtualCircuitPublicPrefixDetails{}

	if cidrBlock, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cidr_block")); ok {
		tmp := cidrBlock.(string)
		result.CidrBlock = &tmp
	}

	return result, nil
}

func CreateVirtualCircuitPublicPrefixDetailsToMap(obj string) map[string]interface{} {
	result := map[string]interface{}{}

	result["cidr_block"] = obj

	return result
}

func (s *VirtualCircuitResourceCrud) mapToCrossConnectMapping(fieldKeyFormat string) (oci_core.CrossConnectMapping, error) {
	result := oci_core.CrossConnectMapping{}

	if bgpMd5AuthKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "bgp_md5auth_key")); ok && bgpMd5AuthKey != "" {
		tmp := bgpMd5AuthKey.(string)
		result.BgpMd5AuthKey = &tmp
	}

	if crossConnectOrCrossConnectGroupId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cross_connect_or_cross_connect_group_id")); ok {
		tmp := crossConnectOrCrossConnectGroupId.(string)
		result.CrossConnectOrCrossConnectGroupId = &tmp
	}

	if customerBgpPeeringIp, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "customer_bgp_peering_ip")); ok {
		tmp := customerBgpPeeringIp.(string)
		result.CustomerBgpPeeringIp = &tmp
	}

	if oracleBgpPeeringIp, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "oracle_bgp_peering_ip")); ok {
		tmp := oracleBgpPeeringIp.(string)
		result.OracleBgpPeeringIp = &tmp
	}

	if vlan, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vlan")); ok {
		tmp := vlan.(int)
		//Vlan value must be greater than or equal to 100. It cannot be specified for certain circuit types.
		if tmp >= 100 {
			result.Vlan = &tmp
		}
	}

	return result, nil
}

func CrossConnectMappingToMap(obj oci_core.CrossConnectMapping) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BgpMd5AuthKey != nil {
		result["bgp_md5auth_key"] = string(*obj.BgpMd5AuthKey)
	}

	if obj.CrossConnectOrCrossConnectGroupId != nil {
		result["cross_connect_or_cross_connect_group_id"] = string(*obj.CrossConnectOrCrossConnectGroupId)
	}

	if obj.CustomerBgpPeeringIp != nil {
		result["customer_bgp_peering_ip"] = string(*obj.CustomerBgpPeeringIp)
	}

	if obj.OracleBgpPeeringIp != nil {
		result["oracle_bgp_peering_ip"] = string(*obj.OracleBgpPeeringIp)
	}

	if obj.Vlan != nil {
		result["vlan"] = int(*obj.Vlan)
	}

	return result
}

func publicPrefixesHashCodeForSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if cidrBlock, ok := m["cidr_block"]; ok && cidrBlock != "" {
		buf.WriteString(fmt.Sprintf("%v-", cidrBlock))
	}
	return hashcode.String(buf.String())
}
