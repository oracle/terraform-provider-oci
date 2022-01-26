// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package waas

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_waas "github.com/oracle/oci-go-sdk/v56/waas"
)

func WaasEdgeSubnetsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readWaasEdgeSubnets,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"edge_subnets": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"cidr": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"region": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_modified": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readWaasEdgeSubnets(d *schema.ResourceData, m interface{}) error {
	sync := &WaasEdgeSubnetsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).WaasClient()

	return tfresource.ReadResource(sync)
}

type WaasEdgeSubnetsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_waas.WaasClient
	Res    *oci_waas.ListEdgeSubnetsResponse
}

func (s *WaasEdgeSubnetsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *WaasEdgeSubnetsDataSourceCrud) Get() error {
	request := oci_waas.ListEdgeSubnetsRequest{}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "waas")

	response, err := s.Client.ListEdgeSubnets(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListEdgeSubnets(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *WaasEdgeSubnetsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("WaasEdgeSubnetsDataSource-", WaasEdgeSubnetsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		edgeSubnet := map[string]interface{}{}

		if r.Cidr != nil {
			edgeSubnet["cidr"] = *r.Cidr
		}

		if r.Region != nil {
			edgeSubnet["region"] = *r.Region
		}

		if r.TimeModified != nil {
			edgeSubnet["time_modified"] = r.TimeModified.String()
		}

		resources = append(resources, edgeSubnet)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, WaasEdgeSubnetsDataSource().Schema["edge_subnets"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("edge_subnets", resources); err != nil {
		return err
	}

	return nil
}
