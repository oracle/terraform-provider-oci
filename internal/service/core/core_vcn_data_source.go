// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CoreVcnDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["vcn_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(CoreVcnResource(), fieldMap, readSingularCoreVcn)
}

func readSingularCoreVcn(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVcnDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

type CoreVcnDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.GetVcnResponse
}

func (s *CoreVcnDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreVcnDataSourceCrud) Get() error {
	request := oci_core.GetVcnRequest{}

	if vcnId, ok := s.D.GetOkExists("vcn_id"); ok {
		tmp := vcnId.(string)
		request.VcnId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.GetVcn(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CoreVcnDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("byoipv6cidr_blocks", s.Res.Byoipv6CidrBlocks)

	if s.Res.CidrBlock != nil {
		s.D.Set("cidr_block", *s.Res.CidrBlock)
	}

	if s.Res.CidrBlocks != nil {
		s.D.Set("cidr_blocks", s.Res.CidrBlocks)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefaultDhcpOptionsId != nil {
		s.D.Set("default_dhcp_options_id", *s.Res.DefaultDhcpOptionsId)
	}

	if s.Res.DefaultRouteTableId != nil {
		s.D.Set("default_route_table_id", *s.Res.DefaultRouteTableId)
	}

	if s.Res.DefaultSecurityListId != nil {
		s.D.Set("default_security_list_id", *s.Res.DefaultSecurityListId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.DnsLabel != nil {
		s.D.Set("dns_label", *s.Res.DnsLabel)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("ipv6cidr_blocks", s.Res.Ipv6CidrBlocks)

	s.D.Set("ipv6private_cidr_blocks", s.Res.Ipv6PrivateCidrBlocks)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.VcnDomainName != nil {
		s.D.Set("vcn_domain_name", *s.Res.VcnDomainName)
	}

	return nil
}
