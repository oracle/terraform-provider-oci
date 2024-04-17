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

func DatabaseManagementExternalMySqlDatabaseDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["external_my_sql_database_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DatabaseManagementExternalMySqlDatabaseResource(), fieldMap, readSingularDatabaseManagementExternalMySqlDatabase)
}

func readSingularDatabaseManagementExternalMySqlDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalMySqlDatabaseDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementExternalMySqlDatabaseDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.GetExternalMySqlDatabaseResponse
}

func (s *DatabaseManagementExternalMySqlDatabaseDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementExternalMySqlDatabaseDataSourceCrud) Get() error {
	request := oci_database_management.GetExternalMySqlDatabaseRequest{}

	if externalMySqlDatabaseId, ok := s.D.GetOkExists("external_my_sql_database_id"); ok {
		tmp := externalMySqlDatabaseId.(string)
		request.ExternalMySqlDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.GetExternalMySqlDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseManagementExternalMySqlDatabaseDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementExternalMySqlDatabaseDataSource-", DatabaseManagementExternalMySqlDatabaseDataSource(), s.D))

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DbName != nil {
		s.D.Set("db_name", *s.Res.DbName)
	}

	if s.Res.ExternalDatabaseId != nil {
		s.D.Set("external_database_id", *s.Res.ExternalDatabaseId)
	}

	return nil
}
