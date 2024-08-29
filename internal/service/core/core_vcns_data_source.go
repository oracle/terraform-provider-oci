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

func CoreVcnsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreVcns,
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
			"virtual_networks": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(CoreVcnResource()),
			},
		},
	}
}

func readCoreVcns(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVcnsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

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

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreVcnsDataSource-", CoreVcnsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		vcn := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		vcn["byoipv6cidr_blocks"] = r.Byoipv6CidrBlocks

		if r.CidrBlock != nil {
			vcn["cidr_block"] = *r.CidrBlock
		}

		if r.CidrBlocks != nil {
			vcn["cidr_blocks"] = r.CidrBlocks
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
			vcn["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
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

		vcn["ipv6cidr_blocks"] = r.Ipv6CidrBlocks

		vcn["ipv6private_cidr_blocks"] = r.Ipv6PrivateCidrBlocks

		if r.SecurityAttributes != nil {
			vcn["security_attributes"] = tfresource.SecurityAttributesToMap(r.SecurityAttributes)
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
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreVcnsDataSource().Schema["virtual_networks"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("virtual_networks", resources); err != nil {
		return err
	}

	return nil
}
