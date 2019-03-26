// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/core"
)

func CoreVcnDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularCoreVcn,
		Schema: map[string]*schema.Schema{
			"vcn_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"cidr_block": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_dhcp_options_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_route_table_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_security_list_id": {
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
			"dns_label": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vcn_domain_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularCoreVcn(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVcnDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return ReadResource(sync)
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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

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

	if s.Res.CidrBlock != nil {
		s.D.Set("cidr_block", *s.Res.CidrBlock)
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
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.DnsLabel != nil {
		s.D.Set("dns_label", *s.Res.DnsLabel)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.VcnDomainName != nil {
		s.D.Set("vcn_domain_name", *s.Res.VcnDomainName)
	}

	return nil
}
