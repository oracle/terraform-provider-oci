// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/database"

	"github.com/oracle/terraform-provider-oci/crud"
)

func DbHomePatchesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDbHomePatches,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
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

func readDbHomePatches(d *schema.ResourceData, m interface{}) error {
	sync := &DbHomePatchesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return crud.ReadResource(sync)
}

type DbHomePatchesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListDbHomePatchesResponse
}

func (s *DbHomePatchesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DbHomePatchesDataSourceCrud) Get() error {
	request := oci_database.ListDbHomePatchesRequest{}

	if dbHomeId, ok := s.D.GetOkExists("db_home_id"); ok {
		tmp := dbHomeId.(string)
		request.DbHomeId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "database")

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

func (s *DbHomePatchesDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
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
		resources = ApplyFilters(f.(*schema.Set), resources, DbHomePatchesDataSource().Schema["patches"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("patches", resources); err != nil {
		panic(err)
	}

	return
}
