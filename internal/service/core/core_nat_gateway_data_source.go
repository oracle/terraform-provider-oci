// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v56/core"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func CoreNatGatewayDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["nat_gateway_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(CoreNatGatewayResource(), fieldMap, readSingularCoreNatGateway)
}

func readSingularCoreNatGateway(d *schema.ResourceData, m interface{}) error {
	sync := &CoreNatGatewayDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

type CoreNatGatewayDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.GetNatGatewayResponse
}

func (s *CoreNatGatewayDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreNatGatewayDataSourceCrud) Get() error {
	request := oci_core.GetNatGatewayRequest{}

	if natGatewayId, ok := s.D.GetOkExists("nat_gateway_id"); ok {
		tmp := natGatewayId.(string)
		request.NatGatewayId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.GetNatGateway(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CoreNatGatewayDataSourceCrud) SetData() error {
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
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.NatIp != nil {
		s.D.Set("nat_ip", *s.Res.NatIp)
	}

	if s.Res.PublicIpId != nil {
		s.D.Set("public_ip_id", *s.Res.PublicIpId)
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
