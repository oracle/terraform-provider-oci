// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
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

func CoreVirtualCircuitDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["virtual_circuit_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(CoreVirtualCircuitResource(), fieldMap, readSingularCoreVirtualCircuit)
}

func readSingularCoreVirtualCircuit(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVirtualCircuitDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

type CoreVirtualCircuitDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.GetVirtualCircuitResponse
}

func (s *CoreVirtualCircuitDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreVirtualCircuitDataSourceCrud) Get() error {
	request := oci_core.GetVirtualCircuitRequest{}

	if virtualCircuitId, ok := s.D.GetOkExists("virtual_circuit_id"); ok {
		tmp := virtualCircuitId.(string)
		request.VirtualCircuitId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.GetVirtualCircuit(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CoreVirtualCircuitDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.BandwidthShapeName != nil {
		s.D.Set("bandwidth_shape_name", *s.Res.BandwidthShapeName)
	}

	s.D.Set("bgp_admin_state", s.Res.BgpAdminState)

	s.D.Set("bgp_ipv6session_state", s.Res.BgpIpv6SessionState)

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

	if s.Res.CustomerAsn != nil {
		s.D.Set("customer_asn", strconv.FormatInt(*s.Res.CustomerAsn, 10))
	}

	if s.Res.CustomerBgpAsn != nil {
		s.D.Set("customer_bgp_asn", *s.Res.CustomerBgpAsn)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.GatewayId != nil {
		s.D.Set("gateway_id", *s.Res.GatewayId)
	}

	s.D.Set("ip_mtu", s.Res.IpMtu)

	if s.Res.IsBfdEnabled != nil {
		s.D.Set("is_bfd_enabled", *s.Res.IsBfdEnabled)
	}

	if s.Res.IsTransportMode != nil {
		s.D.Set("is_transport_mode", *s.Res.IsTransportMode)
	}

	if s.Res.OracleBgpAsn != nil {
		s.D.Set("oracle_bgp_asn", *s.Res.OracleBgpAsn)
	}

	if s.Res.ProviderServiceId != nil {
		s.D.Set("provider_service_id", *s.Res.ProviderServiceId)
	}

	if s.Res.ProviderServiceKeyName != nil {
		s.D.Set("provider_service_key_name", *s.Res.ProviderServiceKeyName)
	}

	s.D.Set("provider_state", s.Res.ProviderState)

	publicPrefixes := []interface{}{}
	for _, item := range s.Res.PublicPrefixes {
		publicPrefixes = append(publicPrefixes, CreateVirtualCircuitPublicPrefixDetailsToMap(item))
	}
	s.D.Set("public_prefixes", publicPrefixes)

	if s.Res.ReferenceComment != nil {
		s.D.Set("reference_comment", *s.Res.ReferenceComment)
	}

	if s.Res.Region != nil {
		s.D.Set("region", *s.Res.Region)
	}

	s.D.Set("routing_policy", s.Res.RoutingPolicy)

	s.D.Set("service_type", s.Res.ServiceType)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	s.D.Set("type", s.Res.Type)

	return nil
}
