// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_migration

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database_migration "github.com/oracle/oci-go-sdk/v65/databasemigration"
)

func DatabaseMigrationMigrationDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["migration_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DatabaseMigrationMigrationResource(), fieldMap, readSingularDatabaseMigrationMigration)
}

func readSingularDatabaseMigrationMigration(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseMigrationMigrationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseMigrationClient()

	return tfresource.ReadResource(sync)
}

type DatabaseMigrationMigrationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_migration.DatabaseMigrationClient
	Res    *oci_database_migration.GetMigrationResponse
}

func (s *DatabaseMigrationMigrationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseMigrationMigrationDataSourceCrud) Get() error {
	request := oci_database_migration.GetMigrationRequest{}

	if migrationId, ok := s.D.GetOkExists("migration_id"); ok {
		tmp := migrationId.(string)
		request.MigrationId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_migration")

	response, err := s.Client.GetMigration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseMigrationMigrationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AdvisorSettings != nil {
		s.D.Set("advisor_settings", []interface{}{AdvisorSettingsToMap(s.Res.AdvisorSettings)})
	} else {
		s.D.Set("advisor_settings", nil)
	}

	if s.Res.AgentId != nil {
		s.D.Set("agent_id", *s.Res.AgentId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CredentialsSecretId != nil {
		s.D.Set("credentials_secret_id", *s.Res.CredentialsSecretId)
	}

	if s.Res.DataTransferMediumDetails != nil {
		s.D.Set("data_transfer_medium_details", []interface{}{DataTransferMediumDetailsToMap(s.Res.DataTransferMediumDetails)})
	} else {
		s.D.Set("data_transfer_medium_details", nil)
	}

	if s.Res.DataTransferMediumDetailsV2 != nil {
		dataTransferMediumDetailsV2Array := []interface{}{}
		if dataTransferMediumDetailsV2Map := DataTransferMediumDetailsV2ToMap(&s.Res.DataTransferMediumDetailsV2); dataTransferMediumDetailsV2Map != nil {
			dataTransferMediumDetailsV2Array = append(dataTransferMediumDetailsV2Array, dataTransferMediumDetailsV2Map)
		}
		s.D.Set("data_transfer_medium_details_v2", dataTransferMediumDetailsV2Array)
	} else {
		s.D.Set("data_transfer_medium_details_v2", nil)
	}

	if s.Res.DatapumpSettings != nil {
		s.D.Set("datapump_settings", []interface{}{DataPumpSettingsToMap(s.Res.DatapumpSettings)})
	} else {
		s.D.Set("datapump_settings", nil)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.DumpTransferDetails != nil {
		s.D.Set("dump_transfer_details", []interface{}{DumpTransferDetailsToMap(s.Res.DumpTransferDetails)})
	} else {
		s.D.Set("dump_transfer_details", nil)
	}

	excludeObjects := []interface{}{}
	for _, item := range s.Res.ExcludeObjects {
		excludeObjects = append(excludeObjects, DatabaseObjectToMap(item))
	}
	s.D.Set("exclude_objects", excludeObjects)

	if s.Res.ExecutingJobId != nil {
		s.D.Set("executing_job_id", *s.Res.ExecutingJobId)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.GoldenGateDetails != nil {
		s.D.Set("golden_gate_details", []interface{}{GoldenGateDetailsToMapPass(s.Res.GoldenGateDetails, s.D)})

	} else {
		s.D.Set("golden_gate_details", nil)
	}

	if s.Res.GoldenGateServiceDetails != nil {
		s.D.Set("golden_gate_service_details", []interface{}{GoldenGateServiceDetailsToMap(s.Res.GoldenGateServiceDetails)})
	} else {
		s.D.Set("golden_gate_service_details", nil)
	}

	includeObjects := []interface{}{}
	for _, item := range s.Res.IncludeObjects {
		includeObjects = append(includeObjects, DatabaseObjectToMap(item))
	}
	s.D.Set("include_objects", includeObjects)

	s.D.Set("lifecycle_details", s.Res.LifecycleDetails)

	if s.Res.SourceContainerDatabaseConnectionId != nil {
		s.D.Set("source_container_database_connection_id", *s.Res.SourceContainerDatabaseConnectionId)
	}

	if s.Res.SourceDatabaseConnectionId != nil {
		s.D.Set("source_database_connection_id", *s.Res.SourceDatabaseConnectionId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TargetDatabaseConnectionId != nil {
		s.D.Set("target_database_connection_id", *s.Res.TargetDatabaseConnectionId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastMigration != nil {
		s.D.Set("time_last_migration", s.Res.TimeLastMigration.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	s.D.Set("type", s.Res.Type)

	if s.Res.VaultDetails != nil {
		s.D.Set("vault_details", []interface{}{VaultDetailsToMap(s.Res.VaultDetails)})
	} else {
		s.D.Set("vault_details", nil)
	}

	s.D.Set("wait_after", s.Res.WaitAfter)

	return nil
}
