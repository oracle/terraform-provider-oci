// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/core"

	"strings"

	"github.com/oracle/terraform-provider-oci/crud"
)

func VirtualCircuitDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularVirtualCircuit,
		Schema: map[string]*schema.Schema{
			"virtual_circuit_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"bandwidth_shape_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"bgp_management": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"bgp_session_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cross_connect_mappings": {
				Type:     schema.TypeList,
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
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"gateway_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"oracle_bgp_asn": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"provider_service_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"provider_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"public_prefixes": {
				Type:     schema.TypeSet,
				Computed: true,
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
			"reference_comment": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"region": {
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
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularVirtualCircuit(d *schema.ResourceData, m interface{}) error {
	sync := &VirtualCircuitDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return crud.ReadResource(sync)
}

type VirtualCircuitDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.GetVirtualCircuitResponse
}

func (s *VirtualCircuitDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *VirtualCircuitDataSourceCrud) Get() error {
	request := oci_core.GetVirtualCircuitRequest{}

	if virtualCircuitId, ok := s.D.GetOkExists("virtual_circuit_id"); ok {
		tmp := virtualCircuitId.(string)
		request.VirtualCircuitId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := s.Client.GetVirtualCircuit(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response

	// VirtualCircuitPublicPrefixes are returned from another API, make a List call on them if the VirtualCircuit is of type 'PUBLIC'
	if strings.EqualFold(string(response.Type), string(oci_core.VirtualCircuitTypePublic)) && s.Res.PublicPrefixes == nil {
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

func (s *VirtualCircuitDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(*s.Res.Id)

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

	publicPrefixes := []interface{}{}
	for _, item := range s.Res.PublicPrefixes {
		publicPrefixes = append(publicPrefixes, CreateVirtualCircuitPublicPrefixDetailsToMap(item))
	}
	s.D.Set("public_prefixes", schema.NewSet(publicPrefixHashCodeForSets, publicPrefixes))

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

	return
}
