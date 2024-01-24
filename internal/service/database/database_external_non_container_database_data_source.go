// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"
)

func DatabaseExternalNonContainerDatabaseDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["external_non_container_database_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DatabaseExternalNonContainerDatabaseResource(), fieldMap, readSingularDatabaseExternalNonContainerDatabase)
}

func readSingularDatabaseExternalNonContainerDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExternalNonContainerDatabaseDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseExternalNonContainerDatabaseDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.GetExternalNonContainerDatabaseResponse
}

func (s *DatabaseExternalNonContainerDatabaseDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseExternalNonContainerDatabaseDataSourceCrud) Get() error {
	request := oci_database.GetExternalNonContainerDatabaseRequest{}

	if externalNonContainerDatabaseId, ok := s.D.GetOkExists("external_non_container_database_id"); ok {
		tmp := externalNonContainerDatabaseId.(string)
		request.ExternalNonContainerDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.GetExternalNonContainerDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseExternalNonContainerDatabaseDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CharacterSet != nil {
		s.D.Set("character_set", *s.Res.CharacterSet)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("database_configuration", s.Res.DatabaseConfiguration)

	s.D.Set("database_edition", s.Res.DatabaseEdition)

	if s.Res.DatabaseManagementConfig != nil {
		s.D.Set("database_management_config", []interface{}{DatabaseManagementConfigToMap(s.Res.DatabaseManagementConfig)})
	} else {
		s.D.Set("database_management_config", nil)
	}

	if s.Res.DatabaseVersion != nil {
		s.D.Set("database_version", *s.Res.DatabaseVersion)
	}

	if s.Res.DbId != nil {
		s.D.Set("db_id", *s.Res.DbId)
	}

	if s.Res.DbPacks != nil {
		s.D.Set("db_packs", *s.Res.DbPacks)
	}

	if s.Res.DbUniqueName != nil {
		s.D.Set("db_unique_name", *s.Res.DbUniqueName)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.NcharacterSet != nil {
		s.D.Set("ncharacter_set", *s.Res.NcharacterSet)
	}

	if s.Res.OperationsInsightsConfig != nil {
		s.D.Set("operations_insights_config", []interface{}{OperationsInsightsConfigToMap(s.Res.OperationsInsightsConfig)})
	} else {
		s.D.Set("operations_insights_config", nil)
	}

	if s.Res.StackMonitoringConfig != nil {
		s.D.Set("stack_monitoring_config", []interface{}{StackMonitoringConfigToMap(s.Res.StackMonitoringConfig)})
	} else {
		s.D.Set("stack_monitoring_config", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeZone != nil {
		s.D.Set("time_zone", *s.Res.TimeZone)
	}

	return nil
}
