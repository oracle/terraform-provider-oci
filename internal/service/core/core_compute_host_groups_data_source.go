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

func CoreComputeHostGroupsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreComputeHostGroups,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compute_host_group_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(CoreComputeHostGroupResource()),
						},
					},
				},
			},
		},
	}
}

func readCoreComputeHostGroups(d *schema.ResourceData, m interface{}) error {
	sync := &CoreComputeHostGroupsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.ReadResource(sync)
}

type CoreComputeHostGroupsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeClient
	Res    *oci_core.ListComputeHostGroupsResponse
}

func (s *CoreComputeHostGroupsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreComputeHostGroupsDataSourceCrud) Get() error {
	request := oci_core.ListComputeHostGroupsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListComputeHostGroups(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListComputeHostGroups(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreComputeHostGroupsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreComputeHostGroupsDataSource-", CoreComputeHostGroupsDataSource(), s.D))
	resources := []map[string]interface{}{}
	computeHostGroup := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ComputeHostGroupSummaryToMap(item))
	}
	computeHostGroup["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, CoreComputeHostGroupsDataSource().Schema["compute_host_group_collection"].Elem.(*schema.Resource).Schema)
		computeHostGroup["items"] = items
	}

	resources = append(resources, computeHostGroup)
	if err := s.D.Set("compute_host_group_collection", resources); err != nil {
		return err
	}

	return nil
}
