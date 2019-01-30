// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/core"
)

func NatGatewayDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularNatGateway,
		Schema: map[string]*schema.Schema{
			"nat_gateway_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"block_traffic": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"defined_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"nat_ip": {
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
			"vcn_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularNatGateway(d *schema.ResourceData, m interface{}) error {
	sync := &NatGatewayDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return ReadResource(sync)
}

type NatGatewayDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.GetNatGatewayResponse
}

func (s *NatGatewayDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *NatGatewayDataSourceCrud) Get() error {
	request := oci_core.GetNatGatewayRequest{}

	if natGatewayId, ok := s.D.GetOkExists("nat_gateway_id"); ok {
		tmp := natGatewayId.(string)
		request.NatGatewayId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := s.Client.GetNatGateway(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *NatGatewayDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.BlockTraffic != nil {
		s.D.Set("block_traffic", *s.Res.BlockTraffic)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.NatIp != nil {
		s.D.Set("nat_ip", *s.Res.NatIp)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.VcnId != nil {
		s.D.Set("vcn_id", *s.Res.VcnId)
	}

	return nil
}
