// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseManagementExternalClustersDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseManagementExternalClusters,
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
			"external_db_system_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_cluster_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(DatabaseManagementExternalClusterResource()),
						},
					},
				},
			},
		},
	}
}

func readDatabaseManagementExternalClusters(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalClustersDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementExternalClustersDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.ListExternalClustersResponse
}

func (s *DatabaseManagementExternalClustersDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementExternalClustersDataSourceCrud) Get() error {
	request := oci_database_management.ListExternalClustersRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if externalDbSystemId, ok := s.D.GetOkExists("external_db_system_id"); ok {
		tmp := externalDbSystemId.(string)
		request.ExternalDbSystemId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.ListExternalClusters(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListExternalClusters(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseManagementExternalClustersDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementExternalClustersDataSource-", DatabaseManagementExternalClustersDataSource(), s.D))
	resources := []map[string]interface{}{}
	externalCluster := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ExternalClusterSummaryToMap(item))
	}
	externalCluster["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseManagementExternalClustersDataSource().Schema["external_cluster_collection"].Elem.(*schema.Resource).Schema)
		externalCluster["items"] = items
	}

	resources = append(resources, externalCluster)
	if err := s.D.Set("external_cluster_collection", resources); err != nil {
		return err
	}

	return nil
}
