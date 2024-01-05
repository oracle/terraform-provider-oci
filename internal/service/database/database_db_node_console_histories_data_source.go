// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseDbNodeConsoleHistoriesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseDbNodeConsoleHistories,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"db_node_id": {
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
			"console_history_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(DatabaseDbNodeConsoleHistoryResource()),
						},
					},
				},
			},
		},
	}
}

func readDatabaseDbNodeConsoleHistories(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDbNodeConsoleHistoriesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseDbNodeConsoleHistoriesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListConsoleHistoriesResponse
}

func (s *DatabaseDbNodeConsoleHistoriesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseDbNodeConsoleHistoriesDataSourceCrud) Get() error {
	request := oci_database.ListConsoleHistoriesRequest{}

	if dbNodeId, ok := s.D.GetOkExists("db_node_id"); ok {
		tmp := dbNodeId.(string)
		request.DbNodeId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_database.ConsoleHistorySummaryLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListConsoleHistories(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListConsoleHistories(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseDbNodeConsoleHistoriesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseDbNodeConsoleHistoriesDataSource-", DatabaseDbNodeConsoleHistoriesDataSource(), s.D))
	resources := []map[string]interface{}{}
	dbNodeConsoleHistory := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ConsoleHistorySummaryToMap(item))
	}
	dbNodeConsoleHistory["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseDbNodeConsoleHistoriesDataSource().Schema["console_history_collection"].Elem.(*schema.Resource).Schema)
		dbNodeConsoleHistory["items"] = items
	}

	resources = append(resources, dbNodeConsoleHistory)
	if err := s.D.Set("console_history_collection", resources); err != nil {
		return err
	}

	return nil
}
