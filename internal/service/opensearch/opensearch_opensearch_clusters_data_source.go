// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package opensearch

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_opensearch "github.com/oracle/oci-go-sdk/v65/opensearch"

	"github.com/oracle/terraform-provider-oci/internal/client"

	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OpensearchOpensearchClustersDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOpensearchOpensearchClusters,
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
			"opensearch_cluster_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(OpensearchOpensearchClusterResource()),
						},
					},
				},
			},
		},
	}
}

func readOpensearchOpensearchClusters(d *schema.ResourceData, m interface{}) error {
	sync := &OpensearchOpensearchClustersDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OpensearchClusterClient()

	return tfresource.ReadResource(sync)
}

type OpensearchOpensearchClustersDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_opensearch.OpensearchClusterClient
	Res    *oci_opensearch.ListOpensearchClustersResponse
}

func (s *OpensearchOpensearchClustersDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OpensearchOpensearchClustersDataSourceCrud) Get() error {
	request := oci_opensearch.ListOpensearchClustersRequest{}

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
		request.LifecycleState = oci_opensearch.OpensearchClusterLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "opensearch")

	response, err := s.Client.ListOpensearchClusters(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListOpensearchClusters(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OpensearchOpensearchClustersDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OpensearchOpensearchClustersDataSource-", OpensearchOpensearchClustersDataSource(), s.D))
	resources := []map[string]interface{}{}
	opensearchCluster := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, OpensearchClusterSummaryToMap(item))
	}
	opensearchCluster["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OpensearchOpensearchClustersDataSource().Schema["opensearch_cluster_collection"].Elem.(*schema.Resource).Schema)
		opensearchCluster["items"] = items
	}

	resources = append(resources, opensearchCluster)
	if err := s.D.Set("opensearch_cluster_collection", resources); err != nil {
		return err
	}

	return nil
}
