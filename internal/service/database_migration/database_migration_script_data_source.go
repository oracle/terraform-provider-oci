// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_migration

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database_migration "github.com/oracle/oci-go-sdk/v65/databasemigration"
	"github.com/oracle/terraform-provider-oci/internal/client"

	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseMigrationScriptDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatabaseMigrationScript,
		Schema: map[string]*schema.Schema{
			"script_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
		},
	}
}

func readSingularDatabaseMigrationScript(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseMigrationScriptDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseMigrationClient()

	return tfresource.ReadResource(sync)
}

type DatabaseMigrationScriptDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_migration.DatabaseMigrationClient
	Res    *oci_database_migration.GetScriptResponse
}

func (s *DatabaseMigrationScriptDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseMigrationScriptDataSourceCrud) Get() error {
	request := oci_database_migration.GetScriptRequest{}

	if scriptId, ok := s.D.GetOkExists("script_id"); ok {
		request.ScriptId = oci_database_migration.GetScriptScriptIdEnum(oci_database_migration.ScriptIdsEnum(scriptId.(string)))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_migration")

	response, err := s.Client.GetScript(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseMigrationScriptDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseMigrationScriptDataSource-", DatabaseMigrationScriptDataSource(), s.D))

	return nil
}
