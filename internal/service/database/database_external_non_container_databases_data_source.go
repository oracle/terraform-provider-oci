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

func DatabaseExternalNonContainerDatabasesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseExternalNonContainerDatabases,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
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
			"external_non_container_databases": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(DatabaseExternalNonContainerDatabaseResource()),
			},
		},
	}
}

func readDatabaseExternalNonContainerDatabases(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExternalNonContainerDatabasesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseExternalNonContainerDatabasesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListExternalNonContainerDatabasesResponse
}

func (s *DatabaseExternalNonContainerDatabasesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseExternalNonContainerDatabasesDataSourceCrud) Get() error {
	request := oci_database.ListExternalNonContainerDatabasesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_database.ExternalDatabaseBaseLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListExternalNonContainerDatabases(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListExternalNonContainerDatabases(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseExternalNonContainerDatabasesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseExternalNonContainerDatabasesDataSource-", DatabaseExternalNonContainerDatabasesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		externalNonContainerDatabase := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.CharacterSet != nil {
			externalNonContainerDatabase["character_set"] = *r.CharacterSet
		}

		externalNonContainerDatabase["database_configuration"] = r.DatabaseConfiguration

		externalNonContainerDatabase["database_edition"] = r.DatabaseEdition

		if r.DatabaseManagementConfig != nil {
			externalNonContainerDatabase["database_management_config"] = []interface{}{DatabaseManagementConfigToMap(r.DatabaseManagementConfig)}
		} else {
			externalNonContainerDatabase["database_management_config"] = nil
		}

		if r.DatabaseVersion != nil {
			externalNonContainerDatabase["database_version"] = *r.DatabaseVersion
		}

		if r.DbId != nil {
			externalNonContainerDatabase["db_id"] = *r.DbId
		}

		if r.DbPacks != nil {
			externalNonContainerDatabase["db_packs"] = *r.DbPacks
		}

		if r.DbUniqueName != nil {
			externalNonContainerDatabase["db_unique_name"] = *r.DbUniqueName
		}

		if r.DefinedTags != nil {
			externalNonContainerDatabase["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			externalNonContainerDatabase["display_name"] = *r.DisplayName
		}

		externalNonContainerDatabase["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			externalNonContainerDatabase["id"] = *r.Id
		}

		if r.LifecycleDetails != nil {
			externalNonContainerDatabase["lifecycle_details"] = *r.LifecycleDetails
		}

		if r.NcharacterSet != nil {
			externalNonContainerDatabase["ncharacter_set"] = *r.NcharacterSet
		}

		if r.OperationsInsightsConfig != nil {
			externalNonContainerDatabase["operations_insights_config"] = []interface{}{OperationsInsightsConfigToMap(r.OperationsInsightsConfig)}
		} else {
			externalNonContainerDatabase["operations_insights_config"] = nil
		}

		if r.StackMonitoringConfig != nil {
			externalNonContainerDatabase["stack_monitoring_config"] = []interface{}{StackMonitoringConfigToMap(r.StackMonitoringConfig)}
		} else {
			externalNonContainerDatabase["stack_monitoring_config"] = nil
		}

		externalNonContainerDatabase["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			externalNonContainerDatabase["time_created"] = r.TimeCreated.String()
		}

		if r.TimeZone != nil {
			externalNonContainerDatabase["time_zone"] = *r.TimeZone
		}

		resources = append(resources, externalNonContainerDatabase)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabaseExternalNonContainerDatabasesDataSource().Schema["external_non_container_databases"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("external_non_container_databases", resources); err != nil {
		return err
	}

	return nil
}
