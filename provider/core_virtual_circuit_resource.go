// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"

	oci_core "github.com/oracle/oci-go-sdk/core"
)

func VirtualCircuitResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
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
				Type:     schema.TypeList,
				Optional: true,
				// @CODEGEN 07/2018: The service does not return publicPrefixes as part of GET or LIST operation on this resource
				// To get or update the publicPrefixes, once has to use the VirtualCircuitPublicPrefix APIs: https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/VirtualCircuitPublicPrefix/
				//Computed: true,
				ForceNew: true,
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

	return crud.CreateResource(d, sync)
}

func readVirtualCircuit(d *schema.ResourceData, m interface{}) error {
	sync := &VirtualCircuitResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return crud.ReadResource(sync)
}

func updateVirtualCircuit(d *schema.ResourceData, m interface{}) error {
	sync := &VirtualCircuitResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return crud.UpdateResource(d, sync)
}

func deleteVirtualCircuit(d *schema.ResourceData, m interface{}) error {
	sync := &VirtualCircuitResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient
	sync.DisableNotFoundRetries = true

	return crud.DeleteResource(d, sync)
}

type VirtualCircuitResourceCrud struct {
	crud.BaseCrud
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
		for i, toBeConverted := range interfaces {
			tmp[i] = mapToCrossConnectMapping(toBeConverted.(map[string]interface{}))
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
		interfaces := publicPrefixes.([]interface{})
		tmp := make([]oci_core.CreateVirtualCircuitPublicPrefixDetails, len(interfaces))
		for i, toBeConverted := range interfaces {
			tmp[i] = mapToCreateVirtualCircuitPublicPrefixDetails(toBeConverted.(map[string]interface{}))
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
		for i, toBeConverted := range interfaces {
			tmp[i] = mapToCrossConnectMapping(toBeConverted.(map[string]interface{}))
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

func (s *VirtualCircuitResourceCrud) SetData() {
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
		s.D.Set("public_prefixes", publicPrefixes)
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

}

func mapToCreateVirtualCircuitPublicPrefixDetails(raw map[string]interface{}) oci_core.CreateVirtualCircuitPublicPrefixDetails {
	result := oci_core.CreateVirtualCircuitPublicPrefixDetails{}

	if cidrBlock, ok := raw["cidr_block"]; ok && cidrBlock != "" {
		tmp := cidrBlock.(string)
		result.CidrBlock = &tmp
	}

	return result
}

func CreateVirtualCircuitPublicPrefixDetailsToMap(obj string) map[string]interface{} {
	result := map[string]interface{}{}

	result["cidr_block"] = obj

	return result
}

func mapToCrossConnectMapping(raw map[string]interface{}) oci_core.CrossConnectMapping {
	result := oci_core.CrossConnectMapping{}

	if bgpMd5AuthKey, ok := raw["bgp_md5auth_key"]; ok && bgpMd5AuthKey != "" {
		tmp := bgpMd5AuthKey.(string)
		result.BgpMd5AuthKey = &tmp
	}

	if crossConnectOrCrossConnectGroupId, ok := raw["cross_connect_or_cross_connect_group_id"]; ok && crossConnectOrCrossConnectGroupId != "" {
		tmp := crossConnectOrCrossConnectGroupId.(string)
		result.CrossConnectOrCrossConnectGroupId = &tmp
	}

	if customerBgpPeeringIp, ok := raw["customer_bgp_peering_ip"]; ok && customerBgpPeeringIp != "" {
		tmp := customerBgpPeeringIp.(string)
		result.CustomerBgpPeeringIp = &tmp
	}

	if oracleBgpPeeringIp, ok := raw["oracle_bgp_peering_ip"]; ok && oracleBgpPeeringIp != "" {
		tmp := oracleBgpPeeringIp.(string)
		result.OracleBgpPeeringIp = &tmp
	}

	if vlan, ok := raw["vlan"]; ok {
		tmp := vlan.(int)
		//Vlan value must be greater than or equal to 100. It cannot be specified for certain circuit types.
		if tmp >= 100 {
			result.Vlan = &tmp
		}
	}

	return result
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
