// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"
)

func DatabaseDbSystemsUpgradeHistoryEntryDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatabaseDbSystemsUpgradeHistoryEntry,
		Schema: map[string]*schema.Schema{
			"db_system_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"upgrade_history_entry_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"action": {
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
	}
}

func readSingularDatabaseDbSystemsUpgradeHistoryEntry(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDbSystemsUpgradeHistoryEntryDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseDbSystemsUpgradeHistoryEntryDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.GetDbSystemUpgradeHistoryEntryResponse
}

func (s *DatabaseDbSystemsUpgradeHistoryEntryDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseDbSystemsUpgradeHistoryEntryDataSourceCrud) Get() error {
	request := oci_database.GetDbSystemUpgradeHistoryEntryRequest{}

	if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
		tmp := dbSystemId.(string)
		request.DbSystemId = &tmp
	}

	if upgradeHistoryEntryId, ok := s.D.GetOkExists("upgrade_history_entry_id"); ok {
		tmp := upgradeHistoryEntryId.(string)
		request.UpgradeHistoryEntryId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.GetDbSystemUpgradeHistoryEntry(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseDbSystemsUpgradeHistoryEntryDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("action", s.Res.Action)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.NewGiVersion != nil {
		s.D.Set("new_gi_version", *s.Res.NewGiVersion)
	}

	if s.Res.NewOsVersion != nil {
		s.D.Set("new_os_version", *s.Res.NewOsVersion)
	}

	if s.Res.OldGiVersion != nil {
		s.D.Set("old_gi_version", *s.Res.OldGiVersion)
	}

	if s.Res.OldOsVersion != nil {
		s.D.Set("old_os_version", *s.Res.OldOsVersion)
	}

	if s.Res.SnapshotRetentionPeriodInDays != nil {
		s.D.Set("snapshot_retention_period_in_days", *s.Res.SnapshotRetentionPeriodInDays)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeEnded != nil {
		s.D.Set("time_ended", s.Res.TimeEnded.String())
	}

	if s.Res.TimeStarted != nil {
		s.D.Set("time_started", s.Res.TimeStarted.String())
	}

	return nil
}
