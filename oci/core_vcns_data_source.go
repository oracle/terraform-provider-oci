// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v27/core"
)

func init() {
	RegisterDatasource("oci_core_vcns", CoreVcnsDataSource())
}

func CoreVcnsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreVcns,
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
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"virtual_networks": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(CoreVcnResource()),
			},
		},
	}
}

func readCoreVcns(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVcnsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient()

	return ReadResource(sync)
}

type CoreVcnsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.ListVcnsResponse
}

func (s *CoreVcnsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreVcnsDataSourceCrud) Get() error {
	request := oci_core.ListVcnsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_core.VcnLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := s.Client.ListVcns(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListVcns(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreVcnsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		vcn := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.CidrBlock != nil {
			vcn["cidr_block"] = *r.CidrBlock
		}

		if r.DefaultDhcpOptionsId != nil {
			vcn["default_dhcp_options_id"] = *r.DefaultDhcpOptionsId
		}

		if r.DefaultRouteTableId != nil {
			vcn["default_route_table_id"] = *r.DefaultRouteTableId
		}

		if r.DefaultSecurityListId != nil {
			vcn["default_security_list_id"] = *r.DefaultSecurityListId
		}

		if r.DefinedTags != nil {
			vcn["defined_tags"] = definedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			vcn["display_name"] = *r.DisplayName
		}

		if r.DnsLabel != nil {
			vcn["dns_label"] = *r.DnsLabel
		}

		vcn["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			vcn["id"] = *r.Id
		}

		if r.Ipv6CidrBlock != nil {
			vcn["ipv6cidr_block"] = *r.Ipv6CidrBlock
			vcn["is_ipv6enabled"] = true
		}

		if r.Ipv6PublicCidrBlock != nil {
			vcn["ipv6public_cidr_block"] = *r.Ipv6PublicCidrBlock
		}

		vcn["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			vcn["time_created"] = r.TimeCreated.String()
		}

		if r.VcnDomainName != nil {
			vcn["vcn_domain_name"] = *r.VcnDomainName
		}

		resources = append(resources, vcn)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, CoreVcnsDataSource().Schema["virtual_networks"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("virtual_networks", resources); err != nil {
		return err
	}

	return nil
}
