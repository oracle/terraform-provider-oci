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

func DatabaseDatabaseUpgradeHistoryEntryDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatabaseDatabaseUpgradeHistoryEntry,
		Schema: map[string]*schema.Schema{
			"database_id": {
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
	}
}

func readSingularDatabaseDatabaseUpgradeHistoryEntry(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDatabaseUpgradeHistoryEntryDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseDatabaseUpgradeHistoryEntryDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.GetDatabaseUpgradeHistoryEntryResponse
}

func (s *DatabaseDatabaseUpgradeHistoryEntryDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseDatabaseUpgradeHistoryEntryDataSourceCrud) Get() error {
	request := oci_database.GetDatabaseUpgradeHistoryEntryRequest{}

	if databaseId, ok := s.D.GetOkExists("database_id"); ok {
		tmp := databaseId.(string)
		request.DatabaseId = &tmp
	}

	if upgradeHistoryEntryId, ok := s.D.GetOkExists("upgrade_history_entry_id"); ok {
		tmp := upgradeHistoryEntryId.(string)
		request.UpgradeHistoryEntryId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.GetDatabaseUpgradeHistoryEntry(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseDatabaseUpgradeHistoryEntryDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("action", s.Res.Action)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Options != nil {
		s.D.Set("options", *s.Res.Options)
	}

	s.D.Set("source", s.Res.Source)

	if s.Res.SourceDbHomeId != nil {
		s.D.Set("source_db_home_id", *s.Res.SourceDbHomeId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TargetDBVersion != nil {
		s.D.Set("target_db_version", *s.Res.TargetDBVersion)
	}

	if s.Res.TargetDatabaseSoftwareImageId != nil {
		s.D.Set("target_database_software_image_id", *s.Res.TargetDatabaseSoftwareImageId)
	}

	if s.Res.TargetDbHomeId != nil {
		s.D.Set("target_db_home_id", *s.Res.TargetDbHomeId)
	}

	if s.Res.TimeEnded != nil {
		s.D.Set("time_ended", s.Res.TimeEnded.String())
	}

	if s.Res.TimeStarted != nil {
		s.D.Set("time_started", s.Res.TimeStarted.String())
	}

	return nil
}
