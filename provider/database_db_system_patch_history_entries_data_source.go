// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/database"

	"github.com/oracle/terraform-provider-oci/crud"
)

func DbSystemPatchHistoryEntriesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDbSystemPatchHistoryEntries,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"db_system_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"patch_history_entries": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"action": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"lifecycle_details": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"patch_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_ended": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_started": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readDbSystemPatchHistoryEntries(d *schema.ResourceData, m interface{}) error {
	sync := &DbSystemPatchHistoryEntriesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return crud.ReadResource(sync)
}

type DbSystemPatchHistoryEntriesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListDbSystemPatchHistoryEntriesResponse
}

func (s *DbSystemPatchHistoryEntriesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DbSystemPatchHistoryEntriesDataSourceCrud) Get() error {
	request := oci_database.ListDbSystemPatchHistoryEntriesRequest{}

	if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
		tmp := dbSystemId.(string)
		request.DbSystemId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "database")

	response, err := s.Client.ListDbSystemPatchHistoryEntries(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDbSystemPatchHistoryEntries(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DbSystemPatchHistoryEntriesDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		dbSystemPatchHistoryEntry := map[string]interface{}{}

		dbSystemPatchHistoryEntry["action"] = r.Action

		if r.Id != nil {
			dbSystemPatchHistoryEntry["id"] = *r.Id
		}

		if r.LifecycleDetails != nil {
			dbSystemPatchHistoryEntry["lifecycle_details"] = *r.LifecycleDetails
		}

		if r.PatchId != nil {
			dbSystemPatchHistoryEntry["patch_id"] = *r.PatchId
		}

		dbSystemPatchHistoryEntry["state"] = r.LifecycleState

		if r.TimeEnded != nil {
			dbSystemPatchHistoryEntry["time_ended"] = r.TimeEnded.String()
		}

		if r.TimeStarted != nil {
			dbSystemPatchHistoryEntry["time_started"] = r.TimeStarted.String()
		}

		resources = append(resources, dbSystemPatchHistoryEntry)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, DbSystemPatchHistoryEntriesDataSource().Schema["patch_history_entries"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("patch_history_entries", resources); err != nil {
		panic(err)
	}

	return
}
