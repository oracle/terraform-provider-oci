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

func DatabaseDatabaseUpgradeHistoryEntriesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseDatabaseUpgradeHistoryEntries,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"upgrade_action": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"database_upgrade_history_entries": {
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
						"options": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"source": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"source_db_home_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"target_db_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"target_database_software_image_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"target_db_home_id": {
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

func readDatabaseDatabaseUpgradeHistoryEntries(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDatabaseUpgradeHistoryEntriesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseDatabaseUpgradeHistoryEntriesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListDatabaseUpgradeHistoryEntriesResponse
}

func (s *DatabaseDatabaseUpgradeHistoryEntriesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseDatabaseUpgradeHistoryEntriesDataSourceCrud) Get() error {
	request := oci_database.ListDatabaseUpgradeHistoryEntriesRequest{}

	if databaseId, ok := s.D.GetOkExists("database_id"); ok {
		tmp := databaseId.(string)
		request.DatabaseId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_database.DatabaseUpgradeHistoryEntrySummaryLifecycleStateEnum(state.(string))
	}

	if upgradeAction, ok := s.D.GetOkExists("upgrade_action"); ok {
		request.UpgradeAction = oci_database.DatabaseUpgradeHistoryEntrySummaryActionEnum(upgradeAction.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListDatabaseUpgradeHistoryEntries(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDatabaseUpgradeHistoryEntries(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseDatabaseUpgradeHistoryEntriesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseDatabaseUpgradeHistoryEntriesDataSource-", DatabaseDatabaseUpgradeHistoryEntriesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		databaseUpgradeHistoryEntry := map[string]interface{}{}

		databaseUpgradeHistoryEntry["action"] = r.Action

		if r.Id != nil {
			databaseUpgradeHistoryEntry["id"] = *r.Id
		}

		if r.LifecycleDetails != nil {
			databaseUpgradeHistoryEntry["lifecycle_details"] = *r.LifecycleDetails
		}

		if r.Options != nil {
			databaseUpgradeHistoryEntry["options"] = *r.Options
		}

		databaseUpgradeHistoryEntry["source"] = r.Source

		if r.SourceDbHomeId != nil {
			databaseUpgradeHistoryEntry["source_db_home_id"] = *r.SourceDbHomeId
		}

		databaseUpgradeHistoryEntry["state"] = r.LifecycleState

		if r.TargetDBVersion != nil {
			databaseUpgradeHistoryEntry["target_db_version"] = *r.TargetDBVersion
		}

		if r.TargetDatabaseSoftwareImageId != nil {
			databaseUpgradeHistoryEntry["target_database_software_image_id"] = *r.TargetDatabaseSoftwareImageId
		}

		if r.TargetDbHomeId != nil {
			databaseUpgradeHistoryEntry["target_db_home_id"] = *r.TargetDbHomeId
		}

		if r.TimeEnded != nil {
			databaseUpgradeHistoryEntry["time_ended"] = r.TimeEnded.String()
		}

		if r.TimeStarted != nil {
			databaseUpgradeHistoryEntry["time_started"] = r.TimeStarted.String()
		}

		resources = append(resources, databaseUpgradeHistoryEntry)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabaseDatabaseUpgradeHistoryEntriesDataSource().Schema["database_upgrade_history_entries"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("database_upgrade_history_entries", resources); err != nil {
		return err
	}

	return nil
}
