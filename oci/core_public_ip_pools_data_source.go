// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v40/core"
)

func init() {
	RegisterDatasource("oci_core_public_ip_pools", CorePublicIpPoolsDataSource())
}

func CorePublicIpPoolsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCorePublicIpPools,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"byoip_range_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"public_ip_pool_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     GetDataSourceItemSchema(CorePublicIpPoolResource()),
						},
					},
				},
			},
		},
	}
}

func readCorePublicIpPools(d *schema.ResourceData, m interface{}) error {
	sync := &CorePublicIpPoolsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient()

	return ReadResource(sync)
}

type CorePublicIpPoolsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.ListPublicIpPoolsResponse
}

func (s *CorePublicIpPoolsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CorePublicIpPoolsDataSourceCrud) Get() error {
	request := oci_core.ListPublicIpPoolsRequest{}

	if byoipRangeId, ok := s.D.GetOkExists("byoip_range_id"); ok {
		tmp := byoipRangeId.(string)
		request.ByoipRangeId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := s.Client.ListPublicIpPools(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListPublicIpPools(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CorePublicIpPoolsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("CorePublicIpPoolsDataSource-", CorePublicIpPoolsDataSource(), s.D))
	resources := []map[string]interface{}{}
	publicIpPool := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, PublicIpPoolSummaryToMap(item))
	}
	publicIpPool["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = ApplyFiltersInCollection(f.(*schema.Set), items, CorePublicIpPoolsDataSource().Schema["public_ip_pool_collection"].Elem.(*schema.Resource).Schema)
		publicIpPool["items"] = items
	}

	resources = append(resources, publicIpPool)
	if err := s.D.Set("public_ip_pool_collection", resources); err != nil {
		return err
	}

	return nil
}
