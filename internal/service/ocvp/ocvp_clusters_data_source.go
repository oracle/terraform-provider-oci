// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ocvp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_ocvp "github.com/oracle/oci-go-sdk/v65/ocvp"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OcvpClustersDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOcvpClusters,
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
			"sddc_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cluster_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(OcvpClusterResource()),
						},
					},
				},
			},
		},
	}
}

func readOcvpClusters(d *schema.ResourceData, m interface{}) error {
	sync := &OcvpClustersDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ClusterClient()

	return tfresource.ReadResource(sync)
}

type OcvpClustersDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_ocvp.ClusterClient
	Res    *oci_ocvp.ListClustersResponse
}

func (s *OcvpClustersDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OcvpClustersDataSourceCrud) Get() error {
	request := oci_ocvp.ListClustersRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if sddcId, ok := s.D.GetOkExists("sddc_id"); ok {
		tmp := sddcId.(string)
		request.SddcId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_ocvp.ListClustersLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "ocvp")

	response, err := s.Client.ListClusters(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListClusters(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OcvpClustersDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OcvpClustersDataSource-", OcvpClustersDataSource(), s.D))
	resources := []map[string]interface{}{}
	cluster := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ClusterSummaryToMap(item))
	}
	cluster["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OcvpClustersDataSource().Schema["cluster_collection"].Elem.(*schema.Resource).Schema)
		cluster["items"] = items
	}

	resources = append(resources, cluster)
	if err := s.D.Set("cluster_collection", resources); err != nil {
		return err
	}

	return nil
}
