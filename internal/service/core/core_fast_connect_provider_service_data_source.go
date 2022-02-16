// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v58/core"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func CoreFastConnectProviderServiceDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularCoreFastConnectProviderService,
		Schema: map[string]*schema.Schema{
			"provider_service_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"bandwith_shape_management": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"customer_asn_management": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"private_peering_bgp_management": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"provider_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"provider_service_key_management": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"provider_service_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"public_peering_bgp_management": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"required_total_cross_connects": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"supported_virtual_circuit_types": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularCoreFastConnectProviderService(d *schema.ResourceData, m interface{}) error {
	sync := &CoreFastConnectProviderServiceDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

type CoreFastConnectProviderServiceDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.GetFastConnectProviderServiceResponse
}

func (s *CoreFastConnectProviderServiceDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreFastConnectProviderServiceDataSourceCrud) Get() error {
	request := oci_core.GetFastConnectProviderServiceRequest{}

	if providerServiceId, ok := s.D.GetOkExists("provider_service_id"); ok {
		tmp := providerServiceId.(string)
		request.ProviderServiceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.GetFastConnectProviderService(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CoreFastConnectProviderServiceDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("bandwith_shape_management", s.Res.BandwithShapeManagement)

	s.D.Set("customer_asn_management", s.Res.CustomerAsnManagement)

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("private_peering_bgp_management", s.Res.PrivatePeeringBgpManagement)

	if s.Res.ProviderName != nil {
		s.D.Set("provider_name", *s.Res.ProviderName)
	}

	s.D.Set("provider_service_key_management", s.Res.ProviderServiceKeyManagement)

	if s.Res.ProviderServiceName != nil {
		s.D.Set("provider_service_name", *s.Res.ProviderServiceName)
	}

	s.D.Set("public_peering_bgp_management", s.Res.PublicPeeringBgpManagement)

	if s.Res.RequiredTotalCrossConnects != nil {
		s.D.Set("required_total_cross_connects", *s.Res.RequiredTotalCrossConnects)
	}

	s.D.Set("supported_virtual_circuit_types", s.Res.SupportedVirtualCircuitTypes)

	s.D.Set("type", s.Res.Type)

	return nil
}
