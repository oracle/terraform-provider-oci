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

func DatabaseManagementExternalClusterInstancesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseManagementExternalClusterInstances,
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
			"external_cluster_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_cluster_instance_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(DatabaseManagementExternalClusterInstanceResource()),
						},
					},
				},
			},
		},
	}
}

func readDatabaseManagementExternalClusterInstances(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalClusterInstancesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementExternalClusterInstancesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.ListExternalClusterInstancesResponse
}

func (s *DatabaseManagementExternalClusterInstancesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementExternalClusterInstancesDataSourceCrud) Get() error {
	request := oci_database_management.ListExternalClusterInstancesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if externalClusterId, ok := s.D.GetOkExists("external_cluster_id"); ok {
		tmp := externalClusterId.(string)
		request.ExternalClusterId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.ListExternalClusterInstances(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListExternalClusterInstances(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseManagementExternalClusterInstancesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementExternalClusterInstancesDataSource-", DatabaseManagementExternalClusterInstancesDataSource(), s.D))
	resources := []map[string]interface{}{}
	externalClusterInstance := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ExternalClusterInstanceSummaryToMap(item))
	}
	externalClusterInstance["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseManagementExternalClusterInstancesDataSource().Schema["external_cluster_instance_collection"].Elem.(*schema.Resource).Schema)
		externalClusterInstance["items"] = items
	}

	resources = append(resources, externalClusterInstance)
	if err := s.D.Set("external_cluster_instance_collection", resources); err != nil {
		return err
	}

	return nil
}
