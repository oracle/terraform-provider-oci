// // Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// // Licensed under the Mozilla Public License v2.0
package database_migration

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database_migration "github.com/oracle/oci-go-sdk/v65/databasemigration"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
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

	s.D.SetId(*s.Res.GetId())
	switch v := (s.Res.Migration).(type) {
	case oci_database_migration.MySqlMigration:
		s.D.Set("database_combination", "MYSQL")

		if v.AdvisorSettings != nil {
			s.D.Set("advisor_settings", []interface{}{MySqlAdvisorSettingsToMap(v.AdvisorSettings)})
		} else {
			s.D.Set("advisor_settings", nil)
		}

		if v.DataTransferMediumDetails != nil {
			dataTransferMediumDetailsArray := []interface{}{}
			if dataTransferMediumDetailsMap := MySqlDataTransferMediumDetailsToMap(&v.DataTransferMediumDetails); dataTransferMediumDetailsMap != nil {
				dataTransferMediumDetailsArray = append(dataTransferMediumDetailsArray, dataTransferMediumDetailsMap)
			}
			s.D.Set("data_transfer_medium_details", dataTransferMediumDetailsArray)
		} else {
			s.D.Set("data_transfer_medium_details", nil)
		}

		if v.GgsDetails != nil {
			s.D.Set("ggs_details", []interface{}{MySqlGgsDeploymentDetailsToMap(v.GgsDetails)})
		} else {
			s.D.Set("ggs_details", nil)
		}

		if v.HubDetails != nil {
			s.D.Set("hub_details", []interface{}{GoldenGateHubDetailsToMap(v.HubDetails)})
		} else {
			s.D.Set("hub_details", nil)
		}

		if v.InitialLoadSettings != nil {
			s.D.Set("initial_load_settings", []interface{}{MySqlInitialLoadSettingsToMap(v.InitialLoadSettings)})
		} else {
			s.D.Set("initial_load_settings", nil)
		}

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

		if v.ExecutingJobId != nil {
			s.D.Set("executing_job_id", *v.ExecutingJobId)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		s.D.Set("lifecycle_details", v.LifecycleDetails)

		if v.SourceDatabaseConnectionId != nil {
			s.D.Set("source_database_connection_id", *v.SourceDatabaseConnectionId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TargetDatabaseConnectionId != nil {
			s.D.Set("target_database_connection_id", *v.TargetDatabaseConnectionId)
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeLastMigration != nil {
			s.D.Set("time_last_migration", v.TimeLastMigration.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		s.D.Set("type", v.Type)

		s.D.Set("wait_after", v.WaitAfter)
	case oci_database_migration.OracleMigration:
		s.D.Set("database_combination", "ORACLE")

		advancedParameters := []interface{}{}
		for _, item := range v.AdvancedParameters {
			advancedParameters = append(advancedParameters, migrationParameterDetailsToMap(item))
		}
		s.D.Set("advanced_parameters", advancedParameters)

		if v.AdvisorSettings != nil {
			s.D.Set("advisor_settings", []interface{}{OracleAdvisorSettingsToMap(v.AdvisorSettings)})
		} else {
			s.D.Set("advisor_settings", nil)
		}

		if v.DataTransferMediumDetails != nil {
			dataTransferMediumDetailsArray := []interface{}{}
			if dataTransferMediumDetailsMap := OracleDataTransferMediumDetailsToMap(&v.DataTransferMediumDetails); dataTransferMediumDetailsMap != nil {
				dataTransferMediumDetailsArray = append(dataTransferMediumDetailsArray, dataTransferMediumDetailsMap)
			}
			s.D.Set("data_transfer_medium_details", dataTransferMediumDetailsArray)
		} else {
			s.D.Set("data_transfer_medium_details", nil)
		}

		if v.GgsDetails != nil {
			s.D.Set("ggs_details", []interface{}{OracleGgsDeploymentDetailsToMap(v.GgsDetails)})
		} else {
			s.D.Set("ggs_details", nil)
		}

		if v.HubDetails != nil {
			s.D.Set("hub_details", []interface{}{GoldenGateHubDetailsToMap(v.HubDetails)})
		} else {
			s.D.Set("hub_details", nil)
		}

		if v.InitialLoadSettings != nil {
			s.D.Set("initial_load_settings", []interface{}{OracleInitialLoadSettingsToMap(v.InitialLoadSettings)})
		} else {
			s.D.Set("initial_load_settings", nil)
		}

		if v.SourceContainerDatabaseConnectionId != nil {
			s.D.Set("source_container_database_connection_id", *v.SourceContainerDatabaseConnectionId)
		}

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

		if v.ExecutingJobId != nil {
			s.D.Set("executing_job_id", *v.ExecutingJobId)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		s.D.Set("lifecycle_details", v.LifecycleDetails)

		if v.SourceDatabaseConnectionId != nil {
			s.D.Set("source_database_connection_id", *v.SourceDatabaseConnectionId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TargetDatabaseConnectionId != nil {
			s.D.Set("target_database_connection_id", *v.TargetDatabaseConnectionId)
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeLastMigration != nil {
			s.D.Set("time_last_migration", v.TimeLastMigration.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		s.D.Set("type", v.Type)

		s.D.Set("wait_after", v.WaitAfter)
	default:
		log.Printf("[WARN] Received 'database_combination' of unknown type %v", s.Res.Migration)
		return nil
	}

	return nil
}
