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

func CoreIpv6sDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreIpv6s,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"ip_address": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vnic_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipv6s": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(CoreIpv6Resource()),
			},
		},
	}
}

func readCoreIpv6s(d *schema.ResourceData, m interface{}) error {
	sync := &CoreIpv6sDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

type CoreIpv6sDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.ListIpv6sResponse
}

func (s *CoreIpv6sDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreIpv6sDataSourceCrud) Get() error {
	request := oci_core.ListIpv6sRequest{}

	if ipAddress, ok := s.D.GetOkExists("ip_address"); ok {
		tmp := ipAddress.(string)
		request.IpAddress = &tmp
	}

	if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
		tmp := subnetId.(string)
		request.SubnetId = &tmp
	}

	if vnicId, ok := s.D.GetOkExists("vnic_id"); ok {
		tmp := vnicId.(string)
		request.VnicId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListIpv6s(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListIpv6s(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreIpv6sDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreIpv6sDataSource-", CoreIpv6sDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		ipv6 := map[string]interface{}{}

		if r.CompartmentId != nil {
			ipv6["compartment_id"] = *r.CompartmentId
		}

		if r.DefinedTags != nil {
			ipv6["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			ipv6["display_name"] = *r.DisplayName
		}

		ipv6["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			ipv6["id"] = *r.Id
		}

		if r.IpAddress != nil {
			ipv6["ip_address"] = *r.IpAddress
		}

		if r.SubnetId != nil {
			ipv6["subnet_id"] = *r.SubnetId
		}

		if r.TimeCreated != nil {
			ipv6["time_created"] = r.TimeCreated.String()
		}

		if r.VnicId != nil {
			ipv6["vnic_id"] = *r.VnicId
		}

		if r.RouteTableId != nil {
			ipv6["route_table_id"] = *r.RouteTableId
		}

		ipv6["ip_state"] = r.IpState
		ipv6["lifetime"] = r.Lifetime

		resources = append(resources, ipv6)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreIpv6sDataSource().Schema["ipv6s"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("ipv6s", resources); err != nil {
		return err
	}

	return nil
}
