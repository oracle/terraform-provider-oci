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

func DatabaseDbSystemsUpgradeHistoryEntriesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseDbSystemsUpgradeHistoryEntries,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"db_system_id": {
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
			"db_system_upgrade_history_entries": {
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
						"new_gi_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"new_os_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"old_gi_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"old_os_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"snapshot_retention_period_in_days": {
							Type:     schema.TypeInt,
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

func readDatabaseDbSystemsUpgradeHistoryEntries(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDbSystemsUpgradeHistoryEntriesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseDbSystemsUpgradeHistoryEntriesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListDbSystemUpgradeHistoryEntriesResponse
}

func (s *DatabaseDbSystemsUpgradeHistoryEntriesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseDbSystemsUpgradeHistoryEntriesDataSourceCrud) Get() error {
	request := oci_database.ListDbSystemUpgradeHistoryEntriesRequest{}

	if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
		tmp := dbSystemId.(string)
		request.DbSystemId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_database.DbSystemUpgradeHistoryEntrySummaryLifecycleStateEnum(state.(string))
	}

	if upgradeAction, ok := s.D.GetOkExists("upgrade_action"); ok {
		request.UpgradeAction = oci_database.DbSystemUpgradeHistoryEntrySummaryActionEnum(upgradeAction.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListDbSystemUpgradeHistoryEntries(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDbSystemUpgradeHistoryEntries(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseDbSystemsUpgradeHistoryEntriesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseDbSystemsUpgradeHistoryEntriesDataSource-", DatabaseDbSystemsUpgradeHistoryEntriesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		dbSystemsUpgradeHistoryEntry := map[string]interface{}{}

		dbSystemsUpgradeHistoryEntry["action"] = r.Action

		if r.Id != nil {
			dbSystemsUpgradeHistoryEntry["id"] = *r.Id
		}

		if r.LifecycleDetails != nil {
			dbSystemsUpgradeHistoryEntry["lifecycle_details"] = *r.LifecycleDetails
		}

		if r.NewGiVersion != nil {
			dbSystemsUpgradeHistoryEntry["new_gi_version"] = *r.NewGiVersion
		}

		if r.NewOsVersion != nil {
			dbSystemsUpgradeHistoryEntry["new_os_version"] = *r.NewOsVersion
		}

		if r.OldGiVersion != nil {
			dbSystemsUpgradeHistoryEntry["old_gi_version"] = *r.OldGiVersion
		}

		if r.OldOsVersion != nil {
			dbSystemsUpgradeHistoryEntry["old_os_version"] = *r.OldOsVersion
		}

		if r.SnapshotRetentionPeriodInDays != nil {
			dbSystemsUpgradeHistoryEntry["snapshot_retention_period_in_days"] = *r.SnapshotRetentionPeriodInDays
		}

		dbSystemsUpgradeHistoryEntry["state"] = r.LifecycleState

		if r.TimeEnded != nil {
			dbSystemsUpgradeHistoryEntry["time_ended"] = r.TimeEnded.String()
		}

		if r.TimeStarted != nil {
			dbSystemsUpgradeHistoryEntry["time_started"] = r.TimeStarted.String()
		}

		resources = append(resources, dbSystemsUpgradeHistoryEntry)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabaseDbSystemsUpgradeHistoryEntriesDataSource().Schema["db_system_upgrade_history_entries"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("db_system_upgrade_history_entries", resources); err != nil {
		return err
	}

	return nil
}
