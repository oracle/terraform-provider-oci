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

func CoreNetworkSecurityGroupVnicsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreNetworkSecurityGroupVnics,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"network_security_group_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"network_security_group_vnics": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"resource_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_associated": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"vnic_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readCoreNetworkSecurityGroupVnics(d *schema.ResourceData, m interface{}) error {
	sync := &CoreNetworkSecurityGroupVnicsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

type CoreNetworkSecurityGroupVnicsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.ListNetworkSecurityGroupVnicsResponse
}

func (s *CoreNetworkSecurityGroupVnicsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreNetworkSecurityGroupVnicsDataSourceCrud) Get() error {
	request := oci_core.ListNetworkSecurityGroupVnicsRequest{}

	if networkSecurityGroupId, ok := s.D.GetOkExists("network_security_group_id"); ok {
		tmp := networkSecurityGroupId.(string)
		request.NetworkSecurityGroupId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListNetworkSecurityGroupVnics(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListNetworkSecurityGroupVnics(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreNetworkSecurityGroupVnicsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreNetworkSecurityGroupVnicsDataSource-", CoreNetworkSecurityGroupVnicsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		networkSecurityGroupVnic := map[string]interface{}{}

		if r.ResourceId != nil {
			networkSecurityGroupVnic["resource_id"] = *r.ResourceId
		}

		if r.TimeAssociated != nil {
			networkSecurityGroupVnic["time_associated"] = r.TimeAssociated.String()
		}

		if r.VnicId != nil {
			networkSecurityGroupVnic["vnic_id"] = *r.VnicId
		}

		resources = append(resources, networkSecurityGroupVnic)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreNetworkSecurityGroupVnicsDataSource().Schema["network_security_group_vnics"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("network_security_group_vnics", resources); err != nil {
		return err
	}

	return nil
}
