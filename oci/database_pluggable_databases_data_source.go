// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v43/database"
)

func init() {
	RegisterDatasource("oci_database_pluggable_databases", DatabasePluggableDatabasesDataSource())
}

func DatabasePluggableDatabasesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabasePluggableDatabases,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"database_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"pdb_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"pluggable_databases": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(DatabasePluggableDatabaseResource()),
			},
		},
	}
}

func readDatabasePluggableDatabases(d *schema.ResourceData, m interface{}) error {
	sync := &DatabasePluggableDatabasesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient()

	return ReadResource(sync)
}

type DatabasePluggableDatabasesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListPluggableDatabasesResponse
}

func (s *DatabasePluggableDatabasesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabasePluggableDatabasesDataSourceCrud) Get() error {
	request := oci_database.ListPluggableDatabasesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if databaseId, ok := s.D.GetOkExists("database_id"); ok {
		tmp := databaseId.(string)
		request.DatabaseId = &tmp
	}

	if pdbName, ok := s.D.GetOkExists("pdb_name"); ok {
		tmp := pdbName.(string)
		request.PdbName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_database.PluggableDatabaseSummaryLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "database")

	response, err := s.Client.ListPluggableDatabases(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListPluggableDatabases(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabasePluggableDatabasesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("DatabasePluggableDatabasesDataSource-", DatabasePluggableDatabasesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		pluggableDatabase := map[string]interface{}{}

		if r.CompartmentId != nil {
			pluggableDatabase["compartment_id"] = *r.CompartmentId
		}

		if r.ConnectionStrings != nil {
			pluggableDatabase["connection_strings"] = []interface{}{PluggableDatabaseConnectionStringsToMap(r.ConnectionStrings)}
		} else {
			pluggableDatabase["connection_strings"] = nil
		}

		if r.ContainerDatabaseId != nil {
			pluggableDatabase["container_database_id"] = *r.ContainerDatabaseId
		}

		if r.DefinedTags != nil {
			pluggableDatabase["defined_tags"] = definedTagsToMap(r.DefinedTags)
		}

		pluggableDatabase["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			pluggableDatabase["id"] = *r.Id
		}

		if r.IsRestricted != nil {
			pluggableDatabase["is_restricted"] = *r.IsRestricted
		}

		if r.LifecycleDetails != nil {
			pluggableDatabase["lifecycle_details"] = *r.LifecycleDetails
		}

		pluggableDatabase["open_mode"] = r.OpenMode

		if r.PdbName != nil {
			pluggableDatabase["pdb_name"] = *r.PdbName
		}

		pluggableDatabase["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			pluggableDatabase["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, pluggableDatabase)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, DatabasePluggableDatabasesDataSource().Schema["pluggable_databases"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("pluggable_databases", resources); err != nil {
		return err
	}

	return nil
}
