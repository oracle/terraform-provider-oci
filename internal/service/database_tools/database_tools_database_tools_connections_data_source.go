// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_tools

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_database_tools "github.com/oracle/oci-go-sdk/v56/databasetools"
)

func DatabaseToolsDatabaseToolsConnectionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseToolsDatabaseToolsConnections,
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
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"database_tools_connection_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(DatabaseToolsDatabaseToolsConnectionResource()),
						},
					},
				},
			},
		},
	}
}

func readDatabaseToolsDatabaseToolsConnections(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseToolsDatabaseToolsConnectionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsClient()

	return tfresource.ReadResource(sync)
}

type DatabaseToolsDatabaseToolsConnectionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_tools.DatabaseToolsClient
	Res    *oci_database_tools.ListDatabaseToolsConnectionsResponse
}

func (s *DatabaseToolsDatabaseToolsConnectionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseToolsDatabaseToolsConnectionsDataSourceCrud) Get() error {
	request := oci_database_tools.ListDatabaseToolsConnectionsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_database_tools.ListDatabaseToolsConnectionsLifecycleStateEnum(state.(string))
	}

	if type_, ok := s.D.GetOkExists("type"); ok {
		interfaces := type_.([]interface{})
		tmp := make([]oci_database_tools.ConnectionTypeEnum, 0)
		for i := range interfaces {
			if interfaces[i] != nil {
				connectionType := interfaces[i].(string)
				if connectionType == "ORACLE_DATABASE" {
					tmp = append(tmp, oci_database_tools.ConnectionTypeOracleDatabase)
				}
			}
		}
		if len(tmp) != 0 {
			request.Type = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_tools")

	response, err := s.Client.ListDatabaseToolsConnections(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDatabaseToolsConnections(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseToolsDatabaseToolsConnectionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseToolsDatabaseToolsConnectionsDataSource-", DatabaseToolsDatabaseToolsConnectionsDataSource(), s.D))
	resources := []map[string]interface{}{}
	databaseToolsConnection := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DatabaseToolsConnectionSummaryToMap(item))
	}
	databaseToolsConnection["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseToolsDatabaseToolsConnectionsDataSource().Schema["database_tools_connection_collection"].Elem.(*schema.Resource).Schema)
		databaseToolsConnection["items"] = items
	}

	resources = append(resources, databaseToolsConnection)
	if err := s.D.Set("database_tools_connection_collection", resources); err != nil {
		return err
	}

	return nil
}
