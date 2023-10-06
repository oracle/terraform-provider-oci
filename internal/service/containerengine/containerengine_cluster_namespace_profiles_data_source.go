// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package containerengine

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_containerengine "github.com/oracle/oci-go-sdk/v65/containerengine"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ContainerengineClusterNamespaceProfilesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readContainerengineClusterNamespaceProfiles,
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
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cluster_namespace_profile_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(ContainerengineClusterNamespaceProfileResource()),
						},
					},
				},
			},
		},
	}
}

func readContainerengineClusterNamespaceProfiles(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineClusterNamespaceProfilesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerEngineClient()

	return tfresource.ReadResource(sync)
}

type ContainerengineClusterNamespaceProfilesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_containerengine.ContainerEngineClient
	Res    *oci_containerengine.ListClusterNamespaceProfilesResponse
}

func (s *ContainerengineClusterNamespaceProfilesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ContainerengineClusterNamespaceProfilesDataSourceCrud) Get() error {
	request := oci_containerengine.ListClusterNamespaceProfilesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_containerengine.ClusterNamespaceProfileLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "containerengine")

	response, err := s.Client.ListClusterNamespaceProfiles(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListClusterNamespaceProfiles(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ContainerengineClusterNamespaceProfilesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ContainerengineClusterNamespaceProfilesDataSource-", ContainerengineClusterNamespaceProfilesDataSource(), s.D))
	resources := []map[string]interface{}{}
	clusterNamespaceProfile := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ClusterNamespaceProfileSummaryToMap(item))
	}
	clusterNamespaceProfile["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, ContainerengineClusterNamespaceProfilesDataSource().Schema["cluster_namespace_profile_collection"].Elem.(*schema.Resource).Schema)
		clusterNamespaceProfile["items"] = items
	}

	resources = append(resources, clusterNamespaceProfile)
	if err := s.D.Set("cluster_namespace_profile_collection", resources); err != nil {
		return err
	}

	return nil
}
