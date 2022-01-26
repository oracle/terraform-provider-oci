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

func CoreVlansDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreVlans,
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
			"vlans": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(CoreVlanResource()),
			},
		},
	}
}

func readCoreVlans(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVlansDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

type CoreVlansDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.ListVlansResponse
}

func (s *CoreVlansDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreVlansDataSourceCrud) Get() error {
	request := oci_core.ListVlansRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_core.VlanLifecycleStateEnum(state.(string))
	}

	if vcnId, ok := s.D.GetOkExists("vcn_id"); ok {
		tmp := vcnId.(string)
		request.VcnId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListVlans(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListVlans(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreVlansDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreVlansDataSource-", CoreVlansDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		vlan := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.AvailabilityDomain != nil {
			vlan["availability_domain"] = *r.AvailabilityDomain
		}

		if r.CidrBlock != nil {
			vlan["cidr_block"] = *r.CidrBlock
		}

		if r.DefinedTags != nil {
			vlan["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			vlan["display_name"] = *r.DisplayName
		}

		vlan["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			vlan["id"] = *r.Id
		}

		vlan["nsg_ids"] = r.NsgIds

		if r.RouteTableId != nil {
			vlan["route_table_id"] = *r.RouteTableId
		}

		vlan["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			vlan["time_created"] = r.TimeCreated.String()
		}

		if r.VcnId != nil {
			vlan["vcn_id"] = *r.VcnId
		}

		if r.VlanTag != nil {
			vlan["vlan_tag"] = *r.VlanTag
		}

		resources = append(resources, vlan)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreVlansDataSource().Schema["vlans"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("vlans", resources); err != nil {
		return err
	}

	return nil
}
