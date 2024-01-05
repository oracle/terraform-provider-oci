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

func OpensearchOpensearchVersionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOpensearchOpensearchVersions,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"opensearch_versions_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"version": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func readOpensearchOpensearchVersions(d *schema.ResourceData, m interface{}) error {
	sync := &OpensearchOpensearchVersionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OpensearchClusterClient()

	return tfresource.ReadResource(sync)
}

type OpensearchOpensearchVersionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_opensearch.OpensearchClusterClient
	Res    *oci_opensearch.ListOpensearchVersionsResponse
}

func (s *OpensearchOpensearchVersionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OpensearchOpensearchVersionsDataSourceCrud) Get() error {
	request := oci_opensearch.ListOpensearchVersionsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "opensearch")

	response, err := s.Client.ListOpensearchVersions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListOpensearchVersions(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OpensearchOpensearchVersionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OpensearchOpensearchVersionsDataSource-", OpensearchOpensearchVersionsDataSource(), s.D))
	resources := []map[string]interface{}{}
	opensearchVersion := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, OpensearchVersionsSummaryToMap(item))
	}
	opensearchVersion["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OpensearchOpensearchVersionsDataSource().Schema["opensearch_versions_collection"].Elem.(*schema.Resource).Schema)
		opensearchVersion["items"] = items
	}

	resources = append(resources, opensearchVersion)
	if err := s.D.Set("opensearch_versions_collection", resources); err != nil {
		return err
	}

	return nil
}

func OpensearchVersionsSummaryToMap(obj oci_opensearch.OpensearchVersionsSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Version != nil {
		result["version"] = string(*obj.Version)
	}

	return result
}
