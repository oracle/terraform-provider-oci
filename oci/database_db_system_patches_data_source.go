// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v27/database"
)

func init() {
	RegisterDatasource("oci_database_db_system_patches", DatabaseDbSystemPatchesDataSource())
}

func DatabaseDbSystemPatchesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseDbSystemPatches,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"db_system_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"patches": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"available_actions": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_action": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"lifecycle_details": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_released": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"version": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readDatabaseDbSystemPatches(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDbSystemPatchesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient()

	return ReadResource(sync)
}

type DatabaseDbSystemPatchesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListDbSystemPatchesResponse
}

func (s *DatabaseDbSystemPatchesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseDbSystemPatchesDataSourceCrud) Get() error {
	request := oci_database.ListDbSystemPatchesRequest{}

	if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
		tmp := dbSystemId.(string)
		request.DbSystemId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "database")

	response, err := s.Client.ListDbSystemPatches(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDbSystemPatches(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseDbSystemPatchesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("DatabaseDbSystemPatchesDataSource-", DatabaseDbSystemPatchesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		dbSystemPatch := map[string]interface{}{}

		dbSystemPatch["available_actions"] = r.AvailableActions

		if r.Description != nil {
			dbSystemPatch["description"] = *r.Description
		}

		if r.Id != nil {
			dbSystemPatch["id"] = *r.Id
		}

		dbSystemPatch["last_action"] = r.LastAction

		if r.LifecycleDetails != nil {
			dbSystemPatch["lifecycle_details"] = *r.LifecycleDetails
		}

		dbSystemPatch["state"] = r.LifecycleState

		if r.TimeReleased != nil {
			dbSystemPatch["time_released"] = r.TimeReleased.String()
		}

		if r.Version != nil {
			dbSystemPatch["version"] = *r.Version
		}

		resources = append(resources, dbSystemPatch)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, DatabaseDbSystemPatchesDataSource().Schema["patches"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("patches", resources); err != nil {
		return err
	}

	return nil
}
