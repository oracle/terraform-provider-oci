// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package cluster_placement_groups

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_cluster_placement_groups "github.com/oracle/oci-go-sdk/v65/clusterplacementgroups"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ClusterPlacementGroupsClusterPlacementGroupsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readClusterPlacementGroupsClusterPlacementGroups,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"ad": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cluster_placement_group_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(ClusterPlacementGroupsClusterPlacementGroupResource()),
						},
					},
				},
			},
		},
	}
}

func readClusterPlacementGroupsClusterPlacementGroups(d *schema.ResourceData, m interface{}) error {
	sync := &ClusterPlacementGroupsClusterPlacementGroupsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ClusterPlacementGroupsCPClient()

	return tfresource.ReadResource(sync)
}

type ClusterPlacementGroupsClusterPlacementGroupsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_cluster_placement_groups.ClusterPlacementGroupsCPClient
	Res    *oci_cluster_placement_groups.ListClusterPlacementGroupsResponse
}

func (s *ClusterPlacementGroupsClusterPlacementGroupsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ClusterPlacementGroupsClusterPlacementGroupsDataSourceCrud) Get() error {
	request := oci_cluster_placement_groups.ListClusterPlacementGroupsRequest{}

	if ad, ok := s.D.GetOkExists("ad"); ok {
		tmp := ad.(string)
		request.Ad = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_cluster_placement_groups.ClusterPlacementGroupLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "cluster_placement_groups")

	response, err := s.Client.ListClusterPlacementGroups(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListClusterPlacementGroups(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ClusterPlacementGroupsClusterPlacementGroupsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ClusterPlacementGroupsClusterPlacementGroupsDataSource-", ClusterPlacementGroupsClusterPlacementGroupsDataSource(), s.D))
	resources := []map[string]interface{}{}
	clusterPlacementGroup := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ClusterPlacementGroupSummaryToMap(item))
	}
	clusterPlacementGroup["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, ClusterPlacementGroupsClusterPlacementGroupsDataSource().Schema["cluster_placement_group_collection"].Elem.(*schema.Resource).Schema)
		clusterPlacementGroup["items"] = items
	}

	resources = append(resources, clusterPlacementGroup)
	if err := s.D.Set("cluster_placement_group_collection", resources); err != nil {
		return err
	}

	return nil
}
