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

func DatabaseManagementExternalMySqlDatabaseConnectorsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseManagementExternalMySqlDatabaseConnectors,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"my_sql_connector_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(DatabaseManagementExternalMySqlDatabaseConnectorResource()),
						},
					},
				},
			},
		},
	}
}

func readDatabaseManagementExternalMySqlDatabaseConnectors(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalMySqlDatabaseConnectorsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementExternalMySqlDatabaseConnectorsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.ListMySqlDatabaseConnectorsResponse
}

func (s *DatabaseManagementExternalMySqlDatabaseConnectorsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementExternalMySqlDatabaseConnectorsDataSourceCrud) Get() error {
	request := oci_database_management.ListMySqlDatabaseConnectorsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.ListMySqlDatabaseConnectors(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListMySqlDatabaseConnectors(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseManagementExternalMySqlDatabaseConnectorsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementExternalMySqlDatabaseConnectorsDataSource-", DatabaseManagementExternalMySqlDatabaseConnectorsDataSource(), s.D))
	resources := []map[string]interface{}{}
	externalMySqlDatabaseConnector := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, MySqlDatabaseConnectorSummaryToMap(item))
	}
	externalMySqlDatabaseConnector["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseManagementExternalMySqlDatabaseConnectorsDataSource().Schema["my_sql_connector_collection"].Elem.(*schema.Resource).Schema)
		externalMySqlDatabaseConnector["items"] = items
	}

	resources = append(resources, externalMySqlDatabaseConnector)
	if err := s.D.Set("my_sql_connector_collection", resources); err != nil {
		return err
	}

	return nil
}
