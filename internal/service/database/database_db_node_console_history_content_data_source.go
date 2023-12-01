// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseDbNodeConsoleHistoryContentDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatabaseDbNodeConsoleHistoryContent,
		Schema: map[string]*schema.Schema{
			"console_history_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"db_node_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
		},
	}
}

func readSingularDatabaseDbNodeConsoleHistoryContent(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDbNodeConsoleHistoryContentDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseDbNodeConsoleHistoryContentDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.GetConsoleHistoryContentResponse
}

func (s *DatabaseDbNodeConsoleHistoryContentDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseDbNodeConsoleHistoryContentDataSourceCrud) Get() error {
	request := oci_database.GetConsoleHistoryContentRequest{}

	if consoleHistoryId, ok := s.D.GetOkExists("console_history_id"); ok {
		tmp := consoleHistoryId.(string)
		request.ConsoleHistoryId = &tmp
	}

	if dbNodeId, ok := s.D.GetOkExists("db_node_id"); ok {
		tmp := dbNodeId.(string)
		request.DbNodeId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.GetConsoleHistoryContent(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseDbNodeConsoleHistoryContentDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseDbNodeConsoleHistoryContentDataSource-", DatabaseDbNodeConsoleHistoryContentDataSource(), s.D))

	return nil
}
