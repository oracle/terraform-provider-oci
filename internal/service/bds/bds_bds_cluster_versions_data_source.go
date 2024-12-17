// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package bds

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_bds "github.com/oracle/oci-go-sdk/v65/bds"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func BdsBdsClusterVersionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readBdsBdsClusterVersions,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"bds_cluster_versions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"bds_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"odh_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readBdsBdsClusterVersions(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsClusterVersionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	return tfresource.ReadResource(sync)
}

type BdsBdsClusterVersionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_bds.BdsClient
	Res    *oci_bds.ListBdsClusterVersionsResponse
}

func (s *BdsBdsClusterVersionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BdsBdsClusterVersionsDataSourceCrud) Get() error {
	request := oci_bds.ListBdsClusterVersionsRequest{}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "bds")

	response, err := s.Client.ListBdsClusterVersions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListBdsClusterVersions(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *BdsBdsClusterVersionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("BdsBdsClusterVersionsDataSource-", BdsBdsClusterVersionsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		bdsClusterVersion := map[string]interface{}{}

		if r.BdsVersion != nil {
			bdsClusterVersion["bds_version"] = *r.BdsVersion
		}

		if r.OdhVersion != nil {
			bdsClusterVersion["odh_version"] = *r.OdhVersion
		}

		resources = append(resources, bdsClusterVersion)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, BdsBdsClusterVersionsDataSource().Schema["bds_cluster_versions"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("bds_cluster_versions", resources); err != nil {
		return err
	}

	return nil
}
