// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/core"

	"github.com/oracle/terraform-provider-oci/crud"
)

// @CODEGEN: Override the resource schema's 'security_list_ids' to make it TypeList instead of TypeSet.
// Avoids a potential breaking change with existing schema.
func SubnetDataSource() *schema.Resource {
	result := SubnetResource()

	result.Schema["security_list_ids"] = &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	}

	return result
}

func SubnetsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSubnets,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": {
				Type:       schema.TypeInt,
				Optional:   true,
				Deprecated: crud.FieldDeprecated("limit"),
			},
			"page": {
				Type:       schema.TypeString,
				Optional:   true,
				Deprecated: crud.FieldDeprecated("page"),
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcn_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subnets": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     SubnetDataSource(),
			},
		},
	}
}

func readSubnets(d *schema.ResourceData, m interface{}) error {
	sync := &SubnetsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return crud.ReadResource(sync)
}

type SubnetsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.ListSubnetsResponse
}

func (s *SubnetsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *SubnetsDataSourceCrud) Get() error {
	request := oci_core.ListSubnetsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if limit, ok := s.D.GetOkExists("limit"); ok {
		tmp := limit.(int)
		request.Limit = &tmp
	}

	if page, ok := s.D.GetOkExists("page"); ok {
		tmp := page.(string)
		request.Page = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_core.SubnetLifecycleStateEnum(state.(string))
	}

	if vcnId, ok := s.D.GetOkExists("vcn_id"); ok {
		tmp := vcnId.(string)
		request.VcnId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

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

func (s *SubnetsDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		subnet := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
			"vcn_id":         *r.VcnId,
		}

		if r.AvailabilityDomain != nil {
			subnet["availability_domain"] = *r.AvailabilityDomain
		}

		if r.CidrBlock != nil {
			subnet["cidr_block"] = *r.CidrBlock
		}

		if r.DefinedTags != nil {
			subnet["defined_tags"] = definedTagsToMap(r.DefinedTags)
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

		if r.VirtualRouterIp != nil {
			subnet["virtual_router_ip"] = *r.VirtualRouterIp
		}

		if r.VirtualRouterMac != nil {
			subnet["virtual_router_mac"] = *r.VirtualRouterMac
		}

		resources = append(resources, subnet)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, SubnetsDataSource().Schema["subnets"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("subnets", resources); err != nil {
		panic(err)
	}

	return
}
