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

func DatabaseDatabasePdbConversionHistoryEntryDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatabaseDatabasePdbConversionHistoryEntry,
		Schema: map[string]*schema.Schema{
			"database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"pdb_conversion_history_entry_id": {
				Type:     schema.TypeString,
				Required: true,
			},
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
	}
}

func readSingularDatabaseDatabasePdbConversionHistoryEntry(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDatabasePdbConversionHistoryEntryDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseDatabasePdbConversionHistoryEntryDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.GetPdbConversionHistoryEntryResponse
}

func (s *DatabaseDatabasePdbConversionHistoryEntryDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseDatabasePdbConversionHistoryEntryDataSourceCrud) Get() error {
	request := oci_database.GetPdbConversionHistoryEntryRequest{}

	if databaseId, ok := s.D.GetOkExists("database_id"); ok {
		tmp := databaseId.(string)
		request.DatabaseId = &tmp
	}

	if pdbConversionHistoryEntryId, ok := s.D.GetOkExists("pdb_conversion_history_entry_id"); ok {
		tmp := pdbConversionHistoryEntryId.(string)
		request.PdbConversionHistoryEntryId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.GetPdbConversionHistoryEntry(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseDatabasePdbConversionHistoryEntryDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("action", s.Res.Action)

	if s.Res.AdditionalCdbParams != nil {
		s.D.Set("additional_cdb_params", *s.Res.AdditionalCdbParams)
	}

	if s.Res.CdbName != nil {
		s.D.Set("cdb_name", *s.Res.CdbName)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.SourceDatabaseId != nil {
		s.D.Set("source_database_id", *s.Res.SourceDatabaseId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("target", s.Res.Target)

	if s.Res.TargetDatabaseId != nil {
		s.D.Set("target_database_id", *s.Res.TargetDatabaseId)
	}

	if s.Res.TimeEnded != nil {
		s.D.Set("time_ended", s.Res.TimeEnded.String())
	}

	if s.Res.TimeStarted != nil {
		s.D.Set("time_started", s.Res.TimeStarted.String())
	}

	return nil
}
