// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ocvp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_ocvp "github.com/oracle/oci-go-sdk/v65/ocvp"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OcvpDatastoreClustersDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOcvpDatastoreClusters,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"cluster_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"datastore_cluster_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"datastore_cluster_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(OcvpDatastoreClusterResource()),
						},
					},
				},
			},
		},
	}
}

func readOcvpDatastoreClusters(d *schema.ResourceData, m interface{}) error {
	sync := &OcvpDatastoreClustersDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatastoreClusterClient()

	return tfresource.ReadResource(sync)
}

type OcvpDatastoreClustersDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_ocvp.DatastoreClusterClient
	Res    *oci_ocvp.ListDatastoreClustersResponse
}

func (s *OcvpDatastoreClustersDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OcvpDatastoreClustersDataSourceCrud) Get() error {
	request := oci_ocvp.ListDatastoreClustersRequest{}

	if clusterId, ok := s.D.GetOkExists("cluster_id"); ok {
		tmp := clusterId.(string)
		request.ClusterId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if datastoreClusterId, ok := s.D.GetOkExists("id"); ok {
		tmp := datastoreClusterId.(string)
		request.DatastoreClusterId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_ocvp.ListDatastoreClustersLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "ocvp")

	response, err := s.Client.ListDatastoreClusters(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDatastoreClusters(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OcvpDatastoreClustersDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OcvpDatastoreClustersDataSource-", OcvpDatastoreClustersDataSource(), s.D))
	resources := []map[string]interface{}{}
	datastoreCluster := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DatastoreClusterSummaryToMap(item))
	}
	datastoreCluster["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OcvpDatastoreClustersDataSource().Schema["datastore_cluster_collection"].Elem.(*schema.Resource).Schema)
		datastoreCluster["items"] = items
	}

	resources = append(resources, datastoreCluster)
	if err := s.D.Set("datastore_cluster_collection", resources); err != nil {
		return err
	}

	return nil
}
