// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_tools

import (
	"context"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database_tools "github.com/oracle/oci-go-sdk/v65/databasetools"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseToolsDatabaseToolsSqlReportDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["database_tools_sql_report_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(DatabaseToolsDatabaseToolsSqlReportResource(), fieldMap, readSingularDatabaseToolsDatabaseToolsSqlReportWithContext)
}

func readSingularDatabaseToolsDatabaseToolsSqlReportWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseToolsDatabaseToolsSqlReportDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DatabaseToolsDatabaseToolsSqlReportDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_tools.DatabaseToolsClient
	Res    *oci_database_tools.GetDatabaseToolsSqlReportResponse
}

func (s *DatabaseToolsDatabaseToolsSqlReportDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseToolsDatabaseToolsSqlReportDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_database_tools.GetDatabaseToolsSqlReportRequest{}

	if databaseToolsSqlReportId, ok := s.D.GetOkExists("database_tools_sql_report_id"); ok {
		tmp := databaseToolsSqlReportId.(string)
		request.DatabaseToolsSqlReportId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_tools")

	response, err := s.Client.GetDatabaseToolsSqlReport(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseToolsDatabaseToolsSqlReportDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.GetId())
	switch v := (s.Res.DatabaseToolsSqlReport).(type) {
	case oci_database_tools.DatabaseToolsSqlReportOracleDatabase:
		s.D.Set("type", "ORACLE_DATABASE")

		columns := []interface{}{}
		for _, item := range v.Columns {
			columns = append(columns, DatabaseToolsSqlReportColumnToMap(item))
		}
		s.D.Set("columns", columns)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Instructions != nil {
			s.D.Set("instructions", *v.Instructions)
		}

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, DbtoolsSqlReportResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		if v.Purpose != nil {
			s.D.Set("purpose", *v.Purpose)
		}

		if v.Source != nil {
			s.D.Set("source", *v.Source)
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.Format(time.RFC3339Nano))
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.Format(time.RFC3339Nano))
		}
		s.D.Set("state", v.LifecycleState)

		variables := []interface{}{}
		for _, item := range v.Variables {
			variables = append(variables, DatabaseToolsSqlReportVariableToMap(item))
		}
		s.D.Set("variables", variables)
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", s.Res.DatabaseToolsSqlReport)
		return nil
	}

	return nil
}
