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

func DatabaseDatabasePdbConversionHistoryEntriesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseDatabasePdbConversionHistoryEntries,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"pdb_conversion_action": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"pdb_conversion_history_entries": {
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
						"additional_cdb_params": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"cdb_name": {
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
						"source_database_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"target": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"target_database_id": {
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

func readDatabaseDatabasePdbConversionHistoryEntries(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDatabasePdbConversionHistoryEntriesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseDatabasePdbConversionHistoryEntriesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListPdbConversionHistoryEntriesResponse
}

func (s *DatabaseDatabasePdbConversionHistoryEntriesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseDatabasePdbConversionHistoryEntriesDataSourceCrud) Get() error {
	request := oci_database.ListPdbConversionHistoryEntriesRequest{}

	if databaseId, ok := s.D.GetOkExists("database_id"); ok {
		tmp := databaseId.(string)
		request.DatabaseId = &tmp
	}

	if pdbConversionAction, ok := s.D.GetOkExists("pdb_conversion_action"); ok {
		request.PdbConversionAction = oci_database.PdbConversionHistoryEntrySummaryActionEnum(pdbConversionAction.(string))
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_database.PdbConversionHistoryEntrySummaryLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListPdbConversionHistoryEntries(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListPdbConversionHistoryEntries(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseDatabasePdbConversionHistoryEntriesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseDatabasePdbConversionHistoryEntriesDataSource-", DatabaseDatabasePdbConversionHistoryEntriesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		databasePdbConversionHistoryEntry := map[string]interface{}{}

		databasePdbConversionHistoryEntry["action"] = r.Action

		if r.AdditionalCdbParams != nil {
			databasePdbConversionHistoryEntry["additional_cdb_params"] = *r.AdditionalCdbParams
		}

		if r.CdbName != nil {
			databasePdbConversionHistoryEntry["cdb_name"] = *r.CdbName
		}

		if r.Id != nil {
			databasePdbConversionHistoryEntry["id"] = *r.Id
		}

		if r.LifecycleDetails != nil {
			databasePdbConversionHistoryEntry["lifecycle_details"] = *r.LifecycleDetails
		}

		if r.SourceDatabaseId != nil {
			databasePdbConversionHistoryEntry["source_database_id"] = *r.SourceDatabaseId
		}

		databasePdbConversionHistoryEntry["state"] = r.LifecycleState

		databasePdbConversionHistoryEntry["target"] = r.Target

		if r.TargetDatabaseId != nil {
			databasePdbConversionHistoryEntry["target_database_id"] = *r.TargetDatabaseId
		}

		if r.TimeEnded != nil {
			databasePdbConversionHistoryEntry["time_ended"] = r.TimeEnded.String()
		}

		if r.TimeStarted != nil {
			databasePdbConversionHistoryEntry["time_started"] = r.TimeStarted.String()
		}

		resources = append(resources, databasePdbConversionHistoryEntry)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabaseDatabasePdbConversionHistoryEntriesDataSource().Schema["pdb_conversion_history_entries"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("pdb_conversion_history_entries", resources); err != nil {
		return err
	}

	return nil
}
