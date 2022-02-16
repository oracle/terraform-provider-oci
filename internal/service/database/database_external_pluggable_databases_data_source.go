// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v58/database"
)

func DatabaseExternalPluggableDatabasesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseExternalPluggableDatabases,
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
			"external_container_database_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_pluggable_databases": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(DatabaseExternalPluggableDatabaseResource()),
			},
		},
	}
}

func readDatabaseExternalPluggableDatabases(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExternalPluggableDatabasesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseExternalPluggableDatabasesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListExternalPluggableDatabasesResponse
}

func (s *DatabaseExternalPluggableDatabasesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseExternalPluggableDatabasesDataSourceCrud) Get() error {
	request := oci_database.ListExternalPluggableDatabasesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if externalContainerDatabaseId, ok := s.D.GetOkExists("external_container_database_id"); ok {
		tmp := externalContainerDatabaseId.(string)
		request.ExternalContainerDatabaseId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_database.ExternalDatabaseBaseLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListExternalPluggableDatabases(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListExternalPluggableDatabases(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseExternalPluggableDatabasesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseExternalPluggableDatabasesDataSource-", DatabaseExternalPluggableDatabasesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		externalPluggableDatabase := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.CharacterSet != nil {
			externalPluggableDatabase["character_set"] = *r.CharacterSet
		}

		externalPluggableDatabase["database_configuration"] = r.DatabaseConfiguration

		externalPluggableDatabase["database_edition"] = r.DatabaseEdition

		if r.DatabaseManagementConfig != nil {
			externalPluggableDatabase["database_management_config"] = []interface{}{DatabaseManagementConfigToMap(r.DatabaseManagementConfig)}
		} else {
			externalPluggableDatabase["database_management_config"] = nil
		}

		if r.DatabaseVersion != nil {
			externalPluggableDatabase["database_version"] = *r.DatabaseVersion
		}

		if r.DbId != nil {
			externalPluggableDatabase["db_id"] = *r.DbId
		}

		if r.DbPacks != nil {
			externalPluggableDatabase["db_packs"] = *r.DbPacks
		}

		if r.DbUniqueName != nil {
			externalPluggableDatabase["db_unique_name"] = *r.DbUniqueName
		}

		if r.DefinedTags != nil {
			externalPluggableDatabase["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			externalPluggableDatabase["display_name"] = *r.DisplayName
		}

		if r.ExternalContainerDatabaseId != nil {
			externalPluggableDatabase["external_container_database_id"] = *r.ExternalContainerDatabaseId
		}

		externalPluggableDatabase["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			externalPluggableDatabase["id"] = *r.Id
		}

		if r.LifecycleDetails != nil {
			externalPluggableDatabase["lifecycle_details"] = *r.LifecycleDetails
		}

		if r.NcharacterSet != nil {
			externalPluggableDatabase["ncharacter_set"] = *r.NcharacterSet
		}

		if r.OperationsInsightsConfig != nil {
			externalPluggableDatabase["operations_insights_config"] = []interface{}{OperationsInsightsConfigToMap(r.OperationsInsightsConfig)}
		} else {
			externalPluggableDatabase["operations_insights_config"] = nil
		}

		if r.SourceId != nil {
			externalPluggableDatabase["source_id"] = *r.SourceId
		}

		externalPluggableDatabase["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			externalPluggableDatabase["time_created"] = r.TimeCreated.String()
		}

		if r.TimeZone != nil {
			externalPluggableDatabase["time_zone"] = *r.TimeZone
		}

		resources = append(resources, externalPluggableDatabase)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabaseExternalPluggableDatabasesDataSource().Schema["external_pluggable_databases"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("external_pluggable_databases", resources); err != nil {
		return err
	}

	return nil
}
