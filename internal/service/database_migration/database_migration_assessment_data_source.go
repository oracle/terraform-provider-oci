// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_migration

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database_migration "github.com/oracle/oci-go-sdk/v65/databasemigration"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseMigrationAssessmentDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["assessment_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DatabaseMigrationAssessmentResource(), fieldMap, readSingularDatabaseMigrationAssessment)
}

func readSingularDatabaseMigrationAssessment(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseMigrationAssessmentDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseMigrationClient()

	return tfresource.ReadResource(sync)
}

type DatabaseMigrationAssessmentDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_migration.DatabaseMigrationClient
	Res    *oci_database_migration.GetAssessmentResponse
}

func (s *DatabaseMigrationAssessmentDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseMigrationAssessmentDataSourceCrud) Get() error {
	request := oci_database_migration.GetAssessmentRequest{}

	if assessmentId, ok := s.D.GetOkExists("assessment_id"); ok {
		tmp := assessmentId.(string)
		request.AssessmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_migration")

	response, err := s.Client.GetAssessment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseMigrationAssessmentDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.GetId())
	switch v := (s.Res.Assessment).(type) {
	case oci_database_migration.MySqlAssessment:
		s.D.Set("database_combination", "MYSQL")

		s.D.Set("acceptable_downtime", v.AcceptableDowntime)

		s.D.Set("assessment_migration_type", v.AssessmentMigrationType)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		s.D.Set("creation_type", v.CreationType)

		s.D.Set("database_data_size", v.DatabaseDataSize)

		s.D.Set("ddl_expectation", v.DdlExpectation)

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

		if v.MigrationId != nil {
			s.D.Set("migration_id", *v.MigrationId)
		}

		s.D.Set("network_speed_megabit_per_second", v.NetworkSpeedMegabitPerSecond)

		if v.SourceDatabaseConnection != nil {
			s.D.Set("source_database_connection", []interface{}{SourceAssessmentConnectionToMap(v.SourceDatabaseConnection)})
		} else {
			s.D.Set("source_database_connection", nil)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TargetDatabaseConnection != nil {
			s.D.Set("target_database_connection", []interface{}{TargetAssessmentConnectionToMap(v.TargetDatabaseConnection)})
		} else {
			s.D.Set("target_database_connection", nil)
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_database_migration.OracleAssessment:
		s.D.Set("database_combination", "ORACLE")

		if v.IsCdbSupported != nil {
			s.D.Set("is_cdb_supported", *v.IsCdbSupported)
		}

		s.D.Set("acceptable_downtime", v.AcceptableDowntime)

		s.D.Set("assessment_migration_type", v.AssessmentMigrationType)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		s.D.Set("creation_type", v.CreationType)

		s.D.Set("database_data_size", v.DatabaseDataSize)

		s.D.Set("ddl_expectation", v.DdlExpectation)

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

		if v.MigrationId != nil {
			s.D.Set("migration_id", *v.MigrationId)
		}

		s.D.Set("network_speed_megabit_per_second", v.NetworkSpeedMegabitPerSecond)

		if v.SourceDatabaseConnection != nil {
			s.D.Set("source_database_connection", []interface{}{SourceAssessmentConnectionToMap(v.SourceDatabaseConnection)})
		} else {
			s.D.Set("source_database_connection", nil)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TargetDatabaseConnection != nil {
			s.D.Set("target_database_connection", []interface{}{TargetAssessmentConnectionToMap(v.TargetDatabaseConnection)})
		} else {
			s.D.Set("target_database_connection", nil)
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	default:
		log.Printf("[WARN] Received 'database_combination' of unknown type %v", s.Res.Assessment)
		return nil
	}

	return nil
}
