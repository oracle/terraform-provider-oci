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

func DatabaseManagementCloudDbSystemDiscoveriesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseManagementCloudDbSystemDiscoveries,
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
			"cloud_db_system_discovery_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(DatabaseManagementCloudDbSystemDiscoveryResource()),
						},
					},
				},
			},
		},
	}
}

func readDatabaseManagementCloudDbSystemDiscoveries(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementCloudDbSystemDiscoveriesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementCloudDbSystemDiscoveriesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.ListCloudDbSystemDiscoveriesResponse
}

func (s *DatabaseManagementCloudDbSystemDiscoveriesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementCloudDbSystemDiscoveriesDataSourceCrud) Get() error {
	request := oci_database_management.ListCloudDbSystemDiscoveriesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.ListCloudDbSystemDiscoveries(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListCloudDbSystemDiscoveries(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseManagementCloudDbSystemDiscoveriesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementCloudDbSystemDiscoveriesDataSource-", DatabaseManagementCloudDbSystemDiscoveriesDataSource(), s.D))
	resources := []map[string]interface{}{}
	cloudDbSystemDiscovery := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, CloudDbSystemDiscoverySummaryToMap(item))
	}
	cloudDbSystemDiscovery["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseManagementCloudDbSystemDiscoveriesDataSource().Schema["cloud_db_system_discovery_collection"].Elem.(*schema.Resource).Schema)
		cloudDbSystemDiscovery["items"] = items
	}

	resources = append(resources, cloudDbSystemDiscovery)
	if err := s.D.Set("cloud_db_system_discovery_collection", resources); err != nil {
		return err
	}

	return nil
}
