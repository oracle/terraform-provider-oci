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

func CoreSubnetDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["subnet_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(CoreSubnetResource(), fieldMap, readSingularCoreSubnet)
}

func readSingularCoreSubnet(d *schema.ResourceData, m interface{}) error {
	sync := &CoreSubnetDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

type CoreSubnetDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.GetSubnetResponse
}

func (s *CoreSubnetDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreSubnetDataSourceCrud) Get() error {
	request := oci_core.GetSubnetRequest{}

	if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
		tmp := subnetId.(string)
		request.SubnetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.GetSubnet(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CoreSubnetDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.CidrBlock != nil {
		s.D.Set("cidr_block", *s.Res.CidrBlock)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DhcpOptionsId != nil {
		s.D.Set("dhcp_options_id", *s.Res.DhcpOptionsId)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.DnsLabel != nil {
		s.D.Set("dns_label", *s.Res.DnsLabel)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Ipv6CidrBlock != nil {
		s.D.Set("ipv6cidr_block", *s.Res.Ipv6CidrBlock)
	}

	if s.Res.Ipv6VirtualRouterIp != nil {
		s.D.Set("ipv6virtual_router_ip", *s.Res.Ipv6VirtualRouterIp)
	}

	if s.Res.ProhibitInternetIngress != nil {
		s.D.Set("prohibit_internet_ingress", *s.Res.ProhibitInternetIngress)
	}

	if s.Res.ProhibitPublicIpOnVnic != nil {
		s.D.Set("prohibit_public_ip_on_vnic", *s.Res.ProhibitPublicIpOnVnic)
	}

	if s.Res.RouteTableId != nil {
		s.D.Set("route_table_id", *s.Res.RouteTableId)
	}

	s.D.Set("security_list_ids", s.Res.SecurityListIds)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SubnetDomainName != nil {
		s.D.Set("subnet_domain_name", *s.Res.SubnetDomainName)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.VcnId != nil {
		s.D.Set("vcn_id", *s.Res.VcnId)
	}

	if s.Res.VirtualRouterIp != nil {
		s.D.Set("virtual_router_ip", *s.Res.VirtualRouterIp)
	}

	if s.Res.VirtualRouterMac != nil {
		s.D.Set("virtual_router_mac", *s.Res.VirtualRouterMac)
	}

	return nil
}
