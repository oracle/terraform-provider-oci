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

func CoreNetworkSecurityGroupsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreNetworkSecurityGroups,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
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
			"vlan_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_security_groups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(CoreNetworkSecurityGroupResource()),
			},
		},
	}
}

func readCoreNetworkSecurityGroups(d *schema.ResourceData, m interface{}) error {
	sync := &CoreNetworkSecurityGroupsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

type CoreNetworkSecurityGroupsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.ListNetworkSecurityGroupsResponse
}

func (s *CoreNetworkSecurityGroupsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreNetworkSecurityGroupsDataSourceCrud) Get() error {
	request := oci_core.ListNetworkSecurityGroupsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_core.NetworkSecurityGroupLifecycleStateEnum(state.(string))
	}

	if vcnId, ok := s.D.GetOkExists("vcn_id"); ok {
		tmp := vcnId.(string)
		request.VcnId = &tmp
	}

	if vlanId, ok := s.D.GetOkExists("vlan_id"); ok {
		tmp := vlanId.(string)
		request.VlanId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListNetworkSecurityGroups(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListNetworkSecurityGroups(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreNetworkSecurityGroupsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreNetworkSecurityGroupsDataSource-", CoreNetworkSecurityGroupsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		networkSecurityGroup := map[string]interface{}{}

		if r.CompartmentId != nil {
			networkSecurityGroup["compartment_id"] = *r.CompartmentId
		}

		if r.DefinedTags != nil {
			networkSecurityGroup["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			networkSecurityGroup["display_name"] = *r.DisplayName
		}

		networkSecurityGroup["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			networkSecurityGroup["id"] = *r.Id
		}

		networkSecurityGroup["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			networkSecurityGroup["time_created"] = r.TimeCreated.String()
		}

		if r.VcnId != nil {
			networkSecurityGroup["vcn_id"] = *r.VcnId
		}

		resources = append(resources, networkSecurityGroup)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreNetworkSecurityGroupsDataSource().Schema["network_security_groups"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("network_security_groups", resources); err != nil {
		return err
	}

	return nil
}
