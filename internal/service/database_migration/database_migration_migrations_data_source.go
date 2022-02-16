// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_migration

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_database_migration "github.com/oracle/oci-go-sdk/v58/databasemigration"
)

func DatabaseMigrationMigrationsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseMigrationMigrations,
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
			"lifecycle_details": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"migration_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(DatabaseMigrationMigrationResource()),
						},
					},
				},
			},
		},
	}
}

func readDatabaseMigrationMigrations(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseMigrationMigrationsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseMigrationClient()

	return tfresource.ReadResource(sync)
}

type DatabaseMigrationMigrationsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_migration.DatabaseMigrationClient
	Res    *oci_database_migration.ListMigrationsResponse
}

func (s *DatabaseMigrationMigrationsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseMigrationMigrationsDataSourceCrud) Get() error {
	request := oci_database_migration.ListMigrationsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if lifecycleDetails, ok := s.D.GetOkExists("lifecycle_details"); ok {
		request.LifecycleDetails = oci_database_migration.ListMigrationsLifecycleDetailsEnum(lifecycleDetails.(string))
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_database_migration.ListMigrationsLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_migration")

	response, err := s.Client.ListMigrations(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListMigrations(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseMigrationMigrationsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseMigrationMigrationsDataSource-", DatabaseMigrationMigrationsDataSource(), s.D))
	resources := []map[string]interface{}{}
	migration := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, MigrationSummaryToMap(item))
	}
	migration["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseMigrationMigrationsDataSource().Schema["migration_collection"].Elem.(*schema.Resource).Schema)
		migration["items"] = items
	}

	resources = append(resources, migration)
	if err := s.D.Set("migration_collection", resources); err != nil {
		return err
	}

	return nil
}
