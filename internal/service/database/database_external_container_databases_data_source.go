// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v56/database"
)

func DatabaseExternalContainerDatabasesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseExternalContainerDatabases,
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
			"external_container_databases": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(DatabaseExternalContainerDatabaseResource()),
			},
		},
	}
}

func readDatabaseExternalContainerDatabases(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExternalContainerDatabasesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseExternalContainerDatabasesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListExternalContainerDatabasesResponse
}

func (s *DatabaseExternalContainerDatabasesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseExternalContainerDatabasesDataSourceCrud) Get() error {
	request := oci_database.ListExternalContainerDatabasesRequest{}

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

	response, err := s.Client.ListExternalContainerDatabases(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListExternalContainerDatabases(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseExternalContainerDatabasesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseExternalContainerDatabasesDataSource-", DatabaseExternalContainerDatabasesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		externalContainerDatabase := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.CharacterSet != nil {
			externalContainerDatabase["character_set"] = *r.CharacterSet
		}

		externalContainerDatabase["database_configuration"] = r.DatabaseConfiguration

		externalContainerDatabase["database_edition"] = r.DatabaseEdition

		if r.DatabaseManagementConfig != nil {
			externalContainerDatabase["database_management_config"] = []interface{}{DatabaseManagementConfigToMap(r.DatabaseManagementConfig)}
		} else {
			externalContainerDatabase["database_management_config"] = nil
		}

		if r.DatabaseVersion != nil {
			externalContainerDatabase["database_version"] = *r.DatabaseVersion
		}

		if r.DbId != nil {
			externalContainerDatabase["db_id"] = *r.DbId
		}

		if r.DbPacks != nil {
			externalContainerDatabase["db_packs"] = *r.DbPacks
		}

		if r.DbUniqueName != nil {
			externalContainerDatabase["db_unique_name"] = *r.DbUniqueName
		}

		if r.DefinedTags != nil {
			externalContainerDatabase["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			externalContainerDatabase["display_name"] = *r.DisplayName
		}

		externalContainerDatabase["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			externalContainerDatabase["id"] = *r.Id
		}

		if r.LifecycleDetails != nil {
			externalContainerDatabase["lifecycle_details"] = *r.LifecycleDetails
		}

		if r.NcharacterSet != nil {
			externalContainerDatabase["ncharacter_set"] = *r.NcharacterSet
		}

		externalContainerDatabase["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			externalContainerDatabase["time_created"] = r.TimeCreated.String()
		}

		if r.TimeZone != nil {
			externalContainerDatabase["time_zone"] = *r.TimeZone
		}

		resources = append(resources, externalContainerDatabase)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabaseExternalContainerDatabasesDataSource().Schema["external_container_databases"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("external_container_databases", resources); err != nil {
		return err
	}

	return nil
}
