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

func CoreSubnetsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreSubnets,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcn_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"subnets": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(CoreSubnetResource()),
			},
		},
	}
}

func readCoreSubnets(d *schema.ResourceData, m interface{}) error {
	sync := &CoreSubnetsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

type CoreSubnetsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.ListSubnetsResponse
}

func (s *CoreSubnetsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreSubnetsDataSourceCrud) Get() error {
	request := oci_core.ListSubnetsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_core.SubnetLifecycleStateEnum(state.(string))
	}

	if vcnId, ok := s.D.GetOkExists("vcn_id"); ok {
		tmp := vcnId.(string)
		request.VcnId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListSubnets(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSubnets(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreSubnetsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreSubnetsDataSource-", CoreSubnetsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		subnet := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.AvailabilityDomain != nil {
			subnet["availability_domain"] = *r.AvailabilityDomain
		}

		if r.CidrBlock != nil {
			subnet["cidr_block"] = *r.CidrBlock
		}

		if r.DefinedTags != nil {
			subnet["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DhcpOptionsId != nil {
			subnet["dhcp_options_id"] = *r.DhcpOptionsId
		}

		if r.DisplayName != nil {
			subnet["display_name"] = *r.DisplayName
		}

		if r.DnsLabel != nil {
			subnet["dns_label"] = *r.DnsLabel
		}

		subnet["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			subnet["id"] = *r.Id
		}

		if r.Ipv6CidrBlock != nil {
			subnet["ipv6cidr_block"] = *r.Ipv6CidrBlock
		}

		subnet["ipv6cidr_blocks"] = r.Ipv6CidrBlocks

		if r.Ipv6VirtualRouterIp != nil {
			subnet["ipv6virtual_router_ip"] = *r.Ipv6VirtualRouterIp
		}

		if r.ProhibitInternetIngress != nil {
			subnet["prohibit_internet_ingress"] = *r.ProhibitInternetIngress
		}

		if r.ProhibitPublicIpOnVnic != nil {
			subnet["prohibit_public_ip_on_vnic"] = *r.ProhibitPublicIpOnVnic
		}

		if r.RouteTableId != nil {
			subnet["route_table_id"] = *r.RouteTableId
		}

		subnet["security_list_ids"] = r.SecurityListIds

		subnet["state"] = r.LifecycleState

		if r.SubnetDomainName != nil {
			subnet["subnet_domain_name"] = *r.SubnetDomainName
		}

		if r.TimeCreated != nil {
			subnet["time_created"] = r.TimeCreated.String()
		}

		if r.VcnId != nil {
			subnet["vcn_id"] = *r.VcnId
		}

		if r.VirtualRouterIp != nil {
			subnet["virtual_router_ip"] = *r.VirtualRouterIp
		}

		if r.VirtualRouterMac != nil {
			subnet["virtual_router_mac"] = *r.VirtualRouterMac
		}

		resources = append(resources, subnet)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreSubnetsDataSource().Schema["subnets"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("subnets", resources); err != nil {
		return err
	}

	return nil
}
