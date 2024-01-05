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

func DatabaseManagementExternalDbSystemDiscoveriesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseManagementExternalDbSystemDiscoveries,
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
			"external_db_system_discovery_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(DatabaseManagementExternalDbSystemDiscoveryResource()),
						},
					},
				},
			},
		},
	}
}

func readDatabaseManagementExternalDbSystemDiscoveries(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalDbSystemDiscoveriesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementExternalDbSystemDiscoveriesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.ListExternalDbSystemDiscoveriesResponse
}

func (s *DatabaseManagementExternalDbSystemDiscoveriesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementExternalDbSystemDiscoveriesDataSourceCrud) Get() error {
	request := oci_database_management.ListExternalDbSystemDiscoveriesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.ListExternalDbSystemDiscoveries(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListExternalDbSystemDiscoveries(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseManagementExternalDbSystemDiscoveriesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementExternalDbSystemDiscoveriesDataSource-", DatabaseManagementExternalDbSystemDiscoveriesDataSource(), s.D))
	resources := []map[string]interface{}{}
	externalDbSystemDiscovery := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ExternalDbSystemDiscoverySummaryToMap(item))
	}
	externalDbSystemDiscovery["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseManagementExternalDbSystemDiscoveriesDataSource().Schema["external_db_system_discovery_collection"].Elem.(*schema.Resource).Schema)
		externalDbSystemDiscovery["items"] = items
	}

	resources = append(resources, externalDbSystemDiscovery)
	if err := s.D.Set("external_db_system_discovery_collection", resources); err != nil {
		return err
	}

	return nil
}
