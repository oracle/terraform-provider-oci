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

func DatabaseDbHomePatchesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseDbHomePatches,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"db_home_id": {
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

func readDatabaseDbHomePatches(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDbHomePatchesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseDbHomePatchesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListDbHomePatchesResponse
}

func (s *DatabaseDbHomePatchesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseDbHomePatchesDataSourceCrud) Get() error {
	request := oci_database.ListDbHomePatchesRequest{}

	if dbHomeId, ok := s.D.GetOkExists("db_home_id"); ok {
		tmp := dbHomeId.(string)
		request.DbHomeId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListDbHomePatches(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDbHomePatches(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseDbHomePatchesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseDbHomePatchesDataSource-", DatabaseDbHomePatchesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		dbHomePatch := map[string]interface{}{}

		dbHomePatch["available_actions"] = r.AvailableActions

		if r.Description != nil {
			dbHomePatch["description"] = *r.Description
		}

		if r.Id != nil {
			dbHomePatch["id"] = *r.Id
		}

		dbHomePatch["last_action"] = r.LastAction

		if r.LifecycleDetails != nil {
			dbHomePatch["lifecycle_details"] = *r.LifecycleDetails
		}

		dbHomePatch["state"] = r.LifecycleState

		if r.TimeReleased != nil {
			dbHomePatch["time_released"] = r.TimeReleased.String()
		}

		if r.Version != nil {
			dbHomePatch["version"] = *r.Version
		}

		resources = append(resources, dbHomePatch)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabaseDbHomePatchesDataSource().Schema["patches"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("patches", resources); err != nil {
		return err
	}

	return nil
}
