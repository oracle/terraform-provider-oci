// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/database"

	"github.com/oracle/terraform-provider-oci/crud"
)

func PatchHistoryEntriesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readPatchHistoryEntries,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"db_home_id": {
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

func readPatchHistoryEntries(d *schema.ResourceData, m interface{}) error {
	sync := &PatchHistoryEntriesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return crud.ReadResource(sync)
}

type PatchHistoryEntriesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListDbHomePatchHistoryEntriesResponse
}

func (s *PatchHistoryEntriesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *PatchHistoryEntriesDataSourceCrud) Get() error {
	request := oci_database.ListDbHomePatchHistoryEntriesRequest{}

	if dbHomeId, ok := s.D.GetOkExists("db_home_id"); ok {
		tmp := dbHomeId.(string)
		request.DbHomeId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "database")

	response, err := s.Client.ListDbHomePatchHistoryEntries(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDbHomePatchHistoryEntries(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *PatchHistoryEntriesDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		patchHistoryEntry := map[string]interface{}{}

		patchHistoryEntry["action"] = r.Action

		if r.Id != nil {
			patchHistoryEntry["id"] = *r.Id
		}

		if r.LifecycleDetails != nil {
			patchHistoryEntry["lifecycle_details"] = *r.LifecycleDetails
		}

		if r.PatchId != nil {
			patchHistoryEntry["patch_id"] = *r.PatchId
		}

		patchHistoryEntry["state"] = r.LifecycleState

		if r.TimeEnded != nil {
			patchHistoryEntry["time_ended"] = r.TimeEnded.String()
		}

		if r.TimeStarted != nil {
			patchHistoryEntry["time_started"] = r.TimeStarted.String()
		}

		resources = append(resources, patchHistoryEntry)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources)
	}

	if err := s.D.Set("patch_history_entries", resources); err != nil {
		panic(err)
	}

	return
}
