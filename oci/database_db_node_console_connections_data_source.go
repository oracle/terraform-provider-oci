// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v36/database"
)

func init() {
	RegisterDatasource("oci_database_db_node_console_connections", DatabaseDbNodeConsoleConnectionsDataSource())
}

func DatabaseDbNodeConsoleConnectionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseDbNodeConsoleConnections,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"db_node_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"console_connections": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(DatabaseDbNodeConsoleConnectionResource()),
			},
		},
	}
}

func readDatabaseDbNodeConsoleConnections(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDbNodeConsoleConnectionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient()

	return ReadResource(sync)
}

type DatabaseDbNodeConsoleConnectionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListConsoleConnectionsResponse
}

func (s *DatabaseDbNodeConsoleConnectionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseDbNodeConsoleConnectionsDataSourceCrud) Get() error {
	request := oci_database.ListConsoleConnectionsRequest{}

	if dbNodeId, ok := s.D.GetOkExists("db_node_id"); ok {
		tmp := dbNodeId.(string)
		request.DbNodeId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "database")

	response, err := s.Client.ListConsoleConnections(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseDbNodeConsoleConnectionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("DatabaseDbNodeConsoleConnectionsDataSource-", DatabaseDbNodeConsoleConnectionsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		dbNodeConsoleConnection := map[string]interface{}{
			"db_node_id": *r.DbNodeId,
		}

		if r.CompartmentId != nil {
			dbNodeConsoleConnection["compartment_id"] = *r.CompartmentId
		}

		if r.ConnectionString != nil {
			dbNodeConsoleConnection["connection_string"] = *r.ConnectionString
		}

		if r.Fingerprint != nil {
			dbNodeConsoleConnection["fingerprint"] = *r.Fingerprint
		}

		if r.Id != nil {
			dbNodeConsoleConnection["id"] = *r.Id
		}

		dbNodeConsoleConnection["state"] = r.LifecycleState

		resources = append(resources, dbNodeConsoleConnection)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, DatabaseDbNodeConsoleConnectionsDataSource().Schema["console_connections"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("console_connections", resources); err != nil {
		return err
	}

	return nil
}
